FROM golang:1.17-alpine AS build
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-vision-client
COPY cmd cmd
COPY pkg pkg
COPY go.mod go.mod
RUN go get -v -t -d ./...
RUN go install ./...

# Start fresh from a smaller image
FROM alpine:3.15
COPY --from=build /go/bin/ssl-vision-cli /app/ssl-vision-cli
ENTRYPOINT ["/app/ssl-vision-cli"]
CMD []
