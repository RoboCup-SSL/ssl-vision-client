FROM node:10.21.0-jessie AS build_node
WORKDIR /tmp/ssl-vision-client
COPY . .
RUN yarn install
RUN yarn build

FROM golang:1.14-alpine AS build_go
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-vision-client
COPY . .
COPY --from=build_node /tmp/ssl-vision-client/dist dist
RUN go get -v -t -d ./...
RUN go get -v github.com/gobuffalo/packr/packr
WORKDIR cmd/ssl-vision-client
RUN GOOS=linux GOARCH=amd64 packr build -o ../../release/ssl-vision-client_linux_amd64

# Start fresh from a smaller image
FROM alpine:3.9
COPY --from=build_go /go/src/github.com/RoboCup-SSL/ssl-vision-client/release/ssl-vision-client_linux_amd64 /app/ssl-vision-client
EXPOSE 8082
ENTRYPOINT ["/app/ssl-vision-client"]
CMD ["/app/ssl-vision-client", "-address", ":8082"]
