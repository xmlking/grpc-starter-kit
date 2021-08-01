# Account Service

This is the Account service

showcase

1. Implements basic CRUD API
2. Multiple handlers, repositories, subscribers
3. Publishing events
4. EntORM data access
5. Config Managment
6. Custom Logging

## Usage

### Build the binary

```bash
make build TARGET=account TYPE=service VERSION=v0.1.1
```

### Run the service

```bash
make run-account
# or
go run service/account/main.go
```

### Build a docker image

```bash
make docker TARGET=account TYPE=service VERSION=v0.1.1
```

### Test the service

```bash
grpcurl -plaintext -protoset <(buf build -o -) list
grpcurl -plaintext -protoset <(buf build -o -) describe mkit.service.account.user.v1.UserService

# test Create API directly
grpcurl -plaintext \
-protoset <(buf build -o -) \
-d '{"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}' \
 0.0.0.0:8080 mkit.service.account.user.v1.UserService/Create

# test Create API directly with TLS
grpcurl -insecure \
-protoset <(buf build -o -) \
-d '{"username": "sumo1", "firstName": "sumo1", "lastName": "demo1", "email": "sumo1@demo.com"}' \
 0.0.0.0:8080 mkit.service.account.user.v1.UserService/Create

# test List API directly
grpcurl -plaintext \
-protoset <(buf build -o -) \
-d '{}' 0.0.0.0:8080 mkit.service.account.user.v1.UserService/List
```
