SOURCES := $(shell find . -type f -name *.go)
TARGET := simulator

build: $(TARGET)

test:
	go test -v ./test/...

$(TARGET): $(SOURCES)
	go mod tidy
	go build -o $(TARGET) cmd/fart-simulator/main.go

.PHONY: build, test
