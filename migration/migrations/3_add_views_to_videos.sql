-- +migrate Up
ALTER TABLE videos ADD COLUMN views BIGINT NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE videos DROP COLUMN views;
