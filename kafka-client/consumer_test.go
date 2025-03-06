package kafkaclient

import (
	"testing"

	"github.com/lkumarjain/benchmark/kafkaclient/confluent"
	"github.com/lkumarjain/benchmark/kafkaclient/franz"
	"github.com/lkumarjain/benchmark/kafkaclient/goka"
	"github.com/lkumarjain/benchmark/kafkaclient/sarama"
	"github.com/lkumarjain/benchmark/kafkaclient/segmentio"
)

func BenchmarkConsumer(b *testing.B) {
	for _, tt := range tests {
		// benchmarkConfluentConsumer(b, tt.name)
		// benchmarkFranzConsumer(b, tt.name)
		// benchmarkGokaConsumer(b, tt.name)
		// benchmarkSaramaConsumer(b, tt.name)
		benchmarkSegmentioConsumer(b, tt.name)
	}
}

func benchmarkConfluentConsumer(b *testing.B, prefix string) {
	topicName := topicName(prefix)
	consumer := confluent.Consumer{Servers: bootstrapServers, EnableEvents: false, Topic: topicName,
		Authenticator: authenticator, UserName: userName, Password: password}

	b.Run(testName(prefix, "Confluent@ConsumePoll"), func(b *testing.B) {
		consumer.Start()
		<-consumer.Message // Added this to wait till Consumer gets ready

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}

		b.StopTimer()
		close(consumer.Done)
	})

	consumer.EnableEvents = true

	b.Run(testName(prefix, "Confluent@ConsumeEvent"), func(b *testing.B) {
		consumer.Start()
		<-consumer.Message // Added this to wait till Consumer gets ready

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}

		b.StopTimer()
		close(consumer.Done)
	})
}

func benchmarkFranzConsumer(b *testing.B, prefix string) {
	topicName := topicName(prefix)

	consumer := franz.Consumer{Servers: bootstrapServers, EnablePartition: false, Topic: topicName,
		Authenticator: authenticator, UserName: userName, Password: password}

	b.Run(testName(prefix, "Franz@ConsumeRecord"), func(b *testing.B) {
		consumer.Start()
		<-consumer.Message // Added this to wait till Consumer gets ready

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}

		b.StopTimer()
		close(consumer.Done)
	})

	consumer.EnablePartition = true

	b.Run(testName(prefix, "Franz@ConsumePartition"), func(b *testing.B) {
		consumer.Start()
		<-consumer.Message // Added this to wait till Consumer gets ready

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}

		b.StopTimer()
		close(consumer.Done)
	})
}

func benchmarkGokaConsumer(b *testing.B, prefix string) {
	topicName := topicName(prefix)

	consumer := goka.Consumer{Servers: bootstrapServers, Topic: topicName,
		Authenticator: authenticator, UserName: userName, Password: password}

	b.Run(testName(prefix, "Goka@Consumer"), func(b *testing.B) {
		consumer.Start()
		<-consumer.Message // Added this to wait till Consumer gets ready

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}

		b.StopTimer()
		close(consumer.Done)
	})
}

func benchmarkSaramaConsumer(b *testing.B, prefix string) {
	topicName := topicName(prefix)

	consumer := sarama.Consumer{Servers: bootstrapServers, Topic: topicName, EnablePartition: false,
		Authenticator: authenticator, UserName: userName, Password: password}

	b.Run(testName(prefix, "Sarama@ConsumerGroup"), func(b *testing.B) {
		consumer.Start()
		<-consumer.Message // Added this to wait till Consumer gets ready

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}

		b.StopTimer()
		close(consumer.Done)
	})

	consumer.EnablePartition = true

	b.Run(testName(prefix, "Sarama@ConsumePartition"), func(b *testing.B) {
		consumer.Start()
		<-consumer.Message // Added this to wait till Consumer gets ready

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}

		close(consumer.Done)
		b.StopTimer()
	})

}

func benchmarkSegmentioConsumer(b *testing.B, prefix string) {
	topicName := topicName(prefix)

	consumer := segmentio.Consumer{Servers: bootstrapServers, Topic: topicName, EnablePartition: false,
		Authenticator: authenticator, UserName: userName, Password: password}

	b.Run(testName(prefix, "Segmentio@ConsumerFetch"), func(b *testing.B) {
		consumer.Start()
		<-consumer.Message // Added this to wait till Consumer gets ready

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			<-consumer.Message
		}

		close(consumer.Done)
		b.StopTimer()
	})
}
