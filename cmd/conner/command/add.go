package command

import (
	"github.com/karta0898098/connor/cmd/conner/flag"
	"github.com/karta0898098/connor/pkg/builder"
	"github.com/karta0898098/connor/pkg/util"
	"github.com/urfave/cli/v2"
)

// NewAddCommand new command for add controller and entity
func NewAddCommand() []*cli.Command {
	return []*cli.Command{
		{
			Name:      "controller",
			Usage:     "add controller",
			UsageText: "add controller and create construct",
			Flags:     flag.NewAddController(),
			Action:    AddController,
		},
		{
			Name:   "entity",
			Usage:  "add entity",
			Flags:  flag.NewAddEntity(),
			Action: AddEntity,
		},
	}
}

// AddController add controller action
func AddController(context *cli.Context) error {

	controller := context.String("name")
	projectName := util.FindProjectName()
	httpEngine := util.FindHttpEngine()


	app := builder.NewApp()
	app.GoMod(projectName).
		WorkingDir()

	if controller != "" {
		name := util.FixControllerName(controller)
		app = app.
			BuildController(name, httpEngine).
			AddControllerModule(name).
			AddHandlerModule(name)
	}

	app.Build()

	return nil
}

// AddEntity add entity
func AddEntity(context *cli.Context) error {

	filePath := context.String("file")
	srv := context.Bool("srv")
	repo := context.Bool("repo")

	projectName := util.FindProjectName()

	app := builder.NewApp()
	app.ProjectName(projectName).
		WorkingDir()

	if filePath != "" {
		model := util.FindModelName(filePath)

		if repo {
			app = app.BuildRepository(model)
			app = app.AddModelWhereAndUpdate(filePath, model)
			app = app.AddRepository(model)
		}

		if srv {
			app = app.BuildService(model)
			app = app.AddServiceModule(model)
		}
	}

	app.Build()
	return nil
}
