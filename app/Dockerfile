FROM golang:1.16.4-alpine3.13

# アップデートとgitのインストール！！
RUN apk update && apk add git

ENV GO111MODULE=on

ENV $APP_PATH = /app
WORKDIR $APP_PATH

RUN go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

EXPOSE 8080

# RUN go run /go/src/server.go
# RUN go mod init main
# RUN go mod tidy

CMD ["go", "run", "server.go"]
# CMD ["./app"]
