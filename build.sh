#!/bin/bash

cd "$(dirname "$(greadlink -f "$0")")/build" || exit 1
env GOARCH=amd64 GOOS=linux go build ..
cp ./xfer-bin "$HOME/.porter/mixins/xfer"
cp ./xfer-bin "$HOME/xfer/"

cd "$(dirname "$(greadlink -f "$0")")/macos" || exit 1
go build ../../