package amqp

import (
	"context"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	ch    *amqp.Channel
	queue string
	fn    func(context.Context, []byte) error
}

func NewConsumer(ch *amqp.Channel, queueName string, fn func(context.Context, []byte) error) (*Consumer, error) {
	_, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &Consumer{ch: ch, queue: queueName, fn: fn}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	msgs, err := c.ch.Consume(c.queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			var attempt int
			for attempt < 3 {
				if err := c.fn(ctx, d.Body); err != nil {
					slog.Warn("handler error", "attempt", attempt+1, "error", err)
					time.Sleep(time.Second * time.Duration(attempt+1))
					attempt++
					continue
				}
				break
			}
		}
	}()
	return nil
}
