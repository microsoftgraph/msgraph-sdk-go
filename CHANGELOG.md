# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

### Changed

## [0.17.0] - 2022-04-05

### Added

- Added support for vendor specific content types.
- Added support for 204 no content responses.

### Changed

- Weekly generation
- Breaking: kiota dependencies have been moved to their own repository.

## [0.16.0] - 2022-03-29

### Changed

- Weekly generation

## [0.15.0] - 2022-03-22

### Added

- Adds support for text endpoints deserialization. (.Count())

### Changed

- Weekly generation.
- Updated core reference for serialization nil check fix.
- Breaking change: simpler API design for page iterator.

## [0.14.0] - 2022-03-15

### Changed

- Updated reference to core for new serialization types
- Weekly generation

## [0.13.0] - 2022-03-08

### Changed

- Weekly generation, based of the new OpenAPI conversion, lots of new endpoints available.
- Fixed an issue with collections responses for OData functions.
- Breaking: Fixed a bug where using the raw request URL would result in invalid requests.
- Breaking: Splat Parsable interface to AdditionalPropertiesHolder.

## [0.12.0] - 2022-03-01

### Changed

- Weekly generation

## [0.11.1] - 2022-02-28

### Changed

- Fixed a bug where http client configuration would impact the default client configuration for other usages.

## [0.11.0] - 2022-02-23

### Changed

- Weekly generation

## [0.10.0] - 2022-02-16

### Changed

- Updated to code 0.0.9 (fixed request body compression, added error deserialization)
- Weekly generation

## [0.9.0] - 2022-02-08

### Changed

- Fixed a deserialization issue with arrays of enums #73
- Updated to core 0.0.8 (request body compression, response body decompression, page iterator)
- Weekly generation

## [0.8.0] - 2022-02-01

### Changed

- Weekly generation
- Fixed a serialization bug with collection properties.

## [0.7.0] - 2022-01-25

### Changed

- Weekly generation
- Breaking: added types for Date-only, time-only and duration instead of using string.

## [0.6.0] - 2022-01-18

### Changed

- Weekly generation

## [0.5.0] - 2022-01-11

### Changed

- Weekly generation

## [0.4.0] - 2022-01-04

Happy new year!

### Changed

- Weekly generation

## [0.3.2] - 2021-12-07

### Changed

- Fixes an issue where escaped properties would not be serialized properly #46

## [0.3.1] - 2021-12-01

### Changed

- Fixes an issue where select query parameters would not work #31.

## [0.3.0] - 2021-11-23

### Changed

- Weekly generation
- Fixes doc comments

## [0.2.1] - 2021-11-19

### Changed

- Made the client options public so they can be used by consumers when customizing middleware

## [0.2.0] - 2021-11-17

### Changed

- Weekly generation

## [0.1.2] - 2021-11-12

### Changed

- Fixes #14 a misalignement with telemetry specification
- Fixes #17 a bug where query parameters would be missing

## [0.1.1] - 2021-11-09

### Changed

- Fixes #9 an issue where deserialization would fail because of nil values

## [0.1.0] - 2021-11-09

### Added

- Weekly generation

## [0.0.3] - 2021-11-08

### Added

- Added support for setting the base url

## [0.0.2] - 2021-11-04

### Changed

- Fixes a bug where scalar collections would not deserialize

## [0.0.1] - 2021-10-22

### Added

- Initial release
