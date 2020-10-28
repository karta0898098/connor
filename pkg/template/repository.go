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
	"context"

	"github.com/karta0898098/kara/db/rw/db"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module for export repository to fx injection
var Module = fx.Provide(
	db.NewConnection,
	NewRepository,
)

// ErrNilTx error nil tx
var ErrNilTx = errors.New("tx is nil, begin first or use Transaction")

// Repository ...
type Repository interface {
	Transaction(
		ctx context.Context,
		callback func(ctx context.Context, txRepo Repository) error,
	) error
	commit() error
	rollback() error
}

type repository struct {
	readDB  *gorm.DB
	writeDB *gorm.DB
	tx      *gorm.DB
}

// NewRepository repository new constructor
func NewRepository(conn *db.Connection) Repository {
	return &repository{
		readDB:  conn.ReadDB,
		writeDB: conn.WriteDB,
	}
}

// forUpdate for transactions rollback
func (repo *repository) forUpdate(forUpdate bool) *gorm.DB {
	tx := repo.getWriteDB()
	if forUpdate {
		tx = tx.Set("gorm:query_option", "FOR UPDATE")
	} else {
		tx = tx.Set("gorm:query_option", "LOCK IN SHARE MODE")
	}
	return tx
}

func (repo *repository) begin() Repository {
	tx := repo.writeDB.Begin()
	return &repository{
		tx:      tx,
		readDB:  repo.readDB,
		writeDB: repo.writeDB,
	}
}


func (repo *repository) getWriteDB() *gorm.DB {
	if repo.tx != nil {
		return repo.tx
	}
	return repo.writeDB
}

func (repo *repository) getReadDB() *gorm.DB {
	if repo.tx != nil {
		return repo.tx
	}
	return repo.readDB
}

func (repo *repository) commit() error {
	if repo.tx == nil {
		return ErrNilTx
	}

	return repo.tx.Commit().Error
}

func (repo *repository) rollback() error {
	if repo.tx == nil {
		return ErrNilTx
	}

	return repo.tx.Rollback().Error
}

// Transaction start rdbms transaction scope
func (repo *repository) Transaction(ctx context.Context, callback func(ctx context.Context, txRepo Repository) error) error {
	var (
		tx          Repository
		callbackErr error
		err         error
	)

	tx = repo.begin()

	callbackErr = callback(ctx, tx)

	if callbackErr != nil {
		err = tx.rollback()
	} else {
		if err = tx.commit(); err != nil {
			err = tx.rollback()
		}
	}

	if err != nil {
		if callbackErr != nil {
			err = errors.Wrapf(err, "transaction callback error reason : %v", callbackErr)
		}
		return err
	}

	return callbackErr
}

`

// Repository template
const Repository = `package repository

import (
	"context"
	"gorm.io/gorm"
    "github.com/karta0898098/kara/exception"
	"github.com/pkg/errors"
	"{{.ProjectName}}/pkg/model"
	"reflect"
)

// {{.Name}}Repository ...
type {{.Name}}Repository interface {
	Get{{.Name}}(ctx context.Context, condition model.Query{{.Name}}) (model.{{.Name}}, error)
	List{{ToPlural .Name}}(ctx context.Context, condition model.Query{{.Name}}) ([]model.{{.Name}} ,error)
	Create{{.Name}}(ctx context.Context, data model.{{.Name}}) (model.{{.Name}}, error)
	Update{{.Name}}(ctx context.Context, condition model.Query{{.Name}}, data interface{}) error
	Delete{{.Name}}(ctx context.Context, condition model.Query{{.Name}}) error
	Count{{.Name}}(ctx context.Context, condition model.Query{{.Name}}) (int64,error)
}


// Get{{.Name}} rdbms get {{ToLowerCamel .Name}}
func (repo *repository) Get{{.Name}}(ctx context.Context, condition model.Query{{.Name}}) (model.{{.Name}}, error) {
	var {{ToLowerCamel .Name}} model.{{.Name}}

	err := repo.
		getReadDB().
		WithContext(ctx).
		Model(&model.{{.Name}}{}).
		Scopes(
			condition.Scope,
		).
		First(&{{ToLowerCamel .Name}}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.{{.Name}}{}, errors.Wrap(exception.ErrResourceNotFound, "get {{ToLowerCamel .Name}} from database error")
		}
		return model.{{.Name}}{}, errors.Wrap(err, "get {{ToLowerCamel .Name}} from database error")
	}

	return {{ToLowerCamel .Name}}, nil
}

// List{{ToPlural .Name}} rdbms list {{ToLowerCamel .Name}}
func (repo *repository) List{{ToPlural .Name}}(ctx context.Context, condition model.Query{{.Name}}, forUpdate bool) ([]model.{{.Name}} , error) {
	var {{ToLowerCamel .Plural}} []model.{{.Name}}

	err := repo.
		getReadDB().
		WithContext(ctx).
		Model(&model.{{.Name}}{}).
		Scopes(condition.Scope).
		Find(&{{ToLowerCamel .Plural}}).Error

	if err != nil {
		{{ToLowerCamel .Plural}} = make([]model.{{.Name}}, 0)
		return {{ToLowerCamel .Plural}}, err
	}
	return {{ToLowerCamel .Plural}}, nil
}

// Create{{.Name}} rdbms create {{ToLowerCamel .Name}}
func (repo *repository) Create{{.Name}}(ctx context.Context, data model.{{.Name}}) (model.{{.Name}}, error) {
	err := repo.
		getWriteDB().
		WithContext(ctx).
		Create(&data).Error

	if err != nil {
		return model.{{.Name}}{}, err
	}

	return data, nil
}

// Update{{.Name}} rdbms update {{ToLowerCamel .Name}}
func (repo *repository) Update{{.Name}}(ctx context.Context, condition model.Query{{.Name}}, data interface{}) error {
	if reflect.DeepEqual(condition, model.Query{{.Name}}{}) {
		return errors.Wrap(exception.ErrInvalidInput, "repository: {{ToLowerCamel .Name}} query condition is nil")
	}

	err := repo.
		getWriteDB().
		WithContext(ctx).
		Model(&model.{{.Name}}{}).
		Scopes(condition.Scope).
		Updates(data).Error

	return err
}

// Delete{{.Name}} rdbms delete {{ToLowerCamel .Name}}
func (repo *repository) Delete{{.Name}}(ctx context.Context, condition model.Query{{.Name}}) error {
	if reflect.DeepEqual(condition, model.Query{{.Name}}{}) {
		return errors.Wrap(exception.ErrInvalidInput, "repository: {{ToLowerCamel .Name}} query condition is nil")
	}

	err := repo.
		getWriteDB().
		WithContext(ctx).
		Scopes(condition.Scope).
		Delete(model.{{.Name}}{}).Error

	if err != nil {
		return err
	}

	return nil
}

// Count{{.Name}} rdbms count {{ToLowerCamel .Name}}
func (repo *repository) Count{{.Name}}(ctx context.Context, condition model.Query{{.Name}}) (int64,error) {
	var count int64

	err := repo.
		getReadDB().
		WithContext(ctx).
		Model(&model.{{.Name}}{}).
		Scopes(condition.Scope).
		Count(&count).Error

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

	for i := 0; i < len(f.Decls); i++ {
		specsTree, ok := f.Decls[i].(*dst.GenDecl)
		if ok {
			for _, spec := range specsTree.Specs {
				typeSpec, ok := spec.(*dst.TypeSpec)
				if ok {
					if typeSpec.Name.Name == "Repository" {
						interfaceType, ok := typeSpec.Type.(*dst.InterfaceType)
						if ok {
							interfaceType.Methods.List = append(interfaceType.Methods.List, &dst.Field{
								Names: []*dst.Ident{},
								Type: &dst.Ident{
									Name: repo,
									Decs: dst.IdentDecorations{},
								},
								Tag:  nil,
								Decs: dst.FieldDecorations{},
							})
						}
					}
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
