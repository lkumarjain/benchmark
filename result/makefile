BINARY_NAME=result

dependencies:
	go mod tidy
	go mod vendor

build:
	go build -o bin/$(BINARY_NAME) cmd/main.go