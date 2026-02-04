-- name: GetSensorMetric :one
SELECT sensor_id, metric_id, value, created_at
FROM sensor_metrics
WHERE sensor_id = $1 AND metric_id = $2
LIMIT 1;

-- name: ListSensorMetrics :many
SELECT sensor_id, metric_id, value, created_at
FROM sensor_metrics
ORDER BY created_at DESC;

-- name: ListSensorMetricsBySensorID :many
SELECT sensor_id, metric_id, value, created_at
FROM sensor_metrics
WHERE sensor_id = $1
ORDER BY created_at DESC;

-- name: CountSensorMetrics :one
SELECT COUNT(*) FROM sensor_metrics;

-- SQL constants for mutations (used in repository)
-- name: CreateSensorMetric :exec
INSERT INTO sensor_metrics (
    sensor_id,
    metric_id,
    value,
    created_at
) VALUES ($1, $2, $3, $4);

-- name: UpdateSensorMetric :exec
UPDATE sensor_metrics
SET
    value = $3
WHERE sensor_id = $1 AND metric_id = $2;

-- name: DeleteSensorMetric :exec
DELETE FROM sensor_metrics
WHERE sensor_id = $1 AND metric_id = $2;
