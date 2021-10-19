#!/bin/bash
# Based on https://stackoverflow.com/a/54946350

pg_restore -U ${POSTGRES_USER} --dbname=${POSTGRES_DB} --verbose --single-transaction < /init.pgdump || exit 1