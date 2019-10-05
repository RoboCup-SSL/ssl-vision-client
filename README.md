[![CircleCI](https://circleci.com/gh/RoboCup-SSL/ssl-vision-client/tree/master.svg?style=svg)](https://circleci.com/gh/RoboCup-SSL/ssl-vision-client/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/RoboCup-SSL/ssl-vision-client?style=flat-square)](https://goreportcard.com/report/github.com/RoboCup-SSL/ssl-vision-client)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/RoboCup-SSL/ssl-vision-client/pkg/vision)
[![Release](https://img.shields.io/github/release/RoboCup-SSL/ssl-vision-client.svg?style=flat-square)](https://github.com/RoboCup-SSL/ssl-vision-client/releases/latest)

# ssl-vision-client

A graphical client for [ssl-vision](https://github.com/RoboCup-SSL/ssl-vision) that receives multicast packages and
shows them in a web-ui.

## Usage
If you just want to use this app, simply download the latest [release binary](https://github.com/RoboCup-SSL/ssl-vision-client/releases/latest). The binary is self-contained. No dependencies are required.

## Development

### Requirements
You need to install following dependencies first: 
 * Go >= 1.9
 * Node
 * Yarn

### Prepare
Download and install to [GOPATH](https://github.com/golang/go/wiki/GOPATH):
```bash
go get -u github.com/RoboCup-SSL/ssl-vision-client/...
```
Switch to project root directory
```bash
cd $GOPATH/src/github.com/RoboCup-SSL/ssl-vision-client/
```
Download dependencies for frontend
```bash
yarn install
```

### Run
Run the backend:
```bash
go run cmd/ssl-vision-client/main.go
```

Run the UI:
```bash
# compile and hot-reload
yarn serve
```
Or use the provided IntelliJ run configurations.

### Build self-contained release binary
First, build the UI resources
```bash
# compile and minify UI
yarn build
```
Then build the backend with `packr`
```bash
# get packr
go get github.com/gobuffalo/packr/packr
# install the binary
cd cmd/ssl-vision-client
packr install
```
