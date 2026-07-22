# RFC 008 - Data Model

## 1. Context
As Appenheimer transitions from a rapid prototype to a production-grade platform, the in-memory dataset must be replaced by PostgreSQL. Before any database infrastructure (migrations, `sqlc`, `pgx`) is implemented, we must establish the absolute source of truth: the Data Model. This schema will dictate the Admin System, the public API, and the future Meilisearch indexing pipeline.

## 2. Goals
- Define a highly normalized, relational structure that supports complex filtering.
- Establish robust constraints (foreign keys, uniqueness).
- Future-proof the model with soft deletes, versioning, and extensibility.
- Optimize for high-read performance while accommodating internal admin CRUD operations.

## 3. Core Entities & Relationships

### `apps` (Primary Entity)
- `id`: `UUID` (UUIDv7). Time-ordered, native UUID standard, excellent Postgres support.
- `slug`: `VARCHAR(255) UNIQUE`. URL-safe identifier (e.g., "visual-studio-code").
- `name`: `VARCHAR(255)`.
- `description`: `TEXT`.
- `organization_id`: `UUID` (Foreign Key -> `organizations.id`).
- `visibility_status`: `ENUM ('draft', 'published', 'archived')`. Admin approval state.
- `health_status`: `ENUM ('operational', 'degraded', 'maintenance', 'outage', 'unknown')`. Public operational state.
- `verification_status`: `ENUM ('verified', 'community', 'pending')`. For user-submitted apps.
- `ranking_weight`: `INTEGER`. Future-proofed deterministic sort priority (search score, popularity, recommendations).
- `icon_key`: `VARCHAR(255)`. Relative object key (e.g., `icons/figma.png`), not absolute URL.
- `metadata`: `JSONB`. Must **never** contain searchable or relational data. Strictly for optional metadata.
- `version`: `BIGINT DEFAULT 1`. For optimistic locking during admin edits.
- `created_at`, `updated_at`: `TIMESTAMPTZ`.
- `deleted_at`: `TIMESTAMPTZ` (Soft Delete).

### `organizations`
- `id`: `UUID` (UUIDv7)
- `slug`: `VARCHAR(255) UNIQUE`
- `name`: `VARCHAR(255)` (e.g., "OpenAI", "Microsoft")
- `website_url`: `VARCHAR(255)`
- Timestamps & Soft Delete.

### `app_links` (One-to-Many)
- `id`: `UUID` (UUIDv7)
- `app_id`: `UUID` (Foreign Key -> `apps.id` ON DELETE CASCADE)
- `link_type`: `VARCHAR(50)` (e.g., 'Website', 'GitHub', 'Documentation', 'Discord')
- `url`: `VARCHAR(255)`

### `media` (One-to-Many)
- `id`: `UUID` (UUIDv7)
- `app_id`: `UUID` (Foreign Key -> `apps.id` ON DELETE CASCADE)
- `media_type`: `VARCHAR(50)` (e.g., 'image', 'video', 'gif')
- `media_key`: `VARCHAR(255)`. Relative object key.
- `order_index`: `INTEGER`. To guarantee deterministic display order in the UI.

### `aliases` (One-to-Many)
- `id`: `UUID` (UUIDv7)
- `app_id`: `UUID` (Foreign Key)
- `alias`: `VARCHAR(255)` (e.g., "VSCode")
- `priority`: `INTEGER`. Search ranking priority (e.g., 100 for VSCode, 20 for Code).

### `app_releases` (One-to-Many)
- `id`: `UUID` (UUIDv7)
- `app_id`: `UUID` (Foreign Key)
- `version`: `VARCHAR(100)`
- `release_date`: `TIMESTAMPTZ`
- `release_notes`: `TEXT`

### Taxonomies: `categories`, `platforms`, `pricing`, `tags`
All of these function identically as highly normalized, independent tables connected to `apps` via many-to-many junction tables.
- **Table Structure**: `id` (UUIDv7), `slug` (UNIQUE), `name` (VARCHAR).
- **Junction Tables**: 
  - `app_categories` (`app_id`, `category_id`)
  - `app_platforms` (`app_id`, `platform_id`)
  - `app_pricing` (`app_id`, `pricing_id`)
  - `app_tags` (`app_id`, `tag_id`)
- **Indexes**: Composite Primary Keys on the junction tables `(app_id, foreign_id)` ensure uniqueness and hyper-fast join performance.

## 4. Technical Constraints
- **IDs**: Use UUIDv7. They generate chronologically, minimizing B-tree fragmentation.
- **Images**: Store relative object keys. The backend constructs absolute URLs, enabling seamless CDN switching, local dev, and signed URLs.
- **Soft Deletes**: `deleted_at` ensures we never permanently lose historical app data.
- **Normalization**: Strict normalization wins. No `TEXT[]` arrays for taxonomies. This ensures referential integrity, trivial admin UI management, and cleaner indexing.
