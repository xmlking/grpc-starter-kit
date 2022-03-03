# ko

**[ko](https://github.com/google/ko)** is a tool for build/publish/deploy container images  for Go applications

**ko** produces [SPDX](https://spdx.dev)-based SBOMs, but it can also produce **CycloneDX** with `--sbom=cyclonedx`.
If you are using ko’s multi-arch functionality, you will get an SBOM for each architecture

### Prerequisites

1. Docker Desktop
2. Cosign (optional)
   ```shell
   brew install cosign
   ```

## Install

```shell
brew install ko
ko version
```

### Completions

```shell
ko completion [bash|zsh|fish|powershell] --help
```

Or, you can source it directly:
```shell
source <(ko completion)
```

### Authenticate to your Registry
This command is same as `docker login ghcr.io --username myusername --password-stdin`<br/>
Will update `.docker/config.json`

```shell
# set you GITHUB_TOKEN here:
GITHUB_TOKEN=ghp_XYZ
echo "$GITHUB_TOKEN" | ko login ghcr.io --username ignored --password-stdin
```

## Configuration

## Usage

### Build

Use `--platform=linux/amd64,linux/arm64` flag for multi-arch functionality<br/>
Use `-v` flag for verbose

```shell
KO_DOCKER_REPO=ghcr.io/xmlking/grpc-starter-kit \
ko build --image-label org.opencontainers.image.source="https://github.com/xmlking/grpc-starter-kit" -B  ./service/greeter/
```

KO_DOCKER_REPO=ghcr.io/mattmoor ko build -B ./test/

Push to local docker registry 
```shell
KO_DOCKER_REPO=ko.local \
ko build -v --image-label org.opencontainers.image.source="https://github.com/xmlking/grpc-starter-kit" -B ./service/greeter/
```

You can download the SBOM [ko](https://github.com/google/ko) produces with Sigstore’s [cosign](https://github.com/sigstore/cosign) via:

```shell
cosign download sbom ghcr.io/xmlking/grpc-starter-kit/greeter:latest --output-file=sbom.spdx
bom document outline sbom.spdx
crane digest ghcr.io/xmlking/grpc-starter-kit/greeter:latest
```

### Publish

```sh
ko publish github.com/xmlking/ko-demo
# publish to local docker repo
ko resolve -L -f deploy/
```

### deps
Print Go module dependency information about the ko-built binary in the image


### Kustomize

```shell
ustomize build config | ko resolve -f -
```


## Reference

- Sumanth's [ko-demo]( https://github.com/xmlking/ko-demo)
