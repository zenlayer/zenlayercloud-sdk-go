package zls

import (
	"gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2023-08-04"
	SERVICE    = "zls"
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

func NewDescribeLogsRequest() (request *DescribeLogsRequest) {
	request = &DescribeLogsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeLogs")

	return
}

// DescribeLogs
// 本接口（DescribeLogs）用于查询审计日志列表。
//
// Possible error codes to return:
func (c *Client) DescribeLogs(request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
	response = NewDescribeLogsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeLogsResponse() (response *DescribeLogsResponse) {
	response = &DescribeLogsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}
