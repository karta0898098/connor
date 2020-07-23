package main

import (
	"fmt"
	"github.com/karta0898098/connor/cmd/conner/command"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "connor"
	app.Usage = "for create web app"
	app.Version = "1.0.1"
	app.Commands = command.New()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Sorry can't create project reason:", err)
	}
}
