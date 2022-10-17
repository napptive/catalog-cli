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

package printer

import (
	"reflect"

	grpc_catalog_common_go "github.com/napptive/grpc-catalog-common-go"
	grpc_catalog_go "github.com/napptive/grpc-catalog-go"
	"github.com/napptive/nerrors/pkg/nerrors"
)

// OpResponseTemplate with the table representation of an OpResponse.
const OpResponseTemplate = `STATUS	INFO
{{.StatusName}}	{{.UserInfo}}
`

// InfoAppResponseTemplate with the table representation of an InfoAppResponse.
const InfoAppResponseTemplate = `APP_ID	NAME
{{.Namespace}}/{{.ApplicationName}}:{{.Tag}}	{{.Metadata.Name}}	

DESCRIPTION
{{.Metadata.Description}}
{{if .Metadata.Requires.Traits}}TRAITS
{{range $name := .Metadata.Requires.Traits}}{{$name}}
{{end}}{{end}}{{if .Metadata.Requires.Scopes}}SCOPES
{{range $name := .Metadata.Requires.Scopes}}{{$name}}
{{end}}{{end}}{{if .Metadata.Requires.K8S}}K8S_ENTITIES
{{range .Metadata.Requires.K8S}}{{.ApiVersion}}/{{.Kind}}
{{end}}{{end}}
README
{{toString .ReadmeFile}}
`

const ApplicationListTemplate = `APPLICATION	NAME	VISIBILITY
{{range $other, $app := .Applications}}{{fromApplicationSummary $app}}{{end}}`

const SummaryResponseTemplate = `NAMESPACES	APPLICATIONS	TAGS
{{.NumNamespaces}}	{{.NumApplications}}	{{.NumTags}}
`

// structTemplates map associating type and template to print it.
var structTemplates = map[reflect.Type]string{
	reflect.TypeOf(&grpc_catalog_common_go.OpResponse{}):       OpResponseTemplate,
	reflect.TypeOf(&grpc_catalog_go.InfoApplicationResponse{}): InfoAppResponseTemplate,
	reflect.TypeOf(&grpc_catalog_go.ApplicationList{}):         ApplicationListTemplate,
	reflect.TypeOf(&grpc_catalog_go.SummaryResponse{}):         SummaryResponseTemplate,
	//
}

// GetTemplate returns a template to print an arbitrary structure in table format.
func GetTemplate(result interface{}) (*string, error) {
	template, exists := structTemplates[reflect.TypeOf(result)]
	if !exists {
		return nil, nerrors.NewUnimplementedError("no template is available to print %s", reflect.TypeOf(result).String())
	}
	return &template, nil
}
