#!/bin/sh

set -e

npm install
npm run build

go get -v -d ./...
cd cmd/ssl-vision-client
go get -v github.com/gobuffalo/packr/packr
packr install
