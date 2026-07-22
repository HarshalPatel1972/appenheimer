
-- name: PatchApp :one
UPDATE apps SET
    name = $2,
    description = $3,
    organization_id = $4,
    health_status = $5,
    ranking_weight = $6,
    icon_key = $7,
    version = version + 1,
    updated_at = NOW(),
    updated_by = $8
WHERE id = $1 AND version = $9
RETURNING *;

-- name: CloneAppAsShadowDraft :one
INSERT INTO apps (
    id, name, slug, description, organization_id, visibility_status, health_status, 
    verification_status, ranking_weight, icon_key, metadata, source_app_id, version, updated_by
)
SELECT 
    $1, name, slug, description, organization_id, 'draft', health_status,
    verification_status, ranking_weight, icon_key, metadata, $2, 1, $3
FROM apps WHERE id = $2
RETURNING *;
