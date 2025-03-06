package sarama

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
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

	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_8_0_0
	cfg.Consumer.Return.Errors = true
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	if c.Authenticator {
		cfg.Net.TLS.Enable = true
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = c.UserName
		cfg.Net.SASL.Password = c.Password
		cfg.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}

	if c.EnablePartition {
		c.consumePartition(cfg)
	} else {
		c.consumeGroup(cfg)
	}
}

func (c *Consumer) consumePartition(cfg *sarama.Config) {
	brokers := strings.Split(c.Servers, ",")

	consumer, err := sarama.NewConsumer(brokers, cfg)
	if err != nil {
		log.Panicf("error creating processor: %v", err)
		return
	}
	partitions, err := consumer.Partitions(c.Topic)
	if err != nil {
		log.Panicf("error creating processor: %v", err)
		return
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(c.Topic, partition, sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("Failed to create consumer: %v\n", err)
			return
		}

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				select {
				case <-c.Done:
					return
				default:
					c.Message <- message
				}
			}
		}(pc)
	}
}

func (c *Consumer) consumeGroup(cfg *sarama.Config) {
	brokers := strings.Split(c.Servers, ",")

	group, err := sarama.NewConsumerGroup(brokers, fmt.Sprintf("sarama-consumer-group-%d", time.Now().UnixNano()), cfg)

	if err != nil {
		log.Panicf("error creating processor: %v", err)
		return
	}

	handler := handler{message: c.Message, done: c.Done}

	go func() {
		defer group.Close()
		run := true

		for run {
			select {
			case <-c.Done:
				run = false
			default:
				group.Consume(context.Background(), []string{c.Topic}, handler)
			}
		}
	}()
}

type handler struct {
	message chan interface{}
	done    chan bool
}

func (h handler) Setup(sarama.ConsumerGroupSession) error { return nil }

func (handler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h handler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	run := true
	for run {
		select {
		case <-h.done:
			run = false
		default:
			h.message <- claim.Messages()
		}
	}
	return nil
}
