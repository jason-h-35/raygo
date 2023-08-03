.DEFAULT_GOAL := run

fmt:
	go fmt ./...

lint: fmt
	go lint ./...

vet: fmt
	go vet ./...

build: vet
	go build raygo.go

run: vet
	go run raygo.go
