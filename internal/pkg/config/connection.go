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

import "github.com/rs/zerolog/log"

// ConnectionConfig contains the configuration elements related to the connection with the Catalog-Manager API.
type ConnectionConfig struct {
	// ServerAddress with the dns/IP of the catalog-manager gRPC server.
	ServerAddress string
	// ServerPort with the port of the catalog-manager gRPC server.
	CatalogPort int

}

// IsValid checks if the configuration options are valid.
func (cc *ConnectionConfig) IsValid() error {
	if err := CheckNotEmpty(cc.ServerAddress, "ServerAddress"); err != nil {
		return err
	}
	if err := CheckPositive(cc.ServerPort, "ServerPort"); err != nil {
		return err
	}

	return nil
}

// Print the configuration using the application logger.
func (cc *ConnectionConfig) Print() {
	log.Info().Str("server", cc.ServerAddress).Int("Port", cc.ServerPort).Msg("Connection options")
}
