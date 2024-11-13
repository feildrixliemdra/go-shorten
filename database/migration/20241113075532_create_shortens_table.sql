-- +goose Up
-- +goose StatementBegin
CREATE TABLE shortens (
    id bigserial PRIMARY KEY,
    original_url text NOT NULL,
    short_url text NOT NULL,
    click_count int DEFAULT 0,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL   

);
CREATE UNIQUE INDEX idx_shortens_short_url ON shortens(short_url);
-- +goose StatementEnd