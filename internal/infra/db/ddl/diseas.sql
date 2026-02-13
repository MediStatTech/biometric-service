-- =========================
-- Diseases
-- =========================
CREATE TABLE diseases (
    disease_id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT NOT NULL,
    code        TEXT NOT NULL UNIQUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- =========================
-- Disease <-> Sensors mapping
-- =========================
CREATE TABLE disease_sensors (
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