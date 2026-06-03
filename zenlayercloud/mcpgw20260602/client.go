/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package mcpgw

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2026-06-02"
	SERVICE    = "mcpgw"
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


func NewDescribeMcpMonthlyBillingRequest() (request *DescribeMcpMonthlyBillingRequest) {
	request = &DescribeMcpMonthlyBillingRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeMcpMonthlyBilling")

	return
}

func NewDescribeMcpMonthlyBillingResponse() (response *DescribeMcpMonthlyBillingResponse) {
	response = &DescribeMcpMonthlyBillingResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeMcpMonthlyBilling 查询MCP网关按月计费明细
func (c *Client) DescribeMcpMonthlyBilling(request *DescribeMcpMonthlyBillingRequest) (response *DescribeMcpMonthlyBillingResponse, err error) {
	response = NewDescribeMcpMonthlyBillingResponse()
	err = c.ApiCall(request, response)
	return
}

