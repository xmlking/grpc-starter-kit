# Prerequisites

You should have:

**golang** installed via **brew**

## third-party

### third-party tools

```bash
# github CLI
brew install hub
# for mac, use brew to install protobuf
brew install protobuf
# VS Code plugin `vscode-proto3` need clang-format
brew install clang-format
# k8s tool similar to helm  (optional)
# generate fill k8s yaml files from overlays
brew install kustomize
# kubeval - validate one or more Kubernetes config files(optional)
brew tap instrumenta/instrumenta
brew install kubeval
# Manage Your lk8s In Style!
brew install derailed/k9s/k9s
# grpc cli client (optional)
brew install grpcurl
# bloomrpc is a UI client for gRPC (optional)
# install `bloomrpc` via `brew` into ~/Applications)
brew cask install --appdir=~/Applications bloomrpc
# gRPC mock server for testing
yarn global add bloomrpc-mock
# CHANGELOG generator
brew tap git-chglog/git-chglog
brew install git-chglog
# buf: proto tool https://buf.build/docs/tour-1
brew tap bufbuild/buf
brew install buf
# git flow
brew install git-flow-avh
```

#### For grpc-web development (optional)  

```bash
GRPC_WEB_VERSION=1.2.1
wget -O ~/Downloads/protoc-gen-grpc-web https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-darwin-x86_64
chmod +x ~/Downloads/protoc-gen-grpc-web
mv  ~/Downloads/protoc-gen-grpc-web /usr/local/bin/protoc-gen-grpc-web

yarn global add grpc-tools
```

### third-party golang tools

```bash
# go better build tool
GO111MODULE=off go get github.com/ahmetb/govvv
# for static check/linter
GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint
# linter and tool for proto files
# (if you use brew to install buf, skip next line)
GO111MODULE=on go get github.com/bufbuild/buf/cmd/buf
# linter/generator and tool for proto files
GO111MODULE=on go get github.com/uber/prototool/cmd/prototool@dev
# kind - kubernetes in docker (optional)
GO111MODULE=on go get sigs.k8s.io/kind
# go lang  build/publish/deploy tool (optional)
GO111MODULE=off go get github.com/google/ko/cmd/ko
# pkger cli
go install github.com/markbates/pkger/cmd/pkger

# fetch protoc plugins into $GOPATH
GO111MODULE=on go get google.golang.org/protobuf/cmd/protoc-gen-go
GO111MODULE=on go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0

# Installing PGV can currently only be done from source:
go get -d github.com/envoyproxy/protoc-gen-validate
cd ~/go/src/github.com/envoyproxy/protoc-gen-validate
git pull
make build

# goup checks if there are any updates for imports in your module.
# the main purpose is using it as a linter in continuous integration or in development process.
# Usage: goup -v -m ./...
GO111MODULE=on go get github.com/rvflash/goup
```

