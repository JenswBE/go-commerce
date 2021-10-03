#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    DELETE FROM "product_categories"
    WHERE "category_id" IN ('1e8d3b8d-fde6-4a66-ace2-14451f46607e');

    DELETE FROM "images"
    WHERE "id" IN (
        '2fc4513d-810b-42d3-8869-001a5b039aa6',
        'f6aca5c2-8aa8-4f14-ad71-26c697859716'
    );

    DELETE FROM "products"
    WHERE "id" IN (
        '0de097d8-8a4f-48d0-8748-758949d27f34',
        '14084ef2-be6b-4fe9-b521-7a3f358fc670',
        '6cc35c44-1b44-4765-ab2f-361a3a195d32'
    );
    
    DELETE FROM "manufacturers"
    WHERE "id" IN ('eb2b656a-0451-4b1a-8ea7-35f213160b29');
    
    DELETE FROM "categories"
    WHERE "id" IN ('1e8d3b8d-fde6-4a66-ace2-14451f46607e');

    DELETE FROM "migrations"
    WHERE "id" IN ('202107302030');
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    INSERT INTO "migrations"
    SELECT * FROM json_populate_recordset (NULL::"migrations",
        '[
            { "id": "202107302030" }
        ]'
    );
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    INSERT INTO "categories"
    SELECT * FROM json_populate_recordset (NULL::"categories",
        '[
            {
                "id": "1e8d3b8d-fde6-4a66-ace2-14451f46607e",
                "created_at": "2021-10-03 12:00:00+00",
                "updated_at": "2021-10-04 14:15:00+00",
                "name": "Promotions",
                "description": "Time to save some money",
                "parent_id": null,
                "order": 0
            }
        ]'
    );
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    INSERT INTO "manufacturers"
    SELECT * FROM json_populate_recordset (NULL::"manufacturers",
        '[
            {
                "id": "eb2b656a-0451-4b1a-8ea7-35f213160b29",
                "created_at": "2021-10-03 12:00:00+00",
                "updated_at": "2021-10-04 14:15:00+00",
                "name": "JenswBE",
                "website_url": "https://jensw.be"
            }
        ]'
    );
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    INSERT INTO "products"
    SELECT * FROM json_populate_recordset (NULL::"products",
        '[
            {
                "id": "0de097d8-8a4f-48d0-8748-758949d27f34",
                "created_at": "2021-10-03 12:00:00+00",
                "updated_at": "2021-10-04 14:15:00+00",
                "name": "Product XYZ",
                "description_short": "The only product you will ever need.",
                "description_long": "This product is specifically designed to be the best in XYZ.",
                "price": 1999,
                "manufacturer_id": "eb2b656a-0451-4b1a-8ea7-35f213160b29",
                "status": "AVAILABLE",
                "stock_count": 5
            },
            {
                "id": "14084ef2-be6b-4fe9-b521-7a3f358fc670",
                "created_at": "2021-10-03 12:00:00+00",
                "updated_at": "2021-10-04 14:15:00+00",
                "name": "Minimal product",
                "description_short": "",
                "description_long": "",
                "price": 2500,
                "manufacturer_id": null,
                "status": "AVAILABLE",
                "stock_count": 8
            },
            {
                "id": "6cc35c44-1b44-4765-ab2f-361a3a195d32",
                "created_at": "2021-10-03 12:00:00+00",
                "updated_at": "2021-10-04 14:15:00+00",
                "name": "Archived product",
                "description_short": "",
                "description_long": "",
                "price": 599,
                "manufacturer_id": "eb2b656a-0451-4b1a-8ea7-35f213160b29",
                "status": "ARCHIVED",
                "stock_count": 0
            }
        ]'
    );
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    INSERT INTO "images"
    SELECT * FROM json_populate_recordset (NULL::"images",
        '[
            {
                "id": "2fc4513d-810b-42d3-8869-001a5b039aa6",
                "owner_id": "0de097d8-8a4f-48d0-8748-758949d27f34",
                "owner_type": "products",
                "extension": ".jpg",
                "order": 0
            },
            {
                "id": "f6aca5c2-8aa8-4f14-ad71-26c697859716",
                "owner_id": "0de097d8-8a4f-48d0-8748-758949d27f34",
                "owner_type": "products",
                "extension": ".jpg",
                "order": 1
            }
        ]'
    );
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    INSERT INTO "product_categories"
    SELECT * FROM json_populate_recordset (NULL::"product_categories",
        '[
            {
                "product_id": "0de097d8-8a4f-48d0-8748-758949d27f34",
                "category_id": "1e8d3b8d-fde6-4a66-ace2-14451f46607e"
            }
        ]'
    );
EOSQL