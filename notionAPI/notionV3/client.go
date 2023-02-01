package notionV3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Thibaut-Padok/yatas-notion/utils"
	kjk "github.com/kjk/notionapi"
)

const (
	notionHost = "https://www.notion.so"
	userAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.3"
	acceptLang = "en-GB,en-US;q=0.9,en;q=0.8"
)

type Client struct {
	AuthToken       string
	HTTPClient      *http.Client
	MinRequestDelay time.Duration
	lastRequestTime time.Time

	httpPostOverride func(uri string, body []byte) ([]byte, error)
	SpaceID          string

	KjkClient *kjk.Client
}

func NewClient(clientToken, pageID string) *Client {
	// Create client with token
	client := &Client{
		AuthToken: clientToken,
	}
	// Create kjk/notionapi client from token
	kjkClient := &kjk.Client{
		AuthToken: clientToken,
	}
	client.KjkClient = kjkClient

	// Get Client workspace for the Client
	page, err := client.KjkClient.DownloadPage(pageID)
	if err != nil {
		log.Printf("Error while triing to get spaceId: %v", err)
	} else {
		client.SpaceID = page.Root().SpaceID
	}
	return client
}

func (c *Client) getHTTPClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	httpClient := *http.DefaultClient
	httpClient.Timeout = time.Second * 30
	return &httpClient
}

func (c *Client) doPost(uri string, body []byte) ([]byte, error) {
	if c.httpPostOverride != nil {
		return c.httpPostOverride(uri, body)
	}
	return c.doPostInternal(uri, body)
}

func (c *Client) doPostInternal(uri string, body []byte) ([]byte, error) {
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

	httpClient := c.getHTTPClient()

	rsp, err = httpClient.Do(req)

	if err != nil {
		log.Printf("httpClient.Do() failed with %s\n", err)
		return nil, err
	}

	if rsp.StatusCode == http.StatusTooManyRequests {
		if nRepeats < 3 {
			log.Printf("retrying '%s' because httpClient.Do() returned %d (%s)\n", uri, rsp.StatusCode, rsp.Status)
			time.Sleep(timeouts[nRepeats])
			nRepeats++
			goto repeatRequest
		}
	}

	if rsp.StatusCode != 200 {
		d, _ := ioutil.ReadAll(rsp.Body)
		log.Printf("Error: status code %s\nBody:\n%s\n", rsp.Status, utils.PrettyPrintJS(d))
		return nil, fmt.Errorf("http.Post('%s') returned non-200 status code of %d", uri, rsp.StatusCode)
	}
	d, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("Error: ioutil.ReadAll() failed with %s\n", err)
		return nil, err
	}
	return d, nil
}

func (c *Client) doNotionAPI(apiURL string, requestData interface{}, result interface{}, rawJSON *map[string]interface{}) error {
	var body []byte
	var err error
	if requestData != nil {
		body, err = utils.Jsonit.MarshalIndent(requestData, "", "  ")
		if err != nil {
			return err
		}
	}
	uri := notionHost + apiURL

	d, err := c.doPost(uri, body)
	if err != nil {
		return err
	}

	err = utils.Jsonit.Unmarshal(d, result)
	if err != nil {
		log.Printf("Error: json.Unmarshal() failed with %s\n. Body:\n%s\n", err, string(d))
		return err
	}
	if rawJSON != nil {
		err = utils.Jsonit.Unmarshal(d, rawJSON)
	}
	return err
}

func (client *Client) saveTransactions(req UpdateRequest) (*UpdateResponse, error) {
	var rsp UpdateResponse
	var err error
	apiURL := "/api/v3/saveTransactions"
	if err = client.doNotionAPI(apiURL, req, &rsp, &rsp.RawJSON); err != nil {
		return nil, err
	}
	return &rsp, nil
}

func (client *Client) GetTableViewType(databaseID string) (string, string, bool) {
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

func (client *Client) GetTableViewProperties(databaseID string) []string {
	kjkClient := client.KjkClient
	page, err := kjkClient.DownloadPage(databaseID)
	var properties []string
	if err != nil {
		log.Printf("Error while getting Database TableView: %v", err)
	} else {
		if len(page.TableViews) > 0 {
			for key := range page.TableViews[0].Collection.Schema {
				properties = append(properties, key)
			}
		}
	}
	return properties
}

func (client *Client) UpdateTableViewList(viewID, desiredType string) error {

	req := TableViewTypeUpdateRequest(client.SpaceID, viewID, desiredType)

	_, err := client.saveTransactions(req)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) LockPage(pageID string) error {
	req := LockPageUpdateRequest(client.SpaceID, pageID)
	_, err := client.saveTransactions(req)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) ShowProperties(viewID string, properties []string) error {
	req := ShowPropertiesUpdateRequest(client.SpaceID, viewID, properties)
	_, err := client.saveTransactions(req)
	if err != nil {
		return err
	}
	return nil
}
