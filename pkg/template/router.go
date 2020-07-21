package template

// Router template ...
const Router = `package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// SetRouter set http router & handler
func SetRouter(router *gin.Engine, handler *Handler) {
}

// SetGRPCService register gRPC handler
func SetGRPCService(server *grpc.Server, handler *Handler) {
}`
