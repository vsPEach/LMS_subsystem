package app

import (
	"github.com/pkg/errors"
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/repository"
	HTTP "github.com/vsPEach/LMS_subsystem/DistributorService/internal/server"
	logger "github.com/vsPEach/LMS_subsystem/DistributorService/pkg"
	"net/http"
	"sync"
)

type Logger interface {
	Error(...any)
	Info(...any)
	Warn(...any)
	Infof(template string, args ...any)
	Errorf(template string, args ...any)
}

func Run(conf config.Config) {
	logg := logger.New(conf.Logger)
	db := repository.NewDatabase()
	server := HTTP.NewServer(db, logg)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.Connect(); err != nil {
			logg.Error(err)
		}
		logg.Info("connected to database")
	}()

	logg.Info("service start")
	wg.Wait()
}
