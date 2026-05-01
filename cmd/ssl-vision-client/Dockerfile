FROM node:24-alpine@sha256:01743339035a5c3c11a373cd7c83aeab6ed1457b55da6a69e014a95ac4e4700b AS build_node
COPY frontend /frontend
WORKDIR /frontend
RUN corepack enable && pnpm install --frozen-lockfile
RUN corepack enable && pnpm run build

FROM golang:1.26-alpine@sha256:2389ebfa5b7f43eeafbd6be0c3700cc46690ef842ad962f6c5bd6be49ed82039 AS build_go
WORKDIR /work
COPY . .
COPY --from=build_node /frontend/dist frontend/dist
RUN go install ./cmd/ssl-vision-client

# Start fresh from a smaller image
FROM alpine:3@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11
COPY --from=build_go /go/bin/ssl-vision-client /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
