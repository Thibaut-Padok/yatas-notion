package main

import (
	"strconv"

	"github.com/jomei/notionapi"
	"github.com/stangirard/yatas/plugins/commons"
)

func CalculatePercent(success int, failure int) string {
	total := success + failure
	if total == 0 {
		return "0"
	}
	return strconv.Itoa((success * 100) / total)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func createCategoriesBlock(test commons.Tests) notionapi.Block {
	// Find the categories
	categories := []string{}
	categoriesSuccess := map[string]int{}
	categoriesFailure := map[string]int{}
	for _, check := range test.Checks {
		for _, category := range check.Categories {
			if !contains(categories, category) {
				categories = append(categories, category)
				categoriesSuccess[category] = 0
				categoriesFailure[category] = 0
			}
			if check.Status == "OK" {
				categoriesSuccess[category]++
			} else {
				categoriesFailure[category]++
			}
		}
	}
	// Calculate the completion scores
	scores := []string{}
	for _, category := range categories {
		scores = append(scores, CalculatePercent(categoriesSuccess[category], categoriesFailure[category])+" %")
	}
	// Write the categories
	block := createCategorieTable(categories, scores)
	return block
}

func createCategorieTable(categories, scores []string) notionapi.Block {
	table := notionapi.TableBlock{
		BasicBlock: notionapi.BasicBlock{
			Type:   notionapi.BlockType("table"),
			Object: notionapi.ObjectType("block"),
		},
		Table: notionapi.Table{
			TableWidth:      2,
			HasColumnHeader: true,
			Children:        createCategorieTableRows(categories, scores),
		},
	}
	return table
}

func createCategorieTableRows(categories, scores []string) []notionapi.Block {
	rows := []notionapi.Block{}
	rows = append(rows, createCategorieTableHeader())
	for key, category := range categories {
		score := scores[key]
		r := createCategorieTableRow(category, score)
		rows = append(rows, r)
	}
	return rows
}

func createCategorieTableHeader() notionapi.TableRowBlock {
	category := notionapi.Text{
		Content: "Category",
	}
	completion := notionapi.Text{
		Content: "Completion",
	}
	row := notionapi.TableRowBlock{
		BasicBlock: notionapi.BasicBlock{
			Type:   notionapi.BlockType("table_row"),
			Object: notionapi.ObjectType("block"),
		},
		TableRow: notionapi.TableRow{
			Cells: [][]notionapi.RichText{
				{{Text: &category}},
				{{Text: &completion}},
			},
		},
	}
	return row
}

func createCategorieTableRow(category, score string) notionapi.TableRowBlock {
	cat := notionapi.Text{
		Content: category,
	}
	completion := notionapi.Text{
		Content: score,
	}
	row := notionapi.TableRowBlock{
		BasicBlock: notionapi.BasicBlock{
			Type:   notionapi.BlockType("table_row"),
			Object: notionapi.ObjectType("block"),
		},
		TableRow: notionapi.TableRow{
			Cells: [][]notionapi.RichText{
				{{Text: &cat}},
				{{Text: &completion}},
			},
		},
	}
	return row
}
