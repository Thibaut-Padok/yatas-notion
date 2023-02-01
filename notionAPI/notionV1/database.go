package notionV1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jomei/notionapi"
)

type DatabaseService interface {
	Create(ctx context.Context, request *DatabaseCreateRequest) (*notionapi.Database, error)
}

type DatabaseClient struct {
	Client *Client
}

type DatabaseCreateRequest struct {
	Parent     notionapi.Parent          `json:"parent"`
	Title      []notionapi.RichText      `json:"title"`
	Icon       *notionapi.Icon           `json:"icon,omitempty"`
	Properties notionapi.PropertyConfigs `json:"properties"`
	IsInline   bool                      `json:"is_inline"`
}

func (dc *DatabaseClient) Create(ctx context.Context, requestBody *DatabaseCreateRequest) (*notionapi.Database, error) {
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
