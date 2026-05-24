CREATE TABLE IF NOT EXISTS disease_metric_overrides (
    disease_id     UUID NOT NULL REFERENCES diseases(disease_id)         ON DELETE CASCADE,
    metric_type_id UUID NOT NULL REFERENCES metric_types(metric_type_id) ON DELETE CASCADE,
    min_value      DOUBLE PRECISION NOT NULL,
    max_value      DOUBLE PRECISION NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (disease_id, metric_type_id),
    CONSTRAINT disease_metric_overrides_range_chk CHECK (min_value < max_value)
);

CREATE INDEX IF NOT EXISTS idx_disease_metric_overrides_metric_type
    ON disease_metric_overrides(metric_type_id);
