-- name: GetSensorPatientMetric :one
SELECT sensor_id, patient_id, metric_id, value, symbol, created_at
FROM sensor_patient_metrics
WHERE sensor_id = $1 AND patient_id = $2 AND metric_id = $3 AND created_at = $4
LIMIT 1;

-- name: ListSensorPatientMetrics :many
SELECT sensor_id, patient_id, metric_id, value, symbol, created_at
FROM sensor_patient_metrics
WHERE sensor_id = $1 AND patient_id = $2
ORDER BY created_at DESC
LIMIT 50;

-- name: ListSensorPatientMetricsByTimeRange :many
SELECT sensor_id, patient_id, metric_id, value, symbol, created_at
FROM sensor_patient_metrics
WHERE sensor_id = $1
    AND patient_id = $2
    AND created_at >= $3
    AND created_at <= $4
ORDER BY created_at DESC;

-- name: ListSensorPatientMetricsBySensor :many
SELECT sensor_id, patient_id, metric_id, value, symbol, created_at
FROM sensor_patient_metrics
WHERE sensor_id = $1
ORDER BY created_at DESC;

-- name: ListSensorPatientMetricsByPatient :many
SELECT sensor_id, patient_id, metric_id, value, symbol, created_at
FROM sensor_patient_metrics
WHERE patient_id = $1
ORDER BY created_at DESC;

-- name: ListSensorPatientMetricsByMetric :many
SELECT sensor_id, patient_id, metric_id, value, symbol, created_at
FROM sensor_patient_metrics
WHERE metric_id = $1
ORDER BY created_at DESC;

-- name: CountSensorPatientMetrics :one
SELECT COUNT(*)
FROM sensor_patient_metrics
WHERE sensor_id = $1 AND patient_id = $2;

-- name: GetLatestSensorPatientMetric :one
SELECT sensor_id, patient_id, metric_id, value, symbol, created_at
FROM sensor_patient_metrics
WHERE sensor_id = $1 AND patient_id = $2 AND metric_id = $3
ORDER BY created_at DESC
LIMIT 1;

-- SQL constants for mutations (used in repository)
-- name: CreateSensorPatientMetric :exec
INSERT INTO sensor_patient_metrics (
    sensor_id,
    patient_id,
    metric_id,
    value,
    symbol,
    created_at
) VALUES ($1, $2, $3, $4, $5, $6);

-- name: DeleteSensorPatientMetric :exec
DELETE FROM sensor_patient_metrics
WHERE sensor_id = $1 AND patient_id = $2 AND metric_id = $3 AND created_at = $4;

-- name: DeleteSensorPatientMetrics :exec
DELETE FROM sensor_patient_metrics
WHERE sensor_id = $1 AND patient_id = $2;

-- name: DeleteSensorPatientMetricsBySensor :exec
DELETE FROM sensor_patient_metrics
WHERE sensor_id = $1;

-- name: DeleteSensorPatientMetricsByPatient :exec
DELETE FROM sensor_patient_metrics
WHERE patient_id = $1;

-- name: DeleteSensorPatientMetricsByMetric :exec
DELETE FROM sensor_patient_metrics
WHERE metric_id = $1;
