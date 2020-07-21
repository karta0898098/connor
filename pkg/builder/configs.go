package builder

import "github.com/karta0898098/connor/pkg/template"

// BuildConfigs build configs/config.go
func (app *AppBuilder) BuildConfigs() *AppBuilder {

	builder := &CodeBuilder{
		Template:    template.Config,
		ProjectName: app.projectName,
		Path:        "configs",
		File:        "config.go",
		Data: H{
			"ProjectName": app.projectName,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// BuildConfigToml build app.toml for lunch app
func (app *AppBuilder) BuildConfigToml() *AppBuilder {

	builder := &CodeBuilder{
		Template:    template.AppToml,
		ProjectName: app.projectName,
		Path:        "deployments/config",
		File:        "app.toml",
		Data: H{
			"ProjectName": app.projectName,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

