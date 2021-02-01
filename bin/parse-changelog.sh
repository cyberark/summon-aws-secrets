#!/bin/bash -ex

# Run from top-level dir
cd "$(dirname "$0")/.." || (echo "Could not cd to parent dir"; exit 1)

docker run --rm \
  --volume "${PWD}/CHANGELOG.md:/CHANGELOG.md"  \
  cyberark/parse-a-changelog
