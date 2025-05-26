package bmc

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2022-11-20"
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

// DescribeZones
// This API is used to query zones available.
//
// Possible error codes to return:
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	response = NewDescribeZonesResponse()
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

// CreateInstances
// This API is used to create one or more instances with a specified configuration.
//
//	 Possible error codes to return:
//		INVALID_PARAMETER_HOSTNAME_EXCEED = "Invalid.Parameter.Hostname.Exceed"
//		INVALID_PARAMETER_HOSTNAME_MALFORMED = "Invalid.Parameter.Hostname.Malformed"
//		INVALID_PARAMETER_INSTANCE_NAME_EXCEED = "Invalid.Parameter.Instance.Name.Exceed"
//		OPERATION_FILED_INTERNET_CHARGE_TYPE_NOT_SUPPORT = "Operation.Filed.Internet.Charge.Type.Not.Support"
//		INVALID_IMAGE_NOT_FOUND = "Invalid.Image.Not.Found"
//		INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
//		INVALID_INSTANCE_TYPE_NOT_FOUND = "Invalid.Instance.Type.Not.Found"
//		INVALID_PARTITION_IMAGE_NOT_SET = "Invalid.Partition.Image.Not.Set"
//		OPERATION_DENIED_INSTANCE_QUOTA_EXCEED = "Operation.Denied.Instance.Quota.Exceed"
//		INVALID_BANDWIDTH_VALUE_EXCEED_MAXIMUM = "Invalid.Bandwidth.Value.Exceed.MaxImum"
//		INVALID_PARAMETER_VALUE_PASSWORD_MALFORMED = "Invalid.Parameter.Value.Password.Malformed"
//		INVALID_PARAMETER_INSTANCE_LOGIN_CONFLICT = "Invalid.Parameter.Instance.Login.Conflict"
//		INVALID_PARAMETER_SSH_KEY_MALFORMED = "Invalid.Parameter.Ssh.Key.Malformed"
//		INVALID_RAID_CONFIG_FAST_CUSTOM_CONFLICT = "Invalid.Raid.Config.Fast.Custom.Conflict"
//		INVALID_INSTANCE_TYPE_RAID_NOT_SUPPORT = "Invalid.Instance.Type.Raid.Not.Support"
//		INVALID_PARAMETER_NIC_NAME_CONFLICT = "Invalid.Parameter.Nic.Name.Conflict"
//		INVALID_PARAMETER_NIC_NAME_MALFORMED = "Invalid.Parameter.Nic.Name.Malformed"
//		INVALID_PARTITION_SIZE_NOT_FULL = "Invalid.Partition.Size.Not.Full"
//		INVALID_PARTITION_DUPLICATE_FILE_PATH = "Invalid.Partition.Duplicate.File.Path"
//		INVALID_PARTITION_MISSING_REQUIRED_FILE_PATH = "Invalid.Partition.Missing.Required.File.Path"
//		INVALID_PARTITION_MALFORMED = "Invalid.Partition.Malformed"
//		INVALID_PARTITION_NO_TYPE = "Invalid.Partition.No.Type"
//		INVALID_PARTITION_INVALID_ORDER = "Invalid.Partition.Invalid.Order"
//		INVALID_PARAMETER_VALUE_RAID_DISK_MISMATCH = "Invalid.Parameter.Value.Raid.Disk.Mismatch"
//		INVALID_PARAMETER_VALUE_RAID_DISK_DISORDER = "Invalid.Parameter.Value.Raid.Disk.Disorder"
//		INVALID_PARAMETER_VALUE_RAID_DISK_SEQUENCE_DUPLICATE = "Invalid.Parameter.Value.Raid.Disk.Sequence.Duplicate"
//		INVALID_PARAMETER_VALUE_RAID_DISK_SEQUENCE_RANGE = "Invalid.Parameter.Value.Raid.Disk.Sequence.Range"
//		RESOURCE_INSUFFICIENT_PRODUCT_SOLD_OUT = "Resource.Insufficient.Product.Sold.out"
//		OPERATION_DENIED_CHARGE_TYPE_NOT_SUPPORT = "Operation.Denied.Charge.Type.Not.Support"
//		RESOURCES_SOLDOUT_INSTANCE_TYPE = "Resource.Soldout.Instance.type"
//		INVALID_CHARGE_TYPE_NOT_SUPPORT = "Invalid.Charge.Type.Not.Support"
//		INVALID_SUBNET_NOT_FOUND = "Invalid.Subnet.Not.Found"
//		INVALID_SUBNET_PRIVATE_IP_INSUFFICIENT = "Invalid.Subnet.Private.Ip.Insufficient"
//		INVALID_SUBNET_ZONE_MISMATCH = "Invalid.Subnet.Zone.Mismatch"
//		INVALID_PARAMETER_SSH_KEY_DUPLICATE = "Invalid.Parameter.Ssh.Key.Duplicate"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
	response = NewCreateInstancesResponse()
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

// DescribeInstanceTypes
// This API is used to query the model configuration of an instance.
//
// Possible error codes to return:
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

// DescribeAvailableResources
// This API is used to describe the status of available resources.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// INVALID_CHARGE_TYPE_NOT_SUPPORT = "Invalid.Charge.Type.Not.Support"
func (c *Client) DescribeAvailableResources(request *DescribeAvailableResourcesRequest) (response *DescribeAvailableResourcesResponse, err error) {
	response = NewDescribeAvailableResourcesResponse()
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

// DescribeImages
// This API is used to view the list of images.
//
// Possible error codes to return:
// INVALID_INSTANCE_TYPE_NOT_FOUND = "Invalid.Instance.Type.Not.Found"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	response = NewDescribeImagesResponse()
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

// DescribeInstances
// This API is used to query the details of instances. You can filter the query results with the instance ID, name, or billing method.
//
// Possible error codes to return:
// INVALID_IMAGE_NOT_FOUND = "Invalid.Image.Not.Found"
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// INVALID_INSTANCE_TYPE_NOT_FOUND = "Invalid.Instance.Type.Not.Found"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	response = NewDescribeInstancesResponse()
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

// StartInstances
// This API is used to start instances.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_RECYCLED = "Operation.Denied.Instance.Recycled"
// OPERATION_DENIED_INSTANCE_NOT_STOPPED = "Operation.Denied.Instance.Not.Stopped"
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

// StopInstances
// This API is used to shut down instances.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_RECYCLED = "Operation.Denied.Instance.Recycled"
// OPERATION_DENIED_INSTANCE_NOT_RUNNING = "Operation.Denied.Instance.Not.Running"
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

// RebootInstances
// This API is used to restart instances.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_RECYCLED = "Operation.Denied.Instance.Recycled"
// OPERATION_DENIED_INSTANCE_STATUS = "Operation.Denied.Instance.Status"
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

func NewReinstallInstanceResponse() (response *ReInstallInstanceResponse) {
	response = &ReInstallInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReinstallInstance
// This API is used to reinstall the operating system of the specified instance.
//
// Possible error codes to return:
// INVALID_PARAMETER_HOSTNAME_EXCEED = "Invalid.Parameter.Hostname.Exceed"
// INVALID_PARAMETER_HOSTNAME_MALFORMED = "Invalid.Parameter.Hostname.Malformed"
// INVALID_PARAMETER_INSTANCE_NAME_EXCEED = "Invalid.Parameter.Instance.Name.Exceed"
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// INVALID_IMAGE_NOT_FOUND = "Invalid.Image.Not.Found"
// INVALID_PARTITION_IMAGE_NOT_SET = "Invalid.Partition.Image.Not.Set"
// INVALID_PARAMETER_VALUE_PASSWORD_MALFORMED = "Invalid.Parameter.Value.Password.Malformed"
// INVALID_PARAMETER_INSTANCE_LOGIN_CONFLICT = "Invalid.Parameter.Instance.Login.Conflict"
// INVALID_PARAMETER_SSH_KEY_MALFORMED = "Invalid.Parameter.Ssh.Key.Malformed"
// INVALID_RAID_CONFIG_FAST_CUSTOM_CONFLICT = "Invalid.Raid.Config.Fast.Custom.Conflict"
// INVALID_INSTANCE_TYPE_RAID_NOT_SUPPORT = "Invalid.Instance.Type.Raid.Not.Support"
// INVALID_PARAMETER_NIC_NAME_CONFLICT = "Invalid.Parameter.Nic.Name.Conflict"
// INVALID_PARAMETER_NIC_NAME_MALFORMED = "Invalid.Parameter.Nic.Name.Malformed"
// INVALID_PARTITION_SIZE_NOT_FULL = "Invalid.Partition.Size.Not.Full"
// INVALID_PARTITION_DUPLICATE_FILE_PATH = "Invalid.Partition.Duplicate.File.Path"
// INVALID_PARTITION_MISSING_REQUIRED_FILE_PATH = "Invalid.Partition.Missing.Required.File.Path"
// INVALID_PARTITION_MALFORMED = "Invalid.Partition.Malformed"
// INVALID_PARTITION_NO_TYPE = "Invalid.Partition.No.Type"
// INVALID_PARTITION_INVALID_ORDER = "Invalid.Partition.Invalid.Order"
// INVALID_PARAMETER_VALUE_RAID_DISK_MISMATCH = "Invalid.Parameter.Value.Raid.Disk.Mismatch"
// INVALID_PARAMETER_VALUE_RAID_DISK_DISORDER = "Invalid.Parameter.Value.Raid.Disk.Disorder"
func (c *Client) ReinstallInstance(request *ReinstallInstanceRequest) (response *ReInstallInstanceResponse, err error) {
	response = NewReinstallInstanceResponse()
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

// TerminateInstance
// This API is used to return an instance.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_RECYCLED = "Operation.Denied.Instance.Recycled"
// OPERATION_DENIED_INSTANCE_CREATING = "Operation.Denied.Instance.Creating"
// OPERATION_DENIED_INSTANCE_STATUS_INSTALLING = "Operation.Denied.Instance.Status.Installing"
// OPERATION_DENIED_INSTANCE_SUBSCRIPTION = "Operation.Denied.Instance.Subscription"
// OPERATION_DENIED_POSTPAID_PROMISE = "Operation.Denied.Postpaid.Promise"
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

// ReleaseInstances
// This API is used to release instances.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_NOT_RECYCLED = "Operation.Denied.Instance.Not.Recycled"
func (c *Client) ReleaseInstances(request *ReleaseInstancesRequest) (response *ReleaseInstancesResponse, err error) {
	response = NewReleaseInstancesResponse()
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

// RenewInstance
// This API is used to renew an instance.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
	response = NewRenewInstanceResponse()
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

// ModifyInstancesAttribute
// The API is used to modify the attributes of one or more instances. Only the instance name to be displayed can be modified for now.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
	response = NewModifyInstancesAttributeResponse()
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

// ModifyInstancesResourceGroup
// This API is used to modify the resource group to which the instances belong.
//
// Possible error codes to return:
// OPERATION_FAILED_RESOURCE_GROUP_NOT_FOUND = "Operation.Failed.Resource.Group.Not.Found"
// OPERATION_FAILED_RESOURCE_NOT_FOUND = "Operation.Failed.Resource.Not.Found"
func (c *Client) ModifyInstancesResourceGroup(request *ModifyInstancesResourceGroupRequest) (response *ModifyInstancesResourceGroupResponse, err error) {
	response = NewModifyInstancesResourceGroupResponse()
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

// DescribeInstanceTraffic
func (c *Client) DescribeInstanceTraffic(request *DescribeInstanceTrafficRequest) (response *DescribeInstanceTrafficResponse, err error) {
	response = NewDescribeInstanceTrafficResponse()
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

// DescribeInstanceInternetStatus
// This API is used to query the bandwidth and traffic package status of an instance.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
func (c *Client) DescribeInstanceInternetStatus(request *DescribeInstanceInternetStatusRequest) (response *DescribeInstanceInternetStatusResponse, err error) {
	response = NewDescribeInstanceInternetStatusResponse()
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

func NewInquiryPriceCreateInstanceRequestResponse() (response *InquiryPriceCreateInstanceResponse) {
	response = &InquiryPriceCreateInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateInstance
// This API is used to query the price of creating instance.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// INVALID_INSTANCE_TYPE_NOT_FOUND = "Invalid.Instance.Type.Not.Found"
// INVALID_INSTANCE_TYPE_ZONE_NO_SELL = "Invalid.Instance.Type.Zone.No.Sell"
// INVALID_INSTANCE_BANDWIDTH_ZONE_NO_SELL = "Invalid.Instance.Bandwidth.Zone.No.Sell"
// INVALID_CHARGE_TYPE_NOT_SUPPORT = "Invalid.Charge.Type.Not.Support"
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
	response = NewInquiryPriceCreateInstanceRequestResponse()
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

// ModifyInstanceBandwidth
// The API is used to modify the bandwidth of an instance.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_NOT_RUNNING = "Operation.Denied.Instance.Not.Running"
// INVALID_PARAMETER_BANDWIDTH_EXCEED = "Invalid.Parameter.Bandwidth.Exceed"
// OPERATION_FAILED_INSTANCE_BANDWIDTH_PROCESSING = "Operation.Failed.Instance.Bandwidth.Processing"
func (c *Client) ModifyInstanceBandwidth(request *ModifyInstanceBandwidthRequest) (response *ModifyInstanceBandwidthResponse, err error) {
	response = NewModifyInstanceBandwidthResponse()
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

// InquiryPriceInstanceBandwidth
// This API is used to query the price of modified instance bandwidth.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INTERNET_CHARGE_TYPE_NOT_SUPPORT = "Operation.Denied.Internet.Charge.Type.Not.Support"
// INVALID_INSTANCE_TYPE_ZONE_NO_SELL = "Invalid.Instance.Type.Zone.No.Sell"
func (c *Client) InquiryPriceInstanceBandwidth(request *InquiryPriceInstanceBandwidthRequest) (response *InquiryPriceInstanceBandwidthResponse, err error) {
	response = NewInquiryPriceInstanceBandwidthResponse()
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

// CancelInstanceBandwidthDowngrade
// The API is used to cancel bandwidth downgrade order.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INTERNET_CHARGE_TYPE_NOT_SUPPORT = "Operation.Denied.Internet.Charge.Type.Not.Support"
// OPERATION_DENIED_DOWNGRADE_NOT_EXIST = "Operation.Denied.Downgrade.Not.Exist"
func (c *Client) CancelInstanceBandwidthDowngrade(request *CancelInstanceBandwidthDowngradeRequest) (response *CancelInstanceBandwidthDowngradeResponse, err error) {
	response = NewCancelInstanceBandwidthDowngradeResponse()
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

// ModifyInstanceTrafficPackage
// The API is used to modify the traffic package of an instance.
//
// Possible error codes to return:
// OPERATION_DENIED_INSTANCE_TRAFFIC_PACKAGE_PROCESSING = "Operation.Denied.Instance.Traffic.Package.Processing"
// OPERATION_DENIED_INTERNET_CHARGE_TYPE_NOT_SUPPORT = "Operation.Denied.Internet.Charge.Type.Not.Support"
// INVALID_PARAMETER_TRAFFIC_PACKAGE_ERROR = "Invalid.Parameter.Traffic.Package.Error"
// OPERATION_FILED_INSTANCE_NOT_EXIST_TRAFFIC_PACKAGE = "Operation.Filed.Instance.Not.Exist.Traffic.Package"
// OPERATION_FAILED_INSTANCE_EXIST_PLAN_TRAFFIC_PACKAGE = "Operation.Failed.Instance.Exist.Plan.Traffic.Package"
// INVALID_PARAMETER_TRAFFIC_PACKAGE_LESS = "Invalid.Parameter.Traffic.Package.Less"
// INVALID_PARAMETER_TRAFFIC_PACKAGE_EXCEED = "Invalid.Parameter.Traffic.Package.Exceed"
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
func (c *Client) ModifyInstanceTrafficPackage(request *ModifyInstanceTrafficPackageRequest) (response *ModifyInstanceTrafficPackageResponse, err error) {
	response = NewModifyInstanceTrafficPackageResponse()
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

// InquiryPriceInstanceTrafficPackage
// This API is used to query the price of modified instance traffic package.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INTERNET_CHARGE_TYPE_NOT_SUPPORT = "Operation.Denied.Internet.Charge.Type.Not.Support"
// INVALID_INSTANCE_TYPE_ZONE_NO_SELL = "Invalid.Instance.Type.Zone.No.Sell"
// INVALID_PARAMETER_TRAFFIC_PACKAGE_EXCEED = "Invalid.Parameter.Traffic.Package.Exceed"
// INVALID_PARAMETER_TRAFFIC_PACKAGE_ERROR = "Invalid.Parameter.Traffic.Package.Error"
func (c *Client) InquiryPriceInstanceTrafficPackage(request *InquiryPriceInstanceTrafficPackageRequest) (response *InquiryPriceInstanceTrafficPackageResponse, err error) {
	response = NewInquiryPriceInstanceTrafficPackageResponse()
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

// CancelInstanceTrafficPackageDowngrade
// The API is used to cancel traffic package downgrade order.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INTERNET_CHARGE_TYPE_NOT_SUPPORT = "Operation.Denied.Internet.Charge.Type.Not.Support"
// OPERATION_DENIED_DOWNGRADE_NOT_EXIST = "Operation.Denied.Downgrade.Not.Exist"
func (c *Client) CancelInstanceTrafficPackageDowngrade(request *CancelInstanceTrafficPackageDowngradeRequest) (response *CancelInstanceTrafficPackageDowngradeResponse, err error) {
	response = NewCancelInstanceTrafficPackageDowngradeResponse()
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

// DescribeInstancesMonitorHealth
// The API is used to cancel traffic package downgrade order.
//
// Possible error codes to return:
func (c *Client) DescribeInstancesMonitorHealth(request *DescribeInstancesMonitorHealthRequest) (response *DescribeInstancesMonitorHealthResponse, err error) {
	response = NewDescribeInstancesMonitorHealthResponse()
	err = c.ApiCall(request, response)
	return
}

//////////////////////////// EIP //////////////////////////////

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

// AllocateEipAddresses
// This API is used to apply for one or multiple elastic IPs.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// INVALID_EIP_TYPE_ZONE_NO_SELL = "Invalid.Eip.Type.Zone.No.Sell"
// MISSING_PARAMETER = "Missing.Parameter"
func (c *Client) AllocateEipAddresses(request *AllocateEipAddressesRequest) (response *AllocateEipAddressesResponse, err error) {
	response = NewAllocateEipAddressesResponse()
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

// DescribeEipAddresses
// This API is used to query the list of elastic IPs. You can query information on elastic IPs according to ID or IP of the elastic IP, and the pricing model.
//
// Possible error codes to return:
func (c *Client) DescribeEipAddresses(request *DescribeEipAddressesRequest) (response *DescribeEipAddressesResponse, err error) {
	response = NewDescribeEipAddressesResponse()
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

// DescribeEipAvailableResources
// This API is used to query elastic IP resources for sale in the zone.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
func (c *Client) DescribeEipAvailableResources(request *DescribeEipAvailableResourcesRequest) (response *DescribeEipAvailableResourcesResponse, err error) {
	response = NewDescribeEipAvailableResourcesResponse()
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

// DescribeInstanceAvailableEipResources
// This API is used to query list of elastic IPs that can be bound to an instance.
//
// Possible error codes to return:
// INVALID_ESXI_NOT_SUPPORT = "Invalid.Esxi.Not.Support"
func (c *Client) DescribeInstanceAvailableEipResources(request *DescribeInstanceAvailableEipResourcesRequest) (response *DescribeInstanceAvailableEipResourcesResponse, err error) {
	response = NewDescribeInstanceAvailableEipResourcesResponse()
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

// TerminateEipAddress
// This API is used to return an elastic IP.
//
// Possible error codes to return:
// INVALID_EIP_NOT_FOUND = "Invalid.Eip.Not.Found"
// OPERATION_DENIED_EIP_SUBSCRIPTION = "Operation.Denied.Eip.Subscription"
// OPERATION_DENIED_EIP_RECYCLED = "Operation.Denied.Eip.Recycled"
func (c *Client) TerminateEipAddress(request *TerminateEipAddressRequest) (response *TerminateEipAddressResponse, err error) {
	response = NewTerminateEipAddressResponse()
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

// ReleaseEipAddresses
// This API is used to release one or multiple elastic IPs.
//
// Possible error codes to return:
// OPERATION_DENIED_EIP_NOT_RECYCLED = "Operation.Denied.Eip.Not.Recycled"
func (c *Client) ReleaseEipAddresses(request *ReleaseEipAddressesRequest) (response *ReleaseEipAddressesResponse, err error) {
	response = NewReleaseEipAddressesResponse()
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

// RenewEipAddress
// This API is used to renew an elastic IP.
//
// Possible error codes to return:
// INVALID_EIP_NOT_FOUND = "Invalid.Eip.Not.Found"
func (c *Client) RenewEipAddress(request *RenewEipAddressRequest) (response *RenewEipAddressResponse, err error) {
	response = NewRenewEipAddressResponse()
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

// AssociateEipAddress
// This API is used to bind an elastic IP to an instance in the same zone.
//
// Possible error codes to return:
// INVALID_EIP_NOT_FOUND = "Invalid.Eip.Not.Found"
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_RECYCLED = "Operation.Denied.Instance.Recycled"
// OPERATION_DENIED_INSTANCE_NOT_RUNNING = "Operation.Denied.Instance.Not.Running"
// OPERATION_DENIED_EIP_ZONE_NOT_SAME = "Operation.Denied.Eip.Zone.Not.Same"
// FAILED_OPERATION_FOR_RECYCLE_RESOURCE = "Failed.Operation.For.Recycle.Resource"
// OPERATION_DENIED_EIP_STATUS_NOT_AVAILABLE = "Operation.Denied.Eip.Status.Not.Available"
// OPERATION_DENIED_EIP_ESXI_INSTANCE_ASSIGN = "Operation.Denied.Eip.Esxi.Instance.Assign"
// OPERATION_DENIED_EIP_INSTANCE_EXCEED_LIMIT = "Operation.Denied.Eip.Instance.Exceed.Limit"
func (c *Client) AssociateEipAddress(request *AssociateEipAddressRequest) (response *AssociateEipAddressResponse, err error) {
	response = NewAssociateEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnassociateEipAddressRequest() (request *UnassociateEipAddressRequest) {
	request = &UnassociateEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnAssociateEipAddress")

	return
}

func NewUnassociateEipAddressResponse() (response *UnassociateEipAddressResponse) {
	response = &UnassociateEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnassociateEipAddress
// This API is used to unbind an elastic IP from an instance.
//
// Possible error codes to return:
// INVALID_EIP_NOT_FOUND = "Invalid.Eip.Not.Found"
// FAILED_OPERATION_FOR_RECYCLE_RESOURCE = "Failed.Operation.For.Recycle.Resource"
// OPERATION_DENIED_EIP_STATUS_NOT_SUPPORT = "Operation.Denied.Eip.Status.Not.Support"
func (c *Client) UnassociateEipAddress(request *UnassociateEipAddressRequest) (response *UnassociateEipAddressResponse, err error) {
	response = NewUnassociateEipAddressResponse()
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

// InquiryPriceCreateEipAddress
// This API is used to query the price of creating elastic IP.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// INVALID_EIP_TYPE_ZONE_NO_SELL = "Invalid.Eip.Type.Zone.No.Sell"
// MISSING_PARAMETER = "Missing.Parameter"
func (c *Client) InquiryPriceCreateEipAddress(request *InquiryPriceCreateEipAddressRequest) (response *InquiryPriceCreateEipAddressResponse, err error) {
	response = NewInquiryPriceCreateEipAddressResponse()
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

// ModifyEipAddressesResourceGroup
// This API is used to modify the resource group to which the EIP belong.
//
// Possible error codes to return:
// OPERATION_FAILED_RESOURCE_GROUP_NOT_FOUND = "Operation.Failed.Resource.Group.Not.Found"
// OPERATION_FAILED_RESOURCE_NOT_FOUND = "Operation.Failed.Resource.Not.Found"
func (c *Client) ModifyEipAddressesResourceGroup(request *ModifyEipAddressesResourceGroupRequest) (response *ModifyEipAddressesResourceGroupResponse, err error) {
	response = NewModifyEipAddressesResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

//////////////////////////// DDoS IP //////////////////////////////

func NewDescribeDdosAvailableResourcesRequest() (request *DescribeDdosIpAvailableResourcesRequest) {
	request = &DescribeDdosIpAvailableResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDdosIpAvailableResources")

	return
}

func NewDescribeDdosAvailableResourcesResponse() (response *DescribeDdosIpAvailableResourcesResponse) {
	response = &DescribeDdosIpAvailableResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeDdosAvailableResources
// This API is used to query DDoS protected IP resources for sale in the zone.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
func (c *Client) DescribeDdosAvailableResources(request *DescribeDdosIpAvailableResourcesRequest) (response *DescribeDdosIpAvailableResourcesResponse, err error) {
	response = NewDescribeDdosAvailableResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceAvailableDdosResourcesRequest() (request *DescribeInstanceAvailableDdosResourcesRequest) {
	request = &DescribeInstanceAvailableDdosResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceAvailableDdosIpResources")

	return
}

func NewDescribeInstanceAvailableDdosResourcesResponse() (response *DescribeInstanceAvailableDdosResourcesResponse) {
	response = &DescribeInstanceAvailableDdosResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceAvailableDdosResources
// This API is used to query list of DDoS protected IPs that can be bound to an instance.
//
// Possible error codes to return:
// INVALID_ESXI_NOT_SUPPORT = "Invalid.Esxi.Not.Support"
func (c *Client) DescribeInstanceAvailableDdosResources(request *DescribeInstanceAvailableDdosResourcesRequest) (response *DescribeInstanceAvailableDdosResourcesResponse, err error) {
	response = NewDescribeInstanceAvailableDdosResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAllocateDdosIpAddressesRequest() (request *AllocateDdosIpAddressesRequest) {
	request = &AllocateDdosIpAddressesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AllocateDdosIpAddresses")

	return
}

func NewAllocateDdosIpAddressesResponse() (response *AllocateDdosIpAddressesResponse) {
	response = &AllocateDdosIpAddressesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AllocateDdosIpAddresses
// This API is used to apply for one or multiple DDoS protected IPs.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// INVALID_DDOS_IP_TYPE_ZONE_NO_SELL = "Invalid.Ddos.Ip.Type.Zone.No.Sell"
// MISSING_PARAMETER = "Missing.Parameter"
func (c *Client) AllocateDdosIpAddresses(request *AllocateDdosIpAddressesRequest) (response *AllocateDdosIpAddressesResponse, err error) {
	response = NewAllocateDdosIpAddressesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDdosIpAddressesRequest() (request *DescribeDdosIpAddressesRequest) {
	request = &DescribeDdosIpAddressesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDdosIpAddresses")

	return
}

func NewDescribeDdosIpAddressesResponse() (response *DescribeDdosIpAddressesResponse) {
	response = &DescribeDdosIpAddressesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeDdosIpAddresses
// This API is used to query the list of DDoS protected IPs. You can query information on DDoS protected IPs according to ID or IP of the DDoS protected IP, and the pricing model.
//
// Possible error codes to return:
func (c *Client) DescribeDdosIpAddresses(request *DescribeDdosIpAddressesRequest) (response *DescribeDdosIpAddressesResponse, err error) {
	response = NewDescribeDdosIpAddressesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminateDdosIpAddressRequest() (request *TerminateDdosIpAddressRequest) {
	request = &TerminateDdosIpAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminateDdosIpAddress")

	return
}

func NewTerminateDdosIpAddressResponse() (response *TerminateDdosIpAddressResponse) {
	response = &TerminateDdosIpAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TerminateDdosIpAddress
// This API is used to return a DDoS protected IP.
//
// Possible error codes to return:
// INVALID_DDOS_IP_NOT_FOUND = "Invalid.Ddos.Ip.Not.Found"
// OPERATION_DENIED_DDOS_IP_SUBSCRIPTION = "Operation.Denied.Ddos.Ip.Subscription"
// OPERATION_DENIED_DDOS_IP_RECYCLED = "Operation.Denied.Ddos.Ip.Recycled"
func (c *Client) TerminateDdosIpAddress(request *TerminateDdosIpAddressRequest) (response *TerminateDdosIpAddressResponse, err error) {
	response = NewTerminateDdosIpAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseDdosIpAddressesRequest() (request *ReleaseDdosIpAddressesRequest) {
	request = &ReleaseDdosIpAddressesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseDdosIpAddresses")

	return
}

func NewReleaseDdosIpAddressesResponse() (response *ReleaseDdosIpAddressesResponse) {
	response = &ReleaseDdosIpAddressesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReleaseDdosIpAddresses
// This API is used to release one or multiple DDoS protected IPs.
//
// Possible error codes to return:
// OPERATION_DENIED_DDOS_IP_NOT_RECYCLED = "Operation.Denied.Ddos.Ip.Not.Recycled"
func (c *Client) ReleaseDdosIpAddresses(request *ReleaseDdosIpAddressesRequest) (response *ReleaseDdosIpAddressesResponse, err error) {
	response = NewReleaseDdosIpAddressesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewDdosIpAddressRequest() (request *RenewDdosIpAddressRequest) {
	request = &RenewDdosIpAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewDdosIpAddress")

	return
}

func NewRenewDdosIpAddressResponse() (response *RenewDdosIpAddressResponse) {
	response = &RenewDdosIpAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewDdosIpAddress
// This API is used to renew a DDoS protected IP.
//
// Possible error codes to return:
// INVALID_DDOS_IP_NOT_FOUND = "Invalid.Ddos.Ip.Not.Found"
func (c *Client) RenewDdosIpAddress(request *RenewDdosIpAddressRequest) (response *RenewDdosIpAddressResponse, err error) {
	response = NewRenewDdosIpAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssociateDdosIpAddressRequest() (request *AssociateDdosIpAddressRequest) {
	request = &AssociateDdosIpAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssociateDdosIpAddress")

	return
}

func NewAssociateDdosIpAddressResponse() (response *AssociateDdosIpAddressResponse) {
	response = &AssociateDdosIpAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssociateDdosIpAddress
// This API is used to bind a DDoS protected IP to an instance in the same zone.
//
// Possible error codes to return:
// INVALID_DDOS_IP_NOT_FOUND = "Invalid.Ddos.Ip.Not.Found"
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_RECYCLED = "Operation.Denied.Instance.Recycled"
// OPERATION_DENIED_INSTANCE_NOT_RUNNING = "Operation.Denied.Instance.Not.Running"
// OPERATION_DENIED_DDOS_IP_ZONE_NOT_SAME = "Operation.Denied.Ddos.Ip.Zone.Not.Same"
// FAILED_OPERATION_FOR_RECYCLE_RESOURCE = "Failed.Operation.For.Recycle.Resource"
// OPERATION_DENIED_DDOS_IP_STATUS_NOT_AVAILABLE = "Operation.Denied.Ddos.Ip.Status.Not.Available"
// OPERATION_DENIED_DDOS_IP_ESXI_INSTANCE_ASSIGN = "Operation.Denied.Ddos.Ip.Esxi.Instance.Assign"
func (c *Client) AssociateDdosIpAddress(request *AssociateDdosIpAddressRequest) (response *AssociateDdosIpAddressResponse, err error) {
	response = NewAssociateDdosIpAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnassociateDdosIpAddressRequest() (request *UnassociateDdosIpAddressRequest) {
	request = &UnassociateDdosIpAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnAssociateDdosIpAddress")

	return
}

func NewUnassociateDdosIpAddressResponse() (response *UnassociateDdosIpAddressResponse) {
	response = &UnassociateDdosIpAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnassociateDdosIpAddress
// This API is used to unbind a DDoS protected IP from an instance.
//
// Possible error codes to return:
// INVALID_DDOS_IP_NOT_FOUND = "Invalid.Ddos.Ip.Not.Found"
// FAILED_OPERATION_FOR_RECYCLE_RESOURCE = "Failed.Operation.For.Recycle.Resource"
// OPERATION_DENIED_DDOS_IP_STATUS_NOT_SUPPORT = "Operation.Denied.Ddos.Ip.Status.Not.Support"
func (c *Client) UnassociateDdosIpAddress(request *UnassociateDdosIpAddressRequest) (response *UnassociateDdosIpAddressResponse, err error) {
	response = NewUnassociateDdosIpAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateDdosIpAddressRequest() (request *InquiryPriceCreateDdosIpAddressRequest) {
	request = &InquiryPriceCreateDdosIpAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateDdosIpAddress")

	return
}

func NewInquiryPriceCreateDdosIpAddressResponse() (response *InquiryPriceCreateDdosIpAddressResponse) {
	response = &InquiryPriceCreateDdosIpAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateDdosIpAddress
// This API is used to query the price of creating DDoS protected IP.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// INVALID_DDOS_IP_TYPE_ZONE_NO_SELL = "Invalid.Ddos.Ip.Type.Zone.No.Sell"
// MISSING_PARAMETER = "Missing.Parameter"
func (c *Client) InquiryPriceCreateDdosIpAddress(request *InquiryPriceCreateDdosIpAddressRequest) (response *InquiryPriceCreateDdosIpAddressResponse, err error) {
	response = NewInquiryPriceCreateDdosIpAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyDdosIpAddressesResourceGroupRequest() (request *ModifyDdosIpAddressesResourceGroupRequest) {
	request = &ModifyDdosIpAddressesResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyDdosIpAddressesResourceGroup")

	return
}

func NewModifyDdosIpAddressesResourceGroupResponse() (response *ModifyDdosIpAddressesResourceGroupResponse) {
	response = &ModifyDdosIpAddressesResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyDdosIpAddressesResourceGroup
// This API is used to modify the resource group to which the DDoS IP belong.
//
// Possible error codes to return:
// OPERATION_FAILED_RESOURCE_GROUP_NOT_FOUND = "Operation.Failed.Resource.Group.Not.Found"
// OPERATION_FAILED_RESOURCE_NOT_FOUND = "Operation.Failed.Resource.Not.Found"
func (c *Client) ModifyDdosIpAddressesResourceGroup(request *ModifyDdosIpAddressesResourceGroupRequest) (response *ModifyDdosIpAddressesResourceGroupResponse, err error) {
	response = NewModifyDdosIpAddressesResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

//////////////////////////// VPC & Subnet //////////////////////////////

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

// DescribeVpcAvailableRegions
// This API is used to query information about availability regions for VPCs.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
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

// CreateVpc
// This API is used to create a VPC.
//
// Possible error codes to return:
// INVALID_VPC_REGION_NOT_FOUND = "Invalid.Vpc.Region.Not.Found"
// OPERATION_DENIED_VPC_REGION_EXCEED_LIMIT = "Operation.Denied.Vpc.Region.Exceed.Limit"
// INVALID_PARAMETER_VPC_LAN_IP_NETMASK = "Invalid.Parameter.Vpc.Lan.Ip.Netmask"
// INVALID_PARAMETER_VPC_CIDR_BLOCK = "Invalid.Parameter.Vpc.Cidr.Block"
// INVALID_PARAMETER_VPC_LAN_IP = "Invalid.Parameter.Vpc.Lan.Ip"
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

// DescribeVpcs
// This API is used to query the details of one or more VPCs. You can filter the query results with the VPC ID, VPC name, subnet ID and availability region ID.
//
// Possible error codes to return:
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

// ModifyVpcsAttribute
// The API is used to modify the attributes of one or more VPCs. Only the VPC name to be displayed can be modified for now.
//
// Possible error codes to return:
// OPERATION_FAILED_RESOURCE_NOT_FOUND = "Operation.Failed.Resource.Not.Found"
func (c *Client) ModifyVpcsAttribute(request *ModifyVpcsAttributeRequest) (response *ModifyVpcsAttributeResponse, err error) {
	response = NewModifyVpcsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyVpcsResourceGroupRequest() (request *ModifyVpcsResourceGroupRequest) {
	request = &ModifyVpcsResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyVpcsResourceGroup")

	return
}

func NewModifyVpcsResourceGroupResponse() (response *ModifyVpcsResourceGroupResponse) {
	response = &ModifyVpcsResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyVpcsResourceGroup(request *ModifyVpcsResourceGroupRequest) (response *ModifyVpcsResourceGroupResponse, err error) {
	response = NewModifyVpcsResourceGroupResponse()
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

// DeleteVpc
// This API is used to delete a VPC.
//
// Possible error codes to return:
// OPERATION_DENIED_VPC_STATUS_NOT_SUPPORT = "Operation.Denied.Vpc.Status.Not.Support"
// OPERATION_DENIED_VPC_EXIST_SUBNET_ASSIGN = "Operation.Denied.Vpc.Exist.Subnet.Assign"
// OPERATION_DENIED_VPC_EXIST_INSTANCE = "Operation.Denied.Vpc.Exist.Instance"
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
	response = NewDeleteVpcResponse()
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

// CreateSubnet
// This API is used to create a subnet.
//
// Possible error codes to return:
// OPERATION_DENIED_ZONE_NOT_SUPPORT_SUBNET = "Operation.Denied.Zone.Not.Support.Subnet"
// INVALID_VPC_NOT_FOUND = "Invalid.Vpc.Not.Found"
// OPERATION_DENIED_VPC_STATUS_NOT_SUPPORT = "Operation.Denied.Vpc.Status.Not.Support"
// OPERATION_DENIED_NO_AVAILABLE_VPC_REGION = "Operation.Denied.No.Available.Vpc.Region"
// OPERATION_DENIED_ZONE_NOT_BELONG_VPC = "Operation.Denied.Zone.Not.Belong.Vpc"
// INVALID_PARAMETER_SUBNET_CIDR_BLOCK = "Invalid.Parameter.Subnet.Cidr.Block"
// INVALID_PARAMETER_SUBNET_LAN_IP_NETMASK = "Invalid.Parameter.Subnet.Lan.Ip.Netmask"
// INVALID_PARAMETER_SUBNET_LAN_IP = "Invalid.Parameter.Subnet.Lan.Ip"
// OPERATION_DENIED_SUBNET_EXCEED_LIMIT = "Operation.Denied.Subnet.Exceed.Limit"
// OPERATION_DENIED_VPC_ZONE_SUBNET_EXCEED_LIMIT = "Operation.Denied.Vpc.Zone.Subnet.Exceed.Limit"
// INVALID_PARAMETER_VPC_SUBNET_OVER_LAP = "Invalid.Parameter.Vpc.Subnet.Over.Lan"
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// INVALID_PARAMETER_SUBNET_CIDR_NOT_BELONG_VPC = "Invalid.Parameter.Subnet.Cidr.Not.Belong.Vpc"
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

// DescribeSubnets
// This API is used to query the details of one or more subnets. You can filter the query results with the subnet ID, VPC ID, zone ID and subnet name.
//
// Possible error codes to return:
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
	response = NewDescribeSubnetsResponse()
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

// DeleteSubnet
// This API is used to delete a subnet.
//
// Possible error codes to return:
// OPERATION_FAILED_RESOURCE_NOT_FOUND = "Operation.Failed.Resource.Not.Found"
// OPERATION_DENIED_SUBNET_EXIST_INSTANCE = "Operation.Denied.Subnet.Exist.Instance"
// OPERATION_DENIED_SUBNET_STATUS_NOT_SUPPORT = "Operation.Denied.Subnet.Status.Not.Support"
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

// ModifySubnetsAttribute
// The API is used to modify the attributes of one or more subnets. Only the subnet name to be displayed can be modified for now.
//
// Possible error codes to return:
// OPERATION_FAILED_RESOURCE_NOT_FOUND = "Operation.Failed.Resource.Not.Found"
func (c *Client) ModifySubnetsAttribute(request *ModifySubnetsAttributeRequest) (response *ModifySubnetsAttributeResponse, err error) {
	response = NewModifySubnetsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySubnetsResourceGroupRequest() (request *ModifySubnetsResourceGroupRequest) {
	request = &ModifySubnetsResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySubnetsResourceGroup")

	return
}

func NewModifySubnetsResourceGroupResponse() (response *ModifySubnetsResourceGroupResponse) {
	response = &ModifySubnetsResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifySubnetsResourceGroup(request *ModifySubnetsResourceGroupRequest) (response *ModifySubnetsResourceGroupResponse, err error) {
	response = NewModifySubnetsResourceGroupResponse()
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

// UnAssociateSubnetInstance
// This API is used to remove an instance from a subnet.
//
// Possible error codes to return:
// INVALID_SUBNET_NOT_FOUND = "Invalid.Subnet.Not.Found"
// OPERATION_DENIED_SUBNET_ASSOCIATING_INSTANCE = "Operation.Denied.Subnet.Associating.Instance"
func (c *Client) UnAssociateSubnetInstance(request *UnAssociateSubnetInstanceRequest) (response *UnAssociateSubnetInstanceResponse, err error) {
	response = NewUnAssociateSubnetInstanceResponse()
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

// AssociateSubnetInstances
// This API is used to add one or more instances into a subnet and assign private IPs to instances.
//
// Possible error codes to return:
// OPERATION_DENIED_INSTANCE_NOT_SUPPORT_SUBNET = "Operation.Denied.Instance.Not.Support.Subnet"
// OPERATION_DENIED_INSTANCE_NOT_RUNNING = "Operation.Denied.Instance.Not.Running"
// INVALID_SUBNET_NOT_FOUND = "Invalid.Subnet.Not.Found"
// OPERATION_DENIED_SUBNET_STATUS_NOT_SUPPORT = "Operation.Denied.Subnet.Status.Not.Support"
// OPERATION_DENIED_DIFFERENT_ZONE = "Operation.Denied.Different.Zone"
// OPERATION_DENIED_SUBNET_EXIST_INSTANCE = "Operation.Denied.Subnet.Exist.Instance"
// INVALID_PARAMETER_DUPLICATE_LAN_IP = "Invalid.Parameter.Duplicate.Lan.Ip"
// INVALID_PARAMETER_LAN_IP_NOT_SUPPORT = "Invalid.Parameter.Lan.Ip.Not.Support"
// OPERATION_DENIED_IP_ASSOCIATED_INSTANCE = "Operation.Denied.Ip.Associated.Instance"
// OPERATION_DENIED_SUBNET_NOT_REPEAT_INSTANCE = "Operation.Denied.Subnet.Not.Repeat.Instance"
func (c *Client) AssociateSubnetInstances(request *AssociateSubnetInstancesRequest) (response *AssociateSubnetInstancesResponse, err error) {
	response = NewAssociateSubnetInstancesResponse()
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

// AssociateVpcSubnet
// This API is used to add a subnet into a VPC.
//
// Possible error codes to return:
// INVALID_VPC_NOT_FOUND = "Invalid.Vpc.Not.Found"
// INVALID_SUBNET_NOT_FOUND = "Invalid.Subnet.Not.Found"
// OPERATION_DENIED_SUBNET_STATUS_NOT_SUPPORT = "Operation.Denied.Subnet.Status.Not.Support"
// OPERATION_DENIED_SUBNET_ASSOCIATED_OTHER_VPC = "Operation.Denied.Subnet.Associated.Other.Vpc"
// OPERATION_DENIED_VPC_STATUS_NOT_SUPPORT = "Operation.Denied.Vpc.Status.Not.Support"
// OPERATION_DENIED_NO_AVAILABLE_VPC_REGION = "Operation.Denied.No.Available.Vpc.Region"
// OPERATION_DENIED_ZONE_NOT_BELONG_VPC = "Operation.Denied.Zone.Not.Belong.Vpc"
// OPERATION_DENIED_VPC_ZONE_SUBNET_EXCEED_LIMIT = "Operation.Denied.Vpc.Zone.Subnet.Exceed.Limit"
// INVALID_PARAMETER_VPC_SUBNET_OVER_LAP = "Invalid.Parameter.Vpc.Subnet.Over.Lan"
func (c *Client) AssociateVpcSubnet(request *AssociateVpcSubnetRequest) (response *AssociateVpcSubnetResponse, err error) {
	response = NewAssociateVpcSubnetResponse()
	err = c.ApiCall(request, response)
	return
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

// DescribeSubnetAvailableResources
// This API is used to query information about zones for subnets.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
func (c *Client) DescribeSubnetAvailableResources(request *DescribeSubnetAvailableResourcesRequest) (response *DescribeSubnetAvailableResourcesResponse, err error) {
	response = NewDescribeSubnetAvailableResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

//////////////////////////// CIDR Block //////////////////////////////

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

// DescribeCidrBlocks
// This API is used to query the details of CIDR blocks.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
func (c *Client) DescribeCidrBlocks(request *DescribeCidrBlocksRequest) (response *DescribeCidrBlocksResponse, err error) {
	response = NewDescribeCidrBlocksResponse()
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

// DescribeCidrBlockIps
// This API is used to query the IPs of a CIDR block.
//
// Possible error codes to return:
func (c *Client) DescribeCidrBlockIps(request *DescribeCidrBlockIpsRequest) (response *DescribeCidrBlockIpsResponse, err error) {
	response = NewDescribeCidrBlockIpsResponse()
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

// DescribeAvailableIpv4Resources
// This API is used to query available IPv4 CIDR blocks.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
func (c *Client) DescribeAvailableIpv4Resources(request *DescribeAvailableIpv4ResourcesRequest) (response *DescribeAvailableIpv4ResourcesResponse, err error) {
	response = NewDescribeAvailableIpv4ResourcesResponse()
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

// DescribeAvailableIpv6Resources
// This API is used to query available IPv6 CIDR blocks.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
func (c *Client) DescribeAvailableIpv6Resources(request *DescribeAvailableIpv6ResourcesRequest) (response *DescribeAvailableIpv6ResourcesResponse, err error) {
	response = NewDescribeAvailableIpv6ResourcesResponse()
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

// DescribeInstanceAvailableCidrBlock
// This API is used to query available CIDR blocks for the instance.
//
// Possible error codes to return:
func (c *Client) DescribeInstanceAvailableCidrBlock(request *DescribeInstanceAvailableCidrBlockRequest) (response *DescribeInstanceAvailableCidrBlockResponse, err error) {
	response = NewDescribeInstanceAvailableCidrBlockResponse()
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

// InquiryPriceCreateIpv4Block
// This API is used to query the price of creating IPv4 CIDR blocks.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// OPERATION_DENIED_UNAVAILABLE_NETMASK = "Operation.Denied.Unavailable.Netmask"
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

// CreateIpv4Block
// This API is used to create one or more IPv4 CIDR blocks.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// OPERATION_DENIED_NETMASK_OUT_OF_STOCK = "Operation.Denied.Netmask.Out.Of.Stock"
// RESOURCE_INSUFFICIENT_PRODUCT_SOLD_OUT = "Resource.Insufficient.Product.Sold.out"
func (c *Client) CreateIpv4Block(request *CreateIpv4BlockRequest) (response *CreateIpv4BlockResponse, err error) {
	response = NewCreateIpv4BlockResponse()
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

// CreateIpv6Block
// This API is used to create one or more IPv6 CIDR blocks.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
// RESOURCE_INSUFFICIENT_PRODUCT_SOLD_OUT = "Resource.Insufficient.Product.Sold.out"
func (c *Client) CreateIpv6Block(request *CreateIpv6BlockRequest) (response *CreateIpv6BlockResponse, err error) {
	response = NewCreateIpv6BlockResponse()
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

// ModifyCidrBlocksAttribute
// The API is used to modify the attributes of one or more CIDR blocks. Only the CIDR block name to be displayed can be modified for now.
//
// Possible error codes to return:
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

// RenewCidrBlock
// This API is used to renew a CIDR block.
//
// Possible error codes to return:
// OPERATION_DENIED_CIDRBLOCK_TYPE_NOT_SUPPORTED = "Operation.Denied.CidrBlock.Type.Not.Support"
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

// TerminateCidrBlock
// This API is used to return a CIDR block.
//
// Possible error codes to return:
// OPERATION_DENIED_CIDRBLOCK_SUBSCRIPTION = "Operation.Denied.CidrBlock.Subscription"
// OPERATION_DENIED_INSTANCE_EXIST = "Operation.Denied.Instance.Exist"
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

// ReleaseCidrBlocks
// This API is used to release one or more IPv4 CIDR blocks.
//
// Possible error codes to return:
// OPERATION_DENIED_CIDRBLOCK_NOT_RECYCLED = "Operation.Denied.CidrBlock.Not.Recycled"
func (c *Client) ReleaseCidrBlocks(request *ReleaseCidrBlocksRequest) (response *ReleaseCidrBlocksResponse, err error) {
	response = NewReleaseCidrBlocksResponse()
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

// BindCidrBlockIps
// This API is used to bind one or more CIDR block IPs to instances.
//
// Possible error codes to return:
// INVALID_CIDRBLOCK_NOT_FOUND = "Invalid.CidrBlock.Not.Found"
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// OPERATION_DENIED_INSTANCE_NOT_RUNNING = "Operation.Denied.Instance.Not.Running"
// OPERATION_DENIED_DIFFERENT_ZONE = "Operation.Denied.Different.Zone"
// OPERATION_DENIED_CIDRBLOCK_RECYCLED = "Operation.Denied.CidrBlock.Recycled"
// OPERATION_DENIED_CIDRBLOCK_IP_COUNT_REACH_LIMIT = "Operation.Denied.CidrBlock.Ip.Count.Reach.Limit"
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

// UnbindCidrBlockIps
// This API is used to unbind one or more CIDR block IPs from instances.
//
// Possible error codes to return:
// OPERATION_DENIED_IP_NOT_BELONG = "Operation.Denied.Ip.Not.Belong"
// OPERATION_DENIED_INVALID_STATUS = "Operation.Denied.Invalid.Status"
// OPERATION_DENIED_CIDRBLOCK_RECYCLED = "Operation.Denied.CidrBlock.Recycled"
func (c *Client) UnbindCidrBlockIps(request *UnbindCidrBlockIpsRequest) (response *UnbindCidrBlockIpsResponse, err error) {
	response = NewUnbindCidrBlockIpsResponse()
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

// DescribeManagedInstances 查询一台或多台实例的信息。用户可以根据实例ID、实例名称等条件来查询实例的详细信息。
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

// DescribeManagedInstanceTraffic
func (c *Client) DescribeManagedInstanceTraffic(request *DescribeManagedInstanceTrafficRequest) (response *DescribeManagedInstanceTrafficResponse, err error) {
	response = NewDescribeManagedInstanceTrafficResponse()
	err = c.ApiCall(request, response)
	return
}