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
package operations

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/napptive/catalog-cli/internal/pkg/connection"
	"github.com/napptive/catalog-cli/internal/pkg/printer"
	"github.com/napptive/catalog-cli/pkg/config"
	grpc_catalog_common_go "github.com/napptive/grpc-catalog-common-go"
	grpc_catalog_go "github.com/napptive/grpc-catalog-go"
	"github.com/napptive/nerrors/pkg/nerrors"
	"github.com/rs/zerolog/log"
)

type Catalog struct {
	*config.AuthToken
	cfg *config.Config
	printer.ResultPrinter
}

func NewCatalog(cfg *config.Config) (*Catalog, error) {
	if err := cfg.IsValid(); err != nil {
		return nil, err
	}
	printer, err := printer.GetPrinter(cfg.PrinterType)
	if err != nil {
		return nil, err
	}
	return &Catalog{
		AuthToken:     config.NewAuthToken(cfg),
		cfg:           cfg,
		ResultPrinter: printer,
	}, nil
}

// loadApp reads the application directory getting all the files and their paths
func (c *Catalog) loadApp(path string, relativePath string) ([]string, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer dir.Close()

	var result []string
	directories, err := dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	for _, dirName := range directories {
		newPath := fmt.Sprintf("%s/%s", path, dirName)
		file, err := os.Stat(newPath)
		if err != nil {
			return nil, err
		}
		if file.IsDir() {
			res, nErr := c.loadApp(newPath, fmt.Sprintf("%s/%s", relativePath, dirName))
			if nErr != nil {
				return nil, nErr
			}
			result = append(result, res...)

		} else {
			result = append(result, fmt.Sprintf("%s/%s", relativePath, dirName))
		}
	}

	return result, nil
}

// Push adds a new application to catalog
func (c *Catalog) Push(applicationID string, path string) error {
	log.Debug().Str("applicationID", applicationID).Str("path", path).Msg("Push received!")

	// Read the path and compose the AddCatalogRequest
	names, err := c.loadApp(path, ".")
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, err)
		return nil
	}
	log.Debug().Interface("names", names).Msg("Files found")

	// Send the request
	// Read the paths and compose the AddCatalogRequest
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, applicationID)
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
		return nil
	}
	defer conn.Close()

	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Get response and print result
	stream, err := client.Add(ctx)
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, err)
		return nil
	}
	for _, fileName := range names {
		readPath := fmt.Sprintf("%s/%s", path, fileName)
		data, err := ioutil.ReadFile(readPath)
		if err != nil {
			PrintResultOrError(c.ResultPrinter, nil, err)
			return nil
		}
		if err := stream.Send(&grpc_catalog_go.AddApplicationRequest{
			ApplicationId: applicationID,
			File: &grpc_catalog_go.FileInfo{
				Path: fileName,
				Data: data,
			},
		}); err != nil {
			PrintResultOrError(c.ResultPrinter, nil, err)
			return nil
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, err)
		return nil
	} else {
		PrintResultOrError(c.ResultPrinter, reply, nil)
	}
	log.Debug().Interface("reply", reply).Msg("Application sent")
	return nil
}

// Pull downloads application files
func (c *Catalog) Pull(applicationID string) error {

	// Connection
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, applicationID)
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
		return nil
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Call Download
	downClient, err := client.Download(ctx, &grpc_catalog_go.DownloadApplicationRequest{
		ApplicationId: applicationID, Compressed: true,
	})
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, err)
		return nil
	}

	// Receive data
	var files []*grpc_catalog_go.FileInfo
	for {
		fileReceived, err := downClient.Recv()
		if err == io.EOF {
			_ = downClient.CloseSend()
			break
		}
		if err != nil {
			PrintResultOrError(c.ResultPrinter, nil, err)
			return nil
		}
		files = append(files, fileReceived)
	}

	// Get the application name
	_, _, appName, _, err := DecomposeApplicationName(applicationID)
	if err != nil {
		// It should not fail, in this case, the file will be named "application.tgz
		appName = "application"
	}

	// Save the files in a tgz file
	err = SaveFile(appName, files[0])
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, err)
		return nil
	}
	PrintResultOrError(c.ResultPrinter, &grpc_catalog_common_go.OpResponse{
		Status:     grpc_catalog_common_go.OpStatus_SUCCESS,
		StatusName: grpc_catalog_common_go.OpStatus_SUCCESS.String(),
		UserInfo:   fmt.Sprintf("application saved on %s", files[0].Path),
	}, nil)

	return nil
}

// Remove deletes an application from catalog repository
func (c *Catalog) Remove(applicationID string) error {

	// Connection
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, applicationID)
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
		return nil
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Call Delete op
	response, err := client.Remove(ctx, &grpc_catalog_go.RemoveApplicationRequest{ApplicationId: applicationID})
	PrintResultOrError(c.ResultPrinter, response, err)
	return nil
}

// Info gets application information
func (c *Catalog) Info(application string) error {
	// Connection
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, application)
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
		return nil
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Call Delete op
	response, err := client.Info(ctx, &grpc_catalog_go.InfoApplicationRequest{ApplicationId: application})
	PrintResultOrError(c.ResultPrinter, response, err)
	return nil
}

// List returns the applications
func (c *Catalog) List(targetNamespace string) error {
	// Connection
	conn, err := connection.GetConnection(&c.cfg.ConnectionConfig)
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
		return nil
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	response, err := client.List(ctx, &grpc_catalog_go.ListApplicationsRequest{
		Namespace: targetNamespace,
	})
	PrintResultOrError(c.ResultPrinter, response, err)
	return nil
}

func (c *Catalog) Summary() error {
	// Connection
	conn, err := connection.GetConnection(&c.cfg.ConnectionConfig)
	if err != nil {
		PrintResultOrError(c.ResultPrinter, nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
		return nil
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Get Summary
	summary, err := client.Summary(ctx, &grpc_catalog_common_go.EmptyRequest{})
	PrintResultOrError(c.ResultPrinter, summary, err)

	return nil
}