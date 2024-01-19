package kafkaclient

import (
	"testing"

	"github.com/lkumarjain/benchmark/kafka-client/confluent"
	"github.com/lkumarjain/benchmark/kafka-client/franz"
	"github.com/lkumarjain/benchmark/kafka-client/goka"
	"github.com/lkumarjain/benchmark/kafka-client/sarama"
	"github.com/lkumarjain/benchmark/kafka-client/segmentio"
)

func BenchmarkProducer(b *testing.B) {
	for _, tt := range tests {
		benchmarkConfluentProducer(b, tt.name, tt.valueGenerator)
		benchmarkFranzProducer(b, tt.name, tt.valueGenerator)
		benchmarkGokaProducer(b, tt.name, tt.valueGenerator)
		benchmarkSaramaProducer(b, tt.name, tt.valueGenerator)
		benchmarkSegmentioProducer(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkConfluentProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	producer := confluent.NewProducer(bootstrapServers)

	b.Run(testName(prefix, "Confluent@Produce"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(topicName, generateKey(prefix, i), valueGenerator(i))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Confluent@ProduceChannel"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(topicName, generateKey(prefix, i), valueGenerator(i))
		}
		producer.Wait()
		b.StopTimer()
	})
}

func benchmarkFranzProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	producer := franz.NewProducer(bootstrapServers)

	b.Run(testName(prefix, "Franz@Produce"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(topicName, generateKey(prefix, i), valueGenerator(i))
		}

		b.StopTimer()
	})

	b.Run(testName(prefix, "Franz@ProduceChannel"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(topicName, generateKey(prefix, i), valueGenerator(i))
		}

		producer.Wait()
		b.StopTimer()
	})
}

func benchmarkGokaProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	producer := goka.NewProducer(bootstrapServers, topicName)

	b.Run(testName(prefix, "Goka@Produce"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(generateKey(prefix, i), valueGenerator(i))
		}

		b.StopTimer()
	})

	b.Run(testName(prefix, "Goka@ProduceChannel"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(generateKey(prefix, i), valueGenerator(i))
		}

		producer.Wait()
		b.StopTimer()
	})
}

func benchmarkSaramaProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	producer := sarama.NewProducer(bootstrapServers)

	b.Run(testName(prefix, "Sarama@Produce"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(topicName, generateKey(prefix, i), valueGenerator(i))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Sarama@ProduceChannel"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(topicName, generateKey(prefix, i), valueGenerator(i))
		}
		producer.Wait()
		b.StopTimer()
	})
}

func benchmarkSegmentioProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	producer := segmentio.NewProducer(bootstrapServers, topicName)

	b.Run(testName(prefix, "Segmentio@Produce"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(generateKey(prefix, i), valueGenerator(i))
		}

		b.StopTimer()
	})

	b.Run(testName(prefix, "Segmentio@ProduceChannel"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(generateKey(prefix, i), valueGenerator(i))
		}

		producer.Wait()
		b.StopTimer()
	})
}
