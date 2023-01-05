package main

import "github.com/jomei/notionapi"

type NotionClient struct {
	NotionClientV1
	NotionClientV3
}

func NewNotionClient(account *NotionAccount) NotionClient {
	// Create a good type token
	// Create the default notionapi/v1 Client
	token := notionapi.Token(account.Token)
	client := notionapi.NewClient(token)

	// Create custom Clients
	clientV1 := NewClientV1(client, token)
	clientV3 := NewClientV3(account.AuthToken, account.PageID)
	notionClient := NotionClient{
		NotionClientV1: *clientV1,
		NotionClientV3: *clientV3,
	}
	return notionClient
}

func (client *NotionClient) GetTableViewType(databaseID string) (string, string, bool) {
	clientV3 := client.NotionClientV3
	return clientV3.GetTableViewType(databaseID)
}

func (client *NotionClient) UpdateTableViewList(viewID, desiredType string) error {
	clientV3 := client.NotionClientV3
	err := clientV3.UpdateTableViewList(viewID, desiredType)
	return err
}

func (client *NotionClient) LockPage(pageID string) error {
	clientV3 := client.NotionClientV3
	err := clientV3.LockPage(pageID)
	return err
}
