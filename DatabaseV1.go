package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jomei/notionapi"
)

type DatabaseV1Service interface {
	Create(ctx context.Context, request *DatabaseV1CreateRequest) (*notionapi.Database, error)
}

type DatabaseV1Client struct {
	Client *NotionClientV1
}

type DatabaseV1CreateRequest struct {
	Parent     notionapi.Parent          `json:"parent"`
	Title      []notionapi.RichText      `json:"title"`
	Icon       *notionapi.Icon           `json:"icon,omitempty"`
	Properties notionapi.PropertyConfigs `json:"properties"`
	IsInline   bool                      `json:"is_inline"`
}

func (dc *DatabaseV1Client) Create(ctx context.Context, requestBody *DatabaseV1CreateRequest) (*notionapi.Database, error) {
	res, err := dc.Client.request(ctx, http.MethodPost, "databases", nil, requestBody)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("Failed to close body, should never happen")
		}
	}()

	var response notionapi.Database
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
