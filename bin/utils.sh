#!/bin/bash -e

function retrieve_cyberark_ca_cert() {
  CURRENT_DIR="$(dirname "${BASH_SOURCE[0]}")"

  pushd "$CURRENT_DIR/.."
    # On CyberArk dev laptops, golang module dependencies are downloaded with a
    # corporate proxy in the middle. For these connections to succeed we need to
    # configure the proxy CA certificate in build containers.
    #
    # To allow this script to also work on non-CyberArk laptops where the CA
    # certificate is not available, we update container certificates based on
    # a (potentially empty) certificate directory, rather than relying on the
    # CA file itself.
    mkdir -p build_ca_certificate

    # Only attempt to extract the certificate if the security
    # command is available.
    #
    # The certificate file must have the .crt extension to be imported
    # by `update-ca-certificates`.
    if command -v security &>/dev/null
    then
      security find-certificate \
        -a -c "CyberArk Root CA" \
        -p > build_ca_certificate/cyberark_root.crt
    fi
  popd
}