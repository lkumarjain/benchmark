package kafkaclient

import (
	"testing"

	"github.com/lkumarjain/benchmark/kafka-client/confluent"
	"github.com/lkumarjain/benchmark/kafka-client/franz"
)

func BenchmarkConfluentConsumer(b *testing.B) {
	consumer := confluent.NewConsumer(bootstrapServers)

	b.Run("Confluent@Consume", func(b *testing.B) {
		message := make(chan interface{}, 1)
		done := make(chan bool, 1)
		go consumer.Consume(topicName, message, done)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-message
		}

		done <- true
		b.StopTimer()
	})

	consumer.Close()
}

func BenchmarkFranzConsumer(b *testing.B) {
	consumer := franz.NewConsumer(bootstrapServers, topicName)

	b.Run("Franz@Consumer", func(b *testing.B) {
		message := make(chan interface{}, 1)
		done := make(chan bool, 1)
		go consumer.Consume(message, done)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-message
		}

		done <- true
		b.StopTimer()
	})
}

// func benchmarkGokaProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
// 	producer := goka.NewProducer(bootstrapServers, topicName)

// 	b.Run(testName(prefix, "Goka@Produce"), func(b *testing.B) {
// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			producer.ProduceSync(generateKey(prefix, i), valueGenerator(i))
// 		}

// 		b.StopTimer()
// 	})

// 	b.Run(testName(prefix, "Goka@ProduceChannel"), func(b *testing.B) {
// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			producer.ProduceAsync(generateKey(prefix, i), valueGenerator(i))
// 		}

// 		producer.Wait()
// 		b.StopTimer()
// 	})
// }

// func benchmarkSaramaProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
// 	producer := sarama.NewProducer(bootstrapServers)

// 	b.Run(testName(prefix, "Sarama@Produce"), func(b *testing.B) {
// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			producer.ProduceSync(topicName, generateKey(prefix, i), valueGenerator(i))
// 		}
// 		b.StopTimer()
// 	})

// 	b.Run(testName(prefix, "Sarama@ProduceChannel"), func(b *testing.B) {
// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			producer.ProduceAsync(topicName, generateKey(prefix, i), valueGenerator(i))
// 		}
// 		producer.Wait()
// 		b.StopTimer()
// 	})
// }

// func benchmarkSegmentioProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
// 	producer := segmentio.NewProducer(bootstrapServers, topicName)

// 	b.Run(testName(prefix, "Segmentio@Produce"), func(b *testing.B) {
// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			producer.ProduceSync(generateKey(prefix, i), valueGenerator(i))
// 		}

// 		b.StopTimer()
// 	})

// 	b.Run(testName(prefix, "Segmentio@ProduceChannel"), func(b *testing.B) {
// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			producer.ProduceAsync(generateKey(prefix, i), valueGenerator(i))
// 		}

// 		producer.Wait()
// 		b.StopTimer()
// 	})
// }
