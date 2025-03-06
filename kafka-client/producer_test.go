package kafkaclient

import (
	"testing"

	"github.com/lkumarjain/benchmark/kafkaclient/confluent"
	"github.com/lkumarjain/benchmark/kafkaclient/franz"
	"github.com/lkumarjain/benchmark/kafkaclient/goka"
	"github.com/lkumarjain/benchmark/kafkaclient/sarama"
	"github.com/lkumarjain/benchmark/kafkaclient/segmentio"
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
	producer := confluent.NewProducer(bootstrapServers, authenticator, userName, password)
	topicName := topicName(prefix)

	b.Run(testName(prefix, "Confluent@ProduceSync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(topicName, generateKey(prefix, i), valueGenerator(i))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Confluent@ProduceAsync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(topicName, generateKey(prefix, i), valueGenerator(i))
		}
		producer.Wait()
		b.StopTimer()
	})
}

func benchmarkFranzProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	producer := franz.NewProducer(bootstrapServers, authenticator, userName, password)
	topicName := topicName(prefix)

	b.Run(testName(prefix, "Franz@ProduceSync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(topicName, generateKey(prefix, i), valueGenerator(i))
		}

		b.StopTimer()
	})

	b.Run(testName(prefix, "Franz@ProduceAsync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(topicName, generateKey(prefix, i), valueGenerator(i))
		}

		producer.Wait()
		b.StopTimer()
	})
}

func benchmarkGokaProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	topicName := topicName(prefix)
	producer := goka.NewProducer(bootstrapServers, topicName, authenticator, userName, password)

	b.Run(testName(prefix, "Goka@ProduceSync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(generateKey(prefix, i), valueGenerator(i))
		}

		b.StopTimer()
	})

	b.Run(testName(prefix, "Goka@ProduceAsync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(generateKey(prefix, i), valueGenerator(i))
		}

		producer.Wait()
		b.StopTimer()
	})
}

func benchmarkSaramaProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	producer := sarama.NewProducer(bootstrapServers, authenticator, userName, password)
	topicName := topicName(prefix)

	b.Run(testName(prefix, "Sarama@ProduceSync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(topicName, generateKey(prefix, i), valueGenerator(i))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Sarama@ProduceAsync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(topicName, generateKey(prefix, i), valueGenerator(i))
		}
		producer.Wait()
		b.StopTimer()
	})
}

func benchmarkSegmentioProducer(b *testing.B, prefix string, valueGenerator func(int) string) {
	topicName := topicName(prefix)

	producer := segmentio.NewProducer(bootstrapServers, topicName, authenticator, userName, password)

	b.Run(testName(prefix, "Segmentio@ProduceSync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceSync(generateKey(prefix, i), valueGenerator(i))
		}

		b.StopTimer()
	})

	b.Run(testName(prefix, "Segmentio@ProduceAsync"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			producer.ProduceAsync(generateKey(prefix, i), valueGenerator(i))
		}

		producer.Wait()
		b.StopTimer()
	})
}
