CREATE TABLE imports (
    id UUID PRIMARY KEY,
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    finished_at TIMESTAMPTZ,
    source VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    report JSONB NOT NULL DEFAULT '{}'::jsonb
);

ALTER TABLE apps ADD COLUMN source VARCHAR(255) NOT NULL DEFAULT 'manual';
