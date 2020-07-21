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

var Module = fx.Provide()
`

// Service template ...
const Service = `package service

import (
	"context"
	"{{.ProjectName}}/pkg/model"
	"{{.ProjectName}}/pkg/repository"
)

// I{{.Name}}Service {{.Name}} service implement...
type I{{.Name}}Service interface {
	Get(ctx context.Context, condition model.Where{{.Name}}) (model.{{.Name}}, error)
	List(ctx context.Context, condition model.Where{{.Name}}) ([]model.{{.Name}}, error)
	Create(ctx context.Context, {{ToLowerCamel .Name}} model.{{.Name}}) (model.{{.Name}}, error)
	Update(ctx context.Context, condition model.Update{{.Name}}) error
	Delete(ctx context.Context, condition model.Where{{.Name}}) error
}

// {{.Name}}Service {{.Name}} service  ...
type {{.Name}}Service struct {
	repo repository.I{{.Name}}Repository
}

// New{{.Name}}Service new service constructor
func New{{.Name}}Service(repo repository.I{{.Name}}Repository) I{{.Name}}Service {
	return &{{.Name}}Service{
		repo: repo,
	}
}

// Get ...
func (srv *{{.Name}}Service) Get(ctx context.Context, where model.Where{{.Name}}) (model.{{.Name}}, error) {
	return srv.repo.Get(ctx, where, nil)
}

// List ...
func (srv *{{.Name}}Service) List(ctx context.Context, where model.Where{{.Name}}) ([]model.{{.Name}}, error) {
	return srv.repo.List(ctx, where, nil)
}

// Create ...
func (srv *{{.Name}}Service) Create(ctx context.Context, {{ToLowerCamel .Name}} model.{{.Name}}) (model.{{.Name}}, error) {
	return srv.repo.Create(ctx, {{ToLowerCamel .Name}}, nil)
}

// Update ...
func (srv *{{.Name}}Service) Update(ctx context.Context, update model.Update{{.Name}}) error {
	return srv.repo.Update(ctx, update.{{.Name}}, update.Where, nil)
}

// Delete ...
func (srv {{.Name}}Service) Delete(ctx context.Context, where model.Where{{.Name}}) error {
	return srv.repo.Delete(ctx, where, nil)
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

