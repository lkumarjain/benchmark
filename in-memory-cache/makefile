dependencies:
	go mod tidy
	go mod vendor

benchmark:
	go test -bench=. -benchmem -count 5 -benchtime=100000x > results/results.out
