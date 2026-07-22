ALTER TABLE outbox_events ADD COLUMN event_version INT NOT NULL DEFAULT 1;
ALTER TABLE outbox_events ADD COLUMN aggregate_id UUID;
ALTER TABLE outbox_events ADD COLUMN attempts INT NOT NULL DEFAULT 0;
ALTER TABLE outbox_events ADD COLUMN last_error TEXT;
ALTER TABLE outbox_events ADD COLUMN next_attempt_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
ALTER TABLE outbox_events ADD COLUMN locked_at TIMESTAMPTZ;
ALTER TABLE outbox_events ADD COLUMN locked_by VARCHAR(255);
