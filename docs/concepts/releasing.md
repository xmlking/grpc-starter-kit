# Releasing

[release page]: https://github.com/xmlking/grpc-starter-kit/releases
[`cloud-build-local`]: https://github.com/GoogleCloudPlatform/cloud-build-local
[google cloud build]: https://cloud.google.com/cloud-build
[semver]: https://semver.org

Scripts and configuration files for publishing a
`kustomize` release on the [release page].

## Build a release locally

Install [`cloud-build-local`], then run

```
./releasing/localbuild.sh
```

to build artifacts under `./dist`.

## Do a real (cloud) release

Get on an up-to-date master branch:

```
git fetch upstream
git checkout master
git rebase upstream/master
```

### review tags

```
git tag -l
git ls-remote --tags upstream
```

### define the new tag

Define the version per [semver] principles; it must start with `v`:

```
version=v3.0.0-pre
```

### if replacing a release...

Must delete the tag before re-pushing it.

Delete the tag locally:

```
git tag --delete $version
```

Delete it upstream:

```
# Disable push protection:
git remote set-url --push upstream git@github.com/xmlking/grpc-starter-kit.git

# The empty space before the colon effectively means delete the tag.
git push upstream :refs/tags/$version

# Enable push protection:
git remote set-url --push upstream no_push
```

Optionally visit the [release page] and delete
(what has now become) the _draft_ release for that
version.

### tag locally

```
git tag -a $version -m "Release $version"
```

### trigger the cloud build

Push the tag:

```
git push upstream $version
```

This triggers a job in [Google Cloud Build] to
put a new release on the [release page].

### Update release notes

Visit the [release page] and edit the release notes as desired.

# New Release Process

## Release

Following command bump **VERSION** number and push `changes` and `tag` to remote<br/>
Then, [GitHub Action](.github/workflows/release.yml) trigger `GoReleaser` process.

> NOTE: make sure you commit all changes before running this command.

```shell
### 
```shell
# dry-run: calculate the next version based on the commit types since the latest tag
cog bump --auto --dry-run 
# calculate the next version based on the commit types since the latest tag
cog bump --auto
```

* check [cog](https://docs.cocogitto.io/guide/#automatic-versioning) docs

## Test release

```shell
# you can verify your .goreleaser.yaml is valid by running the check command:
goreleaser check
# dry-run: lets run a "local-only" release to see if it works using the release command: 
goreleaser release --snapshot --rm-dist --skip-sign
# You can also use GoReleaser to build the binary only for a given GOOS/GOARCH, which is useful for local development:
# goreleaser build --single-target --snapshot --rm-dist
goreleaser build --snapshot --rm-dist
```

## Final Release

In order to release to GitHub, you'll need to export a `GITHUB_TOKEN` environment variable, which should contain a valid
GitHub token with the repository scope.

```shell
export GITHUB_TOKEN="YOUR_GH_TOKEN"
export SLACK_WEBHOOK=https://hooks.slack.com/services/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
# Now, change to main branch, create a tag and push it to GitHub:
git tag v0.1.1 -m "build(release): bump version to v0.1.1"
git push origin v0.1.1
# Now you can run GoReleaser at the root of your repository:
goreleaser release --rm-dist --skip-sign 
# That's all it takes!
```


Test multi-platform docker images
```shell
docker run --rm --platform linux/amd64 \
	ghcr.io/xmlking/grpc-starter-kit/account:latest
docker run --rm --platform linux/arm64 \
	ghcr.io/xmlking/grpc-starter-kit/account:latest
```
