# catalog-cli
A CLI to interact with [Napptive](https://napptive.com/) Catalog.

## Commands Documentation

The `catalog-cli` is a powerful tool that provides a range of methods for interacting with the Napptive Catalog. This documentation outlines the various commands, their usage, and available flags.

**Usage:**
```bash
catalog [flags]
catalog [command]
```
_Example_ To see a list of available applications in the catalog:
```bash
catalog list
 ```

**Available Commands:**
1. `help`: Provides help information about any specific command.
1. `info`: Retrieves principal information about a specific application.
1. `list`: Lists all the applications available in the catalog.
1. `pull`: Pulls a specific application from the catalog.
1. `push`: Pushes an application to the catalog.
1. `remove`: Removes an application from the catalog.

**Flags:**
1. `--catalogAddress`: Specifies the host of the Catalog Manager. (Default: "catalog-manager")
1. `--catalogPort`: Specifies the port of the Catalog Manager. (Default: 7060)
1. `--consoleLogging`: Enables pretty-print logging.
1. `--debug`: Sets the debug level.
1. `-h`, `--help`: Shows help for the `catalog` command.
1. `--output`: Specifies the output format for results: "json" or "table". (Default: "table")
1. `-v`, `--version`: Displays the version information for the `catalog` command.

**Command Usage:** 

For detailed information on a specific command, use the following pattern:
```bash
catalog [command] --help
```
Replace `[command]` with the specific command you want to know more about. For example, to get help for the `pull` command:
```bash
catalog pull --help
```
This command will provide specific usage instructions and additional flags for the `pull` command.

Remember that the `catalog` command is a versatile tool for managing applications in the Napptive Catalog. Use the provided commands and flags to perform various operations on the cataloged applications.

## Development Guide

If you are interested in contributing to the `catalog-cli` project, please read the [CONTRIBUTING](CONTRIBUTING.md) and [DEVELOPMENT](DEVELOPMENT.md) guide.

## Integration with Github Actions

This template is integrated with GitHub Actions. You need to add the secret `CodeClimateRerporterID` in the repository settings.

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
