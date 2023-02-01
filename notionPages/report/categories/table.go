package categories

import (
	"github.com/jomei/notionapi"
)

func createTable(categories, scores []string) notionapi.Block {
	table := notionapi.TableBlock{
		BasicBlock: notionapi.BasicBlock{
			Type:   notionapi.BlockType("table"),
			Object: notionapi.ObjectType("block"),
		},
		Table: notionapi.Table{
			TableWidth:      2,
			HasColumnHeader: true,
			Children:        createTableRows(categories, scores),
		},
	}
	return table
}

func createTableRows(categories, scores []string) []notionapi.Block {
	rows := []notionapi.Block{}
	rows = append(rows, createTableHeader())
	for key, category := range categories {
		score := scores[key]
		r := createTableRow(category, score)
		rows = append(rows, r)
	}
	return rows
}

func createTableHeader() notionapi.TableRowBlock {
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

func createTableRow(category, score string) notionapi.TableRowBlock {
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
