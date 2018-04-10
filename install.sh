#!/usr/bin/env bash

set -e

ARCH=`uname -m`

if [ "${ARCH}" != "x86_64" ]; then
  echo "summon-aws-secrets only works on 64-bit systems"
  echo "exiting installer"
  exit 1
fi

DISTRO=`uname | tr "[:upper:]" "[:lower:]"`

if [ "${DISTRO}" != "linux" ] && [ "${DISTRO}" != "darwin"  ]; then
  echo "This installer only supports Linux and OSX"
  echo "exiting installer"
  exit 1
fi

if test "x$TMPDIR" = "x"; then
  tmp="/tmp"
else
  tmp=$TMPDIR
fi
# secure-ish temp dir creation without having mktemp available (DDoS-able but not expliotable)
tmp_dir="$tmp/install.sh.$$"
(umask 077 && mkdir $tmp_dir) || exit 1

# do_download URL DIR
function do_download(){
  echo "Downloading $1"
  if   [[ $(type -t wget) ]]; then wget -q -c -O "$2" "$1" >/dev/null
  elif [[ $(type -t curl) ]]; then curl -sSL -o "$2" "$1"
  else
    error "Could not find wget or curl"
    return 1
  fi
}

LATEST_VERSION=$(curl -s https://api.github.com/repos/cyberark/summon-aws-secrets/releases/latest | grep -o '"tag_name": "[^"]*' | grep -o '[^"]*$')
BASEURL="https://github.com/cyberark/summon-aws-secrets/releases/download/"
URL=${BASEURL}"${LATEST_VERSION}/summon-aws-secrets-${DISTRO}-amd64.tar.gz"

ZIP_PATH="${tmp_dir}/summon-aws-secrets.tar.gz"
do_download ${URL} ${ZIP_PATH}

echo "Installing summon-aws-secrets ${LATEST_VERSION} into /usr/local/lib/summon"

sudo mkdir -p /usr/local/lib/summon
sudo tar -C /usr/local/lib/summon -zxvf ${ZIP_PATH}

echo "Success!"
echo "Run /usr/local/lib/summon/summon-aws-secrets for usage"
