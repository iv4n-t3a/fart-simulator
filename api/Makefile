APIDIR := generated
PROTODIR := ipc
PROTOC := protoc
PYDIR := ../tools/visualisation/

PROTO := $(shell find $(PROTODIR) -type f -name '*.proto')
TARGET := $(patsubst $(PROTODIR)/%, $(APIDIR)/%, $(PROTO:.proto=_api))

all: $(TARGET)
	@echo $(TARGET)

$(APIDIR)/%_api: $(PROTODIR)/%.proto
	$(PROTOC) --proto_path=. --go_out=. --go-grpc_out=. $<
	python3 -m grpc_tools.protoc --proto_path=. --python_out=$(PYDIR) --grpc_python_out=$(PYDIR) $<

setup:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: all setup
