package bmc

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeZonesRequest struct {
	*common.BaseRequest
	AcceptLanguage string `json:"acceptLanguage,omitempty"`
}

type DescribeZonesResponse struct {
	*common.BaseResponse
	RequestId string                       `json:"requestId,omitempty"`
	Response  *DescribeZonesResponseParams `json:"response"`
}

type DescribeZonesResponseParams struct {
	RequestId string      `json:"requestId,omitempty"`
	ZoneSet   []*ZoneInfo `json:"zoneSet,omitempty"`
}

type ZoneInfo struct {
	ZoneId   string `json:"zoneId,omitempty"`
	ZoneName string `json:"zoneName,omitempty"`
	CityName string `json:"cityName,omitempty"`
	AreaName string `json:"areaName,omitempty"`
}

type CreateInstancesRequest struct {
	*common.BaseRequest
	ZoneId                  string         `json:"zoneId,omitempty"`
	InstanceChargeType      string         `json:"instanceChargeType,omitempty"`
	InstanceChargePrepaid   *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`
	InstanceTypeId          string         `json:"instanceTypeId,omitempty"`
	ImageId                 string         `json:"imageId,omitempty"`
	ResourceGroupId         string         `json:"resourceGroupId,omitempty"`
	InstanceName            string         `json:"instanceName,omitempty"`
	Hostname                string         `json:"hostname,omitempty"`
	Amount                  int            `json:"amount,omitempty"`
	Password                string         `json:"password,omitempty"`
	SshKeys                 []string       `json:"sshKeys,omitempty"`
	InternetChargeType      string         `json:"internetChargeType,omitempty"`
	InternetMaxBandwidthOut int            `json:"internetMaxBandwidthOut,omitempty"`
	TrafficPackageSize      float64        `json:"trafficPackageSize,omitempty"`
	SubnetId                string         `json:"subnetId,omitempty"`
	RaidConfig              *RaidConfig    `json:"raidConfig,omitempty"`
	Partitions              []*Partition   `json:"partitions,omitempty"`
	Nic                     *Nic           `json:"nic,omitempty"`
}

type RaidConfig struct {
	RaidType    *int          `json:"raidType,omitempty"`
	CustomRaids []*CustomRaid `json:"customRaids,omitempty"`
}

type CustomRaid struct {
	RaidType     *int  `json:"raidType,omitempty"`
	DiskSequence []int `json:"diskSequence,omitempty"`
}

type Partition struct {
	FsPath string `json:"fsPath,omitempty"`
	FsType string `json:"fsType,omitempty"`
	Size   int    `json:"size,omitempty"`
}

type Nic struct {
	WanName string `json:"wanName,omitempty"`
	LanName string `json:"lanName,omitempty"`
}

type CreateInstancesResponse struct {
	*common.BaseResponse
	RequestId string                        `json:"requestId,omitempty"`
	Response  *CreateInstanceResponseParams `json:"response"`
}

type CreateInstanceResponseParams struct {
	RequestId     *string   `json:"requestId,omitempty"`
	InstanceIdSet []*string `json:"instanceIdSet,omitempty"`
	OrderNumber   *string   `json:"orderNumber,omitempty"`
}

type ChargePrepaid struct {
	Period int `json:"period,omitempty"`
}

type DescribeInstanceTypesRequest struct {
	*common.BaseRequest
	ImageId             string   `json:"imageId,omitempty"`
	InstanceTypeIds     []string `json:"instanceTypeIds,omitempty"`
	MinimumCpuCoreCount *int     `json:"minimumCpuCoreCount,omitempty"`
	MaximumCpuCoreCount *int     `json:"maximumCpuCoreCount,omitempty"`
	MinimumMemorySize   *int     `json:"minimumMemorySize,omitempty"`
	MaximumMemorySize   *int     `json:"maximumMemorySize,omitempty"`
	MinimumBandwidth    *int     `json:"minimumBandwidth,omitempty"`
	SupportRaids        []int    `json:"supportRaids,omitempty"`
	SupportSubnet       *bool    `json:"supportSubnet,omitempty"`
	MinimumDiskSize     *int     `json:"minimumDiskSize,omitempty"`
	MaximumDiskSize     *int     `json:"maximumDiskSize,omitempty"`
	IsHA                *bool    `json:"isHA,omitempty"`
}

type DescribeInstanceTypesResponse struct {
	*common.BaseResponse
	RequestId string                               `json:"requestId,omitempty"`
	Response  *DescribeInstanceTypesResponseParams `json:"response"`
}

type DescribeInstanceTypesResponseParams struct {
	RequestId     string          `json:"requestId,omitempty"`
	InstanceTypes []*InstanceType `json:"instanceTypes,omitempty"`
}

type InstanceType struct {
	ImageIds         []string          `json:"imageIds,omitempty"`
	InstanceTypeId   string            `json:"instanceTypeId,omitempty"`
	Description      string            `json:"description,omitempty"`
	CpuCoreCount     int               `json:"cpuCoreCount,omitempty"`
	MemorySize       int               `json:"memorySize,omitempty"`
	MaximumBandwidth int               `json:"maximumBandwidth,omitempty"`
	SupportRaids     []int             `json:"supportRaids,omitempty"`
	SupportSubnet    bool              `json:"supportSubnet,omitempty"`
	IsHA             bool              `json:"isHA,omitempty"`
	DiskInfo         *InstanceDiskInfo `json:"diskInfo,omitempty"`
}

type InstanceDiskInfo struct {
	TotalDiskSize   int     `json:"totalDiskSize,omitempty"`
	DiskDescription string  `json:"diskDescription,omitempty"`
	Disks           []*Disk `json:"disks,omitempty"`
}

type Disk struct {
	DiskSize  int `json:"diskSize,omitempty"`
	DiskCount int `json:"diskCount,omitempty"`
}

type DescribeAvailableResourcesRequest struct {
	*common.BaseRequest
	InstanceTypeId     string `json:"instanceTypeId,omitempty"`
	InstanceChargeType string `json:"instanceChargeType,omitempty"`
	ZoneId             string `json:"zoneId,omitempty"`
	SellStatus         string `json:"sellStatus,omitempty"`
}

type DescribeAvailableResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                    `json:"requestId,omitempty"`
	Response  *DescribeAvailableResourcesResponseParams `json:"response"`
}

type DescribeAvailableResourcesResponseParams struct {
	RequestId          string               `json:"requestId,omitempty"`
	AvailableResources []*AvailableResource `json:"availableResources,omitempty"`
}
type AvailableResource struct {
	ZoneId                    string   `json:"zoneId,omitempty"`
	SellStatus                string   `json:"sellStatus,omitempty"`
	InternetChargeTypes       []string `json:"internetChargeTypes,omitempty"`
	InstanceTypeId            string   `json:"instanceTypeId,omitempty"`
	MaximumBandwidthOut       int      `json:"maximumBandwidthOut,omitempty"`
	DefaultBandwidthOut       int      `json:"defaultBandwidthOut,omitempty"`
	DefaultTrafficPackageSize float64  `json:"defaultTrafficPackageSize,omitempty"`
}

type DescribeImagesRequest struct {
	*common.BaseRequest
	ImageIds       []string `json:"imageIds,omitempty"`
	ImageName      string   `json:"imageName,omitempty"`
	Catalog        string   `json:"catalog,omitempty"`
	ImageType      string   `json:"imageType,omitempty"`
	OsType         string   `json:"osType,omitempty"`
	InstanceTypeId string   `json:"instanceTypeId,omitempty"`
}

type DescribeImagesResponse struct {
	*common.BaseResponse
	RequestId string                        `json:"requestId,omitempty"`
	Response  *DescribeImagesResponseParams `json:"response"`
}

type DescribeImagesResponseParams struct {
	RequestId string       `json:"requestId,omitempty"`
	Images    []*ImageInfo `json:"images,omitempty"`
}

type ImageInfo struct {
	ImageId   string `json:"imageId,omitempty"`
	ImageName string `json:"imageName,omitempty"`
	Catalog   string `json:"catalog,omitempty"`
	ImageType string `json:"imageType,omitempty"`
	OsType    string `json:"osType,omitempty"`
}

type DescribeInstancesRequest struct {
	*common.BaseRequest
	InstanceIds        []string `json:"instanceIds,omitempty"`
	ZoneId             string   `json:"zoneId,omitempty"`
	ResourceGroupId    string   `json:"resourceGroupId,omitempty"`
	InstanceTypeId     string   `json:"instanceTypeId,omitempty"`
	InternetChargeType string   `json:"internetChargeType,omitempty"`
	ImageId            string   `json:"imageId,omitempty"`
	SubnetId           string   `json:"subnetId,omitempty"`
	InstanceStatus     string   `json:"instanceStatus,omitempty"`
	InstanceName       string   `json:"instanceName,omitempty"`
	Hostname           string   `json:"hostname,omitempty"`
	PublicIpAddresses  []string `json:"publicIpAddresses,omitempty"`
	PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`
	PageNum            int      `json:"pageNum,omitempty"`
	PageSize           int      `json:"pageSize,omitempty"`
}

type InstanceInfo struct {
	InstanceId             string       `json:"instanceId,omitempty"`
	ZoneId                 string       `json:"zoneId,omitempty"`
	InstanceName           string       `json:"instanceName,omitempty"`
	Hostname               string       `json:"hostname,omitempty"`
	InstanceTypeId         string       `json:"instanceTypeId,omitempty"`
	ImageId                *string      `json:"imageId,omitempty"`
	ImageName              string       `json:"imageName,omitempty"`
	InstanceChargeType     string       `json:"instanceChargeType,omitempty"`
	BandwidthOutMbps       *int         `json:"bandwidthOutMbps,omitempty"`
	TrafficPackageSize     *float64     `json:"trafficPackageSize,omitempty"`
	InternetChargeType     string       `json:"internetChargeType,omitempty"`
	Period                 *int         `json:"period,omitempty"`
	PrimaryPublicIpAddress *string      `json:"primaryPublicIpAddress,omitempty"`
	PublicIpAddresses      []string     `json:"publicIpAddresses,omitempty"`
	PrivateIpAddresses     []string     `json:"privateIpAddresses,omitempty"`
	Ipv6Addresses          []string     `json:"ipv6Addresses,omitempty"`
	SubnetIds              []string     `json:"subnetIds,omitempty"`
	CreateTime             string       `json:"createTime,omitempty"`
	ExpiredTime            *string      `json:"expiredTime,omitempty"`
	ResourceGroupId        string       `json:"resourceGroupId,omitempty"`
	ResourceGroupName      string       `json:"resourceGroupName,omitempty"`
	InstanceStatus         string       `json:"instanceStatus,omitempty"`
	Partitions             []*Partition `json:"partitions,omitempty"`
	RaidConfig             *RaidConfig  `json:"raidConfig,omitempty"`
	Nic                    *Nic         `json:"nic,omitempty"`
}

type DescribeInstancesResponseParams struct {
	RequestId  string          `json:"requestId,omitempty"`
	TotalCount int             `json:"totalCount,omitempty"`
	DataSet    []*InstanceInfo `json:"dataSet,omitempty"`
}

type DescribeInstancesResponse struct {
	*common.BaseResponse
	RequestId string                           `json:"requestId,omitempty"`
	Response  *DescribeInstancesResponseParams `json:"response"`
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

type ReinstallInstanceRequest struct {
	*common.BaseRequest
	InstanceId string       `json:"instanceId,omitempty"`
	ImageId    string       `json:"imageId,omitempty"`
	Hostname   string       `json:"hostname,omitempty"`
	Password   string       `json:"password,omitempty"`
	SshKeys    []string     `json:"sshKeys,omitempty"`
	RaidConfig *RaidConfig  `json:"raidConfig,omitempty"`
	Partitions []*Partition `json:"partitions,omitempty"`
	Nic        *Nic         `json:"nic,omitempty"`
}

type ReInstallInstanceResponse struct {
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

type RenewInstanceRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
}

type RenewInstanceResponse struct {
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

type DescribeInstanceInternetStatusRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
}

type DescribeInstanceInternetStatusResponse struct {
	*common.BaseResponse
	RequestId string                  `json:"requestId,omitempty"`
	Response  *InstanceInternetStatus `json:"response"`
}

type InstanceInternetStatus struct {
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

type InquiryPriceCreateInstanceRequest struct {
	*common.BaseRequest
	ZoneId                  string         `json:"zoneId,omitempty"`
	InstanceTypeId          string         `json:"instanceTypeId,omitempty"`
	InstanceChargeType      string         `json:"instanceChargeType,omitempty"`
	InstanceChargePrepaid   *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`
	TrafficPackageSize      float64        `json:"trafficPackageSize,omitempty"`
	InternetMaxBandwidthOut int            `json:"internetMaxBandwidthOut,omitempty"`
	InternetChargeType      string         `json:"internetChargeType,omitempty"`
}

type InquiryPriceCreateInstanceResponse struct {
	*common.BaseResponse
	RequestId string                                    `json:"requestId,omitempty"`
	Response  *InquiryPriceCreateInstanceResponseParams `json:"response"`
}

type InquiryPriceCreateInstanceResponseParams struct {
	RequestId      string   `json:"requestId,omitempty"`
	InstancePrice  *Price   `json:"instancePrice,omitempty"`
	BandwidthPrice []*Price `json:"bandwidthPrice,omitempty"`
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

type InquiryPriceInstanceBandwidthRequest struct {
	*common.BaseRequest
	InstanceId       string `json:"instanceId,omitempty"`
	BandwidthOutMbps int    `json:"bandwidthOutMbps,omitempty"`
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
	InstanceId       string `json:"instanceId,omitempty"`
	BandwidthOutMbps *int   `json:"bandwidthOutMbps,omitempty"`
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

type ModifyInstanceTrafficPackageRequest struct {
	*common.BaseRequest
	InstanceId         string   `json:"instanceId,omitempty"`
	TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`
}

type ModifyInstanceTrafficPackageSizeResponseParams struct {
	RequestId   string `json:"requestId,omitempty"`
	OrderNumber string `json:"orderNumber,omitempty"`
}

type ModifyInstanceTrafficPackageResponse struct {
	*common.BaseResponse
	RequestId string                                          `json:"requestId,omitempty"`
	Response  *ModifyInstanceTrafficPackageSizeResponseParams `json:"response"`
}

type InquiryPriceInstanceTrafficPackageRequest struct {
	*common.BaseRequest
	InstanceId         string  `json:"instanceId,omitempty"`
	TrafficPackageSize float64 `json:"trafficPackageSize,omitempty"`
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
}

type DescribeEipAddressesRequest struct {
	*common.BaseRequest
	EipIds          []string `json:"eipIds,omitempty"`
	EipChargeType   string   `json:"eipChargeType,omitempty"`
	IpAddress       string   `json:"ipAddress,omitempty"`
	ZoneId          string   `json:"zoneId,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
	EipStatus       string   `json:"eipStatus,omitempty"`
	InstanceId      string   `json:"instanceId,omitempty"`
	InstanceName    string   `json:"instanceName,omitempty"`
	PageSize        int      `json:"pageSize,omitempty"`
	PageNum         int      `json:"pageNum,omitempty"`
}
type DescribeEipAddressesResponse struct {
	*common.BaseResponse
	RequestId string                              `json:"requestId,omitempty"`
	Response  *DescribeEipAddressesResponseParams `json:"response,omitempty"`
}
type DescribeEipAddressesResponseParams struct {
	RequestId  string        `json:"requestId,omitempty"`
	DataSet    []*EipAddress `json:"dataSet,omitempty"`
	TotalCount int           `json:"totalCount,omitempty"`
}

type DescribeEipAvailableResourcesRequest struct {
	*common.BaseRequest
	EipChargeType string `json:"eipChargeType,omitempty"`
	ZoneId        string `json:"zoneId,omitempty"`
}
type DescribeEipAvailableResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                       `json:"requestId,omitempty"`
	Response  *DescribeEipAvailableResourcesResponseParams `json:"response,omitempty"`
}
type DescribeEipAvailableResourcesResponseParams struct {
	RequestId    string                  `json:"requestId,omitempty"`
	EipResources []*EipAvailableResource `json:"eipResources,omitempty"`
}
type EipAvailableResource struct {
	ZoneId string `json:"zoneId,omitempty"`
	Status string `json:"status,omitempty"`
}

type DescribeInstanceAvailableEipResourcesRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
}
type DescribeInstanceAvailableEipResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                               `json:"requestId,omitempty"`
	Response  *DescribeInstanceAvailableEipResourcesResponseParams `json:"response,omitempty"`
}
type DescribeInstanceAvailableEipResourcesResponseParams struct {
	RequestId            string                          `json:"requestId,omitempty"`
	InstanceEipResources []*InstanceAvailableEipResource `json:"instanceEipResources,omitempty"`
}
type InstanceAvailableEipResource struct {
	EipId     string `json:"eipId,omitempty"`
	IpAddress string `json:"ipAddress,omitempty"`
}

type AllocateEipAddressesRequest struct {
	*common.BaseRequest
	EipChargeType    string         `json:"eipChargeType,omitempty"`
	EipChargePrepaid *ChargePrepaid `json:"eipChargePrepaid,omitempty"`
	ZoneId           string         `json:"zoneId,omitempty"`
	ResourceGroupId  string         `json:"resourceGroupId,omitempty"`
	Amount           int            `json:"amount,omitempty"`
}
type AllocateEipAddressesResponse struct {
	*common.BaseResponse
	RequestId string                              `json:"requestId,omitempty"`
	Response  *AllocateEipAddressesResponseParams `json:"response,omitempty"`
}

type AllocateEipAddressesResponseParams struct {
	RequestId   *string   `json:"requestId,omitempty"`
	EipIdSet    []*string `json:"eipIdSet,omitempty"`
	OrderNumber *string   `json:"orderNumber,omitempty"`
}

type TerminateEipAddressRequest struct {
	*common.BaseRequest
	EipId string `json:"eipId,omitempty"`
}

type TerminateEipAddressResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseEipAddressesRequest struct {
	*common.BaseRequest
	EipIds []string `json:"eipIds,omitempty"`
}

type ReleaseEipAddressesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RenewEipAddressRequest struct {
	*common.BaseRequest
	EipId string `json:"eipId,omitempty"`
}

type RenewEipAddressResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type AssociateEipAddressRequest struct {
	*common.BaseRequest
	EipId      string `json:"eipId,omitempty"`
	InstanceId string `json:"instanceId,omitempty"`
}

type AssociateEipAddressResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type UnassociateEipAddressRequest struct {
	*common.BaseRequest
	EipId string `json:"eipId,omitempty"`
}

type UnassociateEipAddressResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type InquiryPriceCreateEipAddressRequest struct {
	*common.BaseRequest
	ZoneId           string         `json:"zoneId,omitempty"`
	EipChargeType    string         `json:"eipChargeType,omitempty"`
	EipChargePrepaid *ChargePrepaid `json:"eipChargePrepaid,omitempty"`
	Amount           int            `json:"amount,omitempty"`
}

type InquiryPriceCreateEipAddressResponse struct {
	*common.BaseResponse
	RequestId string                                      `json:"requestId,omitempty"`
	Response  *InquiryPriceCreateEipAddressResponseParams `json:"response"`
}

type InquiryPriceCreateEipAddressResponseParams struct {
	RequestId string `json:"requestId,omitempty"`
	EipPrice  *Price `json:"eipPrice,omitempty"`
}

type ModifyEipAddressesResourceGroupRequest struct {
	*common.BaseRequest
	EipIds          []string `json:"eipIds,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
}

type ModifyEipAddressesResourceGroupResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeDdosIpAvailableResourcesRequest struct {
	*common.BaseRequest
	DdosIpChargeType string `json:"ddosIpChargeType,omitempty"`
	ZoneId           string `json:"zoneId,omitempty"`
}
type DescribeDdosIpAvailableResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                          `json:"requestId,omitempty"`
	Response  *DescribeDdosIpAvailableResourcesResponseParams `json:"response,omitempty"`
}
type DescribeDdosIpAvailableResourcesResponseParams struct {
	RequestId       string                     `json:"requestId,omitempty"`
	DdosIpResources []*DdosIpAvailableResource `json:"ddosIpResources,omitempty"`
}

type DdosIpAvailableResource struct {
	ZoneId string `json:"zoneId,omitempty"`
	Status string `json:"status,omitempty"`
}

type DescribeInstanceAvailableDdosResourcesRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceId,omitempty"`
}
type DescribeInstanceAvailableDdosResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                                `json:"requestId,omitempty"`
	Response  *DescribeInstanceAvailableDdosResourcesResponseParams `json:"response,omitempty"`
}
type DescribeInstanceAvailableDdosResourcesResponseParams struct {
	RequestId               string                             `json:"requestId,omitempty"`
	InstanceDdosIpResources []*InstanceAvailableDdosIpResource `json:"instanceDdosIpResources,omitempty"`
}
type InstanceAvailableDdosIpResource struct {
	DdosIpId  string `json:"ddosIpId,omitempty"`
	IpAddress string `json:"ipAddress,omitempty"`
}

type AllocateDdosIpAddressesRequest struct {
	*common.BaseRequest
	DdosIpChargeType    string         `json:"ddosIpChargeType,omitempty"`
	DdosIpChargePrepaid *ChargePrepaid `json:"ddosIpChargePrepaid,omitempty"`
	ZoneId              string         `json:"zoneId,omitempty"`
	ResourceGroupId     string         `json:"resourceGroupId,omitempty"`
	Amount              int            `json:"amount,omitempty"`
}
type AllocateDdosIpAddressesResponse struct {
	*common.BaseResponse
	RequestId string                                 `json:"requestId,omitempty"`
	Response  *AllocateDdosIpAddressesResponseParams `json:"response,omitempty"`
}

type AllocateDdosIpAddressesResponseParams struct {
	RequestId   *string   `json:"requestId,omitempty"`
	DdosIdSet   []*string `json:"ddosIdSet,omitempty"`
	OrderNumber *string   `json:"orderNumber,omitempty"`
}

type DescribeDdosIpAddressesResponse struct {
	*common.BaseResponse
	RequestId string                                 `json:"requestId,omitempty"`
	Response  *DescribeDdosIpAddressesResponseParams `json:"response,omitempty"`
}
type DescribeDdosIpAddressesResponseParams struct {
	RequestId  string           `json:"requestId,omitempty"`
	DataSet    []*DdosIpAddress `json:"dataSet,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
}

type DescribeDdosIpAddressesRequest struct {
	*common.BaseRequest
	DdosIpIds        []string `json:"ddosIpIds,omitempty"`
	DdosIpChargeType string   `json:"ddosIpChargeType,omitempty"`
	IpAddress        string   `json:"ipAddress,omitempty"`
	ZoneId           string   `json:"zoneId,omitempty"`
	ResourceGroupId  string   `json:"resourceGroupId,omitempty"`
	DdosIpStatus     string   `json:"ddosIpStatus,omitempty"`
	InstanceId       string   `json:"instanceId,omitempty"`
	InstanceName     string   `json:"instanceName,omitempty"`
	PageSize         int      `json:"pageSize,omitempty"`
	PageNum          int      `json:"pageNum,omitempty"`
}

type DdosIpAddress struct {
	DdosIpId          string `json:"ddosIpId,omitempty"`
	ZoneId            string `json:"zoneId,omitempty"`
	IpAddress         string `json:"ipAddress,omitempty"`
	InstanceId        string `json:"instanceId,omitempty"`
	InstanceName      string `json:"instanceName,omitempty"`
	DdosIpChargeType  string `json:"ddosIpChargeType,omitempty"`
	Period            *int   `json:"period,omitempty"`
	CreateTime        string `json:"createTime,omitempty"`
	ExpiredTime       string `json:"expiredTime,omitempty"`
	ResourceGroupId   string `json:"resourceGroupId,omitempty"`
	ResourceGroupName string `json:"resourceGroupName,omitempty"`
	DdosIpStatus      string `json:"ddosIpStatus,omitempty"`
}

type TerminateDdosIpAddressRequest struct {
	*common.BaseRequest
	DdosIpId string `json:"ddosIpId,omitempty"`
}

type TerminateDdosIpAddressResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseDdosIpAddressesRequest struct {
	*common.BaseRequest
	DdosIpIds []string `json:"ddosIpIds,omitempty"`
}

type ReleaseDdosIpAddressesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RenewDdosIpAddressRequest struct {
	*common.BaseRequest
	DdosIpId string `json:"ddosIpId,omitempty"`
}

type RenewDdosIpAddressResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}
type AssociateDdosIpAddressRequest struct {
	*common.BaseRequest
	DdosIpId   string `json:"ddosIpId,omitempty"`
	InstanceId string `json:"instanceId,omitempty"`
}

type AssociateDdosIpAddressResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type UnassociateDdosIpAddressRequest struct {
	*common.BaseRequest
	DdosIpId string `json:"ddosIpId,omitempty"`
}

type UnassociateDdosIpAddressResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type InquiryPriceCreateDdosIpAddressRequest struct {
	*common.BaseRequest
	ZoneId              string         `json:"zoneId,omitempty"`
	DdosIpChargeType    string         `json:"ddosIpChargeType,omitempty"`
	DdosIpChargePrepaid *ChargePrepaid `json:"ddosIpChargePrepaid,omitempty"`
	Amount              int            `json:"amount,omitempty"`
}

type InquiryPriceCreateDdosIpAddressResponse struct {
	*common.BaseResponse
	RequestId string                                         `json:"requestId,omitempty"`
	Response  *InquiryPriceCreateDdosIpAddressResponseParams `json:"response"`
}

type InquiryPriceCreateDdosIpAddressResponseParams struct {
	RequestId   string `json:"requestId,omitempty"`
	DdosIpPrice *Price `json:"ddosIpPrice,omitempty"`
}

type ModifyDdosIpAddressesResourceGroupRequest struct {
	*common.BaseRequest
	DdosIpIds       []string `json:"ddosIpIds,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
}

type ModifyDdosIpAddressesResourceGroupResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeVpcAvailableRegionsRequest struct {
	*common.BaseRequest
	ZoneId      string `json:"zoneId,omitempty"`
	VpcRegionId string `json:"vpcRegionId,omitempty"`
}

type DescribeVpcAvailableRegionsResponse struct {
	*common.BaseResponse
	RequestId string                                     `json:"requestId,omitempty"`
	Response  *DescribeVpcAvailableRegionsResponseParams `json:"response,omitempty"`
}

type DescribeVpcAvailableRegionsResponseParams struct {
	RequestId    string       `json:"requestId,omitempty"`
	VpcRegionSet []*VpcRegion `json:"vpcRegionSet,omitempty"`
}

type VpcRegion struct {
	VpcRegionId   string   `json:"vpcRegionId,omitempty"`
	VpcRegionName string   `json:"vpcRegionName,omitempty"`
	ZoneIds       []string `json:"zoneIds,omitempty"`
}

type CreateVpcRequest struct {
	*common.BaseRequest
	CidrBlock       string `json:"cidrBlock,omitempty"`
	VpcRegionId     string `json:"vpcRegionId,omitempty"`
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
	VpcName         string `json:"vpcName,omitempty"`
}

type CreateVpcResponse struct {
	*common.BaseResponse
	RequestId string                   `json:"requestId,omitempty"`
	Response  *CreateVpcResponseParams `json:"response,omitempty"`
}

type CreateVpcResponseParams struct {
	RequestId string `json:"requestId,omitempty"`
	VpcId     string `json:"vpcId,omitempty"`
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
	VpcId string `json:"vpcId,omitempty"`
}

type DeleteVpcResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DescribeVpcsRequest struct {
	*common.BaseRequest
	VpcIds          []string `json:"vpcIds,omitempty"`
	CidrBlock       string   `json:"cidrBlock,omitempty"`
	VpcStatus       string   `json:"vpcStatus,omitempty"`
	VpcName         string   `json:"vpcName,omitempty"`
	VpcRegionId     string   `json:"vpcRegionId,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
	PageSize        int      `json:"pageSize,omitempty"`
	PageNum         int      `json:"pageNum,omitempty"`
}

type DescribeVpcsResponse struct {
	*common.BaseResponse
	RequestId string                      `json:"requestId,omitempty"`
	Response  *DescribeVpcsResponseParams `json:"response,omitempty"`
}

type DescribeVpcsResponseParams struct {
	RequestId  string     `json:"requestId,omitempty"`
	TotalCount int        `json:"totalCount,omitempty"`
	DataSet    []*VpcInfo `json:"dataSet,omitempty"`
}

type VpcInfo struct {
	VpcId             string `json:"vpcId,omitempty"`
	VpcRegionId       string `json:"vpcRegionId,omitempty"`
	VpcRegionName     string `json:"vpcRegionName,omitempty"`
	VpcName           string `json:"vpcName,omitempty"`
	CidrBlock         string `json:"cidrBlock,omitempty"`
	ResourceGroupId   string `json:"resourceGroupId,omitempty"`
	ResourceGroupName string `json:"resourceGroupName,omitempty"`
	CreateTime        string `json:"createTime,omitempty"`
	VpcStatus         string `json:"vpcStatus,omitempty"`
}

type CreateSubnetRequest struct {
	*common.BaseRequest
	CidrBlock       string `json:"cidrBlock,omitempty"`
	ZoneId          string `json:"zoneId,omitempty"`
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
	SubnetName      string `json:"subnetName,omitempty"`
	VpcId           string `json:"vpcId,omitempty"`
}

type CreateSubnetResponse struct {
	*common.BaseResponse
	RequestId string                      `json:"requestId,omitempty"`
	Response  *CreateSubnetResponseParams `json:"response,omitempty"`
}

type CreateSubnetResponseParams struct {
	RequestId string `json:"requestId,omitempty"`
	SubnetId  string `json:"subnetId,omitempty"`
}

type Subnet struct {
	SubnetId          string                     `json:"subnetId,omitempty"`
	ZoneId            string                     `json:"zoneId,omitempty"`
	SubnetName        string                     `json:"subnetName,omitempty"`
	VpcId             string                     `json:"vpcId,omitempty"`
	VpcName           string                     `json:"vpcName,omitempty"`
	CidrBlock         string                     `json:"cidrBlock,omitempty"`
	CreateTime        string                     `json:"createTime,omitempty"`
	ResourceGroupId   string                     `json:"resourceGroupId,omitempty"`
	ResourceGroupName string                     `json:"resourceGroupName,omitempty"`
	SubnetStatus      string                     `json:"subnetStatus,omitempty"`
	SubnetInstanceSet []*SubnetAssociateInstance `json:"subnetInstanceSet,omitempty"`
}

type SubnetAssociateInstance struct {
	InstanceId       string `json:"instanceId,omitempty"`
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`
	PrivateIpStatus  string `json:"privateIpStatus,omitempty"`
}

type DescribeSubnetsRequest struct {
	*common.BaseRequest
	SubnetIds       []string `json:"subnetIds,omitempty"`
	CidrBlock       string   `json:"cidrBlock,omitempty"`
	ZoneId          string   `json:"zoneId,omitempty"`
	SubnetName      string   `json:"subnetName,omitempty"`
	VpcId           string   `json:"vpcId,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
	SubnetStatus    string   `json:"subnetStatus,omitempty"`
	PageSize        int      `json:"pageSize,omitempty"`
	PageNum         int      `json:"pageNum,omitempty"`
}

type DescribeSubnetsResponse struct {
	*common.BaseResponse
	RequestId string                         `json:"requestId,omitempty"`
	Response  *DescribeSubnetsResponseParams `json:"response,omitempty"`
}

type DescribeSubnetsResponseParams struct {
	RequestId  string    `json:"requestId,omitempty"`
	DataSet    []*Subnet `json:"dataSet,omitempty"`
	TotalCount int       `json:"totalCount,omitempty"`
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
	} `json:"response,omitempty"`
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
	SubnetId   string `json:"subnetId,omitempty"`
	InstanceId string `json:"instanceId,omitempty"`
}

type UnAssociateSubnetInstanceResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type AssociateSubnetInstancesRequest struct {
	*common.BaseRequest
	SubnetId           string                              `json:"subnetId,omitempty"`
	SubnetInstanceList []*AssociateSubnetInstanceIpAddress `json:"subnetInstanceList,omitempty"`
}

type AssociateSubnetInstanceIpAddress struct {
	InstanceId       string `json:"instanceId,omitempty"`
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`
}

type AssociateSubnetInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type AssociateVpcSubnetRequest struct {
	*common.BaseRequest
	SubnetId string `json:"subnetId,omitempty"`
	VpcId    string `json:"vpcId,omitempty"`
}

type AssociateVpcSubnetResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeSubnetAvailableResourcesRequest struct {
	*common.BaseRequest
	ZoneId string `json:"zoneId,omitempty"`
}

type DescribeSubnetAvailableResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                          `json:"requestId,omitempty"`
	Response  *DescribeSubnetAvailableResourcesResponseParams `json:"response,omitempty"`
}
type DescribeSubnetAvailableResourcesResponseParams struct {
	RequestId string   `json:"requestId,omitempty"`
	ZoneIdSet []string `json:"zoneIdSet,omitempty"`
}

type DescribeCidrBlocksRequest struct {
	*common.BaseRequest
	CidrBlockIds    []string `json:"cidrBlockIds,omitempty"`
	CidrBlock       string   `json:"cidrBlock,omitempty"`
	CidrBlockName   string   `json:"cidrBlockName,omitempty"`
	ZoneId          string   `json:"zoneId,omitempty"`
	CidrBlockType   string   `json:"cidrBlockType,omitempty"`
	Gateway         string   `json:"gateway,omitempty"`
	ChargeType      string   `json:"chargeType,omitempty"`
	ResourceGroupId string   `json:"resourceGroupId,omitempty"`
	PageNum         int      `json:"pageNum,omitempty"`
	PageSize        int      `json:"pageSize,omitempty"`
}

type DescribeCidrBlocksResponse struct {
	*common.BaseResponse
	RequestId string                            `json:"requestId"`
	Response  *DescribeCidrBlocksResponseParams `json:"response"`
}

type DescribeCidrBlocksResponseParams struct {
	RequestId  string           `json:"requestId,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
	DataSet    []*CidrBlockInfo `json:"dataSet,omitempty"`
}

type CidrBlockInfo struct {
	CidrBlockId       string   `json:"cidrBlockId,omitempty"`
	CidrBlockType     string   `json:"cidrBlockType,omitempty"`
	CidrBlockName     string   `json:"cidrBlockName,omitempty"`
	ZoneId            string   `json:"zoneId,omitempty"`
	CidrBlock         string   `json:"cidrBlock,omitempty"`
	Gateway           string   `json:"gateway,omitempty"`
	AvailableIpStart  string   `json:"availableIpStart,omitempty"`
	AvailableIpEnd    string   `json:"availableIpEnd,omitempty"`
	AvailableIpCount  int      `json:"availableIpCount,omitempty"`
	InstanceIds       []string `json:"instanceIds,omitempty"`
	Status            string   `json:"status,omitempty"`
	ChargeType        string   `json:"chargeType,omitempty"`
	CreateTime        string   `json:"createTime,omitempty"`
	ExpireTime        *string  `json:"expireTime,omitempty"`
	ResourceGroupId   string   `json:"resourceGroupId,omitempty"`
	ResourceGroupName string   `json:"resourceGroupName,omitempty"`
}

type DescribeCidrBlockIpsRequest struct {
	*common.BaseRequest
	CidrBlockId string `json:"cidrBlockId,omitempty"`
	InstanceId  string `json:"instanceId,omitempty"`
	Ip          string `json:"ip,omitempty"`
}

type DescribeCidrBlockIpsResponse struct {
	*common.BaseResponse
	RequestId string                              `json:"requestId,omitempty"`
	Response  *DescribeCidrBlockIpsResponseParams `json:"response"`
}

type DescribeCidrBlockIpsResponseParams struct {
	RequestId    string                               `json:"requestId,omitempty"`
	CidrBlockIps []*DescribeCidrBlockIpsResponseParam `json:"CidrBlockIps,omitempty"`
}

type DescribeCidrBlockIpsResponseParam struct {
	CidrBlockId   string  `json:"cidrBlockId,omitempty"`
	CidrBlockType string  `json:"cidrBlockType,omitempty"`
	Ip            string  `json:"ip,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty"`
	Status        string  `json:"status,omitempty"`
}

type DescribeAvailableIpv4ResourcesRequest struct {
	*common.BaseRequest
	ChargeType string `json:"chargeType,omitempty"`
	ZoneId     string `json:"zoneId,omitempty"`
}

type DescribeAvailableIpv4ResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                        `json:"requestId,omitempty"`
	Response  *DescribeAvailableIpv4ResourcesResponseParams `json:"response"`
}

type DescribeAvailableIpv4ResourcesResponseParams struct {
	RequestId              string                                         `json:"requestId,omitempty"`
	AvailableIpv4Resources []*DescribeAvailableIpv4ResourcesResponseParam `json:"availableIpv4Resources,omitempty"`
}

type DescribeAvailableIpv4ResourcesResponseParam struct {
	ZoneId     string `json:"zoneId,omitempty"`
	Netmask    int    `json:"netmask,omitempty"`
	SellStatus string `json:"sellStatus,omitempty"`
}

type DescribeAvailableIpv6ResourcesRequest struct {
	*common.BaseRequest
	ZoneId string `json:"zoneId,omitempty"`
}

type DescribeAvailableIpv6ResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                        `json:"requestId,omitempty"`
	Response  *DescribeAvailableIpv6ResourcesResponseParams `json:"response"`
}

type DescribeAvailableIpv6ResourcesResponseParams struct {
	RequestId              string                                         `json:"requestId,omitempty"`
	AvailableIpv6Resources []*DescribeAvailableIpv6ResourcesResponseParam `json:"availableIpv6Resources,omitempty"`
}

type DescribeAvailableIpv6ResourcesResponseParam struct {
	ZoneId     string `json:"zoneId,omitempty"`
	SellStatus string `json:"sellStatus,omitempty"`
}

type DescribeInstanceAvailableCidrBlockRequest struct {
	*common.BaseRequest
	InstanceId    string `json:"instanceId,omitempty"`
	CidrBlockType string `json:"cidrBlockType,omitempty"`
}

type DescribeInstanceAvailableCidrBlockResponse struct {
	*common.BaseResponse
	RequestId string                                            `json:"requestId,omitempty"`
	Response  *DescribeInstanceAvailableCidrBlockResponseParams `json:"response"`
}

type DescribeInstanceAvailableCidrBlockResponseParams struct {
	RequestId                   string                                             `json:"requestId,omitempty"`
	InstanceAvailableCidrBlocks []*DescribeInstanceAvailableCidrBlockResponseParam `json:"instanceAvailableCidrBlocks,omitempty"`
}

type DescribeInstanceAvailableCidrBlockResponseParam struct {
	CidrBlockId      string   `json:"cidrBlockId,omitempty"`
	ZoneId           string   `json:"zoneId,omitempty"`
	CidrBlockType    string   `json:"cidrBlockType,omitempty"`
	CidrBlock        string   `json:"cidrBlock,omitempty"`
	AvailableIps     []string `json:"availableIps,omitempty"`
	AvailableIpCount int      `json:"availableIpCount,omitempty"`
}

type InquiryPriceCreateIpv4BlockRequest struct {
	*common.BaseRequest
	ZoneId        string         `json:"zoneId,omitempty"`
	ChargeType    string         `json:"chargeType,omitempty"`
	ChargePrepaid *ChargePrepaid `json:"chargePrepaid,omitempty"`
	Netmask       int            `json:"netmask,omitempty"`
	Amount        *int           `json:"amount,omitempty"`
}

type InquiryPriceCreateIpv4BlockResponse struct {
	*common.BaseResponse
	RequestId string                                    `json:"requestId,omitempty"`
	Response  *InquiryPriceCreateIpv4BlockResponseParam `json:"response"`
}

type InquiryPriceCreateIpv4BlockResponseParam struct {
	RequestId string `json:"requestId,omitempty"`
	Price     *Price `json:"price,omitempty"`
}

type CreateIpv4BlockRequest struct {
	*common.BaseRequest
	ZoneId          string         `json:"zoneId,omitempty"`
	Name            string         `json:"name,omitempty"`
	ChargeType      string         `json:"chargeType,omitempty"`
	ChargePrepaid   *ChargePrepaid `json:"chargePrepaid,omitempty"`
	Netmask         int            `json:"netmask,omitempty"`
	Amount          *int           `json:"amount,omitempty"`
	ResourceGroupId string         `json:"resourceGroupId,omitempty"`
}

type CreateIpv4BlockResponse struct {
	*common.BaseResponse
	RequestId string                        `json:"requestId,omitempty"`
	Response  *CreateIpv4BlockResponseParam `json:"response"`
}

type CreateIpv4BlockResponseParam struct {
	RequestId    string   `json:"requestId,omitempty"`
	OrderNumber  string   `json:"orderNumber,omitempty"`
	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`
}

type CreateIpv6BlockRequest struct {
	*common.BaseRequest
	ZoneId          string `json:"zoneId,omitempty"`
	Name            string `json:"name,omitempty"`
	Amount          *int   `json:"amount,omitempty"`
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
}

type CreateIpv6BlockResponse struct {
	*common.BaseResponse
	RequestId string                        `json:"requestId,omitempty"`
	Response  *CreateIpv6BlockResponseParam `json:"response"`
}

type CreateIpv6BlockResponseParam struct {
	RequestId    string   `json:"requestId,omitempty"`
	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`
}

type ModifyCidrBlocksAttributeRequest struct {
	*common.BaseRequest
	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`
	Name         string   `json:"name,omitempty"`
}

type ModifyCidrBlocksAttributeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RenewCidrBlockRequest struct {
	*common.BaseRequest
	CidrBlockId string `json:"cidrBlockId,omitempty"`
}

type RenewCidrBlockResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type TerminateCidrBlockRequest struct {
	*common.BaseRequest
	CidrBlockId string `json:"cidrBlockId,omitempty"`
}

type TerminateCidrBlockResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ReleaseCidrBlocksRequest struct {
	*common.BaseRequest
	CidrBlockIds []string `json:"cidrBlockIds,omitempty"`
}

type ReleaseCidrBlocksResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type BindCidrBlockIpsRequest struct {
	*common.BaseRequest
	CidrBlockId string         `json:"cidrBlockId,omitempty"`
	IpBindList  []*IpBindParam `json:"ipBindList,omitempty"`
}

type IpBindParam struct {
	InstanceId string `json:"instanceId,omitempty"`
	Ip         string `json:"ip,omitempty"`
}

type BindCidrBlockIpsResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type UnbindCidrBlockIpsRequest struct {
	*common.BaseRequest
	CidrBlockId string   `json:"cidrBlockId,omitempty"`
	IpList      []string `json:"ipList,omitempty"`
}

type UnbindCidrBlockIpsResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId,omitempty"`
	Response  struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}
