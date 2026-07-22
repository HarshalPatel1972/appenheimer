import os

base_backend = 'c:/Users/Harshal Patel/Desktop/appenheimer/backend'
q_dir = os.path.join(base_backend, 'database', 'queries')

# Update apps.sql
apps_sql_path = os.path.join(q_dir, 'apps.sql')
with open(apps_sql_path, 'r', encoding='utf-8') as f:
    apps_sql = f.read()

apps_sql = apps_sql.replace(
    'icon_key, metadata, version',
    'icon_key, metadata, version, source'
).replace(
    '$9, $10, $11, $12',
    '$9, $10, $11, $12, $13'
)

# Append UpdateApp (idempotent upsert logic will be done in Go by checking first, or we can use ON CONFLICT but we use UUIDs and slugs)
# Wait, for idempotent insert, we can use ON CONFLICT (slug) DO UPDATE. But let's just do it manually in Go to generate proper diffs/reports.
# We just need an UpdateApp query.
apps_sql += """
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
"""
with open(apps_sql_path, 'w', encoding='utf-8') as f:
    f.write(apps_sql)

# Create imports.sql
imports_sql = """-- name: CreateImport :one
INSERT INTO imports (
    id, source, status, started_at
) VALUES (
    $1, $2, $3, NOW()
) RETURNING *;

-- name: UpdateImport :one
UPDATE imports SET
    status = $2,
    finished_at = NOW(),
    report = $3
WHERE id = $1
RETURNING *;
"""
with open(os.path.join(q_dir, 'imports.sql'), 'w', encoding='utf-8') as f:
    f.write(imports_sql)

# Also need GetOrganizationBySlug, GetPlatformBySlug, etc. for the 'Resolve References' phase
with open(os.path.join(q_dir, 'organizations.sql'), 'a', encoding='utf-8') as f:
    f.write("\n-- name: GetOrganizationBySlug :one\nSELECT * FROM organizations WHERE slug = $1 LIMIT 1;\n")

with open(os.path.join(q_dir, 'taxonomies.sql'), 'a', encoding='utf-8') as f:
    f.write("""
-- name: GetPlatformBySlug :one
SELECT * FROM platforms WHERE slug = $1 LIMIT 1;
-- name: GetCategoryBySlug :one
SELECT * FROM categories WHERE slug = $1 LIMIT 1;
-- name: GetPricingBySlug :one
SELECT * FROM pricing WHERE slug = $1 LIMIT 1;
-- name: GetTagBySlug :one
SELECT * FROM tags WHERE slug = $1 LIMIT 1;
""")

# Create ingestion directories
ing_dir = os.path.join(base_backend, 'internal', 'ingestion')
os.makedirs(ing_dir, exist_ok=True)

with open(os.path.join(ing_dir, 'types.go'), 'w', encoding='utf-8') as f:
    f.write("""package ingestion

type CanonicalApp struct {
	Name         string   `json:"name" yaml:"name"`
	Slug         string   `json:"slug" yaml:"slug"`
	Organization string   `json:"organization" yaml:"organization"`
	Website      string   `json:"website" yaml:"website"`
	Categories   []string `json:"categories" yaml:"categories"`
	Platforms    []string `json:"platforms" yaml:"platforms"`
	Pricing      []string `json:"pricing" yaml:"pricing"`
	Aliases      []string `json:"aliases" yaml:"aliases"`
	Links        map[string]string `json:"links" yaml:"links"`
	Media        []string `json:"media" yaml:"media"`
}
""")

print("Scaffolded DB additions and ingestion types.")
