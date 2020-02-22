#!/bin/bash

# Fail on errors
set -e

# Update to latest protobuf compiler
go get -u github.com/golang/protobuf/protoc-gen-go

for pkgDir in ./proto/*; do
  # Set package name to current directory
  packageName=${pkgDir##*/}

  # compile profobuf files in current directory
  protoc -I"$pkgDir" \
    -I"${GOPATH}/src" \
    --go_out=import_path="${packageName}:pkg/$packageName" \
    "${pkgDir}"/*.proto
done
