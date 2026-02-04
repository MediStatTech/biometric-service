CREATE TABLE IF NOT EXISTS sensors (
  sensor_id    UUID PRIMARY KEY,
  name         TEXT NOT NULL,
  status       TEXT NOT NULL,
  enum_name    TEXT NOT NULL,

  created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at   TIMESTAMPTZ
);
