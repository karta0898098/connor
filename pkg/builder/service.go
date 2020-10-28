package builder

import (
	"github.com/iancoleman/strcase"
	"github.com/karta0898098/connor/pkg/template"
	"os"
)

// BuildServiceModule build service module
func (app *AppBuilder) BuildServiceModule() *AppBuilder {
	builder := &CodeBuilder{
		Template:    template.ServiceModule,
		ProjectName: app.projectName,
		Path:        "pkg/service",
		File:        "module.go",
		Data: H{
			"ProjectName": app.projectName,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// BuildService build service
func (app *AppBuilder) BuildService(service string) *AppBuilder {
	builder := &CodeBuilder{
		Template: template.Service,
		Path:     "pkg/service",
		File:     strcase.ToLowerCamel(service) + ".go",
		Data: H{
			"ProjectName": app.projectName,
			"Name":        strcase.ToCamel(service),
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// AddServiceModule add service to module
func (app *AppBuilder) AddServiceModule(service string) *AppBuilder {
	if _, err := os.Stat("./pkg/service/" + strcase.ToSnake(service) + ".go"); err == nil {
		return app
	}

	builder := &CodeBuilder{
		Template: template.AddServiceModule("New" + strcase.ToCamel(service) + "Service"),
		Path:     "pkg/service",
		File:     "module.go",
	}

	app.actions = append(app.actions, builder)
	return app
}