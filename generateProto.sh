#!/bin/bash

# Fail on errors
set -e

# Update to latest protobuf compiler
go get -u github.com/golang/protobuf/protoc-gen-go

for pkgDir in ./proto/*; do
  # compile protobuf files in current directory
  protoc -I"$pkgDir" \
    -I"${GOPATH}/src" \
    --go_out="$GOPATH/src" \
    "${pkgDir}"/*.proto
done
