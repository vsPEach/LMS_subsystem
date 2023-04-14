package app

import (
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server"
)

func Run(conf config.Config) {
	handler := new(server.Handler)
	serv := server.NewServer(conf.Server, handler)
	serv.Start()
}
