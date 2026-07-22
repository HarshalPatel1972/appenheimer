
-- name: DeleteAppAliases :exec
DELETE FROM aliases WHERE app_id = $1;

-- name: CreateAppAlias :exec
INSERT INTO aliases (id, app_id, name) VALUES ($1, $2, $3);

-- name: DeleteAppLinks :exec
DELETE FROM app_links WHERE app_id = $1;

-- name: CreateAppLink :exec
INSERT INTO app_links (id, app_id, link_type, url) VALUES ($1, $2, $3, $4);

-- name: DeleteAppMedia :exec
DELETE FROM media WHERE app_id = $1;

-- name: CreateAppMedia :exec
INSERT INTO media (id, app_id, object_key, media_type, "order") VALUES ($1, $2, $3, $4, $5);

-- name: DeleteAppCategories :exec
DELETE FROM app_categories WHERE app_id = $1;

-- name: CreateAppCategory :exec
INSERT INTO app_categories (app_id, category_id) VALUES ($1, $2);

-- name: DeleteAppPlatforms :exec
DELETE FROM app_platforms WHERE app_id = $1;

-- name: CreateAppPlatform :exec
INSERT INTO app_platforms (app_id, platform_id) VALUES ($1, $2);

-- name: DeleteAppPricing :exec
DELETE FROM app_pricing WHERE app_id = $1;

-- name: CreateAppPricing :exec
INSERT INTO app_pricing (app_id, pricing_id) VALUES ($1, $2);
