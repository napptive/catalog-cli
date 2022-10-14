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
	"os"
	"strings"

	"github.com/napptive/catalog-cli/v2/internal/pkg/connection"
	"github.com/napptive/catalog-cli/v2/internal/pkg/printer"
	"github.com/napptive/catalog-cli/v2/pkg/config"
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
func (c *Catalog) Push(applicationID string, path string, privateApp bool) error {
	log.Debug().Str("applicationID", applicationID).Str("path", path).Msg("Push received!")

	// Read the path and compose the AddCatalogRequest
	names, err := c.loadApp(path, ".")
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, err)
	}
	log.Debug().Interface("names", names).Msg("Files found")

	// Send the request
	// Read the paths and compose the AddCatalogRequest
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, applicationID)
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
	}
	defer conn.Close()

	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Get response and print result
	stream, err := client.Add(ctx)
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, err)
	}
	for _, fileName := range names {
		readPath := fmt.Sprintf("%s/%s", path, fileName)
		data, err := os.ReadFile(readPath)
		if err != nil {
			return c.ResultPrinter.PrintResultOrError(nil, err)
		}
		if err := stream.Send(&grpc_catalog_go.AddApplicationRequest{
			ApplicationId: applicationID,
			Private:       privateApp,
			File: &grpc_catalog_go.FileInfo{
				Path: fileName,
				Data: data,
			},
		}); err != nil {
			return c.ResultPrinter.PrintResultOrError(nil, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, err)
	}
	log.Debug().Interface("reply", reply).Msg("Application sent")

	return c.ResultPrinter.PrintResultOrError(reply, nil)
}

// Pull downloads application files
func (c *Catalog) Pull(applicationID string) error {

	// Connection
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, applicationID)
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, err)
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
		return c.ResultPrinter.PrintResultOrError(nil, err)
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
			return c.ResultPrinter.PrintResultOrError(nil, err)
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
		return c.ResultPrinter.PrintResultOrError(nil, err)
	}
	return c.ResultPrinter.PrintResultOrError(&grpc_catalog_common_go.OpResponse{
		Status:     grpc_catalog_common_go.OpStatus_SUCCESS,
		StatusName: grpc_catalog_common_go.OpStatus_SUCCESS.String(),
		UserInfo:   fmt.Sprintf("application saved on %s", files[0].Path),
	}, nil)

}

// Remove deletes an application from catalog repository
func (c *Catalog) Remove(applicationID string) error {

	// Connection
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, applicationID)
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Call Delete op
	response, err := client.Remove(ctx, &grpc_catalog_go.RemoveApplicationRequest{ApplicationId: applicationID})
	return c.ResultPrinter.PrintResultOrError(response, err)
}

// Info gets application information
func (c *Catalog) Info(application string) error {
	// Connection
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, application)
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Call Delete op
	response, err := client.Info(ctx, &grpc_catalog_go.InfoApplicationRequest{ApplicationId: application})
	return c.ResultPrinter.PrintResultOrError(response, err)
}

// List returns the applications
func (c *Catalog) List(targetNamespace string) error {
	// Connection
	// adds an empty applicationName to the targetNamespace to use GetConnectionToCatalog method
	conn, err := connection.GetConnectionToCatalog(&c.cfg.ConnectionConfig, fmt.Sprintf("%s/", targetNamespace))
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	response, err := client.List(ctx, &grpc_catalog_go.ListApplicationsRequest{
		Namespace: targetNamespace,
	})
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, nerrors.FromGRPC(err))
	}
	return c.ResultPrinter.PrintResultOrError(response, nil)
}

func (c *Catalog) Summary() error {
	// Connection
	conn, err := connection.GetConnection(&c.cfg.ConnectionConfig)
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Get Summary
	summary, err := client.Summary(ctx, &grpc_catalog_common_go.EmptyRequest{})

	return c.ResultPrinter.PrintResultOrError(summary, err)
}

func (c *Catalog) splitApplicationName(applicationName string) (string, string, error) {
	splited := strings.Split(applicationName, "/")
	if len(splited) != 2 {
		return "", "", nerrors.NewFailedPreconditionError("error in application name: <namespace>/<applicationName>")
	}
	// check if the application name has a tag
	if strings.Contains(splited[1], ":") {
		return "", "", nerrors.NewFailedPreconditionError("error in application name: <namespace>/<applicationName> without tag")
	}

	return splited[0], splited[1], nil
}

// ChangeVisibilty changes the visitibily of an application (for all tags)
func (c *Catalog) ChangeVisibilty(applicationName string, isPrivate bool, isPublic bool) error {

	// validate
	if isPrivate == isPublic {
		return c.ResultPrinter.PrintResultOrError(nil, nerrors.NewInternalError("error changing visibility, choose public or private flag"))
	}

	namespace, app, err := c.splitApplicationName(applicationName)
	if err != nil {
		return err
	}

	conn, err := connection.GetConnection(&c.cfg.ConnectionConfig)
	if err != nil {
		return c.ResultPrinter.PrintResultOrError(nil, nerrors.NewInternalErrorFrom(err, "cannot establish connection with catalog-manager server on %s:%d",
			c.cfg.CatalogAddress, c.cfg.CatalogPort))
	}
	defer conn.Close()

	// Client
	client := grpc_catalog_go.NewCatalogClient(conn)
	ctx, cancel := c.AuthToken.GetContext()
	defer cancel()

	// Get Summary
	opResponse, err := client.Update(ctx, &grpc_catalog_go.UpdateRequest{
		Namespace:       namespace,
		ApplicationName: app,
		Private:         isPrivate,
	})

	return c.ResultPrinter.PrintResultOrError(opResponse, err)
}
