package main

import "github.com/vsPEach/LMS_subsystem/DistributorService/internal/server"

func main() {
	serv := server.NewServer("localhost", "8080")
	serv.Start()
}
