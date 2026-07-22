ALTER TABLE apps
    ADD CONSTRAINT fk_apps_organization
    FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE RESTRICT;

ALTER TABLE app_links
    ADD CONSTRAINT fk_app_links_app
    FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE;

ALTER TABLE media
    ADD CONSTRAINT fk_media_app
    FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE;

ALTER TABLE aliases
    ADD CONSTRAINT fk_aliases_app
    FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE;

ALTER TABLE app_releases
    ADD CONSTRAINT fk_app_releases_app
    FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE;

ALTER TABLE app_categories
    ADD CONSTRAINT fk_app_cat_app FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_app_cat_cat FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE;

ALTER TABLE app_platforms
    ADD CONSTRAINT fk_app_plat_app FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_app_plat_plat FOREIGN KEY (platform_id) REFERENCES platforms(id) ON DELETE CASCADE;

ALTER TABLE app_pricing
    ADD CONSTRAINT fk_app_pri_app FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_app_pri_pri FOREIGN KEY (pricing_id) REFERENCES pricing(id) ON DELETE CASCADE;

ALTER TABLE app_tags
    ADD CONSTRAINT fk_app_tag_app FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_app_tag_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE;
