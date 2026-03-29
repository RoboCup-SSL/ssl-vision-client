FROM golang:1.24-alpine@sha256:8bee1901f1e530bfb4a7850aa7a479d17ae3a18beb6e09064ed54cfd245b7191 AS build_go
WORKDIR /work
COPY . .
RUN go install ./cmd/ssl-vision-cli

# Start fresh from a smaller image
FROM alpine:3@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659
COPY --from=build_go /go/bin/ssl-vision-cli /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
