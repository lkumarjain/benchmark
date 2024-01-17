# Kafka Client Benchmarks

Benchmarks of kafka client libraries for Golang.

## Execute Benchmark

```bash
 go test -bench=. -benchmem -count 5 -benchtime=100000x > results/results.out
```

## Docker File

```bash
docker-compose -p kafka -f docker-compose.yml up
```

## Results

All the [benchmarks](/results.out) are performed in the `Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz` machine with `100K` samples and `5` iterations.

![Average](/in-memory-cache/results/Average_Cache.png)

### Average ns / operation

#### Set Function

![Average_ns_per_operation_set.png](/in-memory-cache/results/Average_ns_per_operation_set.png)

#### Get & Remove Function

![Average_ns_per_operation_get_remove](/in-memory-cache/results/Average_ns_per_operation_get_remove.png)

## Libraries

:warning: Please note that these libraries are benchmarked against storage of sample payloads (i.e. 1, 5, and 10 KB). You are encouraged to benchmark with your custom payloads.

- [confluentinc/confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) - Confluent's Apache Kafka Golang client.
- [twmb/franz-go](https://github.com/twmb/franz-go) - A feature complete, pure Go library for interacting with Kafka.
- [lovoo/goka](https://github.com/lovoo/goka) - Goka is a compact yet powerful distributed stream processing library for Apache Kafka written in Go.
- [IBM/sarama](https://github.com/IBM/sarama) - Sarama is a Go library for Apache Kafka.
- [segmentio/kafka-go](https://github.com/segmentio/kafka-go) - Kafka library in Go.
  
## Credits

- Test data is generated using [mockaroo](https://www.mockaroo.com/)
