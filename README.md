# catalog-cli
A CLI to interact with [Napptive](https://napptive.com/) Catalog.

## Commands Help

The `catalog` is a powerful cli tool that provides a range of commands for interacting with the Napptive Catalog.

```bash
$ catalog --help
The catalog command provides a set of methods to interact with the Napptive Catalog

Usage:
  catalog [flags]
  catalog [command]

Examples:
$ catalog

Available Commands:
  change-visibility Update application visibility
  completion        Generate the autocompletion script for the specified shell
  deploy            Deploy a catalog application in the playground
  help              Help about any command
  info              Get the principal information of an application.
  list              List the applications
  pull              Pull an application from catalog.
  push              Push an application in the catalog.
  remove            Remove an application from catalog.
  summary           Get te catalog summary.

Flags:
      --authEnable                   JWT authentication enable (default true)
      --catalogAddress string        Catalog-manager host (default "catalog.playground.napptive.dev")
      --catalogPort int              Catalog-manager port (default 7060)
      --consoleLogging               Pretty print logging
      --debug                        Set debug level
  -h, --help                         help for catalog
      --output string                Output format in which the results will be returned: json or table (default "table")
      --skipCertValidation           enables ignoring the validation step of the certificate presented by the server
      --usePlaygroundConfiguration   Set to false to avoid reading the .playground.yaml file (default true)
      --useTLS                       TLS connection is expected with the Catalog manager (default true)
  -v, --version                      version for catalog

Use "catalog [command] --help" for more information about a command.
```

## Development Guide

If you are interested in contributing to the `catalog-cli` project, please read the [CONTRIBUTING](CONTRIBUTING.md) and [DEVELOPMENT](DEVELOPMENT.md) guide.

## Integration with Github Actions

This template is integrated with GitHub Actions. You need to add the secret `CodeClimateRerporterID` in the repository settings.

![Check changes in the Main branch](https://github.com/napptive/catalog-cli/workflows/Check%20changes%20in%20the%20Main%20branch/badge.svg)
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
