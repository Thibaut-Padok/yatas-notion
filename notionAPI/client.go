package notionAPI

import (
	"errors"

	"github.com/Thibaut-Padok/yatas-notion/notionAPI/notionV1"
	"github.com/Thibaut-Padok/yatas-notion/notionAPI/notionV3"
	"github.com/jomei/notionapi"
)

type Client struct {
	ClientV1 *notionV1.Client
	ClientV3 *notionV3.Client
}

func NewClient(account *PluginConfig) Client {
	// Create a good type token
	// Create the default notionapi/v1 Client
	token := notionapi.Token(account.Token)
	baseClient := notionapi.NewClient(token)

	// Create custom Clients
	clientV1 := notionV1.NewClient(baseClient, token)
	var client Client
	if account.AuthToken != "" {
		clientV3 := notionV3.NewClient(account.AuthToken, account.PageID)
		client = Client{
			ClientV1: clientV1,
			ClientV3: clientV3,
		}

	} else {
		client = Client{
			ClientV1: clientV1,
		}
	}
	return client
}

func (client *Client) GetTableViewType(databaseID string) (string, string, bool) {
	clientV3 := client.ClientV3
	if clientV3 == nil {
		return "", "", false
	}
	return clientV3.GetTableViewType(databaseID)
}

func (client *Client) GetTableViewProperties(databaseID string) []string {
	clientV3 := client.ClientV3
	var properties []string
	if clientV3 == nil {
		return properties
	}
	return clientV3.GetTableViewProperties(databaseID)
}

func (client *Client) UpdateTableViewList(viewID, desiredType string) error {
	clientV3 := client.ClientV3
	if clientV3 == nil {
		err := errors.New("ClientV3 does not exist, please provide authToken.")
		return err
	}
	err := clientV3.UpdateTableViewList(viewID, desiredType)
	return err
}

func (client *Client) LockPage(pageID string) error {
	clientV3 := client.ClientV3
	if clientV3 == nil {
		err := errors.New("ClientV3 does not exist, please provide authToken.")
		return err
	}
	err := clientV3.LockPage(pageID)
	return err
}

func (client *Client) ShowProperties(viewID string, properties []string) error {
	clientV3 := client.ClientV3
	if clientV3 == nil {
		err := errors.New("ClientV3 does not exist, please provide authToken.")
		return err
	}
	err := clientV3.ShowProperties(viewID, properties)
	return err
}

func (client *Client) ShowAllProperties(viewID, databaseID string) error {
	clientV3 := client.ClientV3
	if clientV3 == nil {
		err := errors.New("ClientV3 does not exist, please provide authToken.")
		return err
	}
	properties := clientV3.GetTableViewProperties(databaseID)
	err := clientV3.ShowProperties(viewID, properties)
	return err
}
