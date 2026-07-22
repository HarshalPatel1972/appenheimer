-- name: ClaimEvents :many
UPDATE outbox_events 
SET status = 'processing', 
    locked_at = NOW(), 
    locked_by = $1 
WHERE id IN (
    SELECT id FROM outbox_events 
    WHERE status = 'pending' 
      AND next_attempt_at <= NOW() 
    ORDER BY created_at ASC 
    LIMIT $2 
    FOR UPDATE SKIP LOCKED
)
RETURNING *;

-- name: MarkEventCompleted :exec
UPDATE outbox_events 
SET status = 'completed', 
    updated_at = NOW() 
WHERE id = $1;

-- name: MarkEventFailed :exec
UPDATE outbox_events 
SET status = 'pending', 
    attempts = attempts + 1, 
    last_error = $2, 
    next_attempt_at = $3, 
    locked_at = NULL, 
    locked_by = NULL,
    updated_at = NOW()
WHERE id = $1;

-- name: MarkEventDead :exec
UPDATE outbox_events 
SET status = 'dead', 
    attempts = attempts + 1, 
    last_error = $2, 
    locked_at = NULL, 
    locked_by = NULL,
    updated_at = NOW()
WHERE id = $1;

-- name: ReclaimStaleLocks :exec
UPDATE outbox_events 
SET status = 'pending', 
    locked_at = NULL, 
    locked_by = NULL,
    updated_at = NOW()
WHERE status = 'processing' 
  AND locked_at < NOW() - INTERVAL '5 MINUTES'
  AND attempts < $1;

-- name: GetWorkerHealth :one
SELECT 
    COUNT(*) FILTER (WHERE status = 'pending') AS pending_events,
    COUNT(*) FILTER (WHERE status = 'dead') AS dead_events,
    COALESCE(MIN(created_at) FILTER (WHERE status = 'pending'), NOW()) AS oldest_pending,
    MAX(updated_at) FILTER (WHERE status = 'completed') AS last_successful
FROM outbox_events;
