package main

import (
	"fmt"
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg)
	serv := server.NewServer(cfg.Address, cfg.Port)
	serv.Start()
}
