package confluent

import (
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	BootstrapServers string
	instance         *kafka.Producer
	wg               *sync.WaitGroup
}

func NewProducer(bootstrapServers string, authenticator bool, userName string, password string) *Producer {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":    bootstrapServers,
		"log.connection.close": false,
		"enable.metrics.push":  false,
	}

	if authenticator {
		cfg.SetKey("sasl.mechanisms", "PLAIN")
		cfg.SetKey("sasl.username", userName)
		cfg.SetKey("sasl.password", password)
		cfg.SetKey("security.protocol", "SASL_SSL")
	}

	p, err := kafka.NewProducer(cfg)
	if err != nil {
		panic(err)
	}

	producer := &Producer{BootstrapServers: bootstrapServers, instance: p, wg: &sync.WaitGroup{}}

	go producer.DeliveryReport()

	return producer
}

func (p *Producer) ProduceSync(topic string, key string, value string) {
	deliveryChan := make(chan kafka.Event, 1)
	p.instance.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(value),
	}, deliveryChan)

	<-deliveryChan
}

func (p *Producer) ProduceAsync(topic string, key string, value string) {
	p.wg.Add(1)
	p.instance.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(value),
	}, nil)
}

func (p *Producer) DeliveryReport() {
	for e := range p.instance.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			p.wg.Done()
			if ev.TopicPartition.Error != nil {
				panic(ev.TopicPartition.Error)
			}
		}
	}
}

func (p *Producer) Wait() {
	p.wg.Wait()
}
