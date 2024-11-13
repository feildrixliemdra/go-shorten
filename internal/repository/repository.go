package repository

import (
	"github.com/jmoiron/sqlx"
	"go-shorten/internal/repository/shorten"
)

type Repository struct {
	ShortenRepository shorten.IShortenRepository
}

type Option struct {
	DB *sqlx.DB
}

func InitiateRepository(opt Option) *Repository {
	return &Repository{
		ShortenRepository: shorten.NewShortenRepository(opt.DB),
	}
}
