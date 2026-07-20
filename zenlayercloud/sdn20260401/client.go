/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package sdn

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2026-04-01"
	SERVICE    = "sdn"
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


func NewCreatePortRequest() (request *CreatePortRequest) {
	request = &CreatePortRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreatePort")

	return
}

func NewCreatePortResponse() (response *CreatePortResponse) {
	response = &CreatePortResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreatePort 创建端口
func (c *Client) CreatePort(request *CreatePortRequest) (response *CreatePortResponse, err error) {
	response = NewCreatePortResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDestroyPortRequest() (request *DestroyPortRequest) {
	request = &DestroyPortRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DestroyPort")

	return
}

func NewDestroyPortResponse() (response *DestroyPortResponse) {
	response = &DestroyPortResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DestroyPort 销毁端口
func (c *Client) DestroyPort(request *DestroyPortRequest) (response *DestroyPortResponse, err error) {
	response = NewDestroyPortResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminatePortRequest() (request *TerminatePortRequest) {
	request = &TerminatePortRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminatePort")

	return
}

func NewTerminatePortResponse() (response *TerminatePortResponse) {
	response = &TerminatePortResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TerminatePort 终止端口
func (c *Client) TerminatePort(request *TerminatePortRequest) (response *TerminatePortResponse, err error) {
	response = NewTerminatePortResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewPortRequest() (request *RenewPortRequest) {
	request = &RenewPortRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewPort")

	return
}

func NewRenewPortResponse() (response *RenewPortResponse) {
	response = &RenewPortResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewPort 恢复端口
func (c *Client) RenewPort(request *RenewPortRequest) (response *RenewPortResponse, err error) {
	response = NewRenewPortResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPortAttributeRequest() (request *ModifyPortAttributeRequest) {
	request = &ModifyPortAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPortAttribute")

	return
}

func NewModifyPortAttributeResponse() (response *ModifyPortAttributeResponse) {
	response = &ModifyPortAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyPortAttribute 修改端口属性
func (c *Client) ModifyPortAttribute(request *ModifyPortAttributeRequest) (response *ModifyPortAttributeResponse, err error) {
	response = NewModifyPortAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePortsRequest() (request *DescribePortsRequest) {
	request = &DescribePortsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePorts")

	return
}

func NewDescribePortsResponse() (response *DescribePortsResponse) {
	response = &DescribePortsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePorts 获取端口列表
func (c *Client) DescribePorts(request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
	response = NewDescribePortsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDataCenterPortPriceRequest() (request *DescribeDataCenterPortPriceRequest) {
	request = &DescribeDataCenterPortPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDataCenterPortPrice")

	return
}

func NewDescribeDataCenterPortPriceResponse() (response *DescribeDataCenterPortPriceResponse) {
	response = &DescribeDataCenterPortPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeDataCenterPortPrice 获取数据中心在售端口类型及价格
func (c *Client) DescribeDataCenterPortPrice(request *DescribeDataCenterPortPriceRequest) (response *DescribeDataCenterPortPriceResponse, err error) {
	response = NewDescribeDataCenterPortPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePortTrafficRequest() (request *DescribePortTrafficRequest) {
	request = &DescribePortTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePortTraffic")

	return
}

func NewDescribePortTrafficResponse() (response *DescribePortTrafficResponse) {
	response = &DescribePortTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePortTraffic 查询端口流量
func (c *Client) DescribePortTraffic(request *DescribePortTrafficRequest) (response *DescribePortTrafficResponse, err error) {
	response = NewDescribePortTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePortUsableVlanRequest() (request *DescribePortUsableVlanRequest) {
	request = &DescribePortUsableVlanRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePortUsableVlan")

	return
}

func NewDescribePortUsableVlanResponse() (response *DescribePortUsableVlanResponse) {
	response = &DescribePortUsableVlanResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePortUsableVlan 查询端口可用vlan
func (c *Client) DescribePortUsableVlan(request *DescribePortUsableVlanRequest) (response *DescribePortUsableVlanResponse, err error) {
	response = NewDescribePortUsableVlanResponse()
	err = c.ApiCall(request, response)
	return
}

func NewQueryCloudOnrampPriceRequest() (request *QueryCloudOnrampPriceRequest) {
	request = &QueryCloudOnrampPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "QueryCloudOnrampPrice")

	return
}

func NewQueryCloudOnrampPriceResponse() (response *QueryCloudOnrampPriceResponse) {
	response = &QueryCloudOnrampPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// QueryCloudOnrampPrice 云连接带宽询价
func (c *Client) QueryCloudOnrampPrice(request *QueryCloudOnrampPriceRequest) (response *QueryCloudOnrampPriceResponse, err error) {
	response = NewQueryCloudOnrampPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewQueryDataCenterPortPriceRequest() (request *QueryDataCenterPortPriceRequest) {
	request = &QueryDataCenterPortPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "QueryDataCenterPortPrice")

	return
}

func NewQueryDataCenterPortPriceResponse() (response *QueryDataCenterPortPriceResponse) {
	response = &QueryDataCenterPortPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// QueryDataCenterPortPrice 数据中心端口询价
func (c *Client) QueryDataCenterPortPrice(request *QueryDataCenterPortPriceRequest) (response *QueryDataCenterPortPriceResponse, err error) {
	response = NewQueryDataCenterPortPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewQueryDataCenterPortPricesRequest() (request *QueryDataCenterPortPricesRequest) {
	request = &QueryDataCenterPortPricesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "QueryDataCenterPortPrices")

	return
}

func NewQueryDataCenterPortPricesResponse() (response *QueryDataCenterPortPricesResponse) {
	response = &QueryDataCenterPortPricesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// QueryDataCenterPortPrices 数据中心端口批量询价
func (c *Client) QueryDataCenterPortPrices(request *QueryDataCenterPortPricesRequest) (response *QueryDataCenterPortPricesResponse, err error) {
	response = NewQueryDataCenterPortPricesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewQueryPrivateConnectPriceRequest() (request *QueryPrivateConnectPriceRequest) {
	request = &QueryPrivateConnectPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "QueryPrivateConnectPrice")

	return
}

func NewQueryPrivateConnectPriceResponse() (response *QueryPrivateConnectPriceResponse) {
	response = &QueryPrivateConnectPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// QueryPrivateConnectPrice 二层网络专线询价
func (c *Client) QueryPrivateConnectPrice(request *QueryPrivateConnectPriceRequest) (response *QueryPrivateConnectPriceResponse, err error) {
	response = NewQueryPrivateConnectPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewQueryPrivateConnectBandwidthPriceRequest() (request *QueryPrivateConnectBandwidthPriceRequest) {
	request = &QueryPrivateConnectBandwidthPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "QueryPrivateConnectBandwidthPrice")

	return
}

func NewQueryPrivateConnectBandwidthPriceResponse() (response *QueryPrivateConnectBandwidthPriceResponse) {
	response = &QueryPrivateConnectBandwidthPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// QueryPrivateConnectBandwidthPrice 二层专线带宽询价
func (c *Client) QueryPrivateConnectBandwidthPrice(request *QueryPrivateConnectBandwidthPriceRequest) (response *QueryPrivateConnectBandwidthPriceResponse, err error) {
	response = NewQueryPrivateConnectBandwidthPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewQueryCloudRouterBandwidthPriceRequest() (request *QueryCloudRouterBandwidthPriceRequest) {
	request = &QueryCloudRouterBandwidthPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "QueryCloudRouterBandwidthPrice")

	return
}

func NewQueryCloudRouterBandwidthPriceResponse() (response *QueryCloudRouterBandwidthPriceResponse) {
	response = &QueryCloudRouterBandwidthPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// QueryCloudRouterBandwidthPrice 三层骨干带宽询价
func (c *Client) QueryCloudRouterBandwidthPrice(request *QueryCloudRouterBandwidthPriceRequest) (response *QueryCloudRouterBandwidthPriceResponse, err error) {
	response = NewQueryCloudRouterBandwidthPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeGoogleVlanUsageRequest() (request *DescribeGoogleVlanUsageRequest) {
	request = &DescribeGoogleVlanUsageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeGoogleVlanUsage")

	return
}

func NewDescribeGoogleVlanUsageResponse() (response *DescribeGoogleVlanUsageResponse) {
	response = &DescribeGoogleVlanUsageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeGoogleVlanUsage 查询Google接入点VLAN使用情况
func (c *Client) DescribeGoogleVlanUsage(request *DescribeGoogleVlanUsageRequest) (response *DescribeGoogleVlanUsageResponse, err error) {
	response = NewDescribeGoogleVlanUsageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeTencentVlanUsageRequest() (request *DescribeTencentVlanUsageRequest) {
	request = &DescribeTencentVlanUsageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeTencentVlanUsage")

	return
}

func NewDescribeTencentVlanUsageResponse() (response *DescribeTencentVlanUsageResponse) {
	response = &DescribeTencentVlanUsageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeTencentVlanUsage 查询腾讯云接入点VLAN使用情况
func (c *Client) DescribeTencentVlanUsage(request *DescribeTencentVlanUsageRequest) (response *DescribeTencentVlanUsageResponse, err error) {
	response = NewDescribeTencentVlanUsageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAliCloudVlanUsageRequest() (request *DescribeAliCloudVlanUsageRequest) {
	request = &DescribeAliCloudVlanUsageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAliCloudVlanUsage")

	return
}

func NewDescribeAliCloudVlanUsageResponse() (response *DescribeAliCloudVlanUsageResponse) {
	response = &DescribeAliCloudVlanUsageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAliCloudVlanUsage 查询阿里云接入点VLAN使用情况
func (c *Client) DescribeAliCloudVlanUsage(request *DescribeAliCloudVlanUsageRequest) (response *DescribeAliCloudVlanUsageResponse, err error) {
	response = NewDescribeAliCloudVlanUsageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeHuaweiCloudVlanUsageRequest() (request *DescribeHuaweiCloudVlanUsageRequest) {
	request = &DescribeHuaweiCloudVlanUsageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeHuaweiCloudVlanUsage")

	return
}

func NewDescribeHuaweiCloudVlanUsageResponse() (response *DescribeHuaweiCloudVlanUsageResponse) {
	response = &DescribeHuaweiCloudVlanUsageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeHuaweiCloudVlanUsage 查询华为云接入点VLAN使用情况
func (c *Client) DescribeHuaweiCloudVlanUsage(request *DescribeHuaweiCloudVlanUsageRequest) (response *DescribeHuaweiCloudVlanUsageResponse, err error) {
	response = NewDescribeHuaweiCloudVlanUsageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAzureVlanUsageRequest() (request *DescribeAzureVlanUsageRequest) {
	request = &DescribeAzureVlanUsageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAzureVlanUsage")

	return
}

func NewDescribeAzureVlanUsageResponse() (response *DescribeAzureVlanUsageResponse) {
	response = &DescribeAzureVlanUsageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAzureVlanUsage 查询Azure接入点VLAN使用情况
func (c *Client) DescribeAzureVlanUsage(request *DescribeAzureVlanUsageRequest) (response *DescribeAzureVlanUsageResponse, err error) {
	response = NewDescribeAzureVlanUsageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeOracleVlanUsageRequest() (request *DescribeOracleVlanUsageRequest) {
	request = &DescribeOracleVlanUsageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeOracleVlanUsage")

	return
}

func NewDescribeOracleVlanUsageResponse() (response *DescribeOracleVlanUsageResponse) {
	response = &DescribeOracleVlanUsageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeOracleVlanUsage 查询Oracle接入点VLAN使用情况
func (c *Client) DescribeOracleVlanUsage(request *DescribeOracleVlanUsageRequest) (response *DescribeOracleVlanUsageResponse, err error) {
	response = NewDescribeOracleVlanUsageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBytePlusVlanUsageRequest() (request *DescribeBytePlusVlanUsageRequest) {
	request = &DescribeBytePlusVlanUsageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBytePlusVlanUsage")

	return
}

func NewDescribeBytePlusVlanUsageResponse() (response *DescribeBytePlusVlanUsageResponse) {
	response = &DescribeBytePlusVlanUsageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBytePlusVlanUsage 查询BytePlus接入点VLAN使用情况
func (c *Client) DescribeBytePlusVlanUsage(request *DescribeBytePlusVlanUsageRequest) (response *DescribeBytePlusVlanUsageResponse, err error) {
	response = NewDescribeBytePlusVlanUsageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCloudBandwidthRequest() (request *ModifyCloudBandwidthRequest) {
	request = &ModifyCloudBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCloudBandwidth")

	return
}

func NewModifyCloudBandwidthResponse() (response *ModifyCloudBandwidthResponse) {
	response = &ModifyCloudBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyCloudBandwidth 修改云连接带宽
func (c *Client) ModifyCloudBandwidth(request *ModifyCloudBandwidthRequest) (response *ModifyCloudBandwidthResponse, err error) {
	response = NewModifyCloudBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudAvailableBandwidthTiersRequest() (request *DescribeCloudAvailableBandwidthTiersRequest) {
	request = &DescribeCloudAvailableBandwidthTiersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudAvailableBandwidthTiers")

	return
}

func NewDescribeCloudAvailableBandwidthTiersResponse() (response *DescribeCloudAvailableBandwidthTiersResponse) {
	response = &DescribeCloudAvailableBandwidthTiersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCloudAvailableBandwidthTiers 查询云连接可用带宽阶梯
func (c *Client) DescribeCloudAvailableBandwidthTiers(request *DescribeCloudAvailableBandwidthTiersRequest) (response *DescribeCloudAvailableBandwidthTiersResponse, err error) {
	response = NewDescribeCloudAvailableBandwidthTiersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAWSRegionsRequest() (request *DescribeAWSRegionsRequest) {
	request = &DescribeAWSRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAWSRegions")

	return
}

func NewDescribeAWSRegionsResponse() (response *DescribeAWSRegionsResponse) {
	response = &DescribeAWSRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAWSRegions 查询AWS接入点区域
func (c *Client) DescribeAWSRegions(request *DescribeAWSRegionsRequest) (response *DescribeAWSRegionsResponse, err error) {
	response = NewDescribeAWSRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeTencentRegionsRequest() (request *DescribeTencentRegionsRequest) {
	request = &DescribeTencentRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeTencentRegions")

	return
}

func NewDescribeTencentRegionsResponse() (response *DescribeTencentRegionsResponse) {
	response = &DescribeTencentRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeTencentRegions 查询腾讯云接入点区域
func (c *Client) DescribeTencentRegions(request *DescribeTencentRegionsRequest) (response *DescribeTencentRegionsResponse, err error) {
	response = NewDescribeTencentRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeGoogleRegionsRequest() (request *DescribeGoogleRegionsRequest) {
	request = &DescribeGoogleRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeGoogleRegions")

	return
}

func NewDescribeGoogleRegionsResponse() (response *DescribeGoogleRegionsResponse) {
	response = &DescribeGoogleRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeGoogleRegions 查询Google接入点区域
func (c *Client) DescribeGoogleRegions(request *DescribeGoogleRegionsRequest) (response *DescribeGoogleRegionsResponse, err error) {
	response = NewDescribeGoogleRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAzureRegionsRequest() (request *DescribeAzureRegionsRequest) {
	request = &DescribeAzureRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAzureRegions")

	return
}

func NewDescribeAzureRegionsResponse() (response *DescribeAzureRegionsResponse) {
	response = &DescribeAzureRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAzureRegions 查询Azure接入点区域
func (c *Client) DescribeAzureRegions(request *DescribeAzureRegionsRequest) (response *DescribeAzureRegionsResponse, err error) {
	response = NewDescribeAzureRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeOracleRegionsRequest() (request *DescribeOracleRegionsRequest) {
	request = &DescribeOracleRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeOracleRegions")

	return
}

func NewDescribeOracleRegionsResponse() (response *DescribeOracleRegionsResponse) {
	response = &DescribeOracleRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeOracleRegions 查询Oracle接入点区域
func (c *Client) DescribeOracleRegions(request *DescribeOracleRegionsRequest) (response *DescribeOracleRegionsResponse, err error) {
	response = NewDescribeOracleRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAliCloudRegionsRequest() (request *DescribeAliCloudRegionsRequest) {
	request = &DescribeAliCloudRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAliCloudRegions")

	return
}

func NewDescribeAliCloudRegionsResponse() (response *DescribeAliCloudRegionsResponse) {
	response = &DescribeAliCloudRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAliCloudRegions 查询阿里云接入点区域
func (c *Client) DescribeAliCloudRegions(request *DescribeAliCloudRegionsRequest) (response *DescribeAliCloudRegionsResponse, err error) {
	response = NewDescribeAliCloudRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeHuaweiCloudRegionsRequest() (request *DescribeHuaweiCloudRegionsRequest) {
	request = &DescribeHuaweiCloudRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeHuaweiCloudRegions")

	return
}

func NewDescribeHuaweiCloudRegionsResponse() (response *DescribeHuaweiCloudRegionsResponse) {
	response = &DescribeHuaweiCloudRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeHuaweiCloudRegions 查询华为云接入点区域
func (c *Client) DescribeHuaweiCloudRegions(request *DescribeHuaweiCloudRegionsRequest) (response *DescribeHuaweiCloudRegionsResponse, err error) {
	response = NewDescribeHuaweiCloudRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBytePlusRegionsRequest() (request *DescribeBytePlusRegionsRequest) {
	request = &DescribeBytePlusRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBytePlusRegions")

	return
}

func NewDescribeBytePlusRegionsResponse() (response *DescribeBytePlusRegionsResponse) {
	response = &DescribeBytePlusRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBytePlusRegions 查询BytePlus接入点区域
func (c *Client) DescribeBytePlusRegions(request *DescribeBytePlusRegionsRequest) (response *DescribeBytePlusRegionsResponse, err error) {
	response = NewDescribeBytePlusRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAWSVlanUsageRequest() (request *DescribeAWSVlanUsageRequest) {
	request = &DescribeAWSVlanUsageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAWSVlanUsage")

	return
}

func NewDescribeAWSVlanUsageResponse() (response *DescribeAWSVlanUsageResponse) {
	response = &DescribeAWSVlanUsageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAWSVlanUsage 查询AWS接入点VLAN使用情况
func (c *Client) DescribeAWSVlanUsage(request *DescribeAWSVlanUsageRequest) (response *DescribeAWSVlanUsageResponse, err error) {
	response = NewDescribeAWSVlanUsageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateCloudRouterRequest() (request *CreateCloudRouterRequest) {
	request = &CreateCloudRouterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateCloudRouter")

	return
}

func NewCreateCloudRouterResponse() (response *CreateCloudRouterResponse) {
	response = &CreateCloudRouterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateCloudRouter 创建三层网络
func (c *Client) CreateCloudRouter(request *CreateCloudRouterRequest) (response *CreateCloudRouterResponse, err error) {
	response = NewCreateCloudRouterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteCloudRouterEdgePointRequest() (request *DeleteCloudRouterEdgePointRequest) {
	request = &DeleteCloudRouterEdgePointRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteCloudRouterEdgePoint")

	return
}

func NewDeleteCloudRouterEdgePointResponse() (response *DeleteCloudRouterEdgePointResponse) {
	response = &DeleteCloudRouterEdgePointResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteCloudRouterEdgePoint 删除连接点
func (c *Client) DeleteCloudRouterEdgePoint(request *DeleteCloudRouterEdgePointRequest) (response *DeleteCloudRouterEdgePointResponse, err error) {
	response = NewDeleteCloudRouterEdgePointResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAddCloudRouterEdgePointsRequest() (request *AddCloudRouterEdgePointsRequest) {
	request = &AddCloudRouterEdgePointsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AddCloudRouterEdgePoints")

	return
}

func NewAddCloudRouterEdgePointsResponse() (response *AddCloudRouterEdgePointsResponse) {
	response = &AddCloudRouterEdgePointsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AddCloudRouterEdgePoints 新增连接点
func (c *Client) AddCloudRouterEdgePoints(request *AddCloudRouterEdgePointsRequest) (response *AddCloudRouterEdgePointsResponse, err error) {
	response = NewAddCloudRouterEdgePointsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCloudRoutersAttributeRequest() (request *ModifyCloudRoutersAttributeRequest) {
	request = &ModifyCloudRoutersAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCloudRoutersAttribute")

	return
}

func NewModifyCloudRoutersAttributeResponse() (response *ModifyCloudRoutersAttributeResponse) {
	response = &ModifyCloudRoutersAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyCloudRoutersAttribute 修改三层网络属性
func (c *Client) ModifyCloudRoutersAttribute(request *ModifyCloudRoutersAttributeRequest) (response *ModifyCloudRoutersAttributeResponse, err error) {
	response = NewModifyCloudRoutersAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRouterAvailableVpcsRequest() (request *DescribeCloudRouterAvailableVpcsRequest) {
	request = &DescribeCloudRouterAvailableVpcsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouterAvailableVpcs")

	return
}

func NewDescribeCloudRouterAvailableVpcsResponse() (response *DescribeCloudRouterAvailableVpcsResponse) {
	response = &DescribeCloudRouterAvailableVpcsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCloudRouterAvailableVpcs 查询可用VPC
func (c *Client) DescribeCloudRouterAvailableVpcs(request *DescribeCloudRouterAvailableVpcsRequest) (response *DescribeCloudRouterAvailableVpcsResponse, err error) {
	response = NewDescribeCloudRouterAvailableVpcsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRouterEdgePointTrafficRequest() (request *DescribeCloudRouterEdgePointTrafficRequest) {
	request = &DescribeCloudRouterEdgePointTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouterEdgePointTraffic")

	return
}

func NewDescribeCloudRouterEdgePointTrafficResponse() (response *DescribeCloudRouterEdgePointTrafficResponse) {
	response = &DescribeCloudRouterEdgePointTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCloudRouterEdgePointTraffic 查询连接点流量
func (c *Client) DescribeCloudRouterEdgePointTraffic(request *DescribeCloudRouterEdgePointTrafficRequest) (response *DescribeCloudRouterEdgePointTrafficResponse, err error) {
	response = NewDescribeCloudRouterEdgePointTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRouterDCToDCTrafficRequest() (request *DescribeCloudRouterDCToDCTrafficRequest) {
	request = &DescribeCloudRouterDCToDCTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouterDCToDCTraffic")

	return
}

func NewDescribeCloudRouterDCToDCTrafficResponse() (response *DescribeCloudRouterDCToDCTrafficResponse) {
	response = &DescribeCloudRouterDCToDCTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCloudRouterDCToDCTraffic 查询数据中心间流量
func (c *Client) DescribeCloudRouterDCToDCTraffic(request *DescribeCloudRouterDCToDCTrafficRequest) (response *DescribeCloudRouterDCToDCTrafficResponse, err error) {
	response = NewDescribeCloudRouterDCToDCTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCloudRouterEdgePointBandwidthRequest() (request *ModifyCloudRouterEdgePointBandwidthRequest) {
	request = &ModifyCloudRouterEdgePointBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCloudRouterEdgePointBandwidth")

	return
}

func NewModifyCloudRouterEdgePointBandwidthResponse() (response *ModifyCloudRouterEdgePointBandwidthResponse) {
	response = &ModifyCloudRouterEdgePointBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyCloudRouterEdgePointBandwidth 修改连接点带宽
func (c *Client) ModifyCloudRouterEdgePointBandwidth(request *ModifyCloudRouterEdgePointBandwidthRequest) (response *ModifyCloudRouterEdgePointBandwidthResponse, err error) {
	response = NewModifyCloudRouterEdgePointBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCloudRouterEdgePointRequest() (request *ModifyCloudRouterEdgePointRequest) {
	request = &ModifyCloudRouterEdgePointRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCloudRouterEdgePoint")

	return
}

func NewModifyCloudRouterEdgePointResponse() (response *ModifyCloudRouterEdgePointResponse) {
	response = &ModifyCloudRouterEdgePointResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyCloudRouterEdgePoint 修改连接点配置
func (c *Client) ModifyCloudRouterEdgePoint(request *ModifyCloudRouterEdgePointRequest) (response *ModifyCloudRouterEdgePointResponse, err error) {
	response = NewModifyCloudRouterEdgePointResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteCloudRouterRequest() (request *DeleteCloudRouterRequest) {
	request = &DeleteCloudRouterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteCloudRouter")

	return
}

func NewDeleteCloudRouterResponse() (response *DeleteCloudRouterResponse) {
	response = &DeleteCloudRouterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteCloudRouter 删除三层网络
func (c *Client) DeleteCloudRouter(request *DeleteCloudRouterRequest) (response *DeleteCloudRouterResponse, err error) {
	response = NewDeleteCloudRouterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDestroyCloudRouterRequest() (request *DestroyCloudRouterRequest) {
	request = &DestroyCloudRouterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DestroyCloudRouter")

	return
}

func NewDestroyCloudRouterResponse() (response *DestroyCloudRouterResponse) {
	response = &DestroyCloudRouterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DestroyCloudRouter 销毁三层网络
func (c *Client) DestroyCloudRouter(request *DestroyCloudRouterRequest) (response *DestroyCloudRouterResponse, err error) {
	response = NewDestroyCloudRouterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewCloudRouterRequest() (request *RenewCloudRouterRequest) {
	request = &RenewCloudRouterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewCloudRouter")

	return
}

func NewRenewCloudRouterResponse() (response *RenewCloudRouterResponse) {
	response = &RenewCloudRouterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewCloudRouter 恢复三层网络
func (c *Client) RenewCloudRouter(request *RenewCloudRouterRequest) (response *RenewCloudRouterResponse, err error) {
	response = NewRenewCloudRouterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRouterAvailablePortsRequest() (request *DescribeCloudRouterAvailablePortsRequest) {
	request = &DescribeCloudRouterAvailablePortsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouterAvailablePorts")

	return
}

func NewDescribeCloudRouterAvailablePortsResponse() (response *DescribeCloudRouterAvailablePortsResponse) {
	response = &DescribeCloudRouterAvailablePortsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCloudRouterAvailablePorts 查询可用端口
func (c *Client) DescribeCloudRouterAvailablePorts(request *DescribeCloudRouterAvailablePortsRequest) (response *DescribeCloudRouterAvailablePortsResponse, err error) {
	response = NewDescribeCloudRouterAvailablePortsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRoutersRequest() (request *DescribeCloudRoutersRequest) {
	request = &DescribeCloudRoutersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouters")

	return
}

func NewDescribeCloudRoutersResponse() (response *DescribeCloudRoutersResponse) {
	response = &DescribeCloudRoutersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCloudRouters 查询三层网络列表
func (c *Client) DescribeCloudRouters(request *DescribeCloudRoutersRequest) (response *DescribeCloudRoutersResponse, err error) {
	response = NewDescribeCloudRoutersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDatacentersRequest() (request *DescribeDatacentersRequest) {
	request = &DescribeDatacentersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDatacenters")

	return
}

func NewDescribeDatacentersResponse() (response *DescribeDatacentersResponse) {
	response = &DescribeDatacentersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeDatacenters 查询数据中心列表
func (c *Client) DescribeDatacenters(request *DescribeDatacentersRequest) (response *DescribeDatacentersResponse, err error) {
	response = NewDescribeDatacentersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeVirtualEdgeDatacentersRequest() (request *DescribeVirtualEdgeDatacentersRequest) {
	request = &DescribeVirtualEdgeDatacentersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeVirtualEdgeDatacenters")

	return
}

func NewDescribeVirtualEdgeDatacentersResponse() (response *DescribeVirtualEdgeDatacentersResponse) {
	response = &DescribeVirtualEdgeDatacentersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeVirtualEdgeDatacenters 查询边缘网关数据中心列表
func (c *Client) DescribeVirtualEdgeDatacenters(request *DescribeVirtualEdgeDatacentersRequest) (response *DescribeVirtualEdgeDatacentersResponse, err error) {
	response = NewDescribeVirtualEdgeDatacentersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBorderGatewayDatacentersRequest() (request *DescribeBorderGatewayDatacentersRequest) {
	request = &DescribeBorderGatewayDatacentersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBorderGatewayDatacenters")

	return
}

func NewDescribeBorderGatewayDatacentersResponse() (response *DescribeBorderGatewayDatacentersResponse) {
	response = &DescribeBorderGatewayDatacentersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBorderGatewayDatacenters 查询边界网关数据中心列表
func (c *Client) DescribeBorderGatewayDatacenters(request *DescribeBorderGatewayDatacentersRequest) (response *DescribeBorderGatewayDatacentersResponse, err error) {
	response = NewDescribeBorderGatewayDatacentersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeVPCDatacentersRequest() (request *DescribeVPCDatacentersRequest) {
	request = &DescribeVPCDatacentersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeVPCDatacenters")

	return
}

func NewDescribeVPCDatacentersResponse() (response *DescribeVPCDatacentersResponse) {
	response = &DescribeVPCDatacentersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeVPCDatacenters 查询 VPC 数据中心列表
func (c *Client) DescribeVPCDatacenters(request *DescribeVPCDatacentersRequest) (response *DescribeVPCDatacentersResponse, err error) {
	response = NewDescribeVPCDatacentersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDatacentersWithServiceRequest() (request *DescribeDatacentersWithServiceRequest) {
	request = &DescribeDatacentersWithServiceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDatacentersWithService")

	return
}

func NewDescribeDatacentersWithServiceResponse() (response *DescribeDatacentersWithServiceResponse) {
	response = &DescribeDatacentersWithServiceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeDatacentersWithService 查询含服务的数据中心列表
func (c *Client) DescribeDatacentersWithService(request *DescribeDatacentersWithServiceRequest) (response *DescribeDatacentersWithServiceResponse, err error) {
	response = NewDescribeDatacentersWithServiceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePrivateConnectsRequest() (request *DescribePrivateConnectsRequest) {
	request = &DescribePrivateConnectsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePrivateConnects")

	return
}

func NewDescribePrivateConnectsResponse() (response *DescribePrivateConnectsResponse) {
	response = &DescribePrivateConnectsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePrivateConnects 获取二层网络专线列表
func (c *Client) DescribePrivateConnects(request *DescribePrivateConnectsRequest) (response *DescribePrivateConnectsResponse, err error) {
	response = NewDescribePrivateConnectsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreatePrivateConnectRequest() (request *CreatePrivateConnectRequest) {
	request = &CreatePrivateConnectRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreatePrivateConnect")

	return
}

func NewCreatePrivateConnectResponse() (response *CreatePrivateConnectResponse) {
	response = &CreatePrivateConnectResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreatePrivateConnect 创建二层网络专线
func (c *Client) CreatePrivateConnect(request *CreatePrivateConnectRequest) (response *CreatePrivateConnectResponse, err error) {
	response = NewCreatePrivateConnectResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPrivateConnectBandwidthRequest() (request *ModifyPrivateConnectBandwidthRequest) {
	request = &ModifyPrivateConnectBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPrivateConnectBandwidth")

	return
}

func NewModifyPrivateConnectBandwidthResponse() (response *ModifyPrivateConnectBandwidthResponse) {
	response = &ModifyPrivateConnectBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyPrivateConnectBandwidth 修改二层网络专线带宽
func (c *Client) ModifyPrivateConnectBandwidth(request *ModifyPrivateConnectBandwidthRequest) (response *ModifyPrivateConnectBandwidthResponse, err error) {
	response = NewModifyPrivateConnectBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePrivateConnectTrafficRequest() (request *DescribePrivateConnectTrafficRequest) {
	request = &DescribePrivateConnectTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePrivateConnectTraffic")

	return
}

func NewDescribePrivateConnectTrafficResponse() (response *DescribePrivateConnectTrafficResponse) {
	response = &DescribePrivateConnectTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePrivateConnectTraffic 查询二层网络专线流量
func (c *Client) DescribePrivateConnectTraffic(request *DescribePrivateConnectTrafficRequest) (response *DescribePrivateConnectTrafficResponse, err error) {
	response = NewDescribePrivateConnectTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeletePrivateConnectRequest() (request *DeletePrivateConnectRequest) {
	request = &DeletePrivateConnectRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeletePrivateConnect")

	return
}

func NewDeletePrivateConnectResponse() (response *DeletePrivateConnectResponse) {
	response = &DeletePrivateConnectResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeletePrivateConnect 删除二层网络专线
func (c *Client) DeletePrivateConnect(request *DeletePrivateConnectRequest) (response *DeletePrivateConnectResponse, err error) {
	response = NewDeletePrivateConnectResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDestroyPrivateConnectRequest() (request *DestroyPrivateConnectRequest) {
	request = &DestroyPrivateConnectRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DestroyPrivateConnect")

	return
}

func NewDestroyPrivateConnectResponse() (response *DestroyPrivateConnectResponse) {
	response = &DestroyPrivateConnectResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DestroyPrivateConnect 销毁二层网络专线
func (c *Client) DestroyPrivateConnect(request *DestroyPrivateConnectRequest) (response *DestroyPrivateConnectResponse, err error) {
	response = NewDestroyPrivateConnectResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePrivateConnectAvailablePortsRequest() (request *DescribePrivateConnectAvailablePortsRequest) {
	request = &DescribePrivateConnectAvailablePortsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePrivateConnectAvailablePorts")

	return
}

func NewDescribePrivateConnectAvailablePortsResponse() (response *DescribePrivateConnectAvailablePortsResponse) {
	response = &DescribePrivateConnectAvailablePortsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePrivateConnectAvailablePorts 查询可加入的数据中心端口
func (c *Client) DescribePrivateConnectAvailablePorts(request *DescribePrivateConnectAvailablePortsRequest) (response *DescribePrivateConnectAvailablePortsResponse, err error) {
	response = NewDescribePrivateConnectAvailablePortsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPrivateConnectsAttributeRequest() (request *ModifyPrivateConnectsAttributeRequest) {
	request = &ModifyPrivateConnectsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPrivateConnectsAttribute")

	return
}

func NewModifyPrivateConnectsAttributeResponse() (response *ModifyPrivateConnectsAttributeResponse) {
	response = &ModifyPrivateConnectsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyPrivateConnectsAttribute 修改二层网络专线属性
func (c *Client) ModifyPrivateConnectsAttribute(request *ModifyPrivateConnectsAttributeRequest) (response *ModifyPrivateConnectsAttributeResponse, err error) {
	response = NewModifyPrivateConnectsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewPrivateConnectRequest() (request *RenewPrivateConnectRequest) {
	request = &RenewPrivateConnectRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewPrivateConnect")

	return
}

func NewRenewPrivateConnectResponse() (response *RenewPrivateConnectResponse) {
	response = &RenewPrivateConnectResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewPrivateConnect 恢复二层网络专线
func (c *Client) RenewPrivateConnect(request *RenewPrivateConnectRequest) (response *RenewPrivateConnectResponse, err error) {
	response = NewRenewPrivateConnectResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryCreatePrivateConnectPriceRequest() (request *InquiryCreatePrivateConnectPriceRequest) {
	request = &InquiryCreatePrivateConnectPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryCreatePrivateConnectPrice")

	return
}

func NewInquiryCreatePrivateConnectPriceResponse() (response *InquiryCreatePrivateConnectPriceResponse) {
	response = &InquiryCreatePrivateConnectPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryCreatePrivateConnectPrice 二层网络专线询价
func (c *Client) InquiryCreatePrivateConnectPrice(request *InquiryCreatePrivateConnectPriceRequest) (response *InquiryCreatePrivateConnectPriceResponse, err error) {
	response = NewInquiryCreatePrivateConnectPriceResponse()
	err = c.ApiCall(request, response)
	return
}

