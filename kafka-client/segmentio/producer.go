package segmentio

import (
	"context"
	"crypto/tls"
	"strings"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

type Producer struct {
	BootstrapServers []string
	wg               *sync.WaitGroup
	syncWriter       *kafka.Writer
	asyncWriter      *kafka.Writer
}

func NewProducer(bootstrapServers string, topic string, authenticator bool, userName string, password string) *Producer {
	brokers := strings.Split(bootstrapServers, ",")

	transport := &kafka.Transport{
		SASL:        plain.Mechanism{Username: userName, Password: password},
		DialTimeout: 10 * time.Second,
		TLS:         &tls.Config{InsecureSkipVerify: true},
	}

	syncWriter := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.RoundRobin{},
		MaxAttempts:  10,
		BatchSize:    100,
		BatchTimeout: 10 * time.Millisecond,
	}

	if authenticator {
		syncWriter.Transport = transport
	}

	asyncWriter := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.RoundRobin{},
		MaxAttempts:  10,
		BatchSize:    100,
		BatchTimeout: 10 * time.Millisecond,
		Async:        true,
	}

	if authenticator {
		asyncWriter.Transport = transport
	}

	producer := &Producer{BootstrapServers: brokers, syncWriter: syncWriter, asyncWriter: asyncWriter, wg: &sync.WaitGroup{}}
	producer.asyncWriter.Completion = producer.DeliveryReport

	return producer
}

func (p *Producer) ProduceSync(key string, value string) {
	err := p.syncWriter.WriteMessages(context.Background(), kafka.Message{Key: []byte(key), Value: []byte(value)})
	if err != nil {
		panic(err)
	}
}

func (p *Producer) ProduceAsync(key string, value string) {
	p.wg.Add(1)
	err := p.asyncWriter.WriteMessages(context.Background(), kafka.Message{Key: []byte(key), Value: []byte(value)})
	if err != nil {
		panic(err)
	}
}

func (p *Producer) DeliveryReport(messages []kafka.Message, err error) {
	if err != nil {
		panic(err)
	}

	for _, v := range messages {
		_ = v
		p.wg.Done()
	}
}

func (p *Producer) Wait() {
	p.wg.Wait()
}
