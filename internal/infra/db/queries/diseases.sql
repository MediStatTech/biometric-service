-- name: GetDisease :one
SELECT disease_id, name, code, created_at, updated_at
FROM diseases
WHERE disease_id = $1
LIMIT 1;

-- name: GetDiseaseByCode :one
SELECT disease_id, name, code, created_at, updated_at
FROM diseases
WHERE code = $1
LIMIT 1;

-- name: ListDiseases :many
SELECT disease_id, name, code, created_at, updated_at
FROM diseases
ORDER BY name ASC;

-- name: CountDiseases :one
SELECT COUNT(*) FROM diseases;

-- SQL constants for mutations (used in repository)
-- name: CreateDisease :exec
INSERT INTO diseases (
    disease_id,
    name,
    code,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5);

-- name: UpdateDisease :exec
UPDATE diseases
SET
    name = $2,
    code = $3,
    updated_at = $4
WHERE disease_id = $1;

-- name: DeleteDisease :exec
DELETE FROM diseases
WHERE disease_id = $1;
