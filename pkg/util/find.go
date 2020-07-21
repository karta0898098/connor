package util

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"io/ioutil"
)

// FindModelName ...
func FindModelName(path string) string {

	code, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("can't find model name reason: ", err)
		return ""
	}

	f, err := decorator.Parse(code)
	if err != nil {
		fmt.Println("can't find model name reason: ", err)
		return ""
	}

	for i := 0; i < len(f.Decls); i++ {
		specsTree, ok := f.Decls[i].(*dst.GenDecl)
		if ok {
			for _, spec := range specsTree.Specs {
				typeSpec, ok := spec.(*dst.TypeSpec)
				if ok {
					_, ok := typeSpec.Type.(*dst.StructType)
					if ok {
						return typeSpec.Name.Name
					}
				}
			}
		}
	}

	return ""
}
