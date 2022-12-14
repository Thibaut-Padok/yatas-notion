package main

import (
	"github.com/stangirard/yatas/plugins/commons"
)

// Notion Config struct
type NotionAccount struct {
	Token     string `yaml:"token"`               // Token of notion connection to use notionapi/v1
	AuthToken string `yaml:"authToken,omitempty"` // Token of web application www.notion.so to use notionapi/v3 (optionnal)

	PageID     string `yaml:"page_id"`         // PageID which contains Yatas database
	DatabaseID string `yaml:"db_id,omitempty"` // DatabaseID which contains Yatas instances
}

func loadNotionPluginConfig(c *commons.Config) NotionAccount {
	var account NotionAccount

	for _, r := range c.PluginConfig {
		var potentialAccount NotionAccount
		isThisPlugin := false

		for key, value := range r {
			switch key {
			case "pluginName":
				if value == "notion" {
					isThisPlugin = true

				}
			case "token":
				potentialAccount.Token = value.(string)
			case "pageID":
				potentialAccount.PageID = value.(string)
			case "authToken":
				potentialAccount.AuthToken = value.(string)
			}
		}
		if isThisPlugin {
			account = potentialAccount
			break
		}
	}
	return account
}
