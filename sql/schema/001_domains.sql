-- +goose Up
CREATE TABLE domains (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    name TEXT NOT NULL UNIQUE,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE domains;