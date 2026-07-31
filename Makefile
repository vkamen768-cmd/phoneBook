.PHONY: run build lint

run:
	go run main.go

build:
	go build -o bin/appAccount main.go

lint:
	golangci-lint run