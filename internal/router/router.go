package router

import (
	"github.com/gin-gonic/gin"
	"go-shorten/internal/config"
	"go-shorten/internal/handler"
)

type router struct {
	rtr     *gin.Engine
	handler *handler.Handler
	cfg     *config.Config
}

func NewRouter(rtr *gin.Engine, cfg *config.Config, handler *handler.Handler) Router {
	return &router{
		rtr,
		handler,
		cfg,
	}
}

type Router interface {
	Init()
}

func (r *router) Init() {
	r.rtr.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.rtr.Group("/v1")
	shortenGroup := v1.Group("/shortens")

	shortenGroup.POST("", r.handler.ShortenHandler.CreateShorten)
	shortenGroup.GET("/:code", r.handler.ShortenHandler.GetShorten)
	shortenGroup.GET("/:code/stats", r.handler.ShortenHandler.GetShortenStats)
	//shortenGroup.DELETE("/:code", r.handler.ShortenHandler.Delete)
}
