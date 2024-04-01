#!/bin/sh

set -e

echo "readable env"
source /app/app.env

echo "start the app"
exec "$@"
