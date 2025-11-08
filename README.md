[![CircleCI](https://circleci.com/gh/RoboCup-SSL/ssl-vision-client/tree/master.svg?style=svg)](https://circleci.com/gh/RoboCup-SSL/ssl-vision-client/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/RoboCup-SSL/ssl-vision-client?style=flat-square)](https://goreportcard.com/report/github.com/RoboCup-SSL/ssl-vision-client)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/RoboCup-SSL/ssl-vision-client/pkg/vision)
[![Release](https://img.shields.io/github/release/RoboCup-SSL/ssl-vision-client.svg?style=flat-square)](https://github.com/RoboCup-SSL/ssl-vision-client/releases/latest)

# ssl-vision-client

A graphical client for [ssl-vision](https://github.com/RoboCup-SSL/ssl-vision) that receives multicast packages and
shows them in a web-ui.

## Usage
If you just want to use this app, simply download the latest [release binary](https://github.com/RoboCup-SSL/ssl-vision-client/releases/latest).
The binary is self-contained. No dependencies are required.

You can also use pre-build docker images:
```shell script
docker pull robocupssl/ssl-vision-client
docker run -p 8082:8082 robocupssl/ssl-vision-client
```

By default, the UI is available at http://localhost:8082

## Development

### Requirements

You need to install following dependencies first:

* Go
* Node

See [.circleci/config.yml](.circleci/config.yml) for compatible versions.

### Frontend

See [frontend/README.md](frontend/README.md)

### Build

Build and install all binaries:

```bash
make install
```

### Run

Run the backend:

```bash
go run cmd/ssl-vision-client/main.go
```

### Update generated protobuf code

Generate the code for the `.proto` files after you've changed anything in a `.proto` file with:

```shell
make proto
```

## Releases
You can find all published releases on the Releases page
 of this repository. Each release contains pre-built binaries (for major platforms). To get up and running quickly, follow the steps below:

1. Navigate to Releases at https://github.com/RoboCup-SSL/ssl-vision-client/releases/.

2. Find the version you want and download the asset labelled something like ssl-vision-client_`<version>`_`<os-arch>`

3. On Linux/MacOS/Windows, open a terminal and run the binary 
You might need to mark it executable on Linux/Mac: 
```shell
chmod +x <file>
```
