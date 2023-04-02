package server

//TODO: create config

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	server  *http.Server
	port    string
	handler *gin.Engine
}

func NewServer(port string, handler *Handler) *Server {
	return &Server{port: port,
		handler: handler.InitRoutes()}
}

func (S *Server) Start() {
	S.server = &http.Server{
		Addr:    "localhost:8080",
		Handler: S.handler,
	}
	err := S.server.ListenAndServe()
	if err != nil {
		log.Fatalf("Can't start server. Error: %s ", err)
	}
}

func (S *Server) Stop(ctx context.Context) error {
	return S.server.Shutdown(ctx)
}
