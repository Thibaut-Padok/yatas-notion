package check

import (
	"context"
	"log"
	"strings"

	"github.com/Thibaut-Padok/yatas-notion/notionAPI"
	"github.com/jomei/notionapi"
	"github.com/stangirard/yatas/plugins/commons"
)

func CreatePage(client *notionAPI.Client, check commons.Check, databaseId string) {
	clientV1 := client.ClientV1
	pageCreateRequest := createPageRequest(check, databaseId)
	page, err := clientV1.Page.Create(context.Background(), &pageCreateRequest)
	if err != nil {
		log.Printf("Error ... %v", err)
	} else {
		log.Printf("Check page created: %v", page.URL)
	}

	// Try to lock page if notionapi/v3 available
	client.LockPage(string(page.ID))
}

func createPageRequest(check commons.Check, databaseId string) notionapi.PageCreateRequest {
	//calculate status boolean
	status := true
	if check.Status == "FAIL" {
		status = false
	}
	notionTitle := notionapi.Text{
		Content: check.Id,
	}
	notionDecription := notionapi.Text{
		Content: check.Description,
	}
	emo := notionapi.Emoji("🥷")
	checkType := strings.Split(check.Id, "_")

	pageCreateRequest := notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: notionapi.DatabaseID(databaseId),
		},
		Properties: notionapi.Properties{
			"Id": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: &notionTitle},
				},
			},
			"Type": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: checkType[1],
				},
			},
			"Status": notionapi.CheckboxProperty{
				Checkbox: status,
			},
			"Description": notionapi.RichTextProperty{
				RichText: []notionapi.RichText{
					{Text: &notionDecription},
				},
			},
		},
		Icon: &notionapi.Icon{
			Type:  `emoji`,
			Emoji: &emo,
		},
		Children: []notionapi.Block{
			createTable(check),
		},
	}
	return pageCreateRequest
}
