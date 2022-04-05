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

// NoPrinter structure with the implementation required to print an empty result.
// Sometimes we could not to print the result
type NoPrinter struct {
}

// NewNoPrinter build a new NoPrinter.
func NewNoPrinter() (ResultPrinter, error) {
	return &NoPrinter{}, nil
}

// Print the result (empty result)
func (jp *NoPrinter) Print(result interface{}) error {
	return nil
}

// PrintResultOrError prints the result using a given printer or the error.
func (jp *NoPrinter) PrintResultOrError(result interface{}, err error) error {
	return err
}

func (jp *NoPrinter) PrintResultOrErrorWithExtendedHeader(result interface{}, opEnv string, err error) error {
	return err
}
