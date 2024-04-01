#!/bin/sh

set -e

echo "start the app"
source /app.env
source /app/app.env

exec "$@"
