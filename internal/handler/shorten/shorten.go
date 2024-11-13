package shorten

import (
	"go-shorten/internal/config"
	"go-shorten/internal/service/shorten"
)

type Handler struct {
	cfg            *config.Config
	shortenService shorten.IShortenService
}

func NewShortenHandler(cfg *config.Config, service shorten.IShortenService) Handler {
	return Handler{
		cfg:            cfg,
		shortenService: service,
	}
}
