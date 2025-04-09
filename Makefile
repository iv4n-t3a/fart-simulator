SOURCES := $(shell find . -type f -name *.go)
TESTS := $(shell find test -type f -name *_test.go)
TARGET := simulator

build: $(TARGET)

test:
	for i in $(TESTS); do \
		go test $$i ; \
	done

$(TARGET): $(SOURCES)
	go mod tidy
	go build -o $(TARGET) cmd/fart-simulator/main.go

.PHONY: build, test
