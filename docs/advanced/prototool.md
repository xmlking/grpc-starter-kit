# prototool

[prototool](https://github.com/uber/prototool) is used to scaffold new **Protobuf** files and generate _Go/Java/JS_ code from it.

## Developer Workflow

### Scaffold

> Create new `proto` file from template

```bash
 prototool create proto/sumo.proto
```

### Generate

> Generating _Go/Java/JS_ from _protobuf_ definitions

```bash
prototool generate proto --dry-run
prototool generate proto
prototool generate proto --debug
```
