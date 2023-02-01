package report

import (
	"github.com/jomei/notionapi"

	"github.com/Thibaut-Padok/yatas-notion/notionAPI/notionV1"
)

func createInlineDatabase(pageID string) notionV1.DatabaseCreateRequest {
	title := notionapi.Text{
		Content: "All checks",
	}
	database := notionV1.DatabaseCreateRequest{
		Parent: notionapi.Parent{PageID: notionapi.PageID(pageID)},
		Title: []notionapi.RichText{
			{
				Text: &title,
			},
		},
		Properties: notionapi.PropertyConfigs{
			"Id": notionapi.TitlePropertyConfig{
				Type:  "title",
				Title: struct{}{},
			},
			"Status": notionapi.CheckboxPropertyConfig{
				Type:     "checkbox",
				Checkbox: struct{}{},
			},
			"Type": notionapi.SelectPropertyConfig{
				Type:   "select",
				Select: notionapi.Select{Options: []notionapi.Option{}},
			},
			"Description": notionapi.RichTextPropertyConfig{
				Type:     "rich_text",
				RichText: struct{}{},
			},
		},
		IsInline: true,
	}
	return database
}
