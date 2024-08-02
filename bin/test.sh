#!/bin/bash -e

cd "$(dirname "$0")"
. ./utils.sh

function finish {
  echo 'Removing environment'
  echo '-----'
  docker compose down -v
}
trap finish EXIT

mkdir -p output

function main() {
  retrieve_cyberark_ca_cert
  docker compose build --pull tester

  docker compose run --rm \
    tester
}

main
