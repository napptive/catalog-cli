# catalog-cli
A CLI to interact with Napptive Catalog.

This component is the entry point into Napptive catalog. 

## Prerequisites

- Component `catalog-manager` deployed

## Commands

To interact with the catalog, we offer several commands in `catalog-cli`.
 ```
 The catalog command provides a set of methods to interact with the Napptive Catalog

Usage:
  catalog [flags]
  catalog [command]

Examples:
$ catalog

Available Commands:
  help        Help about any command
  info        Get the principal information of an application.
  list        List the applications
  pull        Pull an application from catalog.
  push        Push an application in the catalog.
  remove      Remove an application from catalog.

Flags:
      --catalogAddress string   Catalog-manager host (default "catalog-manager")
      --catalogPort int         Catalog-manager port (default 7060)
      --consoleLogging          Pretty print logging
      --debug                   Set debug level
  -h, --help                    help for catalog
      --output string           Output format in which the results will be returned: json or table (default "table")
  -v, --version                 version for catalog

Use "catalog [command] --help" for more information about a command.
 ```


## Layout structure

The layout structure is based on the default golang-template layout.

https://github.com/golang-standards/project-layout

## Usage

A make file is provided with the following targets:

* clean: Remove build files
* test: Run the available tests
* build: Build the files for your local environment
* build-darwin: Build the files for MacOS
* build-linux: Build the files for Linux
* k8s: Generate the Kubernetes deployment files
* docker-prep: Prepare the Dockerfile folder with all the extra files
* docker-build: Build the Dockerfile locally
* docker-push: Push the image to the selected repository. You must make login before to push the docker image.

---
**Important**

If you are developing with MacOS/Darwin, you must install gnu-sed.

```
brew install gnu-sed
```
---

## Integration with Code Climate

This template is integrated with the Code Climate service.

[![Maintainability](https://api.codeclimate.com/v1/badges/d426ab46dd6c71fcb93b/maintainability)](https://codeclimate.com/repos/5f8e9d4ccdd59e0541004d1a/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/d426ab46dd6c71fcb93b/test_coverage)](https://codeclimate.com/repos/5f8e9d4ccdd59e0541004d1a/test_coverage)


## Integration with Github Actions

This template is integrated with GitHub Actions. You need to add the secret `CodeClimateRerporterID` in the repository settings.

![Check changes in the Main branch](https://github.com/napptive/go-template/workflows/Check%20changes%20in%20the%20Main%20branch/badge.svg)

## License

 Copyright 2020 Napptive

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
