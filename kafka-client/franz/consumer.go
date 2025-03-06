package franz

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
)

type Consumer struct {
	Servers         string
	EnablePartition bool
	Topic           string
	Message         chan interface{}
	Done            chan bool
	Authenticator   bool
	UserName        string
	Password        string
}

func (c *Consumer) Start() {
	c.Message = make(chan interface{}, 1)
	c.Done = make(chan bool, 1)

	brokers := strings.Split(c.Servers, ",")

	opts := []kgo.Opt{
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(fmt.Sprintf("franz-consumer-group-%d", time.Now().UnixNano())),
		kgo.ConsumeTopics(c.Topic),
		kgo.DisableAutoCommit(),
	}

	if c.Authenticator {
		tlsDialer := &tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}}
		opts = append(opts, kgo.SASL(plain.Auth{User: c.UserName, Pass: c.Password}.AsMechanism()))
		opts = append(opts, kgo.Dialer(tlsDialer.DialContext))
	}

	consumer, err := kgo.NewClient(opts...)

	if err != nil {
		log.Panicf("error creating processor: %v", err)
		return
	}

	go func() {
		defer consumer.Close()

		run := true

		for run {
			select {
			case <-c.Done:
				run = false
			default:
				run = c.start(consumer)
			}
		}
	}()
}

func (c *Consumer) start(consumer *kgo.Client) bool {
	fetches := consumer.PollFetches(context.Background())

	if fetches.IsClientClosed() {
		return false
	}

	if errs := fetches.Errors(); len(errs) > 0 {
		fmt.Printf("Failed to create fetches: %v\n", errs)
		return false
	}

	if c.EnablePartition {
		return c.consumePartition(fetches)
	}

	return c.consumeRecord(fetches)
}

func (c *Consumer) consumeRecord(fetches kgo.Fetches) bool {
	for _, fetch := range fetches {
		for _, topic := range fetch.Topics {
			for _, partition := range topic.Partitions {
				for _, record := range partition.Records {
					select {
					case <-c.Done:
						return false
					case c.Message <- record.Value:
					}
				}
			}
		}
	}
	return true
}

func (c *Consumer) consumePartition(fetches kgo.Fetches) bool {
	iter := fetches.RecordIter()
	for !iter.Done() {
		select {
		case <-c.Done:
			return false
		case c.Message <- iter.Next():
		}
	}
	return true
}
