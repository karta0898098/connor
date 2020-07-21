package template

import (
	"bytes"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"io/ioutil"
)

// RepositoryModule template ...
const RepositoryModule = `package repository

import (
	"github.com/karta0898098/kara/db"
	"go.uber.org/fx"
)

// Module for export repository to fx injection
var Module = fx.Provide(
	db.NewConnection,
)`

// Repository template
const Repository = `package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/karta0898098/kara/db"
    "github.com/karta0898098/kara/exception"
	"github.com/pkg/errors"
	"{{.ProjectName}}/pkg/model"
	"reflect"
)

// I{{.Name}}Repository ...
type I{{.Name}}Repository interface {
	Get(ctx context.Context, condition model.Where{{.Name}}, tx *gorm.DB) (model.{{.Name}}, error)
	List(ctx context.Context, condition model.Where{{.Name}}, tx *gorm.DB) ([]model.{{.Name}}, error)
	Create(ctx context.Context, data model.{{.Name}}, tx *gorm.DB) (model.{{.Name}}, error)
	Update(ctx context.Context, data model.{{.Name}}, condition model.Where{{.Name}}, tx *gorm.DB) error
	Delete(ctx context.Context, condition model.Where{{.Name}}, tx *gorm.DB) error
}

// {{.Name}}Repository ...
type {{.Name}}Repository struct {
	ReadDB  *gorm.DB
	WriteDB *gorm.DB
}

// New{{.Name}}Repository new constructor
func New{{.Name}}Repository(conn *db.Connection) I{{.Name}}Repository {
	return &{{.Name}}Repository{
		ReadDB:  conn.ReadDB,
		WriteDB: conn.WriteDB,
	}
}

// Get {{ToLowerCamel .Name}} ...
func (repo *{{.Name}}Repository) Get(ctx context.Context, condition model.Where{{.Name}}, tx *gorm.DB) (model.{{.Name}}, error) {

	if tx == nil {
		tx = repo.ReadDB
	}

	var {{ToLowerCamel .Name}} model.{{.Name}}
	err := tx.Model(&model.{{.Name}}{}).Scopes(condition.Where).First(&{{ToLowerCamel .Name}}).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.{{.Name}}{}, errors.Wrap(exception.ErrResourceNotFound, "get {{ToLowerCamel .Name}} from database error")
		}
		return model.{{.Name}}{}, errors.Wrap(err, "get {{ToLowerCamel .Name}} from database error")
	}

	return {{ToLowerCamel .Name}}, nil
}

// List {{ToLowerCamel .Name}} ...
func (repo *{{.Name}}Repository) List(ctx context.Context, condition model.Where{{.Name}}, tx *gorm.DB) ([]model.{{.Name}}, error) {

	if tx == nil {
		tx = repo.ReadDB
	}

	var {{.Plural}} []model.{{.Name}}

	err := tx.Model(&model.{{.Name}}{}).Scopes(condition.Where).Find(&{{.Plural}}).Error
	if err != nil {
		{{.Plural}} = make([]model.{{.Name}}, 0)
		return {{.Plural}}, err
	}
	return {{.Plural}}, nil
}

// Create {{ToLowerCamel .Name}} ...
func (repo *{{.Name}}Repository) Create(ctx context.Context, data model.{{.Name}}, tx *gorm.DB) (model.{{.Name}}, error) {

	if tx == nil {
		tx = repo.WriteDB
	}

	err := tx.Model(&model.{{.Name}}{}).Create(&data).Error
	if err != nil {
		return model.{{.Name}}{}, err
	}

	return data, nil
}

// Update {{ToLowerCamel .Name}} ...
func (repo *{{.Name}}Repository) Update(ctx context.Context, data model.{{.Name}}, condition model.Where{{.Name}}, tx *gorm.DB) error {

	if tx == nil {
		tx = repo.WriteDB
	}

	if reflect.DeepEqual(condition, model.Where{{.Name}}{}) {
		return errors.Wrap(exception.ErrInvalidInput, "repository: {{ToLowerCamel .Name}} where condition is nil")
	}

	err := tx.Model(&model.{{.Name}}{}).Scopes(condition.Where).Update(&data).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete {{ToLowerCamel .Name}} ...
func (repo *{{.Name}}Repository) Delete(ctx context.Context, condition model.Where{{.Name}}, tx *gorm.DB) error {

	if tx == nil {
		tx = repo.WriteDB
	}

	if reflect.DeepEqual(condition, model.Where{{.Name}}{}) {
		return errors.Wrap(exception.ErrInvalidInput, "repository: {{ToLowerCamel .Name}} where condition is nil")
	}

	err := tx.Scopes(condition.Where).Delete(model.{{.Name}}{}).Error
	if err != nil {
		return err
	}

	return nil
}

`

// AddRepositoryModule template ...
func AddRepositoryModule(repo string) string {
	code, _ := ioutil.ReadFile("./pkg/repository/module.go")

	f, err := decorator.Parse(code)
	if err != nil {
		fmt.Println("can't add repository in module reason: ", err)
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
