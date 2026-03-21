-- name: CreateBlacklistDom :one
INSERT INTO blacklist (id, name)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetBlacklistDom :one
SELECT * FROM blacklist WHERE name = $1;

-- name: DeleteBlacklistDoms :exec
DELETE FROM blacklist;

-- name: DeleteBlacklistDom :exec
DELETE FROM blacklist WHERE name = $1;

-- name: GetBlacklistDoms :many
SELECT * FROM blacklist;
