/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package zrm

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2025-10-14"
	SERVICE    = "zrm"
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


func NewCreateTagsRequest() (request *CreateTagsRequest) {
	request = &CreateTagsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateTags")

	return
}

func NewCreateTagsResponse() (response *CreateTagsResponse) {
	response = &CreateTagsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateTags 批量创建标签
func (c *Client) CreateTags(request *CreateTagsRequest) (response *CreateTagsResponse, err error) {
	response = NewCreateTagsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteTagsRequest() (request *DeleteTagsRequest) {
	request = &DeleteTagsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteTags")

	return
}

func NewDeleteTagsResponse() (response *DeleteTagsResponse) {
	response = &DeleteTagsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteTags 批量删除标签
func (c *Client) DeleteTags(request *DeleteTagsRequest) (response *DeleteTagsResponse, err error) {
	response = NewDeleteTagsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
	request = &DescribeTagsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeTags")

	return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
	response = &DescribeTagsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeTags 标签分页列表
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
	response = NewDescribeTagsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTagBindResourcesRequest() (request *TagBindResourcesRequest) {
	request = &TagBindResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TagBindResources")

	return
}

func NewTagBindResourcesResponse() (response *TagBindResourcesResponse) {
	response = &TagBindResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TagBindResources 标签批量绑定资源
func (c *Client) TagBindResources(request *TagBindResourcesRequest) (response *TagBindResourcesResponse, err error) {
	response = NewTagBindResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTagUnbindResourcesRequest() (request *TagUnbindResourcesRequest) {
	request = &TagUnbindResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TagUnbindResources")

	return
}

func NewTagUnbindResourcesResponse() (response *TagUnbindResourcesResponse) {
	response = &TagUnbindResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TagUnbindResources 标签批量解绑资源
func (c *Client) TagUnbindResources(request *TagUnbindResourcesRequest) (response *TagUnbindResourcesResponse, err error) {
	response = NewTagUnbindResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeResourceTagsRequest() (request *DescribeResourceTagsRequest) {
	request = &DescribeResourceTagsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeResourceTags")

	return
}

func NewDescribeResourceTagsResponse() (response *DescribeResourceTagsResponse) {
	response = &DescribeResourceTagsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeResourceTags 获取某个资源下绑定的所有标签列表。
func (c *Client) DescribeResourceTags(request *DescribeResourceTagsRequest) (response *DescribeResourceTagsResponse, err error) {
	response = NewDescribeResourceTagsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyResourceTagsRequest() (request *ModifyResourceTagsRequest) {
	request = &ModifyResourceTagsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyResourceTags")

	return
}

func NewModifyResourceTagsResponse() (response *ModifyResourceTagsResponse) {
	response = &ModifyResourceTagsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyResourceTags 修改某个资源的标签。
func (c *Client) ModifyResourceTags(request *ModifyResourceTagsRequest) (response *ModifyResourceTagsResponse, err error) {
	response = NewModifyResourceTagsResponse()
	err = c.ApiCall(request, response)
	return
}

