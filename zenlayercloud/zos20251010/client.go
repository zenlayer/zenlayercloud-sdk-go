/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package zos

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2025-10-10"
	SERVICE    = "zos"
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


func NewCreateCommandRequest() (request *CreateCommandRequest) {
	request = &CreateCommandRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateCommand")

	return
}

func NewCreateCommandResponse() (response *CreateCommandResponse) {
	response = &CreateCommandResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateCommand 创建命令。
func (c *Client) CreateCommand(request *CreateCommandRequest) (response *CreateCommandResponse, err error) {
	response = NewCreateCommandResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCommandsRequest() (request *DescribeCommandsRequest) {
	request = &DescribeCommandsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCommands")

	return
}

func NewDescribeCommandsResponse() (response *DescribeCommandsResponse) {
	response = &DescribeCommandsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCommands 查询命令列表。
func (c *Client) DescribeCommands(request *DescribeCommandsRequest) (response *DescribeCommandsResponse, err error) {
	response = NewDescribeCommandsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCommandRequest() (request *ModifyCommandRequest) {
	request = &ModifyCommandRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCommand")

	return
}

func NewModifyCommandResponse() (response *ModifyCommandResponse) {
	response = &ModifyCommandResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyCommand 修改命令。
func (c *Client) ModifyCommand(request *ModifyCommandRequest) (response *ModifyCommandResponse, err error) {
	response = NewModifyCommandResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteCommandRequest() (request *DeleteCommandRequest) {
	request = &DeleteCommandRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteCommand")

	return
}

func NewDeleteCommandResponse() (response *DeleteCommandResponse) {
	response = &DeleteCommandResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteCommand 删除命令。
func (c *Client) DeleteCommand(request *DeleteCommandRequest) (response *DeleteCommandResponse, err error) {
	response = NewDeleteCommandResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInvokeCommandRequest() (request *InvokeCommandRequest) {
	request = &InvokeCommandRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InvokeCommand")

	return
}

func NewInvokeCommandResponse() (response *InvokeCommandResponse) {
	response = &InvokeCommandResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InvokeCommand 执行一个已经创建并保存的命令。
func (c *Client) InvokeCommand(request *InvokeCommandRequest) (response *InvokeCommandResponse, err error) {
	response = NewInvokeCommandResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCommandInvocationsRequest() (request *DescribeCommandInvocationsRequest) {
	request = &DescribeCommandInvocationsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCommandInvocations")

	return
}

func NewDescribeCommandInvocationsResponse() (response *DescribeCommandInvocationsResponse) {
	response = &DescribeCommandInvocationsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCommandInvocations 查询执行记录列表
func (c *Client) DescribeCommandInvocations(request *DescribeCommandInvocationsRequest) (response *DescribeCommandInvocationsResponse, err error) {
	response = NewDescribeCommandInvocationsResponse()
	err = c.ApiCall(request, response)
	return
}

