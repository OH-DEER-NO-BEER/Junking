FROM golang:latest

WORKDIR /go/src/Junking/views 
RUN apt-get update && apt-get install -y git && apt-get install -y nodejs npm && npm install
# RUN apk update && apk add git && apk add nodejs && apk add --update npm && npm install && npm run build
WORKDIR /go/src/Junking

ENV GO111MODULE=on
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
EXPOSE 8080

# RUN go run /go/src/server.go
# RUN go mod init main
# RUN go mod tidy

CMD ["go", "run", "server.go"]

# CMD ["./app"]
