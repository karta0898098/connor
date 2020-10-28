package builder

import (
	"github.com/iancoleman/strcase"
	"github.com/karta0898098/connor/pkg/template"
	"os"
	"strings"
)

// BuildControllerModule build default controller module
func (app *AppBuilder) BuildControllerModule() *AppBuilder {

	builder := &CodeBuilder{
		Template:    template.ControllerModule,
		ProjectName: app.projectName,
		Path:        "pkg/handler/controller",
		File:        "module.go",
		Data: H{
			"ProjectName": app.projectName,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// BuildController build controller crud method
func (app *AppBuilder) BuildController(name, httpEngine string) *AppBuilder {
	tmpl := ``

	if httpEngine == "gin" {
		tmpl = template.GinController
	}else {
		tmpl = template.EchoController
	}

	builder := &CodeBuilder{
		Template: tmpl,
		Path:     "pkg/handler/controller",
		File:     strcase.ToSnake(name) + ".go",
		Data: H{
			"Name": name,
		},
	}

	app.actions = append(app.actions, builder)

	return app
}

// AddController add controller to fx.Provide()
func (app *AppBuilder) AddControllerModule(name string) *AppBuilder {

	if _, err := os.Stat("./pkg/handler/controller/" + strings.ToLower(name) + ".go"); err == nil {
		return app
	}

	builder := &CodeBuilder{
		Template: template.AddControllerInModule("New" + name + "Controller"),
		Path:     "pkg/handler/controller",
		File:     "module.go",
	}

	app.actions = append(app.actions, builder)
	return app
}
