# Define the source directory where your .proto files are located
PROTO_DIR := ./

# Define the output directory for generated Go files
OUTPUT_DIR := ./

# Collect all .proto files in the source directory
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Generate Go files for each .proto file
GO_FILES := $(patsubst $(PROTO_DIR)/%.proto,$(OUTPUT_DIR)/%.pb.go,$(PROTO_FILES))

# Command to run the protoc compiler
PROTOC := protoc

all: $(GO_FILES)

$(OUTPUT_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	@mkdir -p $(@D)
	$(PROTOC) --go_out=$(OUTPUT_DIR) $<

clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: all clean