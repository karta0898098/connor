package template

import (
	"bytes"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"io/ioutil"
)

// ControllerModule template ...
const ControllerModule = `package controller

import (
	"go.uber.org/fx"
)

// Module for export controllers to fx injection
var Module = fx.Provide()
`

// GinController template ...
const GinController = `package controller

import (
	"github.com/gin-gonic/gin"
)

// {{.Name}}Controller ...
type {{.Name}}Controller struct{
}

// New{{.Name}}Controller new constructor
func New{{.Name}}Controller() *{{.Name}}Controller{
	return &{{.Name}}Controller{}
}

// Get ...
func (controller *{{.Name}}Controller) Get(c *gin.Context){
	panic("please implement")
}

// List ...
func (controller *{{.Name}}Controller) List(c *gin.Context){
	panic("please implement")
}

// Create ...
func (controller *{{.Name}}Controller) Create(c *gin.Context){
	panic("please implement")
}

// Update ...
func (controller *{{.Name}}Controller) Update(c *gin.Context){
	panic("please implement")
}

// Delete ...
func (controller *{{.Name}}Controller) Delete(c *gin.Context){
	panic("please implement")
}
`

// EchoController template ...
const EchoController =  `package controller

import (
	"github.com/labstack/echo/v4"
)

// {{.Name}}Controller ...
type {{.Name}}Controller struct{
}

// New{{.Name}}Controller new constructor
func New{{.Name}}Controller() *{{.Name}}Controller{
	return &{{.Name}}Controller{}
}

// Get ...
func (controller *{{.Name}}Controller) Get(c echo.Context) error{
	panic("please implement")
}

// List ...
func (controller *{{.Name}}Controller) List(c echo.Context) error{
	panic("please implement")
}

// Create ...
func (controller *{{.Name}}Controller) Create(c echo.Context) error{
	panic("please implement")
}

// Update ...
func (controller *{{.Name}}Controller) Update(c echo.Context) error{
	panic("please implement")
}

// Delete ...
func (controller *{{.Name}}Controller) Delete(c echo.Context) error{
	panic("please implement")
}
`


// AddControllerInModule add controller module
func AddControllerInModule(controller string) string {

	code, _ := ioutil.ReadFile("./pkg/handler/controller/module.go")

	f, err := decorator.Parse(code)
	if err != nil {
		fmt.Println("can't add controller in module reason: ", err)
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
			call.Args = append(call.Args, dst.NewIdent(controller))
			for _, v := range call.Args {
				v := v.(*dst.Ident)
				v.Decs.Before = dst.NewLine
				v.Decs.After = dst.NewLine
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
