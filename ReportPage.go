package main

import (
	"github.com/jomei/notionapi"
)

func createReportPageRequest(client *notionapi.Client, db_id string) notionapi.PageCreateRequest {
	notionTitle := notionapi.Text{
		Content: "Yatas report",
	}
	notionPageHeader := notionapi.Text{
		Content: "YATAS !!",
	}
	notionPagetext := notionapi.Text{
		Content: "This is an new YATAS report.",
	}
	pageYatasCreateRequest := notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: notionapi.DatabaseID(db_id),
		},
		Properties: notionapi.Properties{
			"Name": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: &notionTitle},
				},
			},
		},
		Icon: &notionapi.Icon{
			Type: `external`,
			External: &notionapi.FileObject{
				URL: "https://raw.githubusercontent.com/Thibaut-Padok/yatas-notion/main/docs/auditory.png",
			},
		},

		Children: []notionapi.Block{
			notionapi.Heading2Block{
				BasicBlock: notionapi.BasicBlock{
					Type:   notionapi.BlockType("heading_2"),
					Object: notionapi.ObjectType("block"),
				},
				Heading2: notionapi.Heading{
					RichText: []notionapi.RichText{
						{
							Text: &notionPageHeader,
						},
					},
				},
			},
			notionapi.ParagraphBlock{
				BasicBlock: notionapi.BasicBlock{
					Type:   notionapi.BlockType("paragraph"),
					Object: notionapi.ObjectType("block"),
				},
				Paragraph: notionapi.Paragraph{
					RichText: []notionapi.RichText{
						{
							Text: &notionPagetext,
						},
					},
				},
			},
		},
	}
	return pageYatasCreateRequest
}
