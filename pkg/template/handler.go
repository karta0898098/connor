package template

import (
	"bytes"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"io/ioutil"
	"strings"
)

// Handler template ...
const Handler = `package handler

import (
	"go.uber.org/fx"
	"{{.ProjectName}}/pkg/handler/controller"
)

// Module for export handler to fx injection
var Module = fx.Options(
	fx.Provide(NewHandler),
	controller.Module,
)

// Handler handle all controller
type Handler struct {
}

// Params for handler new constructor
type Params struct {
	fx.In
}

// NewHandler handler new constructor
func NewHandler(h Params) *Handler {
	return &Handler{}
}`

// AddHandlerModule add controller to handler
func AddHandlerModule(controller string) string {

	code, _ := ioutil.ReadFile("./pkg/handler/handler.go")
	f, err := decorator.Parse(code)
	if err != nil {
		fmt.Println("can't add controller in module reason: ", err)
		return ""
	}

	for i := 0; i < len(f.Decls); i++ {
		specsTree, ok := f.Decls[i].(*dst.GenDecl)
		if ok {
			for _, spec := range specsTree.Specs {
				typeSpec, ok := spec.(*dst.TypeSpec)
				if ok {
					if typeSpec.Name.Name == "Handler" {
						structType, ok := typeSpec.Type.(*dst.StructType)
						if ok {
							structType.Fields.List = append(structType.Fields.List,
								&dst.Field{
									Names: []*dst.Ident{
										{
											Name: strings.Title(controller),
											Obj:  nil,
											Path: "",
											Decs: dst.IdentDecorations{},
										},
									},
									Type: &dst.SelectorExpr{
										X: &dst.Ident{
											Name: "*controller",
										},
										Sel:  dst.NewIdent(strings.Title(controller) + "Controller"),
										Decs: dst.SelectorExprDecorations{},
									},
									Tag:  nil,
									Decs: dst.FieldDecorations{},
								},
							)
						}
					}

					if typeSpec.Name.Name == "Params" {
						structType, ok := typeSpec.Type.(*dst.StructType)
						if ok {
							structType.Fields.List = append(structType.Fields.List,
								&dst.Field{
									Names: []*dst.Ident{
										{
											Name: strings.Title(controller),
											Obj:  nil,
											Path: "",
											Decs: dst.IdentDecorations{},
										},
									},
									Type: &dst.SelectorExpr{
										X: &dst.Ident{
											Name: "*controller",
										},
										Sel:  dst.NewIdent(strings.Title(controller) + "Controller"),
										Decs: dst.SelectorExprDecorations{},
									},
									Tag:  nil,
									Decs: dst.FieldDecorations{},
								},
							)
						}
					}
				}
			}
		}
		funcSpec, ok := f.Decls[i].(*dst.FuncDecl)
		if ok {

			for i := 0; i < len(funcSpec.Body.List); i++ {
				rn, ok := funcSpec.Body.List[i].(*dst.ReturnStmt)
				if ok {
					un := rn.Results[0].(*dst.UnaryExpr)
					c := un.X.(*dst.CompositeLit)
					c.Elts = append(c.Elts, &dst.KeyValueExpr{
						Key:   dst.NewIdent(strings.Title(controller)),
						Value: dst.NewIdent("h." + strings.Title(controller)),
					})
				}
			}
		}
	}

	var buf bytes.Buffer
	err = decorator.Fprint(&buf, f)
	if err != nil {
		fmt.Println("can't add controller in module reason: ", err)
		return ""
	}

	return buf.String()
}
