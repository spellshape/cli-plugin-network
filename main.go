package main

import (
	"encoding/gob"
	"os"

	hplugin "github.com/hashicorp/go-plugin"

	"github.com/spellshape/cli/spellshape/services/plugin"

	"github.com/spellshape/cli-plugin-network/cmd"
)

func init() {
	gob.Register(plugin.Manifest{})
	gob.Register(plugin.ExecutedCommand{})
	gob.Register(plugin.ExecutedHook{})
}

type p struct{}

func (p) Manifest() (plugin.Manifest, error) {
	m := plugin.Manifest{
		Name: "network",
	}
	m.ImportCobraCommand(cmd.NewNetwork(), "spellshape")
	return m, nil
}

func (p) Execute(c plugin.ExecutedCommand) error {
	// Instead of a switch on c.Use, we run the root command like if
	// we were in a command line context. This implies to set os.Args
	// correctly.
	// Remove the first arg "spellshape" from OSArgs because our network
	// command root is "network" not "spellshape".
	os.Args = c.OSArgs[1:]
	return cmd.NewNetwork().Execute()
}

func (p) ExecuteHookPre(hook plugin.ExecutedHook) error {
	return nil
}

func (p) ExecuteHookPost(hook plugin.ExecutedHook) error {
	return nil
}

func (p) ExecuteHookCleanUp(hook plugin.ExecutedHook) error {
	return nil
}

func main() {
	pluginMap := map[string]hplugin.Plugin{
		"cli-plugin-network": &plugin.InterfacePlugin{Impl: &p{}},
	}

	hplugin.Serve(&hplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig(),
		Plugins:         pluginMap,
	})
}
