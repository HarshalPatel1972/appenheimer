-- name: CreatePlatform :one
INSERT INTO platforms (id, slug, name) VALUES ($1, $2, $3) RETURNING *;

-- name: LinkAppPlatform :exec
INSERT INTO app_platforms (app_id, platform_id) VALUES ($1, $2);

-- name: GetPlatformBySlug :one
SELECT * FROM platforms WHERE slug = $1 LIMIT 1;
-- name: GetCategoryBySlug :one
SELECT * FROM categories WHERE slug = $1 LIMIT 1;
-- name: GetPricingBySlug :one
SELECT * FROM pricing WHERE slug = $1 LIMIT 1;
-- name: GetTagBySlug :one
SELECT * FROM tags WHERE slug = $1 LIMIT 1;
