package builder

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/karta0898098/connor/pkg/template"
	"io/ioutil"
	"os"
)

// BuildRepositoryModule build repository module
func (app *AppBuilder) BuildRepositoryModule() *AppBuilder {

	builder := &CodeBuilder{
		Template:    template.RepositoryModule,
		ProjectName: app.projectName,
		Path:        "pkg/repository",
		File:        "module.go",
		Data: H{
			"ProjectName": app.projectName,
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// BuildRepository build repository default CRUD
func (app *AppBuilder) BuildRepository(name string) *AppBuilder {

	builder := &CodeBuilder{
		Template: template.Repository,
		Path:     "pkg/repository",
		File:     strcase.ToSnake(name) + ".go",
		Data: H{
			"ProjectName": app.projectName,
			"Name":        strcase.ToCamel(name),
			"Plural":      inflection.Plural(strcase.ToLowerCamel(name)),
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// AddModelWhereAndUpdate add repository where and update condition
func (app *AppBuilder) AddModelWhereAndUpdate(filePath, name string) *AppBuilder {

	code, _ := ioutil.ReadFile(filePath)

	builder := &CodeBuilder{
		Template: template.AddModelQueryCondition(code),
		Path:     "pkg/model",
		File:     strcase.ToSnake(name) + ".go",
		Data: H{
			"ProjectName": app.projectName,
			"Name":        strcase.ToCamel(name),
			"Plural":      inflection.Plural(strcase.ToLowerCamel(name)),
		},
	}

	app.actions = append(app.actions, builder)
	return app
}

// AddRepository add repository to module
func (app *AppBuilder) AddRepository(name string) *AppBuilder {
	if _, err := os.Stat("./pkg/repository/" + strcase.ToSnake(name) + ".go"); err == nil {
		return app
	}

	builder := &CodeBuilder{
		Template: template.AddRepositoryModule(strcase.ToCamel(name) + "Repository"),
		Path:     "pkg/repository",
		File:     "module.go",
	}

	app.actions = append(app.actions, builder)
	return app
}
