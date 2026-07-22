-- name: GetAdminDashboardStats :one
SELECT 
    COUNT(*) FILTER (WHERE visibility_status = 'review') AS pending_reviews,
    COUNT(*) FILTER (WHERE visibility_status = 'needs_changes') AS needs_changes,
    COUNT(*) FILTER (WHERE visibility_status = 'published' AND created_at >= NOW() - INTERVAL '24 HOURS') AS published_today
FROM apps;

-- name: GetReviewQueue :many
SELECT * FROM apps
WHERE visibility_status IN ('draft', 'review', 'needs_changes')
ORDER BY 
    CASE visibility_status
        WHEN 'review' THEN 1
        WHEN 'draft' THEN 2
        WHEN 'needs_changes' THEN 3
        ELSE 4
    END ASC,
    created_at ASC;

-- name: UpdateAppStatus :one
UPDATE apps SET 
    visibility_status = $2,
    updated_at = NOW(),
    updated_by = $3
WHERE id = $1
RETURNING *;

-- name: CreateOutboxEvent :exec
INSERT INTO outbox_events (id, event_type, payload, status)
VALUES ($1, $2, $3, 'pending');

-- name: ArchiveShadowDrafts :exec
UPDATE apps SET visibility_status = 'archived'
WHERE source_app_id = $1 AND visibility_status != 'archived';
