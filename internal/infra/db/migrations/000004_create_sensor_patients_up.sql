CREATE TABLE IF NOT EXISTS sensor_patients (
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

CREATE INDEX IF NOT EXISTS idx_sensor_patients_patient_id
    ON sensor_patients(patient_id);

CREATE INDEX IF NOT EXISTS idx_sensor_patients_status
    ON sensor_patients(status);
