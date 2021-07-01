#!/user/bin/bash

cd /go/src/Junking/views
npm install
npm run build
cd /go/src/Junking
go mod tidy
go run server.go