package template

// GinRouter template ...
const GinRouter = `package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// SetRoute set http router & handler
func SetRoute(router *gin.Engine, handler *Handler) {
}

// SetGRPCService register gRPC handler
func SetGRPCService(server *grpc.Server, handler *Handler) {
}`

// EchoRouter template ...
const EchoRouter = `package handler

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

// SetRoute set http router & handler
func SetRoute(router *echo.Echo, handler *Handler) {
}

// SetGRPCService register gRPC handler
func SetGRPCService(server *grpc.Server, handler *Handler) {
}`