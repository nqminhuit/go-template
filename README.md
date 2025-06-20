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
- [Go 1.22+](https://go.dev/)
- Docker (for services)

Start services with Docker:
```bash
# RabbitMQ
docker run -d --name rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management

# MongoDB
docker run -d --name mongo -p 27017:27017 mongo

# Redis
docker run -d --name redis -p 6379:6379 redis
