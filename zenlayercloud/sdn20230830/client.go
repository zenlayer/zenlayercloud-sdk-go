package sdn

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2023-08-30"
	SERVICE    = "sdn"
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

func NewDescribeDatacentersRequest() (request *DescribeDatacentersRequest) {
	request = &DescribeDatacentersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDatacenters")

	return
}

// DescribeDatacenters
// 本接口用于查询端口列表。
//
// Possible error codes to return:
func (c *Client) DescribeDatacenters(request *DescribeDatacentersRequest) (response *DescribeDatacentersResponse, err error) {
	response = NewDescribeDatacentersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDatacentersResponse() (response *DescribeDatacentersResponse) {
	response = &DescribeDatacentersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewCreatePortRequest() (request *CreatePortRequest) {
	request = &CreatePortRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreatePort")

	return
}

// CreatePort
// 本接口用于创建一个数据中心端口。
//
// Possible error codes to return:
func (c *Client) CreatePort(request *CreatePortRequest) (response *CreatePortResponse, err error) {
	response = NewCreatePortResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreatePortResponse() (response *CreatePortResponse) {
	response = &CreatePortResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeDataCenterPortPriceRequest() (request *DescribeDataCenterPortPriceRequest) {
	request = &DescribeDataCenterPortPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDataCenterPortPrice")

	return
}

// DescribeDataCenterPortPrice
// 本接口用于查询数据中心下可购买的端口类型及其对应的价格。
//
// Possible error codes to return:
func (c *Client) DescribeDataCenterPortPrice(request *DescribeDataCenterPortPriceRequest) (response *DescribeDataCenterPortPriceResponse, err error) {
	response = NewDescribeDataCenterPortPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDataCenterPortPriceResponse() (response *DescribeDataCenterPortPriceResponse) {
	response = &DescribeDataCenterPortPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribePortsRequest() (request *DescribePortsRequest) {
	request = &DescribePortsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePorts")

	return
}

// DescribePorts
// 本接口用于查询端口列表。
//
// Possible error codes to return:
func (c *Client) DescribePorts(request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
	response = NewDescribePortsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePortsResponse() (response *DescribePortsResponse) {
	response = &DescribePortsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribePortTrafficRequest() (request *DescribePortTrafficRequest) {
	request = &DescribePortTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePortTraffic")

	return
}

// DescribePortTraffic
// 本接口用于查询端口流量。
//
// Possible error codes to return:
func (c *Client) DescribePortTraffic(request *DescribePortTrafficRequest) (response *DescribePortTrafficResponse, err error) {
	response = NewDescribePortTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePortTrafficResponse() (response *DescribePortTrafficResponse) {
	response = &DescribePortTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribePortUsableVlanRequest() (request *DescribePortUsableVlanRequest) {
	request = &DescribePortUsableVlanRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePortUsableVlan")

	return
}

// DescribePortUsableVlan
// 本接口用于查询端口可用的VLAN范围。
//
// Possible error codes to return:
func (c *Client) DescribePortUsableVlan(request *DescribePortUsableVlanRequest) (response *DescribePortUsableVlanResponse, err error) {
	response = NewDescribePortUsableVlanResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePortUsableVlanResponse() (response *DescribePortUsableVlanResponse) {
	response = &DescribePortUsableVlanResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDestroyPortRequest() (request *DestroyPortRequest) {
	request = &DestroyPortRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DestroyPort")

	return
}

// DestroyPort
// 本接口用于销毁端口。
//
// Possible error codes to return:
func (c *Client) DestroyPort(request *DestroyPortRequest) (response *DestroyPortResponse, err error) {
	response = NewDestroyPortResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDestroyPortResponse() (response *DestroyPortResponse) {
	response = &DestroyPortResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyPortAttributeRequest() (request *ModifyPortAttributeRequest) {
	request = &ModifyPortAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPortAttribute")

	return
}

// ModifyPortAttribute
// 本接口用于修改端口属性。
//
// Possible error codes to return:
func (c *Client) ModifyPortAttribute(request *ModifyPortAttributeRequest) (response *ModifyPortAttributeResponse, err error) {
	response = NewModifyPortAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPortAttributeResponse() (response *ModifyPortAttributeResponse) {
	response = &ModifyPortAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewTerminatePortRequest() (request *TerminatePortRequest) {
	request = &TerminatePortRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminatePort")

	return
}

// TerminatePort
// 本接口用于终止端口。
//
// Possible error codes to return:
func (c *Client) TerminatePort(request *TerminatePortRequest) (response *TerminatePortResponse, err error) {
	response = NewTerminatePortResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminatePortResponse() (response *TerminatePortResponse) {
	response = &TerminatePortResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewRenewPortRequest() (request *RenewPortRequest) {
	request = &RenewPortRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewPort")

	return
}

// RenewPort
// 调用本接口用于恢复一个数据中心端口。
//
// Possible error codes to return:
func (c *Client) RenewPort(request *RenewPortRequest) (response *RenewPortResponse, err error) {
	response = NewRenewPortResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewPortResponse() (response *RenewPortResponse) {
	response = &RenewPortResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribePrivateConnectsRequest() (request *DescribePrivateConnectsRequest) {
	request = &DescribePrivateConnectsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePrivateConnects")

	return
}

// DescribePrivateConnects
// 调用本接口用于查询一条或多条指定二层网络专线的信息。用户可以根据 ID、名称等信息来进行搜索。
//
// Possible error codes to return:
func (c *Client) DescribePrivateConnects(request *DescribePrivateConnectsRequest) (response *DescribePrivateConnectsResponse, err error) {
	response = NewDescribePrivateConnectsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePrivateConnectsResponse() (response *DescribePrivateConnectsResponse) {
	response = &DescribePrivateConnectsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribePrivateConnectAvailablePortsRequest() (request *DescribePrivateConnectAvailablePortsRequest) {
	request = &DescribePrivateConnectAvailablePortsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePrivateConnectAvailablePorts")

	return
}

// DescribePrivateConnectAvailablePorts
// 调用本接口用于查询可用于创建二层网络专线的子网列表。
//
// Possible error codes to return:
func (c *Client) DescribePrivateConnectAvailablePorts(request *DescribePrivateConnectAvailablePortsRequest) (response *DescribePrivateConnectAvailablePortsResponse, err error) {
	response = NewDescribePrivateConnectAvailablePortsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePrivateConnectAvailablePortsResponse() (response *DescribePrivateConnectAvailablePortsResponse) {
	response = &DescribePrivateConnectAvailablePortsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewCreatePrivateConnectRequest() (request *CreatePrivateConnectRequest) {
	request = &CreatePrivateConnectRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreatePrivateConnect")

	return
}

// CreatePrivateConnect
// 调用本接口用于创建一条二层网络专线。
//
// Possible error codes to return:
func (c *Client) CreatePrivateConnect(request *CreatePrivateConnectRequest) (response *CreatePrivateConnectResponse, err error) {
	response = NewCreatePrivateConnectResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreatePrivateConnectResponse() (response *CreatePrivateConnectResponse) {
	response = &CreatePrivateConnectResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyPrivateConnectsAttributeRequest() (request *ModifyPrivateConnectsAttributeRequest) {
	request = &ModifyPrivateConnectsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPrivateConnectsAttribute")

	return
}

// ModifyPrivateConnectsAttribute
// 调用本接口用于修改一条或多条二层网络专线的基本信息，包括名称。
//
// Possible error codes to return:
func (c *Client) ModifyPrivateConnectsAttribute(request *ModifyPrivateConnectsAttributeRequest) (response *ModifyPrivateConnectsAttributeResponse, err error) {
	response = NewModifyPrivateConnectsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPrivateConnectsAttributeResponse() (response *ModifyPrivateConnectsAttributeResponse) {
	response = &ModifyPrivateConnectsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewInquiryCreatePrivateConnectPriceRequest() (request *InquiryCreatePrivateConnectPriceRequest) {
	request = &InquiryCreatePrivateConnectPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryCreatePrivateConnectPrice")

	return
}

// InquiryCreatePrivateConnectPrice
// 调用本接口用于创建一条二层网络专线询价。
//
// Possible error codes to return:
func (c *Client) InquiryCreatePrivateConnectPrice(request *InquiryCreatePrivateConnectPriceRequest) (response *InquiryCreatePrivateConnectPriceResponse, err error) {
	response = NewInquiryCreatePrivateConnectPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryCreatePrivateConnectPriceResponse() (response *InquiryCreatePrivateConnectPriceResponse) {
	response = &InquiryCreatePrivateConnectPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDeletePrivateConnectRequest() (request *DeletePrivateConnectRequest) {
	request = &DeletePrivateConnectRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeletePrivateConnect")

	return
}

// DeletePrivateConnect
// 调用本接口用于删除一条二层网络专线。
//
// Possible error codes to return:
func (c *Client) DeletePrivateConnect(request *DeletePrivateConnectRequest) (response *DeletePrivateConnectResponse, err error) {
	response = NewDeletePrivateConnectResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeletePrivateConnectResponse() (response *DeletePrivateConnectResponse) {
	response = &DeletePrivateConnectResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDestroyPrivateConnectRequest() (request *DestroyPrivateConnectRequest) {
	request = &DestroyPrivateConnectRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DestroyPrivateConnect")

	return
}

// DestroyPrivateConnect
// 调用本接口用于彻底删除一条二层网络专线。
//
// Possible error codes to return:
func (c *Client) DestroyPrivateConnect(request *DestroyPrivateConnectRequest) (response *DestroyPrivateConnectResponse, err error) {
	response = NewDestroyPrivateConnectResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDestroyPrivateConnectResponse() (response *DestroyPrivateConnectResponse) {
	response = &DestroyPrivateConnectResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewRenewPrivateConnectRequest() (request *RenewPrivateConnectRequest) {
	request = &RenewPrivateConnectRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewPrivateConnect")

	return
}

// RenewPrivateConnect
// 调用本接口用于恢复一条二层网络专线。
//
// Possible error codes to return:
func (c *Client) RenewPrivateConnect(request *RenewPrivateConnectRequest) (response *RenewPrivateConnectResponse, err error) {
	response = NewRenewPrivateConnectResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewPrivateConnectResponse() (response *RenewPrivateConnectResponse) {
	response = &RenewPrivateConnectResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribePrivateConnectTrafficRequest() (request *DescribePrivateConnectTrafficRequest) {
	request = &DescribePrivateConnectTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePrivateConnectTraffic")

	return
}

// DescribePrivateConnectTraffic
// 调用本接口用于查询二层网络专线在指定时间段内的带宽数据。
//
// Possible error codes to return:
func (c *Client) DescribePrivateConnectTraffic(request *DescribePrivateConnectTrafficRequest) (response *DescribePrivateConnectTrafficResponse, err error) {
	response = NewDescribePrivateConnectTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePrivateConnectTrafficResponse() (response *DescribePrivateConnectTrafficResponse) {
	response = &DescribePrivateConnectTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyPrivateConnectBandwidthRequest() (request *ModifyPrivateConnectBandwidthRequest) {
	request = &ModifyPrivateConnectBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyPrivateConnectBandwidth")

	return
}

// ModifyPrivateConnectBandwidth
// 调用本接口用于修改一条二层网络专线的带宽限速。
//
// Possible error codes to return:
func (c *Client) ModifyPrivateConnectBandwidth(request *ModifyPrivateConnectBandwidthRequest) (response *ModifyPrivateConnectBandwidthResponse, err error) {
	response = NewModifyPrivateConnectBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyPrivateConnectBandwidthResponse() (response *ModifyPrivateConnectBandwidthResponse) {
	response = &ModifyPrivateConnectBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeCloudRoutersRequest() (request *DescribeCloudRoutersRequest) {
	request = &DescribeCloudRoutersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouters")

	return
}

// DescribeCloudRouters
// 调用本接口用于查询一个或多个指定三层网络的信息。用户可以根据 ID、名称等信息来进行搜索。
//
// Possible error codes to return:
func (c *Client) DescribeCloudRouters(request *DescribeCloudRoutersRequest) (response *DescribeCloudRoutersResponse, err error) {
	response = NewDescribeCloudRoutersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRoutersResponse() (response *DescribeCloudRoutersResponse) {
	response = &DescribeCloudRoutersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeCloudRouterAvailableVpcsRequest() (request *DescribeCloudRouterAvailableVpcsRequest) {
	request = &DescribeCloudRouterAvailableVpcsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouterAvailableVpcs")

	return
}

// DescribeCloudRouterAvailableVpcs
// 调用本接口用于查询可用于加入三层网络专线的VPC列表。
//
// Possible error codes to return:
func (c *Client) DescribeCloudRouterAvailableVpcs(request *DescribeCloudRouterAvailableVpcsRequest) (response *DescribeCloudRouterAvailableVpcsResponse, err error) {
	response = NewDescribeCloudRouterAvailableVpcsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRouterAvailableVpcsResponse() (response *DescribeCloudRouterAvailableVpcsResponse) {
	response = &DescribeCloudRouterAvailableVpcsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeCloudRouterAvailablePortsRequest() (request *DescribeCloudRouterAvailablePortsRequest) {
	request = &DescribeCloudRouterAvailablePortsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouterAvailablePorts")

	return
}

// DescribeCloudRouterAvailablePorts
// 调用本接口用于查询可用于加入三层网络的物理端口列表。
//
// Possible error codes to return:
func (c *Client) DescribeCloudRouterAvailablePorts(request *DescribeCloudRouterAvailablePortsRequest) (response *DescribeCloudRouterAvailablePortsResponse, err error) {
	response = NewDescribeCloudRouterAvailablePortsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRouterAvailablePortsResponse() (response *DescribeCloudRouterAvailablePortsResponse) {
	response = &DescribeCloudRouterAvailablePortsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewCreateCloudRouterRequest() (request *CreateCloudRouterRequest) {
	request = &CreateCloudRouterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateCloudRouter")

	return
}

// CreateCloudRouter
// 调用本接口用于创建一条三层网络。
//
// Possible error codes to return:
func (c *Client) CreateCloudRouter(request *CreateCloudRouterRequest) (response *CreateCloudRouterResponse, err error) {
	response = NewCreateCloudRouterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateCloudRouterResponse() (response *CreateCloudRouterResponse) {
	response = &CreateCloudRouterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyCloudRoutersAttributeRequest() (request *ModifyCloudRoutersAttributeRequest) {
	request = &ModifyCloudRoutersAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCloudRoutersAttribute")

	return
}

// ModifyCloudRoutersAttribute
// 调用本接口用于修改一个或多个三层网络的基本信息，包括名称和描述信息。
//
// Possible error codes to return:
func (c *Client) ModifyCloudRoutersAttribute(request *ModifyCloudRoutersAttributeRequest) (response *ModifyCloudRoutersAttributeResponse, err error) {
	response = NewModifyCloudRoutersAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCloudRoutersAttributeResponse() (response *ModifyCloudRoutersAttributeResponse) {
	response = &ModifyCloudRoutersAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewAddCloudRouterEdgePointsRequest() (request *AddCloudRouterEdgePointsRequest) {
	request = &AddCloudRouterEdgePointsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AddCloudRouterEdgePoints")

	return
}

// AddCloudRouterEdgePoints
// 调用本接口用于向一个三层网络增加一个或多个新的边缘连接点。
//
// Possible error codes to return:
func (c *Client) AddCloudRouterEdgePoints(request *AddCloudRouterEdgePointsRequest) (response *AddCloudRouterEdgePointsResponse, err error) {
	response = NewAddCloudRouterEdgePointsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAddCloudRouterEdgePointsResponse() (response *AddCloudRouterEdgePointsResponse) {
	response = &AddCloudRouterEdgePointsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDeleteCloudRouterEdgePointRequest() (request *DeleteCloudRouterEdgePointRequest) {
	request = &DeleteCloudRouterEdgePointRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteCloudRouterEdgePoint")

	return
}

// DeleteCloudRouterEdgePoint
// 调用本接口用于删除三层网络中指定的边缘连接点。
//
// Possible error codes to return:
func (c *Client) DeleteCloudRouterEdgePoint(request *DeleteCloudRouterEdgePointRequest) (response *DeleteCloudRouterEdgePointResponse, err error) {
	response = NewDeleteCloudRouterEdgePointResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteCloudRouterEdgePointResponse() (response *DeleteCloudRouterEdgePointResponse) {
	response = &DeleteCloudRouterEdgePointResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDeleteCloudRouterRequest() (request *DeleteCloudRouterRequest) {
	request = &DeleteCloudRouterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteCloudRouter")

	return
}

// DeleteCloudRouter
// 调用本接口用于删除一个三层网络。
//
// Possible error codes to return:
func (c *Client) DeleteCloudRouter(request *DeleteCloudRouterRequest) (response *DeleteCloudRouterResponse, err error) {
	response = NewDeleteCloudRouterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteCloudRouterResponse() (response *DeleteCloudRouterResponse) {
	response = &DeleteCloudRouterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDestroyCloudRouterRequest() (request *DestroyCloudRouterRequest) {
	request = &DestroyCloudRouterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DestroyCloudRouter")

	return
}

// DestroyCloudRouter
// 调用本接口用于彻底删除三层网络。
//
// Possible error codes to return:
func (c *Client) DestroyCloudRouter(request *DestroyCloudRouterRequest) (response *DestroyCloudRouterResponse, err error) {
	response = NewDestroyCloudRouterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDestroyCloudRouterResponse() (response *DestroyCloudRouterResponse) {
	response = &DestroyCloudRouterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewRenewCloudRouterRequest() (request *RenewCloudRouterRequest) {
	request = &RenewCloudRouterRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewCloudRouter")

	return
}

// RenewCloudRouter
// 调用本接口用于恢复一条三层网络。
//
// Possible error codes to return:
func (c *Client) RenewCloudRouter(request *RenewCloudRouterRequest) (response *RenewCloudRouterResponse, err error) {
	response = NewRenewCloudRouterResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewCloudRouterResponse() (response *RenewCloudRouterResponse) {
	response = &RenewCloudRouterResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyCloudRouterEdgePointBandwidthRequest() (request *ModifyCloudRouterEdgePointBandwidthRequest) {
	request = &ModifyCloudRouterEdgePointBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCloudRouterEdgePointBandwidth")

	return
}

// ModifyCloudRouterEdgePointBandwidth
// 调用本接口用于修改三层网络中边缘连接点的带宽限速。
//
// Possible error codes to return:
func (c *Client) ModifyCloudRouterEdgePointBandwidth(request *ModifyCloudRouterEdgePointBandwidthRequest) (response *ModifyCloudRouterEdgePointBandwidthResponse, err error) {
	response = NewModifyCloudRouterEdgePointBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCloudRouterEdgePointBandwidthResponse() (response *ModifyCloudRouterEdgePointBandwidthResponse) {
	response = &ModifyCloudRouterEdgePointBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeCloudRouteEdgePointTrafficRequest() (request *DescribeCloudRouteEdgePointTrafficRequest) {
	request = &DescribeCloudRouteEdgePointTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCloudRouteEdgePointTraffic")

	return
}

// DescribeCloudRouteEdgePointTraffic
// 调用本接口用于查询三层网络连接点在指定时间段内的带宽数据。
//
// Possible error codes to return:
func (c *Client) DescribeCloudRouteEdgePointTraffic(request *DescribeCloudRouteEdgePointTrafficRequest) (response *DescribeCloudRouteEdgePointTrafficResponse, err error) {
	response = NewDescribeCloudRouteEdgePointTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCloudRouteEdgePointTrafficResponse() (response *DescribeCloudRouteEdgePointTrafficResponse) {
	response = &DescribeCloudRouteEdgePointTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyCloudRouterEdgePointRequest() (request *ModifyCloudRouterEdgePointRequest) {
	request = &ModifyCloudRouterEdgePointRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCloudRouterEdgePoint")

	return
}

// ModifyCloudRouterEdgePoint
// 调用本接口用于修改三层网络中边缘连接点的配置信息，包括BPG连接配置，静态路由等。
//
// Possible error codes to return:
func (c *Client) ModifyCloudRouterEdgePoint(request *ModifyCloudRouterEdgePointRequest) (response *ModifyCloudRouterEdgePointResponse, err error) {
	response = NewModifyCloudRouterEdgePointResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCloudRouterEdgePointResponse() (response *ModifyCloudRouterEdgePointResponse) {
	response = &ModifyCloudRouterEdgePointResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}
