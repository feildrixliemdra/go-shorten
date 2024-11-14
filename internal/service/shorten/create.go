package shorten

import (
	"context"
	"go-shorten/internal/model"
	"go-shorten/internal/payload"
	"go-shorten/internal/util"
)

func (s service) Create(ctx context.Context, req payload.CreateShortenRequest) (res payload.CreateShortenResponse, err error) {
	shortURL := util.ShortenURL(req.OriginalURL)

	shortenModel := model.Shorten{
		OriginalURL: req.OriginalURL,
		ShortURL:    shortURL,
	}

	err = s.ShortenRepository.Create(ctx, shortenModel)
	if err != nil {
		return res, err
	}

	res = payload.CreateShortenResponse{
		OriginalURL: req.OriginalURL,
		ShortURL:    shortURL,
	}

	return
}
