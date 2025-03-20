#!/bin/bash

# Check if directory parameter is provided
if [ $# -eq 0 ]; then
    PROTO_DIR="." # Default to current directory
else
    PROTO_DIR="$1" # Use the provided directory
fi

# Output file
OUTPUT_FILE="${PROTO_DIR}/gorpc.proto"
# Delete the output file if it already exists
if [ -f "$OUTPUT_FILE" ]; then
    rm "$OUTPUT_FILE"
    echo "Deleted existing $OUTPUT_FILE"
fi

# Create the header part in the output file
cat >"$OUTPUT_FILE" <<'EOL'
syntax = "proto3";

package gorpc;
option go_package = "./gorpc";

EOL

# First add common.proto content after the header
grep -v 'syntax\|package\|option go_package' common.proto >>"$OUTPUT_FILE"

# Then append all other proto files
for file in "$PROTO_DIR"/*.proto; do
    # Skip common.proto and the output file itself
    if [[ "$(basename "$file")" != "common.proto" && "$(basename "$file")" != "gorpc.proto" ]]; then
        echo "Processing $(basename "$file")..."
        # Skip the header part and common.proto import for other files
        grep -v 'syntax\|package\|option go_package\|import "pb/gorpc/common.proto"' "$file" >>"$OUTPUT_FILE"
    fi
done


echo "Proto files have been merged into $OUTPUT_FILE"

goctl rpc protoc gorpc.proto \
  --go_out=./pb \
  --go-grpc_out=./pb \
  --zrpc_out=./tmp_generate \
  -m \

if [ -d "./client" ]; then
  echo "Removing existing client directory..."
  rm -rf ./client
fi

cp -r ./tmp_generate/client ./client

rm -rf ./tmp_generate

go mod tidy