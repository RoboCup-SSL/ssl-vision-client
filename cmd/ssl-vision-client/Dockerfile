FROM node:24-alpine@sha256:2bdb65ed1dab192432bc31c95f94155ca5ad7fc1392fb7eb7526ab682fa5bf14 AS build_node
COPY frontend /frontend
WORKDIR /frontend
RUN corepack enable && pnpm install --frozen-lockfile
RUN corepack enable && pnpm run build

FROM golang:1.26-alpine@sha256:91eda9776261207ea25fd06b5b7fed8d397dd2c0a283e77f2ab6e91bfa71079d AS build_go
WORKDIR /work
COPY . .
COPY --from=build_node /frontend/dist frontend/dist
RUN go install ./cmd/ssl-vision-client

# Start fresh from a smaller image
FROM alpine:3@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b
COPY --from=build_go /go/bin/ssl-vision-client /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
