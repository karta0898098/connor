package template

// Server template ...
const Server = `package main

import (
	"github.com/karta0898098/kara/grpc"
	"github.com/karta0898098/kara/http"
	"github.com/karta0898098/kara/zlog"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"{{.ProjectName}}/configs"
	"{{.ProjectName}}/pkg/handler"
	"{{.ProjectName}}/pkg/repository"
	"{{.ProjectName}}/pkg/service"
)

// NewServerCommand new lunch server app command
func NewServerCommand() *cobra.Command {
	command := &cobra.Command{
		Run: run,
		Use: "server",
	}
	command.Flags().StringP("config", "c", "", "server config path")
	return command
}

func run(cmd *cobra.Command, args []string) {
	configs.Path, _ = cmd.LocalFlags().GetString("config")

	app := fx.New(
		configs.Module,
		http.Module,
		grpc.Module,
		repository.Module,
		service.Module,
		handler.Module,
		fx.Invoke(zlog.Setup),
		fx.Invoke(handler.SetRouter),
		fx.Invoke(handler.SetGRPCService),
		fx.Invoke(http.RunGin),
		fx.Invoke(grpc.RunGRPC),
	)

	app.Run()
}
`
