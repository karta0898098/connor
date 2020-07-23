package builder

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"os"
	"os/exec"
	"path"
)

// FolderBuilder for build folder structure
type FolderBuilder struct {
	workingDir  string
	projectName string
	packages    []string
	folders     []string
}

// NewFolderBuilder new FolderBuilder constructor
func NewFolderBuilder(workingDir string) *FolderBuilder {
	return &FolderBuilder{workingDir: workingDir}
}

// ProjectName set project name
func (f *FolderBuilder) ProjectName(name string) *FolderBuilder {
	f.projectName = name
	return f
}

// Package set want to add package to go mode
func (f *FolderBuilder) Package(name string) *FolderBuilder {
	f.packages = append(f.packages, name)
	return f
}

// Packages set want to add packages to go mode
func (f *FolderBuilder) Packages(names []string) *FolderBuilder {
	f.packages = append(f.packages, names...)
	return f
}

// Folder set want to add folder for create
func (f *FolderBuilder) Folder(name string) *FolderBuilder {

	absName := ""
	if name != f.projectName {
		absName = path.Join(f.projectName, name)
	}

	f.folders = append(f.folders, absName)
	return f
}

// Folders set want to add folders for create
func (f *FolderBuilder) Folders(names []string) *FolderBuilder {
	absNames := make([]string, 0)

	for _, name := range names {
		if name != f.projectName {
			absNames = append(absNames, path.Join(f.projectName, name))
		}
	}

	f.folders = append(f.folders, absNames...)
	return f
}

// Build ...
func (f *FolderBuilder) Build() {
	err := os.Mkdir(f.projectName, os.ModePerm)

	for i := 0; i < len(f.folders); i++ {
		err := os.MkdirAll(f.folders[i], os.ModePerm)
		if err != nil {
			fmt.Println("can't create folder reason: ", f.folders[i])
		}
	}

	if len(f.packages) > 0 {
		//執行go mod
		cmd := exec.Command("go", "mod", "init", f.projectName)
		cmd.Dir = f.workingDir
		err = cmd.Run()

		if err != nil {
			fmt.Println("can't exec go mod reason: ", err)
		}

		bar := pb.StartNew(len(f.packages))
		for i := 0; i < len(f.packages); i++ {
			process := exec.Command("go", "get", f.packages[i])
			process.Dir = f.workingDir
			process.Run()
			bar.Increment()
		}
		bar.Finish()

		process := exec.Command("go", "mod", "tidy")
		process.Dir = f.workingDir
		process.Run()

		fmt.Println("get go mod package finish")
	}
}
