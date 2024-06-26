all: run

deps:
	go mod tidy

build: deps
	go build cmd/main.go

run: build
	./main

