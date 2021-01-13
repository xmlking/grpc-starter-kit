# Features

## Features/Backlog

- [x] Monorepo - Sharing Code Between Microservices
- [x] gRPC microservices 
- [x] Versatile ingress gateway based on Envoy with Config Manager [esp-v2](https://github.com/GoogleCloudPlatform/esp-v2)
- [ ] Proxy-less Service Discovery and xDS based gRPC Load Balancer with [Traffic Director](https://cloud.google.com/blog/products/networking/traffic-director-supports-proxyless-grpc)
- [x] Input Validation with [protoc-gen-validate (PGV)](https://github.com/envoyproxy/protoc-gen-validate)
- [x] Config - Pluggable Sources, Mergeable Config, Environment Overlays 
- [x] Customizable Logging
- [x] Graph-Based ORM [ent](https://entgo.io/)
- [x] Dependency injection [Container](https://github.com/sarulabs/di), Try [wire](https://itnext.io/mastering-wire-f1226717bbac) next?
- [ ] Flow-Control, [adaptive system protection](https://github.com/alibaba/sentinel-golang/wiki)
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
