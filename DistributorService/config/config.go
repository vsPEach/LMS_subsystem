package config

import (
	"time"
)

var (
	address, port string
)

type Config struct {
	Address   string
	Port      string
	Endpoints []string
	Timeout   time.Duration
}

func NewConfig() *Config {
	return &Config{}
}
