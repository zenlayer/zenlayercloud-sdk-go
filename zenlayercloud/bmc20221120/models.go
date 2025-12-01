package bmc

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeZonesRequest struct {
	*common.BaseRequest

	// The languages of zones available. The optional values are as follows:
	// zh-CN: Chinese
	// en-US: English
	// Default value: en-US.
	AcceptLanguage string `json:"acceptLanguage,omitempty"`

	IsCloudRouterAvailable *bool `json:"isCloudRouterAvailable,omitempty"`
}

type DescribeZonesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeZonesResponseParams `json:"response"`
}

type DescribeZonesResponseParams struct {
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

	// City name of the zone.
	CityName string `json:"cityName,omitempty"`

	// Region name of the zone.
	AreaName string `json:"areaName,omitempty"`

	IsCloudRouterAvailable bool `json:"isCloudRouterAvailable,omitempty"`
}

type CreateInstancesRequest struct {
	*common.BaseRequest

	// Zone ID to which the instances belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Instance pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	InstanceChargeType string `json:"instanceChargeType,omitempty"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the instanceChargeType is PREPAID.
	InstanceChargePrepaid *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`

	// Instance model ID.
	// To view specific values, you can call DescribeInstanceTypes.
	InstanceTypeId string `json:"instanceTypeId,omitempty"`

	// Valid image ID.
	// To obtain valid image ID, you can call DescribeImages , pass in instanceTypeId to retrieve the list of images supported by the current model, and then find the imageId in the response.
	ImageId string `json:"imageId,omitempty"`

	// IpxeUrl
	IpxeUrl string `json:"ipxeUrl,omitempty"`

	// Resource group ID where the instances reside
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Instance name to be displayed.
	// Default value: instance.
	// This parameter can contain up to 64 characters. Only letters, numbers, - and periods (.) are supported.
	// If you purchase multiple instances at the same time, you can specify a pattern string [begin_number,bits].
	// begin_number: the starting value of an ordered number, value range: [0,99999], default value: 0.
	// bits: the number of digits occupied by an ordered value, value range: [1,6], default value: 6.
	// For example, if you purchase 1 instance and the instance name body is server_[3,3], the instance name will be server003; if you purchase 2 instances, the instance names will be server003, server004.
	// Note:
	// Spaces are not supported in the pattern string.
	// Multiple pattern strings are supported, such as server_[3,3]_[1,1].
	InstanceName string `json:"instanceName,omitempty"`

	// Instance hostname.
	// Default value: hostname.
	// This parameter can contain up to 64 . characters. Only letters, numbers, - and periods (.) are supported.
	// If you purchase multiple instances at the same time, you can specify a pattern string [begin_number,bits].
	// begin_number: the starting value of an ordered number, value range: [0,99999], default value: 0.
	// bits: the number of digits occupied by an ordered value, value range: [1,6], default value: 6.
	// For example, if you purchase 1 instance and the instance name body is server_[3,3], the instance name will be server003; if you purchase 2 instances, the instance names will be server003, server004.
	// Note:
	// Spaces are not supported in the pattern string.
	// Multiple pattern strings are supported, such as server_[3,3]_[1,1].
	Hostname string `json:"hostname,omitempty"`

	// The number of instances to be purchased.
	// Value range: [1,100]. Default value: 1.
	Amount int `json:"amount,omitempty"`

	// Instance password.
	// The parameter must be 8-16 characters, including uppercase letters, lowercase letters, numbers and special characters like 1~!@$^*-_=+. This password is also used as the password for IPMI login. Please keep it safe.
	// You must and can only pass the value of either password or sshKeys.
	Password string `json:"password,omitempty"`

	// ID of key pair.
	KeyId *string `json:"keyId,omitempty"`

	// List of SSH keys.
	// sshKeys and password cannot be specified at the same time. If an SSH key is used to log in, password login will be disabled. Up to 5 keys are supported.
	// Note:
	// For instances of Windows and EXSi operating systems, ignore this parameter. Default value is empty. Even if this parameter is filled in, only the value of password will be passed in.
	// If imageId is not specified, then sshKeys will be ignored.
	// SshKeys Deprecated, please using KeyId instead
	SshKeys []string `json:"sshKeys,omitempty"`

	// Network pricing model.
	// See InternetChargeType for details.
	InternetChargeType string `json:"internetChargeType,omitempty"`

	// Public network bandwidth cap (Mbps).
	// Default value: 1 Mbps.
	// The parameter value differs with different instance models. See bandwidth configuration for details.
	InternetMaxBandwidthOut int `json:"internetMaxBandwidthOut,omitempty"`

	// Traffic package size (TB).
	// The parameter is valid only when internetChargeType is ByTrafficPackage.
	// If not passed in, the default value will be the size of the free traffic package.
	TrafficPackageSize float64 `json:"trafficPackageSize,omitempty"`

	// Subnet ID.
	// Call DescribeVpcSubnets to check information on subnets created.
	SubnetId string `json:"subnetId,omitempty"`

	// Disk RAID configuration, including custom RAID settings.
	RaidConfig *RaidConfig `json:"raidConfig,omitempty"`

	// Disk partition configuration.
	// If the operating system is not installed, the partition cannot be set.
	Partitions []*Partition `json:"partitions,omitempty"`

	// NIC configuration.
	Nic *Nic `json:"nic,omitempty"`

	UserData *string `json:"userData,omitempty"`

	// ClusterId Bandwidth cluster ID.
	ClusterId string `json:"clusterId,omitempty"`

	// Whether to enable the instance primary IPv6.
	// If false, the primary IPv6 is not enabled. In this case, elastic IPv6 cannot be configured for the instance.
	// The default value is true.
	EnablePrimaryIPv6 *bool `json:"enablePrimaryIPv6,omitempty"`

	MarketingOptions  *MarketingInfo `json:"marketingOptions,omitempty"`

	Tags *TagAssociation  `json:"tags,omitempty"`
}


// TagAssociation 描述创建资源时同时绑定的标签对的信息。
type TagAssociation struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// Tag 描述一个标签键值对的信息。
type Tag struct {

    // Key 标签键。长度限制：1～128个字符。
    Key *string `json:"key,omitempty"`

    // Value 标签值。长度限制：1～128个字符。
    Value *string `json:"value,omitempty"`

}


type MarketingInfo struct {
	DiscountCode  string `json:"discountCode,omitempty"`
	UsePocVoucher bool `json:"usePocVoucher,omitempty"`
}

type RaidConfig struct {
	// RAID levels for rapid RAID settings.
	// Value possible: 0, 1, 5, 10.
	// Only one of raidType and customRaids can be specified.
	RaidType *int `json:"raidType,omitempty"`

	// RAID levels for custom RAID settings.
	// Only one of raidType and customRaids can be specified.
	CustomRaids []*CustomRaid `json:"customRaids,omitempty"`
}

type CustomRaid struct {
	// RAID levels.
	// Value possible: 0, 1, 5, 10.
	RaidType *int `json:"raidType,omitempty"`

	// Serial number of the disk
	// Numbered sequentially starting from 1. Multiple disk serial numbers should be consecutive.
	DiskSequence []int `json:"diskSequence,omitempty"`
}

type Partition struct {
	// Partition letter.
	// Linux: must start with "/", and the first system partition must be "/".
	// Windows: support C~H, the first system partition must be designated as "C".
	FsPath string `json:"fsPath,omitempty"`

	// File type of partition.
	// Linux: ext2, ext3, ext4, ext type is needed.
	// Windows: NTFS.
	FsType string `json:"fsType,omitempty"`

	// Partition size.
	// Unit: GB.
	Size int `json:"size,omitempty"`
}

type Nic struct {
	// Public NIC name.
	// Only numbers, uppercase and lowercase letters are allowed, starting with a letter with a length limit of 4 to 10 characters.
	// For non-high-availability models, the default public NIC name is wan0. It cannot start with lan.
	// For high-availability models, the default public NIC name is bond0.
	// Public and private NIC names cannot be the same.
	WanName string `json:"wanName,omitempty"`

	// Private NIC name.
	// Only numbers, uppercase and lowercase letters are allowed, starting with a letter with a length limit of 4 to 10 characters.
	// For non-high-availability models, the default private NIC name is lan0. It cannot start with wan.
	// For high-availability models, the default private NIC name is bond1.
	// Public and private NIC names cannot be the same.
	LanName string `json:"lanName,omitempty"`
}

type CreateInstancesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateInstanceResponseParams `json:"response"`
}

type CreateInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"requestId,omitempty"`

	// List of instance IDs.
	// The returned instance ID list does not mean the creation has been completed. The status of the instance will be PENDING or CREATING during the creation. You can call DescribeInstances to query the status of the instance according to dataSet. If the status changes from PENDING or Creating to Running, it means that the instance has been created successfully; CREATE_FAILED means the instance has been created failed.
	InstanceIdSet []*string `json:"instanceIdSet,omitempty"`

	// Number of order.
	OrderNumber *string `json:"orderNumber,omitempty"`
}

type ChargePrepaid struct {
	// Period of subscription.
	// Unit: month.
	Period int `json:"period,omitempty"`
}

type DescribeInstanceTypesRequest struct {
	*common.BaseRequest

	// Instance models that support specified image.
	ImageId string `json:"imageId,omitempty"`

	// Instance model ID.
	// Maximum number: 100
	InstanceTypeIds []string `json:"instanceTypeIds,omitempty"`

	// Desired minimum number of CPU cores.
	// Value range: positive integer.
	MinimumCpuCoreCount *int `json:"minimumCpuCoreCount,omitempty"`

	// Desired maximum number of CPU cores.
	// Value range: positive integer.
	MaximumCpuCoreCount *int `json:"maximumCpuCoreCount,omitempty"`

	// Desired minimum memory size.
	// Value range: positive integer.
	// Unit: GB.
	MinimumMemorySize *int `json:"minimumMemorySize,omitempty"`

	// Desired maximum memory size.
	// Value range: positive integer.
	// Unit: GB.
	MaximumMemorySize *int `json:"maximumMemorySize,omitempty"`

	// Desired minimum public inbound bandwidth cap.
	// Unit: Mbps.
	MinimumBandwidth *int `json:"minimumBandwidth,omitempty"`

	// Supported RAID levels.
	// Value range: 0, 1, 5, 10.
	SupportRaids []int `json:"supportRaids,omitempty"`

	// Subnet supported or not.
	SupportSubnet *bool `json:"supportSubnet,omitempty"`

	// Desired minimum disk size.
	// Value range: positive integer.
	// Unit: GB.
	MinimumDiskSize *int `json:"minimumDiskSize,omitempty"`

	// Desired maximum memory size.
	// Value range: positive integer.
	// Unit: GB.
	MaximumDiskSize *int `json:"maximumDiskSize,omitempty"`

	// High availability supported or not.
	IsHA *bool `json:"isHA,omitempty"`
}

type DescribeInstanceTypesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstanceTypesResponseParams `json:"response"`
}

type DescribeInstanceTypesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Information on instance model.
	InstanceTypes []*InstanceType `json:"instanceTypes,omitempty"`
}

type InstanceType struct {
	// Image ID supported. Hint: only accessible in DescribeInstanceType
	ImageIds []string `json:"imageIds,omitempty"`

	// Instance model ID.
	InstanceTypeId string `json:"instanceTypeId,omitempty"`

	// Model description, including memory size, disk, etc.
	Description string `json:"description,omitempty"`

	// Quantity of CPU cores.
	CpuCoreCount int `json:"cpuCoreCount,omitempty"`

   	CpuDetail string `json:"cpuDetail,omitempty"`

   	CpuCores int `json:"cpuCores,omitempty"`

   	CpuThreads int `json:"cpuThreads,omitempty"`

   	BaseFrequency string `json:"baseFrequency,omitempty"`

	// Memory size.
	// Unit: GB.
	MemorySize int `json:"memorySize,omitempty"`

	// Outbound bandwidth cap.
	// Unit: Mbps.
	MaximumBandwidth int `json:"maximumBandwidth,omitempty"`

	// RAID level supported.
	SupportRaids []int `json:"supportRaids,omitempty"`

	// Subnet supported or not.
	SupportSubnet bool `json:"supportSubnet,omitempty"`

	// High availability supported or not.
	IsHA bool `json:"isHA,omitempty"`

	// Disk size.
	// Unit: GB.
	DiskInfo *InstanceDiskInfo `json:"diskInfo,omitempty"`
}

type InstanceDiskInfo struct {
	// Total disk size.
	// Unit: GB.
	// The available storage space is relatively less than the total disk size to create partitions successfully. What remains actually will be added to the last partition.
	TotalDiskSize int `json:"totalDiskSize,omitempty"`

	// Description of the disks.
	DiskDescription string `json:"diskDescription,omitempty"`

	// Disk information available for RAID and partition.
	// Numbered sequentially. For example, for 880 x 2 and 220 x 2, the disk serial numbers 1, 2, 3, and 4 respectively correspond to disk sizes of 880 GB, 880 GB, 220 GB, and 220 GB.
	Disks []*Disk `json:"disks,omitempty"`
}

type Disk struct {
	// Disk size.
	// Unit: GB.
	DiskSize int `json:"diskSize,omitempty"`

	// Quantity of disks of the size.
	DiskCount int `json:"diskCount,omitempty"`
}

type DescribeAvailableResourcesRequest struct {
	*common.BaseRequest

	// Instance model ID.
	InstanceTypeId string `json:"instanceTypeId,omitempty"`

	// Instance pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	InstanceChargeType string `json:"instanceChargeType,omitempty"`

	// Zone ID.
	ZoneId string `json:"zoneId,omitempty"`

	// Status of sale.
	// SELL: available for sale, stock > 10.
	// SELL_SHORTAGE: available for sale, stock < 10.
	// SOLD_OUT: sold out.
	SellStatus string `json:"sellStatus,omitempty"`
}

type DescribeAvailableResourcesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAvailableResourcesResponseParams `json:"response"`
}

type DescribeAvailableResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Status of available resources.
	AvailableResources []*AvailableResource `json:"availableResources,omitempty"`
}
type AvailableResource struct {
	// Zone ID
	ZoneId string `json:"zoneId,omitempty"`

	// Status for sale.
	// SELL: available for sale, stock > 10.
	// SELL_SHORTAGE: available for sale, stock < 10.
	// SOLD_OUT: sold out.
	SellStatus string `json:"sellStatus,omitempty"`

	// Network pricing model.
	// See InternetChargeType for details.
	InternetChargeTypes []string `json:"internetChargeTypes,omitempty"`

	// ID of the instance model.
	InstanceTypeId string `json:"instanceTypeId,omitempty"`

	// Public outbound bandwidth cap.
	// Unit: Mbps.
	MaximumBandwidthOut int `json:"maximumBandwidthOut,omitempty"`

	// Default free public bandwidth for the pricing model of flat rate.
	// Unit: GB.
	DefaultBandwidthOut int `json:"defaultBandwidthOut,omitempty"`

	// Default free traffic package for the pricing model of data transfer.
	// Unit: TB.
	DefaultTrafficPackageSize float64 `json:"defaultTrafficPackageSize,omitempty"`

	// The inventory quantity.
	Qty int `json:"qty"`
}

type DescribeImagesRequest struct {
	*common.BaseRequest

	// List of image IDs.
	ImageIds []string `json:"imageIds,omitempty"`

	// Image name.
	ImageName string `json:"imageName,omitempty"`

	// Image catalog. The optional values are as follows:
	// CentOS
	// Windows
	// Ubuntu
	// Debian
	// ESXi
	Catalog string `json:"catalog,omitempty"`

	// Image type. The optional values are as follows:
	// PUBLIC_IMAGE: public images
	// CUSTOM_IMAGE: custom images
	// You cannot create a custom image by yourself currently. If necessary, please submit a ticket.
	ImageType string `json:"imageType,omitempty"`

	// Operating system type. The optional values are as follows:
	// Windows
	// Linux
	OsType string `json:"osType,omitempty"`

	// Supported instance model ID.
	InstanceTypeId string `json:"instanceTypeId,omitempty"`
}

type DescribeImagesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeImagesResponseParams `json:"response"`
}

type DescribeImagesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Information on images.
	Images []*ImageInfo `json:"images,omitempty"`
}

type ImageInfo struct {
	// Image ID.
	ImageId string `json:"imageId,omitempty"`

	// Image name.
	ImageName string `json:"imageName,omitempty"`

	// Image catalog.
	// The optional values are as follows:
	// CentOS
	// Windows
	// Ubuntu
	// Debian
	// ESXi
	Catalog string `json:"catalog,omitempty"`

	// Image type.
	// The optional values are as follows:
	// PUBLIC_IMAGE: public image.
	// CUSTOM_IMAGE: custom image.
	// You cannot create a custom image by yourself currently. If necessary, please submit a ticket.
	ImageType string `json:"imageType,omitempty"`

	// Operating system type.
	// The optional values are as follows:
	// Windows
	// Linux
	OsType string `json:"osType,omitempty"`
}

type DescribeInstancesRequest struct {
	*common.BaseRequest
	// Instance IDs.
	// You can query up to 100 instances in each request.
	InstanceIds []string `json:"instanceIds,omitempty"`

	// Zone ID to which the instances belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Resource group ID.
	// If the value is null, then return all the instances in the authorized resource groups.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Instance model ID.
	// You can call DescribeInstanceTypes to obtain the latest specification.
	InstanceTypeId string `json:"instanceTypeId,omitempty"`

	// Network pricing model.
	// See InternetChargeType for details.
	InternetChargeType string `json:"internetChargeType,omitempty"`

	// Image ID.
	ImageId string `json:"imageId,omitempty"`

	// Subnet ID.
	// You can call DescribeVpcSubnets to query information on subnet.
	SubnetId string `json:"subnetId,omitempty"`

	// Instance status.
	// See InstanceStatus for details.
	InstanceStatus string `json:"instanceStatus,omitempty"`

	// Instance name.
	// If the value ends with *, a fuzzy match will be performed on instanceName, otherwise an exact match will be performed.
	InstanceName string `json:"instanceName,omitempty"`

	// Instance hostname.
	// If the value ends with *, a fuzzy match will be performed on hostname, otherwise an exact match will be performed.
	Hostname string `json:"hostname,omitempty"`

	// List of public IPs of the instance.
	PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`

	// List of private IPs of the instance.
	PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`

	// Number of pages returned.
	// Default value: 1
	PageNum int `json:"pageNum,omitempty"`

	// Number of items in the current page result.
	// Default value: 20
	// Maximum value: 1000
	PageSize int `json:"pageSize,omitempty"`

	TagKeys  []string `json:"tagKeys,omitempty"`

    Tags   []*Tag `json:"tags,omitempty"`
}

type InstanceInfo struct {
	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`

	// Zone ID to which the instances belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Instance name to be displayed.
	InstanceName string `json:"instanceName,omitempty"`

	// Instance hostname.
	Hostname string `json:"hostname,omitempty"`

	// Instance model ID.
	InstanceTypeId string `json:"instanceTypeId,omitempty"`

	// Instance model information.
	InstanceType *InstanceType `json:"instanceType,omitempty"`

	// Image ID.
	ImageId *string `json:"imageId,omitempty"`

	// Image name.
	ImageName string `json:"imageName,omitempty"`

	// IpxeUrl.
	IpxeUrl string `json:"ipxeUrl,omitempty"`

	// Instance pricing model.
	// PREPAID: monthly subscription.
	// POSTPAID: pay-as-you-go.
	InstanceChargeType string `json:"instanceChargeType,omitempty"`

	// Public outbound bandwidth.
	// Unit: Mbps.
	// Value 0 means no limit, but not exceeds the upper limit of the instance model supported.
	BandwidthOutMbps *int `json:"bandwidthOutMbps,omitempty"`

	// Traffic package size purchased.
	// Unit: TB.
	TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

	// Network pricing model.
	// See InternetChargeType for details.
	InternetChargeType string `json:"internetChargeType,omitempty"`

	// Period of instance subscription.
	// Unit: month.
	// For postpaid instances, the value is empty.
	Period *int `json:"period,omitempty"`

	// Instance primary IP.
	PrimaryPublicIpAddress *string `json:"primaryPublicIpAddress,omitempty"`

	// List of public IPv4 addresses.
	// If the instance primary IP is not added to the public network interface, the primary IP will not be available., and its value will not be returned.
	PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`

	// List of private IPv4 addresses.
	PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`

	// Instance IPv6 addresses.
	// The value may be empty, which means no available IPv6 address exists.
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty"`

	// List of subnet IDs.
	SubnetIds []string `json:"subnetIds,omitempty"`

	// Creation time.
	// Use UTC time according to the ISO8601 standard. Format: YYYY-MM-DDThh:mm:ssZ.
	CreateTime string `json:"createTime,omitempty"`

	// Expiration time.
	// Use UTC time according to the ISO8601 standard. Format: YYYY-MM-DDThh:mm:ssZ.
	ExpiredTime *string `json:"expiredTime,omitempty"`

	// Resource group ID to which instances belong.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Resource group name to which instances belong.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// Status of instances.
	// See InstanceStatus for details.
	InstanceStatus string `json:"instanceStatus,omitempty"`

	// Partition configuration.
	Partitions []*Partition `json:"partitions,omitempty"`

	// Disk array configuration.
	RaidConfig *RaidConfig `json:"raidConfig,omitempty"`

	// NIC configuration.
	Nic *Nic `json:"nic,omitempty"`

	// 是否自动续费.
	AutoRenew bool `json:"autoRenew,omitempty"`

	KeyId *string `json:"keyId,omitempty"`

	Tags  *Tags    `json:"tags,omitempty"`
}


// Tags 描述资源关联的标签信息。
type Tags struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Number of instances meeting the filtering conditions.
	TotalCount int `json:"totalCount,omitempty"`

	// Information on an instance.
	DataSet []*InstanceInfo `json:"dataSet,omitempty"`
}

type DescribeInstancesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstancesResponseParams `json:"response"`
}

type StartInstancesRequest struct {
	*common.BaseRequest

	// Instance ID(s).
	// To obtain the instance IDs, you can call DescribeInstances and look for instanceId in the response. The maximum number of instances in each request is 100.
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type StartInstancesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type StopInstancesRequest struct {
	*common.BaseRequest

	// Instance ID(s).
	// To obtain the instance IDs, you can call DescribeInstances and look for instanceId in the response. The maximum number of instances in each request is 100.
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type StopInstancesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RebootInstancesRequest struct {
	*common.BaseRequest

	// Instance ID(s).
	// To obtain the instance IDs, you can call DescribeInstances and look for instanceId in the response. The maximum number of instances in each request is 100.
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type RebootInstancesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReinstallInstanceRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`

	// Valid image ID.
	// To obtain valid image ID, you can call DescribeImages , pass in instanceTypeId to retrieve the list of images supported by the current model, and then find the imageId in the response.
	ImageId string `json:"imageId,omitempty"`

	// IpxeUrl。
	IpxeUrl string `json:"ipxeUrl,omitempty"`

	// Instance hostname.
	// Default value: hostname.
	// This parameter can contain up to 64 . characters. Only letters, numbers, - and periods (.) are supported.
	Hostname string `json:"hostname,omitempty"`

	// Instance password.
	// The parameter must be 8-16 characters, including uppercase letters, lowercase letters, numbers and special characters like 1~!@$^*-_=+. This password is also used as the password for IPMI login. Please keep it safe.
	// You must and can only pass the value of either password or sshKeys.
	Password string `json:"password,omitempty"`

	// ID of key pair.
	KeyId *string `json:"keyId,omitempty"`

	// List of SSH keys.
	// sshKeys and password cannot be specified at the same time. If an SSH key is used to log in, password login will be disabled. Up to 5 keys are supported.
	// Note:
	// For instances of Windows and EXSi operating systems, ignore this parameter. Default value is empty. Even if this parameter is filled in, only the value of password will be passed in.
	// If imageId is not specified, then sshKeys will be ignored.
	// SshKeys Deprecated, please using KeyId instead
	SshKeys []string `json:"sshKeys,omitempty"`

	// Disk array configuration.
	RaidConfig *RaidConfig `json:"raidConfig,omitempty"`

	// Disk partition configuration.
	// If the operating system is not installed, the partition cannot be set.
	Partitions []*Partition `json:"partitions,omitempty"`

	// NIC configuration.
	Nic *Nic `json:"nic,omitempty"`

	UserData *string `json:"userData,omitempty"`
}

type ReInstallInstanceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type TerminateInstanceRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}

type TerminateInstanceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseInstancesRequest struct {
	*common.BaseRequest

	// Instance ID(s).
	// To obtain the instance IDs, you can call DescribeInstances and look for instanceId in the response. The maximum number of instances in each request is 100.
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type ReleaseInstancesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RenewInstanceRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}

type RenewInstanceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyInstancesAttributeRequest struct {
	*common.BaseRequest

	// Instance ID(s).
	// To obtain the instance IDs, you can call DescribeInstances and look for instanceId in the response. The maximum number of instances in each request is 100.
	InstanceIds []string `json:"instanceIds,omitempty"`

	// Instance name to be displayed.
	// This parameter can contain up to 64 characters. Only letters, numbers, - and periods (.) are supported.
	InstanceName string `json:"instanceName,omitempty"`
}

type ModifyInstancesAttributeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeInstanceInternetStatusRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}

type DescribeInstanceInternetStatusResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InstanceInternetStatus `json:"response"`
}

type InstanceInternetStatus struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`

	// Instance name.
	InstanceName string `json:"instanceName,omitempty"`

	// Current instance bandwidth.
	InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

	// Modified instance bandwidth.
	ModifiedInternetMaxBandwidthOut *int `json:"modifiedInternetMaxBandwidthOut,omitempty"`

	// Status of instance bandwidth.
	// Processing: modifying.
	// Enable: effective now.
	// WaitToEnable: effective in next billing cycle.
	ModifiedBandwidthStatus string `json:"modifiedBandwidthStatus,omitempty"`

	// Current instance traffic package
	TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

	// Modified instance traffic package
	ModifiedTrafficPackageSize *float64 `json:"modifiedTrafficPackageSize,omitempty"`

	// Status of instance traffic package
	// Processing: modifying.
	// Enable: effective now.
	// WaitToEnable: effective in next billing cycle.
	ModifiedTrafficPackageStatus string `json:"modifiedTrafficPackageStatus,omitempty"`
}

type ModifyInstancesResourceGroupRequest struct {
	*common.BaseRequest

	// List of instance IDs.
	// The maximum number of instances in each request is 100.
	InstanceIds []string `json:"instanceIds,omitempty"`

	// Resource group ID.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
}

type ModifyInstancesResourceGroupResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
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

type InquiryPriceCreateInstanceRequest struct {
	*common.BaseRequest

	// Zone ID of the instance.
	ZoneId string `json:"zoneId,omitempty"`

	// Instance model ID.
	// To view specific values, you can call DescribeInstanceTypes.
	InstanceTypeId string `json:"instanceTypeId,omitempty"`

	// Instance pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	InstanceChargeType string `json:"instanceChargeType,omitempty"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the instanceChargeType is PREPAID.
	InstanceChargePrepaid *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`

	// Traffic package size (TB).
	// The parameter is valid only when internetChargeType is ByTrafficPackage.
	// If not passed in, the default value will be the size of the free traffic package.
	TrafficPackageSize float64 `json:"trafficPackageSize,omitempty"`

	// Public network bandwidth cap (Mbps).
	// Default value: 1 Mbps.
	// The parameter value differs by different instance models. See bandwidth configuration for details.
	InternetMaxBandwidthOut int `json:"internetMaxBandwidthOut,omitempty"`

	// Network pricing model.
	// See InternetChargeType for details.
	InternetChargeType string `json:"internetChargeType,omitempty"`
}

type InquiryPriceCreateInstanceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InquiryPriceCreateInstanceResponseParams `json:"response"`
}

type InquiryPriceCreateInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Price of the instance.
	InstancePrice *Price `json:"instancePrice,omitempty"`

	// Price of public bandwidth.
	// Kinds of prices may exist. For example, traffic package billing method may contain the package price and overage price.
	BandwidthPrice []*Price `json:"bandwidthPrice,omitempty"`
}

type Price struct {
	// Discount.
	// For example, 80.0 means 20% off.
	Discount *float64 `json:"discount,omitempty"`

	// Discount price of prepaid resources.
	// Only used in subscription model. For pay-as-you-go model, the value is empty.
	DiscountPrice *float64 `json:"discountPrice,omitempty"`

	// Original price of prepaid resources.
	// Only used in subscription model. For pay-as-you-go model, the value is empty.
	OriginalPrice *float64 `json:"originalPrice,omitempty"`

	// Original unit price of postpaid resources.
	// Only used in pay-as-you-go model. For tiered billing, the value is empty.
	UnitPrice *float64 `json:"unitPrice,omitempty"`

	// Discount unit price of postpaid resources.
	// Only used in pay-as-you-go model. For tiered billing, the value is empty.
	DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`

	// Unit of postpaid billing.
	// Only used in pay-as-you-go model.
	// Value range:
	// HOUR: you will be billed by hour.
	// DAY: you will be billed by day.
	// MONTH: you will be billed by month. For example, the burstable 95th pricing model.
	ChargeUnit *string `json:"chargeUnit,omitempty"`

	// Tiered price of postpaid billing.
	// Only used in pay-as-you-go model. If it is not tiered price, the value is empty.
	StepPrices []*StepPrice `json:"stepPrices,omitempty"`
}

type StepPrice struct {
	// First price range of a tiered price.
	StepStart *float64 `json:"stepStart,omitempty"`

	// Last price range of a tiered price.
	StepEnd *float64 `json:"stepEnd,omitempty"`

	// Original unit price of current price range.
	// Only used in pay-as-you-go model.
	UnitPrice *float64 `json:"unitPrice,omitempty"`

	// Discount unit price of current price range.
	// Only used in pay-as-you-go model.
	DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`
}

type InquiryPriceInstanceBandwidthRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`

	// Bandwidth size.
	BandwidthOutMbps int `json:"bandwidthOutMbps,omitempty"`
}
type InquiryPriceInstanceBandwidthResponse struct {
	*common.BaseResponse

	// // The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InquiryPriceInstanceBandwidthResponseParams `json:"response"`
}
type InquiryPriceInstanceBandwidthResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Price of public bandwidth.
	// Kinds of prices may exist. For example, traffic package billing method may contain the package price and overage price.
	BandwidthPrice []*Price `json:"bandwidthPrice,omitempty"`
}

type ModifyInstanceBandwidthRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`

	// Bandwidth.
	// Value range: from 1 to maximum supported for the instance model.
	BandwidthOutMbps *int `json:"bandwidthOutMbps,omitempty"`
}

type ModifyInstanceBandwidthResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *ModifyInstanceBandwidthResponseParams `json:"response"`
}

type ModifyInstanceBandwidthResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Number of order.
	OrderNumber string `json:"orderNumber,omitempty"`
}

type CancelInstanceBandwidthDowngradeRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}

type CancelInstanceBandwidthDowngradeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyInstanceTrafficPackageRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`

	// Traffic package size.
	TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`
}

type ModifyInstanceTrafficPackageSizeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Number of order.
	OrderNumber string `json:"orderNumber,omitempty"`
}

type ModifyInstanceTrafficPackageResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *ModifyInstanceTrafficPackageSizeResponseParams `json:"response"`
}

type InquiryPriceInstanceTrafficPackageRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`

	// Traffic package size.
	TrafficPackageSize float64 `json:"trafficPackageSize,omitempty"`
}

type InquiryPriceInstanceTrafficPackageResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InquiryPriceInstanceTrafficPackageResponseParams `json:"response"`
}

type InquiryPriceInstanceTrafficPackageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Price of traffic package.
	// Kinds of prices may exist. For example, traffic package billing method may contain the package price and overage price.
	TrafficPackagePrice []*Price `json:"trafficPackagePrice,omitempty"`
}

type CancelInstanceTrafficPackageDowngradeRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}

type CancelInstanceTrafficPackageDowngradeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeInstancesMonitorHealthRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type DescribeInstancesMonitorHealthResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstancesMonitorHealthResponseParams `json:"response"`
}

type DescribeInstancesMonitorHealthResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Price of traffic package.
	// Kinds of prices may exist. For example, traffic package billing method may contain the package price and overage price.
	MonitorHealthList []*InstanceHealth `json:"monitorHealthList,omitempty"`
}

type InstanceHealth struct {
	// 实例ID。
	InstanceId string `json:"instanceId,omitempty"`

	// CPU状态。
	CpuStatus string `json:"cpuStatus,omitempty"`

	// Disk状态。
	DiskStatus string `json:"diskStatus,omitempty"`

	// Ipmi IP状态。
	IpmiPing string `json:"ipmiPing,omitempty"`

	// Ipmi状态。
	IpmiStatus string `json:"ipmiStatus,omitempty"`

	// Memory状态。
	MemoryStatus string `json:"memoryStatus,omitempty"`

	// Power Supply状态。
	PsuStatus string `json:"psuStatus,omitempty"`

	// 服务器公网口连接的交换机端口的状态。
	WanPortStatus string `json:"wanPortStatus,omitempty"`

	// 风扇状态。
	FanStatus string `json:"fanStatus,omitempty"`

	// 服务器供应商品牌。
	ServerBrand string `json:"serverBrand,omitempty"`

	// 服务器供应商型号。
	ServerModel string `json:"serverModel,omitempty"`

	// 超微 Supermicro 对于刀片机单 CPU 的温度。
	CpuTemp int `json:"cpuTemp,omitempty"`

	// CPU 的温度。
	Cpu0Temp int `json:"cpu0Temp,omitempty"`

	Cpu1Temp int `json:"cpu1Temp,omitempty"`

	Cpu2Temp int `json:"cpu2Temp,omitempty"`

	// 服务器在机房的温度。
	InletTemp int `json:"inletTemp,omitempty"`

	// 温度单位。
	TempUnit string `json:"tempUnit,omitempty"`
}

type EipAddress struct {
	EipId             string `json:"eipId,omitempty"`
	ZoneId            string `json:"zoneId,omitempty"`
	IpAddress         string `json:"ipAddress,omitempty"`
	InstanceId        string `json:"instanceId,omitempty"`
	InstanceName      string `json:"instanceName,omitempty"`
	EipChargeType     string `json:"eipChargeType,omitempty"`
	Period            *int   `json:"period,omitempty"`
	CreateTime        string `json:"createTime,omitempty"`
	ExpiredTime       string `json:"expiredTime,omitempty"`
	ResourceGroupId   string `json:"resourceGroupId,omitempty"`
	ResourceGroupName string `json:"resourceGroupName,omitempty"`
	EipStatus         string `json:"eipStatus,omitempty"`
    Tags              *Tags `json:"tags,omitempty"`
}

type DescribeEipAddressesRequest struct {
	*common.BaseRequest

	// ID list of elastic IPs.
	// The maximum number of elastic IPs in each request is 100.
	EipIds []string `json:"eipIds,omitempty"`

	// Elastic IP pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	EipChargeType string `json:"eipChargeType,omitempty"`

	// IP address.
	IpAddress string `json:"ipAddress,omitempty"`

	// Zone ID to which the elastic IPs belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Resource group ID.
	// If this parameter is not passed in, all the elastic IPs in authorized resource group return.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Status of elastic IPs.
	EipStatus string `json:"eipStatus,omitempty"`

	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`

	// Instance name.
	InstanceName string `json:"instanceName,omitempty"`

	// Number of items in the current page result.
	// Default value: 20
	// Maximum value: 1000
	PageSize int `json:"pageSize,omitempty"`

	// Number of pages returned.
	// Default value: 1
	PageNum int `json:"pageNum,omitempty"`

	TagKeys  []string `json:"tagKeys,omitempty"`

    Tags   []*Tag `json:"tags,omitempty"`
}


type DescribeEipAddressesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeEipAddressesResponseParams `json:"response,omitempty"`
}
type DescribeEipAddressesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// List of elastic IPs.
	DataSet []*EipAddress `json:"dataSet,omitempty"`

	// The total number of elastic IPs.
	TotalCount int `json:"totalCount,omitempty"`
}

type DescribeEipAvailableResourcesRequest struct {
	*common.BaseRequest

	// Elastic IP pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	EipChargeType string `json:"eipChargeType,omitempty"`

	// Zone ID to which the elastic IPs belong.
	// If this parameter is not passed in, the elastic IPs in all zones return.
	ZoneId string `json:"zoneId,omitempty"`
}
type DescribeEipAvailableResourcesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeEipAvailableResourcesResponseParams `json:"response,omitempty"`
}
type DescribeEipAvailableResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Zone list of available elastic IPs.
	EipResources []*EipAvailableResource `json:"eipResources,omitempty"`
}
type EipAvailableResource struct {
	// Zone ID to which the elastic IPs belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Status for sale.
	// SELL: available for sale, stock > 10.
	// SELL_SHORTAGE: available for sale, stock < 10.
	// SOLD_OUT: sold out.
	Status string `json:"status,omitempty"`
}

type DescribeInstanceAvailableEipResourcesRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}
type DescribeInstanceAvailableEipResourcesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstanceAvailableEipResourcesResponseParams `json:"response,omitempty"`
}
type DescribeInstanceAvailableEipResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// List of available elastic IPs that can be bound to an instance.
	InstanceEipResources []*InstanceAvailableEipResource `json:"instanceEipResources,omitempty"`
}
type InstanceAvailableEipResource struct {
	// ID of an elastic IP.
	// To obtain the elastic IP ID, you can call DescribeEipAddresses and look for eipId in the response.
	EipId string `json:"eipId,omitempty"`

	// IP address.
	IpAddress string `json:"ipAddress,omitempty"`
}

type AllocateEipAddressesRequest struct {
	*common.BaseRequest

	// Elastic IP pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	EipChargeType string `json:"eipChargeType,omitempty"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the eipChargeType is PREPAID.
	EipChargePrepaid *ChargePrepaid `json:"eipChargePrepaid,omitempty"`

	// Zone ID to which the elastic IPs belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Resource group ID
	// If the value is not passed in, the elastic IP will be put into the default resource group. If no authorized default resource group found, the request will fail.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Quantity of elastic IPs.
	// Value range: 1-100.
	// Default value: 1.
	Amount int `json:"amount,omitempty"`

	MarketingOptions  *MarketingInfo `json:"marketingOptions,omitempty"`

	Tags *TagAssociation  `json:"tags,omitempty"`
}


type AllocateEipAddressesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *AllocateEipAddressesResponseParams `json:"response,omitempty"`
}

type AllocateEipAddressesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"requestId,omitempty"`

	// ID list of elastic IPs.
	// The returned ID list does not mean the creation has been completed. You can call DescribeEipAddresses to query the status of the elastic IPs. If the status changes from Creating to Available, it means that the elastic IPs have been created successfully.
	EipIdSet []*string `json:"eipIdSet,omitempty"`

	// Number of order.
	// This parameter returns when eipChargeType is PREPAID.
	OrderNumber *string `json:"orderNumber,omitempty"`
}

type TerminateEipAddressRequest struct {
	*common.BaseRequest

	// ID of an elastic IP.
	// To obtain the elastic IP ID, you can call DescribeEipAddresses and look for eipId in the response.
	EipId string `json:"eipId,omitempty"`
}

type TerminateEipAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseEipAddressesRequest struct {
	*common.BaseRequest

	// IDs of elastic IPs.
	// To obtain IDs of the elastic IPs, you can call DescribeEipAddresses and look for eipId in the response.
	EipIds []string `json:"eipIds,omitempty"`
}

type ReleaseEipAddressesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RenewEipAddressRequest struct {
	*common.BaseRequest

	// ID of an elastic IP.
	// To obtain the elastic IP ID, you can call DescribeEipAddresses and look for eipId in the response.
	EipId string `json:"eipId,omitempty"`
}

type RenewEipAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type AssociateEipAddressRequest struct {
	*common.BaseRequest

	// ID of an elastic IP.
	// To obtain the elastic IP ID, you can call DescribeEipAddresses and look for eipId in the response.
	EipId string `json:"eipId,omitempty"`

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}

type AssociateEipAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type UnassociateEipAddressRequest struct {
	*common.BaseRequest

	// ID of an elastic IP.
	// To obtain the elastic IP ID, you can call DescribeEipAddresses and look for eipId in the response.
	EipId string `json:"eipId,omitempty"`
}

type UnassociateEipAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type InquiryPriceCreateEipAddressRequest struct {
	*common.BaseRequest

	// Zone ID to which the elastic IP belongs.
	ZoneId string `json:"zoneId,omitempty"`

	// Elastic IP pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	EipChargeType string `json:"eipChargeType,omitempty"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the eipChargeType is PREPAID.
	EipChargePrepaid *ChargePrepaid `json:"eipChargePrepaid,omitempty"`

	// Quantity of elastic IPs.
	// Value range: 1-100.
	// Default value: 1.
	Amount int `json:"amount,omitempty"`
}

type InquiryPriceCreateEipAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InquiryPriceCreateEipAddressResponseParams `json:"response"`
}

type InquiryPriceCreateEipAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Price of the elastic IP.
	EipPrice *Price `json:"eipPrice,omitempty"`
}

type ModifyEipAddressesResourceGroupRequest struct {
	*common.BaseRequest

	// List of EIP IDs.
	// The maximum number of eip in each request is 100.
	EipIds []string `json:"eipIds,omitempty"`

	// Resource group ID.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
}

type ModifyEipAddressesResourceGroupResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeDdosIpAvailableResourcesRequest struct {
	*common.BaseRequest

	// DDoS protected IP pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	DdosIpChargeType string `json:"ddosIpChargeType,omitempty"`

	// Zone ID to which the DDoS protected IPs belong.
	// If this parameter is not passed in, the DDoS protected IPs in all zones return.
	ZoneId string `json:"zoneId,omitempty"`
}
type DescribeDdosIpAvailableResourcesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeDdosIpAvailableResourcesResponseParams `json:"response,omitempty"`
}
type DescribeDdosIpAvailableResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Zone list of available DDoS protected IPs.
	DdosIpResources []*DdosIpAvailableResource `json:"ddosIpResources,omitempty"`
}

type DdosIpAvailableResource struct {
	// Zone ID to which the DDoS protected IPs belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Status for sale.
	// SELL: available for sale, stock > 10.
	// SELL_SHORTAGE: available for sale, stock < 10.
	// SOLD_OUT: sold out.
	Status string `json:"status,omitempty"`
}

type DescribeInstanceAvailableDdosResourcesRequest struct {
	*common.BaseRequest

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}
type DescribeInstanceAvailableDdosResourcesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstanceAvailableDdosResourcesResponseParams `json:"response,omitempty"`
}
type DescribeInstanceAvailableDdosResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// List of available DDoS protected IPs that can be bound to an instance.
	InstanceDdosIpResources []*InstanceAvailableDdosIpResource `json:"instanceDdosIpResources,omitempty"`
}
type InstanceAvailableDdosIpResource struct {
	// ID of a DDoS protected IP.
	// To obtain the DDoS protected IP ID, you can call DescribeDdosIpAddresses and look for ddosIpId in the response.
	DdosIpId string `json:"ddosIpId,omitempty"`

	// IP address.
	IpAddress string `json:"ipAddress,omitempty"`
}

type AllocateDdosIpAddressesRequest struct {
	*common.BaseRequest

	// DDoS protected IP pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	DdosIpChargeType string `json:"ddosIpChargeType,omitempty"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the ddosIpChargeType is PREPAID.
	DdosIpChargePrepaid *ChargePrepaid `json:"ddosIpChargePrepaid,omitempty"`

	// Zone ID to which the DDoS protected IPs belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Resource group ID
	// If the value is not passed in, the DDoS protected IP will be put into the default resource group. If no authorized default resource group found, the request will fail.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Quantity of DDoS protected IPs.
	// Value range: 1-100.
	// Default value: 1.
	Amount int `json:"amount,omitempty"`
}
type AllocateDdosIpAddressesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *AllocateDdosIpAddressesResponseParams `json:"response,omitempty"`
}

type AllocateDdosIpAddressesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"requestId,omitempty"`

	// ID list of DDoS protected IPs.
	// The returned ID list does not mean the creation has been completed. You can call DescribeDdosIpAddresses to query the status of the DDoS protected IPs. If the status changes from Creating to Available, it means that the DDoS protected IPs have been created successfully.
	DdosIdSet []*string `json:"ddosIdSet,omitempty"`

	// Number of order.
	// This parameter returns when ddosIpChargeType is PREPAID.
	OrderNumber *string `json:"orderNumber,omitempty"`
}

type DescribeDdosIpAddressesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeDdosIpAddressesResponseParams `json:"response,omitempty"`
}
type DescribeDdosIpAddressesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// List of DDoS protected IPs.
	DataSet []*DdosIpAddress `json:"dataSet,omitempty"`

	// The total number of DDoS protected IPs.
	TotalCount int `json:"totalCount,omitempty"`
}

type DescribeDdosIpAddressesRequest struct {
	*common.BaseRequest

	// ID list of DDoS protected IPs.
	// The maximum number of DDoS protected IPs in each request is 100.
	DdosIpIds []string `json:"ddosIpIds,omitempty"`

	// DDoS protected IP pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	DdosIpChargeType string `json:"ddosIpChargeType,omitempty"`

	// IP address.
	IpAddress string `json:"ipAddress,omitempty"`

	// Zone ID to which the DDoS protected IPs belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Resource group ID.
	// If this parameter is not passed in, all the DDoS protected IPs in authorized resource group return.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Status of DDoS protected IP.
	DdosIpStatus string `json:"ddosIpStatus,omitempty"`

	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`

	// Instance name.
	InstanceName string `json:"instanceName,omitempty"`

	// Number of items in the current page result.
	// Default value: 20
	// Maximum value: 1000
	PageSize int `json:"pageSize,omitempty"`

	// Number of pages returned.
	// Default value: 1
	PageNum int `json:"pageNum,omitempty"`
}

type DdosIpAddress struct {
	// ID of the DDoS protected IP.
	DdosIpId string `json:"ddosIpId,omitempty"`

	// Zone ID to which the DDoS protected IPs belong.
	ZoneId string `json:"zoneId,omitempty"`

	// IP address.
	IpAddress string `json:"ipAddress,omitempty"`

	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`

	// Instance name.
	InstanceName string `json:"instanceName,omitempty"`

	// Pricing model.
	// PREPAID: monthly subscription.
	// POSTPAID: pay-as-you-go.
	DdosIpChargeType string `json:"ddosIpChargeType,omitempty"`

	// Period of DDoS protected IP subscription.
	// Unit: month.
	// The value is empty for pay-as-you-go DDoS protected IPs.
	Period *int `json:"period,omitempty"`

	// Creation time.
	// Use UTC time according to the ISO8601 standard. Format: YYYY-MM-DDThh:mm:ssZ.
	CreateTime string `json:"createTime,omitempty"`

	// Expiration time.
	// Use UTC time according to the ISO8601 standard. Format: YYYY-MM-DDThh:mm:ssZ.
	// Note:
	// The value is empty for pay-as-you-go resources.
	ExpiredTime string `json:"expiredTime,omitempty"`

	// Resource group ID.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Resource group name.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// Status of DDoS protected IP.
	DdosIpStatus string `json:"ddosIpStatus,omitempty"`
}

type TerminateDdosIpAddressRequest struct {
	*common.BaseRequest

	// ID of a DDoS protected IP.
	// To obtain the DDoS protected IP ID, you can call DescribeDdosIpAddresses and look for ddosIpId in the response.
	DdosIpId string `json:"ddosIpId,omitempty"`
}

type TerminateDdosIpAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseDdosIpAddressesRequest struct {
	*common.BaseRequest

	// IDs of DDoS protected IPs.
	// To obtain IDs of the DDoS protected IPs, you can call DescribeDdosIpAddresses and look for ddosIpId in the response.
	DdosIpIds []string `json:"ddosIpIds,omitempty"`
}

type ReleaseDdosIpAddressesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RenewDdosIpAddressRequest struct {
	*common.BaseRequest

	// ID of a DDoS protected IP.
	// To obtain the DDoS protected IP ID, you can call DescribeDdosIpAddresses and look for ddosIpId in the response.
	DdosIpId string `json:"ddosIpId,omitempty"`
}

type RenewDdosIpAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}
type AssociateDdosIpAddressRequest struct {
	*common.BaseRequest

	// ID of a DDoS protected IP.
	// To obtain the DDoS protected IP ID, you can call DescribeDdosIpAddresses and look for ddosIpId in the response.
	DdosIpId string `json:"ddosIpId,omitempty"`

	// Instance ID.
	// To obtain the instance ID, you can call DescribeInstances and look for instanceId in the response.
	InstanceId string `json:"instanceId,omitempty"`
}

type AssociateDdosIpAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type UnassociateDdosIpAddressRequest struct {
	*common.BaseRequest

	// ID of a DDoS protected IP.
	// To obtain the DDoS protected IP ID, you can call DescribeDdosIpAddresses and look for ddosIpId in the response.
	DdosIpId string `json:"ddosIpId,omitempty"`
}

type UnassociateDdosIpAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type InquiryPriceCreateDdosIpAddressRequest struct {
	*common.BaseRequest

	// Zone ID to which the DDoS protected IP belongs.
	ZoneId string `json:"zoneId,omitempty"`

	// DDoS protected IP pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	DdosIpChargeType string `json:"ddosIpChargeType,omitempty"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the ddosIpChargeType is PREPAID.
	DdosIpChargePrepaid *ChargePrepaid `json:"ddosIpChargePrepaid,omitempty"`

	// Quantity of DDoS protected IPs.
	// Value range: 1-100.
	// Default value: 1.
	Amount int `json:"amount,omitempty"`
}

type InquiryPriceCreateDdosIpAddressResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InquiryPriceCreateDdosIpAddressResponseParams `json:"response"`
}

type InquiryPriceCreateDdosIpAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Price of the DDoS protected IP.
	DdosIpPrice *Price `json:"ddosIpPrice,omitempty"`
}

type ModifyDdosIpAddressesResourceGroupRequest struct {
	*common.BaseRequest

	// List of DDoS IP IDs.
	// The maximum number of DDoS IP in each request is 100.
	DdosIpIds []string `json:"ddosIpIds,omitempty"`

	// Resource group ID.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
}

type ModifyDdosIpAddressesResourceGroupResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeVpcAvailableRegionsRequest struct {
	*common.BaseRequest

	// Zone ID to which the VPCs belong.
	ZoneId string `json:"zoneId,omitempty"`

	// IDs of availability regions for VPCs.
	VpcRegionId string `json:"vpcRegionId,omitempty"`
}

type DescribeVpcAvailableRegionsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeVpcAvailableRegionsResponseParams `json:"response,omitempty"`
}

type DescribeVpcAvailableRegionsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Information on availability regions for VPCs.
	VpcRegionSet []*VpcRegion `json:"vpcRegionSet,omitempty"`
}

type VpcRegion struct {
	// Availability region ID of VPC.
	VpcRegionId string `json:"vpcRegionId,omitempty"`

	// Availability region name of VPC.
	VpcRegionName string `json:"vpcRegionName,omitempty"`

	// List if zone IDs.
	ZoneIds []string `json:"zoneIds,omitempty"`
}

type CreateVpcRequest struct {
	*common.BaseRequest

	// CIDR block of VPC.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Availability region ID of VPC.
	VpcRegionId string `json:"vpcRegionId,omitempty"`

	// Resource group ID.
	// If the value is not passed in, the VPC will be put into the default resource group. If no authorized default resource group found, the request will fail.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// VPC name to be displayed.
	// This parameter can contain up to 64 characters. Only letters, numbers, - and periods (.) are supported.
	VpcName string `json:"vpcName,omitempty"`

	Tags *TagAssociation  `json:"tags,omitempty"`
}

type CreateVpcResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateVpcResponseParams `json:"response,omitempty"`
}

type CreateVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// VPC ID.
	VpcId string `json:"vpcId,omitempty"`
}

type ModifyVpcsAttributeRequest struct {
	*common.BaseRequest

	// VPC ID(s).
	// To obtain the VPC IDs, you can call DescribeVpcs and look for vpcId in the response.
	// The maximum number of VPCs in each request is 100.
	VpcIds []string `json:"vpcIds,omitempty"`

	// VPC name to be displayed.
	// This parameter can contain up to 64 characters. Only letters, numbers, - and periods (.) are supported.
	VpcName string `json:"vpcName,omitempty"`
}

type ModifyVpcsAttributeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyVpcsResourceGroupRequest struct {
	*common.BaseRequest
	VpcIds          []string `json:"vpcIds,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
}

type ModifyVpcsResourceGroupResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DeleteVpcRequest struct {
	*common.BaseRequest

	// VPC ID.
	VpcId string `json:"vpcId,omitempty"`
}

type DeleteVpcResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DescribeVpcsRequest struct {
	*common.BaseRequest

	// VPC ID(s).
	// The maximum number of VPCs in each request is 100.
	VpcIds []string `json:"vpcIds,omitempty"`

	// CIDR block of VPC.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Status of VPC.
	VpcStatus string `json:"vpcStatus,omitempty"`

	// Name of VPC.
	VpcName string `json:"vpcName,omitempty"`

	// Availability region ID of VPC.
	VpcRegionId string `json:"vpcRegionId,omitempty"`

	// Resource group ID.
	// If the value is null, then return all the VPCs in the authorized resource groups.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Number of items in the current page result.
	// Default value: 20;
	// Maximum value: 1000.
	PageSize int `json:"pageSize,omitempty"`

	// Number of pages returned.
	// Default value: 1.
	PageNum int `json:"pageNum,omitempty"`

	TagKeys  []string `json:"tagKeys,omitempty"`

    Tags   []*Tag `json:"tags,omitempty"`
}

type DescribeVpcsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeVpcsResponseParams `json:"response,omitempty"`
}

type DescribeVpcsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Number of VPCs meeting the filtering conditions.
	TotalCount int `json:"totalCount,omitempty"`

	// Information on VPCs.
	DataSet []*VpcInfo `json:"dataSet,omitempty"`
}

type VpcInfo struct {
	// VPC ID.
	VpcId string `json:"vpcId,omitempty"`

	// Availability region ID of VPC.
	VpcRegionId string `json:"vpcRegionId,omitempty"`

	// Availability region name of VPC.
	VpcRegionName string `json:"vpcRegionName,omitempty"`

	// VPC name.
	VpcName string `json:"vpcName,omitempty"`

	// CIDR block of VPC.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Resource group ID.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Resource group name.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// Creation time.
	// Format: YYYY-MM-DDThh:mm:ssZ.
	CreateTime string `json:"createTime,omitempty"`

	// Status of VPC.
	VpcStatus string `json:"vpcStatus,omitempty"`

	Tags  *Tags    `json:"tags,omitempty"`
}

type CreateSubnetRequest struct {
	*common.BaseRequest

	// CIDR block of the subnet. The optional values are as follows:
	// 10.0.0.0/16, 172.16.0.0/16, 192.168.0.0/16 and their subsets.
	// If the subnet is going to be added into a VPC, the value should be in the range of VPC CIDR block and not overlap with IP ranges of other subnets in this VPC.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Zone ID to which the subnet belongs.
	ZoneId string `json:"zoneId,omitempty"`

	// Resource group ID.
	// If the value is not passed in, the subnet will be put into the default resource group. If no authorized default resource group found, the request will fail.
	// If the subnet is going to be added into a VPC, this parameter will be ignored. The created subnet will be added into the resource group of the VPC.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Subnet name to be displayed.
	// This parameter can contain up to 64 characters. Only letters, numbers, - and periods (.) are supported.
	SubnetName string `json:"subnetName,omitempty"`

	// VPC ID.
	VpcId string `json:"vpcId,omitempty"`
}

type CreateSubnetResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateSubnetResponseParams `json:"response,omitempty"`
}

type CreateSubnetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Subnet的ID。
	SubnetId string `json:"subnetId,omitempty"`
}

type Subnet struct {
	// Subnet ID.
	SubnetId string `json:"subnetId,omitempty"`

	// Zone ID to which the subnet belongs.
	ZoneId string `json:"zoneId,omitempty"`

	// Subnet name.
	SubnetName string `json:"subnetName,omitempty"`

	// VPC ID.
	VpcId string `json:"vpcId,omitempty"`

	// VPC name.
	VpcName string `json:"vpcName,omitempty"`

	// CIDR block of subnet.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Creation time.
	// Use UTC time according to the ISO8601 standard. Format: YYYY-MM-DDThh:mm:ssZ.
	CreateTime string `json:"createTime,omitempty"`

	// Resource group ID.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Resource group name.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// Status of subnet.
	SubnetStatus string `json:"subnetStatus,omitempty"`

	// Information about instances in the subnet.
	SubnetInstanceSet []*SubnetAssociateInstance `json:"subnetInstanceSet,omitempty"`
}

type SubnetAssociateInstance struct {
	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`

	// Private IP address.
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`

	// Status of private IP address.
	PrivateIpStatus string `json:"privateIpStatus,omitempty"`
}

type DescribeSubnetsRequest struct {
	*common.BaseRequest

	// Subnet ID(s).
	// The maximum number of subnets in each request is 100.
	SubnetIds []string `json:"subnetIds,omitempty"`

	// CIDR block of subnet.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Zone ID to which subnets belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Name of subnet.
	SubnetName string `json:"subnetName,omitempty"`

	// VPC ID.
	VpcId string `json:"vpcId,omitempty"`

	// Resource group ID.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Status of subnet.
	SubnetStatus string `json:"subnetStatus,omitempty"`

	// Number of items in the current page result.
	// Default value: 20;
	// Maximum value: 1000.
	PageSize int `json:"pageSize,omitempty"`

	// Number of pages returned.
	// Default value: 1.
	PageNum int `json:"pageNum,omitempty"`
}

type DescribeSubnetsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeSubnetsResponseParams `json:"response,omitempty"`
}

type DescribeSubnetsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Information on subnets.
	DataSet []*Subnet `json:"dataSet,omitempty"`

	// Number of subnets meeting the filtering conditions.
	TotalCount int `json:"totalCount,omitempty"`
}

type DeleteSubnetRequest struct {
	*common.BaseRequest

	// Subnet ID.
	SubnetId string `json:"subnetId,omitempty"`
}

type DeleteSubnetResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type ModifySubnetsAttributeRequest struct {
	*common.BaseRequest

	// Subnet ID(s).
	// To obtain the subnet IDs, you can call DescribeSubnets and look for subnetId in the response.
	// The maximum number of subnets in each request is 100.
	SubnetIds []string `json:"subnetIds,omitempty"`

	// Subnet name to be displayed.
	// This parameter can contain up to 64 characters. Only letters, numbers, - and periods (.) are supported.
	SubnetName string `json:"subnetName,omitempty"`
}

type ModifySubnetsAttributeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifySubnetsResourceGroupRequest struct {
	*common.BaseRequest
	SubnetIds       []string `json:"subnetIds,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
}

type ModifySubnetsResourceGroupResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type UnAssociateSubnetInstanceRequest struct {
	*common.BaseRequest

	// Subnet ID.
	SubnetId string `json:"subnetId,omitempty"`

	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`
}

type UnAssociateSubnetInstanceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type AssociateSubnetInstancesRequest struct {
	*common.BaseRequest

	// Subnet ID.
	SubnetId string `json:"subnetId,omitempty"`

	// List of instances.
	SubnetInstanceList []*AssociateSubnetInstanceIpAddress `json:"subnetInstanceList,omitempty"`
}

type AssociateSubnetInstanceIpAddress struct {
	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`

	// Private IPv4 address.
	// The value should be in the range of subnet CIDR block. If the value is not passed in, any available private IP addresses will be assigned to the instance randomly.
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`
}

type AssociateSubnetInstancesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type AssociateVpcSubnetRequest struct {
	*common.BaseRequest

	// Subnet ID.
	SubnetId string `json:"subnetId,omitempty"`

	// VPC ID.
	VpcId string `json:"vpcId,omitempty"`
}

type AssociateVpcSubnetResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeSubnetAvailableResourcesRequest struct {
	*common.BaseRequest

	// Zone ID.
	// If the value is empty, return information about all the zones available for subnet.
	ZoneId string `json:"zoneId,omitempty"`
}

type DescribeSubnetAvailableResourcesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeSubnetAvailableResourcesResponseParams `json:"response,omitempty"`
}
type DescribeSubnetAvailableResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Information about zones available.
	ZoneIdSet []string `json:"zoneIdSet,omitempty"`
}

type DescribeCidrBlocksRequest struct {
	*common.BaseRequest

	// IDs of the CIDR blocks.
	// You can query up to 100 CIDR blocks in each request.
	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`

	// IP range of CIDR block.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Name of CIDR block.
	CidrBlockName string `json:"cidrBlockName,omitempty"`

	// Zone ID to which the CIDR blocks belong.
	ZoneId string `json:"zoneId,omitempty"`

	// Type of CIDR block. The optional values are as follows:
	// IPv4
	// IPv6
	CidrBlockType string `json:"cidrBlockType,omitempty"`

	// IP address of the gateway.
	Gateway string `json:"gateway,omitempty"`

	// CIDR block pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	ChargeType string `json:"chargeType,omitempty"`

	// Resource group ID.
	// If the value is null, then return all the CIDR blocks in the authorized resource groups.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Number of pages returned.
	// Default value: 1.
	PageNum int `json:"pageNum,omitempty"`

	// Number of items in the current page result.
	// Default value: 20;
	// Maximum value: 1000.
	PageSize int `json:"pageSize,omitempty"`

	TagKeys  []string `json:"tagKeys,omitempty"`

    Tags   []*Tag `json:"tags,omitempty"`
}

type DescribeCidrBlocksResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId"`

	Response *DescribeCidrBlocksResponseParams `json:"response"`
}

type DescribeCidrBlocksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Number of CIDR blocks meeting the filtering conditions.
	TotalCount int `json:"totalCount,omitempty"`

	// Information on a CIDR block.
	DataSet []*CidrBlockInfo `json:"dataSet,omitempty"`
}

type CidrBlockInfo struct {
	// ID of CIDR block.
	CidrBlockId string `json:"cidrBlockId,omitempty"`

	// Type of CIDR block. The optional values are as follows:
	// IPv4
	// IPv6
	CidrBlockType string `json:"cidrBlockType,omitempty"`

	// CIDR block name.
	CidrBlockName string `json:"cidrBlockName,omitempty"`

	// Zone ID to which the CIDR block belong.
	ZoneId string `json:"zoneId,omitempty"`

	// CIDR block.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Gateway address.
	Gateway string `json:"gateway,omitempty"`

	// Available IP starts.
	AvailableIpStart string `json:"availableIpStart,omitempty"`

	// Available IP ends.
	AvailableIpEnd string `json:"availableIpEnd,omitempty"`

	// Quantity of available IPs.
	AvailableIpCount int `json:"availableIpCount,omitempty"`

	// List of instance IDs.
	InstanceIds []string `json:"instanceIds,omitempty"`

	// Status of CIDR block.
	Status string `json:"status,omitempty"`

	// Pricing model.
	// PREPAID: monthly subscription.
	// POSTPAID: pay-as-you-go.
	ChargeType string `json:"chargeType,omitempty"`

	// Creation time.
	// Format: YYYY-MM-DDThh:mm:ssZ.
	CreateTime string `json:"createTime,omitempty"`

	// Expiration time.
	// Format: YYYY-MM-DDThh:mm:ssZ.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Resource group ID.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// Resource group name.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	Tags  *Tags    `json:"tags,omitempty"`
}

type DescribeCidrBlockIpsRequest struct {
	*common.BaseRequest

	// ID of the CIDR block.
	// You can find the cidrBlockId in the response by calling DescribeCidrBlocks.
	CidrBlockId string `json:"cidrBlockId,omitempty"`

	// ID of the instance.
	// You can find the instanceId in the response by calling DescribeInstances.
	InstanceId string `json:"instanceId,omitempty"`

	// IP address.
	Ip string `json:"ip,omitempty"`
}

type DescribeCidrBlockIpsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeCidrBlockIpsResponseParams `json:"response"`
}

type DescribeCidrBlockIpsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// IPs of the CIDR block.
	CidrBlockIps []*DescribeCidrBlockIpsResponseParam `json:"CidrBlockIps,omitempty"`
}

type DescribeCidrBlockIpsResponseParam struct {
	// ID of CIDR block.
	CidrBlockId string `json:"cidrBlockId,omitempty"`

	// Type of CIDR block. The optional values are as follows:
	// IPv4
	// IPv6
	CidrBlockType string `json:"cidrBlockType,omitempty"`

	// IP address.
	Ip string `json:"ip,omitempty"`

	// List of instance IDs.
	InstanceId *string `json:"instanceId,omitempty"`

	// Status of IPs.
	Status string `json:"status,omitempty"`
}

type DescribeAvailableIpv4ResourcesRequest struct {
	*common.BaseRequest

	// CIDR block pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	ChargeType string `json:"chargeType,omitempty"`

	// Zone ID to which the CIDR blocks belong.
	ZoneId string `json:"zoneId,omitempty"`
}

type DescribeAvailableIpv4ResourcesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAvailableIpv4ResourcesResponseParams `json:"response"`
}

type DescribeAvailableIpv4ResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// List of available IPv4 CIDR blocks.
	AvailableIpv4Resources []*DescribeAvailableIpv4ResourcesResponseParam `json:"availableIpv4Resources,omitempty"`
}

type DescribeAvailableIpv4ResourcesResponseParam struct {
	// Zone ID to which the CIDR block belongs.
	ZoneId string `json:"zoneId,omitempty"`

	// Netmask.
	Netmask int `json:"netmask,omitempty"`

	// Status for sale.
	// SELL: available for sale, stock > 10.
	// SELL_SHORTAGE: available for sale, stock < 10.
	// SOLD_OUT: sold out.
	SellStatus string `json:"sellStatus,omitempty"`
}

type DescribeAvailableIpv6ResourcesRequest struct {
	*common.BaseRequest

	// Zone ID to which the CIDR blocks belong.
	ZoneId string `json:"zoneId,omitempty"`
}

type DescribeAvailableIpv6ResourcesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAvailableIpv6ResourcesResponseParams `json:"response"`
}

type DescribeAvailableIpv6ResourcesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// List of available IPv6 CIDR blocks.
	AvailableIpv6Resources []*DescribeAvailableIpv6ResourcesResponseParam `json:"availableIpv6Resources,omitempty"`
}

type DescribeAvailableIpv6ResourcesResponseParam struct {
	// Zone ID to which the CIDR block belongs.
	ZoneId string `json:"zoneId,omitempty"`

	// Status for sale.
	// SELL: available for sale, stock > 10.
	// SELL_SHORTAGE: available for sale, stock < 10.
	// SOLD_OUT: sold out.
	SellStatus string `json:"sellStatus,omitempty"`
}

type DescribeInstanceAvailableCidrBlockRequest struct {
	*common.BaseRequest

	// ID of the instance.
	// You can find the instanceId in the response by calling DescribeInstances.
	InstanceId string `json:"instanceId,omitempty"`

	// Type of CIDR block. The optional values are as follows:
	// IPv4
	// IPv6
	CidrBlockType string `json:"cidrBlockType,omitempty"`
}

type DescribeInstanceAvailableCidrBlockResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstanceAvailableCidrBlockResponseParams `json:"response"`
}

type DescribeInstanceAvailableCidrBlockResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// List of available CIDR blocks for the instance.
	InstanceAvailableCidrBlocks []*DescribeInstanceAvailableCidrBlockResponseParam `json:"instanceAvailableCidrBlocks,omitempty"`
}

type DescribeInstanceAvailableCidrBlockResponseParam struct {
	// ID of CIDR block.
	CidrBlockId string `json:"cidrBlockId,omitempty"`

	// Zone ID to which the CIDR block belongs.
	ZoneId string `json:"zoneId,omitempty"`

	// Type of CIDR block. The optional values are as follows:
	// IPv4
	// IPv6
	CidrBlockType string `json:"cidrBlockType,omitempty"`

	// CIDR block.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// List of available IPs.
	AvailableIps []string `json:"availableIps,omitempty"`

	// Quantity of available IPs.
	AvailableIpCount int `json:"availableIpCount,omitempty"`
}

type InquiryPriceCreateIpv4BlockRequest struct {
	*common.BaseRequest

	// Zone ID to which the CIDR blocks belong.
	ZoneId string `json:"zoneId,omitempty"`

	// CIDR block pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	ChargeType string `json:"chargeType,omitempty"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the chargeType is PREPAID.
	ChargePrepaid *ChargePrepaid `json:"chargePrepaid,omitempty"`

	// Netmasks you want to purchase.
	// Value range: [1,32].
	// You can find the available netmasks in the response by calling DescribeAvailableIpv4Resource.
	Netmask int `json:"netmask,omitempty"`

	// Quantity of IPv4 CIDR blocks you want to purchase.
	// Default value: 1.
	Amount *int `json:"amount,omitempty"`

}

type InquiryPriceCreateIpv4BlockResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InquiryPriceCreateIpv4BlockResponseParam `json:"response"`
}

type InquiryPriceCreateIpv4BlockResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Price of IPv4 CIDR blocks.
	Price *Price `json:"price,omitempty"`
}

type CreateIpv4BlockRequest struct {
	*common.BaseRequest

	// Zone ID to which the CIDR blocks belong.
	ZoneId string `json:"zoneId,omitempty"`

	// CIDR block name to be displayed.
	// This parameter can contain up to 64 characters.
	Name string `json:"name,omitempty"`

	// CIDR block pricing model.
	// PREPAID: subscription
	// POSTPAID: pay-as-you-go
	ChargeType string `json:"chargeType,omitempty"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the chargeType is PREPAID.
	ChargePrepaid *ChargePrepaid `json:"chargePrepaid,omitempty"`

	// Netmasks you want to purchase.
	// Value range: [1,32].
	// You can find the available netmasks in the response by calling DescribeAvailableIpv4Resource.
	Netmask int `json:"netmask,omitempty"`

	// Quantity of IPv4 CIDR blocks you want to purchase.
	// Default value: 1.
	Amount *int `json:"amount,omitempty"`

	// Resource group ID where the CIDR blocks reside.
	// If an available VLAN exists in the specified zone, this parameter will be ignored. The created CIDR blocks will be added to the resource group of the VLAN.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	MarketingOptions  *MarketingInfo `json:"marketingOptions,omitempty"`

	Tags *TagAssociation  `json:"tags,omitempty"`
}

type CreateIpv4BlockResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateIpv4BlockResponseParam `json:"response"`
}

type CreateIpv4BlockResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Number of order.
	OrderNumber string `json:"orderNumber,omitempty"`

	// List of CIDR block IDs.
	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`
}

type CreateIpv6BlockRequest struct {
	*common.BaseRequest

	// Zone ID to which the CIDR blocks belong.
	ZoneId string `json:"zoneId,omitempty"`

	// CIDR block name to be displayed.
	// This parameter can contain up to 64 characters.
	Name string `json:"name,omitempty"`

	// Quantity of IPv6 CIDR blocks you want to purchase.
	// Default value: 1.
	Amount *int `json:"amount,omitempty"`

	// Resource group ID where the CIDR blocks reside.
	// If an available VLAN exists in the specified zone, this parameter will be ignored. The created CIDR blocks will be added to the resource group of the VLAN.
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	Tags *TagAssociation  `json:"tags,omitempty"`
}

type CreateIpv6BlockResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateIpv6BlockResponseParam `json:"response"`
}

type CreateIpv6BlockResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`
}

type ModifyCidrBlocksAttributeRequest struct {
	*common.BaseRequest

	// CIDR block ID(s).
	// To obtain the CIDR block IDs, you can call DescribeCidrBlocks and look for cidrBlockId in the response.
	// The maximum number of CIDR blocks in each request is 100.
	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`

	// CIDR block name to be displayed.
	// This parameter can contain up to 64 characters.
	Name string `json:"name,omitempty"`
}

type ModifyCidrBlocksAttributeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RenewCidrBlockRequest struct {
	*common.BaseRequest

	// ID of the CIDR block.
	// You can find the cidrBlockId in the response by calling DescribeCidrBlocks.
	CidrBlockId string `json:"cidrBlockId,omitempty"`
}

type RenewCidrBlockResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type TerminateCidrBlockRequest struct {
	*common.BaseRequest

	// ID of the CIDR block.
	// You can find the cidrBlockId in the response by calling DescribeCidrBlocks.
	CidrBlockId string `json:"cidrBlockId,omitempty"`
}

type TerminateCidrBlockResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseCidrBlocksRequest struct {
	*common.BaseRequest

	// CIDR block ID(s).
	// To obtain the CIDR block IDs, you can call DescribeCidrBlocks and look for cidrBlockId in the response.
	// The maximum number of CIDR blocks in each request is 100.
	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`
}

type ReleaseCidrBlocksResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type BindCidrBlockIpsRequest struct {
	*common.BaseRequest

	// ID of the CIDR block.
	// You can find the cidrBlockId in the response by calling DescribeCidrBlocks.
	CidrBlockId string `json:"cidrBlockId,omitempty"`

	// List of IPs to be assigned.
	IpBindList []*IpBindParam `json:"ipBindList,omitempty"`
}

type IpBindParam struct {
	// Instance ID.
	InstanceId string `json:"instanceId,omitempty"`

	// CIDR block IP.
	// To obtain the CIDR block IP, you can call DescribeInstanceAvailableCidrBlock and look for availableIps in the response.
	Ip string `json:"ip,omitempty"`
}

type BindCidrBlockIpsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type UnbindCidrBlockIpsRequest struct {
	*common.BaseRequest

	// ID of the CIDR block.
	// You can find the cidrBlockId in the response by calling DescribeCidrBlocks.
	CidrBlockId string `json:"cidrBlockId,omitempty"`

	// List of IPs to be unassigned.
	IpList []string `json:"ipList,omitempty"`
}

type UnbindCidrBlockIpsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}



// DescribeManagedInstancesRequest 实例列表请求参数
type DescribeManagedInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID。取值可以由多个实例ID组成一个。最多支持100个ID查询。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // InstanceName 实例名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // Ip 公网IP地址。
    Ip *string `json:"ip,omitempty"`

    // LanIp 内网IP地址。
    LanIp *string `json:"lanIp,omitempty"`

    // FacName 地域名称。
    FacName *string `json:"facName,omitempty"`

    // CityCode 城市代码。
    CityCode *string `json:"cityCode,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

// DescribeManagedInstancesResponseParams 托管实例查询结果。
type DescribeManagedInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 实例数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 实例列表。
    DataSet []*ManagedInstanceInfo `json:"dataSet,omitempty"`

}

// ManagedInstanceInfo 托管实例信息。
type ManagedInstanceInfo struct {

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 实例名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // FacName 地域名称。
    FacName *string `json:"facName,omitempty"`

    // Ips 公网IP列表。
    Ips []string `json:"ips,omitempty"`

    // LanIps 内网IP列表。
    LanIps []string `json:"lanIps,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

type DescribeManagedInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeManagedInstancesResponseParams `json:"response,omitempty"`
}



type DescribeManagedInstanceTrafficRequest struct {
	*common.BaseRequest

	InstanceId *string `json:"instanceId,omitempty"`

	StartTime *string `json:"startTime,omitempty"`

	EndTime *string `json:"endTime,omitempty"`
}

type DescribeManagedInstanceTrafficResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"requestId,omitempty"`

	Response *InstanceTrafficDataResponse `json:"response"`
}


type DescribeLoadBalancerZonesRequest struct {
	*common.BaseRequest

	// ChargeType
	ChargeType *string `json:"chargeType,omitempty"`

}

type DescribeLoadBalancerZonesResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	ZoneIdSet []string `json:"zoneIdSet,omitempty"`

}

type DescribeLoadBalancerZonesResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *DescribeLoadBalancerZonesResponseParams `json:"response,omitempty"`

}

type DescribeLoadBalancerSpecsRequest struct {
	*common.BaseRequest

	// ChargeType
	ChargeType *string `json:"chargeType,omitempty"`

	// ZoneId
	ZoneId *string `json:"zoneId,omitempty"`

}

type DescribeLoadBalancerSpecsResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	SpecSet []*LoadBalancerSpec `json:"specSet,omitempty"`

}

type LoadBalancerSpec struct {

	SpecName *string `json:"specName,omitempty"`

	MaxConnection *int `json:"maxConnection,omitempty"`

	Cps *int `json:"cps,omitempty"`

	Qps *int `json:"qps,omitempty"`

}

type DescribeLoadBalancerSpecsResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *DescribeLoadBalancerSpecsResponseParams `json:"response,omitempty"`

}

type DescribeLoadBalancersRequest struct {
	*common.BaseRequest

	// LoadBalancerIds
	LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`

	LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	PageNum *int `json:"pageNum,omitempty"`

	PageSize *int `json:"pageSize,omitempty"`

    TagKeys  []string `json:"tagKeys,omitempty"`

    Tags   []*Tag `json:"tags,omitempty"`
}

type DescribeLoadBalancersResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`

	DataSet []*LoadBalancerInfo `json:"dataSet,omitempty"`

}

type LoadBalancerInfo struct {

	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`

	LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	SpecName *string `json:"specName,omitempty"`

	VipList []*LoadBalancerIp `json:"vipList,omitempty"`

	ChargeType *string `json:"chargeType,omitempty"`

	Period *int `json:"period,omitempty"`

	CreateTime *string `json:"createTime,omitempty"`

	ExpiredTime *string `json:"expiredTime,omitempty"`

	Status *string `json:"status,omitempty"`

	ResourceGroupId *string `json:"resourceGroupId,omitempty"`

	ResourceGroupName *string `json:"resourceGroupName,omitempty"`

	MasterIp *string `json:"masterIp,omitempty"`

	BackupIp *string `json:"backupIp,omitempty"`

	IpType *string `json:"ipType,omitempty"`

	Bandwidth *int `json:"bandwidth,omitempty"`

	IsWorking *bool `json:"isWorking,omitempty"`

	ListenerList []*ListenerInfo `json:"listenerList,omitempty"`

	BackendList []*BackendInfo `json:"backendList,omitempty"`

    Tags  *Tags    `json:"tags,omitempty"`
}

type LoadBalancerIp struct {

	VipId *string `json:"vipId,omitempty"`

	IpAddress *string `json:"ipAddress,omitempty"`

	Type *string `json:"type,omitempty"`

	Status *string `json:"status,omitempty"`

}

type ListenerInfo struct {

	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

	ListenerId *string `json:"listenerId,omitempty"`

	ListenerName *string `json:"listenerName,omitempty"`

	Status *string `json:"status,omitempty"`

	Port *string `json:"port,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	BackendProtocol *string `json:"backendProtocol,omitempty"`

	Scheduler *string `json:"scheduler,omitempty"`

	Kind *string `json:"kind,omitempty"`

	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

	Notify *Notify `json:"notify,omitempty"`

	CreateTime *string `json:"createTime,omitempty"`

}

type HealthCheck struct {

	CheckEnabled *bool `json:"checkEnabled,omitempty"`

	CheckType *string `json:"checkType,omitempty"`

	CheckConnectTimeout *int `json:"checkConnectTimeout,omitempty"`

	CheckRetry *int `json:"checkRetry,omitempty"`

	CheckDelayBeforeRetry *int `json:"checkDelayBeforeRetry,omitempty"`

	CheckIntervalTime *int `json:"checkIntervalTime,omitempty"`

	CheckPort *int `json:"checkPort,omitempty"`

	HttpVersion *string `json:"httpVersion,omitempty"`

	HttpCheckPath *string `json:"httpCheckPath,omitempty"`

	HttpCheckDigest *int `json:"httpCheckDigest,omitempty"`

	HttpCode *int `json:"httpCode,omitempty"`

	MiscCheckPath *string `json:"miscCheckPath,omitempty"`

	MiscTimeout *int `json:"miscTimeout,omitempty"`

}

type Notify struct {

	Enable *bool `json:"enable,omitempty"`

	ApiAddress *string `json:"apiAddress,omitempty"`

}

type BackendInfo struct {

	ListenerId *string `json:"listenerId,omitempty"`

	BackendId *string `json:"backendId,omitempty"`

	BackendName *string `json:"backendName,omitempty"`

	Status *string `json:"status,omitempty"`

	Port *string `json:"port,omitempty"`

	Weight *int `json:"weight,omitempty"`

	InstanceId *string `json:"instanceId,omitempty"`

	InstanceName *string `json:"instanceName,omitempty"`

	CreateTime *string `json:"createTime,omitempty"`

}

type DescribeLoadBalancersResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *DescribeLoadBalancersResponseParams `json:"response,omitempty"`

}

type CreateLoadBalancerRequest struct {
	*common.BaseRequest

	ClientToken *string `json:"clientToken,omitempty"`

	// ZoneId
	ZoneId *string `json:"zoneId,omitempty"`

	// LoadBalancerName
	LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	// SpecName
	SpecName *string `json:"specName,omitempty"`

	// ChargeType
	ChargeType *string `json:"chargeType,omitempty"`

	// InstanceChargePrepaid
	InstanceChargePrepaid *InstanceChargePrepaid `json:"instanceChargePrepaid,omitempty"`

	// Bandwidth
	Bandwidth *int `json:"bandwidth,omitempty"`

	// IpType
	IpType *string `json:"ipType,omitempty"`

	VipCount *int `json:"vipCount,omitempty"`

	SubnetId *string `json:"subnetId,omitempty"`

	CidrBlockId *string `json:"cidrBlockId,omitempty"`

	MasterIp *string `json:"masterIp,omitempty"`

	BackupIp *string `json:"backupIp,omitempty"`

    Tags *TagAssociation  `json:"tags,omitempty"`
}


type InstanceChargePrepaid struct {

	// Period
	Period *int `json:"period,omitempty"`

	PeriodUnit *string `json:"periodUnit,omitempty"`

	NoConfirmPay *bool `json:"noConfirmPay,omitempty"`

}

type CreateLoadBalancerResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	OrderNumber *string `json:"orderNumber,omitempty"`

	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type CreateLoadBalancerResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *CreateLoadBalancerResponseParams `json:"response,omitempty"`

}

type ModifyLoadBalancersNameRequest struct {
	*common.BaseRequest

	// LoadBalancerName
	LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	// LoadBalancerIds
	LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

}

type ModifyLoadBalancersNameResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeListenersRequest struct {
	*common.BaseRequest

	// ListenerIds
	ListenerIds []string `json:"listenerIds,omitempty"`

	// LoadBalancerIds
	LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

	ListenerName *string `json:"listenerName,omitempty"`

	PageNum *int `json:"pageNum,omitempty"`

	PageSize *int `json:"pageSize,omitempty"`

}

type DescribeListenersResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`

	DataSet []*ListenerInfo `json:"dataSet,omitempty"`

}

type DescribeListenersResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *DescribeListenersResponseParams `json:"response,omitempty"`

}

type CreateListenerRequest struct {
	*common.BaseRequest

	// LoadBalancerId
	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

	// ListenerName
	ListenerName *string `json:"listenerName,omitempty"`

	// PortList
	PortList []int `json:"portList,omitempty"`

	// ClientToken
	ClientToken *string `json:"clientToken,omitempty"`

	// Protocol
	Protocol *string `json:"protocol,omitempty"`

	// BackendProtocol
	BackendProtocol *string `json:"backendProtocol,omitempty"`

	Scheduler *string `json:"scheduler,omitempty"`

	Kind *string `json:"kind,omitempty"`

	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

	Notify *Notify `json:"notify,omitempty"`

}

type CreateListenerResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	ListenerId *string `json:"listenerId,omitempty"`

}

type CreateListenerResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *CreateListenerResponseParams `json:"response,omitempty"`

}

type ModifyListenerAttributeRequest struct {
	*common.BaseRequest

	// ListenerId
	ListenerId *string `json:"listenerId,omitempty"`

	// ListenerName
	ListenerName *string `json:"listenerName,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	BackendProtocol *string `json:"backendProtocol,omitempty"`

	Scheduler *string `json:"scheduler,omitempty"`

	Kind *string `json:"kind,omitempty"`

	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

	Notify *Notify `json:"notify,omitempty"`

}

type ModifyListenerAttributeResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteListenerRequest struct {
	*common.BaseRequest

	// ListenerId
	ListenerId *string `json:"listenerId,omitempty"`

}

type DeleteListenerResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeBackendsRequest struct {
	*common.BaseRequest

	// ListenerId
	ListenerId *string `json:"listenerId,omitempty"`

	// BackendIds
	BackendIds []string `json:"backendIds,omitempty"`

	BackendName *string `json:"backendName,omitempty"`

	PageNum *int `json:"pageNum,omitempty"`

	PageSize *int `json:"pageSize,omitempty"`

}

type DescribeBackendsResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`

	DataSet []*BackendInfo `json:"dataSet,omitempty"`

}

type DescribeBackendsResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *DescribeBackendsResponseParams `json:"response,omitempty"`

}

type RegisterBackendRequest struct {
	*common.BaseRequest

	// ListenerId
	ListenerId *string `json:"listenerId,omitempty"`

	// BackendName
	BackendName *string `json:"backendName,omitempty"`

	// InstanceId
	InstanceId *string `json:"instanceId,omitempty"`

	// ClientToken
	ClientToken *string `json:"clientToken,omitempty"`

	PortList []int `json:"portList,omitempty"`

	// Weight
	Weight *int `json:"weight,omitempty"`

}

type RegisterBackendResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	BackendId *string `json:"backendId,omitempty"`

}

type RegisterBackendResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *RegisterBackendResponseParams `json:"response,omitempty"`

}

type DeregisterBackendRequest struct {
	*common.BaseRequest

	// BackendId
	BackendId *string `json:"backendId,omitempty"`

}

type DeregisterBackendResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyBackendWeightRequest struct {
	*common.BaseRequest

	// BackendId
	BackendId *string `json:"backendId,omitempty"`

	// Weight
	Weight *int `json:"weight,omitempty"`

}

type ModifyBackendWeightResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type CreateLoadBalancerVIPsRequest struct {
	*common.BaseRequest

	// LoadBalancerId
	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

	// Count
	Count *int `json:"count,omitempty"`

}

type CreateLoadBalancerVIPsResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	OrderNumber *string `json:"orderNumber,omitempty"`

	VipIdSet []string `json:"vipIdSet,omitempty"`

}

type CreateLoadBalancerVIPsResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *CreateLoadBalancerVIPsResponseParams `json:"response,omitempty"`

}

type DeleteLoadBalancerVIPRequest struct {
	*common.BaseRequest

	// VipId
	VipId *string `json:"vipId,omitempty"`

}

type DeleteLoadBalancerVIPResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyLoadBalancerBandwidthRequest struct {
	*common.BaseRequest

	// LoadBalancerId
	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

	// Bandwidth
	Bandwidth *int `json:"bandwidth,omitempty"`

}

type ModifyLoadBalancerBandwidthResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type TerminateLoadBalancerRequest struct {
	*common.BaseRequest

	// LoadBalancerId
	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type TerminateLoadBalancerResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ReleaseLoadBalancerRequest struct {
	*common.BaseRequest

	// LoadBalancerId
	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type ReleaseLoadBalancerResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RestoreLoadBalancerRequest struct {
	*common.BaseRequest

	// LoadBalancerId
	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type RestoreLoadBalancerResponseParams struct {

	RequestId *string `json:"requestId,omitempty"`

	OrderNumber *string `json:"orderNumber,omitempty"`

}

type RestoreLoadBalancerResponse struct {
	*common.BaseResponse

	RequestId *string `json:"requestId,omitempty"`

	Response *RestoreLoadBalancerResponseParams `json:"response,omitempty"`

}
