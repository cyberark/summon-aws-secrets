FROM golang:1.18
MAINTAINER Conjur Inc

# On CyberArk dev laptops, golang module dependencies are downloaded with a
# corporate proxy in the middle. For these connections to succeed we need to
# configure the proxy CA certificate in build containers.
#
# To allow this script to also work on non-CyberArk laptops where the CA
# certificate is not available, we copy the (potentially empty) directory
# and update container certificates based on that, rather than rely on the
# CA file itself.
ADD build_ca_certificate /usr/local/share/ca-certificates/
RUN update-ca-certificates

ENV GOOS=linux
ENV GOARCH=amd64

EXPOSE 8080

RUN apt-get update && \
    apt-get install -y jq

WORKDIR /summon-aws-secrets

RUN go install github.com/jstemmer/go-junit-report@latest && \
    go install github.com/axw/gocov/gocov@latest && \
    go install github.com/AlekSi/gocov-xml@latest && \
    mkdir -p /summon-aws-secrets/output

COPY go.mod go.sum /summon-aws-secrets/
RUN go mod download

COPY . .
