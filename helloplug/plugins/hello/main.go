package main

import (
	// ...

	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/ttsubo2000/study-of-GoPlugin/helloplug/plug"
)

type Hello struct{}

func (h *Hello) Say(name string) (string, error) {
	return fmt.Sprintf("hello %s", name), nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plug.Handshake,
		Plugins: map[string]plugin.Plugin{
			"greeter": &plug.GreeterPlugin{Impl: &Hello{}},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}
