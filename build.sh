#!/bin/bash

APP='summon-aws-secrets'
PKGDIR='output'
OSES=(
  'darwin'
  'freebsd'
  'linux'
  'netbsd'
  'openbsd'
  'solaris'
  'windows'
)
GOARCH='amd64'

for GOOS in "${OSES[@]}"; do
  echo "Building $GOOS-$GOARCH binary"
  echo "-----"

  docker run --rm \
    -v "$PWD:/go/src/$APP" -w "/go/src/$APP" \
    -e "GOOS=$GOOS" -e "GOARCH=$GOARCH" \
    golang:1.9 \
    go build -v -o $PKGDIR/$APP-$GOOS-$GOARCH
done

echo "Building linux-alpine binary"
echo "-----"

docker run --rm \
  -v "$PWD:/go/src/$APP" -w "/go/src/$APP" \
  -e "GOOS=linux" -e "GOARCH=$GOARCH" \
  golang:1.9-alpine \
  go build -v -o $PKGDIR/$APP-linux-alpine-amd64
