package template

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"html/template"
)

// TemplateMap ...
var Map = template.FuncMap{
	"ToLowerCamel":  ToLowerCamel,
	"ToCamel":       ToCamel,
	"ToPlural":      ToPlural,
	"ToLowerPlural": ToLowerPlural,
}

// ToLowerCamel ...
func ToLowerCamel(name string) string {
	return strcase.ToLowerCamel(name)
}

// ToCamel ...
func ToCamel(name string) string {
	return strcase.ToCamel(name)
}

func ToPlural(name string) string {
	return inflection.Plural(name)
}

func ToLowerPlural(name string) string {
	return inflection.Plural(strcase.ToLowerCamel(name))
}
