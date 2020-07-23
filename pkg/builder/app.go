package builder

import (
	"os"
	"path"
)

// AppBuilder to build app
type AppBuilder struct {
	projectName string
	workingDir  string
	actions     []ExecAction
}

// AppBuilder new constructor
func NewApp() *AppBuilder {
	return &AppBuilder{}
}

// H ...
type H map[string]interface{}

// ExecAction exec to build app real action
type ExecAction interface {
	Build()
}

// ProjectName this project name
func (app *AppBuilder) ProjectName(name string) *AppBuilder {
	app.projectName = name
	return app
}

// WorkingDir to set go mod
func (app *AppBuilder) WorkingDir() *AppBuilder {
	nowPath, _ := os.Getwd()
	app.workingDir = path.Join(nowPath, app.projectName)
	return app
}

// Folder create default folder structure & go mod init
func (app *AppBuilder) Folder() *AppBuilder {

	folders := []string{
		app.projectName,
		"cmd",
		"cmd/" + app.projectName,
		"deployments",
		"deployments/config",
		"deployments/docker",
		"configs",
		"pkg",
		"pkg/handler",
		"pkg/handler/controller",
		"pkg/model",
		"pkg/service",
		"pkg/repository",
	}


	builder := NewFolderBuilder(app.workingDir)
	builder.
		ProjectName(app.projectName).
		Folders(folders)

	app.actions = append(app.actions, builder)

	return app
}

func (app *AppBuilder) BuildGoMod() *AppBuilder  {

	packages := []string{
		"github.com/jinzhu/gorm",
		"github.com/karta0898098/kara",
		"github.com/gin-gonic/gin",
		"github.com/labstack/echo/v4",
		"github.com/spf13/cobra",
		"github.com/spf13/viper",
		"github.com/rs/zerolog/log",
		"go.uber.org/fx",
	}

	builder := NewFolderBuilder(app.workingDir)
	builder.
		ProjectName(app.projectName).
		Packages(packages)

	app.actions = append(app.actions, builder)

	return app
}

// Build ...
func (app *AppBuilder) Build() {
	for i := 0; i < len(app.actions); i++ {
		app.actions[i].Build()
	}
}
