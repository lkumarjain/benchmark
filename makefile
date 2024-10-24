dependencies:
	go mod tidy
	go mod vendor

benchmark-expression-evaluation:
	cd  expression-evaluation && go test -bench=. -benchmem -count 5 -benchtime=10000x > results/results.out

benchmark-in-memory-cache:
	cd  in-memory-cache && go test -bench=. -benchmem -count 5 -benchtime=100000x > results/results.out

benchmark-kafka-producer:
	cd  kafka-client && go test -timeout=5h -bench=Producer -benchmem -count 5 -benchtime=10000x > results/producer.out

benchmark-kafka-consumer:
	cd  kafka-client && go test -timeout=5h -bench=Consumer -benchmem -count 5 -benchtime=10000x > results/consumer.out

benchmark-worker-pool:
	cd  worker-pool && go test -timeout=5h -bench=. -benchmem -count 5 -benchtime=1000000x > results/results.out