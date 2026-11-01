FROM node:24-alpine@sha256:8a32c8cf649a49915a66e81e7a609223c708d0adabf288ee3f4077d765542f85 AS build_node
COPY frontend /frontend
WORKDIR /frontend
RUN corepack enable && pnpm install --frozen-lockfile
RUN corepack enable && pnpm run build

FROM golang:1.27-alpine@sha256:d200d2dd76b03a99b52651fd1325aff7dcd35422c35c01bac56571b0507f3deb AS build_go
WORKDIR /work
COPY . .
COPY --from=build_node /frontend/dist frontend/dist
RUN go install ./cmd/ssl-vision-client

# Start fresh from a smaller image
FROM alpine:3@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a
COPY --from=build_go /go/bin/ssl-vision-client /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
