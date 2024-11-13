package model

import "time"

type Shorten struct {
	ID          uint64    `db:"id"`
	OriginalURL string    `db:"original_url"`
	ShortURL    string    `db:"short_url"`
	ClickCount  uint64    `db:"click_count"`
	CreatedAt   time.Time `db:"created_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}

func (t Shorten) TableName() string {
	return "shortens"
}
