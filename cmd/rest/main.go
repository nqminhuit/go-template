package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"go-template/internal/infra/amqp"
	"go-template/internal/infra/mongo"
	"go-template/internal/infra/redis"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	ctx := context.Background()

	client, err := amqp.NewClient("amqp://guest:guest@localhost:5672/")
	if err != nil {
		slog.Error("failed to connect to RabbitMQ", "error", err)
		os.Exit(1)
	}
	defer client.Close()

	producer, err := amqp.NewProducer(client.Channel, "jobs")
	if err != nil {
		slog.Error("failed to create producer", "error", err)
		os.Exit(1)
	}

	mongoClient, err := mongo.NewMongoClient("mongodb://localhost:27017", "myapp")
	if err != nil {
		slog.Error("failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}
	defer mongoClient.Disconnect(ctx)

	redisClient := redis.NewRedisClient("localhost:6379")

	http.HandleFunc("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		job := map[string]string{"task": "say_hello"}
		if err := redisClient.Set(ctx, "last_job", "say_hello"); err != nil {
			slog.Warn("redis set error", "error", err)
		}
		if err := producer.Publish(ctx, job); err != nil {
			slog.Error("publish error", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("job enqueued")
		w.WriteHeader(http.StatusAccepted)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if _, err := redisClient.Get(ctx, "last_job"); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		if err := mongoClient.Client.Ping(ctx, nil); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	slog.Info("REST server listening", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}
