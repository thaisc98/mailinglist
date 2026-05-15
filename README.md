# Mailing List Service

A Go microservice for managing email subscriptions that hosts two APIs — a JSON REST API and a gRPC API. It uses Protocol Buffers as the communication format for the gRPC layer and SQLite as the persistent data store.

## Architecture

```
┌─────────────┐     ┌──────────────────────────────────────────┐
│   client/   │────▶│              server/                     │
└─────────────┘     │  JSON API (:8080) + gRPC API (:8081)    │
                    └──────────────────┬───────────────────────┘
                                       │
                               ┌───────▼───────┐
                               │  mdb/ (SQLite) │
                               └───────────────┘
```

- **`mdb/`** — SQLite data layer (CRUD + batch queries)
- **`jsonapi/`** — HTTP/JSON REST API
- **`grpcapi/`** — gRPC API (generated from `proto/mail.proto`)
- **`client/`** — Example gRPC client
- **`server/`** — Entrypoint that runs both servers concurrently

## Prerequisites

- Go 1.21+
- GCC (required by `go-sqlite3`)
- `protoc` (only needed to regenerate proto files)

### Install protoc

Follow the instructions at [grpc.io/docs/protoc-installation](https://grpc.io/docs/protoc-installation/) or download a pre-built binary from the [protobuf releases page](https://github.com/protocolbuffers/protobuf/releases) and place it in your `$PATH`.

### Install Go protobuf code-gen plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Setup

```bash
git clone <repo-url>
cd mailinglist
go mod download
```

## Running

### Server

```bash
go run ./server
```

The server starts two listeners:

| API   | Default address |
|-------|-----------------|
| JSON  | `:8080`         |
| gRPC  | `:8081`         |

### Client (gRPC example)

In a separate terminal:

```bash
go run ./client
```

## Configuration

All options are set via environment variables:

| Variable               | Default    | Description              |
|------------------------|------------|--------------------------|
| `MAILINGLIST_DB`       | `list.db`  | Path to the SQLite file  |
| `MAILINGLIST_BIND_JSON`| `:8080`    | JSON API bind address    |
| `MAILINGLIST_BIND_GRPC`| `:8081`    | gRPC bind address        |

Example:

```bash
MAILINGLIST_DB=/data/mail.db MAILINGLIST_BIND_JSON=:9090 go run ./server
```

## JSON API

Base URL: `http://localhost:8080`

| Method | Endpoint           | Description                        |
|--------|--------------------|------------------------------------|
| POST   | `/email/create`    | Add a new email address            |
| GET    | `/email/get`       | Look up a single email             |
| GET    | `/email/get_batch` | Paginated list of active emails    |
| PUT    | `/email/update`    | Update confirmation status/opt-out |
| POST   | `/email/delete`    | Opt-out (soft delete) an email     |

All request and response bodies are JSON. The `EmailEntry` shape:

```json
{
  "Id": 1,
  "Email": "user@example.com",
  "ConfirmedAt": "2024-01-15T10:00:00Z",
  "OptOut": false
}
```

Batch request body:

```json
{ "Page": 1, "Count": 20 }
```

## gRPC API

Defined in [`proto/mail.proto`](proto/mail.proto). Service methods mirror the JSON API:

- `CreateEmail`
- `GetEmail`
- `UpdateEmail`
- `DeleteEmail`
- `GetEmailBatch`

### Regenerate proto files

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/mail.proto
```

## Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/mattn/go-sqlite3` | SQLite driver (CGO) |
| `github.com/alexflint/go-arg` | Environment variable / flag parsing |
| `google.golang.org/grpc` | gRPC framework |
| `google.golang.org/protobuf` | Protocol Buffers runtime |
