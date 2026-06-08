package vm

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


type DescribeZonesRequest struct {
    *common.BaseRequest

    // ZoneIds 可用区ID集合。
    ZoneIds []string `json:"zoneIds,omitempty"`

}

type DescribeZonesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeZonesResponseParams `json:"response,omitempty"`

}

type DescribeZonesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ZoneSet 区域信息集合。
    ZoneSet []*ZoneInfo `json:"zoneSet,omitempty"`

}

// ZoneInfo 可用区的基本信息。
type ZoneInfo struct {

    // ZoneId 区域ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ZoneName 区域名称。
    ZoneName *string `json:"zoneName,omitempty"`

    // SupportSecurityGroup 是否支持安全组。
    SupportSecurityGroup *bool `json:"supportSecurityGroup,omitempty"`

    // SupportNetworkType 支持的网络类型。
    // CLASSICS：经典网络。
    // VPC：VPC网络。
    SupportNetworkType *string `json:"supportNetworkType,omitempty"`

    // SupportIpv6 是否支持公网IPv6。
    SupportIpv6 *bool `json:"supportIpv6,omitempty"`

    // SupportCpuPassThrough 是否支持CPU透传。
    SupportCpuPassThrough *bool `json:"supportCpuPassThrough,omitempty"`

    // NetworkLineType 网络线路信息。
    NetworkLineType *string `json:"networkLineType,omitempty"`

}

type InquiryPriceCreateInstanceRequest struct {
    *common.BaseRequest

    // ZoneId 实例所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ImageId 指定有效的镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // InstanceType 实例机型。
    InstanceType *string `json:"instanceType,omitempty"`

    // InstanceChargeType 实例计费类型。
    // PREPAID：预付费。
    // POSTPAID：后付费。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // InstanceChargePrepaid 预付费模式，即包年包月相关参数设置。
    // 若指定实例的付费模式为预付费则该参数必传。
    InstanceChargePrepaid *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`

    // InstanceChargePostpaid 后付费模式相关参数设置。
    InstanceChargePostpaid *ChargePostpaid `json:"instanceChargePostpaid,omitempty"`

    // InternetChargeType 网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // TrafficPackageSize 流量包订购大小，单位TB。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // InternetMaxBandwidthOut 公网出带宽上限，单位Mbps。
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

    // InstanceCount 指定创建实例的数量。
    InstanceCount *int `json:"instanceCount,omitempty"`

    // SystemDisk 系统盘配置。
    SystemDisk *SystemDisk `json:"systemDisk,omitempty"`

    // DataDisks 数据盘配置。
    DataDisks []*DataDisk `json:"dataDisks,omitempty"`

}

// ChargePrepaid 预付费模式，即包年包月相关参数设置。
type ChargePrepaid struct {

    // Period 购买实例的时长。
    // 单位：月。
    Period *int `json:"period,omitempty"`

}

// ChargePostpaid 后付费模式，即按量付费相关参数设置。
type ChargePostpaid struct {

    // Period 后付费时长。
    // 单位：月。
    Period *int `json:"period,omitempty"`

}

// SystemDisk 描述系统盘的基本信息。
type SystemDisk struct {

    // DiskId 磁盘ID。
    DiskId *string `json:"diskId,omitempty"`

    // DiskSize 系统盘大小，单位GB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskCategory 磁盘种类。
    // STANDARD：标准云盘。
    // SSD：固态硬盘。
    // 默认为SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

}

// DataDisk 数据盘的基本信息。
type DataDisk struct {

    // DiskId 磁盘ID。
    DiskId *string `json:"diskId,omitempty"`

    // DiskName 磁盘名称。
    DiskName *string `json:"diskName,omitempty"`

    // DiskSize 数据盘大小，单位GB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskAmount 数据盘数量。
    DiskAmount *int `json:"diskAmount,omitempty"`

    // Portable 是否可拔插。
    Portable *bool `json:"portable,omitempty"`

    // DiskCategory 磁盘种类。
    // STANDARD：标准云盘。
    // SSD：固态硬盘。
    // 默认为SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

    // DiskPrice 数据盘价格。
    DiskPrice *PriceItem `json:"diskPrice,omitempty"`

}

// PriceItem 描述价格的信息。
type PriceItem struct {

    // Discount 折扣大小。
    // 如80.0代表8折。
    Discount *float64 `json:"discount,omitempty"`

    // DiscountPrice 后付费的单元折后价格。
    // 后付费模式使用，如果价格为阶梯价格，该项为null。
    DiscountPrice *float64 `json:"discountPrice,omitempty"`

    // OriginalPrice 预付费的原价。
    // 预付费模式使用，后付费该值为 null。
    OriginalPrice *float64 `json:"originalPrice,omitempty"`

    // UnitPrice 后付费的单元原始价格。
    // 后付费模式使用，如果价格为阶梯价格，该项为null。
    UnitPrice *float64 `json:"unitPrice,omitempty"`

    // DiscountUnitPrice 后付费的单元折后价格。
    // 后付费模式使用，如果价格为阶梯价格，该项为null。
    DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`

    // ChargeUnit 后付费计价单元。
    // 后付费模式使用，可取值范围：<br/>HOUR: 表示计价单元是按每小时来计算。
    // DAY: 表示计价单元是按天来计算。
    // MONTH: 表示计价单元是按月来计算，95计费则是这种。
    ChargeUnit *string `json:"chargeUnit,omitempty"`

    // StepPrices 后付费阶梯价格。
    // 后付费模式使用，如果非阶梯价格，该项为null。
    StepPrices []*StepPrice `json:"stepPrices,omitempty"`

    // AmountUnit 用量单位。
    // 比如Mbps, LCU等。
    // 如果为null, 代表取不到值。
    AmountUnit *string `json:"amountUnit,omitempty"`

    // ExcessUnitPrice 超量原始价格。
    ExcessUnitPrice *float64 `json:"excessUnitPrice,omitempty"`

    // ExcessDiscountUnitPrice 超量折扣后价格。
    ExcessDiscountUnitPrice *float64 `json:"excessDiscountUnitPrice,omitempty"`

    // ExcessAmountUnit 超量用量单位。
    // 如果为null, 代表取不到值。
    ExcessAmountUnit *string `json:"excessAmountUnit,omitempty"`

    // Category 价格所属类别。
    Category *string `json:"category,omitempty"`

}

// StepPrice 描述阶梯价格的信息。
type StepPrice struct {

    // StepStart 阶梯的起始值。
    StepStart *float64 `json:"stepStart,omitempty"`

    // StepEnd 阶梯的到达值。
    // 为null代表最后一级阶梯。
    StepEnd *float64 `json:"stepEnd,omitempty"`

    // UnitPrice 阶梯单价。
    UnitPrice *float64 `json:"unitPrice,omitempty"`

    // DiscountUnitPrice 阶梯折后价。
    DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`

}

type InquiryPriceCreateInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceCreateInstanceResponseParams `json:"response,omitempty"`

}

type InquiryPriceCreateInstanceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstancePrice 实例价格。
    InstancePrice *PriceItem `json:"instancePrice,omitempty"`

    // BandwidthPrice 公网带宽价格。
    BandwidthPrice []*PriceItem `json:"bandwidthPrice,omitempty"`

    // EipPrice 弹性IP价格。
    EipPrice *PriceItem `json:"eipPrice,omitempty"`

    // SystemDiskPrice 系统盘价格。
    SystemDiskPrice *PriceItem `json:"systemDiskPrice,omitempty"`

    // DataDiskPrice 数据盘价格。
    DataDiskPrice *PriceItem `json:"dataDiskPrice,omitempty"`

    // DataDiskPrices 每种规格数据盘的价格。
    DataDiskPrices []*DataDisk `json:"dataDiskPrices,omitempty"`

}

type DescribeZoneInstanceConfigInfosRequest struct {
    *common.BaseRequest

    // InstanceChargeType 实例计费类型。
    // PREPAID：预付费。
    // POSTPAID：后付费。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceType 实例机型。
    InstanceType *string `json:"instanceType,omitempty"`

}

type DescribeZoneInstanceConfigInfosResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeZoneInstanceConfigInfosResponseParams `json:"response,omitempty"`

}

type DescribeZoneInstanceConfigInfosResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceTypeQuotaSet 可用区机型配置列表。
    InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"instanceTypeQuotaSet,omitempty"`

}

// InstanceTypeQuotaItem 描述可用区的机型配置信息。
type InstanceTypeQuotaItem struct {

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceType 实例机型。
    InstanceType *string `json:"instanceType,omitempty"`

    // InstanceTypeName 实例机型名称。
    InstanceTypeName *string `json:"instanceTypeName,omitempty"`

    // CpuCount CPU核数。
    CpuCount *int `json:"cpuCount,omitempty"`

    // Memory 内存大小，单位GiB。
    Memory *int `json:"memory,omitempty"`

    // InternetMaxBandwidthOutLimit 公网出口带宽上限。
    InternetMaxBandwidthOutLimit *int `json:"internetMaxBandwidthOutLimit,omitempty"`

    // Frequency CPU主频。
    Frequency *string `json:"frequency,omitempty"`

    // InternetChargeTypes 支持的网络计费类型列表。
    InternetChargeTypes []string `json:"internetChargeTypes,omitempty"`

}

// DescribeVmInventoryCapacityRequest 
type DescribeVmInventoryCapacityRequest struct {
    *common.BaseRequest

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

}

type DescribeVmInventoryCapacityResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeVmInventoryCapacityResponseParams `json:"response,omitempty"`

}

// DescribeVmInventoryCapacityResponseParams 
type DescribeVmInventoryCapacityResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataSet 可用区库存集合。
    DataSet []*VmInventoryCapacityInfo `json:"dataSet,omitempty"`

}

// VmInventoryCapacityInfo 可用区库存相关信息。
type VmInventoryCapacityInfo struct {

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Capacity 档位。
    // 库存容量根据所有机型可售核数定义。
    Capacity *string `json:"capacity,omitempty"`

}

type CreateInstancesRequest struct {
    *common.BaseRequest

    // ZoneId 实例所在节点ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ImageId 指定有效的镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // InstanceType 实例机型。
    InstanceType *string `json:"instanceType,omitempty"`

    // InstanceName 实例显示名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // Password 实例的密码。
    Password *string `json:"password,omitempty"`

    // KeyId 密钥ID。
    KeyId *string `json:"keyId,omitempty"`

    // ResourceGroupId 实例所在的资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // InstanceChargeType 付费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // InstanceChargePrepaid 预付费模式，即包年包月相关参数设置。
    // 若指定实例的付费模式为预付费则该参数必传。
    InstanceChargePrepaid *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`

    // InstanceChargePostpaid 后付费模式相关参数设置。
    InstanceChargePostpaid *ChargePostpaid `json:"instanceChargePostpaid,omitempty"`

    // InternetChargeType 网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // TrafficPackageSize 流量包订购大小，单位TB。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // InternetMaxBandwidthOut 公网出带宽上限，单位Mbps。
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

    // InstanceCount 指定创建实例的数量。
    InstanceCount *int `json:"instanceCount,omitempty"`

    // SystemDisk 实例系统盘配置信息。
    SystemDisk *SystemDisk `json:"systemDisk,omitempty"`

    // DataDisks 实例数据盘配置信息。
    DataDisks []*DataDisk `json:"dataDisks,omitempty"`

    // SubnetId 私有网络子网ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // EnableIpv6 是否开启公网IPv6。
    EnableIpv6 *bool `json:"enableIpv6,omitempty"`

    // EnableIpv4 是否开启公网IPv4。
    EnableIpv4 *bool `json:"enableIpv4,omitempty"`

    // CpuPassThrough 是否开启CPU穿透。
    CpuPassThrough *bool `json:"cpuPassThrough,omitempty"`

    // InitScript 初始化脚本。
    InitScript *string `json:"initScript,omitempty"`

    // NetworkMode 网卡模式。
    // Vf：物理直通模式。
    // Virtio：软件模拟模式。
    NetworkMode *string `json:"networkMode,omitempty"`

    // DiskPreAllocated 硬盘数据预分配。
    DiskPreAllocated *bool `json:"diskPreAllocated,omitempty"`

    // Nic 网卡配置。
    Nic *Nic `json:"nic,omitempty"`

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // ClusterId 带宽组ID。
    ClusterId *string `json:"clusterId,omitempty"`

    // CidrBlockId CIDR 地址块ID。指定该字段将从CIDR 地址块里分配公网IP
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

    // StartCidrIpv4 CIDR地址段内的起始IP地址。
    // 该字段需要配额`cidrBlockId`一起使用，该字段代表将从该地址起始从地址段中给机器分配公网IP。
    StartCidrIpv4 *string `json:"startCidrIpv4,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

    // Tags 创建实例时关联的标签。
    // 注意：关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

// Nic 网络接口卡配置，包括公网和内网网卡名称设置。
type Nic struct {

    // WanName 公网网卡名称。
    WanName *string `json:"wanName,omitempty"`

    // LanName 内网网卡名称。
    LanName *string `json:"lanName,omitempty"`

}

// MarketingInfo 描述市场活动的相关信息。
type MarketingInfo struct {

    // DiscountCode 使用市场发放的折扣码。
    // 如果折扣码不存在，最终折扣将不会生效。
    DiscountCode *string `json:"discountCode,omitempty"`

    // UsePocVoucher 是否使用POC代金券。
    // 如果系统不存在POC代金券，相关创建流程会失败。
    UsePocVoucher *bool `json:"usePocVoucher,omitempty"`

}

// TagAssociation 描述创建资源时同时绑定的标签对的信息。
type TagAssociation struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// Tag 描述一个标签键值对的信息。
type Tag struct {

    // Key 标签键。
    // 长度限制：1～64个字符。
    Key *string `json:"key,omitempty"`

    // Value 标签值。
    // 长度限制：1～64个字符。
    Value *string `json:"value,omitempty"`

}

type CreateInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateInstancesResponseParams `json:"response,omitempty"`

}

type CreateInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // InstanceIdSet 虚拟机实例ID列表。
    // 当通过本接口来创建实例时会返回该参数，表示一个或多个实例ID。
    // 返回实例ID列表并不代表实例创建成功，可根据 DescribeInstances 接口查询返回的dataSet中对应实例的状态来判断创建是否完成：如果实例状态由DEPLOYING(部署中)或PENDING（待创建）变为RUNNING(运行中)，则为创建成功；如果实例找不到或状态变为CREATE_FAILED，表示创建失败。
    InstanceIdSet []string `json:"instanceIdSet,omitempty"`

    // Instances 随机器创建的数据盘ID集合。
    // 如果请求中没有指定数据盘，返回空数组。
    Instances []*DiskWithInstance `json:"instances,omitempty"`

}

// DiskWithInstance 实例与关联的云盘信息。
type DiskWithInstance struct {

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // DiskIdSet 随机器创建的数据盘ID集合。
    DiskIdSet []string `json:"diskIdSet,omitempty"`

}

type ResetInstancesPasswordRequest struct {
    *common.BaseRequest

    // InstanceIds 虚拟机实例ID集合。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // Password 实例登录密码。
    Password *string `json:"password,omitempty"`

}

type ResetInstancesPasswordResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ResetInstanceRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的虚拟机实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // Password 实例登录密码。
    Password *string `json:"password,omitempty"`

    // KeyId 密钥ID。
    KeyId *string `json:"keyId,omitempty"`

    // ImageId 指定有效的镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // InstanceName 实例显示名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // WanName 公网网卡名称。
    WanName *string `json:"wanName,omitempty"`

    // LanName 内网网卡名称。
    LanName *string `json:"lanName,omitempty"`

    // InitScript 初始化脚本。
    InitScript *string `json:"initScript,omitempty"`

}

type ResetInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyInstancesResourceGroupRequest struct {
    *common.BaseRequest

    // InstanceIds 虚拟机实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type ModifyInstancesResourceGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 虚拟机实例ID。
    // 最多支持100个ID查询。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // ZoneId 实例所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InternetChargeType 网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // InstanceType 实例机型。
    // 具体取值可通过调用接口DescribeZoneInstanceConfigInfos来获得最新的规格表。
    InstanceType *string `json:"instanceType,omitempty"`

    // KeyId 密钥ID。
    KeyId *string `json:"keyId,omitempty"`

    // PublicIpAddresses 公网IPv4地址。
    PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`

    // PrivateIpAddresses 子网内网的IPv4地址。
    PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`

    // InstanceStatus 实例状态。
    InstanceStatus *string `json:"instanceStatus,omitempty"`

    // InstanceName 实例显示名称。
    // 如果该值以*结尾，则对instanceName进行模糊匹配，否则将进行精确匹配。
    InstanceName *string `json:"instanceName,omitempty"`

    // SubnetId 虚拟子网ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    // 默认为1。
    PageNum *int `json:"pageNum,omitempty"`

    // ResourceGroupId 资源组的ID。
    // 如果不传，则返回该用户可见的所有资源组内的实例。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstancesResponseParams `json:"response,omitempty"`

}

// DescribeInstancesResponseParams 
type DescribeInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 实例结果集。
    DataSet []*InstanceInfo `json:"dataSet,omitempty"`

}

// InstanceInfo 实例相关信息。
type InstanceInfo struct {

    // InstanceId 实例唯一ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // ZoneId 实例所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceName 实例显示名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // InstanceType 实例机型ID。
    InstanceType *string `json:"instanceType,omitempty"`

    // CpuCount CPU核数，单位：核。
    CpuCount *int `json:"cpuCount,omitempty"`

    // Memory 实例内存容量，单位：GiB。
    Memory *int `json:"memory,omitempty"`

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // ImageName 镜像名称。
    ImageName *string `json:"imageName,omitempty"`

    // InstanceChargeType 实例计费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // InternetMaxBandwidthOut 公网出口带宽，单位：Mbps。
    // 0代表无限制，但是不会超过机型的最大上限。
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

    // InternetChargeType 网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // Period 购买实例的时长，单位：月。
    // 后付费实例该字段为null。
    Period *int `json:"period,omitempty"`

    // PublicIpAddresses 实例公网IPv4列表。
    // 如果机器的主IP未加入到公网组网接口，那么主IP将无法使用，且该字段也不会返回该IP。
    PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`

    // PublicIpv6Addresses 实例公网IPv6列表。
    PublicIpv6Addresses []string `json:"publicIpv6Addresses,omitempty"`

    // PrivateIpAddresses 实例内网IP列表。
    PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`

    // SubnetId 实例所属的内网子网ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // CreateTime 创建时间。
    // 格式为：YYYY-MM-DDThh:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 到期时间。
    // 格式为：YYYY-MM-DDThh:mm:ssZ。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // ResourceGroupId 实例所属资源组的ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 实例所属资源组的名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // InstanceStatus 实例状态。
    InstanceStatus *string `json:"instanceStatus,omitempty"`

    // TrafficPackageSize 流量包订购大小，单位为TB。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // SecurityGroupIds 实例加入的安全组列表。
    SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

    // SystemDisk 实例系统盘信息。
    SystemDisk *SystemDisk `json:"systemDisk,omitempty"`

    // DataDisks 实例数据盘信息。
    DataDisks []*DataDisk `json:"dataDisks,omitempty"`

    // AutoRenew 是否自动续费。
    // 对于预付费实例，取消订阅后，该字段值将返回false。
    AutoRenew *bool `json:"autoRenew,omitempty"`

    // KeyId 密钥ID。
    // 注意：此字段可能返回null，表示取不到有效值。
    KeyId *string `json:"keyId,omitempty"`

    // Nic 网卡配置。
    Nic *Nic `json:"nic,omitempty"`

    // Tags 资源关联的标签信息。
    Tags *Tags `json:"tags,omitempty"`

}

// Tags 描述资源关联的标签信息。
type Tags struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeInstancesStatusRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID集合。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeInstancesStatusResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstancesStatusResponseParams `json:"response,omitempty"`

}

// DescribeInstancesStatusResponseParams 
type DescribeInstancesStatusResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 实例状态列表。
    DataSet []*InstanceStatus `json:"dataSet,omitempty"`

}

// InstanceStatus 描述实例的状态。
type InstanceStatus struct {

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceStatus 实例状态。
    InstanceStatus *string `json:"instanceStatus,omitempty"`

}

type StartInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的虚拟机实例ID。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type StartInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeVncUrlRequest 
type DescribeVncUrlRequest struct {
    *common.BaseRequest

    // InstanceId 要查询的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type DescribeVncUrlResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeVncUrlResponseParams `json:"response,omitempty"`

}

// DescribeVncUrlResponseParams 
type DescribeVncUrlResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Url VNC地址URL。
    Url *string `json:"url,omitempty"`

}

type ModifyInstancesAttributeRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的实例ID。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // InstanceName 实例名称。
    InstanceName *string `json:"instanceName,omitempty"`

}

type ModifyInstancesAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type StopInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的实例ID。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // ForceShutdown 是否强制关机。
    // 如果不指定默认为是。
    ForceShutdown *bool `json:"forceShutdown,omitempty"`

}

type StopInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RebootInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的实例ID。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type RebootInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyInstanceTypeRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceTypeId 要更换的机型ID。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

}

type ModifyInstanceTypeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ModifyInstanceTypeResponseParams `json:"response,omitempty"`

}

type ModifyInstanceTypeResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type DescribeInstanceTypeStatusRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type DescribeInstanceTypeStatusResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceTypeStatusResponseParams `json:"response,omitempty"`

}

type DescribeInstanceTypeStatusResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceId 虚拟机实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 实例的名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // InstanceType 当前实例的机型。
    InstanceType *string `json:"instanceType,omitempty"`

    // ModifiedInstanceType 实例将要修改的机型。
    ModifiedInstanceType *string `json:"modifiedInstanceType,omitempty"`

    // ModifiedInstanceTypeStatus 实例机型状态。
    // Processing：变更中。
    // Enable：可用。
    // WaitToEnable：下周期变更。
    ModifiedInstanceTypeStatus *string `json:"modifiedInstanceTypeStatus,omitempty"`

}

type InquiryPriceInstanceTrafficPackageRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // TrafficPackageSize 流量包大小。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

}

type InquiryPriceInstanceTrafficPackageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceInstanceTrafficPackageResponseParams `json:"response,omitempty"`

}

type InquiryPriceInstanceTrafficPackageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TrafficPackagePrice 流量包价格。
    TrafficPackagePrice []*PriceItem `json:"trafficPackagePrice,omitempty"`

}

type ModifyInstanceTrafficPackageRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // TrafficPackageSize 流量包大小，单位TB。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

}

type ModifyInstanceTrafficPackageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ModifyInstanceTrafficPackageResponseParams `json:"response,omitempty"`

}

type ModifyInstanceTrafficPackageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type CancelInstanceTrafficPackageDowngradeRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type CancelInstanceTrafficPackageDowngradeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyInstanceBandwidthRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InternetMaxBandwidthOut 出口带宽大小。
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

}

type ModifyInstanceBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ModifyInstanceBandwidthResponseParams `json:"response,omitempty"`

}

type ModifyInstanceBandwidthResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type CancelInstanceBandwidthDowngradeRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type CancelInstanceBandwidthDowngradeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeInstanceInternetStatusRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type DescribeInstanceInternetStatusResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceInternetStatusResponseParams `json:"response,omitempty"`

}

type DescribeInstanceInternetStatusResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceId 虚拟机实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 实例的名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // InternetMaxBandwidthOut 当前实例的公网出口带宽大小。
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

    // ModifiedInternetMaxBandwidthOut 实例将要修改公网出口带宽大小。
    ModifiedInternetMaxBandwidthOut *int `json:"modifiedInternetMaxBandwidthOut,omitempty"`

    // ModifiedBandwidthStatus 实例带宽状态。
    // Processing：变更中。
    // Enable：可用。
    // WaitToEnable：下周期变更。
    ModifiedBandwidthStatus *string `json:"modifiedBandwidthStatus,omitempty"`

    // TrafficPackageSize 当前实例流量包大小，单位TB。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // ModifiedTrafficPackageSize 实例要修改流量包大小，单位TB。
    ModifiedTrafficPackageSize *float64 `json:"modifiedTrafficPackageSize,omitempty"`

    // ModifiedTrafficPackageStatus 实例流量包状态。
    // Processing：变更中。
    // Enable：可用。
    // WaitToEnable：下周期变更。
    ModifiedTrafficPackageStatus *string `json:"modifiedTrafficPackageStatus,omitempty"`

}

type CancelInstanceDowngradeRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type CancelInstanceDowngradeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type TerminateInstanceRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type TerminateInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ReleaseInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的实例ID。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type ReleaseInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type InquiryPriceInstanceBandwidthRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InternetMaxBandwidthOut 出口带宽大小。
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

}

type InquiryPriceInstanceBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceInstanceBandwidthResponseParams `json:"response,omitempty"`

}

type InquiryPriceInstanceBandwidthResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // BandwidthPrice 公网带宽价格。
    BandwidthPrice []*PriceItem `json:"bandwidthPrice,omitempty"`

}

type DescribeInstanceTrafficRequest struct {
    *common.BaseRequest

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // StartTime 查询开始时间。
    // ISO8601标准，UTC时间。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // ISO8601标准，UTC时间。
    EndTime *string `json:"endTime,omitempty"`

}

type DescribeInstanceTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceTrafficResponseParams `json:"response,omitempty"`

}

type DescribeInstanceTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 流量数据列表。
    DataList []*InstanceTrafficData `json:"dataList,omitempty"`

    // In95 入口带宽95值。
    In95 *int64 `json:"in95,omitempty"`

    // In95Time 入口带宽95值时间。
    In95Time *string `json:"in95Time,omitempty"`

    // InAvg 入口带宽平均值。
    InAvg *int64 `json:"inAvg,omitempty"`

    // InMax 入口带宽最大值。
    InMax *int64 `json:"inMax,omitempty"`

    // InMin 入口带宽最小值。
    InMin *int64 `json:"inMin,omitempty"`

    // InTotal 入口带宽总流量。
    InTotal *int64 `json:"inTotal,omitempty"`

    // Out95 出口带宽95值。
    Out95 *int64 `json:"out95,omitempty"`

    // Out95Time 出口带宽95值时间。
    Out95Time *string `json:"out95Time,omitempty"`

    // OutAvg 出口带宽平均值。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出口带宽最大值。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出口带宽最小值。
    OutMin *int64 `json:"outMin,omitempty"`

    // OutTotal 出口带宽总流量。
    OutTotal *int64 `json:"outTotal,omitempty"`

    // MaxBandwidth95ValueMbps 最大带宽95值，单位Mbps。
    MaxBandwidth95ValueMbps *float64 `json:"maxBandwidth95ValueMbps,omitempty"`

    // TotalUnit 总流量单位。
    TotalUnit *string `json:"totalUnit,omitempty"`

    // Unit 带宽值单位。
    Unit *string `json:"unit,omitempty"`

}

// InstanceTrafficData 实例的带宽数据。
type InstanceTrafficData struct {

    // InternetRX 入口流量，单位bps。
    InternetRX *int64 `json:"internetRX,omitempty"`

    // InternetTX 出口流量，单位bps。
    InternetTX *int64 `json:"internetTX,omitempty"`

    // Time 时间。
    Time *string `json:"time,omitempty"`

}

type DescribeInstanceCpuMonitorRequest struct {
    *common.BaseRequest

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // StartTime 查询开始时间。
    // ISO8601标准，UTC时间。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // ISO8601标准，UTC时间。
    EndTime *string `json:"endTime,omitempty"`

}

type DescribeInstanceCpuMonitorResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceCpuMonitorResponseParams `json:"response,omitempty"`

}

type DescribeInstanceCpuMonitorResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList CPU使用率列表。
    DataList []*DescribeInstanceCpuMonitorData `json:"dataList,omitempty"`

}

// DescribeInstanceCpuMonitorData 描述CPU数据的信息。
type DescribeInstanceCpuMonitorData struct {

    // Cpu CPU使用率。
    Cpu *string `json:"cpu,omitempty"`

    // Time 时间。
    Time *string `json:"time,omitempty"`

}

type DescribeDiskCategoryRequest struct {
    *common.BaseRequest

    // InstanceChargeType 实例计费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // ZoneId 可用区ID。
    // 可从DescribeZones接口中获取。
    ZoneId *string `json:"zoneId,omitempty"`

    // DiskCategory 云硬盘种类。
    // STANDARD：标准云盘。
    // SSD：固态硬盘。
    DiskCategory *string `json:"diskCategory,omitempty"`

}

type DescribeDiskCategoryResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeDiskCategoryResponseParams `json:"response,omitempty"`

}

type DescribeDiskCategoryResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CategoryZoneSet 云硬盘类型与可用区关系结果集。
    CategoryZoneSet []*DiskCategory `json:"categoryZoneSet,omitempty"`

}

// DiskCategory 云盘类型。
type DiskCategory struct {

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // CategorySet 该可用区支持的云硬盘种类集合。
    CategorySet []string `json:"categorySet,omitempty"`

}

type InquiryPriceCreateDisksRequest struct {
    *common.BaseRequest

    // ZoneId 可用区ID。
    // 可从DescribeZones接口中获取。
    ZoneId *string `json:"zoneId,omitempty"`

    // DiskSize 云硬盘大小，单位GB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskAmount 云硬盘数量。
    // 最小值与默认值均为1，最大值50。
    DiskAmount *int `json:"diskAmount,omitempty"`

    // ChargeType 付费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // ChargePrepaid 预付费模式，即包年包月相关参数设置。
    // 若指定云硬盘的付费类型为预付费则该参数必传。
    ChargePrepaid *ChargePrepaid `json:"chargePrepaid,omitempty"`

    // DiskCategory 云硬盘种类。
    // STANDARD：标准云盘。
    // SSD：固态硬盘。
    // 默认为SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

}

type InquiryPriceCreateDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceCreateDisksResponseParams `json:"response,omitempty"`

}

type InquiryPriceCreateDisksResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataDiskPrice 云硬盘价格。
    DataDiskPrice *PriceItem `json:"dataDiskPrice,omitempty"`

}

type DescribeDisksRequest struct {
    *common.BaseRequest

    // DiskIds 云硬盘ID集合。
    DiskIds []string `json:"diskIds,omitempty"`

    // DiskName 云硬盘名称。
    DiskName *string `json:"diskName,omitempty"`

    // DiskStatus 云硬盘状态。
    DiskStatus *string `json:"diskStatus,omitempty"`

    // DiskType 云硬盘类型。
    // SYSTEM：系统盘。
    // DATA：数据盘。
    DiskType *string `json:"diskType,omitempty"`

    // DiskCategory 云硬盘种类。
    // STANDARD：标准云盘。
    // SSD：固态硬盘。
    DiskCategory *string `json:"diskCategory,omitempty"`

    // DiskSize 云硬盘大小，单位GB。
    DiskSize *int `json:"diskSize,omitempty"`

    // Portable 是否可拔插。
    // false代表会随实例一起删除。
    // true代表不会随实例一起删除。
    Portable *bool `json:"portable,omitempty"`

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // ZoneId 可用区ID。
    // 可从DescribeZones接口中获取。
    ZoneId *string `json:"zoneId,omitempty"`

    // PageNum 返回的分页数。
    // 默认为1。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeDisksResponseParams `json:"response,omitempty"`

}

type DescribeDisksResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 结果集。
    DataSet []*DiskInfo `json:"dataSet,omitempty"`

}

// DiskInfo 云硬盘信息。
type DiskInfo struct {

    // DiskId 云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

    // DiskName 云硬盘名称。
    DiskName *string `json:"diskName,omitempty"`

    // ZoneId 云盘所属区域。
    ZoneId *string `json:"zoneId,omitempty"`

    // DiskType 云盘类型。
    // SYSTEM：系统盘。
    // DATA：数据盘。
    DiskType *string `json:"diskType,omitempty"`

    // Portable 是否可拔插。
    Portable *bool `json:"portable,omitempty"`

    // DiskCategory 云硬盘种类。
    // STANDARD：标准云盘。
    // SSD：固态硬盘。
    DiskCategory *string `json:"diskCategory,omitempty"`

    // DiskSize 云盘大小，单位GB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskStatus 云盘状态。
    DiskStatus *string `json:"diskStatus,omitempty"`

    // InstanceId 所绑定的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 所绑定的实例名字。
    InstanceName *string `json:"instanceName,omitempty"`

    // ChargeType 付费类型。
    // PREPAID：预付费。
    // POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // CreateTime 创建时间。
    // 格式为：YYYY-MM-DDThh:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 到期时间。
    // 格式为：YYYY-MM-DDThh:mm:ssZ。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // Period 购买实例的时长，单位：月。
    // 后付费实例该字段为null。
    Period *int `json:"period,omitempty"`

    // AutoRenew 是否自动续费。
    // 对于预付费实例，取消订阅后，该字段值将返回false。
    AutoRenew *bool `json:"autoRenew,omitempty"`

    // Tags 资源关联的标签信息。
    Tags *Tags `json:"tags,omitempty"`

}

type CreateDisksRequest struct {
    *common.BaseRequest

    // ChargeType 付费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // ChargePrepaid 预付费模式，即包年包月相关参数设置。
    // 若指定云硬盘的付费模式为预付费则该参数必传。
    ChargePrepaid *ChargePrepaid `json:"chargePrepaid,omitempty"`

    // DiskName 云硬盘名称。
    // 必须以数字或字母开头或结尾，长度1-64字符，仅支持字母、数字、-和英文句点(.)。
    DiskName *string `json:"diskName,omitempty"`

    // DiskSize 云硬盘大小，单位GB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskAmount 云硬盘创建数量。
    // 最小值与默认值均为1，最大值50。
    DiskAmount *int `json:"diskAmount,omitempty"`

    // InstanceId 创建后需要挂载的实例ID。
    // 指定实例ID将会为实例所在的可用区创建数据盘并挂载到实例上。
    InstanceId *string `json:"instanceId,omitempty"`

    // ZoneId 云硬盘所属的可用区ID。
    // 如果指定了instanceId，则该字段无效。
    ZoneId *string `json:"zoneId,omitempty"`

    // ResourceGroupId 云硬盘所在的资源组ID，如不指定则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // DiskCategory 云硬盘种类。
    // STANDARD：标准云盘。
    // SSD：固态硬盘。
    // 默认为SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

    // Tags 创建云硬盘时关联的标签。
    // 注意：关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

type CreateDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateDisksResponseParams `json:"response,omitempty"`

}

type CreateDisksResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DiskIds 云硬盘ID集合。
    DiskIds []string `json:"diskIds,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type ModifyDisksAttributesRequest struct {
    *common.BaseRequest

    // DiskIds 待修改属性的云硬盘ID集合。
    DiskIds []string `json:"diskIds,omitempty"`

    // DiskName 新的云盘名称。
    // 必须以数字或字母开头或结尾，长度1-64字符，仅支持字母、数字、-和英文句点(.)。
    DiskName *string `json:"diskName,omitempty"`

}

type ModifyDisksAttributesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type AttachDisksRequest struct {
    *common.BaseRequest

    // DiskIds 云硬盘ID集合。
    DiskIds []string `json:"diskIds,omitempty"`

    // InstanceId 需要挂载的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type AttachDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DetachDisksRequest struct {
    *common.BaseRequest

    // DiskIds 将要卸载的云硬盘ID集合。
    DiskIds []string `json:"diskIds,omitempty"`

}

type DetachDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ChangeDisksAttachRequest struct {
    *common.BaseRequest

    // DiskIds 云硬盘ID集合。
    DiskIds []string `json:"diskIds,omitempty"`

    // InstanceId 需要挂载的新实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type ChangeDisksAttachResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type TerminateDiskRequest struct {
    *common.BaseRequest

    // DiskId 云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

}

type TerminateDiskResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyDisksResourceGroupRequest struct {
    *common.BaseRequest

    // DiskIds 云硬盘ID列表。
    // 每次请求允许操作的云硬盘数量上限是100。
    DiskIds []string `json:"diskIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type ModifyDisksResourceGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ResizeDiskRequest struct {
    *common.BaseRequest

    // DiskId 云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

    // DiskSize 扩容后的云硬盘大小，单位GB。
    DiskSize *int `json:"diskSize,omitempty"`

}

type ResizeDiskResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ReleaseDiskRequest struct {
    *common.BaseRequest

    // DiskId 云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

}

type ReleaseDiskResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RenewDiskRequest struct {
    *common.BaseRequest

    // DiskId 云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

}

type RenewDiskResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *RenewDiskResponseParams `json:"response,omitempty"`

}

type RenewDiskResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type DescribeImageQuotaRequest struct {
    *common.BaseRequest

    // ZoneIds 可用区ID列表。
    // 可从DescribeZones的zoneId中获取。
    ZoneIds []string `json:"zoneIds,omitempty"`

}

type DescribeImageQuotaResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeImageQuotaResponseParams `json:"response,omitempty"`

}

type DescribeImageQuotaResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Images 镜像配额结果集。
    Images []*ImageQuotaInfo `json:"images,omitempty"`

}

// ImageQuotaInfo 镜像的配额信息。
type ImageQuotaInfo struct {

    // ZoneId 支持创建镜像的区域。
    ZoneId *string `json:"zoneId,omitempty"`

    // Count 当前已配置镜像数。
    Count *int `json:"count,omitempty"`

    // MaxCount 本区域可配置的最大镜像数。
    MaxCount *int `json:"maxCount,omitempty"`

}

type DescribeImagesRequest struct {
    *common.BaseRequest

    // ImageIds 镜像ID集合。
    ImageIds []string `json:"imageIds,omitempty"`

    // ImageName 镜像名称。
    ImageName *string `json:"imageName,omitempty"`

    // ZoneId 可用区ID。
    // 可从DescribeZones的zoneId中获取。
    ZoneId *string `json:"zoneId,omitempty"`

    // Category 镜像所属分类。
    // 可能值：CentOS、Windows、Ubuntu、Debian。
    Category *string `json:"category,omitempty"`

    // ImageType 镜像类型。
    // PUBLIC_IMAGE：公共镜像。
    // CUSTOM_IMAGE：自定义镜像。
    ImageType *string `json:"imageType,omitempty"`

    // OsType 操作系统类型。
    // 可能值：windows、linux。
    OsType *string `json:"osType,omitempty"`

    // ImageStatus 镜像状态。
    // CREATING：创建中。
    // AVAILABLE：可用。
    // UNAVAILABLE：不可用。
    ImageStatus *string `json:"imageStatus,omitempty"`

    // PageNum 返回的分页数。
    // 默认为1。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

}

type DescribeImagesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeImagesResponseParams `json:"response,omitempty"`

}

// DescribeImagesResponseParams 
type DescribeImagesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 镜像结果集。
    DataSet []*ImageInfo `json:"dataSet,omitempty"`

}

// ImageInfo 镜像相关信息。
type ImageInfo struct {

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // ImageName 镜像名称。
    ImageName *string `json:"imageName,omitempty"`

    // ImageType 镜像类型。
    // PUBLIC_IMAGE：公共镜像。
    // CUSTOM_IMAGE：自定义镜像。
    ImageType *string `json:"imageType,omitempty"`

    // ImageSize 镜像大小，单位为GB。
    ImageSize *string `json:"imageSize,omitempty"`

    // ImageDescription 镜像描述。
    ImageDescription *string `json:"imageDescription,omitempty"`

    // ImageVersion 镜像版本。
    ImageVersion *string `json:"imageVersion,omitempty"`

    // ImageStatus 镜像状态。
    // CREATING：创建中。
    // AVAILABLE：可用。
    // UNAVAILABLE：不可用。
    ImageStatus *string `json:"imageStatus,omitempty"`

    // Category 镜像所属分类。
    Category *string `json:"category,omitempty"`

    // OsType 操作系统类型。
    // 可能值：windows、linux。
    OsType *string `json:"osType,omitempty"`

}

type CreateImageRequest struct {
    *common.BaseRequest

    // InstanceId 需要制作镜像的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // ImageName 镜像名称。
    // 长度不超过24位，支持中文、字母、数字或连接符号-_。
    ImageName *string `json:"imageName,omitempty"`

    // ImageDescription 镜像描述。
    // 不超过255个字符。
    ImageDescription *string `json:"imageDescription,omitempty"`

}

type CreateImageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateImageResponseParams `json:"response,omitempty"`

}

type CreateImageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

}

type ModifyImagesAttributesRequest struct {
    *common.BaseRequest

    // ImageIds 镜像ID集合。
    // 可从DescribeImages返回的imageId获取。
    ImageIds []string `json:"imageIds,omitempty"`

    // ImageDescription 新的镜像描述。
    // 描述信息不得超过255个字符。
    ImageDescription *string `json:"imageDescription,omitempty"`

    // ImageName 新的镜像名称。
    // 长度不超过24位，支持中文、字母、数字或连接符号-_。
    ImageName *string `json:"imageName,omitempty"`

}

type ModifyImagesAttributesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteImagesRequest struct {
    *common.BaseRequest

    // ImageIds 将要被删除的镜像ID集合。
    ImageIds []string `json:"imageIds,omitempty"`

}

type DeleteImagesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeSubnetsRequest struct {
    *common.BaseRequest

    // ZoneId 子网所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // SubnetIds 子网 ID。
    // 取值可以由多个Subnet ID组成一个。
    // 最多支持100个ID查询。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // CidrBlock 子网的CIDR。
    // 支持模糊查询。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // SubnetStatus Subnet的状态。
    SubnetStatus *string `json:"subnetStatus,omitempty"`

    // SubnetName 子网的名称。
    // 支持模糊查询。
    SubnetName *string `json:"subnetName,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    // 默认为1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeSubnetsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSubnetsResponseParams `json:"response,omitempty"`

}

// DescribeSubnetsResponseParams 
type DescribeSubnetsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 子网结果集。
    DataSet []*SubnetInfo `json:"dataSet,omitempty"`

}

// SubnetInfo 子网信息。
type SubnetInfo struct {

    // SubnetId Subnet的ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // ZoneId Subnet的机房ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // SubnetName Subnet的名称。
    SubnetName *string `json:"subnetName,omitempty"`

    // SubnetStatus Subnet的状态。
    SubnetStatus *string `json:"subnetStatus,omitempty"`

    // CidrBlockList Subnet的CIDR列表。
    CidrBlockList []string `json:"cidrBlockList,omitempty"`

    // UsageIpCount Subnet的已用IP数。
    UsageIpCount *int `json:"usageIpCount,omitempty"`

    // TotalIpCount Subnet的总IP数。
    TotalIpCount *int `json:"totalIpCount,omitempty"`

    // CreateTime Subnet的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // InstanceIdList Subnet下绑定的实例列表。
    InstanceIdList []string `json:"instanceIdList,omitempty"`

    // SubnetDescription Subnet的描述信息。
    SubnetDescription *string `json:"subnetDescription,omitempty"`

    // CidrBlock Subnet的CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // IsDefault Subnet是否为默认。
    IsDefault *bool `json:"isDefault,omitempty"`

}

type CreateSubnetRequest struct {
    *common.BaseRequest

    // CidrBlock 子网的CIDR。
    // 可选值10.0.0.0/16、172.16.0.0/16和192.168.0.0/16及它们包含的子网。
    // 子网网段不能重叠。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // SubnetName 子网的名称。
    // 范围2到63个字符。
    // 仅支持输入字母、数字、-和英文句点(.)。
    // 且必须以数字或字母开头和结尾。
    SubnetName *string `json:"subnetName,omitempty"`

    // SubnetDescription 子网的描述信息。
    SubnetDescription *string `json:"subnetDescription,omitempty"`

    // ZoneId 子网的节点ID。
    ZoneId *string `json:"zoneId,omitempty"`

}

type CreateSubnetResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateSubnetResponseParams `json:"response,omitempty"`

}

type CreateSubnetResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SubnetId 子网的ID。
    SubnetId *string `json:"subnetId,omitempty"`

}

type ModifySubnetsAttributeRequest struct {
    *common.BaseRequest

    // SubnetIds 一个或多个待操作的Subnet ID。
    // 可通过DescribeSubnets接口返回值中的SubnetId获取。
    // 每次请求批量Subnet的上限为100。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // SubnetName Subnet名称。
    // 范围1到64个字符。
    // 仅支持输入字母、数字、-和英文句点(.)。
    SubnetName *string `json:"subnetName,omitempty"`

}

type ModifySubnetsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteSubnetRequest struct {
    *common.BaseRequest

    // SubnetId 子网的ID。
    SubnetId *string `json:"subnetId,omitempty"`

}

type DeleteSubnetResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeSecurityGroupsRequest struct {
    *common.BaseRequest

    // SecurityGroupIds 安全组ID集合。
    // 最多支持100个ID查询。
    SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

    // SecurityGroupName 安全组名称。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    // 默认为1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeSecurityGroupsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSecurityGroupsResponseParams `json:"response,omitempty"`

}

// DescribeSecurityGroupsResponseParams 
type DescribeSecurityGroupsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 安全组结果集。
    DataSet []*SecurityGroupInfo `json:"dataSet,omitempty"`

}

// SecurityGroupInfo 安全组信息。
type SecurityGroupInfo struct {

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // SecurityGroupName 安全组名称。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // SecurityGroupStatus 安全组状态。
    SecurityGroupStatus *string `json:"securityGroupStatus,omitempty"`

    // CreateTime 创建时间。
    // 格式为：YYYY-MM-DDThh:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // Description 安全组描述。
    Description *string `json:"description,omitempty"`

    // InstanceIds 已绑定实例ID集合。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // RuleInfos 安全组规则。
    RuleInfos []*RuleInfo `json:"ruleInfos,omitempty"`

    // IsDefault 是否默认。
    IsDefault *bool `json:"isDefault,omitempty"`

}

// RuleInfo 安全组规则信息。
type RuleInfo struct {

    // Direction 规则方向。
    // ingress：入方向。
    // egress：出方向。
    Direction *string `json:"direction,omitempty"`

    // Policy 设置访问权限。
    // accept：接受访问。
    // 目前只支持accept。
    Policy *string `json:"policy,omitempty"`

    // Priority 规则优先级。
    Priority *int `json:"priority,omitempty"`

    // IpProtocol 传输层协议。
    // 取值范围：tcp、udp、icmp、all。
    IpProtocol *string `json:"ipProtocol,omitempty"`

    // PortRange 目的端安全组开放的传输层协议相关的端口范围。
    PortRange *string `json:"portRange,omitempty"`

    // CidrIp 源端IP地址范围。
    // 支持CIDR格式和IPv4格式的IP地址范围。
    CidrIp *string `json:"cidrIp,omitempty"`

    // Description 规则描述。
    Description *string `json:"description,omitempty"`

}

type CreateSecurityGroupRequest struct {
    *common.BaseRequest

    // SecurityGroupName 安全组名称。
    // 范围1到64个字符，仅支持字母、数字、-和英文句点(.)。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // RuleInfos 安全组规则。
    RuleInfos []*RuleInfo `json:"ruleInfos,omitempty"`

    // Description 安全组描述。
    // 范围2到256个字符。
    Description *string `json:"description,omitempty"`

}

type CreateSecurityGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateSecurityGroupResponseParams `json:"response,omitempty"`

}

type CreateSecurityGroupResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type ModifySecurityGroupsAttributeRequest struct {
    *common.BaseRequest

    // SecurityGroupName 安全组名称。
    // 范围1到64个字符，仅支持字母、数字、-和英文句点(.)。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // Description 安全组描述。
    // 范围2到256个字符。
    Description *string `json:"description,omitempty"`

    // SecurityGroupIds 一个或多个待操作的安全组ID。
    // 每次请求批量上限为100。
    SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

}

type ModifySecurityGroupsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteSecurityGroupRequest struct {
    *common.BaseRequest

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type DeleteSecurityGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ConfigureSecurityGroupRulesRequest struct {
    *common.BaseRequest

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // RuleInfos 安全组规则。
    RuleInfos []*RuleInfo `json:"ruleInfos,omitempty"`

}

type ConfigureSecurityGroupRulesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RevokeSecurityGroupRulesRequest struct {
    *common.BaseRequest

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // RuleInfos 安全组规则。
    RuleInfos []*RuleInfo `json:"ruleInfos,omitempty"`

}

type RevokeSecurityGroupRulesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type AssociateSecurityGroupInstanceRequest struct {
    *common.BaseRequest

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type AssociateSecurityGroupInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type UnAssociateSecurityGroupInstanceRequest struct {
    *common.BaseRequest

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type UnAssociateSecurityGroupInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeInstanceAvailableSecurityGroupResourcesRequest struct {
    *common.BaseRequest

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type DescribeInstanceAvailableSecurityGroupResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceAvailableSecurityGroupResourcesResponseParams `json:"response,omitempty"`

}

type DescribeInstanceAvailableSecurityGroupResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceAvailableSecurityGroups 实例可绑定的安全组集合。
    InstanceAvailableSecurityGroups []*InstanceAvailableSecurityGroup `json:"instanceAvailableSecurityGroups,omitempty"`

}

// InstanceAvailableSecurityGroup 描述实例可绑定的安全组信息。
type InstanceAvailableSecurityGroup struct {

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // SecurityGroupName 安全组名称。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // IsDefault 安全组是否默认。
    IsDefault *bool `json:"isDefault,omitempty"`

}

type AuthorizeSecurityGroupRulesRequest struct {
    *common.BaseRequest

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // RuleInfos 安全组规则。
    RuleInfos []*RuleInfo `json:"ruleInfos,omitempty"`

}

type AuthorizeSecurityGroupRulesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type AuthorizeSecurityGroupRuleRequest struct {
    *common.BaseRequest

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // Direction 规则方向。
    // ingress：入方向。
    // egress：出方向。
    Direction *string `json:"direction,omitempty"`

    // Policy 设置访问权限。
    // accept（默认值）：接受访问。
    Policy *string `json:"policy,omitempty"`

    // Priority 规则优先级。
    Priority *int `json:"priority,omitempty"`

    // IpProtocol 传输层协议。
    // 取值范围：tcp、udp、icmp、all。
    IpProtocol *string `json:"ipProtocol,omitempty"`

    // PortRange 目的端安全组开放的传输层协议相关的端口范围。
    PortRange *string `json:"portRange,omitempty"`

    // CidrIp 源端IP地址范围。
    CidrIp *string `json:"cidrIp,omitempty"`

    // Description 规则描述。
    Description *string `json:"description,omitempty"`

}

type AuthorizeSecurityGroupRuleResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

