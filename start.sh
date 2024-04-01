#!/bin/sh

set -e

echo "start the app"
source /app/app.env

exec "$@"
