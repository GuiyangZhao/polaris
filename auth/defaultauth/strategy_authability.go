/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package defaultauth

import (
	"context"

	apimodel "github.com/polarismesh/specification/source/go/api/v1/model"
	apisecurity "github.com/polarismesh/specification/source/go/api/v1/security"
	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"

	api "github.com/polarismesh/polaris/common/api/v1"
)

// CreateStrategy creates a new strategy.
func (svr *serverAuthAbility) CreateStrategy(
	ctx context.Context, strategy *apisecurity.AuthStrategy) *apiservice.Response {
	ctx, rsp := svr.verifyAuth(ctx, WriteOp, MustOwner)
	if rsp != nil {
		rsp.AuthStrategy = strategy
		return rsp
	}

	return svr.target.CreateStrategy(ctx, strategy)
}

// UpdateStrategies update a strategy.
func (svr *serverAuthAbility) UpdateStrategies(ctx context.Context,
	reqs []*apisecurity.ModifyAuthStrategy) *apiservice.BatchWriteResponse {
	ctx, rsp := svr.verifyAuth(ctx, WriteOp, MustOwner)
	if rsp != nil {
		resp := api.NewAuthBatchWriteResponse(apimodel.Code_ExecuteSuccess)
		api.Collect(resp, rsp)
		return resp
	}

	return svr.target.UpdateStrategies(ctx, reqs)
}

// DeleteStrategies delete strategy.
func (svr *serverAuthAbility) DeleteStrategies(ctx context.Context,
	reqs []*apisecurity.AuthStrategy) *apiservice.BatchWriteResponse {
	ctx, rsp := svr.verifyAuth(ctx, WriteOp, MustOwner)
	if rsp != nil {
		resp := api.NewAuthBatchWriteResponse(apimodel.Code_ExecuteSuccess)
		api.Collect(resp, rsp)
		return resp
	}

	return svr.target.DeleteStrategies(ctx, reqs)
}

// GetStrategies get strategy list .
func (svr *serverAuthAbility) GetStrategies(ctx context.Context,
	query map[string]string) *apiservice.BatchQueryResponse {
	ctx, rsp := svr.verifyAuth(ctx, ReadOp, NotOwner)
	if rsp != nil {
		return api.NewAuthBatchQueryResponseWithMsg(apimodel.Code(rsp.GetCode().Value), rsp.Info.Value)
	}

	return svr.target.GetStrategies(ctx, query)
}

// GetStrategy get strategy.
func (svr *serverAuthAbility) GetStrategy(
	ctx context.Context, strategy *apisecurity.AuthStrategy) *apiservice.Response {
	ctx, rsp := svr.verifyAuth(ctx, ReadOp, NotOwner)
	if rsp != nil {
		return rsp
	}

	return svr.target.GetStrategy(ctx, strategy)
}

// GetPrincipalResources get principal resources.
func (svr *serverAuthAbility) GetPrincipalResources(ctx context.Context, query map[string]string) *apiservice.Response {
	ctx, rsp := svr.verifyAuth(ctx, ReadOp, NotOwner)
	if rsp != nil {
		return rsp
	}

	return svr.target.GetPrincipalResources(ctx, query)
}
