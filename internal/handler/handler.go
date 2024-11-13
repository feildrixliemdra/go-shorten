package handler

import (
	"go-shorten/internal/config"
	"go-shorten/internal/handler/shorten"
	"go-shorten/internal/service"
)

type Handler struct {
	ShortenHandler shorten.Handler
}

func InitiateHandler(cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		ShortenHandler: shorten.NewShortenHandler(cfg, services.ShortenService),
	}
}
