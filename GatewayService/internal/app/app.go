package app

import (
	"github.com/google/uuid"
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/DTO"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server"
	logger "github.com/vsPEach/LMS_subsystem/DistributorService/pkg"
)

type repo interface {
	Create(file DTO.File) error
	Read(uuid uuid.UUID) error
	ReadAll(uuid uuid.UUID) error
	Update(uuid uuid.UUID) error
	Delete(uuid2 uuid.UUID) error
}

func Run(conf config.Config) {
	logg := logger.New(conf.Logger)
	serv := server.NewServer(conf.Server, server.NewHandler(logg, conf.Endpoints))
	if err := serv.Start(); err != nil {
		logg.Error("Can't start server")
	}
	logg.Info("Server start...")
}
