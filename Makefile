SOURCES := $(shell find . -type f -name *.go)
TARGET := fart-simulator

build: $(TARGET)

$(TARGET): $(SOURCES)
	go mod tidy
	go build -o $(TARGET) cmd/fart-simulator/main.go

.PHONY: build
