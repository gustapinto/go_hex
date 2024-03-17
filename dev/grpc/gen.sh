#!/bin/bash
BASE_PATH=$(pwd)
PROTO_PATH="$BASE_PATH/cmds/grpc/proto"
OUT_PATH="$BASE_PATH/cmds/grpc/gen"

echo "Generating code from .proto files..."

protoc \
    --proto_path="$PROTO_PATH" \
    --go_opt=paths=source_relative \
    --go_out="$OUT_PATH" \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_out="$OUT_PATH" \
    --go-grpc_opt=require_unimplemented_servers=false \
    $PROTO_PATH/*.proto

echo "Finished"