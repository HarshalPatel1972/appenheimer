-- name: CreateApp :one
INSERT INTO apps (
    id, slug, name, description, organization_id, 
    visibility_status, health_status, verification_status,
    ranking_weight, icon_key, metadata, version, source
) VALUES (
    $1, $2, $3, $4, $5, 
    $6, $7, $8, 
    $9, $10, $11, $12, $13
) RETURNING *;

-- name: GetApp :one
SELECT * FROM apps WHERE id = $1 LIMIT 1;

-- name: UpdateApp :one
UPDATE apps SET
    name = $2,
    description = $3,
    organization_id = $4,
    visibility_status = $5,
    health_status = $6,
    verification_status = $7,
    ranking_weight = $8,
    icon_key = $9,
    metadata = $10,
    version = version + 1,
    source = $11,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetAppBySlug :one
SELECT * FROM apps WHERE slug = $1 LIMIT 1;
