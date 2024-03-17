#!/bin/bash

echo "Generating code from .proto files..."

protoc \
    --proto_path=./proto \
    --go_opt=paths=source_relative \
    --go_out=./gen \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_out=./gen \
    --go-grpc_opt=require_unimplemented_servers=false \
    ./proto/*.proto

echo "Finished"