FROM golang:1.16.4-alpine3.13

# アップデートとgitのインストール！！
RUN apk update && apk add git

WORKDIR /go/src

# RUN go run /go/src/server.go

CMD ["go", "run", "server.go"]