FROM golang:1.15.7-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git

EXPOSE 8080