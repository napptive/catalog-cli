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
	"encoding/json"
	"fmt"
)

// JSONPrinter structure with the implementation required to print as JSON a given result.
type JSONPrinter struct {
}

// NewJSONPrinter build a new ResultPrinter whose output is the JSON representation of the object.
func NewJSONPrinter() (ResultPrinter, error) {
	return &JSONPrinter{}, nil
}

// Print the result.
func (jp *JSONPrinter) Print(result interface{}) error {
	res, err := json.Marshal(result)
	if err == nil {
		fmt.Println(string(res))
	}
	return err
}

// PrintResultOrError prints the result using a given printer or the error.
func (jp *JSONPrinter) PrintResultOrError(result interface{}, err error) error {
	return PrintResultOrError(jp, result, err)
}

func (jp *JSONPrinter) PrintResultOrErrorWithExtendedHeader(result interface{}, opEnv string, err error) error {
	return PrintResultOrError(jp, result, err)
}
