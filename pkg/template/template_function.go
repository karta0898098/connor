package template

import (
	"github.com/iancoleman/strcase"
	"html/template"
)

// TemplateMap ...
var Map = template.FuncMap{
	"ToLowerCamel": ToLowerCamel,
	"ToCamel":      ToCamel,
}

// ToLowerCamel ...
func ToLowerCamel(name string) string {
	return strcase.ToLowerCamel(name)
}

// ToCamel ...
func ToCamel(name string) string {
	return strcase.ToCamel(name)
}
