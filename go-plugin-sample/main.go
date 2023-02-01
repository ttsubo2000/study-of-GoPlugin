package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
	provider "github.com/ttsubo2000/study-of-GoPlugin/go-plugin-sample/common"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "PROVIDER_PLUGIN",
	MagicCookieValue: "azure",
}

var pluginMap = map[string]plugin.Plugin{
	"azure": &provider.ProviderPlugin{},
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "client",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin/azure_provider"),
		Logger:          logger,
	})
	defer client.Kill()
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	raw, err := rpcClient.Dispense("azure")
	if err != nil {
		log.Fatal(err)
	}
	azureProvider := raw.(provider.Provider)
	message := "Hello"
	scheme := provider.Scheme{
		message,
	}
	fmt.Println(azureProvider.Create(&scheme))
}
