package main

import (
	"github.com/jomei/notionapi"
	"github.com/stangirard/yatas/plugins/commons"
)

func updatePageRequestWithTitle(test commons.Tests) notionapi.AppendBlockChildrenRequest {
	categoryTitle := notionapi.Text{
		Content: test.Account,
	}
	request := notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{
			notionapi.DividerBlock{
				BasicBlock: notionapi.BasicBlock{
					Type:   notionapi.BlockType("divider"),
					Object: notionapi.ObjectType("block"),
				},
				Divider: notionapi.Divider{},
			},
			notionapi.Heading2Block{
				BasicBlock: notionapi.BasicBlock{
					Type:   notionapi.BlockType("heading_2"),
					Object: notionapi.ObjectType("block"),
				},
				Heading2: notionapi.Heading{
					RichText: []notionapi.RichText{
						{
							Text: &categoryTitle,
						},
					},
				},
			},
		},
	}
	return request
}

func updatePageRequestWithCategories(test commons.Tests) notionapi.AppendBlockChildrenRequest {
	categoryTitle := notionapi.Text{
		Content: "Categories",
	}
	categoriesBlock := createCategoriesBlock(test)
	request := notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{
			notionapi.Heading2Block{
				BasicBlock: notionapi.BasicBlock{
					Type:   notionapi.BlockType("heading_2"),
					Object: notionapi.ObjectType("block"),
				},
				Heading2: notionapi.Heading{
					RichText: []notionapi.RichText{
						{
							Text: &categoryTitle,
						},
					},
				},
			},
			categoriesBlock,
		},
	}
	return request
}
