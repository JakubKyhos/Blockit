-- name: CreateDomain :one
INSERT INTO domains (id, name)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetDomain :one
SELECT * FROM domains WHERE name = $1;

-- name: DeleteDomains :exec
DELETE FROM domains;

-- name: GetDomains :many
SELECT * FROM domains;

-- name: DomainBlockState :one
UPDATE domains
SET is_blocked = $1,
updated_at = NOW()
WHERE name = $2
RETURNING *;

-- name: DomainsBlockedStateGlobal :many
UPDATE domains
SET is_blocked = $1,
updated_at = NOW()
RETURNING *;