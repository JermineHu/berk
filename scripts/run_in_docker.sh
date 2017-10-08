#!/usr/bin/env bash

 docker run -it --rm  -v "../":/usr/src/myapp -v "$GOPATH"/src:/go/src  -w /usr/src/myapp  -e GOOS=linux -e GOARCH=amd64 -e CGO_ENABLED=0  golang:1.8.0-alpine go build -a -v -installsuffix cgo -o berk

 chmod +x ./berk

 mv ./berk ~/berk/xx

docker run --rm --name berk -v ~/berk/:/data  -p 1111:8000 -e "ResourcesPath=/data/berk/" -e "MONGODB_CONNECTION=Jermine:123456@mongoDB:27017/berk"  --link mongo:mongoDB alpine /data/xx