#!/bin/sh

set -e

echo "start the app"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
exec "$@"
