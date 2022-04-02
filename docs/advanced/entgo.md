# entgo.io

[ent](https://entgo.io/docs/getting-started/) is facebook **ORM**  
 
## Installation

```bash
go install entgo.io/ent/cmd/ent@latest
go get -u entgo.io/contrib/entproto
# optional - in project repo, run:
go get ariga.io/ogent@main
```

## Commands 

### Create Your First Schema

```bash
ent init User Profile
# or
go run entgo.io/ent/cmd/ent init User Profile
```

Edit **schema** in `ent/schema` then:

### Generate Assets

Run `entc generate` from the root directory of the project, or use `go generate`:
```bash
go generate ./ent
# temp fix for workspace mode
GOWORK=off go generate ./ent
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
* [Sync Changes to External Data Systems using Ent Hooks](https://entgo.io/blog/2021/11/1/sync-to-external-data-systems-using-hooks/)
