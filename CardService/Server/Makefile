# Define the protobuf source directory and the Go output directory
PROTO_SRC_DIR = .
GO_OUT_DIR = .

# List of .proto files
PROTO_FILES = $(wildcard $(PROTO_SRC_DIR)/*.proto)

# Generate .pb.go and grpc out files
generate_pb:
	protoc --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
        $(PROTO_FILES)

# Clean generated files
clean:
	rm -rf $(GO_OUT_DIR)/*.pb.go

# Default target
all: generate_pb

.PHONY: all generate_pb clean