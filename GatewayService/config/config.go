package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type (
	Config struct {
		Logger    LoggerConf    `mapstructure:"logger"`
		Server    ServerConf    `mapstructure:"server"`
		Database  DatabaseConf  `mapstructure:"database"`
		Endpoints EndpointsConf `mapstructure:"urls"`
	}

	LoggerConf struct {
		Level      zapcore.Level `mapstructure:"level" default:"debug"`
		Encoding   string        `mapstructure:"encoding" default:"console"`
		OutputPath []string      `mapstructure:"output" default:"stdout"`
	}

	ServerConf struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}

	EndpointsConf struct {
		py string `mapstructure:"py" default:""`
		js string `mapstructure:"js" default:""`
	}

	DatabaseConf struct {
		Username string `mapstructure:"username" default:"postgres"`
		Password string `mapstructure:"password" default:"postgres"`
		Host     string `mapstructure:"host" default:"localhost"`
		Port     string `mapstructure:"port" default:"5432"`
		Name     string `mapstructure:"name" default:"postgres"`
		sslMode  string `mapstructure:"ssl_mode" default:"disable"`
	}
)

func NewConfig(path string) (Config, error) {
	viper.SetConfigFile(path)
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}

func (db DatabaseConf) GetConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}
