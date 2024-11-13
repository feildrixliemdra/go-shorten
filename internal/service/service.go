package service

import (
	"go-shorten/internal/config"
	"go-shorten/internal/repository"
	"go-shorten/internal/service/shorten"
)

type Service struct {
	ShortenService shorten.IShortenService
}

func InitiateService(cfg *config.Config, repository *repository.Repository) *Service {
	return &Service{
		ShortenService: shorten.NewShortenService(repository.ShortenRepository),
	}
}
