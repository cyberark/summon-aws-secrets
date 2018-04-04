#!/bin/bash

echo "Running tests"
echo "-----"

go clean -i
go install

go test -v "$(go list ./... | grep -v /vendor/)" | tee output/junit.output

cat output/junit.output | go-junit-report > output/junit.xml

rm output/junit.output
