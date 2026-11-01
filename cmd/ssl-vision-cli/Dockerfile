FROM golang:1.27-alpine@sha256:d200d2dd76b03a99b52651fd1325aff7dcd35422c35c01bac56571b0507f3deb AS build_go
WORKDIR /work
COPY . .
RUN go install ./cmd/ssl-vision-cli

# Start fresh from a smaller image
FROM alpine:3@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a
COPY --from=build_go /go/bin/ssl-vision-cli /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
