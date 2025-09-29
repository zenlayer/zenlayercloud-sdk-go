/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package ccs

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2025-09-01"
	SERVICE    = "ccs"
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


func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
	request = &DescribeKeyPairsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeKeyPairs")

	return
}

func NewDescribeKeyPairsResponse() (response *DescribeKeyPairsResponse) {
	response = &DescribeKeyPairsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeKeyPairs 查询密钥对列表。
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
	response = NewDescribeKeyPairsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewImportKeyPairRequest() (request *ImportKeyPairRequest) {
	request = &ImportKeyPairRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ImportKeyPair")

	return
}

func NewImportKeyPairResponse() (response *ImportKeyPairResponse) {
	response = &ImportKeyPairResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ImportKeyPair 查询密钥对列表。
func (c *Client) ImportKeyPair(request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
	response = NewImportKeyPairResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyKeyPairAttributeRequest() (request *ModifyKeyPairAttributeRequest) {
	request = &ModifyKeyPairAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyKeyPairAttribute")

	return
}

func NewModifyKeyPairAttributeResponse() (response *ModifyKeyPairAttributeResponse) {
	response = &ModifyKeyPairAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyKeyPairAttribute 修改一个密钥对的属性。
func (c *Client) ModifyKeyPairAttribute(request *ModifyKeyPairAttributeRequest) (response *ModifyKeyPairAttributeResponse, err error) {
	response = NewModifyKeyPairAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteKeyPairsRequest() (request *DeleteKeyPairsRequest) {
	request = &DeleteKeyPairsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteKeyPairs")

	return
}

func NewDeleteKeyPairsResponse() (response *DeleteKeyPairsResponse) {
	response = &DeleteKeyPairsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteKeyPairs 删除一个或多个密钥对。
func (c *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
	response = NewDeleteKeyPairsResponse()
	err = c.ApiCall(request, response)
	return
}

