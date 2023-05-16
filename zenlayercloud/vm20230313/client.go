package vm

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2023-03-13"
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

func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateImage")

	return
}

// CreateImages
// This API is used to create a custom image.
//
// Possible error codes to return:
// INVALID_INSTANCE_NOT_FOUND = "Invalid.Instance.Not.Found"
// UNSUPPORTED_OPERATION_INSTANCE_STATE_STARTING = "Unsupported.Operation.Instance.State.Starting"
// LIMIT_EXCEEDED_IMAGE_QUOTA = "Limit.Exceeded.Image.Quota"
// UNSUPPORTED_OPERATION_ZONE_NOT_SUPPORT = "Unsupported.Operation.Zone.Not.Support"
// UNSUPPORTED_OPERATION_DISK_UNAVAILABLE = "Unsupported.Operation.Disk.Unavailable"
// UNSUPPORTED_OPERATION_DISK_MAKING_IMAGE = "Unsupported.Operation.Disk.Making.Image"
func (c *Client) CreateImages(request *CreateImageRequest) (response *CreateImageResponse, err error) {
	response = NewCreateImageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
	request = &DescribeZonesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeZones")

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

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
	response = &DescribeZonesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeImages")

	return
}

// DescribeImages
// This API is used to query the details of images.
//
// Possible error codes to return:
// INVALID_ZONE_NOT_FOUND = "Invalid.Zone.Not.Found"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	response = NewDescribeImagesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeImageQuotaRequest() (request *DescribeImageQuotaRequest) {
	request = &DescribeImageQuotaRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeImageQuota")

	return
}

// DescribeImageQuota
// This API is used to query the quota of images in the zone.
//
// Possible error codes to return:
// INVALID_REGION_NOT_FOUND = "Invalid.Region.Not.Found"
func (c *Client) DescribeImageQuota(request *DescribeImageQuotaRequest) (response *DescribeImageQuotaResponse, err error) {
	response = NewDescribeImageQuotaResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeImageQuotaResponse() (response *DescribeImageQuotaResponse) {
	response = &DescribeImageQuotaResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyImagesAttributesRequest() (request *ModifyImagesAttributesRequest) {
	request = &ModifyImagesAttributesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyImagesAttributes")

	return
}

// ModifyImagesAttributes
// The API is used to modify the attributes of one or more images.
//
// Possible error codes to return:
// INVALID_IMAGE_NOT_FOUND = "Invalid.Image.Not.Found"
// INVALID_IMAGE_STATUS = "Invalid.Image.Status"
func (c *Client) ModifyImagesAttributes(request *ModifyImagesAttributesRequest) (response *ModifyImagesAttributesResponse, err error) {
	response = NewModifyImagesAttributesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyImagesAttributesResponse() (response *ModifyImagesAttributesResponse) {
	response = &ModifyImagesAttributesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
	request = &DeleteImagesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteImages")

	return
}

// DeleteImages
// The API is used to delete one or more images.
//
// Possible error codes to return:
// INVALID_IMAGE_NOT_FOUND = "Invalid.Image.Not.Found"
// INVALID_IMAGE_STATUS = "Invalid.Image.Status"
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
	response = NewDeleteImagesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
	response = &DeleteImagesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

//////////////////////////// SecurityGroup //////////////////////////////

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

// DescribeSecurityGroups
// This API is used to query the details of security groups. You can filter the query results with the security group ID or name.
//
// Possible error codes to return:
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

// ModifySecurityGroupsAttribute
// The API is used to modify the attributes of one or more security groups. Only the security group name to be displayed can be modified for now.
//
// Possible error codes to return:
// OPERATION_DENIED_SECURITY_GROUP_STATUS_NOT_AVAILABLE = "Operation.Denied.Security.Group.Status.Not.Available"
// INVALID_SECURITY_GROUP_NOT_FOUND = "Invalid.Security.Group.Not.Found"
// OPERATION_DENIED_DEFAULT_SECURITY_GROUP_NOT_SUPPORT = "Operation.Denied.Default.Security.Group.Not.Support"
func (c *Client) ModifySecurityGroupsAttribute(request *ModifySecurityGroupsAttributeRequest) (response *ModifySecurityGroupsAttributeResponse, err error) {
	response = NewModifySecurityGroupsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceAvailableSecurityGroupResourceRequest() (request *DescribeInstanceAvailableSecurityGroupResourceRequest) {
	request = &DescribeInstanceAvailableSecurityGroupResourceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceAvailableSecurityGroupResources")

	return
}

func NewDescribeInstanceAvailableSecurityGroupResourceResponse() (response *DescribeInstanceAvailableSecurityGroupResourceResponse) {
	response = &DescribeInstanceAvailableSecurityGroupResourceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceAvailableSecurityGroupResource
// The API is used to obtain security groups available to be applied to an instance.
//
// Possible error codes to return:
func (c *Client) DescribeInstanceAvailableSecurityGroupResource(request *DescribeInstanceAvailableSecurityGroupResourceRequest) (response *DescribeInstanceAvailableSecurityGroupResourceResponse, err error) {
	response = NewDescribeInstanceAvailableSecurityGroupResourceResponse()
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

// CreateSecurityGroup
// This API is used to create a security group.
//
// Possible error codes to return:
// INVALID_PARAMETER_SECURITY_GROUP_RULE_ID_NOT_ALLOW = "Invalid.Parameter.Security.Group.Rule.Id.Not.Allow"
// INVALID_PARAMETER_SECURITY_GROUP_POLICY = "Invalid.Parameter.Security.Group.Policy"
// INVALID_PARAMETER_SECURITY_GROUP_PORT_RANGE = "Invalid.Parameter.Security.Group.Port.Range"
// INVALID_PARAMETER_SECURITY_GROUP_PRIORITY = "Invalid.Parameter.Security.Group.Priority"
// INVALID_PARAMETER_SECURITY_GROUP_SOURCE_CIDR_IP = "Invalid.Parameter.Security.Group.Source.Cidr.Ip"
// OPERATION_DENIED_SECURITY_GROUP_EXIST_REPEAT_RULE = "Operation.Denied.Security.Group.Exist.Repeat.Rule"
// OPERATION_DENIED_SECURITY_GROUP_RULE_EXCEED_LIMIT = "Operation.Denied.Security.Group.Rule.Exceed.Limit"
// OPERATION_DENIED_SECURITY_GROUP_TEAM_EXCEED_LIMIT = "Operation.Denied.Security.Group.Team.Exceed.Limit"
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

// DeleteSecurityGroup
// This API is used to delete a security group.
//
// Possible error codes to return:
// INVALID_SECURITY_GROUP_NOT_FOUND = "Invalid.Security.Group.Not.Found"
// OPERATION_DENIED_SECURITY_GROUP_STATUS_NOT_SUPPORT = "Operation.Denied.Security.Group.Status.Not.Support"
// OPERATION_DENIED_SECURITY_GROUP_EXIST_INSTANCE = "Operation.Denied.Security.Group.Exist.Instance"
// OPERATION_DENIED_SECURITY_GROUP_EXIST_PRE_PRODUCT_INSTANCE = "Operation.Denied.Security.Group.Exist.Pre.Product.Instance"
// OPERATION_DENIED_DEFAULT_SECURITY_GROUP_NOT_SUPPORT = "Operation.Denied.Default.Security.Group.Not.Support"
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
	response = NewDeleteSecurityGroupResponse()
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

func (c *Client) AuthorizeSecurityGroupRules(request *AuthorizeSecurityGroupRulesRequest) (response *AuthorizeSecurityGroupRulesResponse, err error) {
	response = NewAuthorizeSecurityGroupRulesResponse()
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

func (c *Client) AuthorizeSecurityGroupRule(request *AuthorizeSecurityGroupRuleRequest) (response *AuthorizeSecurityGroupRuleResponse, err error) {
	response = NewAuthorizeSecurityGroupRuleResponse()
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

func (c *Client) UnAssociateSecurityGroupInstance(request *UnAssociateSecurityGroupInstanceRequest) (response *UnAssociateSecurityGroupInstanceResponse, err error) {
	response = NewUnAssociateSecurityGroupInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

//////////////////////////// VM //////////////////////////////

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

func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
	response = NewInquiryPriceCreateInstanceResponse()
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

//////////////////////////// Disk //////////////////////////////

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

//////////////////////////// Subnet //////////////////////////////

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
