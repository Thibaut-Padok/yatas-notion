package yatas

import (
	"context"
	"log"

	"github.com/jomei/notionapi"

	"github.com/Thibaut-Padok/yatas-notion/notionAPI"
)

func IsDatabaseExistInPage(client *notionAPI.Client, account *notionAPI.PluginConfig) bool {
	// Getting Default client
	defaultClient := client.ClientV1.JomeiClient

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
				if IsDatabase(client, notionapi.DatabaseID(block.GetID())) {
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

func IsDatabase(client *notionAPI.Client, databaseID notionapi.DatabaseID) bool {
	// Getting Default client
	defaultClient := client.ClientV1.JomeiClient

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
