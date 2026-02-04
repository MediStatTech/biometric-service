-- name: GetSensor :one
SELECT sensor_id, name, status, enum_name, updated_at, created_at
FROM sensors
WHERE sensor_id = $1
LIMIT 1;

-- name: ListSensors :many
SELECT sensor_id, name, status, enum_name, updated_at, created_at
FROM sensors
ORDER BY created_at DESC;

-- name: ListSensorsByStatus :many
SELECT sensor_id, name, status, enum_name, updated_at, created_at
FROM sensors
WHERE status = $1
ORDER BY created_at DESC;

-- name: CountSensors :one
SELECT COUNT(*) FROM sensors;

-- SQL constants for mutations (used in repository)
-- name: CreateSensor :exec
INSERT INTO sensors (
    sensor_id,
    name,
    status,
    enum_name,
    created_at
) VALUES ($1, $2, $3, $4, $5);

-- name: UpdateSensor :exec
UPDATE sensors
SET
    name = $2,
    status = $3,
    enum_name = $4,
    updated_at = $5
WHERE sensor_id = $1;

-- name: DeleteSensor :exec
DELETE FROM sensors
WHERE sensor_id = $1;
