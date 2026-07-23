/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package ipt

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2024-09-01"
	SERVICE    = "ipt"
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


func NewDescribeIPTransitDatacentersRequest() (request *DescribeIPTransitDatacentersRequest) {
	request = &DescribeIPTransitDatacentersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeIPTransitDatacenters")

	return
}

func NewDescribeIPTransitDatacentersResponse() (response *DescribeIPTransitDatacentersResponse) {
	response = &DescribeIPTransitDatacentersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeIPTransitDatacenters 查询IP Transit可连接数据中心
func (c *Client) DescribeIPTransitDatacenters(request *DescribeIPTransitDatacentersRequest) (response *DescribeIPTransitDatacentersResponse, err error) {
	response = NewDescribeIPTransitDatacentersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeIPTransitAvailableAsnsRequest() (request *DescribeIPTransitAvailableAsnsRequest) {
	request = &DescribeIPTransitAvailableAsnsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeIPTransitAvailableAsns")

	return
}

func NewDescribeIPTransitAvailableAsnsResponse() (response *DescribeIPTransitAvailableAsnsResponse) {
	response = &DescribeIPTransitAvailableAsnsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeIPTransitAvailableAsns 查询IP Transit可用 ASN
func (c *Client) DescribeIPTransitAvailableAsns(request *DescribeIPTransitAvailableAsnsRequest) (response *DescribeIPTransitAvailableAsnsResponse, err error) {
	response = NewDescribeIPTransitAvailableAsnsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeIPTransitAvailableCidrBlocksRequest() (request *DescribeIPTransitAvailableCidrBlocksRequest) {
	request = &DescribeIPTransitAvailableCidrBlocksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeIPTransitAvailableCidrBlocks")

	return
}

func NewDescribeIPTransitAvailableCidrBlocksResponse() (response *DescribeIPTransitAvailableCidrBlocksResponse) {
	response = &DescribeIPTransitAvailableCidrBlocksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeIPTransitAvailableCidrBlocks 查询IP Transit可用公网地址段
func (c *Client) DescribeIPTransitAvailableCidrBlocks(request *DescribeIPTransitAvailableCidrBlocksRequest) (response *DescribeIPTransitAvailableCidrBlocksResponse, err error) {
	response = NewDescribeIPTransitAvailableCidrBlocksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryCreateIPTransitPriceRequest() (request *InquiryCreateIPTransitPriceRequest) {
	request = &InquiryCreateIPTransitPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryCreateIPTransitPrice")

	return
}

func NewInquiryCreateIPTransitPriceResponse() (response *InquiryCreateIPTransitPriceResponse) {
	response = &InquiryCreateIPTransitPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryCreateIPTransitPrice iP Transit创建询价
func (c *Client) InquiryCreateIPTransitPrice(request *InquiryCreateIPTransitPriceRequest) (response *InquiryCreateIPTransitPriceResponse, err error) {
	response = NewInquiryCreateIPTransitPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateIPTransitRequest() (request *CreateIPTransitRequest) {
	request = &CreateIPTransitRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateIPTransit")

	return
}

func NewCreateIPTransitResponse() (response *CreateIPTransitResponse) {
	response = &CreateIPTransitResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateIPTransit 创建IP Transit
func (c *Client) CreateIPTransit(request *CreateIPTransitRequest) (response *CreateIPTransitResponse, err error) {
	response = NewCreateIPTransitResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeIPTransitsRequest() (request *DescribeIPTransitsRequest) {
	request = &DescribeIPTransitsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeIPTransits")

	return
}

func NewDescribeIPTransitsResponse() (response *DescribeIPTransitsResponse) {
	response = &DescribeIPTransitsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeIPTransits 查询IP Transit列表
func (c *Client) DescribeIPTransits(request *DescribeIPTransitsRequest) (response *DescribeIPTransitsResponse, err error) {
	response = NewDescribeIPTransitsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyIPTransitBandwidthRequest() (request *ModifyIPTransitBandwidthRequest) {
	request = &ModifyIPTransitBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyIPTransitBandwidth")

	return
}

func NewModifyIPTransitBandwidthResponse() (response *ModifyIPTransitBandwidthResponse) {
	response = &ModifyIPTransitBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyIPTransitBandwidth 修改IP Transit带宽
func (c *Client) ModifyIPTransitBandwidth(request *ModifyIPTransitBandwidthRequest) (response *ModifyIPTransitBandwidthResponse, err error) {
	response = NewModifyIPTransitBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyIPTransitsAttributeRequest() (request *ModifyIPTransitsAttributeRequest) {
	request = &ModifyIPTransitsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyIPTransitsAttribute")

	return
}

func NewModifyIPTransitsAttributeResponse() (response *ModifyIPTransitsAttributeResponse) {
	response = &ModifyIPTransitsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyIPTransitsAttribute 修改IP Transit属性
func (c *Client) ModifyIPTransitsAttribute(request *ModifyIPTransitsAttributeRequest) (response *ModifyIPTransitsAttributeResponse, err error) {
	response = NewModifyIPTransitsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteIPTransitRequest() (request *DeleteIPTransitRequest) {
	request = &DeleteIPTransitRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteIPTransit")

	return
}

func NewDeleteIPTransitResponse() (response *DeleteIPTransitResponse) {
	response = &DeleteIPTransitResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteIPTransit 删除IP Transit
func (c *Client) DeleteIPTransit(request *DeleteIPTransitRequest) (response *DeleteIPTransitResponse, err error) {
	response = NewDeleteIPTransitResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeIPTransitTrafficRequest() (request *DescribeIPTransitTrafficRequest) {
	request = &DescribeIPTransitTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeIPTransitTraffic")

	return
}

func NewDescribeIPTransitTrafficResponse() (response *DescribeIPTransitTrafficResponse) {
	response = &DescribeIPTransitTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeIPTransitTraffic 查询IP Transit流量
func (c *Client) DescribeIPTransitTraffic(request *DescribeIPTransitTrafficRequest) (response *DescribeIPTransitTrafficResponse, err error) {
	response = NewDescribeIPTransitTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryModifyIPTransitPriceRequest() (request *InquiryModifyIPTransitPriceRequest) {
	request = &InquiryModifyIPTransitPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryModifyIPTransitPrice")

	return
}

func NewInquiryModifyIPTransitPriceResponse() (response *InquiryModifyIPTransitPriceResponse) {
	response = &InquiryModifyIPTransitPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryModifyIPTransitPrice IP Transit变配询价
func (c *Client) InquiryModifyIPTransitPrice(request *InquiryModifyIPTransitPriceRequest) (response *InquiryModifyIPTransitPriceResponse, err error) {
	response = NewInquiryModifyIPTransitPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyIPTransitConfigRequest() (request *ModifyIPTransitConfigRequest) {
	request = &ModifyIPTransitConfigRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyIPTransitConfig")

	return
}

func NewModifyIPTransitConfigResponse() (response *ModifyIPTransitConfigResponse) {
	response = &ModifyIPTransitConfigResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyIPTransitConfig 修改IP Transit配置
func (c *Client) ModifyIPTransitConfig(request *ModifyIPTransitConfigRequest) (response *ModifyIPTransitConfigResponse, err error) {
	response = NewModifyIPTransitConfigResponse()
	err = c.ApiCall(request, response)
	return
}

