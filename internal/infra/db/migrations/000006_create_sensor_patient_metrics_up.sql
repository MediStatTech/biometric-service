CREATE TABLE IF NOT EXISTS sensor_patient_metrics (
    sensor_id   UUID NOT NULL,
    patient_id  UUID NOT NULL,
    metric_id   UUID NOT NULL,
    value       DOUBLE PRECISION NOT NULL,
    symbol      VARCHAR(50) NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (sensor_id, patient_id, metric_id, created_at),

    CONSTRAINT sensor_patient_metrics_fk
        FOREIGN KEY (sensor_id, patient_id)
        REFERENCES sensor_patients(sensor_id, patient_id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_sensor_patient_metrics_patient_id
    ON sensor_patient_metrics(patient_id);

CREATE INDEX IF NOT EXISTS idx_sensor_patient_metrics_sensor_id
    ON sensor_patient_metrics(sensor_id);

CREATE INDEX IF NOT EXISTS idx_sensor_patient_metrics_created_at
    ON sensor_patient_metrics(created_at);
