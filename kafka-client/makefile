dependencies:
	go mod tidy
	go mod vendor

benchmark-producer:
	go test -timeout=5h -bench=Producer -benchmem -count 5 -benchtime=10x > results/producer.out

benchmark-consumer:
	go test -timeout=5h -bench=Consumer -benchmem -count 5 -benchtime=10x > results/consumer.out