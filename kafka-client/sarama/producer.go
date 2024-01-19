package sarama

import (
	"strings"
	"sync"

	"github.com/IBM/sarama"
)

type Producer struct {
	BootstrapServers []string
	wg               *sync.WaitGroup
	syncInstance     sarama.SyncProducer
	asyncInstance    sarama.AsyncProducer
}

func NewProducer(bootstrapServers string) *Producer {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_1_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	brokers := strings.Split(bootstrapServers, ",")

	syncInstance, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	asyncInstance, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	producer := &Producer{BootstrapServers: brokers, wg: &sync.WaitGroup{}, syncInstance: syncInstance, asyncInstance: asyncInstance}

	go producer.DeliveryReport()

	return producer
}

func (p *Producer) ProduceSync(topic string, key string, value string) {
	p.syncInstance.SendMessage(&sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.StringEncoder(value)})
}

func (p *Producer) ProduceAsync(topic string, key string, value string) {
	p.wg.Add(1)
	p.asyncInstance.Input() <- &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.StringEncoder(value)}
}

func (p *Producer) DeliveryReport() {
	for range p.asyncInstance.Successes() {
		p.wg.Done()
	}
}

func (p *Producer) Wait() {
	p.wg.Wait()
}
