package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"io/ioutil"
	"os/exec"
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

// FindProjectName ...
func FindProjectName() string {
	var outBuf bytes.Buffer
	process := exec.Command("go", "mod", "edit", "-json")
	process.Stdout = &outBuf
	process.Run()

	var data map[string]interface{}

	_ = json.Unmarshal(outBuf.Bytes(), &data)

	module, _ := data["Module"].(map[string]interface{})

	return module["Path"].(string)
}

// FindHttpEngine ...
func FindHttpEngine() string {

	var outBuf bytes.Buffer
	process := exec.Command("go", "mod", "edit", "-json")
	process.Stdout = &outBuf
	process.Run()

	var data struct {
		Module struct {
			Path string `json:"Path"`
		} `json:"Module"`
		Go      string `json:"Go"`
		Require []struct {
			Path     string `json:"Path"`
			Version  string `json:"Version"`
			Indirect bool   `json:"Indirect,omitempty"`
		} `json:"Require"`
		Exclude interface{} `json:"Exclude"`
		Replace interface{} `json:"Replace"`
	}

	_ = json.Unmarshal(outBuf.Bytes(), &data)

	for _, pkg := range data.Require {
		if pkg.Path == "github.com/gin-gonic/gin" {
			return "gin"
		}

		if pkg.Path == "github.com/labstack/echo/v4" {
			return "echo"
		}
	}

	return "echo"
}
