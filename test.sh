#!/bin/bash -e

function finish {
  echo 'Removing environment'
  echo '-----'
  docker-compose down -v
}
trap finish EXIT

function main() {
  docker-compose build --pull tester
  local secret='
{
  "a": 1,
  "b": "xyz"
}
'
  export AWS_ACCESS_KEY_ID=x
  export AWS_SECRET_ACCESS_KEY=x
  export AWS_TEST_ENDPOINT=http://localstack:4584
  export AWS_DEFAULT_REGION=us-east-1

  docker-compose up -d localstack
  echo "waiting for localstack to be ready"
  until $(docker-compose exec localstack wget -qO /dev/null "http://localhost:8080/"); do
    sleep 2;
    printf ".";
  done
  echo "done"

  docker-compose exec \
   localstack \
   bash -c "
printenv AWS_DEFAULT_REGION
aws --endpoint='${AWS_TEST_ENDPOINT}' \
  secretsmanager create-secret \
  --name production/secret \
  --secret-string '${secret}'
"
  docker-compose run \
    -e AWS_TEST_ENDPOINT \
    -e AWS_ACCESS_KEY_ID \
    -e AWS_SECRET_ACCESS_KEY \
    -e AWS_DEFAULT_REGION \
    -v "$PWD:/summon-aws-secrets" \
    --rm \
    tester
}

main
