package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	ch    *amqp.Channel
	queue amqp.Queue
}

func NewProducer(ch *amqp.Channel, queueName string) (*Producer, error) {
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &Producer{ch: ch, queue: q}, nil
}

func (p *Producer) Publish(ctx context.Context, msg interface{}) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	var attempt int
	for attempt < 3 {
		err := p.ch.PublishWithContext(ctx,
			"", p.queue.Name, false, false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err == nil {
			return nil
		}
		slog.Warn("publish failed", "attempt", attempt+1, "error", err)
		time.Sleep(time.Second * time.Duration(attempt+1))
		attempt++
	}
	return err
}
