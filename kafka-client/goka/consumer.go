package goka

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

type Consumer struct {
	BootstrapServers string
}

func NewConsumer(bootstrapServers string, topic string) *Consumer {
	return &Consumer{BootstrapServers: bootstrapServers}
}

func (c *Consumer) Consume(topic string, message chan interface{}, done chan bool) {
	brokers := strings.Split(c.BootstrapServers, ",")

	cfg := goka.DefaultConfig()
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Version = sarama.V2_8_0_0
	cfg.Consumer.Return.Errors = true
	goka.ReplaceGlobalConfig(cfg)

	cb := consumer{message: message, done: done}

	topicStream := goka.Stream(topic)
	g := goka.DefineGroup("goka-consumer-group", goka.Input(topicStream, new(codec.String), cb.handler), goka.Persist(new(codec.Int64)))

	config := goka.NewTopicManagerConfig()
	config.Table.Replication = 1
	config.CreateTopicTimeout = time.Second * 10
	p, err := goka.NewProcessor(brokers, g, goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(config)))
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go p.Run(ctx)

	<-done
	cancel()
}

type consumer struct {
	message chan interface{}
	done    chan bool
}

func (c consumer) handler(ctx goka.Context, msg interface{}) {
	run := true
	for run {
		select {
		case <-c.done:
			run = false
		default:
			c.message <- msg
		}
	}
}
