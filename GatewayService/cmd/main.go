package main

import (
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server"
)

func main() {
	cfg := config.NewConfig("")
	handler := new(server.Handler)
	serv := server.NewServer(cfg.Port, handler)
	serv.Start()
}
