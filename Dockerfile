FROM golang:1.13-stretch
MAINTAINER Conjur Inc

ENV GOOS=linux
ENV GOARCH=amd64

EXPOSE 8080

RUN apt-get update && \
    apt-get install -y jq

WORKDIR /summon-aws-secrets

RUN go get -u github.com/jstemmer/go-junit-report && \
    go get github.com/smartystreets/goconvey && \
    go get -u github.com/axw/gocov/gocov && \
    go get -u github.com/AlekSi/gocov-xml && \
    mkdir -p /summon-aws-secrets/output

COPY go.mod go.sum /summon-aws-secrets/
RUN go mod download

COPY . .
