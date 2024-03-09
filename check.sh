#!/bin/bash
BASE_PATH=$(pwd)
MODULE_PATHS=(
    "$BASE_PATH/cmds/rest"
    "$BASE_PATH/cmds/grpc"
    "$BASE_PATH/internal"
    "$BASE_PATH/pkg"
)

for MODULE_PATH in "${MODULE_PATHS[@]}";
do
    echo -e "Module $MODULE_PATH:" && cd "$MODULE_PATH" && \
    echo -e "  |  Tidying..." && go mod tidy | sed 's/^/     |  /' && \
    echo -e "  |  Formatting..." && go fmt ./... | sed 's/^/     |  /' && \
    echo -e "  |  Checking..." && go vet ./...   | sed 's/^/     |  /' && \
    echo -e "  |  Testing..." && go test ./...   | sed 's/^/     |  /' | sed 's/\t//' | sed 's/?//' && \
    echo -e "  |  Finished.\n" && cd "$BASE_PATH";
done
