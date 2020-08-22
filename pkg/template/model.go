package template

import (
	"bytes"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"go/token"
	"strconv"
)

const modelWhereUpdate = `
// TableName {{.Name}} in database table name
func ({{ToLowerCamel .Name}} *{{.Name}}) TableName() string {
	return "{{ToLowerCamel .Plural}}"
}

// Query{{.Name}} for repository where condition
type Query{{.Name}} struct {
	{{.Name}} {{.Name}}
	Base condition.Where
	Sorting condition.Sorting
}

// Where for repository where condition
func (option *Query{{.Name}}) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(option.{{.Name}})
	db = db.Scopes(option.Base.Where)
	db = db.Scopes(option.Sorting.Sort)
	return db
}
`

// AddModelWhereAndUpdate add where and update to model
func AddModelWhereAndUpdate(code []byte) string {

	f, err := decorator.Parse(code)
	if err != nil {
		fmt.Println("can't add where condition in model reason: ", err)
		return ""
	}

	importedPkg := 0
	hasGormImported := false
	hasConditionImported := false
	for i := 0; i < len(f.Decls); i++ {
		genDecl, ok := f.Decls[i].(*dst.GenDecl)
		if ok {
			for _, spec := range genDecl.Specs {
				imptSpec, ok := spec.(*dst.ImportSpec)
				if ok {
					importedPkg++
					if imptSpec.Path.Value == strconv.Quote("github.com/jinzhu/gorm") {
						hasGormImported = true
					}

					if imptSpec.Path.Value == strconv.Quote("github.com/karta0898098/kara/db/condition") {
						hasConditionImported = true
					}
				}
			}
		}
	}

	if importedPkg == 0 {
		//No any import
		if !hasGormImported {
			f.Decls = append(f.Decls, &dst.GenDecl{
				Tok:    token.IMPORT,
				Lparen: false,
				Specs: []dst.Spec{
					&dst.ImportSpec{
						Name: nil,
						Path: &dst.BasicLit{
							Kind:  token.STRING,
							Value: strconv.Quote("github.com/jinzhu/gorm"),
						},
					},
				},
				Rparen: false,
				Decs:   dst.GenDeclDecorations{},
			})
		}
		f.Decls[0], f.Decls[1] = f.Decls[1], f.Decls[0]

		for i := 0; i < len(f.Decls); i++ {
			genDecl, ok := f.Decls[i].(*dst.GenDecl)
			if ok {
				if genDecl.Tok == token.IMPORT {
					if !hasConditionImported {
						genDecl.Specs = append(genDecl.Specs, &dst.ImportSpec{
							Path: &dst.BasicLit{
								Kind:  token.STRING,
								Value: strconv.Quote("github.com/karta0898098/kara/db/condition"),
							},
						})
					}
				}
			}
		}
	} else {
		for i := 0; i < len(f.Decls); i++ {
			genDecl, ok := f.Decls[i].(*dst.GenDecl)
			if ok {
				if genDecl.Tok == token.IMPORT {
					if !hasGormImported {
						genDecl.Specs = append(genDecl.Specs, &dst.ImportSpec{
							Path: &dst.BasicLit{
								Kind:  token.STRING,
								Value: strconv.Quote("github.com/jinzhu/gorm"),
							},
						})
					}

					if !hasConditionImported {
						genDecl.Specs = append(genDecl.Specs, &dst.ImportSpec{
							Path: &dst.BasicLit{
								Kind:  token.STRING,
								Value: strconv.Quote("github.com/karta0898098/kara/db/condition"),
							},
						})
					}
				}
			}
		}
	}

	var buf bytes.Buffer
	err = decorator.Fprint(&buf, f)
	if err != nil {
		fmt.Println("can't add model where and update reason: ", err)
		return ""
	}

	buf.Write([]byte(modelWhereUpdate))

	return buf.String()
}
