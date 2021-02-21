#!/bin/bash

SCRIPTS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
PROJECT_DIR="$SCRIPTS_DIR/.."
BUILD_DIR="$PROJECT_DIR/build"
cd "$PROJECT_DIR/cmd/rtmor"
env GOOS=linux GOARCH=386 go build -o "$PROJECT_DIR/build/linux-386/rtmor"
env GOOS=linux GOARCH=amd64 go build -o "$PROJECT_DIR/build/linux-amd64/rtmor"
env GOOS=windows GOARCH=386 go build -o "$PROJECT_DIR/build/windows-386/rtmor.exe"
env GOOS=windows GOARCH=amd64 go build -o "$PROJECT_DIR/build/windows-amd64/rtmor.exe"
env GOOS=darwin GOARCH=amd64 go build -o "$PROJECT_DIR/build/darwin-amd64/rtmor"
cd $PROJECT_DIR