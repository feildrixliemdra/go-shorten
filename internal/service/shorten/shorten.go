package shorten

import (
	"context"
	"go-shorten/internal/payload"
	shortenRepo "go-shorten/internal/repository/shorten"
)

type IShortenService interface {
	Create(ctx context.Context, req payload.CreateShortenRequest) (res payload.CreateShortenResponse, err error)
	GetByCode(ctx context.Context, code string) (res payload.GetShortenResponse, err error)
	GetStats(ctx context.Context, code string) (res payload.GetStatsResponse, err error)
	Delete(ctx context.Context, code string) (err error)
}

type service struct {
	ShortenRepository shortenRepo.IShortenRepository
}

func NewShortenService(shortenRepository shortenRepo.IShortenRepository) IShortenService {
	return &service{
		ShortenRepository: shortenRepository,
	}
}
