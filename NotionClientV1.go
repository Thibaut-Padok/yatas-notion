package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/jomei/notionapi"
)

const (
	apiURL        = "https://api.notion.com"
	apiVersion    = "v1"
	notionVersion = "2022-06-28"
)

type NotionClientV1 struct {
	JomeiClient *notionapi.Client

	httpClient    *http.Client
	baseUrl       *url.URL
	apiVersion    string
	notionVersion string

	Token notionapi.Token

	Database DatabaseV1Service
	Block    notionapi.BlockService
	Page     notionapi.PageService
	User     notionapi.UserService
	Search   notionapi.SearchService
	Comment  notionapi.CommentService
}

func NewClientV1(client *notionapi.Client, token notionapi.Token) *NotionClientV1 {
	u, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}

	cli := &NotionClientV1{
		JomeiClient: client,

		httpClient:    http.DefaultClient,
		Token:         token,
		baseUrl:       u,
		apiVersion:    apiVersion,
		notionVersion: notionVersion,

		Block:   client.Block,
		Page:    client.Page,
		User:    client.User,
		Search:  client.Search,
		Comment: client.Comment,
	}
	cli.Database = &DatabaseV1Client{Client: cli}

	return cli
}

func (c *NotionClientV1) request(ctx context.Context, method string, urlStr string, queryParams map[string]string, requestBody interface{}) (*http.Response, error) {
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s/%s", c.apiVersion, urlStr))
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if requestBody != nil && !reflect.ValueOf(requestBody).IsNil() {
		body, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(body)
	}

	if len(queryParams) > 0 {
		q := u.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token.String()))
	req.Header.Add("Notion-Version", c.notionVersion)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var apiErr notionapi.Error
		err = json.NewDecoder(res.Body).Decode(&apiErr)
		if err != nil {
			return nil, err
		}

		return nil, &apiErr
	}

	return res, nil
}
