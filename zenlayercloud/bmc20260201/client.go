/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package bmc

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2026-02-01"
	SERVICE    = "bmc"
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


func NewDescribeSubnetAvailableResourcesRequest() (request *DescribeSubnetAvailableResourcesRequest) {
	request = &DescribeSubnetAvailableResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSubnetAvailableResources")

	return
}

func NewDescribeSubnetAvailableResourcesResponse() (response *DescribeSubnetAvailableResourcesResponse) {
	response = &DescribeSubnetAvailableResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeSubnetAvailableResources 本接口用于查询可创建Subnet资源的可用区。
func (c *Client) DescribeSubnetAvailableResources(request *DescribeSubnetAvailableResourcesRequest) (response *DescribeSubnetAvailableResourcesResponse, err error) {
	response = NewDescribeSubnetAvailableResourcesResponse()
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

// ModifySubnetsAttribute 本接口用于修改Subnet的属性（目前只支持修改Subnet的名称）。
func (c *Client) ModifySubnetsAttribute(request *ModifySubnetsAttributeRequest) (response *ModifySubnetsAttributeResponse, err error) {
	response = NewModifySubnetsAttributeResponse()
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

// CreateSubnet 本接口用于创建一个私有网络Subnet。
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
	response = NewCreateSubnetResponse()
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

// DescribeSubnets 本接口用于查询一台或多台指定Subnet的信息。用户可以根据Subnet ID、VPC ID、 区域、Subnet 名称等信息来搜索Subnet信息。
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
	response = NewDescribeSubnetsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssociateSubnetInstancesRequest() (request *AssociateSubnetInstancesRequest) {
	request = &AssociateSubnetInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssociateSubnetInstances")

	return
}

func NewAssociateSubnetInstancesResponse() (response *AssociateSubnetInstancesResponse) {
	response = &AssociateSubnetInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssociateSubnetInstances 本接口用于将一台或多台实例加入同一个子网并分配内网IP。
func (c *Client) AssociateSubnetInstances(request *AssociateSubnetInstancesRequest) (response *AssociateSubnetInstancesResponse, err error) {
	response = NewAssociateSubnetInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnAssociateSubnetInstanceRequest() (request *UnAssociateSubnetInstanceRequest) {
	request = &UnAssociateSubnetInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnAssociateSubnetInstance")

	return
}

func NewUnAssociateSubnetInstanceResponse() (response *UnAssociateSubnetInstanceResponse) {
	response = &UnAssociateSubnetInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnAssociateSubnetInstance 本接口用于将某子网下一台实例从Subnet中解绑。
func (c *Client) UnAssociateSubnetInstance(request *UnAssociateSubnetInstanceRequest) (response *UnAssociateSubnetInstanceResponse, err error) {
	response = NewUnAssociateSubnetInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssociateVpcSubnetRequest() (request *AssociateVpcSubnetRequest) {
	request = &AssociateVpcSubnetRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssociateVpcSubnet")

	return
}

func NewAssociateVpcSubnetResponse() (response *AssociateVpcSubnetResponse) {
	response = &AssociateVpcSubnetResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssociateVpcSubnet 本接口用于为VPC添加Subnet。
func (c *Client) AssociateVpcSubnet(request *AssociateVpcSubnetRequest) (response *AssociateVpcSubnetResponse, err error) {
	response = NewAssociateVpcSubnetResponse()
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

// DeleteSubnet 本接口用于删除一个Subnet。
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
	response = NewDeleteSubnetResponse()
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

// DescribeZones 查询可用区信息。​
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	response = NewDescribeZonesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
	request = &DescribeInstanceTypesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceTypes")

	return
}

func NewDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
	response = &DescribeInstanceTypesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceTypes 查询实例机型配置。
func (c *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
	response = NewDescribeInstanceTypesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAvailableResourcesRequest() (request *DescribeAvailableResourcesRequest) {
	request = &DescribeAvailableResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAvailableResources")

	return
}

func NewDescribeAvailableResourcesResponse() (response *DescribeAvailableResourcesResponse) {
	response = &DescribeAvailableResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAvailableResources 查询售卖实例和带宽的可用资源。
func (c *Client) DescribeAvailableResources(request *DescribeAvailableResourcesRequest) (response *DescribeAvailableResourcesResponse, err error) {
	response = NewDescribeAvailableResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
	request = &InquiryPriceCreateInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateInstance")

	return
}

func NewInquiryPriceCreateInstanceResponse() (response *InquiryPriceCreateInstanceResponse) {
	response = &InquiryPriceCreateInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateInstance 创建一个实例询价。
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
	response = NewInquiryPriceCreateInstanceResponse()
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

// DescribeImages 查看镜像列表。
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	response = NewDescribeImagesResponse()
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

// CreateInstances 创建一个或多个指定配置的实例。
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
	response = NewCreateInstancesResponse()
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

// DescribeInstances 查询一台或多台实例的信息。用户可以根据实例ID、实例名称或者实例计费模式等条件来查询实例的详细信息。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	response = NewDescribeInstancesResponse()
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

// ModifyInstancesAttribute 本接口用于修改实例的属性（目前只支持修改实例的显示名称）。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
	response = NewModifyInstancesAttributeResponse()
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

// StartInstances 本接口用于启动一个或多个实例。
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

// StopInstances 本接口用于关闭一个或多个实例。
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	response = NewStopInstancesResponse()
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

// RebootInstances 本接口用于重启一个或多个实例。
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	response = NewRebootInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReinstallInstanceRequest() (request *ReinstallInstanceRequest) {
	request = &ReinstallInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReinstallInstance")

	return
}

func NewReinstallInstanceResponse() (response *ReinstallInstanceResponse) {
	response = &ReinstallInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReinstallInstance 本接口用于重装一个实例。
func (c *Client) ReinstallInstance(request *ReinstallInstanceRequest) (response *ReinstallInstanceResponse, err error) {
	response = NewReinstallInstanceResponse()
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

// ModifyInstancesResourceGroup 修改实例所属的资源组。
func (c *Client) ModifyInstancesResourceGroup(request *ModifyInstancesResourceGroupRequest) (response *ModifyInstancesResourceGroupResponse, err error) {
	response = NewModifyInstancesResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminateInstanceRequest() (request *TerminateInstanceRequest) {
	request = &TerminateInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminateInstance")

	return
}

func NewTerminateInstanceResponse() (response *TerminateInstanceResponse) {
	response = &TerminateInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TerminateInstance 本接口用于退还一个实例。
func (c *Client) TerminateInstance(request *TerminateInstanceRequest) (response *TerminateInstanceResponse, err error) {
	response = NewTerminateInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
	request = &RenewInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewInstance")

	return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
	response = &RenewInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewInstance 本接口用于续费一个实例。
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
	response = NewRenewInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseInstancesRequest() (request *ReleaseInstancesRequest) {
	request = &ReleaseInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseInstances")

	return
}

func NewReleaseInstancesResponse() (response *ReleaseInstancesResponse) {
	response = &ReleaseInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReleaseInstances 本接口用于释放一个或多个实例。
func (c *Client) ReleaseInstances(request *ReleaseInstancesRequest) (response *ReleaseInstancesResponse, err error) {
	response = NewReleaseInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceInstanceBandwidthRequest() (request *InquiryPriceInstanceBandwidthRequest) {
	request = &InquiryPriceInstanceBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceInstanceBandwidth")

	return
}

func NewInquiryPriceInstanceBandwidthResponse() (response *InquiryPriceInstanceBandwidthResponse) {
	response = &InquiryPriceInstanceBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceInstanceBandwidth 实例修改带宽询价。
func (c *Client) InquiryPriceInstanceBandwidth(request *InquiryPriceInstanceBandwidthRequest) (response *InquiryPriceInstanceBandwidthResponse, err error) {
	response = NewInquiryPriceInstanceBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstanceBandwidthRequest() (request *ModifyInstanceBandwidthRequest) {
	request = &ModifyInstanceBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstanceBandwidth")

	return
}

func NewModifyInstanceBandwidthResponse() (response *ModifyInstanceBandwidthResponse) {
	response = &ModifyInstanceBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyInstanceBandwidth 修改实例带宽。
func (c *Client) ModifyInstanceBandwidth(request *ModifyInstanceBandwidthRequest) (response *ModifyInstanceBandwidthResponse, err error) {
	response = NewModifyInstanceBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCancelInstanceBandwidthDowngradeRequest() (request *CancelInstanceBandwidthDowngradeRequest) {
	request = &CancelInstanceBandwidthDowngradeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CancelInstanceBandwidthDowngrade")

	return
}

func NewCancelInstanceBandwidthDowngradeResponse() (response *CancelInstanceBandwidthDowngradeResponse) {
	response = &CancelInstanceBandwidthDowngradeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CancelInstanceBandwidthDowngrade 取消带宽降配订单。
func (c *Client) CancelInstanceBandwidthDowngrade(request *CancelInstanceBandwidthDowngradeRequest) (response *CancelInstanceBandwidthDowngradeResponse, err error) {
	response = NewCancelInstanceBandwidthDowngradeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceInstanceTrafficPackageRequest() (request *InquiryPriceInstanceTrafficPackageRequest) {
	request = &InquiryPriceInstanceTrafficPackageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceInstanceTrafficPackage")

	return
}

func NewInquiryPriceInstanceTrafficPackageResponse() (response *InquiryPriceInstanceTrafficPackageResponse) {
	response = &InquiryPriceInstanceTrafficPackageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceInstanceTrafficPackage 实例修改流量包询价。
func (c *Client) InquiryPriceInstanceTrafficPackage(request *InquiryPriceInstanceTrafficPackageRequest) (response *InquiryPriceInstanceTrafficPackageResponse, err error) {
	response = NewInquiryPriceInstanceTrafficPackageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstanceTrafficPackageRequest() (request *ModifyInstanceTrafficPackageRequest) {
	request = &ModifyInstanceTrafficPackageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstanceTrafficPackage")

	return
}

func NewModifyInstanceTrafficPackageResponse() (response *ModifyInstanceTrafficPackageResponse) {
	response = &ModifyInstanceTrafficPackageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyInstanceTrafficPackage 修改实例流量包。
func (c *Client) ModifyInstanceTrafficPackage(request *ModifyInstanceTrafficPackageRequest) (response *ModifyInstanceTrafficPackageResponse, err error) {
	response = NewModifyInstanceTrafficPackageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCancelInstanceTrafficPackageDowngradeRequest() (request *CancelInstanceTrafficPackageDowngradeRequest) {
	request = &CancelInstanceTrafficPackageDowngradeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CancelInstanceTrafficPackageDowngrade")

	return
}

func NewCancelInstanceTrafficPackageDowngradeResponse() (response *CancelInstanceTrafficPackageDowngradeResponse) {
	response = &CancelInstanceTrafficPackageDowngradeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CancelInstanceTrafficPackageDowngrade 取消流量包降配订单。
func (c *Client) CancelInstanceTrafficPackageDowngrade(request *CancelInstanceTrafficPackageDowngradeRequest) (response *CancelInstanceTrafficPackageDowngradeResponse, err error) {
	response = NewCancelInstanceTrafficPackageDowngradeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceInternetStatusRequest() (request *DescribeInstanceInternetStatusRequest) {
	request = &DescribeInstanceInternetStatusRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceInternetStatus")

	return
}

func NewDescribeInstanceInternetStatusResponse() (response *DescribeInstanceInternetStatusResponse) {
	response = &DescribeInstanceInternetStatusResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceInternetStatus 查询实例带宽、流量包状态。
func (c *Client) DescribeInstanceInternetStatus(request *DescribeInstanceInternetStatusRequest) (response *DescribeInstanceInternetStatusResponse, err error) {
	response = NewDescribeInstanceInternetStatusResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceTrafficRequest() (request *DescribeInstanceTrafficRequest) {
	request = &DescribeInstanceTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceTraffic")

	return
}

func NewDescribeInstanceTrafficResponse() (response *DescribeInstanceTrafficResponse) {
	response = &DescribeInstanceTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceTraffic 查询实例指定时间段内的流量信息。
func (c *Client) DescribeInstanceTraffic(request *DescribeInstanceTrafficRequest) (response *DescribeInstanceTrafficResponse, err error) {
	response = NewDescribeInstanceTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstancesMonitorHealthRequest() (request *DescribeInstancesMonitorHealthRequest) {
	request = &DescribeInstancesMonitorHealthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstancesMonitorHealth")

	return
}

func NewDescribeInstancesMonitorHealthResponse() (response *DescribeInstancesMonitorHealthResponse) {
	response = &DescribeInstancesMonitorHealthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstancesMonitorHealth 查询实例的硬件状态信息。
func (c *Client) DescribeInstancesMonitorHealth(request *DescribeInstancesMonitorHealthRequest) (response *DescribeInstancesMonitorHealthResponse, err error) {
	response = NewDescribeInstancesMonitorHealthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAvailableIpv4ResourcesRequest() (request *DescribeAvailableIpv4ResourcesRequest) {
	request = &DescribeAvailableIpv4ResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAvailableIpv4Resources")

	return
}

func NewDescribeAvailableIpv4ResourcesResponse() (response *DescribeAvailableIpv4ResourcesResponse) {
	response = &DescribeAvailableIpv4ResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAvailableIpv4Resources 查询可售的Ipv4 Cidr Block资源。
func (c *Client) DescribeAvailableIpv4Resources(request *DescribeAvailableIpv4ResourcesRequest) (response *DescribeAvailableIpv4ResourcesResponse, err error) {
	response = NewDescribeAvailableIpv4ResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateIpv4BlockRequest() (request *InquiryPriceCreateIpv4BlockRequest) {
	request = &InquiryPriceCreateIpv4BlockRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateIpv4Block")

	return
}

func NewInquiryPriceCreateIpv4BlockResponse() (response *InquiryPriceCreateIpv4BlockResponse) {
	response = &InquiryPriceCreateIpv4BlockResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateIpv4Block 创建Ipv4 CidrBlock询价。
func (c *Client) InquiryPriceCreateIpv4Block(request *InquiryPriceCreateIpv4BlockRequest) (response *InquiryPriceCreateIpv4BlockResponse, err error) {
	response = NewInquiryPriceCreateIpv4BlockResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateIpv4BlockRequest() (request *CreateIpv4BlockRequest) {
	request = &CreateIpv4BlockRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateIpv4Block")

	return
}

func NewCreateIpv4BlockResponse() (response *CreateIpv4BlockResponse) {
	response = &CreateIpv4BlockResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateIpv4Block 创建一个或多个Ipv4 Cidr Block。
func (c *Client) CreateIpv4Block(request *CreateIpv4BlockRequest) (response *CreateIpv4BlockResponse, err error) {
	response = NewCreateIpv4BlockResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAvailableIpv6ResourcesRequest() (request *DescribeAvailableIpv6ResourcesRequest) {
	request = &DescribeAvailableIpv6ResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAvailableIpv6Resources")

	return
}

func NewDescribeAvailableIpv6ResourcesResponse() (response *DescribeAvailableIpv6ResourcesResponse) {
	response = &DescribeAvailableIpv6ResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAvailableIpv6Resources 查询可售的Ipv6 Cidr Block资源。
func (c *Client) DescribeAvailableIpv6Resources(request *DescribeAvailableIpv6ResourcesRequest) (response *DescribeAvailableIpv6ResourcesResponse, err error) {
	response = NewDescribeAvailableIpv6ResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateIpv6BlockRequest() (request *CreateIpv6BlockRequest) {
	request = &CreateIpv6BlockRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateIpv6Block")

	return
}

func NewCreateIpv6BlockResponse() (response *CreateIpv6BlockResponse) {
	response = &CreateIpv6BlockResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateIpv6Block 创建一个或多个IPv6 Cidr Block。
func (c *Client) CreateIpv6Block(request *CreateIpv6BlockRequest) (response *CreateIpv6BlockResponse, err error) {
	response = NewCreateIpv6BlockResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrBlocksRequest() (request *DescribeCidrBlocksRequest) {
	request = &DescribeCidrBlocksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrBlocks")

	return
}

func NewDescribeCidrBlocksResponse() (response *DescribeCidrBlocksResponse) {
	response = &DescribeCidrBlocksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCidrBlocks 查询一个或多个Cidr Block的信息。
func (c *Client) DescribeCidrBlocks(request *DescribeCidrBlocksRequest) (response *DescribeCidrBlocksResponse, err error) {
	response = NewDescribeCidrBlocksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCidrBlocksAttributeRequest() (request *ModifyCidrBlocksAttributeRequest) {
	request = &ModifyCidrBlocksAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCidrBlocksAttribute")

	return
}

func NewModifyCidrBlocksAttributeResponse() (response *ModifyCidrBlocksAttributeResponse) {
	response = &ModifyCidrBlocksAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyCidrBlocksAttribute 本接口用于修改一个或多个Cidr Block的属性（目前只支持修改名称）。
func (c *Client) ModifyCidrBlocksAttribute(request *ModifyCidrBlocksAttributeRequest) (response *ModifyCidrBlocksAttributeResponse, err error) {
	response = NewModifyCidrBlocksAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewCidrBlockRequest() (request *RenewCidrBlockRequest) {
	request = &RenewCidrBlockRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewCidrBlock")

	return
}

func NewRenewCidrBlockResponse() (response *RenewCidrBlockResponse) {
	response = &RenewCidrBlockResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewCidrBlock 本接口用于续费一个Cidr Block。
func (c *Client) RenewCidrBlock(request *RenewCidrBlockRequest) (response *RenewCidrBlockResponse, err error) {
	response = NewRenewCidrBlockResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminateCidrBlockRequest() (request *TerminateCidrBlockRequest) {
	request = &TerminateCidrBlockRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminateCidrBlock")

	return
}

func NewTerminateCidrBlockResponse() (response *TerminateCidrBlockResponse) {
	response = &TerminateCidrBlockResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TerminateCidrBlock 本接口用于退还一个Cidr Block。
func (c *Client) TerminateCidrBlock(request *TerminateCidrBlockRequest) (response *TerminateCidrBlockResponse, err error) {
	response = NewTerminateCidrBlockResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseCidrBlocksRequest() (request *ReleaseCidrBlocksRequest) {
	request = &ReleaseCidrBlocksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseCidrBlocks")

	return
}

func NewReleaseCidrBlocksResponse() (response *ReleaseCidrBlocksResponse) {
	response = &ReleaseCidrBlocksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReleaseCidrBlocks 释放一个或多个Ipv4 Cidr Block。
func (c *Client) ReleaseCidrBlocks(request *ReleaseCidrBlocksRequest) (response *ReleaseCidrBlocksResponse, err error) {
	response = NewReleaseCidrBlocksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceAvailableCidrBlockRequest() (request *DescribeInstanceAvailableCidrBlockRequest) {
	request = &DescribeInstanceAvailableCidrBlockRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceAvailableCidrBlock")

	return
}

func NewDescribeInstanceAvailableCidrBlockResponse() (response *DescribeInstanceAvailableCidrBlockResponse) {
	response = &DescribeInstanceAvailableCidrBlockResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceAvailableCidrBlock 查询实例可用的Cidr Block。
func (c *Client) DescribeInstanceAvailableCidrBlock(request *DescribeInstanceAvailableCidrBlockRequest) (response *DescribeInstanceAvailableCidrBlockResponse, err error) {
	response = NewDescribeInstanceAvailableCidrBlockResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrBlockIpsRequest() (request *DescribeCidrBlockIpsRequest) {
	request = &DescribeCidrBlockIpsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrBlockIps")

	return
}

func NewDescribeCidrBlockIpsResponse() (response *DescribeCidrBlockIpsResponse) {
	response = &DescribeCidrBlockIpsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCidrBlockIps 查询一个Cidr Block的IP列表。
func (c *Client) DescribeCidrBlockIps(request *DescribeCidrBlockIpsRequest) (response *DescribeCidrBlockIpsResponse, err error) {
	response = NewDescribeCidrBlockIpsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewBindCidrBlockIpsRequest() (request *BindCidrBlockIpsRequest) {
	request = &BindCidrBlockIpsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "BindCidrBlockIps")

	return
}

func NewBindCidrBlockIpsResponse() (response *BindCidrBlockIpsResponse) {
	response = &BindCidrBlockIpsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// BindCidrBlockIps 实例绑定一个或多个Cidr Block IP。
func (c *Client) BindCidrBlockIps(request *BindCidrBlockIpsRequest) (response *BindCidrBlockIpsResponse, err error) {
	response = NewBindCidrBlockIpsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnbindCidrBlockIpsRequest() (request *UnbindCidrBlockIpsRequest) {
	request = &UnbindCidrBlockIpsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnbindCidrBlockIps")

	return
}

func NewUnbindCidrBlockIpsResponse() (response *UnbindCidrBlockIpsResponse) {
	response = &UnbindCidrBlockIpsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnbindCidrBlockIps 实例解绑一个或多个Cidr Block IP。
func (c *Client) UnbindCidrBlockIps(request *UnbindCidrBlockIpsRequest) (response *UnbindCidrBlockIpsResponse, err error) {
	response = NewUnbindCidrBlockIpsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateByoipRequest() (request *CreateByoipRequest) {
	request = &CreateByoipRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateByoip")

	return
}

func NewCreateByoipResponse() (response *CreateByoipResponse) {
	response = &CreateByoipResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateByoip 创建一个BYO IP。
func (c *Client) CreateByoip(request *CreateByoipRequest) (response *CreateByoipResponse, err error) {
	response = NewCreateByoipResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateEipAddressRequest() (request *InquiryPriceCreateEipAddressRequest) {
	request = &InquiryPriceCreateEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateEipAddress")

	return
}

func NewInquiryPriceCreateEipAddressResponse() (response *InquiryPriceCreateEipAddressResponse) {
	response = &InquiryPriceCreateEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateEipAddress 创建EIP询价。
func (c *Client) InquiryPriceCreateEipAddress(request *InquiryPriceCreateEipAddressRequest) (response *InquiryPriceCreateEipAddressResponse, err error) {
	response = NewInquiryPriceCreateEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipAddressesRequest() (request *DescribeEipAddressesRequest) {
	request = &DescribeEipAddressesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEipAddresses")

	return
}

func NewDescribeEipAddressesResponse() (response *DescribeEipAddressesResponse) {
	response = &DescribeEipAddressesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEipAddresses 查询一台或多台指定EIP的信息。用户可以根据EIP ID、IP或者计费模式等信息来搜索EIP的信息。
func (c *Client) DescribeEipAddresses(request *DescribeEipAddressesRequest) (response *DescribeEipAddressesResponse, err error) {
	response = NewDescribeEipAddressesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAllocateEipAddressesRequest() (request *AllocateEipAddressesRequest) {
	request = &AllocateEipAddressesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AllocateEipAddresses")

	return
}

func NewAllocateEipAddressesResponse() (response *AllocateEipAddressesResponse) {
	response = &AllocateEipAddressesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AllocateEipAddresses 创建一个或多个EIP。
func (c *Client) AllocateEipAddresses(request *AllocateEipAddressesRequest) (response *AllocateEipAddressesResponse, err error) {
	response = NewAllocateEipAddressesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyEipAddressesResourceGroupRequest() (request *ModifyEipAddressesResourceGroupRequest) {
	request = &ModifyEipAddressesResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyEipAddressesResourceGroup")

	return
}

func NewModifyEipAddressesResourceGroupResponse() (response *ModifyEipAddressesResourceGroupResponse) {
	response = &ModifyEipAddressesResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyEipAddressesResourceGroup 修改弹性IP所属的资源组。
func (c *Client) ModifyEipAddressesResourceGroup(request *ModifyEipAddressesResourceGroupRequest) (response *ModifyEipAddressesResourceGroupResponse, err error) {
	response = NewModifyEipAddressesResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceAvailableEipResourcesRequest() (request *DescribeInstanceAvailableEipResourcesRequest) {
	request = &DescribeInstanceAvailableEipResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceAvailableEipResources")

	return
}

func NewDescribeInstanceAvailableEipResourcesResponse() (response *DescribeInstanceAvailableEipResourcesResponse) {
	response = &DescribeInstanceAvailableEipResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceAvailableEipResources 查询实例可绑定的EIP列表。
func (c *Client) DescribeInstanceAvailableEipResources(request *DescribeInstanceAvailableEipResourcesRequest) (response *DescribeInstanceAvailableEipResourcesResponse, err error) {
	response = NewDescribeInstanceAvailableEipResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssociateEipAddressRequest() (request *AssociateEipAddressRequest) {
	request = &AssociateEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssociateEipAddress")

	return
}

func NewAssociateEipAddressResponse() (response *AssociateEipAddressResponse) {
	response = &AssociateEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssociateEipAddress 将EIP绑定到同区域的机器实例上。
func (c *Client) AssociateEipAddress(request *AssociateEipAddressRequest) (response *AssociateEipAddressResponse, err error) {
	response = NewAssociateEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnAssociateEipAddressRequest() (request *UnAssociateEipAddressRequest) {
	request = &UnAssociateEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnAssociateEipAddress")

	return
}

func NewUnAssociateEipAddressResponse() (response *UnAssociateEipAddressResponse) {
	response = &UnAssociateEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnAssociateEipAddress 将EIP上已绑定的机器实例解绑。
func (c *Client) UnAssociateEipAddress(request *UnAssociateEipAddressRequest) (response *UnAssociateEipAddressResponse, err error) {
	response = NewUnAssociateEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipAvailableResourcesRequest() (request *DescribeEipAvailableResourcesRequest) {
	request = &DescribeEipAvailableResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEipAvailableResources")

	return
}

func NewDescribeEipAvailableResourcesResponse() (response *DescribeEipAvailableResourcesResponse) {
	response = &DescribeEipAvailableResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEipAvailableResources 查询区域可购买EIP资源。
func (c *Client) DescribeEipAvailableResources(request *DescribeEipAvailableResourcesRequest) (response *DescribeEipAvailableResourcesResponse, err error) {
	response = NewDescribeEipAvailableResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminateEipAddressRequest() (request *TerminateEipAddressRequest) {
	request = &TerminateEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminateEipAddress")

	return
}

func NewTerminateEipAddressResponse() (response *TerminateEipAddressResponse) {
	response = &TerminateEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TerminateEipAddress 本接口用于退还一个EIP。
func (c *Client) TerminateEipAddress(request *TerminateEipAddressRequest) (response *TerminateEipAddressResponse, err error) {
	response = NewTerminateEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewEipAddressRequest() (request *RenewEipAddressRequest) {
	request = &RenewEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewEipAddress")

	return
}

func NewRenewEipAddressResponse() (response *RenewEipAddressResponse) {
	response = &RenewEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewEipAddress 本接口用于续费一个EIP。
func (c *Client) RenewEipAddress(request *RenewEipAddressRequest) (response *RenewEipAddressResponse, err error) {
	response = NewRenewEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseEipAddressesRequest() (request *ReleaseEipAddressesRequest) {
	request = &ReleaseEipAddressesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseEipAddresses")

	return
}

func NewReleaseEipAddressesResponse() (response *ReleaseEipAddressesResponse) {
	response = &ReleaseEipAddressesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReleaseEipAddresses 释放一个或多个EIP。
func (c *Client) ReleaseEipAddresses(request *ReleaseEipAddressesRequest) (response *ReleaseEipAddressesResponse, err error) {
	response = NewReleaseEipAddressesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeVpcAvailableRegionsRequest() (request *DescribeVpcAvailableRegionsRequest) {
	request = &DescribeVpcAvailableRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeVpcAvailableRegions")

	return
}

func NewDescribeVpcAvailableRegionsResponse() (response *DescribeVpcAvailableRegionsResponse) {
	response = &DescribeVpcAvailableRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeVpcAvailableRegions 本接口用于查询支持VPC组网的节点区域信息。
func (c *Client) DescribeVpcAvailableRegions(request *DescribeVpcAvailableRegionsRequest) (response *DescribeVpcAvailableRegionsResponse, err error) {
	response = NewDescribeVpcAvailableRegionsResponse()
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

// CreateVpc 本接口用于创建一个私有网络VPC。
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
	response = NewCreateVpcResponse()
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

// DescribeVpcs 本接口用于查询一个或多个指定VPC的信息。用户可以根据VPC ID、Subnet ID、 VPC节点ID、VPC名称等信息来搜索VPC信息。
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
	response = NewDescribeVpcsResponse()
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

// ModifyVpcsAttribute 本接口用于修改VPC的属性（目前只支持修改VPC的名称）。
func (c *Client) ModifyVpcsAttribute(request *ModifyVpcsAttributeRequest) (response *ModifyVpcsAttributeResponse, err error) {
	response = NewModifyVpcsAttributeResponse()
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

// DeleteVpc 本接口用于删除一个Vpc。
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
	response = NewDeleteVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeLoadBalancerZonesRequest() (request *DescribeLoadBalancerZonesRequest) {
	request = &DescribeLoadBalancerZonesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeLoadBalancerZones")

	return
}

func NewDescribeLoadBalancerZonesResponse() (response *DescribeLoadBalancerZonesResponse) {
	response = &DescribeLoadBalancerZonesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeLoadBalancerZones 查询支持创建负载均衡的区域。
func (c *Client) DescribeLoadBalancerZones(request *DescribeLoadBalancerZonesRequest) (response *DescribeLoadBalancerZonesResponse, err error) {
	response = NewDescribeLoadBalancerZonesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeLoadBalancerSpecsRequest() (request *DescribeLoadBalancerSpecsRequest) {
	request = &DescribeLoadBalancerSpecsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeLoadBalancerSpecs")

	return
}

func NewDescribeLoadBalancerSpecsResponse() (response *DescribeLoadBalancerSpecsResponse) {
	response = &DescribeLoadBalancerSpecsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeLoadBalancerSpecs 查询可用区节点下可用的负载均衡规格列表。
func (c *Client) DescribeLoadBalancerSpecs(request *DescribeLoadBalancerSpecsRequest) (response *DescribeLoadBalancerSpecsResponse, err error) {
	response = NewDescribeLoadBalancerSpecsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateLoadBalancer")

	return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateLoadBalancer 创建负载均衡实例。
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = NewCreateLoadBalancerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
	request = &DescribeLoadBalancersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeLoadBalancers")

	return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
	response = &DescribeLoadBalancersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeLoadBalancers 查询一台或多台指定负载均衡实例的信息。用户可以根据负载均衡实例的ID、可用区等信息来搜索负载均衡实例的信息。
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
	response = NewDescribeLoadBalancersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyLoadBalancersNameRequest() (request *ModifyLoadBalancersNameRequest) {
	request = &ModifyLoadBalancersNameRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyLoadBalancersName")

	return
}

func NewModifyLoadBalancersNameResponse() (response *ModifyLoadBalancersNameResponse) {
	response = &ModifyLoadBalancersNameResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyLoadBalancersName 修改负载均衡实例的名称。
func (c *Client) ModifyLoadBalancersName(request *ModifyLoadBalancersNameRequest) (response *ModifyLoadBalancersNameResponse, err error) {
	response = NewModifyLoadBalancersNameResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminateLoadBalancerRequest() (request *TerminateLoadBalancerRequest) {
	request = &TerminateLoadBalancerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminateLoadBalancer")

	return
}

func NewTerminateLoadBalancerResponse() (response *TerminateLoadBalancerResponse) {
	response = &TerminateLoadBalancerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TerminateLoadBalancer 终止负载均衡实例。
func (c *Client) TerminateLoadBalancer(request *TerminateLoadBalancerRequest) (response *TerminateLoadBalancerResponse, err error) {
	response = NewTerminateLoadBalancerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRestoreLoadBalancerRequest() (request *RestoreLoadBalancerRequest) {
	request = &RestoreLoadBalancerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RestoreLoadBalancer")

	return
}

func NewRestoreLoadBalancerResponse() (response *RestoreLoadBalancerResponse) {
	response = &RestoreLoadBalancerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RestoreLoadBalancer 恢复被删除的负载均衡实例。
func (c *Client) RestoreLoadBalancer(request *RestoreLoadBalancerRequest) (response *RestoreLoadBalancerResponse, err error) {
	response = NewRestoreLoadBalancerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseLoadBalancerRequest() (request *ReleaseLoadBalancerRequest) {
	request = &ReleaseLoadBalancerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseLoadBalancer")

	return
}

func NewReleaseLoadBalancerResponse() (response *ReleaseLoadBalancerResponse) {
	response = &ReleaseLoadBalancerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReleaseLoadBalancer 释放负载均衡实例。
func (c *Client) ReleaseLoadBalancer(request *ReleaseLoadBalancerRequest) (response *ReleaseLoadBalancerResponse, err error) {
	response = NewReleaseLoadBalancerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
	request = &CreateListenerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateListener")

	return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
	response = &CreateListenerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateListener 在网络型负载均衡实例中创建TCP、UDP或TCPSSL监听。
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
	response = NewCreateListenerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
	request = &DescribeListenersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeListeners")

	return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
	response = &DescribeListenersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeListeners 查询一台或多台指定Listener的信息。用户可以根据负载均衡实例的ID、监听器的ID等信息来搜索监听器的信息。
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
	response = NewDescribeListenersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyListenerAttributeRequest() (request *ModifyListenerAttributeRequest) {
	request = &ModifyListenerAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyListenerAttribute")

	return
}

func NewModifyListenerAttributeResponse() (response *ModifyListenerAttributeResponse) {
	response = &ModifyListenerAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyListenerAttribute 修改监听器配置。
func (c *Client) ModifyListenerAttribute(request *ModifyListenerAttributeRequest) (response *ModifyListenerAttributeResponse, err error) {
	response = NewModifyListenerAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
	request = &DeleteListenerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteListener")

	return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
	response = &DeleteListenerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteListener 删除负载均衡监听器。
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
	response = NewDeleteListenerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateLoadBalancerVIPsRequest() (request *CreateLoadBalancerVIPsRequest) {
	request = &CreateLoadBalancerVIPsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateLoadBalancerVIPs")

	return
}

func NewCreateLoadBalancerVIPsResponse() (response *CreateLoadBalancerVIPsResponse) {
	response = &CreateLoadBalancerVIPsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateLoadBalancerVIPs 创建负载均衡的VIP。
func (c *Client) CreateLoadBalancerVIPs(request *CreateLoadBalancerVIPsRequest) (response *CreateLoadBalancerVIPsResponse, err error) {
	response = NewCreateLoadBalancerVIPsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteLoadBalancerVIPRequest() (request *DeleteLoadBalancerVIPRequest) {
	request = &DeleteLoadBalancerVIPRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteLoadBalancerVIP")

	return
}

func NewDeleteLoadBalancerVIPResponse() (response *DeleteLoadBalancerVIPResponse) {
	response = &DeleteLoadBalancerVIPResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteLoadBalancerVIP 删除负载均衡的IP。
func (c *Client) DeleteLoadBalancerVIP(request *DeleteLoadBalancerVIPRequest) (response *DeleteLoadBalancerVIPResponse, err error) {
	response = NewDeleteLoadBalancerVIPResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBackendsRequest() (request *DescribeBackendsRequest) {
	request = &DescribeBackendsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBackends")

	return
}

func NewDescribeBackendsResponse() (response *DescribeBackendsResponse) {
	response = &DescribeBackendsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBackends 查询一台或多台指定后端配置服务器的信息。用户可以根据监听器的ID、后端配置服务器的ID等信息来搜索后端配置服务器的信息。
func (c *Client) DescribeBackends(request *DescribeBackendsRequest) (response *DescribeBackendsResponse, err error) {
	response = NewDescribeBackendsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRegisterBackendRequest() (request *RegisterBackendRequest) {
	request = &RegisterBackendRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RegisterBackend")

	return
}

func NewRegisterBackendResponse() (response *RegisterBackendResponse) {
	response = &RegisterBackendResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RegisterBackend 创建将一台后端服务绑定到负载均衡的监听器。
func (c *Client) RegisterBackend(request *RegisterBackendRequest) (response *RegisterBackendResponse, err error) {
	response = NewRegisterBackendResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeregisterBackendRequest() (request *DeregisterBackendRequest) {
	request = &DeregisterBackendRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeregisterBackend")

	return
}

func NewDeregisterBackendResponse() (response *DeregisterBackendResponse) {
	response = &DeregisterBackendResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeregisterBackend 删除负载均衡后端配置服务器。
func (c *Client) DeregisterBackend(request *DeregisterBackendRequest) (response *DeregisterBackendResponse, err error) {
	response = NewDeregisterBackendResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeManagedInstancesRequest() (request *DescribeManagedInstancesRequest) {
	request = &DescribeManagedInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeManagedInstances")

	return
}

func NewDescribeManagedInstancesResponse() (response *DescribeManagedInstancesResponse) {
	response = &DescribeManagedInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeManagedInstances 本接口用于查询一台或多台实例的信息。用户可以根据实例ID、实例名称等条件来查询实例的详细信息。
func (c *Client) DescribeManagedInstances(request *DescribeManagedInstancesRequest) (response *DescribeManagedInstancesResponse, err error) {
	response = NewDescribeManagedInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeManagedInstanceTrafficRequest() (request *DescribeManagedInstanceTrafficRequest) {
	request = &DescribeManagedInstanceTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeManagedInstanceTraffic")

	return
}

func NewDescribeManagedInstanceTrafficResponse() (response *DescribeManagedInstanceTrafficResponse) {
	response = &DescribeManagedInstanceTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeManagedInstanceTraffic 查询托管实例指定时间段内的流量信息。
func (c *Client) DescribeManagedInstanceTraffic(request *DescribeManagedInstanceTrafficRequest) (response *DescribeManagedInstanceTrafficResponse, err error) {
	response = NewDescribeManagedInstanceTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

