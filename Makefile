all: build test lint

download:
	go mod download

build: download
	go build

test:
	go test -race ./...

lint:
	golangci-lint run

