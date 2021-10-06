CREATE EXTENSION postgres_fdw;
CREATE SERVER localhost FOREIGN DATA WRAPPER postgres_fdw OPTIONS (host 'localhost', dbname 'migrate', port '5432');
CREATE USER MAPPING FOR go_commerce SERVER localhost OPTIONS (user 'go_commerce', password 'go_commerce');
CREATE SCHEMA IF NOT EXISTS migrate;
IMPORT FOREIGN SCHEMA public FROM SERVER localhost INTO migrate;

INSERT INTO "public"."categories"
(
    id,
    "name",
    "description",
    "order"
)
SELECT
    id,
    "name",
    "description",
    sort_order
FROM "migrate"."categories";

INSERT INTO "public"."manufacturers"
(
    id,
    "name",
    website_url
)
SELECT
    id,
    "name",
    website_url
FROM "migrate"."manufacturers";

INSERT INTO "public"."products"
(
    id,
    created_at,
    updated_at,
    "name",
    description_short,
    description_long,
    price,
    manufacturer_id,
    "status",
    stock_count
)
SELECT
    id,
    created_at,
    updated_at,
    "name",
    COALESCE(description_short, ''),
    COALESCE(description_long, ''),
    price,
    manufacturer_id,
    CASE WHEN "status" IS NULL OR "status" = '' THEN 'AVAILABLE' ELSE "status" END,
    stock_count
FROM "migrate"."products";

INSERT INTO "public"."product_categories"
(
    product_id,
    category_id
)
SELECT
    product_id,
    category_id
FROM "migrate"."category_products";

-- sudo ls -1 /opt/appdata/bjoetiek/backend/images | grep -v fit | cut -d. -f1 > images.txt
-- sed 's/.*/uuidgen/e' images.txt | paste -d ' ' - images.txt | sed "s/ /', '/" | sed "s/^/INSERT INTO \"public\".\"images\" VALUES ('/" | sed "s/$/', 'products', '.png', 0);/" > migrate-images.sql
UPDATE images
SET owner_type = 'manufacturers'
WHERE id IN (
    SELECT i.id
    FROM images i
    JOIN manufacturers m ON i.owner_id = m.id
)

UPDATE images
SET owner_type = 'categories'
WHERE id IN (
    SELECT i.id
    FROM images i
    JOIN categories c ON i.owner_id = c.id
)