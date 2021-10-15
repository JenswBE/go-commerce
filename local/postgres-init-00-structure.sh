#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE TABLE "public"."categories" (
    "id" uuid NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "name" text,
    "description" text,
    "parent_id" uuid,
    "order" bigint,
    CONSTRAINT "categories_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE TABLE "public"."images" (
    "id" uuid NOT NULL,
    "owner_id" uuid,
    "owner_type" text,
    "extension" text,
    "order" bigint,
    CONSTRAINT "images_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE TABLE "public"."manufacturers" (
    "id" uuid NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "name" text,
    "website_url" text,
    CONSTRAINT "manufacturers_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE TABLE "public"."migrations" (
    "id" character varying(255) NOT NULL,
    CONSTRAINT "migrations_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE TABLE "public"."product_categories" (
    "product_id" uuid NOT NULL,
    "category_id" uuid NOT NULL,
    CONSTRAINT "product_categories_pkey" PRIMARY KEY ("product_id", "category_id")
) WITH (oids = false);


CREATE TABLE "public"."products" (
    "id" uuid NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "name" text,
    "description_short" text,
    "description_long" text,
    "price" bigint,
    "manufacturer_id" uuid,
    "status" text,
    "stock_count" bigint,
    CONSTRAINT "products_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


ALTER TABLE ONLY "public"."categories" ADD CONSTRAINT "fk_categories_children" FOREIGN KEY (parent_id) REFERENCES categories(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."product_categories" ADD CONSTRAINT "fk_product_categories_category" FOREIGN KEY (category_id) REFERENCES categories(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."product_categories" ADD CONSTRAINT "fk_product_categories_product" FOREIGN KEY (product_id) REFERENCES products(id) NOT DEFERRABLE;
EOSQL