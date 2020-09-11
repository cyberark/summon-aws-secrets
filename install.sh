#!/usr/bin/env bash

set -e
set -o pipefail

error() {
  echo "ERROR: $@" 1>&2
  echo "Exiting installer" 1>&2
  exit 1
}

ARCH=`uname -m`

if [ "${ARCH}" != "x86_64" ]; then
  error "summon-aws-secrets only works on 64-bit systems"
fi

DISTRO=`uname | tr "[:upper:]" "[:lower:]"`

if [ "${DISTRO}" != "linux" ] && [ "${DISTRO}" != "darwin"  ]; then
  error "This installer only supports Linux and OSX"
fi

tmp="/tmp"
if [ ! -z "$TMPDIR" ]; then
  tmp=$TMPDIR
fi

# secure-ish temp dir creation without having mktemp available (DDoS-able but not exploitable)
tmp_dir="$tmp/install.sh.$$"
(umask 077 && mkdir $tmp_dir) || exit 1

# do_download URL DIR
do_download() {
  echo "Downloading $1"
  if [[ $(command -v wget) ]]; then
    wget -q -O "$2" "$1" >/dev/null
  elif [[ $(command -v curl) ]]; then
    curl --fail -sSL -o "$2" "$1" &>/dev/null || true
  else
    error "Could not find wget or curl"
    exit 1
  fi
}

# get_latest_version
get_latest_version() {
  local LATEST_VERSION_URL="https://api.github.com/repos/cyberark/summon-aws-secrets/releases/latest"
  local latest_payload

  if [[ $(command -v wget) ]]; then
    latest_payload=$(wget -q -O - "$LATEST_VERSION_URL")
  elif [[ $(command -v curl) ]]; then
    latest_payload=$(curl --fail -sSL "$LATEST_VERSION_URL")
  else
    error "Could not find wget or curl"
    exit 1
  fi

  echo "$latest_payload" | # Get latest release from GitHub api
    grep '"tag_name":' | # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/' # Pluck JSON value
}

LATEST_VERSION=$(get_latest_version)

echo "Using version number: $LATEST_VERSION"

BASEURL="https://github.com/cyberark/summon-aws-secrets/releases/download/"
URL=${BASEURL}"${LATEST_VERSION}/summon-aws-secrets-${DISTRO}-amd64.tar.gz"

ZIP_PATH="${tmp_dir}/summon-aws-secrets.tar.gz"
do_download ${URL} ${ZIP_PATH}

echo "Installing summon-aws-secrets ${LATEST_VERSION} into /usr/local/lib/summon"

if sudo -h >/dev/null 2>&1; then
  sudo mkdir -p /usr/local/lib/summon
  sudo tar -C /usr/local/lib/summon -o -zxvf ${ZIP_PATH} >/dev/null
else
  mkdir -p /usr/local/lib/summon
  tar -C /usr/local/lib/summon -o -zxvf ${ZIP_PATH} >/dev/null
fi

echo "Success!"
echo "Run /usr/local/lib/summon/summon-aws-secrets for usage"
