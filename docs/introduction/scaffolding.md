# Scaffolding

> setup new golang project workspace from scratch using `go-micro` cli

```bash
mkdir -p ~/Developer/Work/go/grpc-starter-kit
cd ~/Developer/Work/go/grpc-starter-kit
go mod init github.com/xmlking/grpc-starter-kit
mkdir service api fnc

# scaffold modules
micro new --fqdn gkit.service.account --type srv --alias account service/account
micro new --fqdn gkit.service.greeter --type srv --alias greeter service/greeter
micro new --fqdn gkit.service.emailer --type srv --alias emailer service/emailer
# micro new --fqdn="gkit.service.emailer" --type="service"  \
# --alias="emailer"  --plugin=client/selector=static:broker=nats service/emailer
micro new --fqdn gkit.service.recorder --type srv --alias recorder service/recorder
```

## Setup project

### GitFlow setup

```bash
git flow init -d
```

### CHANGELOG generator setup

```bash
git-chglog --init
```
