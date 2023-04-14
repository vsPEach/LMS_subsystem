package server

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/DTO"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server/requests"
	"io"
	"strings"
)

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()
	routes.POST("/api/", h.Redirect)
	return routes
}

func (*Handler) Redirect(ctx *gin.Context) {
	var item DTO.Item
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(500, err)
	}
	err = json.Unmarshal(data, &item)
	if err != nil {
		ctx.JSON(500, err)
	}
	go distribute(item.ToStringSlice())
}

func distribute(body []string) {
	python := bytes.Buffer{}
	js := bytes.Buffer{}
	for _, s := range body {
		for _, lang := range strings.Split(s, ",") {
			switch strings.Split(lang, ":")[1] {
			case "py":
				python.Write([]byte(s + "\n"))
			case "js":
				js.Write([]byte(s))
			}
		}
	}
	err := requests.Request(python.Bytes(), "https://eoyg9isams0wupi.m.pipedream.net")
	if err != nil {
		return
	}
}
