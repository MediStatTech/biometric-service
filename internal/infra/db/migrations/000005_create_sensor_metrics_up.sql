CREATE TABLE IF NOT EXISTS sensor_metrics (
  sensor_id  UUID NOT NULL,
  metric_id  UUID NOT NULL,
  value      DOUBLE PRECISION NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT pk_sensor_metrics PRIMARY KEY (sensor_id, metric_id),
  CONSTRAINT fk_sensor_metrics_sensor
    FOREIGN KEY (sensor_id)
    REFERENCES sensors(sensor_id)
    ON DELETE CASCADE
);
