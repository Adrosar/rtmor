#!/bin/bash

SCRIPTS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
PROJECT_DIR="$SCRIPTS_DIR/.."
BUILD_DIR="$PROJECT_DIR/build"
cd "$PROJECT_DIR/cmd/rtmor"
go build -o "$BUILD_DIR"
cd $PROJECT_DIR