#!/bin/bash

set -e
set -x

CGO_ENABLED=0 GOAMD64=v4 go build -o wirez \
  -ldflags="-s -w -extldflags '-static'" \
  -tags "netgo osusergo release" \
  -trimpath .


echo "Build complete: ./wirez"
