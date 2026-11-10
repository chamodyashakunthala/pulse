CREATE TABLE services (
  id         SERIAL PRIMARY KEY,
  name       TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE deploys (
  id          SERIAL PRIMARY KEY,
  service_id  INT NOT NULL REFERENCES services(id),
  version     TEXT NOT NULL,
  deployed_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX idx_deploys_service_time ON deploys (service_id, deployed_at DESC);

CREATE TABLE metrics_minute (
  service_id     INT NOT NULL REFERENCES services(id),
  bucket_start   TIMESTAMPTZ NOT NULL,
  request_count  INT NOT NULL DEFAULT 0,
  error_count    INT NOT NULL DEFAULT 0,
  latency_sum_ms BIGINT NOT NULL DEFAULT 0,
  latency_max_ms INT NOT NULL DEFAULT 0,
  latency_hist   JSONB NOT NULL DEFAULT '{}',
  PRIMARY KEY (service_id, bucket_start)
);

CREATE TABLE incidents (
  id             SERIAL PRIMARY KEY,
  service_id     INT NOT NULL REFERENCES services(id),
  kind           TEXT NOT NULL CHECK (kind IN ('latency','error_rate')),
  status         TEXT NOT NULL DEFAULT 'open' CHECK (status IN ('open','resolved')),
  baseline_value DOUBLE PRECISION NOT NULL,
  observed_value DOUBLE PRECISION NOT NULL,
  opened_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
  resolved_at    TIMESTAMPTZ
);
CREATE INDEX idx_incidents_service_status ON incidents (service_id, status);

CREATE TABLE incident_reports (
  id             SERIAL PRIMARY KEY,
  incident_id    INT NOT NULL REFERENCES incidents(id),
  status         TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending','done','failed')),
  probable_cause TEXT,
  evidence       JSONB,
  model          TEXT,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);