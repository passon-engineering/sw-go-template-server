#!/bin/bash

PACKAGE_NAME="sw-go-template-server"
echo Build file: "$BASH_SOURCE"
echo Build DIR: "$(dirname "$BASH_SOURCE")"

SCRIPT_DIR=$(dirname "$BASH_SOURCE")

GOOS=linux GOARCH=amd64 go build -o $SCRIPT_DIR/$PACKAGE_NAME-amd64-linux $SCRIPT_DIR/main.go
#GOOS=linux GOARCH=arm64 go build -o $SCRIPT_DIR/$PACKAGE_NAME-arm64-linux $SCRIPT_DIR/main.go
#GOOS=darwin GOARCH=arm64 go build -o $SCRIPT_DIR/$PACKAGE_NAME-arm64-darwin $SCRIPT_DIR/main.go