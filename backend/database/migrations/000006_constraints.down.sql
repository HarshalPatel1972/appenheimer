ALTER TABLE app_tags DROP CONSTRAINT fk_app_tag_tag, DROP CONSTRAINT fk_app_tag_app;
ALTER TABLE app_pricing DROP CONSTRAINT fk_app_pri_pri, DROP CONSTRAINT fk_app_pri_app;
ALTER TABLE app_platforms DROP CONSTRAINT fk_app_plat_plat, DROP CONSTRAINT fk_app_plat_app;
ALTER TABLE app_categories DROP CONSTRAINT fk_app_cat_cat, DROP CONSTRAINT fk_app_cat_app;
ALTER TABLE app_releases DROP CONSTRAINT fk_app_releases_app;
ALTER TABLE aliases DROP CONSTRAINT fk_aliases_app;
ALTER TABLE media DROP CONSTRAINT fk_media_app;
ALTER TABLE app_links DROP CONSTRAINT fk_app_links_app;
ALTER TABLE apps DROP CONSTRAINT fk_apps_organization;
