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

package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/golang/protobuf/proto"
	apifault "github.com/polarismesh/specification/source/go/api/v1/fault_tolerance"
	apimodel "github.com/polarismesh/specification/source/go/api/v1/model"
	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"
	apitraffic "github.com/polarismesh/specification/source/go/api/v1/traffic_manage"

	httpcommon "github.com/polarismesh/polaris/apiserver/httpserver/http"
	api "github.com/polarismesh/polaris/common/api/v1"
	"github.com/polarismesh/polaris/common/utils"
)

const (
	defaultReadAccess string = "default-read"
	defaultAccess     string = "default"
)

// GetNamingConsoleAccessServer 注册管理端接口
func (h *HTTPServerV1) GetNamingConsoleAccessServer(include []string) (*restful.WebService, error) {
	consoleAccess := []string{defaultAccess}

	ws := new(restful.WebService)

	ws.Path("/naming/v1").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	// 如果为空，则开启全部接口
	if len(include) == 0 {
		include = consoleAccess
	}

	var hasDefault = false
	for _, item := range include {
		if item == defaultAccess {
			hasDefault = true
			break
		}
	}
	for _, item := range include {
		switch item {
		case defaultReadAccess:
			if !hasDefault {
				h.addDefaultReadAccess(ws)
			}
		case defaultAccess:
			h.addDefaultAccess(ws)
		default:
			log.Errorf("method %s does not exist in HTTPServerV1 console access", item)
			return nil, fmt.Errorf("method %s does not exist in HTTPServerV1 console access", item)
		}
	}
	return ws, nil
}

// addDefaultReadAccess 增加默认读接口
func (h *HTTPServerV1) addDefaultReadAccess(ws *restful.WebService) {
	// 管理端接口：只包含读接口
	ws.Route(enrichGetNamespacesApiDocs(ws.GET("/namespaces").To(h.GetNamespaces)))

	ws.Route(enrichGetNamespaceTokenApiDocs(ws.GET("/namespace/token").To(h.GetNamespaceToken)))

	ws.Route(enrichGetServicesApiDocs(ws.GET("/services").To(h.GetServices)))
	ws.Route(enrichGetServicesCountApiDocs(ws.GET("/services/count").To(h.GetServicesCount)))
	ws.Route(enrichGetServiceTokenApiDocs(ws.GET("/service/token").To(h.GetServiceToken)))
	ws.Route(enrichGetServiceAliasesApiDocs(ws.GET("/service/aliases").To(h.GetServiceAliases)))
	ws.Route(enrichGetCircuitBreakerByServiceApiDocs(ws.GET("/service/circuitbreaker").
		To(h.GetCircuitBreakerByService)))
	ws.Route(enrichGetServiceOwnerApiDocs(ws.POST("/service/owner").To(h.GetServiceOwner)))

	ws.Route(enrichGetInstancesApiDocs(ws.GET("/instances").To(h.GetInstances)))
	ws.Route(enrichGetInstancesCountApiDocs(ws.GET("/instances/count").To(h.GetInstancesCount)))

	ws.Route(enrichGetRoutingsApiDocs(ws.GET("/routings").To(h.GetRoutings)))

	ws.Route(enrichGetRateLimitsApiDocs(ws.GET("/ratelimits").To(h.GetRateLimits)))

	ws.Route(enrichGetCircuitBreakersApiDocs(ws.GET("/circuitbreaker").To(h.GetCircuitBreaker)))
	ws.Route(enrichGetCircuitBreakerVersionsApiDocs(ws.GET("/circuitbreaker/versions").
		To(h.GetCircuitBreakerVersions)))
	ws.Route(enrichGetMasterCircuitBreakersApiDocs(ws.GET("/circuitbreakers/master").
		To(h.GetMasterCircuitBreakers)))
	ws.Route(enrichGetReleaseCircuitBreakersApiDocs(ws.GET("/circuitbreakers/release").
		To(h.GetReleaseCircuitBreakers)))
	ws.Route(enrichGetCircuitBreakerTokensApiDocs(
		ws.GET("/circuitbreaker/token").To(h.GetCircuitBreakerToken)))

	ws.Route(enrichGetCircuitBreakerRulesApiDocs(ws.GET("/circuitbreaker/rules").To(h.GetCircuitBreakerRules)))

	ws.Route(enrichGetFaultDetectRulesApiDocs(ws.GET("/faultdetectors").To(h.GetFaultDetectRules)))
}

// addDefaultAccess 增加默认接口
func (h *HTTPServerV1) addDefaultAccess(ws *restful.WebService) {
	// 管理端接口：增删改查请求全部操作存储层
	ws.Route(enrichCreateNamespacesApiDocs(ws.POST("/namespaces").To(h.CreateNamespaces)))
	ws.Route(enrichDeleteNamespacesApiDocs(ws.POST("/namespaces/delete").To(h.DeleteNamespaces)))
	ws.Route(enrichUpdateNamespacesApiDocs(ws.PUT("/namespaces").To(h.UpdateNamespaces)))
	ws.Route(enrichGetNamespacesApiDocs(ws.GET("/namespaces").To(h.GetNamespaces)))
	ws.Route(enrichGetNamespaceTokenApiDocs(ws.GET("/namespace/token").To(h.GetNamespaceToken)))
	ws.Route(enrichUpdateNamespaceTokenApiDocs(ws.PUT("/namespace/token").To(h.UpdateNamespaceToken)))

	ws.Route(enrichCreateServicesApiDocs(ws.POST("/services").To(h.CreateServices)))
	ws.Route(enrichDeleteServicesApiDocs(ws.POST("/services/delete").To(h.DeleteServices)))
	ws.Route(enrichUpdateServicesApiDocs(ws.PUT("/services").To(h.UpdateServices)))
	ws.Route(enrichGetServicesApiDocs(ws.GET("/services").To(h.GetServices)))
	ws.Route(enrichGetAllServicesApiDocs(ws.GET("/services/all").To(h.GetAllServices)))
	ws.Route(enrichGetServicesCountApiDocs(ws.GET("/services/count").To(h.GetServicesCount)))
	ws.Route(enrichGetServiceTokenApiDocs(ws.GET("/service/token").To(h.GetServiceToken)))
	ws.Route(enrichUpdateServiceTokenApiDocs(ws.PUT("/service/token").To(h.UpdateServiceToken)))
	ws.Route(enrichCreateServiceAliasApiDocs(ws.POST("/service/alias").To(h.CreateServiceAlias)))
	ws.Route(enrichUpdateServiceAliasApiDocs(ws.PUT("/service/alias").To(h.UpdateServiceAlias)))
	ws.Route(enrichGetServiceAliasesApiDocs(ws.GET("/service/aliases").To(h.GetServiceAliases)))
	ws.Route(enrichDeleteServiceAliasesApiDocs(ws.POST("/service/aliases/delete").To(h.DeleteServiceAliases)))
	ws.Route(enrichGetCircuitBreakerByServiceApiDocs(ws.GET("/service/circuitbreaker").
		To(h.GetCircuitBreakerByService)))
	ws.Route(enrichGetServiceOwnerApiDocs(ws.POST("/service/owner").To(h.GetServiceOwner)))

	ws.Route(enrichCreateInstancesApiDocs(ws.POST("/instances").To(h.CreateInstances)))
	ws.Route(enrichDeleteInstancesApiDocs(ws.POST("/instances/delete").To(h.DeleteInstances)))
	ws.Route(enrichDeleteInstancesByHostApiDocs(
		ws.POST("/instances/delete/host").To(h.DeleteInstancesByHost)))
	ws.Route(enrichUpdateInstancesApiDocs(ws.PUT("/instances").To(h.UpdateInstances)))
	ws.Route(enrichUpdateInstancesIsolateApiDocs(
		ws.PUT("/instances/isolate/host").To(h.UpdateInstancesIsolate)))
	ws.Route(enrichGetInstancesApiDocs(ws.GET("/instances").To(h.GetInstances)))
	ws.Route(enrichGetInstancesCountApiDocs(ws.GET("/instances/count").To(h.GetInstancesCount)))
	ws.Route(enrichGetInstanceLabelsApiDocs(ws.GET("/instances/labels").To(h.GetInstanceLabels)))

	ws.Route(enrichCreateRoutingsApiDocs(ws.POST("/routings").To(h.CreateRoutings)))
	ws.Route(enrichDeleteRoutingsApiDocs(ws.POST("/routings/delete").To(h.DeleteRoutings)))
	ws.Route(enrichUpdateRoutingsApiDocs(ws.PUT("/routings").To(h.UpdateRoutings)))
	ws.Route(enrichGetRoutingsApiDocs(ws.GET("/routings").To(h.GetRoutings)))

	ws.Route(enrichCreateRateLimitsApiDocs(ws.POST("/ratelimits").To(h.CreateRateLimits)))
	ws.Route(enrichDeleteRateLimitsApiDocs(ws.POST("/ratelimits/delete").To(h.DeleteRateLimits)))
	ws.Route(enrichUpdateRateLimitsApiDocs(ws.PUT("/ratelimits").To(h.UpdateRateLimits)))
	ws.Route(enrichGetRateLimitsApiDocs(ws.GET("/ratelimits").To(h.GetRateLimits)))
	ws.Route(enrichEnableRateLimitsApiDocs(ws.PUT("/ratelimits/enable").To(h.EnableRateLimits)))

	ws.Route(enrichCreateCircuitBreakersApiDocs(ws.POST("/circuitbreakers").To(h.CreateCircuitBreakers)))
	ws.Route(enrichCreateCircuitBreakerVersionsApiDocs(ws.POST("/circuitbreakers/version").
		To(h.CreateCircuitBreakerVersions)))
	ws.Route(enrichDeleteCircuitBreakersApiDocs(
		ws.POST("/circuitbreakers/delete").To(h.DeleteCircuitBreakers)))
	ws.Route(enrichUpdateCircuitBreakersApiDocs(
		ws.PUT("/circuitbreakers").To(h.UpdateCircuitBreakers)))
	ws.Route(enrichReleaseCircuitBreakersApiDocs(
		ws.POST("/circuitbreakers/release").To(h.ReleaseCircuitBreakers)))
	ws.Route(enrichUnBindCircuitBreakersApiDocs(
		ws.POST("/circuitbreakers/unbind").To(h.UnBindCircuitBreakers)))
	ws.Route(enrichGetCircuitBreakersApiDocs(ws.GET("/circuitbreaker").To(h.GetCircuitBreaker)))
	ws.Route(enrichGetCircuitBreakerVersionsApiDocs(ws.GET("/circuitbreaker/versions").
		To(h.GetCircuitBreakerVersions)))
	ws.Route(enrichGetMasterCircuitBreakersApiDocs(ws.GET("/circuitbreakers/master").
		To(h.GetMasterCircuitBreakers)))
	ws.Route(enrichGetReleaseCircuitBreakersApiDocs(ws.GET("/circuitbreakers/release").
		To(h.GetReleaseCircuitBreakers)))
	ws.Route(enrichGetCircuitBreakerTokensApiDocs(
		ws.GET("/circuitbreaker/token").To(h.GetCircuitBreakerToken)))

	ws.Route(enrichGetCircuitBreakerRulesApiDocs(
		ws.GET("/circuitbreaker/rules").To(h.GetCircuitBreakerRules)))
	ws.Route(enrichCreateCircuitBreakerRulesApiDocs(
		ws.POST("/circuitbreaker/rules").To(h.CreateCircuitBreakerRules)))
	ws.Route(enrichUpdateCircuitBreakerRulesApiDocs(
		ws.PUT("/circuitbreaker/rules").To(h.UpdateCircuitBreakerRules)))
	ws.Route(enrichDeleteCircuitBreakerRulesApiDocs(
		ws.POST("/circuitbreaker/rules/delete").To(h.DeleteCircuitBreakerRules)))
	ws.Route(enrichEnableCircuitBreakerRulesApiDocs(
		ws.PUT("/circuitbreaker/rules/enable").To(h.EnableCircuitBreakerRules)))

	ws.Route(enrichGetFaultDetectRulesApiDocs(ws.GET("/faultdetectors").To(h.GetFaultDetectRules)))
	ws.Route(enrichCreateFaultDetectRulesApiDocs(ws.POST("/faultdetectors").To(h.CreateFaultDetectRules)))
	ws.Route(enrichUpdateFaultDetectRulesApiDocs(ws.PUT("/faultdetectors").To(h.UpdateFaultDetectRules)))
	ws.Route(enrichDeleteFaultDetectRulesApiDocs(
		ws.POST("/faultdetectors/delete").To(h.DeleteFaultDetectRules)))
}

// CreateNamespaces 创建命名空间
func (h *HTTPServerV1) CreateNamespaces(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var namespaces NamespaceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apimodel.Namespace{}
		namespaces = append(namespaces, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namespaceServer.CreateNamespaces(ctx, namespaces))
}

// DeleteNamespaces 删除命名空间
func (h *HTTPServerV1) DeleteNamespaces(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var namespaces NamespaceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apimodel.Namespace{}
		namespaces = append(namespaces, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namespaceServer.DeleteNamespaces(ctx, namespaces)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// UpdateNamespaces 修改命名空间
func (h *HTTPServerV1) UpdateNamespaces(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var namespaces NamespaceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apimodel.Namespace{}
		namespaces = append(namespaces, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namespaceServer.UpdateNamespaces(ctx, namespaces)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// GetNamespaces 查询命名空间
func (h *HTTPServerV1) GetNamespaces(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	ret := h.namespaceServer.GetNamespaces(handler.ParseHeaderContext(), req.Request.URL.Query())
	handler.WriteHeaderAndProto(ret)
}

// GetNamespaceToken 命名空间token的获取
func (h *HTTPServerV1) GetNamespaceToken(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	token := req.HeaderParameter("Polaris-Token")
	ctx := context.WithValue(context.Background(), utils.StringContext("polaris-token"), token)

	queryParams := httpcommon.ParseQueryParams(req)
	namespace := &apimodel.Namespace{
		Name:  utils.NewStringValue(queryParams["name"]),
		Token: utils.NewStringValue(queryParams["token"]),
	}

	ret := h.namespaceServer.GetNamespaceToken(ctx, namespace)
	handler.WriteHeaderAndProto(ret)
}

// UpdateNamespaceToken 更新命名空间的token
func (h *HTTPServerV1) UpdateNamespaceToken(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var namespace apimodel.Namespace
	ctx, err := handler.Parse(&namespace)
	if err != nil {
		handler.WriteHeaderAndProto(api.NewResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namespaceServer.UpdateNamespaceToken(ctx, &namespace)
	handler.WriteHeaderAndProto(ret)
}

// CreateServices 创建服务
func (h *HTTPServerV1) CreateServices(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var services ServiceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Service{}
		services = append(services, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.CreateServices(ctx, services))
}

// DeleteServices 删除服务
func (h *HTTPServerV1) DeleteServices(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var services ServiceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Service{}
		services = append(services, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.DeleteServices(ctx, services)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// UpdateServices 修改服务
func (h *HTTPServerV1) UpdateServices(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var services ServiceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Service{}
		services = append(services, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateServices(ctx, services)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}
	handler.WriteHeaderAndProto(ret)
}

// GetAllServices 查询服务
func (h *HTTPServerV1) GetAllServices(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ctx := handler.ParseHeaderContext()
	ret := h.namingServer.GetAllServices(ctx, queryParams)
	handler.WriteHeaderAndProto(ret)
}

// GetServices 查询服务
func (h *HTTPServerV1) GetServices(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ctx := handler.ParseHeaderContext()
	ret := h.namingServer.GetServices(ctx, queryParams)
	handler.WriteHeaderAndProto(ret)
}

// GetServicesCount 查询服务总数
func (h *HTTPServerV1) GetServicesCount(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}
	ret := h.namingServer.GetServicesCount(handler.ParseHeaderContext())
	handler.WriteHeaderAndProto(ret)
}

// GetServiceToken 获取服务token
func (h *HTTPServerV1) GetServiceToken(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}
	token := req.HeaderParameter("Polaris-Token")
	ctx := context.WithValue(context.Background(), utils.StringContext("polaris-token"), token)

	queryParams := httpcommon.ParseQueryParams(req)
	service := &apiservice.Service{
		Name:      utils.NewStringValue(queryParams["name"]),
		Namespace: utils.NewStringValue(queryParams["namespace"]),
		Token:     utils.NewStringValue(queryParams["token"]),
	}

	ret := h.namingServer.GetServiceToken(ctx, service)
	handler.WriteHeaderAndProto(ret)
}

// UpdateServiceToken 更新服务token
func (h *HTTPServerV1) UpdateServiceToken(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var service apiservice.Service
	ctx, err := handler.Parse(&service)
	if err != nil {
		handler.WriteHeaderAndProto(api.NewResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.UpdateServiceToken(ctx, &service))
}

// CreateServiceAlias service alias
func (h *HTTPServerV1) CreateServiceAlias(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var alias apiservice.ServiceAlias
	ctx, err := handler.Parse(&alias)
	if err != nil {
		handler.WriteHeaderAndProto(api.NewResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.CreateServiceAlias(ctx, &alias))
}

// UpdateServiceAlias 修改服务别名
func (h *HTTPServerV1) UpdateServiceAlias(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var alias apiservice.ServiceAlias
	ctx, err := handler.Parse(&alias)
	if err != nil {
		handler.WriteHeaderAndProto(api.NewResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateServiceAlias(ctx, &alias)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// DeleteServiceAliases 删除服务别名
func (h *HTTPServerV1) DeleteServiceAliases(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var aliases ServiceAliasArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.ServiceAlias{}
		aliases = append(aliases, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}
	ret := h.namingServer.DeleteServiceAliases(ctx, aliases)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// GetServiceAliases 根据源服务获取服务别名
func (h *HTTPServerV1) GetServiceAliases(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetServiceAliases(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// CreateInstances 创建服务实例
func (h *HTTPServerV1) CreateInstances(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var instances InstanceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Instance{}
		instances = append(instances, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.CreateInstances(ctx, instances))
}

// DeleteInstances 删除服务实例
func (h *HTTPServerV1) DeleteInstances(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var instances InstanceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Instance{}
		instances = append(instances, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.DeleteInstances(ctx, instances)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// DeleteInstancesByHost 根据host删除服务实例
func (h *HTTPServerV1) DeleteInstancesByHost(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var instances InstanceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Instance{}
		instances = append(instances, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.DeleteInstancesByHost(ctx, instances)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// UpdateInstances 修改服务实例
func (h *HTTPServerV1) UpdateInstances(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var instances InstanceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Instance{}
		instances = append(instances, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateInstances(ctx, instances)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// UpdateInstancesIsolate 修改服务实例的隔离状态
func (h *HTTPServerV1) UpdateInstancesIsolate(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var instances InstanceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Instance{}
		instances = append(instances, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateInstancesIsolate(ctx, instances)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// GetInstances 查询服务实例
func (h *HTTPServerV1) GetInstances(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetInstances(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// GetInstancesCount 查询服务实例
func (h *HTTPServerV1) GetInstancesCount(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	ret := h.namingServer.GetInstancesCount(handler.ParseHeaderContext())
	handler.WriteHeaderAndProto(ret)
}

// GetInstanceLabels 查询某个服务下所有实例的标签信息
func (h *HTTPServerV1) GetInstanceLabels(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	ret := h.namingServer.GetInstanceLabels(handler.ParseHeaderContext(), httpcommon.ParseQueryParams(req))
	handler.WriteHeaderAndProto(ret)
}

// CreateRoutings 创建规则路由
func (h *HTTPServerV1) CreateRoutings(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var routings RoutingArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apitraffic.Routing{}
		routings = append(routings, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.CreateRoutingConfigs(ctx, routings)
	handler.WriteHeaderAndProto(ret)
}

// DeleteRoutings 删除规则路由
func (h *HTTPServerV1) DeleteRoutings(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var routings RoutingArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apitraffic.Routing{}
		routings = append(routings, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.DeleteRoutingConfigs(ctx, routings)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// UpdateRoutings 修改规则路由
func (h *HTTPServerV1) UpdateRoutings(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var routings RoutingArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apitraffic.Routing{}
		routings = append(routings, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateRoutingConfigs(ctx, routings)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// GetRoutings 查询规则路由
func (h *HTTPServerV1) GetRoutings(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetRoutingConfigs(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// CreateRateLimits 创建限流规则
func (h *HTTPServerV1) CreateRateLimits(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var rateLimits RateLimitArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apitraffic.Rule{}
		rateLimits = append(rateLimits, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.CreateRateLimits(ctx, rateLimits))
}

// DeleteRateLimits 删除限流规则
func (h *HTTPServerV1) DeleteRateLimits(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var rateLimits RateLimitArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apitraffic.Rule{}
		rateLimits = append(rateLimits, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.DeleteRateLimits(ctx, rateLimits)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}
	handler.WriteHeaderAndProto(ret)
}

// EnableRateLimits 激活限流规则
func (h *HTTPServerV1) EnableRateLimits(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}
	var rateLimits RateLimitArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apitraffic.Rule{}
		rateLimits = append(rateLimits, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}
	ret := h.namingServer.EnableRateLimits(ctx, rateLimits)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// UpdateRateLimits 修改限流规则
func (h *HTTPServerV1) UpdateRateLimits(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var rateLimits RateLimitArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apitraffic.Rule{}
		rateLimits = append(rateLimits, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateRateLimits(ctx, rateLimits)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// GetRateLimits 查询限流规则
func (h *HTTPServerV1) GetRateLimits(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetRateLimits(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// CreateCircuitBreakers 创建熔断规则
func (h *HTTPServerV1) CreateCircuitBreakers(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var circuitBreakers CircuitBreakerArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.CircuitBreaker{}
		circuitBreakers = append(circuitBreakers, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.CreateCircuitBreakers(ctx, circuitBreakers)
	handler.WriteHeaderAndProto(ret)
}

// CreateCircuitBreakerVersions 创建熔断规则版本
func (h *HTTPServerV1) CreateCircuitBreakerVersions(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var circuitBreakers CircuitBreakerArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.CircuitBreaker{}
		circuitBreakers = append(circuitBreakers, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchQueryResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.CreateCircuitBreakerVersions(ctx, circuitBreakers))
}

// DeleteCircuitBreakers 删除熔断规则
func (h *HTTPServerV1) DeleteCircuitBreakers(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var circuitBreakers CircuitBreakerArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.CircuitBreaker{}
		circuitBreakers = append(circuitBreakers, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.DeleteCircuitBreakers(ctx, circuitBreakers)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}
	handler.WriteHeaderAndProto(ret)
}

// UpdateCircuitBreakers 修改熔断规则
func (h *HTTPServerV1) UpdateCircuitBreakers(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var circuitBreakers CircuitBreakerArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.CircuitBreaker{}
		circuitBreakers = append(circuitBreakers, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateCircuitBreakers(ctx, circuitBreakers)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}
	handler.WriteHeaderAndProto(ret)
}

// ReleaseCircuitBreakers 发布熔断规则
func (h *HTTPServerV1) ReleaseCircuitBreakers(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var configRelease ConfigReleaseArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.ConfigRelease{}
		configRelease = append(configRelease, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.ReleaseCircuitBreakers(ctx, configRelease)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}
	handler.WriteHeaderAndProto(ret)
}

// UnBindCircuitBreakers 解绑熔断规则
func (h *HTTPServerV1) UnBindCircuitBreakers(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var configRelease ConfigReleaseArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.ConfigRelease{}
		configRelease = append(configRelease, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UnBindCircuitBreakers(ctx, configRelease)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}
	handler.WriteHeaderAndProto(ret)
}

// GetCircuitBreaker 根据id和version获取熔断规则
func (h *HTTPServerV1) GetCircuitBreaker(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetCircuitBreaker(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// GetCircuitBreakerVersions 查询熔断规则的所有版本
func (h *HTTPServerV1) GetCircuitBreakerVersions(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetCircuitBreakerVersions(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// GetMasterCircuitBreakers 查询master熔断规则
func (h *HTTPServerV1) GetMasterCircuitBreakers(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetMasterCircuitBreakers(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// GetReleaseCircuitBreakers 根据规则id查询已发布的熔断规则
func (h *HTTPServerV1) GetReleaseCircuitBreakers(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetReleaseCircuitBreakers(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// GetCircuitBreakerByService 根据服务查询绑定熔断规则
func (h *HTTPServerV1) GetCircuitBreakerByService(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetCircuitBreakerByService(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// GetServiceOwner 根据服务获取服务负责人
func (h *HTTPServerV1) GetServiceOwner(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var services ServiceArr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apiservice.Service{}
		services = append(services, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.GetServiceOwner(ctx, services))
}

// GetCircuitBreakerToken 获取熔断规则token
func (h *HTTPServerV1) GetCircuitBreakerToken(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}
	token := req.HeaderParameter("Polaris-Token")
	ctx := context.WithValue(context.Background(), utils.StringContext("polaris-token"), token)

	queryParams := httpcommon.ParseQueryParams(req)
	circuitBreaker := &apifault.CircuitBreaker{
		Id:      utils.NewStringValue(queryParams["id"]),
		Version: utils.NewStringValue("master"),
		Token:   utils.NewStringValue(queryParams["token"]),
	}
	ret := h.namingServer.GetCircuitBreakerToken(ctx, circuitBreaker)
	handler.WriteHeaderAndProto(ret)
}

// CreateCircuitBreakerRules create the circuitbreaker rues
func (h *HTTPServerV1) CreateCircuitBreakerRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var circuitBreakerRules CircuitBreakerRuleAttr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.CircuitBreakerRule{}
		circuitBreakerRules = append(circuitBreakerRules, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.CreateCircuitBreakerRules(ctx, circuitBreakerRules))
}

// DeleteCircuitBreakerRules delete the circuitbreaker rues
func (h *HTTPServerV1) DeleteCircuitBreakerRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var circuitBreakerRules CircuitBreakerRuleAttr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.CircuitBreakerRule{}
		circuitBreakerRules = append(circuitBreakerRules, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.DeleteCircuitBreakerRules(ctx, circuitBreakerRules)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}
	handler.WriteHeaderAndProto(ret)
}

// EnableCircuitBreakerRules enable the circuitbreaker rues
func (h *HTTPServerV1) EnableCircuitBreakerRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}
	var circuitBreakerRules CircuitBreakerRuleAttr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.CircuitBreakerRule{}
		circuitBreakerRules = append(circuitBreakerRules, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}
	ret := h.namingServer.EnableCircuitBreakerRules(ctx, circuitBreakerRules)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// UpdateCircuitBreakerRules update the circuitbreaker rues
func (h *HTTPServerV1) UpdateCircuitBreakerRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var circuitBreakerRules CircuitBreakerRuleAttr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.CircuitBreakerRule{}
		circuitBreakerRules = append(circuitBreakerRules, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateCircuitBreakerRules(ctx, circuitBreakerRules)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// GetCircuitBreakerRules query the circuitbreaker rues
func (h *HTTPServerV1) GetCircuitBreakerRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetCircuitBreakerRules(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}

// CreateFaultDetectRules create the fault detect rues
func (h *HTTPServerV1) CreateFaultDetectRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var faultDetectRules FaultDetectRuleAttr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.FaultDetectRule{}
		faultDetectRules = append(faultDetectRules, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	handler.WriteHeaderAndProto(h.namingServer.CreateFaultDetectRules(ctx, faultDetectRules))
}

// DeleteFaultDetectRules delete the fault detect rues
func (h *HTTPServerV1) DeleteFaultDetectRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var faultDetectRules FaultDetectRuleAttr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.FaultDetectRule{}
		faultDetectRules = append(faultDetectRules, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.DeleteFaultDetectRules(ctx, faultDetectRules)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}
	handler.WriteHeaderAndProto(ret)
}

// UpdateFaultDetectRules update the fault detect rues
func (h *HTTPServerV1) UpdateFaultDetectRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	var faultDetectRules FaultDetectRuleAttr
	ctx, err := handler.ParseArray(func() proto.Message {
		msg := &apifault.FaultDetectRule{}
		faultDetectRules = append(faultDetectRules, msg)
		return msg
	})
	if err != nil {
		handler.WriteHeaderAndProto(api.NewBatchWriteResponseWithMsg(apimodel.Code_ParseException, err.Error()))
		return
	}

	ret := h.namingServer.UpdateFaultDetectRules(ctx, faultDetectRules)
	if code := api.CalcCode(ret); code != http.StatusOK {
		handler.WriteHeaderAndProto(ret)
		return
	}

	handler.WriteHeaderAndProto(ret)
}

// GetFaultDetectRules query the fault detect rues
func (h *HTTPServerV1) GetFaultDetectRules(req *restful.Request, rsp *restful.Response) {
	handler := &httpcommon.Handler{
		Request:  req,
		Response: rsp,
	}

	queryParams := httpcommon.ParseQueryParams(req)
	ret := h.namingServer.GetFaultDetectRules(handler.ParseHeaderContext(), queryParams)
	handler.WriteHeaderAndProto(ret)
}
