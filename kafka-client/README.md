# Kafka Client Benchmarks

Benchmarks of kafka client libraries for Golang.

## Execute Benchmark

```bash
go test -timeout=5h -bench=Producer -benchmem -count 5 -benchtime=10000x > results/producer_results.out
go test -timeout=5h -bench=Consumer -benchmem -count 5 -benchtime=100000x > results/consumer_results.out
```

## Docker File

```bash
docker-compose -p kafka -f docker-compose.yml up
```

## Results

All the [benchmarks](/kafka-client/results) are performed in the `Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz` machine with a standalone Kafka running in the docker.

### Producer

All the [benchmarks](/kafka-client/results/producer_results.out) are performed with `10K` samples size and `5` iterations.

#### Microsecond / Operation

| Sync Producer                                                       | Async Producer                                                        |
| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
| ![SyncProducerTime.png](/kafka-client/results/SyncProducerTime.png) | ![AsyncProducerTime.png](/kafka-client/results/AsyncProducerTime.png) |

#### Memory Allocation / Operation

| Sync Producer                                                                    | Async Producer                                                                     |
| -------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| ![SyncProducerTime.png](/kafka-client/results/SyncProducerMemoryAllocations.png) | ![AsyncProducerTime.png](/kafka-client/results/AsyncProducerMemoryAllocations.png) |

#### Bytes / Operation

| Sync Producer                                                        | Async Producer                                                         |
| -------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| ![SyncProducerTime.png](/kafka-client/results/SyncProducerBytes.png) | ![AsyncProducerTime.png](/kafka-client/results/AsyncProducerBytes.png) |

## Libraries

:warning: Please note that these libraries are benchmarked against storage of sample payloads (i.e. 1, 5, and 10 KB). You are encouraged to benchmark with your custom payloads.

- [confluentinc/confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) - Confluent's Apache Kafka Golang client.
- [twmb/franz-go](https://github.com/twmb/franz-go) - A feature complete, pure Go library for interacting with Kafka.
- [lovoo/goka](https://github.com/lovoo/goka) - Goka is a compact yet powerful distributed stream processing library for Apache Kafka written in Go.
- [IBM/sarama](https://github.com/IBM/sarama) - Sarama is a Go library for Apache Kafka.
- [segmentio/kafka-go](https://github.com/segmentio/kafka-go) - Kafka library in Go.
  
## Credits

- Test data is generated using [mockaroo](https://www.mockaroo.com/)
