CREATE UNIQUE INDEX idx_organizations_slug ON organizations (slug);
CREATE UNIQUE INDEX idx_apps_slug ON apps (slug);
CREATE UNIQUE INDEX idx_categories_slug ON categories (slug);
CREATE UNIQUE INDEX idx_platforms_slug ON platforms (slug);
CREATE UNIQUE INDEX idx_pricing_slug ON pricing (slug);
CREATE UNIQUE INDEX idx_tags_slug ON tags (slug);

CREATE INDEX idx_apps_org_id ON apps (organization_id);
CREATE INDEX idx_app_links_app_id ON app_links (app_id);
CREATE INDEX idx_media_app_id ON media (app_id);
CREATE INDEX idx_aliases_app_id ON aliases (app_id);
CREATE INDEX idx_app_releases_app_id ON app_releases (app_id);
