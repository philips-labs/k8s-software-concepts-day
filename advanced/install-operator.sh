#!/bin/bash

# This script follows instructions from https://v1-11-x.sdk.operatorframework.io/docs/installation/

# Download the binary for your platform:
export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac)
export OS=$(uname | awk '{print tolower($0)}')
export RELEASE_VERSION=v1.11.0 # Search all project source for "1-11" and change all if this version updates
export OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download/$RELEASE_VERSION
curl -LO ${OPERATOR_SDK_DL_URL}/operator-sdk_${OS}_${ARCH}

# Verify the downloaded binary
# Import the operator-sdk release GPG key
gpg --keyserver keyserver.ubuntu.com --recv-keys 052996E2A20B5C7E

# Download the checksums file and its signature, then verify the signature:
curl -LO ${OPERATOR_SDK_DL_URL}/checksums.txt
curl -LO ${OPERATOR_SDK_DL_URL}/checksums.txt.asc
gpg -u "Operator SDK (release) <cncf-operator-sdk@cncf.io>" --verify checksums.txt.asc

# Make sure the checksums match
grep operator-sdk_${OS}_${ARCH} checksums.txt | sha256sum -c -
if [ $? -ne 0 ]; then
    echo "Checksum mismatch with Operator SDK installation source."
    return 1
fi

# Install release binary into PATH
chmod +x operator-sdk_${OS}_${ARCH} && sudo mv operator-sdk_${OS}_${ARCH} /usr/local/bin/operator-sdk
