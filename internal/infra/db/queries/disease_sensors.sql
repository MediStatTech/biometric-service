-- name: GetDiseaseSensor :one
SELECT disease_id, sensor_id, created_at, updated_at
FROM disease_sensors
WHERE disease_id = $1 AND sensor_id = $2
LIMIT 1;

-- name: ListDiseaseSensorsByDisease :many
SELECT disease_id, sensor_id, created_at, updated_at
FROM disease_sensors
WHERE disease_id = $1
ORDER BY created_at DESC;

-- name: ListDiseaseSensorsBySensor :many
SELECT disease_id, sensor_id, created_at, updated_at
FROM disease_sensors
WHERE sensor_id = $1
ORDER BY created_at DESC;

-- name: ListAllDiseaseSensors :many
SELECT disease_id, sensor_id, created_at, updated_at
FROM disease_sensors
ORDER BY created_at DESC;

-- name: CountDiseaseSensorsByDisease :one
SELECT COUNT(*)
FROM disease_sensors
WHERE disease_id = $1;

-- name: CountDiseaseSensorsBySensor :one
SELECT COUNT(*)
FROM disease_sensors
WHERE sensor_id = $1;

-- SQL constants for mutations (used in repository)
-- name: CreateDiseaseSensor :exec
INSERT INTO disease_sensors (
    disease_id,
    sensor_id,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4);

-- name: UpdateDiseaseSensor :exec
UPDATE disease_sensors
SET
    updated_at = $3
WHERE disease_id = $1 AND sensor_id = $2;

-- name: DeleteDiseaseSensor :exec
DELETE FROM disease_sensors
WHERE disease_id = $1 AND sensor_id = $2;

-- name: DeleteDiseaseSensorsByDisease :exec
DELETE FROM disease_sensors
WHERE disease_id = $1;

-- name: DeleteDiseaseSensorsBySensor :exec
DELETE FROM disease_sensors
WHERE sensor_id = $1;
