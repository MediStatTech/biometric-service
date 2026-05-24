ALTER TABLE sensor_patient_metrics
    DROP CONSTRAINT IF EXISTS sensor_patient_metrics_metric_fk;

ALTER TABLE sensor_patient_metrics
    ADD CONSTRAINT sensor_patient_metrics_metric_fk
    FOREIGN KEY (metric_id)
    REFERENCES metric_types(metric_type_id)
    ON DELETE CASCADE;
