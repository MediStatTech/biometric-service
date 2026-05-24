-- name: GetMetricType :one
SELECT metric_type_id, sensor_id, code, name, symbol, min_value, max_value, created_at, updated_at
FROM metric_types
WHERE metric_type_id = $1
LIMIT 1;

-- name: GetMetricTypeBySensorAndCode :one
SELECT metric_type_id, sensor_id, code, name, symbol, min_value, max_value, created_at, updated_at
FROM metric_types
WHERE sensor_id = $1 AND code = $2
LIMIT 1;

-- name: ListMetricTypes :many
SELECT metric_type_id, sensor_id, code, name, symbol, min_value, max_value, created_at, updated_at
FROM metric_types
ORDER BY sensor_id, code;

-- name: ListMetricTypesBySensor :many
SELECT metric_type_id, sensor_id, code, name, symbol, min_value, max_value, created_at, updated_at
FROM metric_types
WHERE sensor_id = $1
ORDER BY code;

-- name: CountMetricTypesBySensor :one
SELECT COUNT(*) FROM metric_types
WHERE sensor_id = $1;

-- name: CreateMetricType :exec
INSERT INTO metric_types (
    metric_type_id,
    sensor_id,
    code,
    name,
    symbol,
    min_value,
    max_value,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: UpdateMetricType :exec
UPDATE metric_types
SET
    code       = $2,
    name       = $3,
    symbol     = $4,
    min_value  = $5,
    max_value  = $6,
    updated_at = $7
WHERE metric_type_id = $1;

-- name: DeleteMetricType :exec
DELETE FROM metric_types
WHERE metric_type_id = $1;

-- name: DeleteMetricTypesBySensor :exec
DELETE FROM metric_types
WHERE sensor_id = $1;
