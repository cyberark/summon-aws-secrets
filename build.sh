#!/bin/bash

# Platforms to build: https://golang.org/doc/install/source#environment
PLATFORMS=(
  'darwin:amd64'     # MacOS
  # 'dragonfly:amd64'  # Dragonfly https://www.dragonflybsd.org/
#  'freebsd:amd64'
  # 'linux:386'
  'linux:amd64'
  # 'linux:arm'
  # 'linux:arm64'
#  'netbsd:amd64'
#  'openbsd:amd64'
#  'solaris:amd64'
  # 'windows:386'
#  'windows:amd64'
)

echo "Creating summon-aws-secrets binaries in output/"
docker-compose build --pull builder

for platform in "${PLATFORMS[@]}"; do
  GOOS=${platform%%:*}
  GOARCH=${platform#*:}

  echo "-----"
  echo "GOOS=$GOOS, GOARCH=$GOARCH"
  echo "....."

  docker-compose run --rm \
    -e GOOS=$GOOS -e GOARCH=$GOARCH \
    builder \
      build -v -o output/summon-aws-secrets-$GOOS-$GOARCH
done
