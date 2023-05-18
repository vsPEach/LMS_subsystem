package server

import (
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/repository"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server/handlers"
	"net/http"
)

type Logger interface {
	Error(...any)
	Info(...any)
	Warn(...any)
	Infof(template string, args ...any)
	Errorf(template string, args ...any)
}

type Server struct {
	server http.Server
	logg   Logger
}

func NewServer(database *repository.Database, logger Logger) *Server {
	h := handlers.NewHTTPHandler(database, logger)
	return &Server{server: http.Server{
		Addr:    "https://gateway-service-bx88.onrender.com:80",
		Handler: h.Routes(),
	}}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
