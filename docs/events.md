# Event API Documentation

## Overview

This document describes how to test the event-related API endpoints and set up database migrations.

## Setting Up golang-migrate

### Install golang-migrate CLI

```bash
# Linux (using snap)
sudo snap install golang-migrate

# Or download binary
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz
sudo mv migrate.linux-amd64 /usr/local/bin/migrate

# macOS
brew install golang-migrate

# Windows (using scoop)
scoop install migrate
```

### Run Migrations

```bash
# From the backend directory
cd backend

# Run all migrations
migrate -path pkg/db/migrations/sqlite -database "file:data/app.db" up

# Run specific number of migrations
migrate -path pkg/db/migrations/sqlite -database "file:data/app.db" up 1

# Rollback last migration
migrate -path pkg/db/migrations/sqlite -database "file:data/app.db" down

# Check current version
migrate -path pkg/db/migrations/sqlite -database "file:data/app.db" version

# Force a specific version (if needed)
migrate -path pkg/db/migrations/sqlite -database "file:data/app.db" force VERSION_NUMBER
```

### Create New Migration

```bash
# Create a new migration
migrate create -ext sql -dir pkg/db/migrations/sqlite -seq create_events_table

# This creates two files:
# - pkg/db/migrations/sqlite/000009_create_events_table.up.sql
# - pkg/db/migrations/sqlite/000009_create_events_table.down.sql
```

## Running the Server

```bash
# Create data directory
mkdir -p data

# Run the server
go run ./cmd/server

# Server will start on http://localhost:8080
```

## API Endpoints

### Create an Event

```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "group_id": "12345678-1234-1234-1234-123456789012",
    "title": "Group Meeting",
    "description": "Monthly group meeting",
    "date_time": "2024-01-15T18:00:00Z"
  }'
```

**Response:** `201 Created`
```json
{
  "id": "generated-uuid",
  "group_id": "12345678-1234-1234-1234-123456789012",
  "title": "Group Meeting",
  "description": "Monthly group meeting",
  "date_time": "2024-01-15T18:00:00Z",
  "created_by": "user-uuid-from-session",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Get an Event

```bash
curl -X GET http://localhost:8080/events/12345678-1234-1234-1234-123456789012
```

**Response:** `200 OK`
```json
{
  "id": "12345678-1234-1234-1234-123456789012",
  "group_id": "group-uuid",
  "title": "Event Title",
  "description": "Event Description",
  "date_time": "2024-01-15T18:00:00Z",
  "created_by": "user-uuid",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### List Group Events

```bash
curl -X GET http://localhost:8080/groups/12345678-1234-1234-1234-123456789012/events
```

**Response:** `200 OK`
```json
[
  {
    "id": "event-uuid-1",
    "group_id": "12345678-1234-1234-1234-123456789012",
    "title": "Event 1",
    "description": "Description 1",
    "date_time": "2024-01-15T18:00:00Z",
    "created_by": "user-uuid",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### Update an Event

```bash
curl -X PUT http://localhost:8080/events/12345678-1234-1234-1234-123456789012 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Event Title",
    "description": "Updated description",
    "date_time": "2024-01-16T19:00:00Z"
  }'
```

**Response:** `200 OK`
```json
{
  "id": "12345678-1234-1234-1234-123456789012",
  "group_id": "group-uuid",
  "title": "Updated Event Title",
  "description": "Updated description",
  "date_time": "2024-01-16T19:00:00Z",
  "created_by": "user-uuid",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-02T00:00:00Z"
}
```

### Delete an Event

```bash
curl -X DELETE http://localhost:8080/events/12345678-1234-1234-1234-123456789012
```

**Response:** `204 No Content`

### Create/Update RSVP Response

```bash
curl -X POST http://localhost:8080/events/12345678-1234-1234-1234-123456789012/responses \
  -H "Content-Type: application/json" \
  -d '{"response": "going"}'
```

**Response:** `201 Created`
```json
{
  "id": "response-uuid",
  "event_id": "12345678-1234-1234-1234-123456789012",
  "user_id": "user-uuid-from-session",
  "response": "going",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

**Valid response values:** `going`, `not_going`, `maybe`

### Get Event Responses

```bash
curl -X GET http://localhost:8080/events/12345678-1234-1234-1234-123456789012/responses
```

**Response:** `200 OK`
```json
[
  {
    "id": "response-uuid",
    "event_id": "12345678-1234-1234-1234-123456789012",
    "user_id": "user-uuid",
    "response": "going",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### Delete RSVP Response

```bash
curl -X DELETE http://localhost:8080/events/12345678-1234-1234-1234-123456789012/responses
```

**Response:** `204 No Content`

## Authentication

The following endpoints require authentication (user must be logged in):
- `POST /events` - Creates event with `created_by` from session
- `POST /events/{id}/responses` - Creates response with `user_id` from session
- `DELETE /events/{id}/responses` - Deletes response for logged-in user

The authentication middleware sets the `userID` in the request context. Without it, these endpoints will return `401 Unauthorized`.

## Error Responses

| Status Code | Description |
|-------------|-------------|
| 400 | Bad Request - Invalid input data |
| 401 | Unauthorized - Missing or invalid session |
| 404 | Not Found - Event or response doesn't exist |
| 500 | Internal Server Error - Database or server error |