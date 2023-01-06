package main

import (
	"strings"

	"github.com/jomei/notionapi"
	"github.com/stangirard/yatas/plugins/commons"
)

func createPageCheckRequest(check commons.Check, databaseId string) notionapi.PageCreateRequest {
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
			createCheckTable(check),
		},
	}
	return pageCreateRequest
}
