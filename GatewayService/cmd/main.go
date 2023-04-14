package main

import (
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/app "
	"log"
)

func main() {
	cfg, err := config.NewConfig("")
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
