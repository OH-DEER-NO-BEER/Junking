#!/user/bin/bash

cd /go/src/Junking/views
npm install
npm run build
cd /go/src/Junking
go run server.go