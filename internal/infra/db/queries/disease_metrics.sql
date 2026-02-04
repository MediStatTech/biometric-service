-- name: GetDiseaseMetric :one
SELECT disease_id, metric_id, name, enum_name, updated_at, created_at
FROM disease_metrics
WHERE disease_id = $1 AND metric_id = $2
LIMIT 1;

-- name: ListDiseaseMetrics :many
SELECT disease_id, metric_id, name, enum_name, updated_at, created_at
FROM disease_metrics
ORDER BY disease_id, name;

-- name: ListDiseaseMetricsByDiseaseID :many
SELECT disease_id, metric_id, name, enum_name, updated_at, created_at
FROM disease_metrics
WHERE disease_id = $1
ORDER BY name;

-- name: CountDiseaseMetrics :one
SELECT COUNT(*) FROM disease_metrics;

-- SQL constants for mutations (used in repository)
-- name: CreateDiseaseMetric :exec
INSERT INTO disease_metrics (
    disease_id,
    metric_id,
    name,
    enum_name,
    created_at
) VALUES ($1, $2, $3, $4, $5);

-- name: UpdateDiseaseMetric :exec
UPDATE disease_metrics
SET
    name = $3,
    enum_name = $4,
    updated_at = $5
WHERE disease_id = $1 AND metric_id = $2;

-- name: DeleteDiseaseMetric :exec
DELETE FROM disease_metrics
WHERE disease_id = $1 AND metric_id = $2;
