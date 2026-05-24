CREATE TABLE IF NOT EXISTS metric_types (
    metric_type_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sensor_id      UUID NOT NULL REFERENCES sensors(sensor_id) ON DELETE CASCADE,
    code           TEXT NOT NULL,
    name           TEXT NOT NULL,
    symbol         TEXT NOT NULL,
    min_value      DOUBLE PRECISION NOT NULL,
    max_value      DOUBLE PRECISION NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT metric_types_sensor_code_unq UNIQUE (sensor_id, code),
    CONSTRAINT metric_types_range_chk       CHECK (min_value < max_value)
);

CREATE INDEX IF NOT EXISTS idx_metric_types_sensor_id
    ON metric_types(sensor_id);
