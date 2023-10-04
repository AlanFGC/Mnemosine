# Define the protobuf source directory and the Go output directory
PROTO_SRC_DIR = .
GO_OUT_DIR = ./generated

# List of .proto files
PROTO_FILES = $(wildcard $(PROTO_SRC_DIR)/*.proto)

# Generate .pb.go files
generate_pb:
	protoc --go_out=$(GO_OUT_DIR) --proto_path=$(PROTO_SRC_DIR) $(PROTO_FILES)

# Clean generated files
clean:
	rm -rf $(GO_OUT_DIR)/*.go

# Default target
all: generate_pb

.PHONY: all generate_pb clean