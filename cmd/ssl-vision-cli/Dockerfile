FROM golang:1.26-alpine@sha256:3ad57304ad93bbec8548a0437ad9e06a455660655d9af011d58b993f6f615648 AS build_go
WORKDIR /work
COPY . .
RUN go install ./cmd/ssl-vision-cli

# Start fresh from a smaller image
FROM alpine:3@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b
COPY --from=build_go /go/bin/ssl-vision-cli /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
