package goka

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

type Consumer struct {
	Servers string
	Topic   string
	Message chan interface{}
	Done    chan bool
}

func (c *Consumer) Start(wg *sync.WaitGroup) {
	c.Message = make(chan interface{}, 1)
	c.Done = make(chan bool, 1)

	brokers := strings.Split(c.Servers, ",")

	cfg := goka.DefaultConfig()
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Version = sarama.V2_8_0_0
	cfg.Consumer.Return.Errors = true
	goka.ReplaceGlobalConfig(cfg)

	topicStream := goka.Stream(c.Topic)
	g := goka.DefineGroup("goka-consumer-group", goka.Input(topicStream, new(codec.String), c.handler), goka.Persist(new(codec.Int64)))

	config := goka.NewTopicManagerConfig()
	config.Table.Replication = 1
	config.CreateTopicTimeout = time.Second * 10

	p, err := goka.NewProcessor(brokers, g, goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(config)))
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
		wg.Done()
		return
	}

	wg.Done()
	ctx, cancel := context.WithCancel(context.Background())
	go p.Run(ctx)
	<-c.Done
	cancel()
}

func (c *Consumer) handler(ctx goka.Context, msg interface{}) {
	run := true
	for run {
		select {
		case <-c.Done:
			run = false
		default:
			c.Message <- msg
		}
	}
}
