package franz

import (
	"context"
	"strings"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	BootstrapServers []string
	wg               *sync.WaitGroup
	instance         *kgo.Client
}

func NewProducer(bootstrapServers string) *Producer {
	brokers := strings.Split(bootstrapServers, ",")
	instance, err := kgo.NewClient(kgo.SeedBrokers(brokers...))
	if err != nil {
		panic(err)
	}

	producer := &Producer{BootstrapServers: brokers, wg: &sync.WaitGroup{}, instance: instance}

	return producer
}

func (p *Producer) ProduceSync(topic string, key string, value string) {
	p.instance.ProduceSync(context.Background(), &kgo.Record{Topic: topic, Key: []byte(key), Value: []byte(value)})
}

func (p *Producer) ProduceAsync(topic string, key string, value string) {
	p.wg.Add(1)
	p.instance.Produce(context.Background(), &kgo.Record{Topic: topic, Key: []byte(key), Value: []byte(value)}, p.DeliveryReport)
}

func (p *Producer) DeliveryReport(_ *kgo.Record, _ error) {
	p.wg.Done()
}

func (p *Producer) Wait() {
	p.wg.Wait()
}
