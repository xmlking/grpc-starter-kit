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
GRPC_WEB_VERSION=1.3.0
wget -O ~/Downloads/protoc-gen-grpc-web "https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-darwin-x86_64"
chmod +x ~/Downloads/protoc-gen-grpc-web
mv  ~/Downloads/protoc-gen-grpc-web /usr/local/bin/protoc-gen-grpc-web

yarn global add grpc-tools
```

### third-party golang tools

```bash
# go better build tool
go install github.com/ahmetb/govvv@latest
# for static check/linter
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
# linter and tool for proto files
# *** (if you use brew to install buf, skip next line) ***
go install github.com/bufbuild/buf/cmd/buf@latest
# kind - kubernetes in docker (optional)
go install sigs.k8s.io/kind@latest
# go lang  build/publish/deploy tool (optional)
go install github.com/google/ko/cmd/ko@latest

# fetch protoc plugins into $GOPATH
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# to add tags to go struct
go install github.com/srikrsna/protoc-gen-gotag@latest

# Installing PGV can currently only be done from source: 
# from user's home directory, run
go get -d github.com/envoyproxy/protoc-gen-validate
cd ~/go/src/github.com/envoyproxy/protoc-gen-validate
git pull
make build

# goup checks if there are any updates for imports in your module.
# the main purpose is using it as a linter in continuous integration or in development process.
# Usage: goup -v -m ./...
go install github.com/rvflash/goup@latest
```

### Working with golang 1.18 (beta)

```bash
go install golang.org/dl/gotip@latest
gotip download
```

This will build the latest beta go SDK in `/Users/<username>/sdk/gotip` and copy **gotip** binary into `/Users/<username>/go/bin/gotip`.<br/>
`/Users/<username>/go/bin` should be already in your path.
