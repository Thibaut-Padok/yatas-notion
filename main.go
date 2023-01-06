package main

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/stangirard/yatas/plugins/commons"
)

type YatasPlugin struct {
	logger hclog.Logger
}

// Don't remove this function
func (g *YatasPlugin) Run(c *commons.Config) []commons.Tests {
	g.logger.Debug("message from Yatas Notion Plugin")
	var err error
	if err != nil {
		panic(err)
	}
	var checksAll []commons.Tests

	checks, err := runPlugin(c, "notion")
	if err != nil {
		g.logger.Error("Error running plugins", "error", err)
	}
	checksAll = append(checksAll, checks...)
	return checksAll
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  2,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	gob.Register([]interface{}{})
	gob.Register(map[string]interface{}{})
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	yatasPlugin := &YatasPlugin{
		logger: logger,
	}
	// pluginMap is the map of plugins we can dispense.
	// Name of your plugin
	var pluginMap = map[string]plugin.Plugin{
		"notion": &commons.YatasPlugin{Impl: yatasPlugin},
	}

	logger.Debug("Message from plugin", "YATAS-NOTION", "!")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

// Function that runs the checks or things to dot
func runPlugin(c *commons.Config, plugin string) ([]commons.Tests, error) {
	var checksAll []commons.Tests

	// Load notion account values
	var account = loadNotionPluginConfig(c)

	// Init client
	client := NewNotionClient(&account)

	//Init Yatas database
	isDatabaseExist := initYatasDatabase(&client, &account)

	if !isDatabaseExist {
		log.Printf("Failed to get Database. Stop notion reporting!")
		return checksAll, nil
	}

	// Create Yatas report instance
	CreateNotionReport(c.Tests, account, &client)

	return checksAll, nil
}
