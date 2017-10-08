#!/usr/bin/env bash

# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -o berk

 docker run -it --rm  -v "$PWD/..":/usr/src/myapp -v "$GOPATH"/src:/go/src  -w /usr/src/myapp  -e GOOS=linux -e GOARCH=amd64 -e CGO_ENABLED=0  golang:alpine go build -a -v -installsuffix cgo -o berk