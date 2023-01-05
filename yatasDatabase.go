package main

import (
	"context"
	"log"
	"time"

	"github.com/jomei/notionapi"
)

func initYatasDatabase(client *NotionClient, account *NotionAccount) bool {
	// If exists retrieve the Yatas databaseID else create a new inline database.
	if isYatasDatabaseExistInPage(client, account) {
		log.Printf("[LOADED] DatabaseID = %v", account.DatabaseID)
		return true
	} else {
		// Needs to create Yatas database
		log.Printf("Start database creation")
		return createYatasInlineDatabase(client, account)
	}
}

func isYatasDatabaseExistInPage(client *NotionClient, account *NotionAccount) bool {
	// Getting Default client
	defaultClient := client.NotionClientV1.JomeiClient

	// Type casting
	pageID := notionapi.PageID(account.PageID)
	blockID := notionapi.BlockID(account.PageID)
	pagination := notionapi.Pagination{
		PageSize: 10,
	}

	// Getting pages
	_, pageError := defaultClient.Page.Get(context.Background(), pageID)
	blocks, blockError := defaultClient.Block.GetChildren(context.Background(), blockID, &pagination)

	// Is the provided page exist ?
	if pageError != nil {
		log.Printf("Error getting the page : %v  ... %v", account.PageID, pageError)
		return false
	}

	// Is the Yatas database exist in the provided page?
	if blockError != nil {
		log.Printf("Error getting the block : %v  ... %v", blockID, blockError)
		return false
	} else {
		for _, block := range blocks.Results {
			if block.GetType() == "child_database" {
				// Is the Yatas database ?
				if isYatasDatabase(client, notionapi.DatabaseID(block.GetID())) {
					// Update account to re-use databaseID
					log.Printf("Yatas database found !")
					account.DatabaseID = string(block.GetID())
					return true
				}
			}
		}
	}
	return false
}

func isYatasDatabase(client *NotionClient, databaseID notionapi.DatabaseID) bool {
	// Getting Default client
	defaultClient := client.NotionClientV1.JomeiClient

	// Getting the database
	db, dbError := defaultClient.Database.Get(context.Background(), databaseID)

	// Is the database exist ?
	if dbError != nil {
		log.Printf("Error getting the database : %v  ... %v", databaseID, dbError)
	} else {
		log.Printf("Getting database %v successed ! %v", databaseID, *(&(db.Title[0].Text).Content))

		// Is the database have the good name ?
		if *(&(db.Title[0].Text).Content) == "Yatas instances" {
			return true
		}
	}
	return false
}

func createYatasInlineDatabase(client *NotionClient, account *NotionAccount) bool {
	// Get ClientV1 to create inline databases
	clientv1 := client.NotionClientV1

	// Create a inline database creation request
	inlineDatabase := yatasInlineDatabaseCreateRequest(account.PageID)

	// Try to create inline database
	newDatabase, err := clientv1.Database.Create(context.Background(), &inlineDatabase)

	// Is the database created with success ?
	if err != nil {
		log.Printf("Error during the creation of the database ... %v", err)
		return false
	} else {
		account.DatabaseID = string(newDatabase.ID)
		//Wait the db to be created
		time.Sleep(1 * time.Second)

		// Make the database as a list
		viewID, viewType, exist := client.GetTableViewType(account.DatabaseID)
		if exist {
			if viewType != "list" {
				client.UpdateTableViewList(viewID, "list")
			}
		}
		return true
	}
}

func yatasInlineDatabaseCreateRequest(pageID string) DatabaseV1CreateRequest {
	title := notionapi.Text{
		Content: "Yatas instances",
	}
	emoji := notionapi.Emoji("🧠")
	database := DatabaseV1CreateRequest{
		Parent: notionapi.Parent{PageID: notionapi.PageID(pageID)},
		Title: []notionapi.RichText{
			{
				Text: &title,
			},
		},
		Icon: &notionapi.Icon{
			Type:  `emoji`,
			Emoji: &emoji,
		},
		Properties: notionapi.PropertyConfigs{
			"Name": notionapi.TitlePropertyConfig{
				Type:  "title",
				Title: struct{}{},
			},
		},
		IsInline: true,
	}
	return database
}
