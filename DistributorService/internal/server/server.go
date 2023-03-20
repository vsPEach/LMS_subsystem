package server

import (
	"log"
	"net/http"
)

type Server struct {
	Port    string
	Address string
}

func (S *Server) Start() {
	err := http.ListenAndServe(S.Address+":"+S.Port, nil)
	if err != nil {
		log.Fatalf("Can't start server. Error: %s ", err)
	}
}

func NewServer(address, port string) *Server {
	return &Server{
		Address: address,
		Port:    port,
	}
}
