-- +goose Up
CREATE TABLE whitelist (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    expires_at TIMESTAMP NULL,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE whitelist;