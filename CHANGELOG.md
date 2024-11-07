# Changelog

## [0.10.0] - 2024-11-07

### Added

- description attribute to impart_list resource

### Changed

- impart_list implementation to use PUT lists and PATCH list items endpoints

## [0.9.0] - 2024-10-29

### Added

- impart_label resource
- impart_tag_metadata resource
- labels attribute to rule scripts, rule test cases, lists, monitors, tags
- description attribute to test cases messages and assertions

## [0.8.3] - 2024-09-12

### Changed

- Added required attribute to impart_rule_test_case resource

### Fixed

- Errors on empty descriptions for impart_rule_test_case and impart_rule_script resources
- Array validation for impart_rule_test_case

## [0.8.2] - 2024-08-23

### Changed

- Added blocking_effect to impart_rule_script resource

## [0.8.1] - 2024-08-06

### Changed

- Added assertions to impart_rule_test_case resource

## [0.8.0] - 2024-07-24

### Changed

- Added impart_rule_test_case resource

## [0.7.0] - 2024-07-17

### Changed

- Added list functionality attribute

## [0.6.4] - 2024-07-10

### Changed

- Ignore api new fields fields
- Updated dependencies

## [0.6.3] - 2024-06-25

### Fixed

- Unnecessary diff for list of ips
- Rule script dependencies plan

## [0.6.2] - 2024-06-24

### Fixed

- Handle list ordering
- Ip list state
- Ignore items when not set

## [0.6.1] - 2024-06-11

### Fixed

- rule script content udpate

## [0.6.0] - 2024-06-11

### Added

- impart_list to manage lists
- impart_api_binding disabled attribute

## [0.5.0] - 2024-04-09

- Add impart_rule_script_dependencies resource

## [0.4.0] - 2024-03-28

- Add impart_connector Terraform data source
- Add impart_notification_template Terraform resource
- Add impart_monitor Terraform resource

## [0.3.0] - 2024-02-14

### Changed

- Add json support for the log binding configuration
- Add pattern_type field
- Rename grok_pattern to pattern

## [0.2.1] - 2024-01-16

### Changed

- Generate examples for log and api bindings

## [0.2.0] - 2024-01-16

### Changed

- Fixed source hash validation
- Renamed impart_binding to impart_api_binding
- Added advanced optios to impart_api_binding
- Added impart_log_binding

## [0.1.2] - 2023-12-29

### Changed

- Updated deps

## [0.1.1] - 2023-05-25

### Changed

- Updated docs

## [0.1.0] - 2023-05-02

### Added

- Initial release
