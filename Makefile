.DEFAULT_GOAL := build

fmt:
	go fmt ./...

lint: fmt
	go lint ./...

vet: fmt
	go vet ./...

build: vet
	go build raygo.go
