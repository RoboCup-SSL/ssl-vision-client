FROM node:22-alpine@sha256:f2dc6eea95f787e25f173ba9904c9d0647ab2506178c7b5b7c5a3d02bc4af145 AS build_node
COPY frontend frontend
WORKDIR frontend
RUN npm install
RUN npm run build

FROM golang:1.21-alpine AS build_go
ARG cmd=ssl-vision-client
WORKDIR work
COPY . .
COPY --from=build_node frontend/dist frontend/dist
RUN go install ./cmd/${cmd}

# Start fresh from a smaller image
FROM alpine:3
ARG cmd=ssl-vision-client
COPY --from=build_go /go/bin/${cmd} /app
USER 1000
ENTRYPOINT ["/app"]
CMD []
