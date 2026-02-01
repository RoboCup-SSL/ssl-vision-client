FROM golang:1.24-alpine@sha256:12c199a889439928e36df7b4c5031c18bfdad0d33cdeae5dd35b2de369b5fbf5 AS build_go
ARG cmd=ssl-vision-cli
WORKDIR work
COPY . .
RUN go install ./cmd/${cmd}

# Start fresh from a smaller image
FROM alpine:3@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659
ARG cmd=ssl-vision-cli
COPY --from=build_go /go/bin/${cmd} /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
