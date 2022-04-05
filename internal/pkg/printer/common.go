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

package printer

import (
	"fmt"

	"github.com/napptive/nerrors/pkg/nerrors"
	"github.com/rs/zerolog"
)

// ResultPrinter defines the operations that a printer must define. Multiple printer are
// offered depending on the desired output format.
type ResultPrinter interface {
	// Print the result.
	Print(result interface{}) error
	// PrintResultOrError prints the result using a given printer or the error.
	PrintResultOrError(result interface{}, err error) error
	// PrintResultOrErrorWithExtendedHeader prints the result indicating the environment where the user is logged
	// and the environment where the user is operating
	PrintResultOrErrorWithExtendedHeader(result interface{}, opEnv string, err error) error
}

// GetPrinter creates a ResultPrinter attending to the user preferences.
func GetPrinter(printerType string) (ResultPrinter, error) {
	switch printerType {
	case "json":
		return NewJSONPrinter()
	case "table":
		return NewTablePrinter()
	case "noPrinter":
		return NewNoPrinter()
	}
	return nil, nerrors.NewUnavailableError("printer type not supported: [%s]", printerType)
}

// PrintResultOrError prints the result using a given printer or the error.
func PrintResultOrError(printer ResultPrinter, result interface{}, err error) error {
	if err != nil {
		return err
	}
	return printer.Print(result)
}

func PrintError(err error) {
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		fmt.Println(nerrors.FromError(err).StackTraceToString())
	} else {
		fmt.Println(err.Error())
	}
}
