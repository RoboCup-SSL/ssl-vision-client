FROM node:24-alpine@sha256:2bdb65ed1dab192432bc31c95f94155ca5ad7fc1392fb7eb7526ab682fa5bf14 AS build_node
COPY frontend /frontend
WORKDIR /frontend
RUN corepack enable && pnpm install --frozen-lockfile
RUN corepack enable && pnpm run build

FROM golang:1.26-alpine@sha256:3ad57304ad93bbec8548a0437ad9e06a455660655d9af011d58b993f6f615648 AS build_go
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
