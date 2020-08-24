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
	"github.com/karta0898098/kara/db/rw/db"
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
	"github.com/karta0898098/kara/db/rw/db"
    "github.com/karta0898098/kara/exception"
	"github.com/pkg/errors"
	"{{.ProjectName}}/pkg/model"
	"reflect"
)

// {{.Name}}Repository ...
type {{.Name}}Repository interface {
	Begin() {{.Name}}Repository
	Commit() error
	Rollback() error
	Get(ctx context.Context, condition model.Query{{.Name}}, forUpdate bool) (model.{{.Name}}, error)
	List(ctx context.Context, condition model.Query{{.Name}}, forUpdate bool) ([]model.{{.Name}} ,error)
	Create(ctx context.Context, data model.{{.Name}}) (model.{{.Name}}, error)
	Update(ctx context.Context, condition model.Query{{.Name}}, data interface{}) error
	Delete(ctx context.Context, condition model.Query{{.Name}}) error
	Count(ctx context.Context, condition model.Query{{.Name}}) (int,error)
}

// {{ToLowerCamel .Name}}Repository ...
type {{ToLowerCamel .Name}}Repository struct {
	readDB  *gorm.DB
	writeDB *gorm.DB
}


// New{{.Name}}Repository new constructor
func New{{.Name}}Repository(conn *db.Connection) {{.Name}}Repository {
	return &{{ToLowerCamel .Name}}Repository{
		readDB:  conn.ReadDB,
		writeDB: conn.WriteDB,
	}
}

// Begin for transactions get tx
func (repo *{{ToLowerCamel .Name}}Repository) Begin() {{.Name}}Repository {
	tx := repo.writeDB.Begin()
	return &{{ToLowerCamel .Name}}Repository{
		readDB:  tx,
		writeDB: tx,
	}
}

// Commit for transactions commit
func (repo *{{ToLowerCamel .Name}}Repository) Commit() error {
	return repo.writeDB.Commit().Error
}

// Rollback for transactions rollback
func (repo *{{ToLowerCamel .Name}}Repository) Rollback() error {
	return repo.writeDB.Rollback().Error
}

// forUpdate for transactions rollback
func (repo *{{ToLowerCamel .Name}}Repository) forUpdate(forUpdate bool) *gorm.DB  {
	if forUpdate {
		repo.readDB = repo.readDB.Set("gorm:query_option", "FOR UPDATE")
	} else {
		repo.readDB = repo.readDB.Set("gorm:query_option", "LOCK IN SHARE MODE")
	}
	return repo.readDB
}


// Get {{ToLowerCamel .Name}} ...
func (repo *{{ToLowerCamel .Name}}Repository) Get(ctx context.Context, condition model.Query{{.Name}}, forUpdate bool) (model.{{.Name}}, error) {


	var {{ToLowerCamel .Name}} model.{{.Name}}
	err := repo.forUpdate(forUpdate).Model(&model.{{.Name}}{}).Scopes(condition.Where).First(&{{ToLowerCamel .Name}}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.{{.Name}}{}, errors.Wrap(exception.ErrResourceNotFound, "get {{ToLowerCamel .Name}} from database error")
		}
		return model.{{.Name}}{}, errors.Wrap(err, "get {{ToLowerCamel .Name}} from database error")
	}

	return {{ToLowerCamel .Name}}, nil
}

// List {{ToLowerCamel .Name}} ...
func (repo *{{ToLowerCamel .Name}}Repository) List(ctx context.Context, condition model.Query{{.Name}}, forUpdate bool) ([]model.{{.Name}} , error) {

	var {{ToLowerCamel .Plural}} []model.{{.Name}}

	err := repo.forUpdate(forUpdate).Model(&model.{{.Name}}{}).Scopes(condition.Where).Find(&{{ToLowerCamel .Plural}}).Error
	if err != nil {
		{{ToLowerCamel .Plural}} = make([]model.{{.Name}}, 0)
		return {{ToLowerCamel .Plural}}, err
	}
	return {{ToLowerCamel .Plural}}, nil
}

// Create {{ToLowerCamel .Name}} ...
func (repo *{{ToLowerCamel .Name}}Repository) Create(ctx context.Context, data model.{{.Name}}) (model.{{.Name}}, error) {

	err := repo.writeDB.Model(&model.{{.Name}}{}).Create(&data).Error
	if err != nil {
		return model.{{.Name}}{}, err
	}

	return data, nil
}

// Update {{ToLowerCamel .Name}} ...
func (repo *{{ToLowerCamel .Name}}Repository) Update(ctx context.Context, condition model.Query{{.Name}}, data interface{}) error {

	if reflect.DeepEqual(condition, model.Query{{.Name}}{}) {
		return errors.Wrap(exception.ErrInvalidInput, "repository: {{ToLowerCamel .Name}} query condition is nil")
	}

	err := repo.writeDB.Model(&model.{{.Name}}{}).Scopes(condition.Where).Updates(data).Error
	return err
}

// Delete {{ToLowerCamel .Name}} ...
func (repo *{{ToLowerCamel .Name}}Repository) Delete(ctx context.Context, condition model.Query{{.Name}}) error {

	if reflect.DeepEqual(condition, model.Query{{.Name}}{}) {
		return errors.Wrap(exception.ErrInvalidInput, "repository: {{ToLowerCamel .Name}} query condition is nil")
	}

	err := repo.writeDB.Scopes(condition.Where).Delete(model.{{.Name}}{}).Error
	if err != nil {
		return err
	}

	return nil
}

// Count {{ToLowerCamel .Name}} ...
func (repo *{{ToLowerCamel .Name}}Repository) Count(ctx context.Context, condition model.Query{{.Name}}) (int,error) {

	var count int

	if reflect.DeepEqual(condition, model.Query{{.Name}}{}) {
		return count,errors.Wrap(exception.ErrInvalidInput, "repository: {{ToLowerCamel .Name}} query condition is nil")
	}

	err := repo.readDB.Model(&model.{{.Name}}{}).Scopes(condition.Where).Count(&count).Error
	if err != nil{
		return count,err
	}
	return count,nil
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
