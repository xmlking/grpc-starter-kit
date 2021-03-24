# grpc-starter-kit

Microservices starter kit for **Golang**, aims to be developer friendly.

[![GoDoc](https://godoc.org/github.com/xmlking/grpc-starter-kit?status.svg)](https://godoc.org/github.com/xmlking/grpc-starter-kit)
[![Go](https://img.shields.io/github/go-mod/go-version/xmlking/grpc-starter-kit/develop)](https://golang.org/dl/)
[![Renovate dependency Status](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://renovatebot.com/)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

### Build Status
[![Build Status](https://github.com/xmlking/grpc-starter-kit/workflows/Test%20on%20Push/badge.svg?branch=develop)](https://github.com/xmlking/grpc-starter-kit/actions?query=branch%3Adevelop+event%3Apush+workflow%3A%22Test+on+Push%22)

### Quality
[![Go Report Card](https://goreportcard.com/badge/github.com/xmlking/grpc-starter-kit)](https://goreportcard.com/report/github.com/xmlking/grpc-starter-kit)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/xmlking/grpc-starter-kit.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/xmlking/grpc-starter-kit/alerts/)
[![codecov](https://codecov.io/gh/xmlking/grpc-starter-kit/branch/develop/graph/badge.svg)](https://codecov.io/gh/xmlking/grpc-starter-kit)
[![Language grade: Go](https://img.shields.io/lgtm/grade/go/g/xmlking/grpc-starter-kit.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/xmlking/grpc-starter-kit/context:go)


## Overview

<p align="center">
  <img src="docs/images/deployment.drawio.svg" width="60%">
</p>

### What you get

- [x] Monorepo - Sharing Code Between Microservices
- [x] gRPC microservices
- [x] Versatile ingress gateway based on Envoy with Config Manager [esp-v2](https://github.com/GoogleCloudPlatform/esp-v2)
- [ ] Proxy-less Service Discovery and xDS based gRPC Load Balancer with [Traffic Director](https://cloud.google.com/blog/products/networking/traffic-director-supports-proxyless-grpc)
- [x] Input Validation with [protoc-gen-validate (PGV)](https://github.com/envoyproxy/protoc-gen-validate)
- [x] Add/replace struct tags on generated protobuf messages [protoc-gen-gotag (PGGT)](https://github.com/srikrsna/protoc-gen-gotag)
- [x] Config - Pluggable Sources, Mergeable Config, Environment Overlays
- [x] Customizable Logging
- [x] Flexible [errors](https://github.com/cockroachdb/errors) lib: _PII-free, gRPC middleware, opt-in Sentry.io reporting_
- [x] Graph-Based ORM [ent](https://entgo.io/)
- [x] CRUD API with [ent](https://entgo.io/blog/2021/03/18/generating-a-grpc-server-with-ent/)
- [x] Dependency injection [Container](https://github.com/sarulabs/di), Try [wire](https://itnext.io/mastering-wire-f1226717bbac) next?
- [ ] Adaptive System Protection / Adaptive concurrency limits with 
        [Alibaba's Sentinel](https://github.com/alibaba/sentinel-golang/wiki), 
        [Netflix's concurrency-limits](https://medium.com/@NetflixTechBlog/performance-under-load-3e6fa9a60581), 
        [go-concurrency-limits](https://github.com/platinummonkey/go-concurrency-limits)
- [x] multi-stage-multi-target Dockerfile
- [x] One Step _build/publish/deploy_ with [ko](https://github.com/google/ko)
- [x] BuildInfo with [govvv](https://github.com/ahmetb/govvv)
- [x] Linting with [GolangCI](https://github.com/golangci/golangci-lint) linters aggregator
- [x] Linting Protos with [Buf](https://buf.build/docs/introduction)
- [x] Linting rest with [super-linter](https://github.com/github/super-linter/blob/master/docs/disabling-linters.md)
- [x] CICD Pipelines with [GitHub Actions](https://github.com/features/actions)
- [x] Kubernetes _Matrix Deployment_ with [Kustomize](https://kustomize.io/)
- [ ] Add k8s [healthchecks](https://github.com/heptiolabs/healthcheck) with [cmux](https://medium.com/@drgarcia1986/listen-grpc-and-http-requests-on-the-same-port-263c40cb45ff)
- [x] Feature Flags (enable/disable with zero cost)
- [ ] Observability via [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-go)
- [ ] Service Mesh with [Istio](https://istio.io/)
- [ ] GraphQL Gateway with [gqlgen](https://gqlgen.com/), [rejoiner](https://github.com/google/rejoiner),[gqlgen](https://github.com/Shpota/skmz)
- [ ] Switch to [Bazel Build](https://bazel.build/)
- [ ] Graceful / zero downtime upgrades [tableflip](https://github.com/cloudflare/tableflip)

## Getting Started

### Prerequisite

Refer [prerequisites](docs/introduction/prerequisites.md) docs

### Initial Setup

Also Refer [scaffolding](docs/introduction/scaffolding.md) docs

> clone the repo

```bash
git clone https://github.com/xmlking/grpc-starter-kit ~/Developer/Work/go/grpc-starter-kit
# pull dependencies (when every time `go.mod` changed)
go mod download
```

### Run

#### Database

By default, this project use embedded `sqlite3` database. if you want to use **postgreSQL**,

- start **postgres** via `docker-compose` command provided below
- uncomment `postgres` import statement and comment `sqlite` in `main.go`
- start micro server with `export export CONFY_FILES=/config/config.yml,/config/config.pg.yml` flag <br/>
  i.e., `CONFY_FILES=/config/config.yml,/config/config.pg.yml go run service/account/main.go`

```bash
# to start postgres in foreground
docker-compose up postgres
# to stop postgres
docker-compose down
# if needed, remove `postgres_data` volume to recreate database next time, when you start.
docker system prune --volumes
```

#### Services

> Node: `--server_address=localhost:5501x --broker_address=localhost:5502x` required only when you are behind VPN a.k.a `Work From Home`

```bash
# dev mode
make run-account
# or
make run-account ARGS="--server_address=localhost:55011 --broker_address=localhost:55021"
# or
go run srv/account/main.go \
--configDir deploy/bases/account-srv/config \
--server_address=localhost:55011 --broker_address=localhost:55021

make run-greeter
# or
make run-emailer ARGS="--server_address=localhost:55012 --broker_address=localhost:55022"

make run-emailer
# or
make run-emailer ARGS="--server_address=localhost:55013 --broker_address=localhost:55023"


# integration tests for config module via CMD
make run TARGET=demo TYPE=cmd
go run cmd/demo/main.go --help
go run cmd/demo/main.go --database_host=1.1.1.1 --database_port=7777

export APP_ENV=production
go run cmd/demo/main.go
```

### Test

Refer [testing](docs/testing/testing.md) docs

## GitOps

### Make

Refer [makefile](docs/introduction/makefile.md) docs

### Docker

Refer [docker](docs/devops/docker.md) docs

### Release

Refer [releasing](docs/concepts/releasing.md) docs

### Deploy

```bash
make docker DOCKER_REGISTRY=ghcr.io DOCKER_CONTEXT_PATH=xmlking/grpc-starter-kit
docker rmi $(docker images -f "dangling=true" -q)

# make kustomize OVERLAY=e2e NS=default VERSION=v0.1.0-440-g6c7fb7a
make kustomize
kubectl apply -f build/kubernetes.yaml

POD_NAME=$(kubectl get pods  -lapp.kubernetes.io/name=account-srv -o jsonpath='{.items[0].metadata.name}')
kubectl logs -f -c srv $POD_NAME

kubectl delete -f build/kubernetes.yaml
```

## Reference

### Project Docs

1. [prerequisites](docs/introduction/prerequisites.md)
2. [scaffolding](docs/introduction/scaffolding.md)
3. [makefile](docs/introduction/makefile.md)
4. [testing](docs/testing/testing.md)
5. [docker](docs/devops/docker.md)
6. [gitops](docs/advanced/gitops.md)
7. [Protobuf Style Guide](https://buf.build/docs/style-guide)
8. [Google Protobuf Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### External Docs
1. [Go Repo Layout](https://christine.website/blog/within-go-repo-layout-2020-09-07)
1. [examples](https://github.com/micro/examples) - example usage code for micro
1. [microhq](https://github.com/microhq) - a place for prebuilt microservices
1. [explorer](https://micro.mu/explore/) - which aggregates micro based open source projects
1. [micro-plugins](https://github.com/micro/go-plugins) extensible micro plugins
1. [step-by-step-guide-micro](https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro)
1. [micro-in-cn](https://github.com/micro-in-cn/tutorials/tree/master/examples)
1. [Platform Web](https://github.com/micro-in-cn/platform-web)
1. [grpc template](https://github.com/vtolstov/micro-template-grpc)
1. [Simple API backed by PostgresQL, Golang and gRPC](https://medium.com/@vptech/complexity-is-the-bane-of-every-software-engineer-e2878d0ad45a)
1. [securing gRPC connections with TLS](https://itnext.io/practical-guide-to-securing-grpc-connections-with-go-and-tls-part-2-994ef93b8ea9) via [certify](https://github.com/johanbrandhorst/certify)
## 🔗 Credits
- [atlas-app-toolkit](https://github.com/infobloxopen/atlas-app-toolkit)
- [dapr](https://github.com/dapr/dapr)
- [goyave](https://github.com/System-Glitch/goyave)
