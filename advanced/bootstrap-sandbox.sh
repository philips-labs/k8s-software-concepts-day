#!/usr/bin/env bash

git config --global user.name "Software Concepts"
git config --global user.email "software.concepts@philips.com"

./install-operator.sh
./install-registry.sh

snap install go --classic

export GOROOT=/snap/go/current
export PATH=/snap/bin:$PATH
