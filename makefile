benchmark-expression-evaluation:
	cd result && make dependencies && make build
	cd expression-evaluation && make dependencies && make benchmark && make generate-graph

benchmark-in-memory-cache:
	cd result && make dependencies && make build
	cd in-memory-cache && make dependencies && make benchmark

benchmark-kafka-producer:
	cd result && make dependencies && make build
	cd  kafka-client && make dependencies && make benchmark-producer

benchmark-kafka-consumer:
	cd result && make dependencies && make build
	cd  kafka-client && make dependencies && make benchmark-consumer

benchmark-worker-pool:
	cd result && make dependencies && make build
	cd  worker-pool && make dependencies && make benchmark && make generate-graph