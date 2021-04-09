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
package config

import "github.com/napptive/nerrors/pkg/nerrors"

// CheckNotEmpty returns an error if the given attribute is empty.
func CheckNotEmpty(attribute string, attributeName string) error {
	if attribute == "" {
		return nerrors.NewInvalidArgumentError("%s cannot be empty", attributeName)
	}
	return nil
}

// CheckPositive returns an error if the given value is less or equal than zero.
func CheckPositive(attribute int, attributeName string) error {
	if attribute <= 0 {
		return nerrors.NewInvalidArgumentError("%s must be a positive number", attributeName)
	}
	return nil
}

