package template

import (
	"bytes"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"io/ioutil"
)

// ServiceModule template ...
const ServiceModule = `package service

import "go.uber.org/fx"

// Module for fx provide service
var Module = fx.Provide()
`

// Service template ...
const Service = `package service

import (
	"context"
	"{{.ProjectName}}/pkg/model"
	"{{.ProjectName}}/pkg/repository"
)

// {{.Name}}Service {{.Name}} service implement...
type {{.Name}}Service interface {
	Get(ctx context.Context, condition model.Query{{.Name}}) (model.{{.Name}}, error)
	List(ctx context.Context, condition model.Query{{.Name}}) ([]model.{{.Name}}, error)
	Create(ctx context.Context, {{ToLowerCamel .Name}} model.{{.Name}}) (model.{{.Name}}, error)
	Update(ctx context.Context, condition model.Query{{.Name}} , data interface{}) error
	Delete(ctx context.Context, condition model.Query{{.Name}}) error
}

// {{ToLowerCamel .Name}}Service {{ToLowerCamel .Name}} service  ...
type {{ToLowerCamel .Name}}Service struct {
	repo repository.{{.Name}}Repository
}

// New{{.Name}}Service new service constructor
func New{{.Name}}Service(repo repository.{{.Name}}Repository) {{.Name}}Service {
	return &{{ToLowerCamel .Name}}Service{
		repo: repo,
	}
}

// Get ...
func (srv *{{ToLowerCamel .Name}}Service) Get(ctx context.Context, where model.Query{{.Name}}) (model.{{.Name}}, error) {
	return srv.repo.Get{{.Name}}(ctx, where, false)
}

// List ...
func (srv *{{ToLowerCamel .Name}}Service) List(ctx context.Context, where model.Query{{.Name}}) ([]model.{{.Name}}, error) {
	return srv.repo.List{{ToPlural .Name}}(ctx, where, false)
}

// Create ...
func (srv *{{ToLowerCamel .Name}}Service) Create(ctx context.Context, {{ToLowerCamel .Name}} model.{{.Name}}) (model.{{.Name}}, error) {
	return srv.repo.Create{{.Name}}(ctx, {{ToLowerCamel .Name}})
}

// Update ...
func (srv *{{ToLowerCamel .Name}}Service) Update(ctx context.Context, where model.Query{{.Name}},data interface{}) error {
	return srv.repo.Update{{.Name}}(ctx, where, data)
}

// Delete ...
func (srv {{ToLowerCamel .Name}}Service) Delete(ctx context.Context, where model.Query{{.Name}}) error {
	return srv.repo.Delete{{.Name}}(ctx, where)
}
`

// AddServiceModule template ...
func AddServiceModule(repo string) string {
	code, _ := ioutil.ReadFile("./pkg/service/module.go")

	f, err := decorator.Parse(code)
	if err != nil {
		fmt.Println("can't add service in module reason: ", err)
		return ""
	}

	spec := f.Decls[1].(*dst.GenDecl).Specs

	// scan all go source code
	// find to var Module = fx.Provide()
	for index, item := range spec {

		if item.(*dst.ValueSpec).Names[index].Name == "Module" {
			call := item.(*dst.ValueSpec).Values[0].(*dst.CallExpr)

			call.Decs.Before = dst.EmptyLine
			call.Decs.After = dst.EmptyLine
			call.Args = append(call.Args, dst.NewIdent(repo))
			for _, v := range call.Args {
				v, ok := v.(*dst.Ident)
				if ok {
					v.Decs.Before = dst.NewLine
					v.Decs.After = dst.NewLine
				}
			}
		}
	}

	var buf bytes.Buffer
	err = decorator.Fprint(&buf, f)
	if err != nil {
		fmt.Println("can't add service in module reason: ", err)
		return ""
	}

	return buf.String()
}

