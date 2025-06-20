package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewClient(url string) (*Client, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	return &Client{Conn: conn, Channel: ch}, nil
}

func (c *Client) Close() {
	c.Channel.Close()
	c.Conn.Close()
}
