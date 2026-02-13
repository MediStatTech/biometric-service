CREATE TABLE IF NOT EXISTS disease_sensors (
    disease_id  UUID NOT NULL,
    sensor_id   UUID NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (disease_id, sensor_id),

    CONSTRAINT disease_sensors_disease_fk
        FOREIGN KEY (disease_id)
        REFERENCES diseases(disease_id)
        ON DELETE CASCADE,

    CONSTRAINT disease_sensors_sensor_fk
        FOREIGN KEY (sensor_id)
        REFERENCES sensors(sensor_id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_disease_sensors_sensor_id
    ON disease_sensors(sensor_id);

CREATE INDEX IF NOT EXISTS idx_disease_sensors_disease_id
    ON disease_sensors(disease_id);
