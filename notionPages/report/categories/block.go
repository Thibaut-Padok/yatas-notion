package categories

import (
	"github.com/jomei/notionapi"
	"github.com/stangirard/yatas/plugins/commons"

	"github.com/Thibaut-Padok/yatas-notion/utils"
)

func CreateBlock(test commons.Tests) notionapi.Block {
	// Find the categories
	categories := []string{}
	categoriesSuccess := map[string]int{}
	categoriesFailure := map[string]int{}
	for _, check := range test.Checks {
		for _, category := range check.Categories {
			if !utils.Contains(categories, category) {
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
		scores = append(scores, utils.CalculatePercent(categoriesSuccess[category], categoriesFailure[category])+" %")
	}
	// Write the categories
	block := createTable(categories, scores)
	return block
}
