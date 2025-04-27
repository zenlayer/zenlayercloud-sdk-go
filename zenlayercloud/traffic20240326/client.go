/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package traffic

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2024-03-26"
	SERVICE    = "traffic"
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


func NewDescribeBandwidthClustersRequest() (request *DescribeBandwidthClustersRequest) {
	request = &DescribeBandwidthClustersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBandwidthClusters")

	return
}

func NewDescribeBandwidthClustersResponse() (response *DescribeBandwidthClustersResponse) {
	response = &DescribeBandwidthClustersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBandwidthClusters 查询一个或多个共享带宽包的信息。用户可以根据共享带宽包ID、名称或者城市名称等条件来查询共享带宽包的详细信息。
func (c *Client) DescribeBandwidthClusters(request *DescribeBandwidthClustersRequest) (response *DescribeBandwidthClustersResponse, err error) {
	response = NewDescribeBandwidthClustersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBandwidthClusterAreasRequest() (request *DescribeBandwidthClusterAreasRequest) {
	request = &DescribeBandwidthClusterAreasRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBandwidthClusterAreas")

	return
}

func NewDescribeBandwidthClusterAreasResponse() (response *DescribeBandwidthClusterAreasResponse) {
	response = &DescribeBandwidthClusterAreasResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBandwidthClusterAreas 查询共享带宽包的地区信息。
func (c *Client) DescribeBandwidthClusterAreas(request *DescribeBandwidthClusterAreasRequest) (response *DescribeBandwidthClusterAreasResponse, err error) {
	response = NewDescribeBandwidthClusterAreasResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBandwidthClusterTrafficRequest() (request *DescribeBandwidthClusterTrafficRequest) {
	request = &DescribeBandwidthClusterTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBandwidthClusterTraffic")

	return
}

func NewDescribeBandwidthClusterTrafficResponse() (response *DescribeBandwidthClusterTrafficResponse) {
	response = &DescribeBandwidthClusterTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBandwidthClusterTraffic 查询带宽组指定时间段内的流量信息。
func (c *Client) DescribeBandwidthClusterTraffic(request *DescribeBandwidthClusterTrafficRequest) (response *DescribeBandwidthClusterTrafficResponse, err error) {
	response = NewDescribeBandwidthClusterTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryBandwidthClusterPriceRequest() (request *InquiryBandwidthClusterPriceRequest) {
	request = &InquiryBandwidthClusterPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryBandwidthClusterPrice")

	return
}

func NewInquiryBandwidthClusterPriceResponse() (response *InquiryBandwidthClusterPriceResponse) {
	response = &InquiryBandwidthClusterPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryBandwidthClusterPrice 查询共享带宽包价格。
func (c *Client) InquiryBandwidthClusterPrice(request *InquiryBandwidthClusterPriceRequest) (response *InquiryBandwidthClusterPriceResponse, err error) {
	response = NewInquiryBandwidthClusterPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateBandwidthClusterRequest() (request *CreateBandwidthClusterRequest) {
	request = &CreateBandwidthClusterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateBandwidthCluster")

	return
}

func NewCreateBandwidthClusterResponse() (response *CreateBandwidthClusterResponse) {
	response = &CreateBandwidthClusterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateBandwidthCluster 创建一个共享带宽包。
func (c *Client) CreateBandwidthCluster(request *CreateBandwidthClusterRequest) (response *CreateBandwidthClusterResponse, err error) {
	response = NewCreateBandwidthClusterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteBandwidthClustersRequest() (request *DeleteBandwidthClustersRequest) {
	request = &DeleteBandwidthClustersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteBandwidthClusters")

	return
}

func NewDeleteBandwidthClustersResponse() (response *DeleteBandwidthClustersResponse) {
	response = &DeleteBandwidthClustersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteBandwidthClusters 删除一个或多个共享带宽包。
func (c *Client) DeleteBandwidthClusters(request *DeleteBandwidthClustersRequest) (response *DeleteBandwidthClustersResponse, err error) {
	response = NewDeleteBandwidthClustersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUpdateBandwidthClusterCommitBandwidthRequest() (request *UpdateBandwidthClusterCommitBandwidthRequest) {
	request = &UpdateBandwidthClusterCommitBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UpdateBandwidthClusterCommitBandwidth")

	return
}

func NewUpdateBandwidthClusterCommitBandwidthResponse() (response *UpdateBandwidthClusterCommitBandwidthResponse) {
	response = &UpdateBandwidthClusterCommitBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UpdateBandwidthClusterCommitBandwidth 修改带宽包的保底带宽。
func (c *Client) UpdateBandwidthClusterCommitBandwidth(request *UpdateBandwidthClusterCommitBandwidthRequest) (response *UpdateBandwidthClusterCommitBandwidthResponse, err error) {
	response = NewUpdateBandwidthClusterCommitBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

