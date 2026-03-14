-- +goose Up
CREATE TABLE blacklist (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE blacklist;