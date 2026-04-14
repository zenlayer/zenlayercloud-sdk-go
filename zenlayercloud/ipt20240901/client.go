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

// DescribeIPTransitDatacenters 本接口用于连接IP Transit 服务支持的数据中心。
func (c *Client) DescribeIPTransitDatacenters(request *DescribeIPTransitDatacentersRequest) (response *DescribeIPTransitDatacentersResponse, err error) {
	response = NewDescribeIPTransitDatacentersResponse()
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

// InquiryCreateIPTransitPrice 创建一条IP Transit 的询价。
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

// CreateIPTransit 创建一条IP Transit。
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

// DescribeIPTransits 本接口用于查询IP Transit资源列表。
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

// ModifyIPTransitBandwidth 修改一条IP Transit的带宽限速。
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

// ModifyIPTransitsAttribute 修改IP Transit的基本信息，包括名称和备注。
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

// DeleteIPTransit 删除一条IP Transit。
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

// DescribeIPTransitTraffic 查询IP Transit在指定时间段内的带宽数据。
func (c *Client) DescribeIPTransitTraffic(request *DescribeIPTransitTrafficRequest) (response *DescribeIPTransitTrafficResponse, err error) {
	response = NewDescribeIPTransitTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

