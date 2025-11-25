/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package zdns

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2025-11-01"
	SERVICE    = "zdns"
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


func NewAddPrivateZoneRequest() (request *AddPrivateZoneRequest) {
	request = &AddPrivateZoneRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AddPrivateZone")

	return
}

func NewAddPrivateZoneResponse() (response *AddPrivateZoneResponse) {
	response = &AddPrivateZoneResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AddPrivateZone 创建内网权威域名。
func (c *Client) AddPrivateZone(request *AddPrivateZoneRequest) (response *AddPrivateZoneResponse, err error) {
	response = NewAddPrivateZoneResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePrivateZonesRequest() (request *DescribePrivateZonesRequest) {
	request = &DescribePrivateZonesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePrivateZones")

	return
}

func NewDescribePrivateZonesResponse() (response *DescribePrivateZonesResponse) {
	response = &DescribePrivateZonesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePrivateZones 查询内网权威域名列表。
func (c *Client) DescribePrivateZones(request *DescribePrivateZonesRequest) (response *DescribePrivateZonesResponse, err error) {
	response = NewDescribePrivateZonesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPrivateZoneRequest() (request *ModifyPrivateZoneRequest) {
	request = &ModifyPrivateZoneRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPrivateZone")

	return
}

func NewModifyPrivateZoneResponse() (response *ModifyPrivateZoneResponse) {
	response = &ModifyPrivateZoneResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyPrivateZone 修改内网权威域名信息。包括备注，是否开启子域名递归代理等。
func (c *Client) ModifyPrivateZone(request *ModifyPrivateZoneRequest) (response *ModifyPrivateZoneResponse, err error) {
	response = NewModifyPrivateZoneResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeletePrivateZoneRequest() (request *DeletePrivateZoneRequest) {
	request = &DeletePrivateZoneRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeletePrivateZone")

	return
}

func NewDeletePrivateZoneResponse() (response *DeletePrivateZoneResponse) {
	response = &DeletePrivateZoneResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeletePrivateZone 删除内网权威域名。
func (c *Client) DeletePrivateZone(request *DeletePrivateZoneRequest) (response *DeletePrivateZoneResponse, err error) {
	response = NewDeletePrivateZoneResponse()
	err = c.ApiCall(request, response)
	return
}

func NewBindPrivateZoneVpcRequest() (request *BindPrivateZoneVpcRequest) {
	request = &BindPrivateZoneVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "BindPrivateZoneVpc")

	return
}

func NewBindPrivateZoneVpcResponse() (response *BindPrivateZoneVpcResponse) {
	response = &BindPrivateZoneVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// BindPrivateZoneVpc 内网权威域名额外绑定VPC。
func (c *Client) BindPrivateZoneVpc(request *BindPrivateZoneVpcRequest) (response *BindPrivateZoneVpcResponse, err error) {
	response = NewBindPrivateZoneVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnbindPrivateZoneVpcRequest() (request *UnbindPrivateZoneVpcRequest) {
	request = &UnbindPrivateZoneVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnbindPrivateZoneVpc")

	return
}

func NewUnbindPrivateZoneVpcResponse() (response *UnbindPrivateZoneVpcResponse) {
	response = &UnbindPrivateZoneVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnbindPrivateZoneVpc 内网权威域名解除VPC绑定。
func (c *Client) UnbindPrivateZoneVpc(request *UnbindPrivateZoneVpcRequest) (response *UnbindPrivateZoneVpcResponse, err error) {
	response = NewUnbindPrivateZoneVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAddPrivateZoneRecordRequest() (request *AddPrivateZoneRecordRequest) {
	request = &AddPrivateZoneRecordRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AddPrivateZoneRecord")

	return
}

func NewAddPrivateZoneRecordResponse() (response *AddPrivateZoneRecordResponse) {
	response = &AddPrivateZoneRecordResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AddPrivateZoneRecord 新增内网权威域名解析记录。
func (c *Client) AddPrivateZoneRecord(request *AddPrivateZoneRecordRequest) (response *AddPrivateZoneRecordResponse, err error) {
	response = NewAddPrivateZoneRecordResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePrivateZoneRecordsRequest() (request *DescribePrivateZoneRecordsRequest) {
	request = &DescribePrivateZoneRecordsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePrivateZoneRecords")

	return
}

func NewDescribePrivateZoneRecordsResponse() (response *DescribePrivateZoneRecordsResponse) {
	response = &DescribePrivateZoneRecordsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePrivateZoneRecords 查询内网权威域名的解析记录列表。
func (c *Client) DescribePrivateZoneRecords(request *DescribePrivateZoneRecordsRequest) (response *DescribePrivateZoneRecordsResponse, err error) {
	response = NewDescribePrivateZoneRecordsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPrivateZoneRecordRequest() (request *ModifyPrivateZoneRecordRequest) {
	request = &ModifyPrivateZoneRecordRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPrivateZoneRecord")

	return
}

func NewModifyPrivateZoneRecordResponse() (response *ModifyPrivateZoneRecordResponse) {
	response = &ModifyPrivateZoneRecordResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyPrivateZoneRecord 修改内网权威域名解析记录。
func (c *Client) ModifyPrivateZoneRecord(request *ModifyPrivateZoneRecordRequest) (response *ModifyPrivateZoneRecordResponse, err error) {
	response = NewModifyPrivateZoneRecordResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPrivateZoneRecordsStatusRequest() (request *ModifyPrivateZoneRecordsStatusRequest) {
	request = &ModifyPrivateZoneRecordsStatusRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPrivateZoneRecordsStatus")

	return
}

func NewModifyPrivateZoneRecordsStatusResponse() (response *ModifyPrivateZoneRecordsStatusResponse) {
	response = &ModifyPrivateZoneRecordsStatusResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyPrivateZoneRecordsStatus 修改内网权威域名解析记录的生效状态。
func (c *Client) ModifyPrivateZoneRecordsStatus(request *ModifyPrivateZoneRecordsStatusRequest) (response *ModifyPrivateZoneRecordsStatusResponse, err error) {
	response = NewModifyPrivateZoneRecordsStatusResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeletePrivateZoneRecordRequest() (request *DeletePrivateZoneRecordRequest) {
	request = &DeletePrivateZoneRecordRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeletePrivateZoneRecord")

	return
}

func NewDeletePrivateZoneRecordResponse() (response *DeletePrivateZoneRecordResponse) {
	response = &DeletePrivateZoneRecordResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeletePrivateZoneRecord 删除内网权威域名解析记录。
func (c *Client) DeletePrivateZoneRecord(request *DeletePrivateZoneRecordRequest) (response *DeletePrivateZoneRecordResponse, err error) {
	response = NewDeletePrivateZoneRecordResponse()
	err = c.ApiCall(request, response)
	return
}

