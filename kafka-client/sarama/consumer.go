package sarama

import (
	"context"
	"strings"

	"github.com/IBM/sarama"
)

type Consumer struct {
	BootstrapServers string
	group            sarama.ConsumerGroup
}

func NewConsumer(bootstrapServers string) *Consumer {
	brokers := strings.Split(bootstrapServers, ",")

	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_8_0_0
	cfg.Consumer.Return.Errors = true
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	group, err := sarama.NewConsumerGroup(brokers, "sarama-consumer-group", cfg)

	if err != nil {
		panic(err)
	}

	return &Consumer{BootstrapServers: bootstrapServers, group: group}
}

func (c *Consumer) Consume(topic string, message chan interface{}, done chan bool, ready chan bool) {
	run := true
	con := consumer{message: message, done: done, ready: ready}

	for run {
		select {
		case <-done:
			run = false
		default:
			c.group.Consume(context.Background(), []string{topic}, con)
		}
	}
}

func (c *Consumer) Close() {
	c.group.Close()
}

type consumer struct {
	message chan interface{}
	done    chan bool
	ready   chan bool
}

func (c consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (c consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	run := true
	for run {
		select {
		case <-c.done:
			run = false
		default:
			c.message <- claim.Messages()
		}
	}
	return nil
}
