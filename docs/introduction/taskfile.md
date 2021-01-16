# Taskfile

using **Taskfile**

You can use `--force` or `-f` if you want to force a task to run even when up-to-date.

Also, `task --status [tasks]...` will exit with a non-zero exit code if any of the tasks are not up-to-date.

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

### test


### run
