package franz

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer struct {
	Servers         string
	EnablePartition bool
	Topic           string
	Message         chan interface{}
	Done            chan bool
}

func (c *Consumer) Start() {
	c.Message = make(chan interface{}, 1)
	c.Done = make(chan bool, 1)

	brokers := strings.Split(c.Servers, ",")

	consumer, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(fmt.Sprintf("franz-consumer-group-%d", time.Now().UnixNano())),
		kgo.ConsumeTopics(c.Topic),
		kgo.DisableAutoCommit(),
	)

	if err != nil {
		fmt.Printf("Failed to create consumer: %v\n", err)
		return
	}

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
