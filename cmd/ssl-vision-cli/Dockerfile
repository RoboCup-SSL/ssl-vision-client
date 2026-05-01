FROM golang:1.26-alpine@sha256:2389ebfa5b7f43eeafbd6be0c3700cc46690ef842ad962f6c5bd6be49ed82039 AS build_go
WORKDIR /work
COPY . .
RUN go install ./cmd/ssl-vision-cli

# Start fresh from a smaller image
FROM alpine:3@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11
COPY --from=build_go /go/bin/ssl-vision-cli /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
