#!/bin/bash
set -euo pipefail

cd "$(dirname "$0")" || exit 1

MODULE="git.livesys.org/go-service/go-service-pb/gozero_gen"
PROTO_DIR="./proto"
OUTPUT_DIR="./gozero_gen"

# Validate proto directory exists
if [[ ! -d "${PROTO_DIR}" ]]; then
    echo "Error: Proto directory ${PROTO_DIR} not found" >&2
    exit 1
fi

rm -rf "${OUTPUT_DIR}"
mkdir -p "${OUTPUT_DIR}"  # 新增目录创建命令

find "${PROTO_DIR}" -name '*.proto' -print0 | xargs -0 protoc \
    --proto_path="${PROTO_DIR}" \
    --go_out="${OUTPUT_DIR}" \
    --go_opt=module="${MODULE}" \
    --go-grpc_out="${OUTPUT_DIR}" \
    --go-grpc_opt=module="${MODULE}"

echo "Protobuf code generation completed successfully."