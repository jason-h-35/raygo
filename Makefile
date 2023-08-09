.DEFAULT_GOAL := run

fmt:
	go fmt ./...

lint: fmt
	go lint ./...

vet: fmt
	go vet ./...

cover: fmt
	go test -v -cover -coverprofile=c.out ./...
	go tool cover -html=c.out

test: vet
	go test ./...

build: vet
	go build raygo.go

run: vet
	go run raygo.go
