#!/bin/bash -ex

cd "$(dirname "$0")"

docker run --rm \
  --volume "${PWD}/CHANGELOG.md:/CHANGELOG.md"  \
  cyberark/parse-a-changelog
