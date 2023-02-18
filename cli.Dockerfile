FROM golang:1.20-alpine AS build
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-vision-client
COPY cmd cmd
COPY pkg pkg
COPY go.mod go.mod
COPY go.sum go.sum
RUN go install -v ./cmd/ssl-vision-cli

# Start fresh from a smaller image
FROM alpine:3
COPY --from=build /go/bin/ssl-vision-cli /app/ssl-vision-cli
ENTRYPOINT ["/app/ssl-vision-cli"]
CMD []
