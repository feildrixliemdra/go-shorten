package shorten

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-shorten/internal/model"
)

type IShortenRepository interface {
	Create(ctx context.Context, shortenModel model.Shorten) (err error)
	GetByCode(ctx context.Context, code string) (res model.Shorten, err error)
	GetStats(ctx context.Context, code string) (res model.Shorten, err error)
	Delete(ctx context.Context, code string) (err error)
}

type repository struct {
	DB *sqlx.DB
}

func NewShortenRepository(db *sqlx.DB) IShortenRepository {
	return &repository{
		DB: db,
	}
}

func (r repository) Create(ctx context.Context, shortenModel model.Shorten) (err error) {
	query := `INSERT INTO 
    			shortens (original_url, short_url)
    			VALUES ($1, $2)`
	_, err = r.DB.ExecContext(ctx, query, shortenModel.OriginalURL, shortenModel.ShortURL)

	return err
}

func (r repository) GetByCode(ctx context.Context, code string) (res model.Shorten, err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetStats(ctx context.Context, code string) (res model.Shorten, err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, code string) (err error) {
	//TODO implement me
	panic("implement me")
}
