#!/bin/bash

snap install go --classic

export GOROOT=/snap/go/current
export PATH=/snap/bin:$PATH
