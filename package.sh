#!/bin/bash -e

GLOB='summon-aws-secrets-*-amd64'

echo "==> Packaging..."

rm -rf output/dist && mkdir -p output/dist

pushd output

for binary_name in $GLOB; do
  pushd dist

  cp ../$binary_name summon-aws-secrets && \
  tar -cvzf $binary_name.tar.gz summon-aws-secrets && \
  rm -f summon-aws-secrets

  popd
done

popd

# # Make the checksums
echo "==> Checksumming..."
pushd output/dist
shasum -a256 * > SHA256SUMS.txt
popd

export AWS_ACCESS_KEY_ID=AKIAJUOXDVEEWZTTTFQQ
export AWS_SECRET_ACCESS_KEY=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
export AWS_REGION=us-west-1