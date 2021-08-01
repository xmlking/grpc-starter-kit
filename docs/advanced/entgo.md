# entgo.io

[ent](https://entgo.io/docs/getting-started/) is facebook **ORM**  
 
## Installation

```bash
go install entgo.io/ent/cmd/ent@latest
# project repo, run:
go get -u github.com/masseelch/elk
```

## Commands 

### Create Your First Schema

```bash
ent init User Profile
# or
go run -mod=mod entgo.io/ent/cmd/ent init User Profile
```

Edit **schema** in `ent/schema` then:

### Generate Assets

Run `entc generate` from the root directory of the project, or use `go generate`:
```bash
go generate ./ent
# or
ent generate --idtype string ./ent/schema
```

# Schema Description

```bash
ent describe ./ent/schema
```

## Reference

* [Generate a fully-working Go gRPC server in two minutes with Ent](https://entgo.io/blog/2021/03/18/generating-a-grpc-server-with-ent/)
* [Generate a fully-working Go CRUD HTTP API with Ent](https://entgo.io/blog/2021/07/29/generate-a-fully-working-go-crud-http-api-with-ent/)
