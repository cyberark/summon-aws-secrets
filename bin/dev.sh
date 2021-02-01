#!/usr/bin/env bash

function finish {
  echo 'Removing environment'
  echo '-----'
  docker-compose down -v
}
trap finish EXIT

function main() {

  docker-compose build --pull tester

  docker-compose run --rm \
    --entrypoint "bash -c './bin/convey.sh& bash'" \
    --service-ports \
    tester
}

main
