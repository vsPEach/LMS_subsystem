package config

import "flag"

var (
	address, port string
)

func flagSetup() {
	flag.StringVar(&address, "address", "localhost", "Service address")
	flag.StringVar(&port, "port", "8080", "Service port")
}

type Config struct {
	Address string
	Port    string
}

func NewConfig() *Config {
	flag.Parse()
	return &Config{
		Address: address,
		Port:    port,
	}
}
