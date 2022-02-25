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

package operations

import (
	"github.com/napptive/catalog-cli/v2/internal/pkg/connection"
	"github.com/napptive/catalog-cli/v2/internal/pkg/printer"
	"github.com/napptive/catalog-cli/v2/pkg/config"
	grpc_catalog_go "github.com/napptive/grpc-catalog-go"
	"github.com/napptive/nerrors/pkg/nerrors"
)

type Deploy struct {
	*config.AuthToken
	cfg *config.Config
	printer.ResultPrinter
}

// NewDeploy creates a new structure to faciliate the deploy operations.
func NewDeploy(cfg *config.Config) (*Deploy, error) {
	if err := cfg.IsValid(); err != nil {
		return nil, err
	}
	printer, err := printer.GetPrinter(cfg.PrinterType)
	if err != nil {
		return nil, err
	}
	return &Deploy{
		AuthToken:     config.NewAuthToken(cfg),
		cfg:           cfg,
		ResultPrinter: printer,
	}, nil
}

// Deploy triggers the deployment of the application in the selected environment.
func (d *Deploy) Deploy(applicationID string, targetEnvQualifiedName string, targetPlaygroundAPI string) error {
	// Connection
	conn, err := connection.GetConnection(&d.cfg.ConnectionConfig)
	if err != nil {
		return d.ResultPrinter.PrintResultOrError(nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			d.cfg.CatalogAddress, d.cfg.CatalogPort))
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewApplicationsClient(conn)
	ctx, cancel := d.AuthToken.GetContext()
	defer cancel()
	response, err := client.Deploy(ctx, &grpc_catalog_go.DeployApplicationRequest{
		ApplicationId:                  applicationID,
		TargetEnvironmentQualifiedName: targetEnvQualifiedName,
		TargetPlaygroundApiUrl:         targetPlaygroundAPI,
	})
	return d.ResultPrinter.PrintResultOrError(response, err)
}
