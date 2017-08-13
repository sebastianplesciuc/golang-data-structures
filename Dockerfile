FROM golang:latest

ADD . /opt/go/src/github.com/sebastianplesciuc/golang-data-structures
ENV GOPATH /opt/go

WORKDIR /opt/go/src/github.com/sebastianplesciuc/golang-data-structures

RUN bin/build.sh

ENTRYPOINT build/main
