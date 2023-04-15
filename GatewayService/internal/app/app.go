package app

import (
	"github.com/google/uuid"
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	logger "github.com/vsPEach/LMS_subsystem/DistributorService/internal"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/DTO"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server"
)

type repo interface {
	Create(file DTO.File) error
	Read(uuid uuid.UUID) error
	ReadAll(uuid uuid.UUID) error
	Update(uuid uuid.UUID) error
	Delete(uuid2 uuid.UUID)
}

func Run(conf config.Config) {
	logg := logger.New(conf.Logger)
	handler := new(server.Handler)
	serv := server.NewServer(conf.Server, handler)
	_ = serv.Start()
	logg.Info("Server start...")
}
