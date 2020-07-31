#!/bin/bash -e

echo "Running tests"
echo "-----"

go clean -i
go install

go test --coverprofile=output/c.out -v ./... | tee output/junit.output

cat output/junit.output | go-junit-report > output/junit.xml

gocov convert output/c.out | gocov-xml > output/coverage.xml

rm output/junit.output
# Leave c.out in place as that file is what is passed to Code Climate
