package kafkaclient

import (
	"testing"

	"github.com/lkumarjain/benchmark/kafka-client/confluent"
	"github.com/lkumarjain/benchmark/kafka-client/franz"
	"github.com/lkumarjain/benchmark/kafka-client/goka"
	"github.com/lkumarjain/benchmark/kafka-client/sarama"
	"github.com/lkumarjain/benchmark/kafka-client/segmentio"
)

func BenchmarkConfluentConsumer(b *testing.B) {
	consumer := confluent.NewConsumer(bootstrapServers)
	message := make(chan interface{}, 1)
	done := make(chan bool, 1)
	go consumer.Consume(topicName, message, done)

	b.Run("Confluent@Consume", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-message
		}
		b.StopTimer()
	})

	done <- true
	consumer.Close()
}
func BenchmarkGokaConsumer(b *testing.B) {
	consumer := goka.NewConsumer(bootstrapServers, topicName)
	message := make(chan interface{}, 1)
	done := make(chan bool, 1)
	go consumer.Consume(topicName, message, done)

	b.Run("Goka@Consumer", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-message
		}
		b.StopTimer()
	})

	close(done)
}

func BenchmarkFranzConsumer(b *testing.B) {
	consumer := franz.NewConsumer(bootstrapServers, topicName)
	message := make(chan interface{}, 1)
	done := make(chan bool, 1)
	go consumer.Consume(message, done)

	b.Run("Franz@Consumer", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-message
		}
		b.StopTimer()
	})

	done <- true
}

func BenchmarkSaramaConsumer(b *testing.B) {
	consumer := sarama.NewConsumer(bootstrapServers)
	message := make(chan interface{}, 1)
	done := make(chan bool, 1)
	ready := make(chan bool, 1)

	go consumer.Consume(topicName, message, done, ready)

	b.Run("Sarama@Consumer", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-message
		}
		b.StopTimer()
	})

	close(done)
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
