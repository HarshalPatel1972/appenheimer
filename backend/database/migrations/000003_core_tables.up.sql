CREATE TABLE organizations (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    website_url VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE apps (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    organization_id UUID NOT NULL,
    visibility_status visibility_status NOT NULL DEFAULT 'draft',
    health_status health_status NOT NULL DEFAULT 'unknown',
    verification_status verification_status NOT NULL DEFAULT 'pending',
    ranking_weight INTEGER NOT NULL DEFAULT 0,
    icon_key VARCHAR(255) NOT NULL DEFAULT '',
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    version BIGINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE app_links (
    id UUID PRIMARY KEY,
    app_id UUID NOT NULL,
    link_type VARCHAR(50) NOT NULL,
    url VARCHAR(255) NOT NULL
);

CREATE TABLE media (
    id UUID PRIMARY KEY,
    app_id UUID NOT NULL,
    media_type VARCHAR(50) NOT NULL,
    media_key VARCHAR(255) NOT NULL,
    order_index INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE aliases (
    id UUID PRIMARY KEY,
    app_id UUID NOT NULL,
    alias VARCHAR(255) NOT NULL,
    priority INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE app_releases (
    id UUID PRIMARY KEY,
    app_id UUID NOT NULL,
    version VARCHAR(100) NOT NULL,
    release_date TIMESTAMPTZ NOT NULL,
    release_notes TEXT NOT NULL DEFAULT ''
);
