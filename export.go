package main

import (
	"context"
	"log"

	"github.com/jomei/notionapi"
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
		for _, test := range tests {
			// Add a separator and a title for each 'Test'
			addTitleTable(defaultClient, notionapi.BlockID(page.ID), test)

			// Create a new inline database for each 'Test'
			dbCreate := createInlineDatabase(page.ID.String())
			db, err := clientV1.Database.Create(context.Background(), &dbCreate)
			if err != nil {
				log.Printf("Error during the creation of the database ... %v", err)
			} else {

				for _, check := range test.Checks {
					// Create a new page in the inline database for each 'Check'
					createPageCheck(client, check, db.ID.String())
				}
				// Create a new Categories table for each 'Test'
				addCategoriesTable(defaultClient, notionapi.BlockID(page.ID), test)
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
