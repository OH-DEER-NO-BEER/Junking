FROM golang:1.16.4-alpine3.13

RUN apk update && apk add git

WORKDIR /go/src/Junking
ENV GO111MODULE=on
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
EXPOSE 8080

# RUN go run /go/src/server.go
# RUN go mod init main
# RUN go mod tidy

# CMD ["go", "run", "server.go"]

# CMD ["./app"]
