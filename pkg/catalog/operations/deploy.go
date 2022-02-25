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
	"github.com/napptive/catalog-cli/v2/internal/pkg/printer"
	"github.com/napptive/catalog-cli/v2/pkg/config"
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
