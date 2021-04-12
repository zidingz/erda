// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package orchestrator

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var ORCHESTRATOR_DEPLOYMENT_CANCEL = apis.ApiSpec{
	Path:         "/api/deployments/<deploymentId>/actions/cancel",
	BackendPath:  "/api/deployments/<deploymentId>/actions/cancel",
	Host:         "orchestrator.marathon.l4lb.thisdcos.directory:8081",
	Scheme:       "http",
	Method:       "POST",
	CheckLogin:   false, // TODO: 现在是因为 pipeline action 是从集群内走公网掉，只能放开登录认证，后面需要重构成内网调用
	RequestType:  apistructs.DeploymentCancelRequest{},
	ResponseType: apistructs.DeploymentCancelResponse{},
	Doc:          `取消部署`,
}