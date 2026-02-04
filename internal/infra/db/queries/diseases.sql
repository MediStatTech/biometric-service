-- name: GetDisease :one
SELECT disease_id, name, updated_at, created_at
FROM diseases
WHERE disease_id = $1
LIMIT 1;

-- name: ListDiseases :many
SELECT disease_id, name, updated_at, created_at
FROM diseases
ORDER BY name;

-- name: CountDiseases :one
SELECT COUNT(*) FROM diseases;

-- SQL constants for mutations (used in repository)
-- name: CreateDisease :exec
INSERT INTO diseases (
    disease_id,
    name,
    created_at
) VALUES ($1, $2, $3);

-- name: UpdateDisease :exec
UPDATE diseases
SET
    name = $2,
    updated_at = $3
WHERE disease_id = $1;

-- name: DeleteDisease :exec
DELETE FROM diseases
WHERE disease_id = $1;
