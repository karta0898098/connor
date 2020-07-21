package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"

	appTemplate "github.com/karta0898098/connor/pkg/template"
)

// CodeBuilder to write code file
type CodeBuilder struct {
	Data        interface{}
	Template    string
	ProjectName string
	Path        string
	File        string
}

// Build create source code
func (c *CodeBuilder) Build() {

	var buf bytes.Buffer
	t, err := template.New(c.File).Funcs(appTemplate.Map).Parse(c.Template)
	if err != nil {
		fmt.Printf("can't write:%v reason:%v\n", c.File, err)
		return
	}
	err = t.Execute(&buf, c.Data)

	fileName := path.Join(c.ProjectName, c.Path, c.File)
	err = ioutil.WriteFile(fileName, buf.Bytes(), os.ModePerm)
	if err != nil {
		fmt.Printf("can't write:%v reason:%v\n", c.File, err)
		return
	}

	fmt.Printf("create file %s success \n", c.File)
}
