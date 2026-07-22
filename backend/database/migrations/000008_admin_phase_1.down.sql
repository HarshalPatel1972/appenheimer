DROP TABLE outbox_events;
ALTER TABLE apps DROP COLUMN updated_by;
ALTER TABLE apps DROP COLUMN source_app_id;
-- Note: Postgres does not support DROP VALUE from ENUM easily.
