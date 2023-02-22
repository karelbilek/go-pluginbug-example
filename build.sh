#!/usr/bin/env bash

set -ex

pushd ./work/myplugin
go build -buildmode=plugin .
popd

mv ./work/myplugin/myplugin.so ./mymain

pushd ./mymain
go run main.go