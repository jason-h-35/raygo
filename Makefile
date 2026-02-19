.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

check: vet test

cover:
	go test -v -cover -coverprofile=c.out ./...
	go tool cover -html=c.out

build:
	go build raygo.go

b: build

run:
	go run raygo.go

r: run

t: test
