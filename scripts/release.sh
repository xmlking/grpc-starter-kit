#!/usr/bin/env bash

# Usage:
#   hack/release.sh $VERSION
#
# Example:
#   hack/release.sh v0.1.1

VERSION=$1

go install github.com/ahmetb/govvv@latest
govvv build -o build/account-service service/account/main.go -version $VERSION
git tag $VERSION
git push origin $VERSION
