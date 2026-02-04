CREATE TABLE IF NOT EXISTS disease_metrics (
  disease_id   UUID NOT NULL,
  metric_id    UUID NOT NULL,
  name         TEXT NOT NULL,
  enum_name    TEXT NOT NULL,
  created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at   TIMESTAMPTZ,

  CONSTRAINT pk_disease_metrics PRIMARY KEY (disease_id, metric_id),
  CONSTRAINT fk_disease_metrics_disease
    FOREIGN KEY (disease_id)
    REFERENCES diseases(disease_id)
    ON DELETE CASCADE
);
