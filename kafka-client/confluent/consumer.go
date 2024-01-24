package confluent

import (
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	BootstrapServers string
	instance         *kafka.Consumer
	wg               *sync.WaitGroup
}

func NewConsumer(bootstrapServers string) *Consumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers":        bootstrapServers,
		"group.id":                 "confluent-consumer-group",
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.commit":       "true",
		"enable.auto.offset.store": "false",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		fmt.Printf("Failed to create consumer: %v\n", err)
		return nil
	}

	return &Consumer{BootstrapServers: bootstrapServers, instance: consumer, wg: &sync.WaitGroup{}}
}

func (c *Consumer) Consume(topic string, message chan interface{}, done chan bool) {
	c.instance.SubscribeTopics([]string{topic}, nil)

	run := true

	for run {
		select {
		case <-done:
			run = false
		default:
			event := c.instance.Poll(1000)
			if event == nil {
				continue
			}
			switch e := event.(type) {
			case *kafka.Message:
				message <- e
			case kafka.Error:
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			}
		}
	}
}

func (c *Consumer) Close() {
	c.instance.Close()
}
