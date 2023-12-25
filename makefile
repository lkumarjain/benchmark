dependencies:
	go mod tidy
	go mod vendor

benchmark-expression-evaluation:
	cd  expression-evaluation && go test -bench=. -benchmem -count 5 -benchtime=10000x > results/results.out

benchmark-in-memory-cache:
	cd  in-memory-cache && go test -bench=. -benchmem -count 5 -benchtime=100000x > results/results.out