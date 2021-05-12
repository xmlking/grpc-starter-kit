#!/usr/bin/env zx

await $`cat go.mod | grep module`

let branch = await $`git branch --show-current`
await $`echo deploy --branch=${branch}`

await Promise.all([
  $`sleep 1; echo 1`,
  $`sleep 2; echo 2`,
  $`sleep 3; echo 3`,
])

let name = 'foo bar'
await $`mkdir /tmp/${name}`

cd('docs')
await $`pwd`
