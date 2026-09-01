/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package zsp

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2026-08-01"
	SERVICE    = "zsp"
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


func NewDescribeTicketsRequest() (request *DescribeTicketsRequest) {
	request = &DescribeTicketsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeTickets")

	return
}

func NewDescribeTicketsResponse() (response *DescribeTicketsResponse) {
	response = &DescribeTicketsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeTickets 查询当前团队提交的工单列表。
func (c *Client) DescribeTickets(request *DescribeTicketsRequest) (response *DescribeTicketsResponse, err error) {
	response = NewDescribeTicketsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateTicketRequest() (request *CreateTicketRequest) {
	request = &CreateTicketRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateTicket")

	return
}

func NewCreateTicketResponse() (response *CreateTicketResponse) {
	response = &CreateTicketResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateTicket 创建一个工单。
func (c *Client) CreateTicket(request *CreateTicketRequest) (response *CreateTicketResponse, err error) {
	response = NewCreateTicketResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeTicketDetailRequest() (request *DescribeTicketDetailRequest) {
	request = &DescribeTicketDetailRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeTicketDetail")

	return
}

func NewDescribeTicketDetailResponse() (response *DescribeTicketDetailResponse) {
	response = &DescribeTicketDetailResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeTicketDetail 查询指定工单的详细信息与沟通记录。
func (c *Client) DescribeTicketDetail(request *DescribeTicketDetailRequest) (response *DescribeTicketDetailResponse, err error) {
	response = NewDescribeTicketDetailResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeTicketCreationQuotaRequest() (request *DescribeTicketCreationQuotaRequest) {
	request = &DescribeTicketCreationQuotaRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeTicketCreationQuota")

	return
}

func NewDescribeTicketCreationQuotaResponse() (response *DescribeTicketCreationQuotaResponse) {
	response = &DescribeTicketCreationQuotaResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeTicketCreationQuota 查询当前团队一小时内是否还可以创建工单。
func (c *Client) DescribeTicketCreationQuota(request *DescribeTicketCreationQuotaRequest) (response *DescribeTicketCreationQuotaResponse, err error) {
	response = NewDescribeTicketCreationQuotaResponse()
	err = c.ApiCall(request, response)
	return
}

