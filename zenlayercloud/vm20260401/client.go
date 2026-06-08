/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package vm

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2026-04-01"
	SERVICE    = "vm"
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

// DescribeZones 查询可用地区。
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	response = NewDescribeZonesResponse()
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

// InquiryPriceCreateInstance 创建一台虚拟机实例询价。
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
	response = NewInquiryPriceCreateInstanceResponse()
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

// DescribeZoneInstanceConfigInfos 查询售卖可用区的机型信息。
func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
	response = NewDescribeZoneInstanceConfigInfosResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeVmInventoryCapacityRequest() (request *DescribeVmInventoryCapacityRequest) {
	request = &DescribeVmInventoryCapacityRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeVmInventoryCapacity")

	return
}

func NewDescribeVmInventoryCapacityResponse() (response *DescribeVmInventoryCapacityResponse) {
	response = &DescribeVmInventoryCapacityResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeVmInventoryCapacity 获取可用区库存。
func (c *Client) DescribeVmInventoryCapacity(request *DescribeVmInventoryCapacityRequest) (response *DescribeVmInventoryCapacityResponse, err error) {
	response = NewDescribeVmInventoryCapacityResponse()
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

// CreateInstances 创建一个或多个指定配置的虚拟机实例。
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
	response = NewCreateInstancesResponse()
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

// ResetInstancesPassword 将虚拟机实例操作系统的密码重置为用户指定的密码。
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
	response = NewResetInstancesPasswordResponse()
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

// ResetInstance 重装一台虚拟机实例上的操作系统。
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
	response = NewResetInstanceResponse()
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

// ModifyInstancesResourceGroup 修改虚拟机实例所属的资源组。
func (c *Client) ModifyInstancesResourceGroup(request *ModifyInstancesResourceGroupRequest) (response *ModifyInstancesResourceGroupResponse, err error) {
	response = NewModifyInstancesResourceGroupResponse()
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

// DescribeInstances 查询一台或多台虚拟机实例的信息。用户可以根据实例ID、实例名称或者实例计费模式等条件来查询实例的详细信息。
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

// DescribeInstancesStatus 查询一台或多台虚拟机实例的状态。
func (c *Client) DescribeInstancesStatus(request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
	response = NewDescribeInstancesStatusResponse()
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

// StartInstances 启动一个或多个虚拟机实例。
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	response = NewStartInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeVncUrlRequest() (request *DescribeVncUrlRequest) {
	request = &DescribeVncUrlRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeVncUrl")

	return
}

func NewDescribeVncUrlResponse() (response *DescribeVncUrlResponse) {
	response = &DescribeVncUrlResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeVncUrl 获取实例VNC地址。
func (c *Client) DescribeVncUrl(request *DescribeVncUrlRequest) (response *DescribeVncUrlResponse, err error) {
	response = NewDescribeVncUrlResponse()
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

// ModifyInstancesAttribute 修改虚拟机实例的属性（目前只支持修改实例的显示名称）。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
	response = NewModifyInstancesAttributeResponse()
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

// StopInstances 关闭一个或多个虚拟机实例。
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

// RebootInstances 重启一个或多个虚拟机实例。
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	response = NewRebootInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstanceTypeRequest() (request *ModifyInstanceTypeRequest) {
	request = &ModifyInstanceTypeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstanceType")

	return
}

func NewModifyInstanceTypeResponse() (response *ModifyInstanceTypeResponse) {
	response = &ModifyInstanceTypeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyInstanceType 修改一台虚拟机实例的机型。
func (c *Client) ModifyInstanceType(request *ModifyInstanceTypeRequest) (response *ModifyInstanceTypeResponse, err error) {
	response = NewModifyInstanceTypeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceTypeStatusRequest() (request *DescribeInstanceTypeStatusRequest) {
	request = &DescribeInstanceTypeStatusRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceTypeStatus")

	return
}

func NewDescribeInstanceTypeStatusResponse() (response *DescribeInstanceTypeStatusResponse) {
	response = &DescribeInstanceTypeStatusResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceTypeStatus 查询实例变配后的机型状态。
func (c *Client) DescribeInstanceTypeStatus(request *DescribeInstanceTypeStatusRequest) (response *DescribeInstanceTypeStatusResponse, err error) {
	response = NewDescribeInstanceTypeStatusResponse()
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

// InquiryPriceInstanceTrafficPackage 虚拟机实例修改流量包询价。
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

// ModifyInstanceTrafficPackage 修改虚拟机实例流量包大小。
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

// CancelInstanceTrafficPackageDowngrade 取消虚拟机实例流量包降配订单。
func (c *Client) CancelInstanceTrafficPackageDowngrade(request *CancelInstanceTrafficPackageDowngradeRequest) (response *CancelInstanceTrafficPackageDowngradeResponse, err error) {
	response = NewCancelInstanceTrafficPackageDowngradeResponse()
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

// ModifyInstanceBandwidth 修改虚拟机实例的公网出口带宽。
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

// CancelInstanceBandwidthDowngrade 取消虚拟机实例带宽降配订单。
func (c *Client) CancelInstanceBandwidthDowngrade(request *CancelInstanceBandwidthDowngradeRequest) (response *CancelInstanceBandwidthDowngradeResponse, err error) {
	response = NewCancelInstanceBandwidthDowngradeResponse()
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

func NewCancelInstanceDowngradeRequest() (request *CancelInstanceDowngradeRequest) {
	request = &CancelInstanceDowngradeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CancelInstanceDowngrade")

	return
}

func NewCancelInstanceDowngradeResponse() (response *CancelInstanceDowngradeResponse) {
	response = &CancelInstanceDowngradeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CancelInstanceDowngrade 取消虚拟机实例降配订单。
func (c *Client) CancelInstanceDowngrade(request *CancelInstanceDowngradeRequest) (response *CancelInstanceDowngradeResponse, err error) {
	response = NewCancelInstanceDowngradeResponse()
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

// TerminateInstance 退还一个虚拟机实例。
func (c *Client) TerminateInstance(request *TerminateInstanceRequest) (response *TerminateInstanceResponse, err error) {
	response = NewTerminateInstanceResponse()
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

// ReleaseInstances 释放一个或多个虚拟机实例。
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

// InquiryPriceInstanceBandwidth 虚拟机实例修改带宽询价。
func (c *Client) InquiryPriceInstanceBandwidth(request *InquiryPriceInstanceBandwidthRequest) (response *InquiryPriceInstanceBandwidthResponse, err error) {
	response = NewInquiryPriceInstanceBandwidthResponse()
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

func NewDescribeInstanceCpuMonitorRequest() (request *DescribeInstanceCpuMonitorRequest) {
	request = &DescribeInstanceCpuMonitorRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceCpuMonitor")

	return
}

func NewDescribeInstanceCpuMonitorResponse() (response *DescribeInstanceCpuMonitorResponse) {
	response = &DescribeInstanceCpuMonitorResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceCpuMonitor 查询实例指定时间段内的CPU使用率。
func (c *Client) DescribeInstanceCpuMonitor(request *DescribeInstanceCpuMonitorRequest) (response *DescribeInstanceCpuMonitorResponse, err error) {
	response = NewDescribeInstanceCpuMonitorResponse()
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

// DescribeDiskCategory 查询云硬盘支持的类型。
func (c *Client) DescribeDiskCategory(request *DescribeDiskCategoryRequest) (response *DescribeDiskCategoryResponse, err error) {
	response = NewDescribeDiskCategoryResponse()
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

// InquiryPriceCreateDisks 创建云硬盘询价。
func (c *Client) InquiryPriceCreateDisks(request *InquiryPriceCreateDisksRequest) (response *InquiryPriceCreateDisksResponse, err error) {
	response = NewInquiryPriceCreateDisksResponse()
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

// DescribeDisks 查询云硬盘列表。
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	response = NewDescribeDisksResponse()
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

// CreateDisks 创建一个或多个云硬盘。
func (c *Client) CreateDisks(request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
	response = NewCreateDisksResponse()
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

// ModifyDisksAttributes 修改一个或多个云硬盘属性。
func (c *Client) ModifyDisksAttributes(request *ModifyDisksAttributesRequest) (response *ModifyDisksAttributesResponse, err error) {
	response = NewModifyDisksAttributesResponse()
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

// AttachDisks 挂载云硬盘到云主机实例。
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
	response = NewAttachDisksResponse()
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

// DetachDisks 从云主机实例上卸载云硬盘。
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
	response = NewDetachDisksResponse()
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

// ChangeDisksAttach 将一个或多个已经挂载到一台实例的云硬盘挂载到另外一台实例上。
func (c *Client) ChangeDisksAttach(request *ChangeDisksAttachRequest) (response *ChangeDisksAttachResponse, err error) {
	response = NewChangeDisksAttachResponse()
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

// TerminateDisk 终止云硬盘。
func (c *Client) TerminateDisk(request *TerminateDiskRequest) (response *TerminateDiskResponse, err error) {
	response = NewTerminateDiskResponse()
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

// ModifyDisksResourceGroup 修改云硬盘所属的资源组。
func (c *Client) ModifyDisksResourceGroup(request *ModifyDisksResourceGroupRequest) (response *ModifyDisksResourceGroupResponse, err error) {
	response = NewModifyDisksResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
	request = &ResizeDiskRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ResizeDisk")

	return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
	response = &ResizeDiskResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ResizeDisk 将一个云硬盘扩容到指定大小。
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
	response = NewResizeDiskResponse()
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

// ReleaseDisk 释放云硬盘。
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

// RenewDisk 续费云硬盘。
func (c *Client) RenewDisk(request *RenewDiskRequest) (response *RenewDiskResponse, err error) {
	response = NewRenewDiskResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeImageQuotaRequest() (request *DescribeImageQuotaRequest) {
	request = &DescribeImageQuotaRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeImageQuota")

	return
}

func NewDescribeImageQuotaResponse() (response *DescribeImageQuotaResponse) {
	response = &DescribeImageQuotaResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeImageQuota 查询可创建镜像的配额。
func (c *Client) DescribeImageQuota(request *DescribeImageQuotaRequest) (response *DescribeImageQuotaResponse, err error) {
	response = NewDescribeImageQuotaResponse()
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

func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateImage")

	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateImage 创建自定义镜像。
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
	response = NewCreateImageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyImagesAttributesRequest() (request *ModifyImagesAttributesRequest) {
	request = &ModifyImagesAttributesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyImagesAttributes")

	return
}

func NewModifyImagesAttributesResponse() (response *ModifyImagesAttributesResponse) {
	response = &ModifyImagesAttributesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyImagesAttributes 修改镜像属性。
func (c *Client) ModifyImagesAttributes(request *ModifyImagesAttributesRequest) (response *ModifyImagesAttributesResponse, err error) {
	response = NewModifyImagesAttributesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
	request = &DeleteImagesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteImages")

	return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
	response = &DeleteImagesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteImages 删除一个或多个镜像。
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
	response = NewDeleteImagesResponse()
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

// DescribeSubnets 查询一台或多台指定子网的信息。用户可以根据Subnet ID、区域、Subnet名称等信息来搜索Subnet信息。
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
	response = NewDescribeSubnetsResponse()
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

// CreateSubnet 创建子网。
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
	response = NewCreateSubnetResponse()
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

// ModifySubnetsAttribute 修改子网的属性（目前只支持修改子网的名称）。
func (c *Client) ModifySubnetsAttribute(request *ModifySubnetsAttributeRequest) (response *ModifySubnetsAttributeResponse, err error) {
	response = NewModifySubnetsAttributeResponse()
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

// DeleteSubnet 删除一个子网。
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
	response = NewDeleteSubnetResponse()
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

// DescribeSecurityGroups 查询一个或多个指定安全组的信息。用户可以根据安全组ID、安全组名称等信息来搜索安全组信息。
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
	response = NewDescribeSecurityGroupsResponse()
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

// CreateSecurityGroup 创建安全组。
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
	response = NewCreateSecurityGroupResponse()
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

// ModifySecurityGroupsAttribute 修改安全组的属性（目前只支持修改安全组的名称）。
func (c *Client) ModifySecurityGroupsAttribute(request *ModifySecurityGroupsAttributeRequest) (response *ModifySecurityGroupsAttributeResponse, err error) {
	response = NewModifySecurityGroupsAttributeResponse()
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

// DeleteSecurityGroup 删除安全组。
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

// ConfigureSecurityGroupRules 配置安全组的规则。
func (c *Client) ConfigureSecurityGroupRules(request *ConfigureSecurityGroupRulesRequest) (response *ConfigureSecurityGroupRulesResponse, err error) {
	response = NewConfigureSecurityGroupRulesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRevokeSecurityGroupRulesRequest() (request *RevokeSecurityGroupRulesRequest) {
	request = &RevokeSecurityGroupRulesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RevokeSecurityGroupRules")

	return
}

func NewRevokeSecurityGroupRulesResponse() (response *RevokeSecurityGroupRulesResponse) {
	response = &RevokeSecurityGroupRulesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RevokeSecurityGroupRules 移除安全组的规则。
func (c *Client) RevokeSecurityGroupRules(request *RevokeSecurityGroupRulesRequest) (response *RevokeSecurityGroupRulesResponse, err error) {
	response = NewRevokeSecurityGroupRulesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssociateSecurityGroupInstanceRequest() (request *AssociateSecurityGroupInstanceRequest) {
	request = &AssociateSecurityGroupInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssociateSecurityGroupInstance")

	return
}

func NewAssociateSecurityGroupInstanceResponse() (response *AssociateSecurityGroupInstanceResponse) {
	response = &AssociateSecurityGroupInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssociateSecurityGroupInstance 为安全组绑定实例。
func (c *Client) AssociateSecurityGroupInstance(request *AssociateSecurityGroupInstanceRequest) (response *AssociateSecurityGroupInstanceResponse, err error) {
	response = NewAssociateSecurityGroupInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnAssociateSecurityGroupInstanceRequest() (request *UnAssociateSecurityGroupInstanceRequest) {
	request = &UnAssociateSecurityGroupInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnAssociateSecurityGroupInstance")

	return
}

func NewUnAssociateSecurityGroupInstanceResponse() (response *UnAssociateSecurityGroupInstanceResponse) {
	response = &UnAssociateSecurityGroupInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnAssociateSecurityGroupInstance 为安全组解绑实例。
func (c *Client) UnAssociateSecurityGroupInstance(request *UnAssociateSecurityGroupInstanceRequest) (response *UnAssociateSecurityGroupInstanceResponse, err error) {
	response = NewUnAssociateSecurityGroupInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceAvailableSecurityGroupResourcesRequest() (request *DescribeInstanceAvailableSecurityGroupResourcesRequest) {
	request = &DescribeInstanceAvailableSecurityGroupResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceAvailableSecurityGroupResources")

	return
}

func NewDescribeInstanceAvailableSecurityGroupResourcesResponse() (response *DescribeInstanceAvailableSecurityGroupResourcesResponse) {
	response = &DescribeInstanceAvailableSecurityGroupResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceAvailableSecurityGroupResources 获取实例可绑定的安全组。
func (c *Client) DescribeInstanceAvailableSecurityGroupResources(request *DescribeInstanceAvailableSecurityGroupResourcesRequest) (response *DescribeInstanceAvailableSecurityGroupResourcesResponse, err error) {
	response = NewDescribeInstanceAvailableSecurityGroupResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAuthorizeSecurityGroupRulesRequest() (request *AuthorizeSecurityGroupRulesRequest) {
	request = &AuthorizeSecurityGroupRulesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AuthorizeSecurityGroupRules")

	return
}

func NewAuthorizeSecurityGroupRulesResponse() (response *AuthorizeSecurityGroupRulesResponse) {
	response = &AuthorizeSecurityGroupRulesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AuthorizeSecurityGroupRules 新增安全组的规则。
func (c *Client) AuthorizeSecurityGroupRules(request *AuthorizeSecurityGroupRulesRequest) (response *AuthorizeSecurityGroupRulesResponse, err error) {
	response = NewAuthorizeSecurityGroupRulesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAuthorizeSecurityGroupRuleRequest() (request *AuthorizeSecurityGroupRuleRequest) {
	request = &AuthorizeSecurityGroupRuleRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AuthorizeSecurityGroupRule")

	return
}

func NewAuthorizeSecurityGroupRuleResponse() (response *AuthorizeSecurityGroupRuleResponse) {
	response = &AuthorizeSecurityGroupRuleResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AuthorizeSecurityGroupRule 新增安全组单条规则。
func (c *Client) AuthorizeSecurityGroupRule(request *AuthorizeSecurityGroupRuleRequest) (response *AuthorizeSecurityGroupRuleResponse, err error) {
	response = NewAuthorizeSecurityGroupRuleResponse()
	err = c.ApiCall(request, response)
	return
}

