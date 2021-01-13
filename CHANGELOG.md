# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<a name="unreleased"></a>
## [Unreleased]


<a name="v0.1.0"></a>
## v0.1.0 - 2021-01-13
### Build
- **actions:** updated github actions
- **actions:** updated github actions
- **clog:** updating changelog
- **clog:** updating changelog
- **clog:** updating changelog
- **clog:** updating changelog
- **config:** updated k8s config and envoy config file
- **deploy:** polish envoy config files
- **deps:** bump actions/upload-release-asset from v1.0.1 to v1.0.2
- **deps:** bump actions/setup-go from v1 to v2.1.3
- **deps:** update actions/checkout requirement to v2.3.4
- **deps:** updated deps
- **deps:** updated google opentelemetry-operations-go to latest master
- **deps:** upgrade protoc-gen-go -> v1.24.0 protoc -> v3.12.3
- **deps:** updated to latest toolkit
- **deps:** updated to latest toolkit
- **deps:** bump actions/labeler from v2 to v3
- **deps:** updated opentelemetry and buf
- **docker:** switching to ghcr.io from docker.pkg.github.com
- **dockerfile:** improved docker build speed by caching more layers
- **entgo:** generate entgo code from latest ent CLI
- **makefile:** fix proto_clean command
- **refactor:** rename micro-starter-kit to grpc-starter-kit

### Ci
- **actions:** switching to GitHub Linter

### Docs
- **all:** updated README with grpc commands
- **deploy:** diagram update with Envoy Proxy
- **drawio:** updated README
- **drawio:** useing xxx.grawio.svg extention
- **entgo:** updated docs for entgo
- **makefile:** updated docs for make proto tasks
- **readme:** updated deployment diagram
- **readme:** updated to use deployment.drawio.svg
- **util:** updated tls util

### Feat
- **build:** experimenting Taskfile task runner and build tool
- **cedemo:** adding cedemo
- **emailer:** added cloudevents as message broker
- **errors:** switching to cockroachdb/errors from standard errors package
- **grpc:** experimenting k8s-resolver
- **linter:** adding Linter GitHub Action
- **middleware:** adding tags middleware
- **middleware:** adding middleware/Interceptors
- **orm:** adding https://entgo.io/ as ORM
- **orm:** replaced GORM with ENTGO
- **recorder:** adding translog middleware and recorder service back
- **rpclog:** adding ExcludeMethods option for rpclog Interceptors
- **telemetry:** adding telemetry example service `play`
- **tls:** adding password support for TLS, also added metadata form services
- **toolkit:** adding framework code to toolkit module
- **utils:** adding pkger ReadFile function

### Fix
- **broker:** shutdown
- **cloudevents:** updated cloudevents version
- **config:** config now support loading multiple files in the order
- **eventing:** polish emailer service with cloudevents

### Refactor
- **account:** account tested with native gRPC
- **account:** using DI for more components
- **broker:** moved broker code to https://github.com/xmlking/broker
- **config:** polish
- **config:** polish
- **config:** moved ClientConn creation to config.go
- **crypto:** moving crypto and some utils to toolkit
- **docker:** updated docker base image and dockerfiles
- **emailer:** improve error handling with cloudevents.Result
- **greeter:** switched greeter to native gRPC
- **internal:** refactor shared to internal, now using buf for codegen
- **logger:** mode docs for logger and config
- **proto:** polish protodef
- **proto:** polish protodef
- **proto:** polish protodef
- **toolkit:** moved middleware to toolkit
- **toolkit:** toolkit moved to its own module
- **wip:** work in pogress

### Style
- **fmt:** formate
- **format:** fix code format
- **format:** format code
- **format:** format with gofmt
- **format:** format proto
- **format:** format proto files with proto_format
- **gofmt:** format code with gofmt

### Test
- **gorm:** adding  Unit Tests for GORM With Sqlmock
- **grpcurl:** docs on grpcurl testing via Envoy
- **test:** test


[Unreleased]: https://github.com/xmlking/grpc-starter-kit/compare/v0.1.0...HEAD
