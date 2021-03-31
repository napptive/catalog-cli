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
	"github.com/napptive/catalog-cli/internal/app/catalog/operations"
	"github.com/spf13/cobra"
)

var catalogPushCmdLongHelp = `Push an application in the catalog. \
The application should be named: [catalog/]repoName/appName[:tag] `

var catalogPushCmdShortHelp = `Push an application in the catalog.`

var pushCmd = &cobra.Command{
	Use:   "push <[catalog/]repoName/appName[:tag]> <application_path>",
	Long:  catalogPushCmdLongHelp,
	Short: catalogPushCmdShortHelp,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		catalog, err := operations.NewCatalog(&cfg)
		if err != nil {
			return err
		}
		return catalog.Push(args[0], args[1])
	},
}

var catalogPullCmdLongHelp = `Pull an application from catalog.`

var catalogPullCmdShortHelp = `Pull an application from catalog.`

var pullCmd = &cobra.Command{
	Use:   "pull <[catalog/]repoName/appName[:tag]>",
	Long:  catalogPullCmdLongHelp,
	Short: catalogPullCmdShortHelp,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		catalog, err := operations.NewCatalog(&cfg)
		if err != nil {
			return err
		}
		return catalog.Pull(args[0])
	},
}

var catalogRemoveCmdLongHelp = `Remove an application from catalog.`

var catalogRemoveCmdShortHelp = `Remove an application from catalog.`

var removeCmd = &cobra.Command{
	Use:   "remove <[catalog/]repoName/appName[:tag]>",
	Long:  catalogRemoveCmdLongHelp,
	Short: catalogRemoveCmdShortHelp,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error{
		catalog, err := operations.NewCatalog(&cfg)
		if err != nil {
			return err
		}
		return catalog.Remove(args[0])
	},
}

var catalogInfoCmdLongHelp = `Get the principal information of an application.`

var catalogInfoCmdShortHelp = `Get the principal information of an application.`

var infoCmd = &cobra.Command{
	Use:   "info <[catalog/]repoName/appName[:tag]>",
	Long:  catalogInfoCmdLongHelp,
	Short: catalogInfoCmdShortHelp,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error{
		catalog, err := operations.NewCatalog(&cfg)
		if err != nil {
			return err
		}
		return catalog.Info(args[0])
	},
}

var catalogListCmdLongHelp = `List the applications stored in the catalog`

var catalogListCmdShortHelp = `List the applications`

var listCmd = &cobra.Command{
	Use:   "list",
	Long:  catalogInfoCmdLongHelp,
	Short: catalogInfoCmdShortHelp,
	RunE: func(cmd *cobra.Command, args []string) error{
		catalog, err := operations.NewCatalog(&cfg)
		if err != nil {
			return err
		}
		return catalog.List()
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(listCmd)

}
