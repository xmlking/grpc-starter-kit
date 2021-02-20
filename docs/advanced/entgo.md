# entgo.io

[ent](https://entgo.io/docs/getting-started/) is facebook **ORM**  
 
## Installation

```bash
go install entgo.io/ent/cmd/ent
```

## Commands 

### Create Your First Schema

```bash
ent init User Profile
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
