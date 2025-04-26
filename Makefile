SOURCES := $(shell find . -type f -name *.go)
TARGET := simulator

build: $(TARGET)

run:
	go mod tidy
	go run cmd/fart-simulator/main.go

test:
	go mod tidy
	go test -v ./test/...

api:
	make -C api

$(TARGET): $(SOURCES)
	go mod tidy
	go build -o $(TARGET) cmd/fart-simulator/main.go

.PHONY: build, api, run, test
