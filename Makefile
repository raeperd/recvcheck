all: build test lint

download:
	go mod download

build: download
	go build -C cmd/recvcheck

test:
	go test -race -coverprofile=coverage.out ./...

lint:
	golangci-lint run

