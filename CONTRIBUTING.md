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
To test against a specific database type, run `./test/test {db type}`.
