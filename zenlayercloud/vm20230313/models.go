package vm

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type CreateImageRequest struct {
	*common.BaseRequest

	// ID of instance to be mirrored.
	InstanceId string `json:"instanceId,omitempty"`

	// Image name to be displayed.
	//This parameter must contain up to 24 characters. Only Chinese characters, letters, numbers, - and _ are supported.
	ImageName string `json:"imageName,omitempty"`

	// Image description.
	// This parameter must contain up to 255 characters.
	ImageDescription string `json:"imageDescription,omitempty"`
}

type CreateImageResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateImageResponseParams `json:"response"`
}

type CreateImageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Image ID.
	ImageId string `json:"imageId,omitempty"`
}

type DescribeZonesRequest struct {
	*common.BaseRequest

	// List of zone IDs.
	ZoneIds []string `json:"zoneIds,omitempty"`
}

type DescribeZonesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeZonesResponseParam `json:"response"`
}

type DescribeZonesResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// The list of zones available.
	ZoneSet []*ZoneInfo `json:"zoneSet,omitempty"`
}

type ZoneInfo struct {
	// Zone ID. For example, SEL-A.
	ZoneId string `json:"zoneId,omitempty"`

	// Zone name.
	ZoneName string `json:"zoneName,omitempty"`

	// Zone support securityGroup.
	SupportSecurityGroup bool `json:"supportSecurityGroup,omitempty"`

	// Zone support networkType
	SupportNetworkType string `json:"supportNetworkType,omitempty"`

	SupportIpv6 bool `json:"supportIpv6,omitempty"`

	SupportCpuPassThrough bool `json:"supportCpuPassThrough,omitempty"`
}

type DescribeImageRequest struct {
	*common.BaseRequest
	ImageId string `json:"imageId,omitempty"`
}

type DescribeImagesRequest struct {
	*common.BaseRequest

	// ID list of images.
	// Call DescribeImages and find imageId in the response.
	ImageIds []string `json:"imageIds,omitempty"`

	// Image name.
	ImageName string `json:"imageName,omitempty"`

	// Zone ID.
	// Call DescribeZones and find zoneId in the response.
	ZoneId string `json:"zoneId,omitempty"`

	// Image category.
	// Available values:
	// CentOS
	// Windows
	// Ubuntu
	// Debian
	Category string `json:"category,omitempty"`

	// Image type.
	// PUBLIC_IMAGE: the default images.
	// CUSTOM_IMAGE: the newly created images by yourself.
	ImageType string `json:"imageType,omitempty"`

	// Operating system type.
	// Available values:
	// Windows
	// Linux
	OsType string `json:"osType,omitempty"`

	// Image status.
	// CREATING: creating.
	// AVAILABLE: able to use.
	// UNAVAILABLE: unable to use.
	ImageStatus string `json:"imageStatus,omitempty"`

	// Number of pages returned.
	// Default value: 1.
	PageNum int `json:"pageNum,omitempty"`

	// Number of items in the current page result.
	// Default value: 20;
	// Maximum value: 1000.
	PageSize int `json:"pageSize,omitempty"`
}

type DescribeImageResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeImageResponseParams `json:"response"`
}

type DescribeImagesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeImagesResponseParams `json:"response"`
}

type DescribeImageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Image ID.
	ImageId string `json:"imageId,omitempty"`

	// Image name.
	ImageName string `json:"imageName,omitempty"`

	// Image type.
	// PUBLIC_IMAGE: the default images.
	// CUSTOM_IMAGE: the newly created images by yourself.
	ImageType string `json:"imageType,omitempty"`

	// Image size.
	// Unit: GB.
	ImageSize string `json:"imageSize,omitempty"`

	// Image description.
	ImageDescription string `json:"imageDescription,omitempty"`

	// Image version.
	ImageVersion string `json:"imageVersion,omitempty"`

	// Image status.
	// CREATING: creating.
	// AVAILABLE: able to use.
	// UNAVAILABLE: unable to use.
	ImageStatus string `json:"imageStatus,omitempty"`

	// Image category.
	// Available values:
	// CentOS
	// Windows
	// Ubuntu
	// Debian
	Category string `json:"category,omitempty"`

	// Operating system type.
	// Available values:
	// Windows
	// Linux
	OsType string `json:"osType,omitempty"`
}

type DescribeImagesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Information on an image.
	DataSet []*ImageInfo `json:"dataSet,omitempty"`

	// Number of images meeting the filtering conditions.
	TotalCount int `json:"totalCount,omitempty"`
}

type ImageInfo struct {
	// Image ID.
	ImageId string `json:"imageId,omitempty"`

	// Image name.
	ImageName string `json:"imageName,omitempty"`

	// Image type.
	// PUBLIC_IMAGE: the default images.
	// CUSTOM_IMAGE: the newly created images by yourself.
	ImageType string `json:"imageType,omitempty"`

	// Image size.
	// Unit: GB.
	ImageSize string `json:"imageSize,omitempty"`

	// Image description.
	ImageDescription string `json:"imageDescription,omitempty"`

	// Image version.
	ImageVersion string `json:"imageVersion,omitempty"`

	// Image status.
	// CREATING: creating.
	// AVAILABLE: able to use.
	// UNAVAILABLE: unable to use.
	ImageStatus string `json:"imageStatus,omitempty"`

	// Image category.
	// Available values:
	// CentOS
	// Windows
	// Ubuntu
	// Debian
	Category string `json:"category,omitempty"`

	// Operating system type.
	// Available values:
	// Windows
	// Linux
	OsType string `json:"osType,omitempty"`
}

type DescribeImageQuotaRequest struct {
	*common.BaseRequest

	// Zone ID.
	// Call DescribeZones and find zoneId in the response.
	ZoneIds []*string `json:"zoneIds,omitempty"`
}

type DescribeImageQuotaResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeImageQuotaResponseParam `json:"response"`
}

type DescribeImageQuotaResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Information on image quota.
	Images []*ImageQuotaInfo `json:"images,omitempty"`
}

type ImageQuotaInfo struct {
	// Zone supported for images.
	ZoneId string `json:"zoneId,omitempty"`

	// Quantity of current images.
	Count int `json:"count,omitempty"`

	// Maximum quantity of images in current zone.
	MaxCount int `json:"maxCount,omitempty"`
}

type ModifyImagesAttributesRequest struct {
	*common.BaseRequest

	// ID list of images.
	// Call DescribeImages and find imageId in the response.
	ImageIds []string `json:"imageIds,omitempty"`

	// Image name to be displayed.
	// This parameter must contain up to 24 characters. Only Chinese characters, letters, numbers, - and _ are supported.
	ImageName string `json:"imageName,omitempty"`

	// Image description.
	// This parameter must contain up to 255 characters.
	ImageDescription string `json:"imageDescription,omitempty"`
}

type ModifyImagesAttributesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DeleteImagesRequest struct {
	*common.BaseRequest

	// ID list of images to be deleted.
	ImageIds []string `json:"imageIds,omitempty"`
}

type DeleteImagesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeSecurityGroupsRequest struct {
	*common.BaseRequest

	// Security group ID(s).
	// You can query up to 100 security groups in each request.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

	// Security group name.
	SecurityGroupName string `json:"securityGroupName,omitempty"`

	// Number of pages returned.
	// Default value: 1.
	PageNum int `json:"pageNum,omitempty"`

	// Number of items in the current page result.
	// Default value: 20;
	// Maximum value: 1000.
	PageSize int `json:"pageSize,omitempty"`
}

type DescribeSecurityGroupsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeSecurityGroupsResponseParams `json:"response,omitempty"`
}

type DescribeSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Information on a security group.
	DataSet []*SecurityGroupInfo `json:"dataSet,omitempty"`

	// Number of security groups meeting the filtering conditions.
	TotalCount int `json:"totalCount,omitempty"`
}

type SecurityGroupInfo struct {
	// Security group ID.
	SecurityGroupId string `json:"securityGroupId,omitempty"`

	// Security group name.
	SecurityGroupName string `json:"securityGroupName,omitempty"`

	// Status of security group.
	SecurityGroupStatus string `json:"securityGroupStatus,omitempty"`

	// Creation time.
	// Format: YYYY-MM-DDThh:mm:ssZ.
	CreateTime string `json:"createTime,omitempty"`

	// Security group description.
	Description string `json:"description,omitempty"`

	// ID list of applied instances.
	InstanceIds []*string `json:"instanceIds,omitempty"`

	// Security group rules.
	RuleInfos []*RuleInfo `json:"ruleInfos,omitempty"`

	// Whether the security  group is default or not.
	IsDefault bool `json:"isDefault,omitempty"`
}

type RuleInfo struct {
	// Traffic directions.
	// ingress: inbound rules.
	// egress: outbound rules.
	Direction string `json:"direction,omitempty"`

	// Access strategy.
	// Available values:
	// accept: allows access by default.
	// Only accept strategy is supported for now.
	Policy string `json:"policy,omitempty"`

	// Transport protocol. The value is case sensitive.
	// Available value:
	// tcp: TCP protocol.
	// udp: UDP protocol.
	// icmp: ICMP protocol.
	// all: all protocols supported.
	IpProtocol string `json:"ipProtocol,omitempty"`

	// Destination port range.
	// Available values:
	// For TCP and UDP protocols: The value ranges from 1 to 65535. Use a slash (/) to separate the start port number and the end port number. Examples: 1/200; incorrect example: 200/1.
	// For ICMP protocol: -1/-1.
	// For all protocols: -1/-1.
	PortRange string `json:"portRange,omitempty"`

	// Source IP address range.
	// Default value: 0.0.XX.XX/0.
	CidrIp string `json:"cidrIp,omitempty"`

	Description string `json:"description,omitempty"`
}

type ModifySecurityGroupsAttributeRequest struct {
	*common.BaseRequest

	// Security group ID(s).
	// To obtain the security group IDs, you can call DescribeSecurityGroups and look for securityGroupId in the response.
	// The maximum number of security groups in each request is 100.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

	// Security group name to be displayed.
	// This parameter can contain up to 64 characters. Only letters, numbers, - and periods (.) are supported.
	SecurityGroupName string `json:"securityGroupName,omitempty"`

	// Security group description.
	// This parameter must contain 2 to 255 characters.
	Description *string `json:"description,omitempty"`
}

type ModifySecurityGroupsAttributeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeInstanceAvailableSecurityGroupResourceRequest struct {
	*common.BaseRequest

	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`
}

type DescribeInstanceAvailableSecurityGroupResourceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstanceAvailableSecurityGroupResourceParams `json:"response,omitempty"`
}

type DescribeInstanceAvailableSecurityGroupResourceParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// All available security groups.
	InstanceAvailableSecurityGroups []*InstanceAvailableSecurityGroup `json:"instanceAvailableSecurityGroups,omitempty"`
}

type InstanceAvailableSecurityGroup struct {
	// Security group ID.
	SecurityGroupId string `json:"securityGroupId,omitempty"`

	// Security group name.
	SecurityGroupName string `json:"securityGroupName,omitempty"`

	// Whether the security  group is default or not.
	IsDefault bool `json:"isDefault,omitempty"`
}

type CreateSecurityGroupRequest struct {
	*common.BaseRequest

	// Security group name to be displayed.
	// This parameter can contain up to 64 characters. Only letters, numbers, - and periods (.) are supported.
	SecurityGroupName string `json:"securityGroupName,omitempty"`

	// Security group rules.
	RuleInfos []*RuleInfo `json:"ruleInfos,omitempty"`

	// Security group description.
	// This parameter must contain 2 to 255 characters.
	Description string `json:"description,omitempty"`
}

type CreateSecurityGroupResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateSecurityGroupParams `json:"response,omitempty"`
}

type CreateSecurityGroupParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Security group ID.
	SecurityGroupId string `json:"securityGroupId,omitempty"`
}

type DeleteSecurityGroupRequest struct {
	*common.BaseRequest

	// Security group ID.
	SecurityGroupId string `json:"securityGroupId,omitempty"`
}

type DeleteSecurityGroupResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type AuthorizeSecurityGroupRulesRequest struct {
	*common.BaseRequest
	SecurityGroupId string      `json:"securityGroupId,omitempty"`
	RuleInfos       []*RuleInfo `json:"ruleInfos,omitempty"`
}

type AuthorizeSecurityGroupRulesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type ConfigureSecurityGroupRulesRequest struct {
	*common.BaseRequest
	SecurityGroupId string      `json:"securityGroupId,omitempty"`
	RuleInfos       []*RuleInfo `json:"ruleInfos,omitempty"`
}

type ConfigureSecurityGroupRulesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type AuthorizeSecurityGroupRuleRequest struct {
	*common.BaseRequest
	SecurityGroupId string `json:"securityGroupId,omitempty"`
	Direction       string `json:"direction,omitempty"`
	Policy          string `json:"policy,omitempty"`
	IpProtocol      string `json:"ipProtocol,omitempty"`
	PortRange       string `json:"portRange,omitempty"`
	CidrIp          string `json:"cidrIp,omitempty"`
	Description     string `json:"description,omitempty"`
}

type AuthorizeSecurityGroupRuleResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type RevokeSecurityGroupRulesRequest struct {
	*common.BaseRequest
	SecurityGroupId string      `json:"securityGroupId,omitempty"`
	RuleInfos       []*RuleInfo `json:"ruleInfos,omitempty"`
}

type RevokeSecurityGroupRulesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type AssociateSecurityGroupInstanceRequest struct {
	*common.BaseRequest
	SecurityGroupId string `json:"securityGroupId,omitempty"`
	InstanceId      string `json:"instanceId,omitempty"`
}

type AssociateSecurityGroupInstanceResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type UnAssociateSecurityGroupInstanceRequest struct {
	*common.BaseRequest
	SecurityGroupId string `json:"securityGroupId,omitempty"`
	InstanceId      string `json:"instanceId,omitempty"`
}

type UnAssociateSecurityGroupInstanceResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type ModifyInstancesResourceGroupRequest struct {
	*common.BaseRequest
	InstanceIds     []string `json:"instanceIds,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
}

type ModifyInstancesResourceGroupResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type StartInstancesRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type StartInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type StopInstancesRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
	ForceShutdown *bool   `json:"forceShutdown,omitempty"`

}

type StopInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RebootInstancesRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type RebootInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ResetInstancesPasswordRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
	Password    string   `json:"password,omitempty"`
}

type ResetInstancesPasswordResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ResetInstanceRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
	ImageId    string `json:"imageId,omitempty"`
	Password   string `json:"password,omitempty"`
	KeyId      string `json:"keyId,omitempty"`
	WanName    string `json:"wanName,omitempty"`
	LanName    string `json:"lanName,omitempty"`
}

type ResetInstanceResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type TerminateInstanceRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
}

type TerminateInstanceResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseInstancesRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type ReleaseInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyInstancesAttributeRequest struct {
	*common.BaseRequest
	InstanceIds  []string `json:"instanceIds,omitempty"`
	InstanceName string   `json:"instanceName,omitempty"`
}

type ModifyInstancesAttributeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type InquiryPriceInstanceBandwidthRequest struct {
	*common.BaseRequest
	InstanceId              string `json:"instanceId,omitempty"`
	InternetMaxBandwidthOut int    `json:"internetMaxBandwidthOut,omitempty"`
}

type InquiryPriceInstanceBandwidthResponse struct {
	*common.BaseResponse
	RequestId string                                       `json:"requestId,omitempty"`
	Response  *InquiryPriceInstanceBandwidthResponseParams `json:"response"`
}

type InquiryPriceInstanceBandwidthResponseParams struct {
	RequestId      string   `json:"requestId,omitempty"`
	BandwidthPrice []*Price `json:"bandwidthPrice,omitempty"`
}

type ModifyInstanceBandwidthRequest struct {
	*common.BaseRequest
	InstanceId              string `json:"instanceId,omitempty"`
	InternetMaxBandwidthOut int    `json:"internetMaxBandwidthOut,omitempty"`
}

type ModifyInstanceBandwidthResponse struct {
	*common.BaseResponse
	RequestId string                                 `json:"requestId,omitempty"`
	Response  *ModifyInstanceBandwidthResponseParams `json:"response"`
}

type ModifyInstanceBandwidthResponseParams struct {
	RequestId   string `json:"requestId,omitempty"`
	OrderNumber string `json:"orderNumber,omitempty"`
}

type CancelInstanceBandwidthDowngradeRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
}

type CancelInstanceBandwidthDowngradeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type InquiryPriceInstanceTrafficPackageRequest struct {
	*common.BaseRequest
	InstanceId         string   `json:"instanceId,omitempty"`
	TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`
}

type InquiryPriceInstanceTrafficPackageResponse struct {
	*common.BaseResponse
	RequestId string                                            `json:"requestId,omitempty"`
	Response  *InquiryPriceInstanceTrafficPackageResponseParams `json:"response"`
}

type InquiryPriceInstanceTrafficPackageResponseParams struct {
	RequestId           string   `json:"requestId,omitempty"`
	TrafficPackagePrice []*Price `json:"trafficPackagePrice,omitempty"`
}

type ModifyInstanceTrafficPackageRequest struct {
	*common.BaseRequest
	InstanceId         string   `json:"instanceId,omitempty"`
	TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`
}

type ModifyInstanceTrafficPackageResponse struct {
	*common.BaseResponse
	RequestId string                                      `json:"requestId,omitempty"`
	Response  *ModifyInstanceTrafficPackageResponseParams `json:"response"`
}

type ModifyInstanceTrafficPackageResponseParams struct {
	RequestId   string `json:"requestId,omitempty"`
	OrderNumber string `json:"orderNumber,omitempty"`
}

type CancelInstanceTrafficPackageDowngradeRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
}

type CancelInstanceTrafficPackageDowngradeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeInstancesRequest struct {
	*common.BaseRequest
	InstanceIds        []string `json:"instanceIds,omitempty"`
	ZoneId             string   `json:"zoneId,omitempty"`
	ResourceGroupId    string   `json:"resourceGroupId,omitempty"`
	InstanceType       string   `json:"instanceType,omitempty"`
	InternetChargeType string   `json:"internetChargeType,omitempty"`
	ImageId            string   `json:"imageId,omitempty"`
	KeyId              string   `json:"keyId,omitempty"`
	SubnetId           string   `json:"subnetId,omitempty"`
	InstanceStatus     string   `json:"instanceStatus,omitempty"`
	InstanceName       string   `json:"instanceName,omitempty"`
	SecurityGroupId    string   `json:"securityGroupId,omitempty"`
	PublicIpAddresses  []string `json:"publicIpAddresses,omitempty"`
	PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`
	PageSize           int      `json:"pageSize,omitempty"`
	PageNum            int      `json:"pageNum,omitempty"`
}

type DescribeInstancesResponse struct {
	*common.BaseResponse
	RequestId string                           `json:"requestId,omitempty"`
	Response  *DescribeInstancesResponseParams `json:"response,omitempty"`
}

type DescribeInstancesResponseParams struct {
	RequestId  string          `json:"requestId,omitempty"`
	DataSet    []*InstanceInfo `json:"dataSet,omitempty"`
	TotalCount int             `json:"totalCount,omitempty"`
}

type InstanceInfo struct {
	InstanceId              string      `json:"instanceId,omitempty"`
	ZoneId                  string      `json:"zoneId,omitempty"`
	InstanceName            string      `json:"instanceName,omitempty"`
	InstanceType            string      `json:"instanceType,omitempty"`
	CpuCount                int         `json:"cpuCount,omitempty"`
	Memory                  int         `json:"memory,omitempty"`
	ImageId                 string      `json:"imageId,omitempty"`
	ImageName               string      `json:"imageName,omitempty"`
	InstanceChargeType      string      `json:"instanceChargeType,omitempty"`
	InternetMaxBandwidthOut int         `json:"internetMaxBandwidthOut,omitempty"`
	InternetChargeType      string      `json:"internetChargeType,omitempty"`
	Period                  *int        `json:"period,omitempty"`
	PublicIpAddresses       []string    `json:"publicIpAddresses,omitempty"`
	PublicIpv6Addresses     []string    `json:"publicIpv6Addresses,omitempty"`
	PrivateIpAddresses      []string    `json:"privateIpAddresses,omitempty"`
	SubnetId                string      `json:"subnetId,omitempty"`
	CreateTime              string      `json:"createTime,omitempty"`
	ExpiredTime             *string     `json:"expiredTime,omitempty"`
	ResourceGroupId         string      `json:"resourceGroupId,omitempty"`
	ResourceGroupName       string      `json:"resourceGroupName,omitempty"`
	InstanceStatus          string      `json:"instanceStatus,omitempty"`
	TrafficPackageSize      *float64    `json:"trafficPackageSize,omitempty"`
	SecurityGroupIds        []string    `json:"securityGroupIds,omitempty"`
	SystemDisk              *SystemDisk `json:"systemDisk,omitempty"`
	DataDisks               []*DataDisk `json:"dataDisks,omitempty"`
	AutoRenew               bool        `json:"autoRenew,omitempty"`
	KeyId                   string      `json:"keyId,omitempty"`
	Nic                     *Nic        `json:"nic,omitempty"`
}

type SystemDisk struct {
	DiskId       string `json:"diskId,omitempty"`
	DiskSize     int    `json:"diskSize,omitempty"`
	DiskCategory string `json:"diskCategory,omitempty"`
}

type DataDisk struct {
	DiskId       string `json:"diskId,omitempty"`
	DiskSize     int    `json:"diskSize,omitempty"`
	DiskName     string `json:"diskName,omitempty"`
	Portable     bool   `json:"portable,omitempty"`
	DiskCategory string `json:"diskCategory,omitempty"`
}

type DiskCategory struct {
	ZoneId      string   `json:"zoneId,omitempty"`
	CategorySet []string `json:"categorySet,omitempty"`
}

type DescribeInstancesStatusRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
	PageSize    int      `json:"pageSize,omitempty"`
	PageNum     int      `json:"pageNum,omitempty"`
}

type DescribeInstancesStatusResponse struct {
	*common.BaseResponse
	RequestId string                                 `json:"requestId,omitempty"`
	Response  *DescribeInstancesStatusResponseParams `json:"response,omitempty"`
}

type DescribeInstancesStatusResponseParams struct {
	RequestId  string            `json:"requestId,omitempty"`
	DataSet    []*InstanceStatus `json:"dataSet,omitempty"`
	TotalCount int               `json:"totalCount,omitempty"`
}

type InstanceStatus struct {
	InstanceId     string `json:"instanceId,omitempty"`
	InstanceStatus string `json:"instanceStatus,omitempty"`
}

type Price struct {
	Discount          *float64     `json:"discount,omitempty"`
	DiscountPrice     *float64     `json:"discountPrice,omitempty"`
	OriginalPrice     *float64     `json:"originalPrice,omitempty"`
	UnitPrice         *float64     `json:"unitPrice,omitempty"`
	DiscountUnitPrice *float64     `json:"discountUnitPrice,omitempty"`
	ChargeUnit        *string      `json:"chargeUnit,omitempty"`
	StepPrices        []*StepPrice `json:"stepPrices,omitempty"`
}

type StepPrice struct {
	StepStart         *float64 `json:"stepStart,omitempty"`
	StepEnd           *float64 `json:"stepEnd,omitempty"`
	UnitPrice         *float64 `json:"unitPrice,omitempty"`
	DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`
}

type CreateInstancesRequest struct {
	*common.BaseRequest
	ZoneId                  string         `json:"zoneId,omitempty"`
	InstanceChargeType      string         `json:"instanceChargeType,omitempty"`
	InstanceChargePrepaid   *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`
	InstanceType            string         `json:"instanceType,omitempty"`
	ImageId                 string         `json:"imageId,omitempty"`
	ResourceGroupId         string         `json:"resourceGroupId,omitempty"`
	InstanceName            string         `json:"instanceName,omitempty"`
	InstanceCount           int            `json:"instanceCount,omitempty"`
	Password                string         `json:"password,omitempty"`
	KeyId                   string         `json:"keyId,omitempty"`
	InternetChargeType      string         `json:"internetChargeType,omitempty"`
	InternetMaxBandwidthOut int            `json:"internetMaxBandwidthOut,omitempty"`
	TrafficPackageSize      *float64       `json:"trafficPackageSize,omitempty"`
	SubnetId                string         `json:"subnetId,omitempty"`
	EnableIpv6              bool           `json:"enableIpv6,omitempty"`
	CpuPassThrough          bool           `json:"cpuPassThrough,omitempty"`
	InitScript              string         `json:"initScript,omitempty"`
	SystemDisk              *SystemDisk    `json:"systemDisk,omitempty"`
	DataDisks               []*DataDisk    `json:"dataDisks,omitempty"`
	SecurityGroupId         string         `json:"securityGroupId,omitempty"`
	Nic                     *Nic           `json:"nic,omitempty"`
	ClusterId               string         `json:"clusterId,omitempty"`
	NetworkMode             string         `json:"networkMode,omitempty"`
	DiskPreAllocated        bool           `json:"diskPreAllocated,omitempty"`
}

type Nic struct {
	// Public NIC name.
	// Only numbers, uppercase and lowercase letters are allowed, starting with a letter with a length limit of 4 to 15 characters.
	WanName string `json:"wanName,omitempty"`

	// Private NIC name.
	// Only numbers, uppercase and lowercase letters are allowed, starting with a letter with a length limit of 4 to 15 characters.
	LanName string `json:"lanName,omitempty"`
}

type CreateInstancesResponse struct {
	*common.BaseResponse
	RequestId string                         `json:"requestId,omitempty"`
	Response  *CreateInstancesResponseParams `json:"response"`
}

type CreateInstancesResponseParams struct {
	RequestId     *string   `json:"requestId,omitempty"`
	InstanceIdSet []*string `json:"instanceIdSet,omitempty"`
	OrderNumber   *string   `json:"orderNumber,omitempty"`
}

type ChargePrepaid struct {
	Period int `json:"period,omitempty"`
}

type DescribeZoneInstanceConfigInfosRequest struct {
	*common.BaseRequest
	InstanceChargeType string `json:"instanceChargeType,omitempty"`
	ZoneId             string `json:"zoneId,omitempty"`
	InstanceType       string `json:"instanceType,omitempty"`
}

type DescribeZoneInstanceConfigInfosResponse struct {
	*common.BaseResponse
	RequestId string                                         `json:"requestId,omitempty"`
	Response  *DescribeZoneInstanceConfigInfosResponseParams `json:"response"`
}

type DescribeZoneInstanceConfigInfosResponseParams struct {
	RequestId            string                   `json:"requestId,omitempty"`
	InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"instanceTypeQuotaSet,omitempty"`
}

type InstanceTypeQuotaItem struct {
	ZoneId                       string   `json:"zoneId,omitempty"`
	InstanceType                 string   `json:"instanceType,omitempty"`
	CpuCount                     int      `json:"cpuCount,omitempty"`
	Memory                       int      `json:"memory,omitempty"`
	Frequency                    string   `json:"frequency,omitempty"`
	InternetMaxBandwidthOutLimit int      `json:"internetMaxBandwidthOutLimit,omitempty"`
	InstanceTypeName             string   `json:"instanceTypeName,omitempty"`
	InternetChargeTypes          []string `json:"internetChargeTypes,omitempty"`
}

type InquiryPriceCreateInstanceRequest struct {
	*common.BaseRequest
	ZoneId                  string         `json:"zoneId,omitempty"`
	InstanceType            string         `json:"instanceType,omitempty"`
	ImageId                 string         `json:"imageId,omitempty"`
	InstanceChargeType      string         `json:"instanceChargeType,omitempty"`
	InternetChargeType      string         `json:"internetChargeType,omitempty"`
	InstanceChargePrepaid   *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`
	TrafficPackageSize      *float64       `json:"trafficPackageSize,omitempty"`
	InternetMaxBandwidthOut int            `json:"internetMaxBandwidthOut,omitempty"`
	SystemDisk              *SystemDisk    `json:"systemDisk,omitempty"`
	DataDisks               []*DataDisk    `json:"dataDisks,omitempty"`
}

type InquiryPriceCreateInstanceResponse struct {
	*common.BaseResponse
	RequestId string                                    `json:"requestId,omitempty"`
	Response  *InquiryPriceCreateInstanceResponseParams `json:"response"`
}

type InquiryPriceCreateInstanceResponseParams struct {
	RequestId       string   `json:"requestId,omitempty"`
	InstancePrice   Price    `json:"instancePrice,omitempty"`
	BandwidthPrice  []*Price `json:"bandwidthPrice,omitempty"`
	SystemDiskPrice *Price   `json:"systemDiskPrice,omitempty"`
	DataDiskPrice   *Price   `json:"dataDiskPrice,omitempty"`
}

type DescribeInstanceInternetStatusRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
}

type DescribeInstanceInternetStatusResponse struct {
	*common.BaseResponse
	RequestId string                                        `json:"requestId,omitempty"`
	Response  *DescribeInstanceInternetStatusResponseParams `json:"response"`
}

type DescribeInstanceInternetStatusResponseParams struct {
	RequestId                       string   `json:"requestId,omitempty"`
	InstanceId                      string   `json:"instanceId,omitempty"`
	InstanceName                    string   `json:"instanceName,omitempty"`
	InternetMaxBandwidthOut         *int     `json:"internetMaxBandwidthOut,omitempty"`
	ModifiedInternetMaxBandwidthOut *int     `json:"modifiedInternetMaxBandwidthOut,omitempty"`
	ModifiedBandwidthStatus         string   `json:"modifiedBandwidthStatus,omitempty"`
	TrafficPackageSize              *float64 `json:"trafficPackageSize,omitempty"`
	ModifiedTrafficPackageSize      *float64 `json:"modifiedTrafficPackageSize,omitempty"`
	ModifiedTrafficPackageStatus    string   `json:"modifiedTrafficPackageStatus,omitempty"`
}

type DescribeInstanceTrafficRequest struct {
	*common.BaseRequest

	InstanceId string `json:"instanceId,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`
}

type DescribeInstanceTrafficResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InstanceTrafficDataResponse `json:"response"`
}

type InstanceTrafficDataResponse struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	DataList []*InstanceTrafficData `json:"dataList,omitempty"`

	In95 int64 `json:"in95,omitempty"`

	In95Time string `json:"in95Time,omitempty"`

	InAvg int64 `json:"inAvg,omitempty"`

	InMax int64 `json:"inMax,omitempty"`

	InMin int64 `json:"inMin,omitempty"`

	InTotal int64 `json:"inTotal,omitempty"`

	MaxBandwidth95ValueMbps float64 `json:"maxBandwidth95ValueMbps,omitempty"`

	Out95 int64 `json:"out95,omitempty"`

	Out95Time string `json:"out95Time,omitempty"`

	OutAvg int64 `json:"outAvg,omitempty"`

	OutMax int64 `json:"outMax,omitempty"`

	OutMin int64 `json:"outMin,omitempty"`

	OutTotal int64 `json:"outTotal,omitempty"`

	TotalUnit string `json:"totalUnit,omitempty"`

	Unit string `json:"unit,omitempty"`
}

type InstanceTrafficData struct {
	InternetRX int64 `json:"internetRX,omitempty"`

	InternetTX int64 `json:"internetTX,omitempty"`

	Time string `json:"time,omitempty"`
}

type DescribeInstanceCpuMonitorRequest struct {
	*common.BaseRequest

	InstanceId string `json:"instanceId,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`
}

type DescribeInstanceCpuMonitorResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InstanceCpuMonitorResponse `json:"response"`
}

type InstanceCpuMonitorResponse struct {

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	DataList []*InstanceCpuMonitorData `json:"dataList,omitempty"`
}

type InstanceCpuMonitorData struct {
	Cpu string `json:"cpu,omitempty"`

	Time string `json:"time,omitempty"`
}

type ModifyInstanceTypeRequest struct {
	*common.BaseRequest

	InstanceId string `json:"instanceId,omitempty"`

	InstanceTypeId string `json:"instanceTypeId,omitempty"`
}

type ModifyInstanceTypeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *ModifyInstanceTypeResponseParams `json:"response"`
}

type ModifyInstanceTypeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	OrderNumber string `json:"orderNumber,omitempty"`
}

type CancelInstanceDowngradeRequest struct {
	*common.BaseRequest

	InstanceId string `json:"instanceId,omitempty"`
}

type CancelInstanceDowngradeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeInstanceTypeStatusRequest struct {
	*common.BaseRequest

	InstanceId string `json:"instanceId,omitempty"`
}

type DescribeInstanceTypeStatusResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstanceTypeStatusResponseParams `json:"response"`
}

type DescribeInstanceTypeStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	InstanceId string `json:"instanceId,omitempty"`

	InstanceName string `json:"instanceName,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`

	ModifiedInstanceType string `json:"modifiedInstanceType,omitempty"`

	ModifiedInstanceTypeStatus string `json:"modifiedInstanceTypeStatus,omitempty"`
}

type CreateDisksRequest struct {
	*common.BaseRequest
	ChargeType      string         `json:"chargeType,omitempty"`
	ChargePrepaid   *ChargePrepaid `json:"chargePrepaid,omitempty"`
	DiskName        string         `json:"diskName,omitempty"`
	DiskSize        *int           `json:"diskSize,omitempty"`
	DiskCategory    string         `json:"diskCategory,omitempty"`
	InstanceId      string         `json:"instanceId,omitempty"`
	ZoneId          string         `json:"zoneId,omitempty"`
	DiskAmount      *int           `json:"diskAmount,omitempty"`
	ResourceGroupId string         `json:"resourceGroupId,omitempty"`
}

type CreateDisksResponse struct {
	*common.BaseResponse
	RequestId string                    `json:"requestId,omitempty"`
	Response  *CreateDiskResponseParams `json:"response"`
}

type CreateDiskResponseParams struct {
	RequestId   string   `json:"requestId,omitempty"`
	DiskIds     []string `json:"diskIds,omitempty"`
	OrderNumber string   `json:"orderNumber,omitempty"`
}

type DescribeDisksRequest struct {
	*common.BaseRequest
	DiskIds      []string `json:"diskIds,omitempty"`
	DiskName     string   `json:"diskName,omitempty"`
	DiskStatus   string   `json:"diskStatus,omitempty"`
	DiskType     string   `json:"diskType,omitempty"`
	DiskSize     *int     `json:"diskSize,omitempty"`
	DiskCategory string   `json:"diskCategory,omitempty"`
	Portable     *bool    `json:"portable,omitempty"`
	InstanceId   string   `json:"instanceId,omitempty"`
	ZoneId       string   `json:"zoneId,omitempty"`
	PageSize     int      `json:"pageSize,omitempty"`
	PageNum      int      `json:"pageNum,omitempty"`
}

type DiskStatus struct {
	DiskStatus string `json:"diskStatus,omitempty"`
}

type DescribeDisksResponse struct {
	*common.BaseResponse
	RequestId string                       `json:"requestId,omitempty"`
	Response  *DescribeDisksResponseParams `json:"response"`
}

type DescribeDisksResponseParams struct {
	RequestId  string      `json:"requestId,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	DataSet    []*DiskInfo `json:"dataSet,omitempty"`
}

type DiskInfo struct {
	DiskId       string `json:"diskId,omitempty"`
	DiskName     string `json:"diskName,omitempty"`
	ZoneId       string `json:"zoneId,omitempty"`
	DiskType     string `json:"diskType,omitempty"`
	Portable     bool   `json:"portable,omitempty"`
	DiskCategory string `json:"diskCategory,omitempty"`
	DiskSize     int    `json:"diskSize,omitempty"`
	DiskStatus   string `json:"diskStatus,omitempty"`
	InstanceId   string `json:"instanceId,omitempty"`
	InstanceName string `json:"instanceName,omitempty"`
	ChargeType   string `json:"chargeType,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	ExpiredTime  string `json:"expiredTime,omitempty"`
	Period       *int   `json:"period,omitempty"`
	AutoRenew    bool   `json:"autoRenew,omitempty"`
}

type AttachDisksRequest struct {
	*common.BaseRequest
	DiskIds    []string `json:"diskIds,omitempty"`
	InstanceId string   `json:"instanceId,omitempty"`
}

type AttachDisksResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ChangeDisksAttachRequest struct {
	*common.BaseRequest
	DiskIds    []string `json:"diskIds,omitempty"`
	InstanceId string   `json:"instanceId,omitempty"`
}

type ChangeDisksAttachResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DetachDisksRequest struct {
	*common.BaseRequest
	DiskIds []string `json:"diskIds,omitempty"`
}

type DetachDisksResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyDisksAttributesRequest struct {
	*common.BaseRequest
	DiskIds  []string `json:"diskIds,omitempty"`
	DiskName string   `json:"diskName,omitempty"`
}

type ModifyDisksAttributesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyDisksResourceGroupRequest struct {
	*common.BaseRequest
	DiskIds         []string `json:"diskIds,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
}

type ModifyDisksResourceGroupResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type InquiryPriceCreateDisksRequest struct {
	*common.BaseRequest
	ZoneId        string         `json:"zoneId,omitempty"`
	DiskSize      int            `json:"diskSize,omitempty"`
	DiskAmount    *int           `json:"diskAmount,omitempty"`
	ChargeType    string         `json:"chargeType,omitempty"`
	ChargePrepaid *ChargePrepaid `json:"chargePrepaid,omitempty"`
	DiskCategory  string         `json:"diskCategory,omitempty"`
}

type InquiryPriceCreateDisksResponse struct {
	*common.BaseResponse
	RequestId string                                 `json:"requestId,omitempty"`
	Response  *InquiryPriceCreateDisksResponseParams `json:"response"`
}

type InquiryPriceCreateDisksResponseParams struct {
	RequestId     string `json:"requestId,omitempty"`
	DataDiskPrice *Price `json:"dataDiskPrice,omitempty"`
}

type TerminateDiskRequest struct {
	*common.BaseRequest
	DiskId string `json:"diskId,omitempty"`
}

type TerminateDiskResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseDiskRequest struct {
	*common.BaseRequest
	DiskId string `json:"diskId,omitempty"`
}

type ReleaseDiskResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RenewDiskRequest struct {
	*common.BaseRequest
	DiskId string `json:"diskId,omitempty"`
}

type RenewDiskResponse struct {
	*common.BaseResponse
	RequestId string                   `json:"requestId,omitempty"`
	Response  *RenewDiskResponseParams `json:"response"`
}

type RenewDiskResponseParams struct {
	RequestId   string `json:"requestId,omitempty"`
	OrderNumber string `json:"orderNumber,omitempty"`
}

type DescribeDiskCategoryRequest struct {
	*common.BaseRequest
	ZoneId             string `json:"zoneId,omitempty"`
	InstanceChargeType string `json:"instanceChargeType,omitempty"`
	DiskCategory       string `json:"diskCategory,omitempty"`
}

type DescribeDiskCategoryResponse struct {
	*common.BaseResponse
	RequestId string                              `json:"requestId,omitempty"`
	Response  *DescribeDiskCategoryResponseParams `json:"response"`
}

type DescribeDiskCategoryResponseParams struct {
	RequestId       string         `json:"requestId,omitempty"`
	CategoryZoneSet []DiskCategory `json:"categoryZoneSet,omitempty"`
}

type CreateVpcRequest struct {
	*common.BaseRequest
	ZoneId          string `json:"zoneId,omitempty"`
	VpcName         string `json:"vpcName,omitempty"`
	VpcCidrBlock    string `json:"vpcCidrBlock,omitempty"`
	SubnetName      string `json:"subnetName,omitempty"`
	SubnetCidrBlock string `json:"subnetCidrBlock,omitempty"`
}

type CreateVpcResponse struct {
	*common.BaseResponse
	RequestId string                   `json:"requestId,omitempty"`
	Response  *CreateVpcResponseParams `json:"response"`
}

type CreateVpcResponseParams struct {
	RequestId string `json:"requestId,omitempty"`
	VpcId     string `json:"vpcId,omitempty"`
}

type DeleteVpcRequest struct {
	*common.BaseRequest
	VpcId string `json:"vpcId,omitempty"`
}

type DeleteVpcResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyVpcsAttributeRequest struct {
	*common.BaseRequest
	VpcIds  []string `json:"vpcIds,omitempty"`
	VpcName string   `json:"vpcName,omitempty"`
}

type ModifyVpcsAttributeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeVpcsRequest struct {
	*common.BaseRequest
	VpcIds    []string `json:"vpcIds,omitempty"`
	CidrBlock string   `json:"cidrBlock,omitempty"`
	ZoneId    string   `json:"zoneId,omitempty"`
	VpcStatus string   `json:"vpcStatus,omitempty"`
	VpcName   string   `json:"vpcName,omitempty"`
	PageSize  int      `json:"pageSize,omitempty"`
	PageNum   int      `json:"pageNum,omitempty"`
}

type DescribeVpcsResponse struct {
	*common.BaseResponse
	RequestId string                      `json:"requestId,omitempty"`
	Response  *DescribeVpcsResponseParams `json:"response"`
}

type DescribeVpcsResponseParams struct {
	RequestId  string            `json:"requestId,omitempty"`
	DataSet    []*VpcNetworkInfo `json:"dataSet,omitempty"`
	TotalCount int               `json:"totalCount,omitempty"`
}

type VpcNetworkInfo struct {
	VpcId        string   `json:"vpcId,omitempty"`
	ZoneId       string   `json:"zoneId,omitempty"`
	VpcName      string   `json:"vpcName,omitempty"`
	VpcStatus    string   `json:"vpcStatus,omitempty"`
	CidrBlock    string   `json:"cidrBlock,omitempty"`
	SubnetIdList []string `json:"subnetIdList,omitempty"`
	CreateTime   string   `json:"createTime,omitempty"`
	IsDefault    bool     `json:"isDefault,omitempty"`
}

type CreateVpcSubnetRequest struct {
	*common.BaseRequest
	CidrBlock  string `json:"cidrBlock,omitempty"`
	SubnetName string `json:"subnetName,omitempty"`
	VpcId      string `json:"vpcId,omitempty"`
}

type CreateVpcSubnetResponse struct {
	*common.BaseResponse
	RequestId string                         `json:"requestId,omitempty"`
	Response  *CreateVpcSubnetResponseParams `json:"response"`
}

type CreateVpcSubnetResponseParams struct {
	RequestId string `json:"requestId,omitempty"`
	SubnetId  string `json:"subnetId,omitempty"`
}

type DeleteVpcSubnetRequest struct {
	*common.BaseRequest
	SubnetId string `json:"subnetId,omitempty"`
}

type DeleteVpcSubnetResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyVpcSubnetsAttributeRequest struct {
	*common.BaseRequest
	SubnetIds  []string `json:"subnetIds,omitempty"`
	SubnetName string   `json:"subnetName,omitempty"`
}

type ModifyVpcSubnetsAttributeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeVpcSubnetsRequest struct {
	*common.BaseRequest
	SubnetIds    []string `json:"subnetIds,omitempty"`
	CidrBlock    string   `json:"cidrBlock,omitempty"`
	VpcId        string   `json:"vpcId,omitempty"`
	SubnetStatus string   `json:"subnetStatus,omitempty"`
	SubnetName   string   `json:"subnetName,omitempty"`
	PageSize     int      `json:"pageSize,omitempty"`
	PageNum      int      `json:"pageNum,omitempty"`
}

type DescribeVpcSubnetsResponse struct {
	*common.BaseResponse
	RequestId string                            `json:"requestId,omitempty"`
	Response  *DescribeVpcSubnetsResponseParams `json:"response"`
}

type DescribeVpcSubnetsResponseParams struct {
	RequestId  string           `json:"requestId,omitempty"`
	DataSet    []*VpcSubnetInfo `json:"dataSet,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
}

type VpcSubnetInfo struct {
	SubnetId       string   `json:"subnetId,omitempty"`
	VpcId          string   `json:"vpcId,omitempty"`
	SubnetName     string   `json:"subnetName,omitempty"`
	SubnetStatus   string   `json:"subnetStatus,omitempty"`
	CidrBlock      string   `json:"cidrBlock,omitempty"`
	InstanceIdList []string `json:"instanceIdList,omitempty"`
	CreateTime     string   `json:"createTime,omitempty"`
	UsageIpCount   int      `json:"usageIpCount,omitempty"`
	TotalIpCount   int      `json:"totalIpCount,omitempty"`
	IsDefault      bool     `json:"isDefault,omitempty"`
}

type CreateSubnetRequest struct {
	*common.BaseRequest
	CidrBlock         string `json:"cidrBlock,omitempty"`
	SubnetName        string `json:"subnetName,omitempty"`
	ZoneId            string `json:"zoneId,omitempty"`
	SubnetDescription string `json:"subnetDescription,omitempty"`
}

type CreateSubnetResponse struct {
	*common.BaseResponse
	RequestId string                      `json:"requestId,omitempty"`
	Response  *CreateSubnetResponseParams `json:"response"`
}

type CreateSubnetResponseParams struct {
	RequestId string `json:"requestId,omitempty"`
	SubnetId  string `json:"subnetId,omitempty"`
}

type DeleteSubnetRequest struct {
	*common.BaseRequest
	SubnetId string `json:"subnetId,omitempty"`
}

type DeleteSubnetResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifySubnetsAttributeRequest struct {
	*common.BaseRequest
	SubnetIds  []string `json:"subnetIds,omitempty"`
	SubnetName string   `json:"subnetName,omitempty"`
}

type ModifySubnetsAttributeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeSubnetsRequest struct {
	*common.BaseRequest
	SubnetIds       []string `json:"subnetIds,omitempty"`
	CidrBlock       string   `json:"cidrBlock,omitempty"`
	ZoneId          string   `json:"zoneId,omitempty"`
	SubnetStatus    string   `json:"subnetStatus,omitempty"`
	SubnetName      string   `json:"subnetName,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
	NetworkId       string   `json:"networkId,omitempty"`
	PageNum         int      `json:"pageNum,omitempty"`
	PageSize        int      `json:"pageSize,omitempty"`
}

type DescribeSubnetsResponse struct {
	*common.BaseResponse
	RequestId string                         `json:"requestId,omitempty"`
	Response  *DescribeSubnetsResponseParams `json:"response"`
}

type DescribeSubnetsResponseParams struct {
	RequestId  string        `json:"requestId,omitempty"`
	DataSet    []*SubnetInfo `json:"dataSet,omitempty"`
	TotalCount int           `json:"totalCount,omitempty"`
}

type SubnetInfo struct {
	SubnetId          string   `json:"subnetId,omitempty"`
	ZoneId            string   `json:"zoneId,omitempty"`
	SubnetName        string   `json:"subnetName,omitempty"`
	SubnetDescription string   `json:"subnetDescription,omitempty"`
	SubnetStatus      string   `json:"subnetStatus,omitempty"`
	CidrBlock         string   `json:"cidrBlock,omitempty"`
	CidrBlockList     []string `json:"cidrBlockList,omitempty"`
	UsageIpCount      int      `json:"usageIpCount,omitempty"`
	TotalIpCount      int      `json:"totalIpCount,omitempty"`
	CreateTime        string   `json:"createTime,omitempty"`
	InstanceIdList    []string `json:"instanceIdList,omitempty"`
}

type DescribeKeyPairsRequest struct {
	*common.BaseRequest
	KeyIds   []string `json:"keyIds,omitempty"`
	KeyName  string   `json:"keyName,omitempty"`
	PageNum  int      `json:"pageNum,omitempty"`
	PageSize int      `json:"pageSize,omitempty"`
}

type DescribeKeyPairsResponse struct {
	*common.BaseResponse
	RequestId string                          `json:"requestId,omitempty"`
	Response  *DescribeKeyPairsResponseParams `json:"response"`
}

type DescribeKeyPairsResponseParams struct {
	RequestId  string     `json:"requestId,omitempty"`
	DataSet    []*KeyPair `json:"dataSet,omitempty"`
	TotalCount int        `json:"totalCount,omitempty"`
}

type KeyPair struct {
	// 密钥对ID。
	KeyId string `json:"keyId,omitempty"`

	// 密钥对名称。
	KeyName string `json:"keyName,omitempty"`

	// 密钥对的公钥内容。
	PublicKey string `json:"publicKey,omitempty"`

	// 密钥对描述信息。
	KeyDescription string `json:"keyDescription,omitempty"`

	// 创建时间。
	CreateTime string `json:"createTime,omitempty"`
}

type ModifyKeyPairAttributeRequest struct {
	*common.BaseRequest

	// 密钥对ID。
	KeyId string `json:"keyId,omitempty"`

	// 密钥对描述信息。
	KeyDescription *string `json:"keyDescription,omitempty"`
}

type ModifyKeyPairAttributeResponse struct {
	*common.BaseResponse

	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DeleteKeyPairsRequest struct {
	*common.BaseRequest

	// 一个或多个待操作的密钥对ID。
	KeyIds []string `json:"keyIds,omitempty"`
}

type DeleteKeyPairsResponse struct {
	*common.BaseResponse

	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ImportKeyPairRequest struct {
	*common.BaseRequest

	// 密钥对名称。
	KeyName string `json:"keyName,omitempty"`

	// 密钥对的公钥内容
	PublicKey string `json:"publicKey,omitempty"`

	// 密钥对描述信息。
	KeyDescription *string `json:"keyDescription,omitempty"`
}

type ImportKeyPairResponse struct {
	*common.BaseResponse

	RequestId string `json:"requestId,omitempty"`

	Response *ImportKeyPairResponseParams `json:"response"`
}

type ImportKeyPairResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// 镜像ID。
	KeyId string `json:"keyId,omitempty"`
}
