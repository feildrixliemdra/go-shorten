package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-shorten/internal/config"
	"go-shorten/internal/handler"
	"go-shorten/internal/router"
)

func InitiateGinRouter(cfg *config.Config, handler *handler.Handler) *gin.Engine {
	r := gin.Default()

	//init router
	route := router.NewRouter(r, cfg, handler)
	route.Init()

	gin.SetMode(cfg.App.ReleaseMode)

	return r
}
