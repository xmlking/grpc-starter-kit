# Make

using Makefile

> use `-n` flag for `dry-run`, `-s` or '--silent' flag to suppress echoing<br />

> `make release/*` `make deploy/*` are manually invoked.

## Targets

- default **VERSION** is `git tag`
- default **TYPE** is `service`

### proto

> codegen from proto

```bash
make proto
```

```bash
make proto_lint
make proto_breaking
make proto_format
make proto_check
make proto_clean
make proto_gen
```

### test

```bash
# unit tests
make test-unit TARGET=account
make test-unit TARGET=emailer
make test-unit TARGET=account TYPE=api
make test-unit TARGET=config TYPE=shared
make test-unit TARGET=demo TYPE=cmd
make test-unit

# integration tests
make test-inte TARGET=account
make test-inte TARGET=emailer
make test-inte TARGET=emailer TIMEOUT=30s
make test-inte

# end-to-end tests
make test-e2e TARGET=account
make test-e2e TARGET=emailer
make test-e2e
# e2e tests in CI envelopment with micro gRPC proxy
MICRO_PROXY_ADDRESS="localhost:8081" make test-e2e

# generate code coverage
make test-cover
# benchmark testing
make test-bench
make test-race

```

### e2e tests on CI

> trigger e2e tests on GitHub Actions

```bash
make deploy/e2e GITHUB_TOKEN=123...
```

### run

```bash
make run-account
make run TARGET=emailer
make run-emailer
make run-greeter ARGS="--server_address=127.0.0.1:8080"
make run-micro-cmd ARGS="--api_address=0.0.0.0:8088 api"
make run-demo-cmd
```

### lint

```bash
# lint all
make lint
# lint account service
make lint-account
make lint-account-service
# go-mod-upgrade checks if there are any updates for imports in your module.
# the main purpose is using it as a linter in continuous integration or in development process.
make sync
```

### build

```bash
# use git tag as VERSION
make build VERSION=v0.1.1
make build TARGET=account VERSION=v0.1.1
make build TARGET=account TYPE=service VERSION=v0.1.1
make build TARGET=emailer TYPE=service VERSION=v0.1.1
make build TARGET=account TYPE=api VERSION=v0.1.1
make build-account VERSION=v0.1.1
make build-account-api VERSION=v0.1.1
```

### release

> push tag to git

```bash
make release VERSION=v0.1.1 GITHUB_TOKEN=123...
```

### docker

#### Base
This will build and publish base image: `ghcr.io/xmlking/grpc-starter-kit/base`
Optionally set `nerdctl build --progress=plain` in `Makefile` to show **RUN** command output, to debug **nerdctl** builds

```bash
make docker_base VERSION=v0.2.1
# Sign and push the image with Keyless mode.Login if needed
# echo $GITHUB_PACKAGES_TOKEN | nerdctl login ghcr.io -u xmlking --password-stdin
nerdctl push --sign=cosign ghcr.io/xmlking/grpc-starter-kit/base:v0.2.1
nerdctl push --sign=cosign ghcr.io/xmlking/grpc-starter-kit/base:latest

# Verify the image with Keyless mode
nerdctl pull --verify=cosign ghcr.io/xmlking/grpc-starter-kit/base:latest
```

#### Services
Optionally set `nerdctl build --progress=plain` in `Makefile` to show **RUN** command output, to debug **nerdctl** builds

```bash
make docker-account VERSION=v0.1.5
make docker-account-service VERSION=v0.1.5
make docker TARGET=account VERSION=v0.1.5
make docker TARGET=account TYPE=service VERSION=v0.1.5
make docker TARGET=account DOCKER_REGISTRY=us.gcr.io DOCKER_CONTEXT_PATH=<MY_PROJECT_ID>/grpc-starter-kit
make docker TARGET=account DOCKER_REGISTRY=us.gcr.io DOCKER_CONTEXT_PATH=<MY_PROJECT_ID>/grpc-starter-kit BASE_VERSION=v0.2.1
# short hand for TARGET and TYPE args
make docker-emailer-service

# build all docker images for docker-compose
make docker
make docker DOCKER_REGISTRY=us.gcr.io
make docker VERSION=v0.3.2 BASE_VERSION=v0.1.0
make docker DOCKER_REGISTRY=ghcr.io DOCKER_CONTEXT_PATH=xmlking/grpc-starter-kit
make docker DOCKER_REGISTRY=ghcr.io DOCKER_CONTEXT_PATH=xmlking/grpc-starter-kit VERSION=v0.2.9
make docker DOCKER_REGISTRY=ghcr.io DOCKER_CONTEXT_PATH=xmlking/grpc-starter-kit VERSION=v0.2.9 BASE_VERSION=v0.1.0

# publish all microservices images
make docker_push

# remove all previous microservices images and any dangling images
make docker_clean

# explore and run
dive ghcr.io/xmlking/grpc-starter-kit/account-service:v0.1.5
nerdctl run --rm -it ghcr.io/xmlking/grpc-starter-kit/account-service:v0.1.5
# when built from `gcr.io/distroless/static:debug`, you can debug with:
nerdctl run --rm -it --entrypoint=sh ghcr.io/xmlking/grpc-starter-kit/account-service:v0.1.5 
# Sign and push with Keyless mode
# echo $GITHUB_PACKAGES_TOKEN | nerdctl login ghcr.io -u xmlking --password-stdin
nerdctl push --sign=cosign ghcr.io/xmlking/grpc-starter-kit/account-service:v0.1.5
# Verify the image with Keyless mode
nerdctl pull --verify=cosign ghcr.io/xmlking/grpc-starter-kit/account-service:v0.1.5
```


### kustomize

> generate `build/kubernetes.yaml` for given `overlay` and `namespace` using **kustomize**

```bash
make kustomize OVERLAY=production NS=default VERSION=v1.0.1
make kustomize OVERLAY=production NS=default
make kustomize OVERLAY=production
make kustomize NS=default
# defaults: ENV=local,  NS=default, VERSION=git tag
make kustomize
# build yaml files for local, prod overlays into ./build
make build/kustomize VERSION=v0.2.5
```

### Release

```bash
make release/draft VERSION=v0.1.1
```

At this point, you should inspect the release in the Github web UI. If it looks reasonable, proceed:

```bash
make release/publish GITHUB_TOKEN=123...
```

### Deploy

```bash
make deploy/e2e GITHUB_TOKEN=123...
make deploy/prod GITHUB_TOKEN=123...
```

### Reference

1. [A Makefile for your Go project](https://vincent.bernat.ch/en/blog/2019-makefile-build-golang)
