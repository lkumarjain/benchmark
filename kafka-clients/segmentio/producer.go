package segmentio

import (
	"context"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Producer struct {
	BootstrapServers []string
	writer           *kafka.Writer
}

func NewProducer(bootstrapServers string, topic string) *Producer {
	brokers := strings.Split(bootstrapServers, ",")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:       brokers,
		Topic:         topic,
		Balancer:      &kafka.Hash{},
		BatchTimeout:  time.Duration(1) * time.Millisecond,
		QueueCapacity: 1000,
		BatchSize:     1000000,
	})

	producer := &Producer{BootstrapServers: brokers, writer: writer}
	return producer
}

func (p *Producer) Produce(key string, value string) {
	p.writer.WriteMessages(context.Background(), kafka.Message{Key: []byte(key), Value: []byte(value)})
}
