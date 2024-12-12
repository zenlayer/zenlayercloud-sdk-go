package user

import (
	"gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2024-05-29"
	SERVICE    = "user"
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

func NewDescribeResourceGroupsRequest() (request *DescribeResourceGroupsRequest) {
	request = &DescribeResourceGroupsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeResourceGroups")

	return
}

// DescribeResourceGroups
// 调用本接口用于查询一个或多个资源组的信息。
// Possible error codes to return:
func (c *Client) DescribeResourceGroups(request *DescribeResourceGroupsRequest) (response *DescribeResourceGroupsResponse, err error) {
	response = NewDescribeResourceGroupsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeResourceGroupsResponse() (response *DescribeResourceGroupsResponse) {
	response = &DescribeResourceGroupsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeResourceTypesRequest() (request *DescribeResourceTypesRequest) {
	request = &DescribeResourceTypesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeResourceTypes")

	return
}

// DescribeResourceTypes
// 调用本接口用于查询全量的资源类型的信息。
// Possible error codes to return:
func (c *Client) DescribeResourceTypes(request *DescribeResourceTypesRequest) (response *DescribeResourceTypesResponse, err error) {
	response = NewDescribeResourceTypesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeResourceTypesResponse() (response *DescribeResourceTypesResponse) {
	response = &DescribeResourceTypesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeResourcesByGroupRequest() (request *DescribeResourcesByGroupRequest) {
	request = &DescribeResourcesByGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeResourcesByGroup")

	return
}

// DescribeResourcesByGroup
// 调用本接口用于查询某个资源组下资源的信息。
// Possible error codes to return:
func (c *Client) DescribeResourcesByGroup(request *DescribeResourcesByGroupRequest) (response *DescribeResourcesByGroupResponse, err error) {
	response = NewDescribeResourcesByGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeResourcesByGroupResponse() (response *DescribeResourcesByGroupResponse) {
	response = &DescribeResourcesByGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}
