package main

import (
	"github.com/asafcyi/Stechhelm/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "stechhelm"
	app.Description = "The JFrog Platform Protector."
	app.Version = "v1.0.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetAuditCommand(),
		commands.GetGraphCommand(),
	}
}
