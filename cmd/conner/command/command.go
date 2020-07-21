package command

import (
	"github.com/karta0898098/connor/cmd/conner/flag"
	"github.com/urfave/cli/v2"
)

// New for cli command
func New() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "init",
			Usage:  "init default project scaffold",
			Flags:  flag.NewInit(),
			Action: Init,
		},
		{
			Name:        "add",
			Usage:       "add you need file to project",
			Subcommands: NewAddCommand(),
		},
	}
}
