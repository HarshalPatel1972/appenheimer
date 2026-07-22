ALTER TABLE outbox_events DROP COLUMN locked_by;
ALTER TABLE outbox_events DROP COLUMN locked_at;
ALTER TABLE outbox_events DROP COLUMN next_attempt_at;
ALTER TABLE outbox_events DROP COLUMN last_error;
ALTER TABLE outbox_events DROP COLUMN attempts;
ALTER TABLE outbox_events DROP COLUMN aggregate_id;
ALTER TABLE outbox_events DROP COLUMN event_version;
