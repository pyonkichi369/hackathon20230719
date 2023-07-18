FROM golang:1.19.1-alpine

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/chatgpt

WORKDIR /go/src/github.com/chatgpt

ADD . /go/src/github.com/chatgpt

RUN go install golang.org/x/tools/cmd/goimports@latest
