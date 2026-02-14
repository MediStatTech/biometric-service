CREATE TABLE sensors (
    sensor_id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT NOT NULL,
    code        TEXT NOT NULL UNIQUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE sensor_patients (
    sensor_id   UUID NOT NULL,
    patient_id  UUID NOT NULL,
    status      TEXT NOT NULL DEFAULT 'active',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (sensor_id, patient_id),

    CONSTRAINT sensor_patients_sensor_fk
        FOREIGN KEY (sensor_id) REFERENCES sensors(sensor_id) ON DELETE CASCADE,

    CONSTRAINT sensor_patients_status_chk
        CHECK (status IN ('active', 'inactive'))
);

CREATE INDEX idx_sensor_patients_patient_id
    ON sensor_patients(patient_id);

CREATE INDEX idx_sensor_patients_status
    ON sensor_patients(status);


CREATE TABLE sensor_patient_metrics (
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

CREATE INDEX idx_sensor_patient_metrics_patient_id
    ON sensor_patient_metrics(patient_id);

CREATE INDEX idx_sensor_patient_metrics_sensor_id
    ON sensor_patient_metrics(sensor_id);

CREATE INDEX idx_sensor_patient_metrics_created_at
    ON sensor_patient_metrics(created_at);