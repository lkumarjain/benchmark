package segmentio

import (
	"context"
	"strings"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Producer struct {
	BootstrapServers []string
	wg               *sync.WaitGroup
	syncWriter       *kafka.Writer
	asyncWriter      *kafka.Writer
}

func NewProducer(bootstrapServers string, topic string) *Producer {
	brokers := strings.Split(bootstrapServers, ",")

	syncWriter := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.RoundRobin{},
		MaxAttempts:  10,
		BatchSize:    100,
		BatchTimeout: 10 * time.Millisecond,
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

	producer := &Producer{BootstrapServers: brokers, syncWriter: syncWriter, asyncWriter: asyncWriter, wg: &sync.WaitGroup{}}
	producer.asyncWriter.Completion = producer.DeliveryReport

	return producer
}

func (p *Producer) ProduceSync(key string, value string) {
	p.syncWriter.WriteMessages(context.Background(), kafka.Message{Key: []byte(key), Value: []byte(value)})
}

func (p *Producer) ProduceAsync(key string, value string) {
	p.wg.Add(1)
	p.asyncWriter.WriteMessages(context.Background(), kafka.Message{Key: []byte(key), Value: []byte(value)})
}

func (p *Producer) DeliveryReport(messages []kafka.Message, err error) {
	for _, v := range messages {
		_ = v
		p.wg.Done()
	}
}

func (p *Producer) Wait() {
	p.wg.Wait()
}
