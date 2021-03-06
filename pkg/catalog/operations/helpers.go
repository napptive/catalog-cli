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
	"archive/tar"
	"compress/gzip"
	"fmt"
	"os"
	"strings"

	grpc_catalog_go "github.com/napptive/grpc-catalog-go"
	"github.com/napptive/nerrors/pkg/nerrors"
	"github.com/rs/zerolog/log"
)
/*
// PrintResultOrError prints the result using a given printer or the error.
func PrintResultOrError(printer printer.ResultPrinter, result interface{}, err error) {
	if err != nil {
		if zerolog.GlobalLevel() == zerolog.DebugLevel {
			fmt.Println(nerrors.FromError(err).StackTraceToString())
		} else {
			fmt.Println(err.Error())
		}
	} else {
		if pErr := printer.Print(result); pErr != nil {
			if zerolog.GlobalLevel() == zerolog.DebugLevel {
				fmt.Println(nerrors.FromError(pErr).StackTraceToString())
			} else {
				fmt.Println(pErr.Error())
			}
		}
	}
}
*/


func SaveFile (resultFile string, file *grpc_catalog_go.FileInfo) error {
	// Create output file
	out, err := os.Create(fmt.Sprintf(file.Path))
	if err != nil {
		return nerrors.NewInternalErrorFrom(err, "Error creating file")
	}
	defer out.Close()

	if _, err = out.Write(file.Data); err != nil {
		return nerrors.NewInternalErrorFrom(err, "Error writing file")
	}

	return nil
}

// SaveAndCompressFiles save the all the application files in a tgz file
func SaveAndCompressFiles(resultFile string, files []*grpc_catalog_go.FileInfo) error {

	// Create output file
	out, err := os.Create(fmt.Sprintf("%s.tgz", resultFile))
	if err != nil {
		return nerrors.FromError(err)
	}
	defer out.Close()

	// Create new Writers for gzip and tar
	// These writers are chained. Writing to the tar writer will
	// write to the gzip writer which in turn will write to
	// the "buf" writer
	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Iterate over files and add them to the tar archive
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Path,
			Mode: 0600,
			Size: int64(len(file.Data)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return nerrors.FromError(err)
		}
		if _, err := tw.Write([]byte(file.Data)); err != nil {
			return nerrors.FromError(err)
		}
	}
	log.Debug().Msg("Archive created successfully")
	return nil
}

// DecomposeApplicationName decompose [catalogURL/]repoName/applicationName[:version] to
// catalogURL, repoName, applicationName and version
func DecomposeApplicationName(appName string) (string, string, string, string, error) {
	applicationName := ""
	repoName := ""
	catalogName := ""
	version := "latest"

	names := strings.Split(appName, "/")
	if len(names) != 2 && len(names) != 3 {
		return "", "", "", "", nerrors.NewFailedPreconditionError(
			"incorrect format for application name. [catalogURL/]namespace/appName[:tag]")
	}

	// if len == 2 -> no url informed.
	if len(names) == 3 {
		catalogName = names[0]
	}
	repoName = names[len(names)-2]

	// get the version -> appName[:tag]
	sp := strings.Split(names[len(names)-1], ":")
	if len(sp) == 1 {
		applicationName = sp[0]
	} else if len(sp) == 2 {
		applicationName = sp[0]
		version = sp[1]
	} else {
		return "", "", "", "", nerrors.NewFailedPreconditionError(
			"incorrect format for application name. [catalogURL/]namespace/appName[:tag]")
	}

	return catalogName, repoName, applicationName, version, nil
}
