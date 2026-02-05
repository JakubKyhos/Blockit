-- name: CreateWhitelistDom :one
INSERT INTO whitelist (id, name)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: CreateWhitelistDomTemp :one
INSERT INTO whitelist (id, expires_at, name)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetWhitelistDom :one
SELECT * FROM whitelist WHERE name = $1 AND (expires_at IS NULL OR expires_at > now());

-- name: DeleteWhitelistDoms :exec
DELETE FROM whitelist;

-- name: DeleteWhitelistDom :exec
DELETE FROM whitelist WHERE name = $1;

-- name: DeleteWhitelistTempDoms :exec
DELETE FROM whitelist WHERE expires_at IS NOT NULL AND expires_at <= now();


-- name: GetWhitelistDoms :many
SELECT * FROM whitelist;
