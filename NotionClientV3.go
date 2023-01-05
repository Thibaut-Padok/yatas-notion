package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	kjk "github.com/kjk/notionapi"
)

const (
	notionHost = "https://www.notion.so"
	userAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.3"
	acceptLang = "en-GB,en-US;q=0.9,en;q=0.8"
)

type NotionClientV3 struct {
	AuthToken       string
	HTTPClient      *http.Client
	MinRequestDelay time.Duration
	lastRequestTime time.Time

	httpPostOverride func(uri string, body []byte) ([]byte, error)
	SpaceID          string

	KjkClient *kjk.Client
}

func NewClientV3(clientToken, pageID string) *NotionClientV3 {
	// Create client with token
	client := &NotionClientV3{
		AuthToken: clientToken,
	}
	// Create kjk/notionapi client from token
	kjkClient := &kjk.Client{
		AuthToken: clientToken,
	}
	client.KjkClient = kjkClient

	// Get Client workspace for the NotionClientV3
	page, err := client.KjkClient.DownloadPage(pageID)
	if err != nil {
		log.Printf("Error while triing to get spaceId: %v", err)
	} else {
		client.SpaceID = page.Root().SpaceID
	}
	return client
}

func (c *NotionClientV3) getHTTPClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	httpNotionClientV3 := *http.DefaultClient
	httpNotionClientV3.Timeout = time.Second * 30
	return &httpNotionClientV3
}

func (c *NotionClientV3) doPost(uri string, body []byte) ([]byte, error) {
	if c.httpPostOverride != nil {
		return c.httpPostOverride(uri, body)
	}
	return c.doPostInternal(uri, body)
}

func (c *NotionClientV3) doPostInternal(uri string, body []byte) ([]byte, error) {
	nRepeats := 0
	timeouts := []time.Duration{time.Second * 3, time.Second * 5, time.Second * 10}
repeatRequest:
	br := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", uri, br)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept-Language", acceptLang)

	if c.AuthToken != "" {
		req.Header.Set("cookie", fmt.Sprintf("token_v2=%v", c.AuthToken))
	}
	var rsp *http.Response

	httpNotionClientV3 := c.getHTTPClient()

	rsp, err = httpNotionClientV3.Do(req)

	if err != nil {
		log.Printf("httpNotionClientV3.Do() failed with %s\n", err)
		return nil, err
	}

	if rsp.StatusCode == http.StatusTooManyRequests {
		if nRepeats < 3 {
			log.Printf("retrying '%s' because httpNotionClientV3.Do() returned %d (%s)\n", uri, rsp.StatusCode, rsp.Status)
			time.Sleep(timeouts[nRepeats])
			nRepeats++
			goto repeatRequest
		}
	}

	if rsp.StatusCode != 200 {
		d, _ := ioutil.ReadAll(rsp.Body)
		log.Printf("Error: status code %s\nBody:\n%s\n", rsp.Status, PrettyPrintJS(d))
		return nil, fmt.Errorf("http.Post('%s') returned non-200 status code of %d", uri, rsp.StatusCode)
	}
	d, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("Error: ioutil.ReadAll() failed with %s\n", err)
		return nil, err
	}
	return d, nil
}

func (c *NotionClientV3) doNotionAPI(apiURL string, requestData interface{}, result interface{}, rawJSON *map[string]interface{}) error {
	var body []byte
	var err error
	if requestData != nil {
		body, err = jsonit.MarshalIndent(requestData, "", "  ")
		if err != nil {
			return err
		}
	}
	uri := notionHost + apiURL

	d, err := c.doPost(uri, body)
	if err != nil {
		return err
	}

	err = jsonit.Unmarshal(d, result)
	if err != nil {
		log.Printf("Error: json.Unmarshal() failed with %s\n. Body:\n%s\n", err, string(d))
		return err
	}
	if rawJSON != nil {
		err = jsonit.Unmarshal(d, rawJSON)
	}
	return err
}

func (client *NotionClientV3) saveTransactions(req UpdateRequest) (*UpdateResponse, error) {
	var rsp UpdateResponse
	var err error
	apiURL := "/api/v3/saveTransactions"
	if err = client.doNotionAPI(apiURL, req, &rsp, &rsp.RawJSON); err != nil {
		return nil, err
	}
	return &rsp, nil
}

func (client *NotionClientV3) GetTableViewType(databaseID string) (string, string, bool) {
	kjkClient := client.KjkClient
	page, err := kjkClient.DownloadPage(databaseID)
	if err != nil {
		log.Printf("Error while getting Database TableView: %v", err)
	} else {
		if len(page.TableViews) > 0 {
			collectionView := page.TableViews[0].CollectionView
			return collectionView.ID, collectionView.Type, true
		}
	}
	return "", "", false
}

func (client *NotionClientV3) UpdateTableViewList(viewID, desiredType string) error {

	req := TableViewTypeUpdateRequest(client.SpaceID, viewID, desiredType)

	_, err := client.saveTransactions(req)
	if err != nil {
		return err
	}
	return nil
}

func (client *NotionClientV3) LockPage(pageID string) error {
	req := LockPageUpdateRequest(client.SpaceID, pageID)
	_, err := client.saveTransactions(req)
	if err != nil {
		return err
	}
	return nil
}
