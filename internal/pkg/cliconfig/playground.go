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

package cliconfig

// PlaygroundConfig with a simplified configuration structure for the playground.
type PlaygroundConfig struct {
	CurrentInstallation *string
	Installations       []*Installation
}

// Installation structure representing a target environment where the user is able to interact with using the plaground CLI.
type Installation struct {
	Name             string
	ConnectionConfig *ConnectionConfig
}

// ConnectionConfig contains the configuration elements related to the connection with the Playground API.
type ConnectionConfig struct {
	// ServerAddress with the dns/IP of the playground gRPC server.
	ServerAddress string
	// ServerPort with the port of the playground gRPC server.
	ServerPort int
	// UseTLS indicates that a TLS connection is expected with the Playground API.
	UseTLS bool
	// SkipCertValidation flag that enables ignoring the validation step of the certificate presented by the server.
	SkipCertValidation bool
	// CatalogAddress with the dns/IP of the catalog-manager gRPC server.
	CatalogAddress string
	// CatalogPort with the port of the catalog-manager gRPC server.
	CatalogPort int
	// ClientCA with a valid client CA
	ClientCA string
}

// GetSelectedConnectionConfig retrieves the selected configuration from the playground configuration.
func (pc PlaygroundConfig) GetSelectedConnectionConfig() *ConnectionConfig {
	for _, inst := range pc.Installations {
		if inst.Name == *pc.CurrentInstallation {
			return inst.ConnectionConfig
		}
	}
	return nil
}
