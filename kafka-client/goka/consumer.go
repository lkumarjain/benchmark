package goka

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

type Consumer struct {
	Servers       string
	Topic         string
	Message       chan interface{}
	Done          chan bool
	Authenticator bool
	UserName      string
	Password      string
}

func (c *Consumer) Start() {
	c.Message = make(chan interface{}, 1)
	c.Done = make(chan bool, 1)

	brokers := strings.Split(c.Servers, ",")

	cfg := goka.DefaultConfig()
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Version = sarama.V2_8_0_0
	cfg.Consumer.Return.Errors = true

	if c.Authenticator {
		cfg.Net.TLS.Enable = true
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = c.UserName
		cfg.Net.SASL.Password = c.Password
		cfg.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}

	goka.ReplaceGlobalConfig(cfg)

	topicStream := goka.Stream(c.Topic)
	group := goka.Group(fmt.Sprintf("goka-consumer-group-%d", time.Now().UnixNano()))
	g := goka.DefineGroup(group,
		goka.Input(topicStream, new(codec.String), c.handler), goka.Persist(new(codec.Int64)))

	config := goka.NewTopicManagerConfig()
	config.Table.Replication = 1
	config.CreateTopicTimeout = time.Minute

	log := log.New(io.Discard, "", log.LstdFlags)

	p, err := goka.NewProcessor(brokers, g,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(config)),
		goka.WithLogger(log))
	if err != nil {
		log.Panicf("error creating processor: %v", err)
		return
	}

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		go p.Run(ctx)
		<-c.Done
		cancel()
	}()
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
