// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dop

import (
	"net/http"
	"net/url"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var GetAccess = apis.ApiSpec{
	Path:         "/api/api-access/<accessID>",
	BackendPath:  "/api/api-access/<accessID>",
	Host:         APIMAddr,
	Scheme:       "http",
	Method:       http.MethodGet,
	CheckLogin:   true,
	CheckToken:   true,
	RequestType:  nil,
	ResponseType: nil,
	Doc:          "get access",
	Parameters: &apis.Parameters{
		Tag:    "apim",
		Header: nil,
		QueryValues: url.Values{
			"accessID": nil,
		},
		Body:     nil,
		Response: &apistructs.GetAccessRspAccess{},
	},
}
