package builder

import "github.com/karta0898098/connor/pkg/template"

// BuildConfigs build configs/config.go
func (app *AppBuilder) BuildConfigs(useYaml bool) *AppBuilder {

	configType := "toml"
	if useYaml {
		configType = "yaml"
	}
	builder := &CodeBuilder{
		Template:    template.Config,
		ProjectName: app.projectName,
		Path:        "configs",
		File:        "config.go",
		Data: H{
			"ProjectName": app.projectName,
			"ConfigType":  configType,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// BuildConfig build app.toml for lunch app
func (app *AppBuilder) BuildConfig(useYaml bool) *AppBuilder {

	builder := &CodeBuilder{}
	if !useYaml {
		builder = &CodeBuilder{
			Template:    template.AppToml,
			ProjectName: app.projectName,
			Path:        "deployments/config",
			File:        "app.toml",
			Data: H{
				"ProjectName": app.projectName,
			},
		}
	} else {
		builder = &CodeBuilder{
			Template:    template.AppYaml,
			ProjectName: app.projectName,
			Path:        "deployments/config",
			File:        "app.yaml",
			Data: H{
				"ProjectName": app.projectName,
			},
		}
	}

	app.actions = append(app.actions, builder)
	return app
}
