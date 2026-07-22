ALTER TYPE visibility_status ADD VALUE IF NOT EXISTS 'review';
ALTER TYPE visibility_status ADD VALUE IF NOT EXISTS 'needs_changes';

ALTER TABLE apps ADD COLUMN source_app_id UUID REFERENCES apps(id) ON DELETE SET NULL;
ALTER TABLE apps ADD COLUMN updated_by UUID;

CREATE TABLE outbox_events (
    id UUID PRIMARY KEY,
    event_type VARCHAR(50) NOT NULL,
    payload JSONB NOT NULL DEFAULT '{}'::jsonb,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
