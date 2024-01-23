package franz

import (
	"context"
	"strings"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer struct {
	BootstrapServers string
	instance         *kgo.Client
	wg               *sync.WaitGroup
}

func NewConsumer(bootstrapServers string, topicName string) *Consumer {
	brokers := strings.Split(bootstrapServers, ",")

	instance, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup("franz-consumer-group"),
		kgo.ConsumeTopics(topicName),
		kgo.DisableAutoCommit(),
	)

	if err != nil {
		panic(err)
	}

	return &Consumer{BootstrapServers: bootstrapServers, instance: instance, wg: &sync.WaitGroup{}}
}

func (c *Consumer) Consume(message chan interface{}, done chan bool) {
	run := true

	for run {
		select {
		case <-done:
			run = false
		default:
			fetches := c.instance.PollFetches(context.Background())
			if fetches.IsClientClosed() {
				return
			}
			fetches.EachRecord(func(r *kgo.Record) {
				select {
				case <-done:
					run = false
					return
				case message <- r:
				}
			})
		}
	}
}

func (c *Consumer) Close() {
	c.instance.Close()
}
