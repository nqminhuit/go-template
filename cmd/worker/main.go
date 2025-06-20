package main

import (
	"context"
	"encoding/json"
	"log/slog"
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

	mongoClient, err := mongo.NewMongoClient("mongodb://localhost:27017", "myapp")
	if err != nil {
		slog.Error("failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}
	defer mongoClient.Disconnect(ctx)

	redisClient := redis.NewRedisClient("localhost:6379")

	handler := func(ctx context.Context, body []byte) error {
		var job map[string]string
		if err := json.Unmarshal(body, &job); err != nil {
			return err
		}
		slog.Info("job received", "job", job)
		if err := redisClient.Set(ctx, "last_processed_job", job["task"]); err != nil {
			slog.Warn("redis set error", "error", err)
		}
		coll := mongoClient.DB.Collection("jobs")
		_, err := coll.InsertOne(ctx, job)
		return err
	}

	consumer, err := amqp.NewConsumer(client.Channel, "jobs", handler)
	if err != nil {
		slog.Error("failed to create consumer", "error", err)
		os.Exit(1)
	}
	if err := consumer.Start(ctx); err != nil {
		slog.Error("failed to start consumer", "error", err)
		os.Exit(1)
	}
	slog.Info("Worker listening for jobs...")
	select {}
}
