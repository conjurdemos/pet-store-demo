# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Security
- Switch to alpine base image, upgrade Spring Boot to 3.3.1
  [conjurdemos/pet-store-demo#77](https://github.com/conjurdemos/pet-store-demo/pull/77)

## [1.2.4] - 2024-03-22

### Security
- Upgrade Spring Boot to 3.2.4 for CVE-2024-22259 [conjurdemos/pet-store-demo#76](https://github.com/conjurdemos/pet-store-demo/pull/76)

## [1.2.3] - 2024-03-07

### Security
- Upgrade postgres connector to v42.7.2 for CVE-2023-1597, aa well as Spring Framework to 3.2.3, com.microsoft.sqlserver to 12.6.1, and maven enforcer to 3.4.1 [conjurdemos/pet-store-demo#74](https://github.com/conjurdemos/pet-store-demo/pull/74)

## [1.2.2] - 2024-02-09

### Changed
- Shifted tests to use `docker compose` instead of `docker-compose` 
  [conjurdemos/pet-store-demo#70](https://github.com/conjurdemos/pet-store-demo/pull/70)

### Security
- Upgrade spring-boot and other dependencies to latest versions
  [conjurdemos/pet-store-demo#71](https://github.com/conjurdemos/pet-store-demo/pull/71)
- Upgrade spring-boot to 3.2.1 and Java to 11
  [conjurdemos/pet-store-demo#70](https://github.com/conjurdemos/pet-store-demo/pull/70)
- Upgrade spring-boot to 3.0.7 for CVE-2023-20883
  [conjurdemos/pet-store-demo#63](https://github.com/conjurdemos/pet-store-demo/pull/63)
- Upgrade mysql-connector-java to v8.0.33
  [conjurdemos/pet-store-demo#62](https://github.com/conjurdemos/pet-store-demo/pull/62)
- Update DB drivers and test base images
  [conjurdemos/pet-store-demo#61](https://github.com/conjurdemos/pet-store-demo/pull/61)

## [1.2.1] - 2023-04-24

### Security
- Updated Spring boot to 3.0.6 and Dockerfile to eclipse-temurin
  [conjurdemos/pet-store-demo#60](https://github.com/conjurdemos/pet-store-demo/pull/60)
- Updated Spring boot to 3.0.2 and Dockerfile to openjdk:21
  [conjurdemos/pet-store-demo#58](https://github.com/conjurdemos/pet-store-demo/pull/58)
- Updated postgresql to 42.5.1 to resolve CVE-2022-41946
  [conjurdemos/pet-store-demo#57](https://github.com/conjurdemos/pet-store-demo/pull/57)
- Updated Spring boot to 2.7.5 to pull in fixes for jackson-databind for
   CVE-2022-42003 and CVE-2022-42004
   [conjurdemos/pet-store-demo#56](https://github.com/conjurdemos/pet-store-demo/pull/56)
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

[Unreleased]: https://github.com/conjurdemos/pet-store-demo/compare/v1.2.4...HEAD
[1.2.4]: https://github.com/conjurdemos/pet-store-demo/compare/v1.2.3...v1.2.4
[1.2.3]: https://github.com/conjurdemos/pet-store-demo/compare/v1.2.2...v1.2.3
[1.2.2]: https://github.com/conjurdemos/pet-store-demo/compare/v1.2.1...v1.2.2
[1.2.1]: https://github.com/conjurdemos/pet-store-demo/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/conjurdemos/pet-store-demo/compare/v1.1.0...v1.2.0
