-- name: GetSensor :one
SELECT sensor_id, name, code, created_at, updated_at
FROM sensors
WHERE sensor_id = $1
LIMIT 1;

-- name: GetSensorByCode :one
SELECT sensor_id, name, code, created_at, updated_at
FROM sensors
WHERE code = $1
LIMIT 1;

-- name: ListSensors :many
SELECT sensor_id, name, code, created_at, updated_at
FROM sensors
ORDER BY created_at DESC;

-- name: CountSensors :one
SELECT COUNT(*) FROM sensors;

-- SQL constants for mutations (used in repository)
-- name: CreateSensor :exec
INSERT INTO sensors (
    sensor_id,
    name,
    code,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5);

-- name: UpdateSensor :exec
UPDATE sensors
SET
    name = $2,
    code = $3,
    updated_at = $4
WHERE sensor_id = $1;

-- name: DeleteSensor :exec
DELETE FROM sensors
WHERE sensor_id = $1;
