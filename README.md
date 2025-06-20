# 🛠️ RabbitMQ Worker Example with MongoDB, Redis, and Structured Logging

This is a minimal Go example showcasing how to build a distributed system with:
- 🟡 RabbitMQ (message queue)
- 🟢 MongoDB (persistent storage)
- 🔴 Redis (caching)
- ✅ Health checks
- 🔁 Retry logic
- 📋 Structured logging with `log/slog`

---

## 📦 Components

- `cmd/rest/`: REST API that enqueues jobs to RabbitMQ and stores metadata in Redis.
- `cmd/worker/`: Worker that consumes jobs, processes them, and persists to MongoDB.
- `internal/infra/`: Modular packages for Redis, Mongo, and RabbitMQ clients.

---

## 🚀 Getting Started

### Prerequisites

Install:
- [Go 1.24+](https://go.dev/)
- Docker (for services)

Usage:

``` shell
make build         # Build binaries
make podman-build  # Build container images
make up            # Launch Mongo, Redis, RabbitMQ, rest, worker
make down          # Tear them all down
make run-rest      # Run rest locally
make run-worker    # Run worker locally
make fmt           # Format code
make tidy          # Tidy go.mod
make lint          # Run the linter
```
