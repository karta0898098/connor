package builder

import "github.com/karta0898098/connor/pkg/template"

// BuildMain build main.go
func (app *AppBuilder) BuildMain() *AppBuilder {
	builder := &CodeBuilder{
		Template:    template.Main,
		ProjectName: app.projectName,
		Path:        "cmd/" + app.projectName,
		File:        app.projectName + ".go",
		Data: H{
			"ProjectName": app.projectName,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// BuildServer build server.go
func (app *AppBuilder) BuildServer(http string) *AppBuilder {
	engine := ""
	if http == "gin" {
		engine = "RunGin"
	} else {
		engine = "RunEcho"
	}

	builder := &CodeBuilder{
		Template:    template.Server,
		ProjectName: app.projectName,
		Path:        "cmd/" + app.projectName,
		File:        "server.go",
		Data: H{
			"ProjectName": app.gomodPath,
			"Engine":      engine,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}
