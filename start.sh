#!/bin/sh

# Run migrations
migrate -path /app/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

# Start the application
./main