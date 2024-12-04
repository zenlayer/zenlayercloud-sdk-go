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

func NewCreateVpcRequest() (request *CreateVpcRequest) {
	request = &CreateVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateVpc")

	return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
	response = &CreateVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
	response = NewCreateVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteVpcRequest() (request *DeleteVpcRequest) {
	request = &DeleteVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteVpc")

	return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
	response = &DeleteVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
	response = NewDeleteVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyVpcsAttributeRequest() (request *ModifyVpcsAttributeRequest) {
	request = &ModifyVpcsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyVpcsAttribute")

	return
}

func NewModifyVpcsAttributeResponse() (response *ModifyVpcsAttributeResponse) {
	response = &ModifyVpcsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVpcsAttribute(request *ModifyVpcsAttributeRequest) (response *ModifyVpcsAttributeResponse, err error) {
	response = NewModifyVpcsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
	request = &DescribeVpcsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeVpcs")

	return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
	response = &DescribeVpcsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
	response = NewDescribeVpcsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSubnetRegionsRequest() (request *DescribeSubnetRegionsRequest) {
	request = &DescribeSubnetRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSubnetRegions")

	return
}

func NewDescribeSubnetRegionsResponse() (response *DescribeSubnetRegionsResponse) {
	response = &DescribeSubnetRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubnetRegions(request *DescribeSubnetRegionsRequest) (response *DescribeSubnetRegionsResponse, err error) {
	response = NewDescribeSubnetRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
	request = &CreateSubnetRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateSubnet")

	return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
	response = &CreateSubnetResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
	response = NewCreateSubnetResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
	request = &DeleteSubnetRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteSubnet")

	return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
	response = &DeleteSubnetResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
	response = NewDeleteSubnetResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySubnetsAttributeRequest() (request *ModifySubnetsAttributeRequest) {
	request = &ModifySubnetsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySubnetsAttribute")

	return
}

func NewModifySubnetsAttributeResponse() (response *ModifySubnetsAttributeResponse) {
	response = &ModifySubnetsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifySubnetsAttribute(request *ModifySubnetsAttributeRequest) (response *ModifySubnetsAttributeResponse, err error) {
	response = NewModifySubnetsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
	request = &DescribeSubnetsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSubnets")

	return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
	response = &DescribeSubnetsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
	response = NewDescribeSubnetsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySubnetStackTypeRequest() (request *ModifySubnetStackTypeRequest) {
	request = &ModifySubnetStackTypeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySubnetStackType")

	return
}

func NewModifySubnetStackTypeResponse() (response *ModifySubnetStackTypeResponse) {
	response = &ModifySubnetStackTypeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifySubnetStackType(request *ModifySubnetStackTypeRequest) (response *ModifySubnetStackTypeResponse, err error) {
	response = NewModifySubnetStackTypeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
	request = &CreateRouteRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateRoute")

	return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
	response = &CreateRouteResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
	response = NewCreateRouteResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
	request = &DeleteRouteRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteRoute")

	return
}

func NewDeleteRouteResponse() (response *DeleteRouteResponse) {
	response = &DeleteRouteResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DeleteRoute(request *DeleteRouteRequest) (response *DeleteRouteResponse, err error) {
	response = NewDeleteRouteResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeRoutesRequest() (request *DescribeRoutesRequest) {
	request = &DescribeRoutesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeRoutes")

	return
}

func NewDescribeRoutesResponse() (response *DescribeRoutesResponse) {
	response = &DescribeRoutesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeRoutes(request *DescribeRoutesRequest) (response *DescribeRoutesResponse, err error) {
	response = NewDescribeRoutesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstances")

	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	response = NewDescribeInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstancesStatusRequest() (request *DescribeInstancesStatusRequest) {
	request = &DescribeInstancesStatusRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstancesStatus")

	return
}

func NewDescribeInstancesStatusResponse() (response *DescribeInstancesStatusResponse) {
	response = &DescribeInstancesStatusResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstancesStatus(request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
	response = NewDescribeInstancesStatusResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstancesResourceGroupRequest() (request *ModifyInstancesResourceGroupRequest) {
	request = &ModifyInstancesResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstancesResourceGroup")

	return
}

func NewModifyInstancesResourceGroupResponse() (response *ModifyInstancesResourceGroupResponse) {
	response = &ModifyInstancesResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstancesResourceGroup(request *ModifyInstancesResourceGroupRequest) (response *ModifyInstancesResourceGroupResponse, err error) {
	response = NewModifyInstancesResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
	request = &ModifyInstancesAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstancesAttribute")

	return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
	response = &ModifyInstancesAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
	response = NewModifyInstancesAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
	request = &RebootInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RebootInstances")

	return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
	response = &RebootInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	response = NewRebootInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
	request = &ResetInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ResetInstance")

	return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
	response = &ResetInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
	response = NewResetInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
	request = &ResetInstancesPasswordRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ResetInstancesPassword")

	return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
	response = &ResetInstancesPasswordResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
	response = NewResetInstancesPasswordResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StartInstances")

	return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
	response = &StartInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	response = NewStartInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StopInstances")

	return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
	response = &StopInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	response = NewStopInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrsRequest() (request *DescribeCidrsRequest) {
	request = &DescribeCidrsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrs")

	return
}

func NewDescribeCidrsResponse() (response *DescribeCidrsResponse) {
	response = &DescribeCidrsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCidrs(request *DescribeCidrsRequest) (response *DescribeCidrsResponse, err error) {
	response = NewDescribeCidrsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrRegionsRequest() (request *DescribeCidrRegionsRequest) {
	request = &DescribeCidrRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrRegions")

	return
}

func NewDescribeCidrRegionsResponse() (response *DescribeCidrRegionsResponse) {
	response = &DescribeCidrRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCidrRegions(request *DescribeCidrRegionsRequest) (response *DescribeCidrRegionsResponse, err error) {
	response = NewDescribeCidrRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeOwnCidrPriceRequest() (request *DescribeOwnCidrPriceRequest) {
	request = &DescribeOwnCidrPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeOwnCidrPrice")

	return
}

func NewDescribeOwnCidrPriceResponse() (response *DescribeOwnCidrPriceResponse) {
	response = &DescribeOwnCidrPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeOwnCidrPrice(request *DescribeOwnCidrPriceRequest) (response *DescribeOwnCidrPriceResponse, err error) {
	response = NewDescribeOwnCidrPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrPriceRequest() (request *DescribeCidrPriceRequest) {
	request = &DescribeCidrPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrPrice")

	return
}

func NewDescribeCidrPriceResponse() (response *DescribeCidrPriceResponse) {
	response = &DescribeCidrPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCidrPrice(request *DescribeCidrPriceRequest) (response *DescribeCidrPriceResponse, err error) {
	response = NewDescribeCidrPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrUsedIpsRequest() (request *DescribeCidrUsedIpsRequest) {
	request = &DescribeCidrUsedIpsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrUsedIps")

	return
}

func NewDescribeCidrUsedIpsResponse() (response *DescribeCidrUsedIpsResponse) {
	response = &DescribeCidrUsedIpsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeCidrUsedIps(request *DescribeCidrUsedIpsRequest) (response *DescribeCidrUsedIpsResponse, err error) {
	response = NewDescribeCidrUsedIpsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateOwnCidrRequest() (request *CreateOwnCidrRequest) {
	request = &CreateOwnCidrRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateOwnCidr")

	return
}

func NewCreateOwnCidrResponse() (response *CreateOwnCidrResponse) {
	response = &CreateOwnCidrResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateOwnCidr(request *CreateOwnCidrRequest) (response *CreateOwnCidrResponse, err error) {
	response = NewCreateOwnCidrResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateCidrRequest() (request *CreateCidrRequest) {
	request = &CreateCidrRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateCidr")

	return
}

func NewCreateCidrResponse() (response *CreateCidrResponse) {
	response = &CreateCidrResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateCidr(request *CreateCidrRequest) (response *CreateCidrResponse, err error) {
	response = NewCreateCidrResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteCidrRequest() (request *DeleteCidrRequest) {
	request = &DeleteCidrRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteCidr")

	return
}

func NewDeleteCidrResponse() (response *DeleteCidrResponse) {
	response = &DeleteCidrResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DeleteCidr(request *DeleteCidrRequest) (response *DeleteCidrResponse, err error) {
	response = NewDeleteCidrResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewCidrRequest() (request *RenewCidrRequest) {
	request = &RenewCidrRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewCidr")

	return
}

func NewRenewCidrResponse() (response *RenewCidrResponse) {
	response = &RenewCidrResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) RenewCidr(request *RenewCidrRequest) (response *RenewCidrResponse, err error) {
	response = NewRenewCidrResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnAssignCidrIpRequest() (request *UnAssignCidrIpRequest) {
	request = &UnAssignCidrIpRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnAssignCidrIp")

	return
}

func NewUnAssignCidrIpResponse() (response *UnAssignCidrIpResponse) {
	response = &UnAssignCidrIpResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) UnAssignCidrIp(request *UnAssignCidrIpRequest) (response *UnAssignCidrIpResponse, err error) {
	response = NewUnAssignCidrIpResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssignCidrIpRequest() (request *AssignCidrIpRequest) {
	request = &AssignCidrIpRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignCidrIp")

	return
}

func NewAssignCidrIpResponse() (response *AssignCidrIpResponse) {
	response = &AssignCidrIpResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) AssignCidrIp(request *AssignCidrIpRequest) (response *AssignCidrIpResponse, err error) {
	response = NewAssignCidrIpResponse()
	err = c.ApiCall(request, response)
	return
}

func NewBatchAssignCidrIpRequest() (request *BatchAssignCidrIpRequest) {
	request = &BatchAssignCidrIpRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "BatchAssignCidrIp")

	return
}

func NewBatchAssignCidrIpResponse() (response *BatchAssignCidrIpResponse) {
	response = &BatchAssignCidrIpResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) BatchAssignCidrIp(request *BatchAssignCidrIpRequest) (response *BatchAssignCidrIpResponse, err error) {
	response = NewBatchAssignCidrIpResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAvailableCidrIpRequest() (request *AvailableCidrIpRequest) {
	request = &AvailableCidrIpRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AvailableCidrIp")

	return
}

func NewAvailableCidrIpResponse() (response *AvailableCidrIpResponse) {
	response = &AvailableCidrIpResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) AvailableCidrIp(request *AvailableCidrIpRequest) (response *AvailableCidrIpResponse, err error) {
	response = NewAvailableCidrIpResponse()
	err = c.ApiCall(request, response)
	return
}

func NewConfigEgressIpRequest() (request *ConfigEgressIpRequest) {
	request = &ConfigEgressIpRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ConfigEgressIp")

	return
}

func NewConfigEgressIpResponse() (response *ConfigEgressIpResponse) {
	response = &ConfigEgressIpResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ConfigEgressIp(request *ConfigEgressIpRequest) (response *ConfigEgressIpResponse, err error) {
	response = NewConfigEgressIpResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeOwnCidrsRequest() (request *DescribeOwnCidrsRequest) {
	request = &DescribeOwnCidrsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeOwnCidrs")

	return
}

func NewDescribeOwnCidrsResponse() (response *DescribeOwnCidrsResponse) {
	response = &DescribeOwnCidrsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeOwnCidrs(request *DescribeOwnCidrsRequest) (response *DescribeOwnCidrsResponse, err error) {
	response = NewDescribeOwnCidrsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAvailableLanIpRequest() (request *AvailableLanIpRequest) {
	request = &AvailableLanIpRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AvailableLanIp")

	return
}

func NewAvailableLanIpResponse() (response *AvailableLanIpResponse) {
	response = &AvailableLanIpResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) AvailableLanIp(request *AvailableLanIpRequest) (response *AvailableLanIpResponse, err error) {
	response = NewAvailableLanIpResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDiskRegionsRequest() (request *DescribeDiskRegionsRequest) {
	request = &DescribeDiskRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDiskRegions")

	return
}

func NewDescribeDiskRegionsResponse() (response *DescribeDiskRegionsResponse) {
	response = &DescribeDiskRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDiskRegions(request *DescribeDiskRegionsRequest) (response *DescribeDiskRegionsResponse, err error) {
	response = NewDescribeDiskRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateDisksRequest() (request *CreateDisksRequest) {
	request = &CreateDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateDisks")

	return
}

func NewCreateDisksResponse() (response *CreateDisksResponse) {
	response = &CreateDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateDisks(request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
	response = NewCreateDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDisks")

	return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
	response = &DescribeDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	response = NewDescribeDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
	request = &AttachDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AttachDisks")

	return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
	response = &AttachDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
	response = NewAttachDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewChangeDisksAttachRequest() (request *ChangeDisksAttachRequest) {
	request = &ChangeDisksAttachRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ChangeDisksAttach")

	return
}

func NewChangeDisksAttachResponse() (response *ChangeDisksAttachResponse) {
	response = &ChangeDisksAttachResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ChangeDisksAttach(request *ChangeDisksAttachRequest) (response *ChangeDisksAttachResponse, err error) {
	response = NewChangeDisksAttachResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
	request = &DetachDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DetachDisks")

	return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
	response = &DetachDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
	response = NewDetachDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyDisksAttributesRequest() (request *ModifyDisksAttributesRequest) {
	request = &ModifyDisksAttributesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyDisksAttributes")

	return
}

func NewModifyDisksAttributesResponse() (response *ModifyDisksAttributesResponse) {
	response = &ModifyDisksAttributesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDisksAttributes(request *ModifyDisksAttributesRequest) (response *ModifyDisksAttributesResponse, err error) {
	response = NewModifyDisksAttributesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateDisksRequest() (request *InquiryPriceCreateDisksRequest) {
	request = &InquiryPriceCreateDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateDisks")

	return
}

func NewInquiryPriceCreateDisksResponse() (response *InquiryPriceCreateDisksResponse) {
	response = &InquiryPriceCreateDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) InquiryPriceCreateDisks(request *InquiryPriceCreateDisksRequest) (response *InquiryPriceCreateDisksResponse, err error) {
	response = NewInquiryPriceCreateDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminateDiskRequest() (request *TerminateDiskRequest) {
	request = &TerminateDiskRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminateDisk")

	return
}

func NewTerminateDiskResponse() (response *TerminateDiskResponse) {
	response = &TerminateDiskResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) TerminateDisk(request *TerminateDiskRequest) (response *TerminateDiskResponse, err error) {
	response = NewTerminateDiskResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseDiskRequest() (request *ReleaseDiskRequest) {
	request = &ReleaseDiskRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseDisk")

	return
}

func NewReleaseDiskResponse() (response *ReleaseDiskResponse) {
	response = &ReleaseDiskResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ReleaseDisk(request *ReleaseDiskRequest) (response *ReleaseDiskResponse, err error) {
	response = NewReleaseDiskResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewDiskRequest() (request *RenewDiskRequest) {
	request = &RenewDiskRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewDisk")

	return
}

func NewRenewDiskResponse() (response *RenewDiskResponse) {
	response = &RenewDiskResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) RenewDisk(request *RenewDiskRequest) (response *RenewDiskResponse, err error) {
	response = NewRenewDiskResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyDisksResourceGroupRequest() (request *ModifyDisksResourceGroupRequest) {
	request = &ModifyDisksResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyDisksResourceGroup")

	return
}

func NewModifyDisksResourceGroupResponse() (response *ModifyDisksResourceGroupResponse) {
	response = &ModifyDisksResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyDisksResourceGroup(request *ModifyDisksResourceGroupRequest) (response *ModifyDisksResourceGroupResponse, err error) {
	response = NewModifyDisksResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDiskCategoryRequest() (request *DescribeDiskCategoryRequest) {
	request = &DescribeDiskCategoryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDiskCategory")

	return
}

func NewDescribeDiskCategoryResponse() (response *DescribeDiskCategoryResponse) {
	response = &DescribeDiskCategoryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeDiskCategory(request *DescribeDiskCategoryRequest) (response *DescribeDiskCategoryResponse, err error) {
	response = NewDescribeDiskCategoryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
	request = &DescribeSecurityGroupsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSecurityGroups")

	return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
	response = &DescribeSecurityGroupsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
	response = NewDescribeSecurityGroupsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySecurityGroupsAttributeRequest() (request *ModifySecurityGroupsAttributeRequest) {
	request = &ModifySecurityGroupsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySecurityGroupsAttribute")

	return
}

func NewModifySecurityGroupsAttributeResponse() (response *ModifySecurityGroupsAttributeResponse) {
	response = &ModifySecurityGroupsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifySecurityGroupsAttribute(request *ModifySecurityGroupsAttributeRequest) (response *ModifySecurityGroupsAttributeResponse, err error) {
	response = NewModifySecurityGroupsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSecurityGroupRuleRequest() (request *DescribeSecurityGroupRuleRequest) {
	request = &DescribeSecurityGroupRuleRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSecurityGroupRule")

	return
}

func NewDescribeSecurityGroupRuleResponse() (response *DescribeSecurityGroupRuleResponse) {
	response = &DescribeSecurityGroupRuleResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSecurityGroupRule(request *DescribeSecurityGroupRuleRequest) (response *DescribeSecurityGroupRuleResponse, err error) {
	response = NewDescribeSecurityGroupRuleResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
	request = &CreateSecurityGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateSecurityGroup")

	return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
	response = &CreateSecurityGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
	response = NewCreateSecurityGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
	request = &DeleteSecurityGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteSecurityGroup")

	return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
	response = &DeleteSecurityGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
	response = NewDeleteSecurityGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewConfigureSecurityGroupRulesRequest() (request *ConfigureSecurityGroupRulesRequest) {
	request = &ConfigureSecurityGroupRulesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ConfigureSecurityGroupRules")

	return
}

func NewConfigureSecurityGroupRulesResponse() (response *ConfigureSecurityGroupRulesResponse) {
	response = &ConfigureSecurityGroupRulesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ConfigureSecurityGroupRules(request *ConfigureSecurityGroupRulesRequest) (response *ConfigureSecurityGroupRulesResponse, err error) {
	response = NewConfigureSecurityGroupRulesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssignSecurityGroupVpcRequest() (request *AssignSecurityGroupVpcRequest) {
	request = &AssignSecurityGroupVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignSecurityGroupVpc")

	return
}

func NewAssignSecurityGroupVpcResponse() (response *AssignSecurityGroupVpcResponse) {
	response = &AssignSecurityGroupVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) AssignSecurityGroupVpc(request *AssignSecurityGroupVpcRequest) (response *AssignSecurityGroupVpcResponse, err error) {
	response = NewAssignSecurityGroupVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnAssignSecurityGroupVpcRequest() (request *UnAssignSecurityGroupVpcRequest) {
	request = &UnAssignSecurityGroupVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnAssignSecurityGroupVpc")

	return
}

func NewUnAssignSecurityGroupVpcResponse() (response *UnAssignSecurityGroupVpcResponse) {
	response = &UnAssignSecurityGroupVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) UnAssignSecurityGroupVpc(request *UnAssignSecurityGroupVpcRequest) (response *UnAssignSecurityGroupVpcResponse, err error) {
	response = NewUnAssignSecurityGroupVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeNetworkInterfacesRequest() (request *DescribeNicsRequest) {
	request = &DescribeNicsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeNetworkInterfaces")

	return
}

func NewDescribeNetworkInterfacesResponse() (response *DescribeNicsResponse) {
	response = &DescribeNicsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNetworkInterfaces(request *DescribeNicsRequest) (response *DescribeNicsResponse, err error) {
	response = NewDescribeNetworkInterfacesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyNetworkInterfacesAttributeRequest() (request *ModifyNicsAttributeRequest) {
	request = &ModifyNicsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyNetworkInterfacesAttribute")

	return
}

func NewModifyNetworkInterfacesAttributeResponse() (response *ModifyNicsAttributeResponse) {
	response = &ModifyNicsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyNetworkInterfacesAttribute(request *ModifyNicsAttributeRequest) (response *ModifyNicsAttributeResponse, err error) {
	response = NewModifyNetworkInterfacesAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateNetworkInterfaceRequest() (request *CreateNicRequest) {
	request = &CreateNicRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateNetworkInterface")

	return
}

func NewCreateNetworkInterfaceResponse() (response *CreateNicResponse) {
	response = &CreateNicResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateNetworkInterface(request *CreateNicRequest) (response *CreateNicResponse, err error) {
	response = NewCreateNetworkInterfaceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteNetworkInterfaceRequest() (request *DeleteNicRequest) {
	request = &DeleteNicRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteNetworkInterface")

	return
}

func NewDeleteNetworkInterfaceResponse() (response *DeleteNicResponse) {
	response = &DeleteNicResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DeleteNetworkInterface(request *DeleteNicRequest) (response *DeleteNicResponse, err error) {
	response = NewDeleteNetworkInterfaceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAttachNetworkInterfaceRequest() (request *AttachNicRequest) {
	request = &AttachNicRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AttachNetworkInterface")

	return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNicResponse) {
	response = &AttachNicResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) AttachNetworkInterface(request *AttachNicRequest) (response *AttachNicResponse, err error) {
	response = NewAttachNetworkInterfaceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssignNetworkInterfaceIpv6Request() (request *AssignNicIpv6Request) {
	request = &AssignNicIpv6Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignNetworkInterfaceIpv6")

	return
}

func NewAssignNetworkInterfaceIpv6Response() (response *AssignNicIpv6Response) {
	response = &AssignNicIpv6Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) AssignNetworkInterfaceIpv6(request *AssignNicIpv6Request) (response *AssignNicIpv6Response, err error) {
	response = NewAssignNetworkInterfaceIpv6Response()
	err = c.ApiCall(request, response)
	return
}

func NewUnassignNetworkInterfaceIpv4Request() (request *UnAssignNicIpv4Request) {
	request = &UnAssignNicIpv4Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnassignNetworkInterfaceIpv4")

	return
}

func NewUnassignNetworkInterfaceIpv4Response() (response *UnAssignNicIpv4Response) {
	response = &UnAssignNicIpv4Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) UnassignNetworkInterfaceIpv4(request *UnAssignNicIpv4Request) (response *UnAssignNicIpv4Response, err error) {
	response = NewUnassignNetworkInterfaceIpv4Response()
	err = c.ApiCall(request, response)
	return
}

func NewAssignNetworkInterfaceIpv4Request() (request *AssignNicIpv4Request) {
	request = &AssignNicIpv4Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignNetworkInterfaceIpv4")

	return
}

func NewAssignNetworkInterfaceIpv4Response() (response *AssignNicIpv4Response) {
	response = &AssignNicIpv4Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) AssignNetworkInterfaceIpv4(request *AssignNicIpv4Request) (response *AssignNicIpv4Response, err error) {
	response = NewAssignNetworkInterfaceIpv4Response()
	err = c.ApiCall(request, response)
	return
}

func NewBatchAssignNetworkInterfaceIpv4Request() (request *BatchAssignNicIpv4Request) {
	request = &BatchAssignNicIpv4Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "BatchAssignNetworkInterfaceIpv4")

	return
}

func NewBatchAssignNetworkInterfaceIpv4Response() (response *BatchAssignNicIpv4Response) {
	response = &BatchAssignNicIpv4Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) BatchAssignNetworkInterfaceIpv4(request *BatchAssignNicIpv4Request) (response *BatchAssignNicIpv4Response, err error) {
	response = NewBatchAssignNetworkInterfaceIpv4Response()
	err = c.ApiCall(request, response)
	return
}

func NewDetachNetworkInterfaceRequest() (request *DetachNicRequest) {
	request = &DetachNicRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DetachNetworkInterface")

	return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNicResponse) {
	response = &DetachNicResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DetachNetworkInterface(request *DetachNicRequest) (response *DetachNicResponse, err error) {
	response = NewDetachNetworkInterfaceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeNetworkInterfaceRegionsRequest() (request *DescribeNicRegionsRequest) {
	request = &DescribeNicRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeNetworkInterfaceRegions")

	return
}

func NewDescribeNetworkInterfaceRegionsResponse() (response *DescribeNicRegionsResponse) {
	response = &DescribeNicRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeNetworkInterfaceRegions(request *DescribeNicRegionsRequest) (response *DescribeNicRegionsResponse, err error) {
	response = NewDescribeNetworkInterfaceRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPricePublicIpv6Request() (request *InquiryPricePublicIpv6Request) {
	request = &InquiryPricePublicIpv6Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPricePublicIpv6")

	return
}

func NewInquiryPricePublicIpv6Response() (response *InquiryPricePublicIpv6Response) {
	response = &InquiryPricePublicIpv6Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) InquiryPricePublicIpv6(request *InquiryPricePublicIpv6Request) (response *InquiryPricePublicIpv6Response, err error) {
	response = NewInquiryPricePublicIpv6Response()
	err = c.ApiCall(request, response)
	return
}

func NewCreateZecInstancesRequest() (request *CreateZecInstancesRequest) {
	request = &CreateZecInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateZecInstances")

	return
}

func NewCreateZecInstancesResponse() (response *CreateZecInstancesResponse) {
	response = &CreateZecInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateZecInstances(request *CreateZecInstancesRequest) (response *CreateZecInstancesResponse, err error) {
	response = NewCreateZecInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
	request = &CreateInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateInstances")

	return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
	response = &CreateInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
	response = NewCreateInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeZoneInstanceConfigInfosRequest() (request *DescribeZoneInstanceConfigInfosRequest) {
	request = &DescribeZoneInstanceConfigInfosRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeZoneInstanceConfigInfos")

	return
}

func NewDescribeZoneInstanceConfigInfosResponse() (response *DescribeZoneInstanceConfigInfosResponse) {
	response = &DescribeZoneInstanceConfigInfosResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
	response = NewDescribeZoneInstanceConfigInfosResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
	request = &DescribeZonesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeZones")

	return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
	response = &DescribeZonesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	response = NewDescribeZonesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeImages")

	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	response = NewDescribeImagesResponse()
	err = c.ApiCall(request, response)
	return
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

func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
	response = NewDescribeKeyPairsResponse()
	err = c.ApiCall(request, response)
	return
}
