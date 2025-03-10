#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status
set -x  # Print commands before executing them

# Build binary with maximum performance optimizations
CGO_ENABLED=0 go build -o wirez \
  -ldflags="-s -w -extldflags '-static'" \
  -gcflags="all=-B" \
  -asmflags="all=-D GOAMD64=v3" \
  -tags=netgo,osusergo,release \
  -trimpath .

echo "Build complete: ./wirez"
