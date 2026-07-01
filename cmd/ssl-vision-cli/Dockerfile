FROM golang:1.26-alpine@sha256:91eda9776261207ea25fd06b5b7fed8d397dd2c0a283e77f2ab6e91bfa71079d AS build_go
WORKDIR /work
COPY . .
RUN go install ./cmd/ssl-vision-cli

# Start fresh from a smaller image
FROM alpine:3@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b
COPY --from=build_go /go/bin/ssl-vision-cli /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
