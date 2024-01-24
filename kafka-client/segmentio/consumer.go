package segmentio

import (
	"context"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

type Consumer struct {
	BootstrapServers string
	reader           *kafka.Reader
}

func NewConsumer(bootstrapServers string, topic string) *Consumer {
	brokers := strings.Split(bootstrapServers, ",")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "segmentio-consumer-group",
	})

	return &Consumer{BootstrapServers: bootstrapServers, reader: reader}
}

func (c *Consumer) Consume(message chan interface{}, done chan bool) {
	run := true

	for run {
		select {
		case <-done:
			run = false
		default:
			msg, _ := c.reader.FetchMessage(context.Background())
			message <- msg
		}
	}
}
