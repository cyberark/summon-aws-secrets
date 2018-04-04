FROM golang:1.8
MAINTAINER Conjur Inc

RUN apt-get update && apt-get install jq
RUN go get -u github.com/jstemmer/go-junit-report
RUN go get github.com/tools/godep
RUN go get github.com/smartystreets/goconvey

RUN mkdir -p /go/src/github.com/cyberark/summon-aws-secrets/output
WORKDIR /go/src/github.com/cyberark/summon-aws-secrets

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64

EXPOSE 8080
