package main

import (
	"context"
	"log"

	"github.com/stangirard/yatas/plugins/commons"
)

func CreateNotionReport(tests []commons.Tests, account NotionAccount, client *NotionClient) {
	// client := notionapi.NewClient(notionapi.Token(account.Token))
	// cli := NewLocalClient(client, notionapi.Token(account.Token))
	// Get the different clients
	clientV1 := client.NotionClientV1
	defaultClient := clientV1.JomeiClient

	// Create a new page corresponding to the yatas report
	newYatasReport := createReportPageRequest(defaultClient, account.DatabaseID)
	page, err := client.Page.Create(context.Background(), &newYatasReport)
	if err != nil {
		log.Printf("Error during the creation of the Yamas page ... %v", err)
	} else {

		// Create the database inside the yatas report page
		dbCreate := createInlineDatabase(page.ID.String())
		db, err := clientV1.Database.Create(context.Background(), &dbCreate)
		if err != nil {
			log.Printf("Error during the creation of the database ... %v", err)
		} else {

			// Create a new page for each test
			for _, test := range tests {
				for _, check := range test.Checks {
					createPageCheck(client, check, db.ID.String())
				}
			}
		}
		// Try to lock page if notionapi/v3 available
		client.LockPage(page.ID.String())
	}
}

func createPageCheck(client *NotionClient, check commons.Check, databaseId string) {
	clientV1 := client.NotionClientV1
	pageCreateRequest := createPageCheckRequest(check, databaseId)
	page, err := clientV1.Page.Create(context.Background(), &pageCreateRequest)
	if err != nil {
		log.Printf("Error ... %v", err)
	} else {
		log.Printf("Check page created: %v", page.URL)
	}

	// Try to lock page if notionapi/v3 available
	client.LockPage(string(page.ID))
}
