# mockery

**mockery** provides the ability to easily generate mocks for golang interfaces used by **testify** package

## Installation

```bash
go install github.com/vektra/mockery/v2@latest
```

## Usage

`--testonly`    generate a mock in a _test.go file
`--inpackage`   generate a mock that goes inside the original package
`--case`        name the mocked file using casing convention [camel, snake, underscore] (default "camel")
`--note`        comment to insert into prologue of each generated file
`--keeptree`    keep the tree structure of the original interface files into a different repository. Must be used with XX
`--tags string` space-separated list of additional build tags to use
` -d, --dry-run` Do a dry run, don't modify any files

Annotate Interfaces with `//go:generate mockery ...`

```go
//go:generate mockery --name=UserRepository --case=snake --tags="mock" --inpackage --testonly
type UserRepository interface {
    Exist(ctx context.Context, model *ent.User) (bool, error)
    List(ctx context.Context, limit, page int, sort string, model *ent.User) (total int, users []*ent.User, err error)
    Get(ctx context.Context, id uuid.UUID) (*ent.User, error)
    Create(ctx context.Context, model *ent.User) (*ent.User, error)
    Update(ctx context.Context, model *ent.User) (*ent.User, error)
    DeleteFull(ctx context.Context, model *ent.User) (*ent.User, error)
    Delete(ctx context.Context, id uuid.UUID) (*ent.User, error)
    Count(ctx context.Context) (int, error)
}

```

Run `go generate` to produce mock objects

```bash
# example
go generate ./service/account/repository
go generate ./service/emailer/service

# generate all, this will also generate `ent`,  `wire` and `version` code
go generate ./...

go generate ./internal/version
```
