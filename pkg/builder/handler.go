package builder

import (
	"github.com/karta0898098/connor/pkg/template"
	"os"
	"strings"
)

// BuildHandlerModule create build handler.go
func (app *AppBuilder) BuildHandlerModule() *AppBuilder {

	builder := &CodeBuilder{
		Template:    template.Handler,
		ProjectName: app.projectName,
		Path:        "pkg/handler",
		File:        "handler.go",
		Data: H{
			"ProjectName": app.gomodPath,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// BuildHandlerRouter build handler router.go
func (app *AppBuilder) BuildHandlerRouter(http string) *AppBuilder {

	tmpl := ""
	if http == "gin" {
		tmpl = template.GinRouter
	} else {
		tmpl = template.EchoRouter
	}

	builder := &CodeBuilder{
		Template:    tmpl,
		ProjectName: app.projectName,
		Path:        "pkg/handler",
		File:        "router.go",
		Data: H{
			"ProjectName": app.projectName,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// AddHandlerModule add controller to handler
func (app *AppBuilder) AddHandlerModule(name string) *AppBuilder {
	if _, err := os.Stat("./pkg/handler/controller/" + strings.ToLower(name) + ".go"); err == nil {
		return app
	}

	builder := &CodeBuilder{
		Template: template.AddHandlerModule(name),
		Path:     "pkg/handler",
		File:     "handler.go",
	}

	app.actions = append(app.actions, builder)
	return app
}
