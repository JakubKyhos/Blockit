-- name: CreateDomain :one
INSERT INTO domains (id, created_at, updated_at, name, is_blocked)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetDomain :one
SELECT * FROM domains WHERE name = $1;

-- name: DeleteDomains :exec
DELETE FROM domains;

-- name: GetDomains :many
SELECT * FROM domains;