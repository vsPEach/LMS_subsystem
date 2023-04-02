package server

import (
	"github.com/gin-gonic/gin"
	"io"
)

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()
	routes.POST("/api", func(context *gin.Context) {
		jsonData, err := io.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, "can't read file")
		}
		context.JSON(200, "OK!")
	})
	return routes
}
