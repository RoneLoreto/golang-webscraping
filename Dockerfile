FROM golang:1.15.6-alpine3.12
RUN apk add --update --no-cache \
  alpine-sdk \
  git && mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
ENV GO111MODULE=on
