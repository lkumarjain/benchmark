package kafkaclient

import (
	"sync"
	"testing"

	"github.com/lkumarjain/benchmark/kafka-client/confluent"
	"github.com/lkumarjain/benchmark/kafka-client/franz"
	"github.com/lkumarjain/benchmark/kafka-client/goka"
	"github.com/lkumarjain/benchmark/kafka-client/sarama"
	"github.com/lkumarjain/benchmark/kafka-client/segmentio"
)

func BenchmarkConfluentConsumer(b *testing.B) {
	consumer := confluent.Consumer{Servers: bootstrapServers, EnableEvents: false, Topic: topicName}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go consumer.Start(wg)
	wg.Wait()

	b.Run("Confluent@ConsumePoll", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}
		b.StopTimer()
	})

	close(consumer.Done)

	consumer.EnableEvents = true

	wg.Add(1)
	go consumer.Start(wg)
	wg.Wait()

	b.Run("Confluent@ConsumeEvent", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}
		b.StopTimer()
	})

	close(consumer.Done)
}

func BenchmarkFranzConsumer(b *testing.B) {
	consumer := franz.Consumer{Servers: bootstrapServers, EnablePartition: false, Topic: topicName}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go consumer.Start(wg)
	wg.Wait()

	b.Run("Franz@ConsumeRecord", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}
		b.StopTimer()
	})

	close(consumer.Done)

	consumer.EnablePartition = true

	wg.Add(1)
	go consumer.Start(wg)
	wg.Wait()

	b.Run("Confluent@ConsumePartition", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}
		b.StopTimer()
	})

	close(consumer.Done)
}

func BenchmarkGokaConsumer(b *testing.B) {
	consumer := goka.Consumer{Servers: bootstrapServers, Topic: topicName}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go consumer.Start(wg)
	wg.Wait()

	b.Run("Goka@Consumer", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}
		b.StopTimer()
	})

	close(consumer.Done)
}

func BenchmarkSaramaConsumer(b *testing.B) {
	consumer := sarama.Consumer{Servers: bootstrapServers, Topic: topicName, EnablePartition: false}

	wg := &sync.WaitGroup{}
	// wg.Add(1)
	// go consumer.Start(wg)
	// wg.Wait()

	// b.Run("Sarama@ConsumerGroup", func(b *testing.B) {
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		<-consumer.Message
	// 	}
	// 	b.StopTimer()
	// })

	// close(consumer.Done)

	consumer.EnablePartition = true

	wg.Add(1)
	go consumer.Start(wg)
	wg.Wait()

	b.Run("Sarama@ConsumePartition", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}
		b.StopTimer()
	})

	close(consumer.Done)
}

func BenchmarkSegmentioConsumer(b *testing.B) {
	consumer := segmentio.NewConsumer(bootstrapServers, topicName)
	message := make(chan interface{}, 1)
	done := make(chan bool, 1)
	go consumer.Consume(message, done)

	b.Run("Segmentio@Consumer", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-message
		}
		b.StopTimer()
	})

	done <- true
}
