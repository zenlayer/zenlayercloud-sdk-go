/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package aigw

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2026-04-14"
	SERVICE    = "aigw"
)

type Client struct {
	common.Client
}

func NewClientWithSecretKey(secretKeyId, secretKeyPassword string) (client *Client, err error) {
	return NewClient(common.NewConfig(), secretKeyId, secretKeyPassword)
}

func NewClient(config *common.Config, secretKeyId, secretKeyPassword string) (client *Client, err error) {
	client = &Client{}

	err = client.InitWithCredential(common.NewCredential(secretKeyId, secretKeyPassword))
	if err != nil {
		return nil, err
	}
	err = client.WithConfig(config)

	if err != nil {
		return nil, err
	}
	return client, nil
}
func NewClientWithToken(token string) (client *Client, err error) {
	return NewClientWithTokenAndConfig(common.NewConfig(), token)
}

func NewClientWithTokenAndConfig(config *common.Config, token string) (client *Client, err error) {
	client = &Client{}

	err = client.InitWithTokenCredential(common.NewTokenCredential(token))
	if err != nil {
		return nil, err
	}
	err = client.WithConfig(config)

	if err != nil {
		return nil, err
	}
	return client, nil
}



func NewDescribeAiGatewaysRequest() (request *DescribeAiGatewaysRequest) {
	request = &DescribeAiGatewaysRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiGateways")

	return
}

func NewDescribeAiGatewaysResponse() (response *DescribeAiGatewaysResponse) {
	response = &DescribeAiGatewaysResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiGateways 分页查询ai网关列表
func (c *Client) DescribeAiGateways(request *DescribeAiGatewaysRequest) (response *DescribeAiGatewaysResponse, err error) {
	response = NewDescribeAiGatewaysResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateAiGatewayRequest() (request *CreateAiGatewayRequest) {
	request = &CreateAiGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateAiGateway")

	return
}

func NewCreateAiGatewayResponse() (response *CreateAiGatewayResponse) {
	response = &CreateAiGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateAiGateway 创建AI网关
func (c *Client) CreateAiGateway(request *CreateAiGatewayRequest) (response *CreateAiGatewayResponse, err error) {
	response = NewCreateAiGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStartAiGatewayRequest() (request *StartAiGatewayRequest) {
	request = &StartAiGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StartAiGateway")

	return
}

func NewStartAiGatewayResponse() (response *StartAiGatewayResponse) {
	response = &StartAiGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// StartAiGateway 启动AI网关
func (c *Client) StartAiGateway(request *StartAiGatewayRequest) (response *StartAiGatewayResponse, err error) {
	response = NewStartAiGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStopAiGatewayRequest() (request *StopAiGatewayRequest) {
	request = &StopAiGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StopAiGateway")

	return
}

func NewStopAiGatewayResponse() (response *StopAiGatewayResponse) {
	response = &StopAiGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// StopAiGateway 停止AI网关
func (c *Client) StopAiGateway(request *StopAiGatewayRequest) (response *StopAiGatewayResponse, err error) {
	response = NewStopAiGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteAiGatewayRequest() (request *DeleteAiGatewayRequest) {
	request = &DeleteAiGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteAiGateway")

	return
}

func NewDeleteAiGatewayResponse() (response *DeleteAiGatewayResponse) {
	response = &DeleteAiGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteAiGateway 删除AI网关
func (c *Client) DeleteAiGateway(request *DeleteAiGatewayRequest) (response *DeleteAiGatewayResponse, err error) {
	response = NewDeleteAiGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiGatewayModelsRequest() (request *DescribeAiGatewayModelsRequest) {
	request = &DescribeAiGatewayModelsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiGatewayModels")

	return
}

func NewDescribeAiGatewayModelsResponse() (response *DescribeAiGatewayModelsResponse) {
	response = &DescribeAiGatewayModelsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiGatewayModels 查询AI网关模型
func (c *Client) DescribeAiGatewayModels(request *DescribeAiGatewayModelsRequest) (response *DescribeAiGatewayModelsResponse, err error) {
	response = NewDescribeAiGatewayModelsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAiGatewayModelsRequest() (request *ModifyAiGatewayModelsRequest) {
	request = &ModifyAiGatewayModelsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAiGatewayModels")

	return
}

func NewModifyAiGatewayModelsResponse() (response *ModifyAiGatewayModelsResponse) {
	response = &ModifyAiGatewayModelsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyAiGatewayModels 修改AI网关模型
func (c *Client) ModifyAiGatewayModels(request *ModifyAiGatewayModelsRequest) (response *ModifyAiGatewayModelsResponse, err error) {
	response = NewModifyAiGatewayModelsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiGatewayIpAclRequest() (request *DescribeAiGatewayIpAclRequest) {
	request = &DescribeAiGatewayIpAclRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiGatewayIpAcl")

	return
}

func NewDescribeAiGatewayIpAclResponse() (response *DescribeAiGatewayIpAclResponse) {
	response = &DescribeAiGatewayIpAclResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiGatewayIpAcl 查询AI网关IP访问控制
func (c *Client) DescribeAiGatewayIpAcl(request *DescribeAiGatewayIpAclRequest) (response *DescribeAiGatewayIpAclResponse, err error) {
	response = NewDescribeAiGatewayIpAclResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAiGatewayIpAclRequest() (request *ModifyAiGatewayIpAclRequest) {
	request = &ModifyAiGatewayIpAclRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAiGatewayIpAcl")

	return
}

func NewModifyAiGatewayIpAclResponse() (response *ModifyAiGatewayIpAclResponse) {
	response = &ModifyAiGatewayIpAclResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyAiGatewayIpAcl 修改AI网关IP访问控制
func (c *Client) ModifyAiGatewayIpAcl(request *ModifyAiGatewayIpAclRequest) (response *ModifyAiGatewayIpAclResponse, err error) {
	response = NewModifyAiGatewayIpAclResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiGatewayExpireTimeRequest() (request *DescribeAiGatewayExpireTimeRequest) {
	request = &DescribeAiGatewayExpireTimeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiGatewayExpireTime")

	return
}

func NewDescribeAiGatewayExpireTimeResponse() (response *DescribeAiGatewayExpireTimeResponse) {
	response = &DescribeAiGatewayExpireTimeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiGatewayExpireTime 查询AI网关过期时间
func (c *Client) DescribeAiGatewayExpireTime(request *DescribeAiGatewayExpireTimeRequest) (response *DescribeAiGatewayExpireTimeResponse, err error) {
	response = NewDescribeAiGatewayExpireTimeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAiGatewayExpireTimeRequest() (request *ModifyAiGatewayExpireTimeRequest) {
	request = &ModifyAiGatewayExpireTimeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAiGatewayExpireTime")

	return
}

func NewModifyAiGatewayExpireTimeResponse() (response *ModifyAiGatewayExpireTimeResponse) {
	response = &ModifyAiGatewayExpireTimeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyAiGatewayExpireTime 修改AI网关过期时间
func (c *Client) ModifyAiGatewayExpireTime(request *ModifyAiGatewayExpireTimeRequest) (response *ModifyAiGatewayExpireTimeResponse, err error) {
	response = NewModifyAiGatewayExpireTimeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiGatewayTokenLimitRequest() (request *DescribeAiGatewayTokenLimitRequest) {
	request = &DescribeAiGatewayTokenLimitRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiGatewayTokenLimit")

	return
}

func NewDescribeAiGatewayTokenLimitResponse() (response *DescribeAiGatewayTokenLimitResponse) {
	response = &DescribeAiGatewayTokenLimitResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiGatewayTokenLimit 查询AI网关Token限制
func (c *Client) DescribeAiGatewayTokenLimit(request *DescribeAiGatewayTokenLimitRequest) (response *DescribeAiGatewayTokenLimitResponse, err error) {
	response = NewDescribeAiGatewayTokenLimitResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAiGatewayTokenLimitRequest() (request *ModifyAiGatewayTokenLimitRequest) {
	request = &ModifyAiGatewayTokenLimitRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAiGatewayTokenLimit")

	return
}

func NewModifyAiGatewayTokenLimitResponse() (response *ModifyAiGatewayTokenLimitResponse) {
	response = &ModifyAiGatewayTokenLimitResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyAiGatewayTokenLimit 修改AI网关Token限制
func (c *Client) ModifyAiGatewayTokenLimit(request *ModifyAiGatewayTokenLimitRequest) (response *ModifyAiGatewayTokenLimitResponse, err error) {
	response = NewModifyAiGatewayTokenLimitResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAiGatewayNameRequest() (request *ModifyAiGatewayNameRequest) {
	request = &ModifyAiGatewayNameRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAiGatewayName")

	return
}

func NewModifyAiGatewayNameResponse() (response *ModifyAiGatewayNameResponse) {
	response = &ModifyAiGatewayNameResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyAiGatewayName 修改AI网关名称
func (c *Client) ModifyAiGatewayName(request *ModifyAiGatewayNameRequest) (response *ModifyAiGatewayNameResponse, err error) {
	response = NewModifyAiGatewayNameResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiModelsRequest() (request *DescribeAiModelsRequest) {
	request = &DescribeAiModelsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiModels")

	return
}

func NewDescribeAiModelsResponse() (response *DescribeAiModelsResponse) {
	response = &DescribeAiModelsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiModels 查询ai模型列表
func (c *Client) DescribeAiModels(request *DescribeAiModelsRequest) (response *DescribeAiModelsResponse, err error) {
	response = NewDescribeAiModelsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiUsageDetailDataRequest() (request *DescribeAiUsageDetailDataRequest) {
	request = &DescribeAiUsageDetailDataRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiUsageDetailData")

	return
}

func NewDescribeAiUsageDetailDataResponse() (response *DescribeAiUsageDetailDataResponse) {
	response = &DescribeAiUsageDetailDataResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiUsageDetailData 查询AI网关用量明细数据
func (c *Client) DescribeAiUsageDetailData(request *DescribeAiUsageDetailDataRequest) (response *DescribeAiUsageDetailDataResponse, err error) {
	response = NewDescribeAiUsageDetailDataResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiMonthlyTotalCostRequest() (request *DescribeAiMonthlyTotalCostRequest) {
	request = &DescribeAiMonthlyTotalCostRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiMonthlyTotalCost")

	return
}

func NewDescribeAiMonthlyTotalCostResponse() (response *DescribeAiMonthlyTotalCostResponse) {
	response = &DescribeAiMonthlyTotalCostResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiMonthlyTotalCost 查询AI网关月总费用
func (c *Client) DescribeAiMonthlyTotalCost(request *DescribeAiMonthlyTotalCostRequest) (response *DescribeAiMonthlyTotalCostResponse, err error) {
	response = NewDescribeAiMonthlyTotalCostResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiModelDailyCostRequest() (request *DescribeAiModelDailyCostRequest) {
	request = &DescribeAiModelDailyCostRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiModelDailyCost")

	return
}

func NewDescribeAiModelDailyCostResponse() (response *DescribeAiModelDailyCostResponse) {
	response = &DescribeAiModelDailyCostResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiModelDailyCost 查询AI网关月模型日费用
func (c *Client) DescribeAiModelDailyCost(request *DescribeAiModelDailyCostRequest) (response *DescribeAiModelDailyCostResponse, err error) {
	response = NewDescribeAiModelDailyCostResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiUsageDataRequest() (request *DescribeAiUsageDataRequest) {
	request = &DescribeAiUsageDataRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiUsageData")

	return
}

func NewDescribeAiUsageDataResponse() (response *DescribeAiUsageDataResponse) {
	response = &DescribeAiUsageDataResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiUsageData 查询AI网关用量统计数据
func (c *Client) DescribeAiUsageData(request *DescribeAiUsageDataRequest) (response *DescribeAiUsageDataResponse, err error) {
	response = NewDescribeAiUsageDataResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAiModelDailyCacheHitRateRequest() (request *DescribeAiModelDailyCacheHitRateRequest) {
	request = &DescribeAiModelDailyCacheHitRateRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAiModelDailyCacheHitRate")

	return
}

func NewDescribeAiModelDailyCacheHitRateResponse() (response *DescribeAiModelDailyCacheHitRateResponse) {
	response = &DescribeAiModelDailyCacheHitRateResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAiModelDailyCacheHitRate 查询AI网关日模型缓存命中率
func (c *Client) DescribeAiModelDailyCacheHitRate(request *DescribeAiModelDailyCacheHitRateRequest) (response *DescribeAiModelDailyCacheHitRateResponse, err error) {
	response = NewDescribeAiModelDailyCacheHitRateResponse()
	err = c.ApiCall(request, response)
	return
}

