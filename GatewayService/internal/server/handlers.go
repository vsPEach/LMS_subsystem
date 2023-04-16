package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/DTO"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server/requests"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Handler struct {
	logg *zap.SugaredLogger
	conf config.EndpointsConf
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()
	routes.POST("/api/", h.Redirect)
	return routes
}

func (h *Handler) Redirect(ctx *gin.Context) {
	var item DTO.Item
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.logg.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
	}
	err = json.Unmarshal(data, &item)
	if err != nil {
		h.logg.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
	}
	h.distribute(item.Files)
}

func NewHandler(logger *zap.SugaredLogger, conf config.EndpointsConf) *Handler {
	return &Handler{
		logg: logger,
		conf: conf,
	}
}

func (h *Handler) distribute(files []DTO.File) {
	python, js := make([]DTO.File, 0, 10), make([]DTO.File, 0, 10)
	for _, file := range files {
		switch file.Lang {
		case "py":
			python = append(python, file)
		case "js":
			js = append(js, file)
		}
	}
	go h.sender(python, "https://eoalsyff94oteab.m.pipedream.net/")
	go h.sender(js, "https://eoalsyff94oteab.m.pipedream.net/")
}

func (h *Handler) sender(files []DTO.File, url string) {
	bytes, err := json.Marshal(files)
	if err != nil {
		h.logg.Error(err)
	}
	err = requests.Request(bytes, url)
	if err != nil {
		h.logg.Error()
	}
}
