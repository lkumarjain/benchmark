package segmentio

import (
	"context"
	"fmt"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Consumer struct {
	Servers         string
	Topic           string
	EnablePartition bool
	Message         chan interface{}
	Done            chan bool
}

func (c *Consumer) Start() {
	c.Message = make(chan interface{}, 1)
	c.Done = make(chan bool, 1)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(c.Servers, ","),
		Topic:   c.Topic,
		GroupID: fmt.Sprintf("segmentio-consumer-group-%d", time.Now().UnixNano()),
	})

	go func() {
		defer reader.Close()

		run := true

		for run {
			select {
			case <-c.Done:
				run = false
			default:
				msg, _ := reader.FetchMessage(context.Background())
				c.Message <- msg
			}
		}
	}()
}
