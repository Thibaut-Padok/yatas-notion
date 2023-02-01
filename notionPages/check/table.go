package check

import (
	"github.com/jomei/notionapi"
	"github.com/stangirard/yatas/plugins/commons"
)

func createTable(check commons.Check) notionapi.Block {
	table := notionapi.TableBlock{
		BasicBlock: notionapi.BasicBlock{
			Type:   notionapi.BlockType("table"),
			Object: notionapi.ObjectType("block"),
		},
		Table: notionapi.Table{
			TableWidth:      3,
			HasColumnHeader: true,
			Children:        createRowsTable(check.Results),
		},
	}
	return table
}

func createRowsTable(results []commons.Result) []notionapi.Block {
	rows := []notionapi.Block{}
	rows = append(rows, createHeaderRowTable())
	for _, result := range results {
		r := createRowTable(result)
		rows = append(rows, r)
	}
	return rows
}

func createHeaderRowTable() notionapi.TableRowBlock {
	id := notionapi.Text{
		Content: "ID",
	}
	status := notionapi.Text{
		Content: "Status",
	}
	message := notionapi.Text{
		Content: "Message",
	}
	row := notionapi.TableRowBlock{
		BasicBlock: notionapi.BasicBlock{
			Type:   notionapi.BlockType("table_row"),
			Object: notionapi.ObjectType("block"),
		},
		TableRow: notionapi.TableRow{
			Cells: [][]notionapi.RichText{
				{{Text: &id}},
				{{Text: &status}},
				{{Text: &message}},
			},
		},
	}
	return row
}

func createRowTable(res commons.Result) notionapi.TableRowBlock {
	emojiStatus := ""
	if res.Status == "FAIL" {
		emojiStatus = "❌"
	} else {
		emojiStatus = "✅"
	}
	id := notionapi.Text{
		Content: res.ResourceID,
	}
	status := notionapi.Text{
		Content: emojiStatus,
	}
	message := notionapi.Text{
		Content: res.Message,
	}
	row := notionapi.TableRowBlock{
		BasicBlock: notionapi.BasicBlock{
			Type:   notionapi.BlockType("table_row"),
			Object: notionapi.ObjectType("block"),
		},
		TableRow: notionapi.TableRow{
			Cells: [][]notionapi.RichText{
				{{Text: &id}},
				{{Text: &status}},
				{{Text: &message}},
			},
		},
	}
	return row
}
