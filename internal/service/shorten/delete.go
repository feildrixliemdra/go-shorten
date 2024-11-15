package shorten

import (
	"context"
)

func (s service) Delete(ctx context.Context, code string) (err error) {
	return s.ShortenRepository.Delete(ctx, code)
}
