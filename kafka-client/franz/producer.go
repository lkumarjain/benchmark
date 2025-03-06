package franz

import (
	"context"
	"crypto/tls"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
)

type Producer struct {
	BootstrapServers []string
	wg               *sync.WaitGroup
	instance         *kgo.Client
}

func NewProducer(bootstrapServers string, authenticator bool, userName string, password string) *Producer {
	brokers := strings.Split(bootstrapServers, ",")
	opts := []kgo.Opt{kgo.SeedBrokers(brokers...)}

	if authenticator {
		tlsDialer := &tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}}
		opts = append(opts, kgo.SASL(plain.Auth{User: userName, Pass: password}.AsMechanism()))
		opts = append(opts, kgo.Dialer(tlsDialer.DialContext))
	}

	instance, err := kgo.NewClient(opts...)
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
