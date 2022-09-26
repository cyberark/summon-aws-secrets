# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Security
- Updated dependencies (aws-sdk-go -> v1.44.106, testify -> 1.8.0) and added replace
  to force golang.org/x/net to latest version for CVE-2022-27664
  [cyberark/summon-aws-secrets#65](https://github.com/cyberark/summon-aws-secrets/pull/65)

## [0.4.3] - 2022-06-08
### Changed
- Updated dependencies in go.mod (github.com/aws/aws-sdk-go -> v1.44.30, 
  github.com/stretchr/testify -> 1.7.2)
  [cyberark/summon-aws-secrets#64](https://github.com/cyberark/summon-aws-secrets/pull/64)

## [0.4.2] - 2022-06-01
### Changed
- Updated to Golang 1.18
  [cyberark/summon-aws-secrets#60](https://github.com/cyberark/summon-aws-secrets/pull/60)

## [0.4.1] - 2021-05-10
### Added
- Added a build for Apple Silicon
  [cyberark/summon-aws-secrets#55](https://github.com/cyberark/summon-aws-secrets/issues/55)

### Changed
- The project Golang version is updated from the end-of-life v1.13 to v1.17.
  [cyberark/summon-aws-secrets#54](https://github.com/cyberark/summon-aws-secrets/pull/54)

## [0.4.0] - 2020-09-11
### Changed
- Update aws-sdk-go to v1.34.6 to enable using new environment variables that AWS supports, such as
  `AWS_ROLE_ARN` which is required for EKS.
  [PR cyberark/summon-aws-secrets#37](https://github.com/cyberark/summon-aws-secrets/pull/37)
- Upgraded build/test images to use Golang 1.13 instead of 1.11.
  [cyberark/summon-aws-secrets#27](https://github.com/cyberark/summon-aws-secrets/issues/27)

### Fixed
- Made installer script more robust to target environments.
  [cyberark/summon-aws-secrets#16](https://github.com/cyberark/summon-aws-secrets/issues/16)

## [0.3.0] - 2019-03-06
### Added
- Converted to go modules
- Added ability to use `#` as a delimiter for variable IDs

## [0.2.0] - 2018-08-30
### Added
- Fetch AWS config from ec2metadata if on an ec2 instance [PR#5](https://github.com/cyberark/summon-aws-secrets/pull/5)

## 0.1.0 - 2018-04-04
### Added
- Initial release

[Unreleased]: https://github.com/cyberark/summon-aws-secrets/compare/v0.4.2...HEAD
[0.4.2]: https://github.com/cyberark/summon-aws-secrets/compare/v0.4.1...v0.4.2
[0.4.1]: https://github.com/cyberark/summon-aws-secrets/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/cyberark/summon-aws-secrets/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/cyberark/summon-aws-secrets/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/cyberark/summon-aws-secrets/compare/v0.1.0...v0.2.0
