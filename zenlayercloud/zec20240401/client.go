package zec

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2024-04-01"
	SERVICE    = "zec"
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

func NewDescribeInstanceMonitorDataRequest() (request *DescribeInstanceMonitorDataRequest) {
	request = &DescribeInstanceMonitorDataRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceMonitorData")

	return
}

func NewDescribeInstanceMonitorDataResponse() (response *DescribeInstanceMonitorDataResponse) {
	response = &DescribeInstanceMonitorDataResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceMonitorData(request *DescribeInstanceMonitorDataRequest) (response *DescribeInstanceMonitorDataResponse, err error) {
	response = NewDescribeInstanceMonitorDataResponse()
	err = c.ApiCall(request, response)
	return
}
