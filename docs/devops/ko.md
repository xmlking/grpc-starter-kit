# ko

**[ko](https://github.com/google/ko)** is a tool for build/publish/deploy container images  for Go applications

**ko** produces [SPDX](https://spdx.dev)-based SBOMs, but it can also produce **CycloneDX** with `--sbom=cyclonedx`.
If you are using ko’s multi-arch functionality, you will get an SBOM for each architecture

## Prerequisites

1. Docker Desktop
2. Cosign (optional)
   ```shell
   brew install cosign
   ```
3. bom (optional) - Create/View SPDX-compliant Bill of Materials
4. crane (optional) - [crane](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane.md) is a tool for interacting with remote images and registries.

## Install

```shell
# install ko
brew install ko
ko version
# install bom
go install sigs.k8s.io/bom/cmd/bom@latest
# install crane
brew install crane
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
# set you GITHUB_PACKAGES_TOKEN here:
export GITHUB_PACKAGES_TOKEN=ghp_YOUR_TOKEN
echo "$GITHUB_PACKAGES_TOKEN" | ko login ghcr.io --username ignored --password-stdin
```

## Configuration

## Usage

### Build

Use `--platform=linux/amd64,linux/arm64` or `--platform=all` flag for multi-arch functionality<br/>
Use `-v` flag for verbose<br/>
Use `GGCR_EXPERIMENT_ESTARGZ=1` for optimize images for [eStargz support](https://github.com/containerd/stargz-snapshotter/blob/v0.7.0/docs/stargz-estargz.md)

```shell
KO_DOCKER_REPO=ghcr.io/xmlking/grpc-starter-kit \
ko build --image-label org.opencontainers.image.source="https://github.com/xmlking/grpc-starter-kit"  -B  ./service/greeter/
```

KO_DOCKER_REPO=ghcr.io/mattmoor ko build -B ./test/

Push to local docker registry 
```shell

KO_DOCKER_REPO=ko.local \
#KO_DOCKER_REPO=k3s.local \ 
ko build -v --image-label org.opencontainers.image.source="https://github.com/xmlking/grpc-starter-kit"  -B ./service/greeter/
```

You can download the SBOM [ko](https://github.com/google/ko) produces with Sigstore’s [cosign](https://github.com/sigstore/cosign) via:

```shell
# Get the digest of an image
crane digest ghcr.io/xmlking/grpc-starter-kit/greeter:latest
# cosign download sbom ghcr.io/xmlking/grpc-starter-kit/greeter:latest --output-file=/tmp/sbom.spdx
cosign download sbom ghcr.io/xmlking/grpc-starter-kit/greeter@sha256:eaef37a8b9422d50dbf5c5b6366ea4a3e1cce2b0d4b5632998cf1ed842aad578 --output-file=/tmp/sbom.spdx
# show deps
bom document outline /tmp/sbom.spdx
# cleanup
rm -f /tmp/sbom.spdx
# Downloading the signature
cosign download signature ghcr.io/xmlking/grpc-starter-kit/greeter:latest
# Downloading the attestations both the sbom and build provenance are here as an attestation…
cosign download attestation ghcr.io/xmlking/grpc-starter-kit/greeter:latest
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
kustomize build config | ko resolve -f -
```

## Deploy

will replace `image: ko://github.com/xmlking/grpc-starter-kit/service/account` in yaml with latest built image and deploy it to k8s

```shell
ko apply -f config/
```

To teardown resources applied using ko apply, you can run `ko delete`:
```shell
ko delete -f config/
```

## Reference

- Sumanth's [ko-demo]( https://github.com/xmlking/ko-demo)
