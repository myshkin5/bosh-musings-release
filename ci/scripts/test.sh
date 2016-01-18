#!/usr/bin/env bash

set -e -x

cd bosh-musings-release

export GOPATH="${PWD}"
export PATH="$PATH:$GOPATH/bin"

git submodule update --init --recursive

go get golang.org/x/tools/cmd/goimports
go get github.com/onsi/ginkgo/ginkgo

./scripts/test
