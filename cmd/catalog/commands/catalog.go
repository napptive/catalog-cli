/**
 * Copyright 2021 Napptive
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package commands

import (
	"github.com/napptive/catalog-cli/v2/pkg/catalog/operations"
	"github.com/spf13/cobra"
)

var catalogPushCmdLongHelp = `Push an application in the catalog. \
The application should be named: [catalog/]namespace/appName[:tag] `

var catalogPushCmdShortHelp = `Push an application in the catalog.`

var pushCmd = &cobra.Command{
	Use:   "push <[catalog/]namespace/appName[:tag]> <application_path>",
	Long:  catalogPushCmdLongHelp,
	Short: catalogPushCmdShortHelp,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		catalog, err := operations.NewCatalog(&cfg)
		crashOnError(err)
		crashOnError(catalog.Push(args[0], args[1]))
	},
}

var catalogPullCmdLongHelp = `Pull an application from catalog.`

var catalogPullCmdShortHelp = `Pull an application from catalog.`

var pullCmd = &cobra.Command{
	Use:   "pull <[catalog/]namespace/appName[:tag]>",
	Long:  catalogPullCmdLongHelp,
	Short: catalogPullCmdShortHelp,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		catalog, err := operations.NewCatalog(&cfg)
		crashOnError(err)
		crashOnError(catalog.Pull(args[0]))
	},
}

var catalogRemoveCmdLongHelp = `Remove an application from catalog.`

var catalogRemoveCmdShortHelp = `Remove an application from catalog.`

var removeCmd = &cobra.Command{
	Use:   "remove <[catalog/]namespace/appName[:tag]>",
	Long:  catalogRemoveCmdLongHelp,
	Short: catalogRemoveCmdShortHelp,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		catalog, err := operations.NewCatalog(&cfg)
		crashOnError(err)
		crashOnError(catalog.Remove(args[0]))
	},
}

var catalogInfoCmdLongHelp = `Get the principal information of an application.`

var catalogInfoCmdShortHelp = `Get the principal information of an application.`

var infoCmd = &cobra.Command{
	Use:   "info <[catalog/]namespace/appName[:tag]>",
	Long:  catalogInfoCmdLongHelp,
	Short: catalogInfoCmdShortHelp,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		catalog, err := operations.NewCatalog(&cfg)
		crashOnError(err)
		crashOnError(catalog.Info(args[0]))
	},
}

var catalogListCmdLongHelp = `List the applications stored in the catalog`

var catalogListCmdShortHelp = `List the applications`

var listCmd = &cobra.Command{
	Use:   "list [namespace]",
	Long:  catalogListCmdLongHelp,
	Short: catalogListCmdShortHelp,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		catalog, err := operations.NewCatalog(&cfg)
		crashOnError(err)
		targetNamespace := ""
		if len(args) == 1 {
			targetNamespace = args[0]
		}
		crashOnError(catalog.List(targetNamespace))
	},
}

var catalogSummaryCmdLongHelp = `Get te catalog summary. # Namespaces, # Applications and # Tags`

var catalogSummaryCmdShortHelp = `Get te catalog summary.`

var summaryCmd = &cobra.Command{
	Use:     "summary",
	Long:    catalogSummaryCmdLongHelp,
	Short:   catalogSummaryCmdShortHelp,
	Aliases: []string{"sum"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		catalog, err := operations.NewCatalog(&cfg)
		crashOnError(err)
		crashOnError(catalog.Summary())
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(summaryCmd)

	rootCmd.AddCommand(listCmd)
}
