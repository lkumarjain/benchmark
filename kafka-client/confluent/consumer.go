package confluent

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	Servers      string
	EnableEvents bool
	Topic        string
	Message      chan interface{}
	Done         chan bool
}

func (c *Consumer) Start() {
	c.Message = make(chan interface{}, 1)
	c.Done = make(chan bool, 1)

	config := &kafka.ConfigMap{
		"bootstrap.servers":        c.Servers,
		"group.id":                 fmt.Sprintf("confluent-consumer-group-%d", time.Now().UnixNano()),
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.commit":       "true",
		"enable.auto.offset.store": "false",
		"go.events.channel.enable": c.EnableEvents,
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		fmt.Printf("Failed to create consumer: %v\n", err)
		return
	}

	consumer.SubscribeTopics([]string{c.Topic}, nil)

	go func() {
		run := true

		for run {
			select {
			case <-c.Done:
				run = false
			default:
				run = c.start(consumer)
			}
		}

		consumer.Close()
	}()

}

func (c *Consumer) start(consumer *kafka.Consumer) bool {
	var event kafka.Event

	if c.EnableEvents {
		event = <-consumer.Events()
	} else {
		event = consumer.Poll(1000)
	}

	if event == nil {
		return true
	}

	switch e := event.(type) {
	case *kafka.Message:
		c.Message <- e
	case kafka.Error:
		if e.Code() == kafka.ErrAllBrokersDown {
			return false
		}
	}

	return true
}
