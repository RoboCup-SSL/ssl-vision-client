FROM node:18-alpine AS build_node
COPY frontend /tmp/ssl-vision-client/frontend
WORKDIR /tmp/ssl-vision-client/frontend
RUN npm install
RUN npm run build

FROM golang:1.20-alpine AS build_go
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-vision-client
COPY cmd cmd
COPY pkg pkg
COPY frontend/embed.go frontend/embed.go
COPY go.mod .
COPY go.sum .
COPY --from=build_node /tmp/ssl-vision-client/frontend/dist frontend/dist
RUN go install -v ./cmd/ssl-vision-client

# Start fresh from a smaller image
FROM alpine:3
COPY --from=build_go /go/bin/ssl-vision-client /app/ssl-vision-client
USER 1000
EXPOSE 8082
ENTRYPOINT ["/app/ssl-vision-client"]
CMD []
