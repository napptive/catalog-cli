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

package config

import (
	"github.com/rs/zerolog/log"
)

// Config structure with all the options required by the service and service components.
type Config struct {
	ConnectionConfig
	Version string
	Commit  string
	Debug   bool

	// PrinterType defines how results are to be shown.
	PrinterType string
}

// IsValid checks if the configuration options are valid.
func (c *Config) IsValid() error {
	if err := c.ConnectionConfig.IsValid(); err != nil {
		return err
	}
	return nil
}

// Print the configuration using the application logger.
func (c *Config) Print() {
	// Use logger to print the configuration
	log.Info().Str("version", c.Version).Str("commit", c.Commit).Msg("Build information")
	log.Info().Bool("debug", c.Debug).Str("printer", c.PrinterType).Msg("internal configuration")

	c.ConnectionConfig.Print()
}
