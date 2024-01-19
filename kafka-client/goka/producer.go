package goka

import (
	"log"
	"strings"
	"sync"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

type Producer struct {
	BootstrapServers []string
	wg               *sync.WaitGroup
	instance         *goka.Emitter
}

func NewProducer(bootstrapServers string, topic string) *Producer {
	brokers := strings.Split(bootstrapServers, ",")
	emitter, err := goka.NewEmitter(brokers, goka.Stream(topic), new(codec.String))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}

	producer := &Producer{BootstrapServers: brokers, wg: &sync.WaitGroup{}, instance: emitter}

	return producer
}

func (p *Producer) ProduceSync(key string, value string) {
	p.instance.EmitSync(key, value)
}

func (p *Producer) ProduceAsync(key string, value string) {
	p.wg.Add(1)
	promise, err := p.instance.Emit(key, value)
	if err != nil {
		p.wg.Done()
		return
	}

	promise.Then(p.DeliveryReport)
}

func (p *Producer) DeliveryReport(err error) {
	defer p.wg.Done()
}

func (p *Producer) Wait() {
	p.wg.Wait()
}
