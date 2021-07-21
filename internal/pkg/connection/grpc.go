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
package connection

import (
	"crypto/tls"
	"fmt"
	"github.com/napptive/nerrors/pkg/nerrors"
	"google.golang.org/grpc/credentials"
	"strings"

	"github.com/napptive/catalog-cli/pkg/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// GetConnectionToCatalog creates
func GetConnectionToCatalog(cfg *config.ConnectionConfig, applicationID string) (*grpc.ClientConn, error) {
	// Get CatalogURl from ApplicationID (this method validates the applicationID format too)
	catalogURL, err := GetURL(cfg, applicationID)
	if err != nil {
		return nil, err
	}
	if cfg.UseTLS {
		return GetTLSConnection(cfg, catalogURL)
	}
	return GetNonTLSConnection(cfg, catalogURL)

}

// GetTLSConnection returns a TLS wrapped connection with the playground server.
func GetTLSConnection(cfg *config.ConnectionConfig, address string) (*grpc.ClientConn, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: cfg.SkipCertValidation,
	}
	tlsCredentials := credentials.NewTLS(tlsConfig)
	return grpc.Dial(address, grpc.WithTransportCredentials(tlsCredentials))
}

// GetConnection creates a connection with a gRPC server.
func GetConnection(cfg *config.ConnectionConfig) (*grpc.ClientConn, error) {
	addr := cfg.GetEffectiveAddress()
	if cfg.UseTLS {
		return GetTLSConnection(cfg, addr)
	}
	return GetNonTLSConnection(cfg, addr)
}

// GetNonTLSConnection returns a plain connection with the playground server.
func GetNonTLSConnection(cfg *config.ConnectionConfig, address string) (*grpc.ClientConn, error) {
	log.Debug().Msg("using insecure connection with the Catalog-Manager")
	return grpc.Dial(address, grpc.WithInsecure())
}

// GetURL returns CatalogURL from [catalogURL/]repoName/applicationName[:version]
func GetURL(cfg *config.ConnectionConfig, appName string) (string, error) {

	names := strings.Split(appName, "/")
	if len(names) != 2 && len(names) != 3 {
		return "", nerrors.NewFailedPreconditionError(
			"Incorrect format for application name. It should be [catalogURL/]namespace/appName[:tag]")
	}

	if len(names) == 3 {
		// Check if the URL has the port, if not -> append cfg.CatalogPort to the URL
		url := names[0]
		if ! strings.Contains(url, ":") {
			url = fmt.Sprintf("%s:%d", url, cfg.CatalogPort)
		}
		return url, nil
	}
	// if len == 2 -> no url informed, default
	return fmt.Sprintf("%s:%d", cfg.CatalogAddress, cfg.CatalogPort), nil
}
