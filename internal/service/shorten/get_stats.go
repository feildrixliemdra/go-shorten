package shorten

import (
	"context"
	"database/sql"
	"errors"
	"go-shorten/internal/constant/errormsg"
	"go-shorten/internal/payload"
)

func (s service) GetStats(ctx context.Context, code string) (res payload.GetStatsResponse, err error) {
	shorten, err := s.ShortenRepository.GetStats(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, errormsg.ErrShortenNotFound
		}
		return res, err
	}

	res.OriginalURL = shorten.OriginalURL
	res.ClickCount = shorten.ClickCount

	return
}
