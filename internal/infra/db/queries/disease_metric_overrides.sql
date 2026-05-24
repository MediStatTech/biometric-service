-- name: GetDiseaseMetricOverride :one
SELECT disease_id, metric_type_id, min_value, max_value, created_at, updated_at
FROM disease_metric_overrides
WHERE disease_id = $1 AND metric_type_id = $2
LIMIT 1;

-- name: ListDiseaseMetricOverridesByDisease :many
SELECT disease_id, metric_type_id, min_value, max_value, created_at, updated_at
FROM disease_metric_overrides
WHERE disease_id = $1;

-- name: ListDiseaseMetricOverridesByMetricType :many
SELECT disease_id, metric_type_id, min_value, max_value, created_at, updated_at
FROM disease_metric_overrides
WHERE metric_type_id = $1;

-- name: ListAllDiseaseMetricOverrides :many
SELECT disease_id, metric_type_id, min_value, max_value, created_at, updated_at
FROM disease_metric_overrides;

-- name: CreateDiseaseMetricOverride :exec
INSERT INTO disease_metric_overrides (
    disease_id,
    metric_type_id,
    min_value,
    max_value,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateDiseaseMetricOverride :exec
UPDATE disease_metric_overrides
SET
    min_value  = $3,
    max_value  = $4,
    updated_at = $5
WHERE disease_id = $1 AND metric_type_id = $2;

-- name: DeleteDiseaseMetricOverride :exec
DELETE FROM disease_metric_overrides
WHERE disease_id = $1 AND metric_type_id = $2;

-- name: DeleteDiseaseMetricOverridesByDisease :exec
DELETE FROM disease_metric_overrides
WHERE disease_id = $1;
