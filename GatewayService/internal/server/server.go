package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	config  config.ServerConf
	logg    zap.SugaredLogger
	server  *http.Server
	handler *gin.Engine
}

func NewServer(conf config.ServerConf, handler *Handler) *Server {
	return &Server{config: conf,
		handler: handler.InitRoutes()}
}

func (S *Server) Start() error {
	S.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", S.config.Host, S.config.Port),
		Handler: S.handler,
	}
	if err := S.server.ListenAndServe(); err != nil {
		return errors.New("Can't start server. Error: %s " + err.Error())
	}
	return nil
}

func (S *Server) Stop(ctx context.Context) error {
	return S.server.Shutdown(ctx)
}
