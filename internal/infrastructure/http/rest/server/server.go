package server

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

	v1 := g.Group("/v1")

	users := v1.Group("/users")
	users.POST("", cfg.Controller.CreateUser)
	users.GET("", cfg.Controller.ListUsers)
	users.PUT("/:userID", cfg.Controller.UpdateUser)
	users.DELETE("/:userID", cfg.Controller.DeleteUser)
	users.GET("/:userID", cfg.Controller.GetUser)

	v2 := g.Group("/v2")

	usersv2 := v2.Group("/users")
	usersv2.POST("", cfg.Controller.CreateUser)

	return &http.Server{
		Addr:    cfg.Addr,
		Handler: g,
	}
}
