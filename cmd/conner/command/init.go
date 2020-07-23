package command

import (
	"github.com/karta0898098/connor/pkg/builder"
	"github.com/urfave/cli/v2"
)

// Init for init project action
func Init(context *cli.Context) error {

	name := context.String("name")
	httpEngine := context.String("http")

	app := builder.NewApp()
	app.ProjectName(name).
		WorkingDir().
		Folder().
		BuildConfigs().
		BuildConfigToml().
		BuildControllerModule().
		BuildHandlerModule().
		BuildHandlerRouter(httpEngine).
		BuildServiceModule().
		BuildRepositoryModule().
		BuildServer(httpEngine).
		BuildMain().
		BuildGoMod().
		Build()

	return nil
}
