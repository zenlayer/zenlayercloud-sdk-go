package maintenance

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2025-03-10"
	SERVICE    = "maintenance"
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

func NewDescribeMaintenanceAlertsRequest() (request *DescribeMaintenanceAlertsRequest) {
	request = &DescribeMaintenanceAlertsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeMaintenanceAlerts")

	return
}

func NewDescribeMaintenanceAlertsResponse() (response *DescribeMaintenanceAlertsResponse) {
	response = &DescribeMaintenanceAlertsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeMaintenanceAlerts(request *DescribeMaintenanceAlertsRequest) (response *DescribeMaintenanceAlertsResponse, err error) {
	response = NewDescribeMaintenanceAlertsResponse()
	err = c.ApiCall(request, response)
	return
}
