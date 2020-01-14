# Contributing to the Pet Store Demo

## Prerequisites

Before getting started, you should install some developer tools.

1. [git][get-git] to manage source code
1. [Docker][get-docker] to manage dependencies and runtime environments

[get-docker]: https://docs.docker.com/engine/installation
[get-git]: https://git-scm.com/downloads

## Building
Run `./bin/build`. Requires Docker.

## Testing
To test against a specific database type, run `./test/test {db type}`, where
**db type** is `mysql`, `postgres`, or `mssql`.

## Releasing
- Update the [VERSION](VERSION)
- Update the [CHANGELOG](CHANGELOG.md)
- Submit your changes in a PR
- Once the PR has been reviewed and merged, at a git tag to the repo
- Add a github release for the new tag, and copy/paste the CHANGELOG data
  for the version into the gh release notes
