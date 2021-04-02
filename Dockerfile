FROM golang:latest

ADD . /go/src/github.com/takahiro0530/picture_management_server

RUN go install github.com/takahiro0530/picture_management_server

EXPOSE 50051