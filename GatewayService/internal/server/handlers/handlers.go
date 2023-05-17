package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/models"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/repository"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/server/middlewares"
	"github.com/vsPEach/LMS_subsystem/DistributorService/pkg/utils"
	"log"
	"net/http"
)

type Logger interface {
	Error(...any)
	Info(...any)
	Warn(...any)
	Infof(template string, args ...any)
	Errorf(template string, args ...any)
}

type HTTPHandler struct {
	engine   *gin.Engine
	database *repository.Database
	logger   Logger
}

func NewHTTPHandler(database *repository.Database, logger Logger) *HTTPHandler {
	return &HTTPHandler{engine: gin.New(), database: database, logger: logger}
}

func (h *HTTPHandler) Routes() *gin.Engine {
	h.engine.Use(middlewares.Logging(h.logger))
	api := h.engine.Group("/auth")
	api.POST("/signin", h.Login)
	api.POST("/signup", h.Register)
	secured := h.engine.Group("/api").Use(middlewares.Auth())
	secured.GET("/ex1", h.Redirect)
	secured.GET("/ex2", h.Redirect)
	log.Print(h.engine.Routes())
	return h.engine
}

func (h *HTTPHandler) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	err := h.database.Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	tokenString, err := utils.GenerateJWT(user.Email, user.Password, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *HTTPHandler) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	err := h.database.Read(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	tokenString, err := utils.GenerateJWT(user.Email, user.Password, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *HTTPHandler) Redirect(ctx *gin.Context) {

}
