package template

// Main template ...
const Main = `package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {
	root := &cobra.Command{}
	root.AddCommand(NewServerCommand())
	if err := root.Execute(); err != nil {
		os.Exit(-1)
	}
}
`
