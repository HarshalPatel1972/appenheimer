CREATE TABLE categories (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE app_categories (
    app_id UUID NOT NULL,
    category_id UUID NOT NULL,
    PRIMARY KEY (app_id, category_id)
);

CREATE TABLE platforms (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE app_platforms (
    app_id UUID NOT NULL,
    platform_id UUID NOT NULL,
    PRIMARY KEY (app_id, platform_id)
);

CREATE TABLE pricing (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE app_pricing (
    app_id UUID NOT NULL,
    pricing_id UUID NOT NULL,
    PRIMARY KEY (app_id, pricing_id)
);

CREATE TABLE tags (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE app_tags (
    app_id UUID NOT NULL,
    tag_id UUID NOT NULL,
    PRIMARY KEY (app_id, tag_id)
);
