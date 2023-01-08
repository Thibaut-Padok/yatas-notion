package main

import (
	"errors"

	"github.com/jomei/notionapi"
)

type NotionClient struct {
	*NotionClientV1
	*NotionClientV3
}

func NewNotionClient(account *NotionAccount) NotionClient {
	// Create a good type token
	// Create the default notionapi/v1 Client
	token := notionapi.Token(account.Token)
	client := notionapi.NewClient(token)

	// Create custom Clients
	clientV1 := NewClientV1(client, token)
	var notionClient NotionClient
	if account.AuthToken != "" {
		clientV3 := NewClientV3(account.AuthToken, account.PageID)
		notionClient = NotionClient{
			NotionClientV1: clientV1,
			NotionClientV3: clientV3,
		}

	} else {
		notionClient = NotionClient{
			NotionClientV1: clientV1,
		}
	}
	return notionClient
}

func (client *NotionClient) GetTableViewType(databaseID string) (string, string, bool) {
	clientV3 := client.NotionClientV3
	if clientV3 == nil {
		return "", "", false
	}
	return clientV3.GetTableViewType(databaseID)
}

func (client *NotionClient) GetTableViewProperties(databaseID string) []string {
	clientV3 := client.NotionClientV3
	var properties []string
	if clientV3 == nil {
		return properties
	}
	return clientV3.GetTableViewProperties(databaseID)
}

func (client *NotionClient) UpdateTableViewList(viewID, desiredType string) error {
	clientV3 := client.NotionClientV3
	if clientV3 == nil {
		err := errors.New("NotionClientV3 does not exist, please provide authToken.")
		return err
	}
	err := clientV3.UpdateTableViewList(viewID, desiredType)
	return err
}

func (client *NotionClient) LockPage(pageID string) error {
	clientV3 := client.NotionClientV3
	if clientV3 == nil {
		err := errors.New("NotionClientV3 does not exist, please provide authToken.")
		return err
	}
	err := clientV3.LockPage(pageID)
	return err
}

func (client *NotionClient) ShowProperties(viewID string, properties []string) error {
	clientV3 := client.NotionClientV3
	if clientV3 == nil {
		err := errors.New("NotionClientV3 does not exist, please provide authToken.")
		return err
	}
	err := clientV3.ShowProperties(viewID, properties)
	return err
}

func (client *NotionClient) ShowAllProperties(viewID, databaseID string) error {
	clientV3 := client.NotionClientV3
	if clientV3 == nil {
		err := errors.New("NotionClientV3 does not exist, please provide authToken.")
		return err
	}
	properties := clientV3.GetTableViewProperties(databaseID)
	err := clientV3.ShowProperties(viewID, properties)
	return err
}
