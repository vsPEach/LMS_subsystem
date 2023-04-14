package main

import (
	"flag"
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/app "
	"log"
)

var (
	path string
)

func init() {
	flag.StringVar(&path, "path", "./config/config.yml", "")
}

func main() {
	cfg, err := config.NewConfig("")
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
