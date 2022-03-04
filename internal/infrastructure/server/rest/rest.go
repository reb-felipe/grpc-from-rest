package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	ListUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
}

type Config struct {
	Controller Controller
	Addr       string
}

func NewServer(cfg *Config) *http.Server {
	g := gin.Default()

	users := g.Group("/users")
	users.POST("", cfg.Controller.CreateUser)
	users.GET("", cfg.Controller.ListUsers)
	users.PUT("/:userID", cfg.Controller.UpdateUser)
	users.DELETE("/:userID", cfg.Controller.DeleteUser)
	users.GET("/:userID", cfg.Controller.GetUser)

	return &http.Server{
		Addr:    cfg.Addr,
		Handler: g,
	}
}
