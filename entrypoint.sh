#!/bin/bash
set -e

echo "DB URL: $DB_DSN"

# Run database migrations
echo "Running migrations..."
migrate -path /migrations -database "$DB_DSN" up

# Start the Go application
echo "Starting the application..."
exec /app/myapp
