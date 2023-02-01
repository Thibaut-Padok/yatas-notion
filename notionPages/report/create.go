package report

import (
	"context"
	"log"

	"github.com/jomei/notionapi"
	"github.com/stangirard/yatas/plugins/commons"

	"github.com/Thibaut-Padok/yatas-notion/notionAPI"
	checkPage "github.com/Thibaut-Padok/yatas-notion/notionPages/check"
)

func Create(tests []commons.Tests, account notionAPI.PluginConfig, client *notionAPI.Client) {
	// Get the different clients
	clientV1 := client.ClientV1
	defaultClient := clientV1.JomeiClient

	// Create a new page corresponding to the yatas report
	newYatasReport := createPageRequest(defaultClient, account.DatabaseID)
	page, err := client.ClientV1.Page.Create(context.Background(), &newYatasReport)
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
					checkPage.CreatePage(client, check, db.ID.String())
				}
				// Create a new Categories table for each 'Test'
				addCategoriesTable(defaultClient, notionapi.BlockID(page.ID), test)
			}
		}
		// Try to lock page if notionapi/v3 available
		client.LockPage(page.ID.String())
	}
}
