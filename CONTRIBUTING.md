# Contributing to the Pet Store Demo

For general contribution and community guidelines, please see the [community repo](https://github.com/cyberark/community).

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

## Contributing

1. [Fork the project](https://help.github.com/en/github/getting-started-with-github/fork-a-repo)
2. [Clone your fork](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/cloning-a-repository)
3. Make local changes to your fork by editing files
3. [Commit your changes](https://help.github.com/en/github/managing-files-in-a-repository/adding-a-file-to-a-repository-using-the-command-line)
4. [Push your local changes to the remote server](https://help.github.com/en/github/using-git/pushing-commits-to-a-remote-repository)
5. [Create new Pull Request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request-from-a-fork)

From here your pull request will be reviewed and once you've responded to all
feedback it will be merged into the project. Congratulations, you're a
contributor!
