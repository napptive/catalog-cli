/*
 Copyright 2022 Napptive

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package commands

import (
	"github.com/napptive/catalog-cli/v2/pkg/catalog/operations"
	"github.com/spf13/cobra"
)

var deployCmdLongHelp = `Deploy an application in an associated Playground.
This command requires the login token obtained by executing

playground login
`

var deployCmdShortHelp = `Deploy a catalog application in the playground`

var deployCmd = &cobra.Command{
	Use:   "deploy <[catalog/]namespace/appName[:tag]> <account>/<environment> <target_playground>",
	Long:  deployCmdLongHelp,
	Short: deployCmdShortHelp,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		catalog, err := operations.NewCatalog(&cfg)
		crashOnError(err)
		crashOnError(catalog.Push(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
