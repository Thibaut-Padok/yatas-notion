package yatas

import (
	"context"
	"log"
	"time"

	"github.com/jomei/notionapi"

	"github.com/Thibaut-Padok/yatas-notion/notionAPI"
	"github.com/Thibaut-Padok/yatas-notion/notionAPI/notionV1"
)

func InitDatabase(client *notionAPI.Client, account *notionAPI.PluginConfig) bool {
	// If exists retrieve the Yatas databaseID else create a new inline database.
	if IsDatabaseExistInPage(client, account) {
		log.Printf("[LOADED] DatabaseID = %v", account.DatabaseID)
		return true
	} else {
		// Needs to create Yatas database
		log.Printf("Start database creation")
		return createInlineDatabase(client, account)
	}
}

func createInlineDatabase(client *notionAPI.Client, account *notionAPI.PluginConfig) bool {
	// Get ClientV1 to create inline databases
	clientv1 := client.ClientV1

	// Create a inline database creation request
	inlineDatabase := inlineDatabaseCreateRequest(account.PageID)

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
			client.ShowAllProperties(viewID, account.DatabaseID)
		}
		return true
	}
}

func inlineDatabaseCreateRequest(pageID string) notionV1.DatabaseCreateRequest {
	title := notionapi.Text{
		Content: "Yatas instances",
	}
	emoji := notionapi.Emoji("🧠")
	database := notionV1.DatabaseCreateRequest{
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
			"Created time": notionapi.CreatedTimePropertyConfig{
				Type:        "created_time",
				CreatedTime: struct{}{},
			},
		},
		IsInline: true,
	}
	return database
}
