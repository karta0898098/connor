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

// Where{{.Name}} for repository where condition
type Where{{.Name}} struct {
	{{.Name}} {{.Name}}
}

// Update{{.Name}} for repository update condition
type Update{{.Name}} struct {
	Where Where{{.Name}}
	{{.Name}}  {{.Name}}
}

// Where for repository where condition
func (where *Where{{.Name}}) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.{{.Name}})
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
	hasImported := false
	for i := 0; i < len(f.Decls); i++ {
		genDecl, ok := f.Decls[i].(*dst.GenDecl)
		if ok {
			for _, spec := range genDecl.Specs {
				imptSpec, ok := spec.(*dst.ImportSpec)
				if ok {
					importedPkg++
					if imptSpec.Path.Value == strconv.Quote("github.com/jinzhu/gorm") {
						hasImported = true
					}
				}
			}
		}
	}

	if !hasImported {
		if importedPkg == 0 {
			//No any import
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
			f.Decls[0], f.Decls[1] = f.Decls[1], f.Decls[0]
		} else {
			for i := 0; i < len(f.Decls); i++ {
				genDecl, ok := f.Decls[i].(*dst.GenDecl)
				if ok {
					if genDecl.Tok == token.IMPORT {
						genDecl.Specs = append(genDecl.Specs, &dst.ImportSpec{
							Path: &dst.BasicLit{
								Kind:  token.STRING,
								Value: strconv.Quote("github.com/jinzhu/gorm"),
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
