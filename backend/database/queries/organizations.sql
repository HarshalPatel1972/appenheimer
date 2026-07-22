-- name: CreateOrganization :one
INSERT INTO organizations (
    id, slug, name, website_url
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetOrganization :one
SELECT * FROM organizations WHERE id = $1 LIMIT 1;

-- name: GetOrganizationBySlug :one
SELECT * FROM organizations WHERE slug = $1 LIMIT 1;
