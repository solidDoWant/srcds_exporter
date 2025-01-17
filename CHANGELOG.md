## 1.5.1 / 2023-03-03

* [BUGFIX] fix changelog entry that caused the 1.5.0 release build to fail

## 1.5.0 / 2023-03-03

* [BUGFIX] tweak status parser regex - Seems that CS:GO changed the `status` command's output slightly. Resolves #7

## 1.4.3 / 2022-05-31

* [SECURITY] update gopkg.in/yaml.v3 to v3.0.1 (CVE-2022-28948)

## 1.4.2 / 2022-05-04

* [ENHANCEMENT] update minimum go version to 1.18
* [ENHANCEMENT] update all dependencies

## 1.4.1 / 2022-04-17

* [BUGFIX] tweak server query connection logic

## 1.4.0 / 2022-04-17

* [ENHANCEMENT] add server query support

## 1.3.2 / 2021-10-08

* [BUGFIX] fix issues with connections not being correctly created
* [ENHANCEMENT] improve server connection error log messages

## 1.3.1 / 2021-10-08

* [BUGFIX] try to address #4 issue by updating go-rcon library

## 1.3.0 / 2021-09-12

* [ENHANCEMENT] fix go linter issues and rework cmd entrypoint
* [ENHANCEMENT] update CI code and only push to quay and ghcr

## 1.2.12 / 2021-09-11

* [ENHANCEMENT] ci: use go 1.16 for build

## 1.2.11 / 2021-01-06

* [ENHANCEMENT] ci: use go 1.15 for build

## 1.2.10 / 2020-10-26

* [ENHANCEMENT] ci: fix changelog for new release

## 1.2.9 / 2020-10-26

* [ENHANCEMENT] ci: update to go 1.15

## 1.2.8 / 2020-06-07

* [ENHANCEMENT] ci: use GitHub Actions + promu

## v1.2.7 / 2020-06-07

* [ENHANCEMENT] ci: use GitHub Actions + promu

## 1.2.6 / 2020-02-24

* [ENHANCEMENT] ci: fixed CI release upload
