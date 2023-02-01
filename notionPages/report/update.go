package report

import (
	"context"
	"log"

	"github.com/jomei/notionapi"
	"github.com/stangirard/yatas/plugins/commons"
)

func addCategoriesTable(client *notionapi.Client, blockID notionapi.BlockID, test commons.Tests) {
	updateRequest := updateRequestWithCategories(test)
	blocks, err := client.Block.AppendChildren(context.Background(), blockID, &updateRequest)
	if err != nil {
		log.Printf("Error while triing to add categories table ... %v", err)
	} else {
		log.Printf("New blocks added ... %v", blocks)
	}
}

func addTitleTable(client *notionapi.Client, blockID notionapi.BlockID, test commons.Tests) {
	updateRequest := updateRequestWithTitle(test)
	blocks, err := client.Block.AppendChildren(context.Background(), blockID, &updateRequest)
	if err != nil {
		log.Printf("Error while triing to add a title ... %v", err)
	} else {
		log.Printf("New Title added ... %v", blocks)
	}
}
