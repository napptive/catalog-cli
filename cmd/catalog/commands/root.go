/**
 * Copyright 2020 Napptive
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
	"fmt"
	"os"
	"os/user"

	"github.com/napptive/catalog-cli/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg config.Config

// tokenViper in charge of reading the JWT token returned on a login operation.
var tokenViper = viper.New()

var DefaultConfigLocation = []string{"."}

var debugLevel bool
var consoleLogging bool

var rootCmdLongHelp = "The catalog command provides a set of methods to interact with the Napptive Catalog"
var rootCmdShortHelp = "Catalog command"
var rootCmdExample = `$ catalog`
var rootCmdUse = "catalog"

var rootCmd = &cobra.Command{
	Use:     rootCmdUse,
	Example: rootCmdExample,
	Short:   rootCmdShortHelp,
	Long:    rootCmdLongHelp,
	Version: "NaN",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&debugLevel, "debug", false, "Set debug level")
	rootCmd.PersistentFlags().BoolVar(&consoleLogging, "consoleLogging", false, "Pretty print logging")

	// noPrinter allowed yet, but this value is only for internal use
	rootCmd.PersistentFlags().StringVar(&cfg.PrinterType, "output", "table", "Output format in which the results will be returned: json or table")

	rootCmd.PersistentFlags().StringVar(&cfg.CatalogAddress, "catalogAddress", "catalog-manager", "Catalog-manager host")
	rootCmd.PersistentFlags().IntVar(&cfg.CatalogPort, "catalogPort", 7060, "Catalog-manager port")
	rootCmd.PersistentFlags().BoolVar(&cfg.AuthEnable, "authEnable", true, "JWT authentication enable")

}

// Execute the user command
func Execute(version string, commit string) {
	versionTemplate := fmt.Sprintf("%s [%s] ", version, commit)
	rootCmd.SetVersionTemplate(versionTemplate)
	cfg.Version = version
	cfg.Commit = commit

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	setupLogging()
	if cfg.AuthEnable {
		readConfiguration()
	}
}

// setupLogging sets the debug level and console logging if required.
func setupLogging() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debugLevel {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if consoleLogging {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}

// getConfigLocations resolves the location of platform dependend directories such as the user home.
func getConfigLocations() []string {
	result := DefaultConfigLocation
	usr, err := user.Current()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to determine user home")
	}
	result = append(result, fmt.Sprintf("%s/.napptive", usr.HomeDir))
	return result
}

func readConfiguration() {

	// token configuration
	tokenViper.SetConfigName(".token")
	tokenViper.SetConfigType("yaml")

	for _, location := range getConfigLocations() {
		tokenViper.AddConfigPath(location)
	}
	// Load the token information if any
	if err := tokenViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal().Err(err).Msg("unable to read token file")
		} else {
			log.Debug().Msg("CLI is not authenticated")
		}
	} else {
		log.Debug().Str("path", tokenViper.ConfigFileUsed()).Msg("token loaded")
	}

	// Unmarshal resulting values in the CLI configuration.
	if err := tokenViper.Unmarshal(&cfg.TokenConfig); err != nil {
		log.Fatal().Err(err).Msg("unable to unmarshal resolved token into config structure. Check structure/file structure for a mismatch")
	}

}
