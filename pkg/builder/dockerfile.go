package builder

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/karta0898098/connor/pkg/template"
)

//BuildDockerfile build dockerfile
func (app *AppBuilder) BuildDockerfile() *AppBuilder {
	builder := &CodeBuilder{
		Template:    template.Dockerfile,
		ProjectName: app.projectName,
		Path:        "deployments/docker",
		File:        fmt.Sprintf("%s.dockerfile", strcase.ToLowerCamel(app.projectName)),
		Data: H{
			"ProjectName": strcase.ToLowerCamel(app.projectName),
		},
	}
	app.actions = append(app.actions, builder)
	return app
}
