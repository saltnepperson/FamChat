#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 migration_name"
  exit 1
fi

MIGRATION_DIRECTORY="db/migrations"

mkdir -p "$MIGRATION_DIRECTORY"

TIMESTAMP=$(date +"%Y%m%d%H%M%S")

UP_FILE="${MIGRATION_DIRECTORY}/${TIMESTAMP}_$1.up.sql"
DOWN_FILE="${MIGRATION_DIRECTORY}/${TIMESTAMP}_$1.down.sql"

touch "$UP_FILE"
touch "$DOWN_FILE"

echo "Created migration files:"
echo "  $UP_FILE"
echo "  $DOWN_FILE"
