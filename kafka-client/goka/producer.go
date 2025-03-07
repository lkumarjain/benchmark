package goka

import (
	"log"
	"strings"
	"sync"

	"github.com/IBM/sarama"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

type Producer struct {
	BootstrapServers []string
	wg               *sync.WaitGroup
	instance         *goka.Emitter
}

func NewProducer(bootstrapServers string, topic string, authenticator bool, userName string, password string) *Producer {
	brokers := strings.Split(bootstrapServers, ",")

	config := goka.DefaultConfig()
	config.Version = sarama.V3_6_0_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	if authenticator {
		config.Net.TLS.Enable = true
		config.Net.SASL.Enable = true
		config.Net.SASL.User = userName
		config.Net.SASL.Password = password
		config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}

	emitter, err := goka.NewEmitter(brokers, goka.Stream(topic),
		new(codec.String), goka.WithEmitterProducerBuilder(goka.ProducerBuilderWithConfig(config)))

	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}

	producer := &Producer{BootstrapServers: brokers, wg: &sync.WaitGroup{}, instance: emitter}

	return producer
}

func (p *Producer) ProduceSync(key string, value string) {
	err := p.instance.EmitSync(key, value)
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
	}
}

func (p *Producer) ProduceAsync(key string, value string) {
	p.wg.Add(1)
	promise, err := p.instance.Emit(key, value)
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
		p.wg.Done()
		return
	}

	promise.Then(p.DeliveryReport)
}

func (p *Producer) DeliveryReport(err error) {
	defer p.wg.Done()
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
	}
}

func (p *Producer) Wait() {
	p.wg.Wait()
}
