# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Security
- Updated Spring boot to 2.7.4 to pull in fixes for jackson-databind for
  CVE-2022-42003 and CVE-2022-42004
  [conjurdemos/pet-store-demo#55](https://github.com/conjurdemos/pet-store-demo/pull/55)
- Updated all dependency versions in pom.xml and added maven-enforcer-plugin
  [conjurdemos/pet-store-demo#54](https://github.com/conjurdemos/pet-store-demo/pull/54)
- Upgraded Postgres to 42.4.1 to resolve CVE-2022-31197
  [conjurdemos/pet-store-demo#51](https://github.com/conjurdemos/pet-store-demo/pull/51)
- Upgraded Spring to 2.6.9
  [conjurdemos/pet-store-demo#49](https://github.com/conjurdemos/pet-store-demo/pull/49)
- Upgraded Spring to 2.6.7 & Maven/Ruby containers to latest versions
  [conjurdemos/pet-store-demo#48](https://github.com/conjurdemos/pet-store-demo/pull/48)
- Upgraded Postgres to 42.3.2 to resolve CVE-2022-21724
  [conjurdemos/pet-store-demo#45](https://github.com/conjurdemos/pet-store-demo/pull/45)

## [1.2.0] - 2021-08-04

### Added
- Includes binaries for 'summon' and 'summon-conjur' in the Pet Store
  application Docker image. This avoids the extra step of building a custom
  application image to include these binaries in scenarios where they are
  needed.
  [conjurdemos/pet-store-demo#35](https://github.com/conjurdemos/pet-store-demo/pull/35)

## Security
- Update Spring Boot to 2.5.3 to address multiple high & critical security issues
  [conjurdemos/pet-store-demo#37](https://github.com/conjurdemos/pet-store-demo/pull/37)

## [1.1.0] - 2020-01-13

### Added
- Added support for MSSQL

## [1.0.0] - 2018-05-17

The first tagged version.

[Unreleased]: https://github.com/cyberark/secretless-broker/compare/v1.2.0...HEAD
[1.2.0]: https://github.com/cyberark/secretless-broker/compare/v1.1.0...v1.2.0
