-- name: GetSensorPatient :one
SELECT sensor_id, patient_id, status, created_at, updated_at
FROM sensor_patients
WHERE sensor_id = $1 AND patient_id = $2
LIMIT 1;

-- name: ListSensorPatientsBySensor :many
SELECT sensor_id, patient_id, status, created_at, updated_at
FROM sensor_patients
WHERE sensor_id = $1
ORDER BY created_at DESC;

-- name: ListSensorPatientsByPatient :many
SELECT sensor_id, patient_id, status, created_at, updated_at
FROM sensor_patients
WHERE patient_id = $1
ORDER BY created_at DESC;

-- name: ListSensorPatientsByStatus :many
SELECT sensor_id, patient_id, status, created_at, updated_at
FROM sensor_patients
WHERE status = $1
ORDER BY created_at DESC;

-- name: ListAllSensorPatients :many
SELECT sensor_id, patient_id, status, created_at, updated_at
FROM sensor_patients
ORDER BY created_at DESC;

-- name: CountSensorPatientsBySensor :one
SELECT COUNT(*)
FROM sensor_patients
WHERE sensor_id = $1;

-- name: CountSensorPatientsByPatient :one
SELECT COUNT(*)
FROM sensor_patients
WHERE patient_id = $1;

-- SQL constants for mutations (used in repository)
-- name: CreateSensorPatient :exec
INSERT INTO sensor_patients (
    sensor_id,
    patient_id,
    status,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5);

-- name: UpdateSensorPatient :exec
UPDATE sensor_patients
SET
    status = $3,
    updated_at = $4
WHERE sensor_id = $1 AND patient_id = $2;

-- name: DeleteSensorPatient :exec
DELETE FROM sensor_patients
WHERE sensor_id = $1 AND patient_id = $2;

-- name: DeleteSensorPatientsBySensor :exec
DELETE FROM sensor_patients
WHERE sensor_id = $1;

-- name: DeleteSensorPatientsByPatient :exec
DELETE FROM sensor_patients
WHERE patient_id = $1;
