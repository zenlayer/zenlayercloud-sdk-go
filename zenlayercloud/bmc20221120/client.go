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

func (c *Client) ModifyInstancesResourceGroup(request *ModifyInstancesResourceGroupRequest) (response *ModifyInstancesResourceGroupResponse, err error) {
	response = NewModifyInstancesResourceGroupResponse()
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

func (c *Client) CancelInstanceTrafficPackageDowngrade(request *CancelInstanceTrafficPackageDowngradeRequest) (response *CancelInstanceTrafficPackageDowngradeResponse, err error) {
	response = NewCancelInstanceTrafficPackageDowngradeResponse()
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

func (c *Client) UnbindCidrBlockIps(request *UnbindCidrBlockIpsRequest) (response *UnbindCidrBlockIpsResponse, err error) {
	response = NewUnbindCidrBlockIpsResponse()
	err = c.ApiCall(request, response)
	return
}
