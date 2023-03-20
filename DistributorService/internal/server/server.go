package server

import "net/http"

type Server struct {
	Port    string
	Address string
}

func (S *Server) Start() {
	err := http.ListenAndServe(S.Address+":"+S.Port, nil)
	if err != nil {
		return
	}
}

func NewServer(address, port string) *Server {
	return &Server{
		Address: address,
		Port:    port,
	}
}
