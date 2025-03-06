package segmentio

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

type Consumer struct {
	Servers         string
	Topic           string
	EnablePartition bool
	Message         chan interface{}
	Done            chan bool
	Authenticator   bool
	UserName        string
	Password        string
}

func (c *Consumer) Start() {
	c.Message = make(chan interface{}, 1)
	c.Done = make(chan bool, 1)

	config := kafka.ReaderConfig{
		Brokers: strings.Split(c.Servers, ","),
		Topic:   c.Topic,
		GroupID: fmt.Sprintf("segmentio-consumer-group-%d", time.Now().UnixNano()),
	}

	if c.Authenticator {
		dialer := &kafka.Dialer{
			Timeout:       10 * time.Second,
			DualStack:     true,
			SASLMechanism: plain.Mechanism{Username: c.UserName, Password: c.Password},
			TLS:           &tls.Config{InsecureSkipVerify: true},
		}

		config.Dialer = dialer
	}

	reader := kafka.NewReader(config)

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
