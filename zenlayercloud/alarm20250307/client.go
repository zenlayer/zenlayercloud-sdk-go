package alarm

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2025-03-07"
	SERVICE    = "alarm"
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

func NewDescribeIpBlockEventsRequest() (request *DescribeIpBlockEventsRequest) {
	request = &DescribeIpBlockEventsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeIpBlockEvents")

	return
}

func NewDescribeIpBlockEventsResponse() (response *DescribeIpBlockEventsResponse) {
	response = &DescribeIpBlockEventsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeIpBlockEvents(request *DescribeIpBlockEventsRequest) (response *DescribeIpBlockEventsResponse, err error) {
	response = NewDescribeIpBlockEventsResponse()
	err = c.ApiCall(request, response)
	return
}
