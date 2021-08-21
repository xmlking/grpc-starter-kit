# Taskfile

**Taskfile** is a task runner like `Makefile` but have many benefits 

**[Features](https://taskfile.dev/#/?id=features)**
1. Truly cross-platform: works with MacOS/Linux/Windows seamlessly 
1. Task dependencies 
    1. Run tasks in parallel or sequential based on your needs
1. Prevent unnecessary
    1. work Great for code generation: you can easily prevent a task from running if a given set of files haven’t changed since last run (based either on its timestamp or content).
1. Each task in a specific directory if needed 

## Install

```bash
brew install go-task/tap/go-task
```

## Create

Use this [guide](https://github.com/go-task/task/blob/master/docs/usage.md) to add new tasks.
We recommend you to follow this [Style Guide](https://github.com/go-task/task/blob/master/docs/styleguide.md)

## Use 

You can use `--force` or `-f` if you want to force a task to run even when up-to-date.

Also, `task --status [tasks]...` will exit with a non-zero exit code if any of the tasks are not up-to-date.

Dry run mode (`--dry`) compiles and steps through each task, printing the commands that would be run without executing them.

With the flags `--watch` or `-w` task will watch for file changes and run the task again.

```bash
task --list
task --summary  proto:check
task proto --dry
```

### proto

> codegen from proto

```bash
task proto:default
```

```bash
task proto:lint
task proto:breaking
task proto:format
task proto:check
task proto:clean
task proto:generate
```

### go

```bash
task go:cli-deps
task go:clean
task go:format
task go:lint
task go:ci
task go:packages
task go:ent
```
#### test

```bash
task go:test
task go:test-unit
task go:test-inte
task go:test-e2e
task go:test-race
task go:test-cover
```

### go-mod

```bash
task mod:default
task mod:sync
task mod:verify
task mod:outdated
task mod:release
```

### k3d

```bash
task k3d:create # create cluster
task k3d:delete # delete cluster
task k3d:list # list clusters
task k3d:start # start cluster
task k3d:stop # stop cluster
```

### run

### GitHub Actions

If you want to install Task in GitHub Actions you can try using
[this action](https://github.com/arduino/actions/tree/master/setup-taskfile)
by the Arduino team:

```yaml
- name: Install Task
  uses: Arduino/actions/setup-taskfile@master
```

### Reference
1. [Style Guide](https://github.com/go-task/task/blob/master/docs/styleguide.md)
1. [Usage](https://github.com/go-task/task/blob/master/docs/usage.md)
