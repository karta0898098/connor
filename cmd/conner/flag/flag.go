package flag

import "github.com/urfave/cli/v2"

// NewInit 指令集
func NewInit() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Required: true,
			Usage:    "project name",
		},
	}
}

// NewAddController 新增controller
func NewAddController() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Required: true,
			Usage:    "controller name",
		},
	}
}

// NewAddEntity 新增entity
func NewAddEntity() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Required: true,
			Usage:    "input file path",
		},
		&cli.BoolFlag{
			Name:     "repository",
			Aliases:  []string{"repo"},
			Required: false,
			Usage:    "add repository",
		},
		&cli.BoolFlag{
			Name:     "service",
			Aliases:  []string{"srv"},
			Required: false,
			Usage:    "add service",
		},
	}
}
