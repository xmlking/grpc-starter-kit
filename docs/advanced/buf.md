# Buf

[Buf](https://buf.build/) is a tool for __Protobuf__ files:

- [Linter](https://buf.build/docs/lint-usage) that enforces good API design choices and structure.
- [Breaking change detector](https://buf.build/docs/breaking-usage) that enforces compatibility at the source code or wire level
- Configurable file [builder](https://buf.build/docs/build-overview) that produces [Images](https://buf.build/docs/build-images) our extension of [FileDescriptorSets](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/descriptor.proto)

## Prerequisites

```bash
# buf: proto tool https://buf.build/docs/tour-1
# or use `go install` to install Buf
go install github.com/bufbuild/buf/cmd/buf@latest

# Install protoc plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/srikrsna/protoc-gen-gotag@latest
```

## Developer Workflow

### Info

```bash
# To list all files Buf is configured to use:
buf ls-files
# To see your currently configured lint or breaking checkers:
buf config ls-lint-rules
buf config ls-breaking-rules
# To see all available lint checkers independent of configuration/defaults:
buf config ls-lint-rules --all
```

### Build

```bash
# check
buf build -o /dev/null
buf build -o image.bin
```

### Lint

```bash
buf lint
# We can also output errors in a format you can then copy into your buf.yaml file
buf lint --error-format=config-ignore-yaml
# Run breaking change detection
# for dev local
buf breaking --against image.bin
buf breaking --against '.git#branch=main'
buf breaking --against '.git#branch=main,subdir=proto/mkit'

# for CI
export HTTPS_GIT=https://github.com/xmlking/yeti.git
buf breaking --against "$(HTTPS_GIT)#branch=main"
```

### Generate

Cleanup
```bash
rm -rf dkit
rm -rf mkit
```

Generate
```bash
buf generate proto/dkit
buf generate proto/mkit
buf generate
# FIXME: https://github.com/bufbuild/buf/issues/560  
# WORKAROUND: https://github.com/srikrsna/protoc-gen-gotag/issues/26
buf generate --template buf.gen.tag.yaml
# to generate into `gen` directory
buf generate -o gen
```

### Buf Modules
Buf support modules [dependencies](https://docs.buf.build/tour/add-a-dependency) and modules registry [BSR](https://docs.buf.build/bsr/overview)

#### Login to buf registry. 
This command will create `/Users/<username>/.netrc` file

```bash
# fill the username and token
BUF_USER=
BUF_API_TOKEN=
echo ${BUF_API_TOKEN} | buf registry login --username ${BUF_USER} --token-stdin
```

#### Create buf lock files
one-time-setup
```bash
cd proto/dkit
buf mod update
cd proto/mkit
buf mod update
```
This will pull deps into `$HOME/.cache/buf`

#### Create a Repository
```bash
buf beta registry repository create buf.build/chintha/dkit --visibility private
buf beta registry repository create buf.build/chintha/dkit --visibility public

buf beta registry repository create buf.build/chintha/mkit --visibility private
```

#### List Modules
```bash
buf beta registry repository list  buf.build --page-size 100
buf beta registry repository get    buf.build/googleapis/googleapis
```

#### Push the Module
```bash
cd proto/dkit
buf push
cd proto/mkit
buf push
```

### Format

```bash
# FIXME buf don't have proto formatter yet 
prototool format -w proto;
```

## Tools

### grpcurl

```bash
# To use Buf-produced FileDescriptorSets with grpcurl on the fly:
grpcurl -protoset <(buf build -o -) ...
```

### ghz

```bash
# To use Buf-produced FileDescriptorSets with ghz on the fly:
ghz --protoset <(buf build -o -) ...
```

## Reference

1. [Style Guide](https://buf.build/docs/style-guide)
2. [Buf docs](https://buf.build/docs/introduction)
3. [Buf Example](https://github.com/bufbuild/buf-example/blob/master/Makefile)
4. [Buf Schema Registry](https://buf.build/docs/roadmap)
5. [Why adopt ProtoBuf](https://itnext.io/a-minimalist-guide-to-protobuf-1f24fbca0e2d)
