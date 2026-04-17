package bmc

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


type DescribeSubnetAvailableResourcesRequest struct {
    *common.BaseRequest

    // ZoneId 可用区ID。
    // 不传则查询所有可创建Subnet的可用区。
    ZoneId *string `json:"zoneId,omitempty"`

}

type DescribeSubnetAvailableResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSubnetAvailableResourcesResponseParams `json:"response,omitempty"`

}

type DescribeSubnetAvailableResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ZoneIdSet zone结果集。
    ZoneIdSet []string `json:"zoneIdSet,omitempty"`

}

type ModifySubnetsAttributeRequest struct {
    *common.BaseRequest

    // SubnetName Subnet名称。
    // 范围1到64个字符。仅支持输入字母、数字、-和英文句点(.)。
    SubnetName *string `json:"subnetName,omitempty"`

    // SubnetIds 一个或多个待操作的Subnet ID。
    // 可通过DescribeSubnets接口返回值中的subnetId获取。
    // 每次请求批量Subnet的上限为100。
    SubnetIds []string `json:"subnetIds,omitempty"`

}

type ModifySubnetsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type CreateSubnetRequest struct {
    *common.BaseRequest

    // ZoneId 可用区的节点ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // CidrBlock 子网的CIDR。
    // 可选值10.0.0.0/16、172.16.0.0/16和192.168.0.0/16及它们包含的子网。
    // 如果指定了VPC ID，那么子网网段必须在VPC的CIDR范围之内，且不能和VPC下其他子网网段有重叠。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // SubnetName 子网的名称。
    // 范围1到64个字符。仅支持输入字母、数字、-和英文句点(.)。
    SubnetName *string `json:"subnetName,omitempty"`

    // ResourceGroupId 子网所在的资源组ID。
    // 如果不指定，则会放入默认资源组， 如果用户没有默认资源组权限， 则请求将会失败。
    // 如果指定VPC，则会忽略该参数自动使用与VPC相同的资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // VpcId VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

}

type CreateSubnetResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateSubnetResponseParams `json:"response,omitempty"`

}

type CreateSubnetResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SubnetId Subnet的ID。
    SubnetId *string `json:"subnetId,omitempty"`

}

type DescribeSubnetsRequest struct {
    *common.BaseRequest

    // SubnetIds Subnet ID。
    // 取值可以由多个Subnet ID组成一个。最多支持100个ID查询。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // SubnetName Subnet的名称。
    SubnetName *string `json:"subnetName,omitempty"`

    // CidrBlock Subnet的CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // SubnetStatus Subnet的状态。
    SubnetStatus *string `json:"subnetStatus,omitempty"`

    // ZoneId Subnet所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // VpcId VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

    // ResourceGroupId 资源组的ID。
    // 如果不传，则返回该用户可见的所有资源组内的Subnet。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeSubnetsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSubnetsResponseParams `json:"response,omitempty"`

}

type DescribeSubnetsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet SubnetInfo结果集。
    DataSet []*SubnetInfo `json:"dataSet,omitempty"`

}

// SubnetInfo Subnet Info的信息。
type SubnetInfo struct {

    // SubnetId Subnet唯一ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // SubnetName Subnet的名称。
    SubnetName *string `json:"subnetName,omitempty"`

    // ZoneId Subnet所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // AvailableIpCount Subnet下可用的IP数量。
    AvailableIpCount *int `json:"availableIpCount,omitempty"`

    // CidrBlock Subnet的CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // SubnetStatus Subnet的状态。
    SubnetStatus *string `json:"subnetStatus,omitempty"`

    // CreateTime 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-ddTHH:mm:ssZ`。
    CreateTime *string `json:"createTime,omitempty"`

    // VpcSubnetStatus VPC与Subnet的绑定状态。
    VpcSubnetStatus *string `json:"vpcSubnetStatus,omitempty"`

    // VpcId VPC唯一ID。
    VpcId *string `json:"vpcId,omitempty"`

    // VpcName VPC的名称。
    VpcName *string `json:"vpcName,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // SubnetInstanceSet Subnet下实例集合。
    SubnetInstanceSet []*SubnetInstance `json:"subnetInstanceSet,omitempty"`

}

// SubnetInstance Subnet Instance信息。
type SubnetInstance struct {

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // PrivateIpAddress 私网IP。
    PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

    // PrivateIpStatus 私网IP与实例的绑定状态。
    PrivateIpStatus *string `json:"privateIpStatus,omitempty"`

}

type AssociateSubnetInstancesRequest struct {
    *common.BaseRequest

    // SubnetId Subnet ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // SubnetInstanceList Subnet绑定实例集合。
    SubnetInstanceList []*AssociateSubnetInstance `json:"subnetInstanceList,omitempty"`

}

// AssociateSubnetInstance Subnet绑定实例。
type AssociateSubnetInstance struct {

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // PrivateIpAddress 内网IPv4地址。该地址必须在子网的CIDR范围内。 如果不指定内网地址，系统会会寻找CIDR中未用的内网地址下发到实例。
    PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

}

type AssociateSubnetInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type UnAssociateSubnetInstanceRequest struct {
    *common.BaseRequest

    // SubnetId Subnet的ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // InstanceId 实例的ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type UnAssociateSubnetInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type AssociateVpcSubnetRequest struct {
    *common.BaseRequest

    // SubnetId Subnet的ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // VpcId VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

}

type AssociateVpcSubnetResponse struct {
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

type DescribeZonesRequest struct {
    *common.BaseRequest

    // AcceptLanguage 接收的区域地域的语言。可选值如下：
    // <ul><li>zh-CN：中文</li><li>en-US：英文</li></ul>默认值：en-US。
    AcceptLanguage *string `json:"acceptLanguage,omitempty"`

    // IsCloudRouterAvailable 根据可用区是否支持SDN三层网络进行筛选。可选值如下：
    // <ul><li>true：支持</li><li>false：不支持</li></ul>
    IsCloudRouterAvailable *bool `json:"isCloudRouterAvailable,omitempty"`

}

type DescribeZonesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeZonesResponseParams `json:"response,omitempty"`

}

type DescribeZonesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ZoneSet 可用区列表。
    ZoneSet []*Zone `json:"zoneSet,omitempty"`

}

// Zone 可用区信息， 包含可用区所在的城市等信息。
type Zone struct {

    // ZoneId 可用区ID。
    // 例如：SEL-A
    ZoneId *string `json:"zoneId,omitempty"`

    // ZoneName 可用区的名称。
    ZoneName *string `json:"zoneName,omitempty"`

    // CityName 可用区的城市名称。
    CityName *string `json:"cityName,omitempty"`

    // AreaName 可用区所在的大区名称。
    AreaName *string `json:"areaName,omitempty"`

    // IsCloudRouterAvailable 可用区是否支持SDN三层网络。
    IsCloudRouterAvailable *bool `json:"isCloudRouterAvailable,omitempty"`

}

type DescribeInstanceTypesRequest struct {
    *common.BaseRequest

    // ImageId 查询实例机型时，支持某镜像的机型。
    ImageId *string `json:"imageId,omitempty"`

    // InstanceTypeIds 实例机型ID。
    // 数量不超过100个。
    InstanceTypeIds []string `json:"instanceTypeIds,omitempty"`

    // MinimumCpuCoreCount 查询实例机型时，期望最小CPU内核的数目。
    // 取值范围：正整数。
    MinimumCpuCoreCount *int `json:"minimumCpuCoreCount,omitempty"`

    // MaximumCpuCoreCount 查询实例机型时，期望最大CPU内核的数目。
    // 取值范围：正整数。
    MaximumCpuCoreCount *int `json:"maximumCpuCoreCount,omitempty"`

    // MinimumMemorySize 查询实例机型时，期望最小内存大小。
    // 取值范围：正整数。
    // 单位：GB。
    MinimumMemorySize *int `json:"minimumMemorySize,omitempty"`

    // MaximumMemorySize 查询实例机型时，期望最大内存大小。
    // 取值范围：正整数。
    // 单位：GB。
    MaximumMemorySize *int `json:"maximumMemorySize,omitempty"`

    // MinimumBandwidth 查询实例机型时，期望最小公网入方向带宽限制。
    // 单位：Mbps。
    MinimumBandwidth *int `json:"minimumBandwidth,omitempty"`

    // SupportRaids 查询实例机型时，对实例机型做raid时所支持的raid类型。
    // Raid可选值包括：0, 1, 5, 10。
    SupportRaids []int `json:"supportRaids,omitempty"`

    // SupportSubnet 查询实例机型时，机型是否支持内网组网。
    SupportSubnet *bool `json:"supportSubnet,omitempty"`

    // MinimumDiskSize 查询实例机型时，期望最小磁盘大小。
    // 取值范围：正整数。
    // 单位：GB。
    MinimumDiskSize *int `json:"minimumDiskSize,omitempty"`

    // MaximumDiskSize 查询实例机型时，期望最大磁盘大小。
    // 取值范围：正整数。
    // 单位：GB。
    MaximumDiskSize *int `json:"maximumDiskSize,omitempty"`

    // IsHA 查询实例机型时，是否是高可用机型。
    IsHA *bool `json:"isHA,omitempty"`

}

type DescribeInstanceTypesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceTypesResponseParams `json:"response,omitempty"`

}

type DescribeInstanceTypesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceTypes 查询的机型信息。
    InstanceTypes []*InstanceType `json:"instanceTypes,omitempty"`

}

// InstanceType 机型的配置信息。包括机型的cpu、内存、是否支持组内网等等。
type InstanceType struct {

    // ImageIds 机型支持的镜像ID。
    // 仅在DescribeInstanceType可取值
    ImageIds []string `json:"imageIds,omitempty"`

    // InstanceTypeId 实例机型ID。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

    // Description 机型描述。
    // 一般包括内存大小，硬盘。
    Description *string `json:"description,omitempty"`

    // CpuCoreCount CPU数量。
    CpuCoreCount *int `json:"cpuCoreCount,omitempty"`

    // CpuDetail CPU详情。
    CpuDetail *string `json:"cpuDetail,omitempty"`

    // CpuCores CPU核心数。
    CpuCores *int `json:"cpuCores,omitempty"`

    // CpuThreads CPU线程。
    CpuThreads *int `json:"cpuThreads,omitempty"`

    // BaseFrequency CPU基础频率。
    BaseFrequency *string `json:"baseFrequency,omitempty"`

    // MemorySize 内存大小。
    // 单位：GB。
    MemorySize *int `json:"memorySize,omitempty"`

    // MaximumBandwidth 机型支持的最大出口带宽。
    // 单位：Mbps。
    MaximumBandwidth *int `json:"maximumBandwidth,omitempty"`

    // SupportRaids 机型支持的raid。
    SupportRaids []int `json:"supportRaids,omitempty"`

    // SupportSubnet 是否支持内网组网。
    SupportSubnet *bool `json:"supportSubnet,omitempty"`

    // IsHA 是否是高可用机型。
    IsHA *bool `json:"isHA,omitempty"`

    // DiskInfo 硬盘配置。
    // 单位：GB。
    DiskInfo *InstanceDiskInfo `json:"diskInfo,omitempty"`

}

// InstanceDiskInfo 机型硬盘信息。
type InstanceDiskInfo struct {

    // TotalDiskSize 机型的硬盘总大小。
    // 单位：GB。
    // totalDiskSize的大小一般小于描述的信息，系统为了分区能够成功预留了一小部分。如果采用自定义分区，最后的一个分区将会获得剩余的所有磁盘大小。
    TotalDiskSize *int `json:"totalDiskSize,omitempty"`

    // DiskDescription 机型硬盘的描述信息。
    DiskDescription *string `json:"diskDescription,omitempty"`

    // Disks 可用于raid和分区的磁盘信息。
    // 按顺序标号。比如880 x 2、 220 x2，其磁盘序号1,2,3,4 分别对应的磁盘大小为880，880，220，220。
    Disks []*Disk `json:"disks,omitempty"`

}

// Disk 硬盘块信息。
type Disk struct {

    // DiskSize 硬盘的大小。
    // 单位：GB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskCount 该硬盘大小的硬盘的数量。
    DiskCount *int `json:"diskCount,omitempty"`

}

type DescribeAvailableResourcesRequest struct {
    *common.BaseRequest

    // InstanceChargeType 实例计费类型。
    // PREPAID：预付费，即包年包月。 POSTPAID：后付费。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // InstanceTypeId 实例机型ID。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // SellStatus 售卖的状态。
    // <ul><li>SELL：表示实例可购买，且库存>10。</li><li>SELL_SHORTAGE: 表示可购买，但是库存<10台。</li><li>SOLD_OUT：表示实例已售罄。</li></ul>
    SellStatus *string `json:"sellStatus,omitempty"`

}

type DescribeAvailableResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAvailableResourcesResponseParams `json:"response,omitempty"`

}

type DescribeAvailableResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // AvailableResources 可用资源的集合。
    AvailableResources []*AvailableResource `json:"availableResources,omitempty"`

}

// AvailableResource 可售卖的实例资源信息。描述了哪些可用区有哪些机型可以售卖。
type AvailableResource struct {

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // SellStatus 售卖的状态。
    // <ul><li>SELL：表示实例可购买，且库存>10。</li><li>SELL_SHORTAGE: 表示可购买，但是库存<10台。</li><li>SOLD_OUT：表示实例已售罄。</li></ul>
    SellStatus *string `json:"sellStatus,omitempty"`

    // InternetChargeTypes 网络计费类型。
    InternetChargeTypes []string `json:"internetChargeTypes,omitempty"`

    // InstanceTypeId 机型ID。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

    // MaximumBandwidthOut 最大的公网出口带宽限制。
    // 单位：Mbps。
    MaximumBandwidthOut *int `json:"maximumBandwidthOut,omitempty"`

    // DefaultBandwidthOut 固定带宽计费方式时默认赠送公网带宽。
    // 单位：GB。
    DefaultBandwidthOut *int `json:"defaultBandwidthOut,omitempty"`

    // DefaultTrafficPackageSize 流量包计费方式时默认增送的流量包大小。
    // 单位：TB。
    DefaultTrafficPackageSize *float64 `json:"defaultTrafficPackageSize,omitempty"`

    // Qty 库存数量。
    Qty *int `json:"qty,omitempty"`

}

type InquiryPriceCreateInstanceRequest struct {
    *common.BaseRequest

    // ZoneId 实例所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceTypeId 实例机型ID。
    // 具体取值可通过调用接口DescribeInstanceTypes来获得最新的规格表。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

    // InstanceChargeType 实例计费类型。
    // PREPAID：预付费，即包年包月。 POSTPAID：后付费。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // InstanceChargePrepaid 预付费模式，即包年包月相关参数设置。
    // 通过该参数可以指定包年包月实例的购买时长等属性。
    // 若指定实例的付费模式为预付费则该参数必传。
    InstanceChargePrepaid *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`

    // TrafficPackageSize 流量包订购大小。
    // 单位为TB。该值仅限当 internetChargeType = ByTrafficPackage 生效。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // InternetMaxBandwidthOut 公网出带宽上限。
    // 单位：Mbps。
    // 默认值：1Mbps。
    // 不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

    // InternetChargeType 网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

}

// ChargePrepaid 描述预付费模式，即包年包月相关参数。包括购买时长等逻辑。
type ChargePrepaid struct {

    // Period 购买实例的时长。
    // 单位：月。
    Period *int `json:"period,omitempty"`

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
    // 可能有多个价格，比如流量包计费，包含包的价格和用量超出包后的价格
    BandwidthPrice []*PriceItem `json:"bandwidthPrice,omitempty"`

    // PrimaryIpPrice 主IP的价格。
    PrimaryIpPrice *PriceItem `json:"primaryIpPrice,omitempty"`

}

// PriceItem 价格。描述了购买资源的预付费或后付费价格信息。
type PriceItem struct {

    // Discount 折扣大小。
    // 如80.0代表8折。
    Discount *float64 `json:"discount,omitempty"`

    // DiscountPrice 预付费的折扣价。
    // 预付费模式使用，后付费该值为 null。
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
    // 后付费模式使用，可取值范围：
    // HOUR: 表示计价单元是按每小时来计算。
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

// StepPrice 后付费阶梯价格。描述了价格的一个阶梯的信息。
type StepPrice struct {

    // StepStart 阶梯用量的开始。
    StepStart *float64 `json:"stepStart,omitempty"`

    // StepEnd 阶梯用量的结束。
    StepEnd *float64 `json:"stepEnd,omitempty"`

    // UnitPrice 当前阶梯的单元原始价格。
    // 后付费模式使用。
    UnitPrice *float64 `json:"unitPrice,omitempty"`

    // DiscountUnitPrice 当前阶梯的单元折后价格。
    // 后付费模式使用。
    DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`

}

type DescribeImagesRequest struct {
    *common.BaseRequest

    // ImageIds 镜像ID。
    ImageIds []string `json:"imageIds,omitempty"`

    // ImageName 镜像名称。
    ImageName *string `json:"imageName,omitempty"`

    // Catalog 镜像所属分类。
    // 可能值：
    // <ul><li>centos</li><li>windows</li><li>ubuntu</li><li>debian</li><li>esxi</li></ul>
    Catalog *string `json:"catalog,omitempty"`

    // ImageType 镜像类型。
    // PUBLIC_IMAGE: 公共镜像。
    // CUSTOM_IMAGE：自定义镜像。
    // 目前不支持自主的创建自定义镜像，如有需求，请提交support。
    ImageType *string `json:"imageType,omitempty"`

    // OsType 操作系统类型。
    // 可能值：
    // <ul><li>windows</li><li>linux</li></ul>
    OsType *string `json:"osType,omitempty"`

    // InstanceTypeId 支持的机型ID。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

}

type DescribeImagesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeImagesResponseParams `json:"response,omitempty"`

}

type DescribeImagesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Images 结果集。
    Images []*ImageInfo `json:"images,omitempty"`

}

// ImageInfo 镜像相关信息。
type ImageInfo struct {

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // ImageName 镜像名称。
    ImageName *string `json:"imageName,omitempty"`

    // Catalog 镜像所属分类。
    // 可能值：
    // <ul><li>centos</li><li>windows</li><li>ubuntu</li><li>debian</li><li>esxi</li></ul>
    Catalog *string `json:"catalog,omitempty"`

    // ImageType 镜像类型。
    // PUBLIC_IMAGE: 公共镜像。
    // CUSTOM_IMAGE：自定义镜像。
    // 目前不支持自主的创建自定义镜像，可联系support沟通。
    ImageType *string `json:"imageType,omitempty"`

    // OsType 操作系统类型。
    // 可能值：
    // <ul><li>windows</li><li>linux</li></ul>
    OsType *string `json:"osType,omitempty"`

}

type CreateInstancesRequest struct {
    *common.BaseRequest

    // ZoneId 实例所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceChargeType 付费类型。
    // PREPAID：预付费，即包年包月 POSTPAID：后付费
    // 默认只支持预付费计费方式， 如果需要后付费， 请联系Support开通后付费权限。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // InstanceChargePrepaid 预付费模式。
    // 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长等属性。若指定实例的付费模式为预付费则该参数必传。
    InstanceChargePrepaid *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`

    // InstanceTypeId 实例机型ID。
    // 具体取值可通过调用接口DescribeInstanceTypes来获得最新的规格表。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

    // ImageId 指定有效的镜像ID。
    // 可通过以下方式获取可用的镜像ID：通过调用接口 DescribeImages
    //                 ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的ImageId字段。
    //                 也可以不指定镜像，如果不指定镜像，后续可以通过IPMI进行安装。使用iPXE安装镜像，请指定ipxeUrl字段，且该字段不必传。
    ImageId *string `json:"imageId,omitempty"`

    // ResourceGroupId 实例所在的资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // InstanceName 实例显示名称。
    // 不得超过64个字符。仅支持输入字母、数字、-和英文句点(.)。
    // 购买多台实例，可以指定模式串[begin_number,bits]。begin_number：有序数值的起始值，取值支持[0,99999]，默认值为0。bits：有序数值所占的位数，取值支持[1,6]，默认值为6。注意模式串中不得有空格。购买1台时，例如server_[3,3]实例显示为server003；购买2台时，实例显示名分别为server003，server004。支持指定多个模式串，如server_[3,3]_[1,1]。
    //                 
    // 默认值为 instance。
    InstanceName *string `json:"instanceName,omitempty"`

    // Hostname 实例的主机名。
    // 不得超过64个字符。仅支持输入字母、数字、-和英文句点(.) 。
    // 购买多台实例，可以指定模式串[begin_number,bits]。begin_number：有序数值的起始值，取值支持[0,99999]，默认值为0。bits：有序数值所占的位数，取值支持[1,6]，默认值为6。注意模式串中不得有空格。购买1台时，例如server_[3,3]主机名为server003；购买2台时，实例主机名分别为server003，server004。支持指定多个模式串，如server_[3,3]_[1,1]。
    //                 
    // 默认值为hostname。
    Hostname *string `json:"hostname,omitempty"`

    // Amount 指定创建实例的数量。
    Amount *int `json:"amount,omitempty"`

    // Password 实例的密码。
    // 必须是 8-16 个字符，包含大写字母、小写字母、数字和特殊字符。特殊符号可以是：1~!@$^*-_=+。该密码也是作为IPMI登录的密码。请妥善保管。
    // 密钥ID与密码必须并且只能指定其中一个。
    Password *string `json:"password,omitempty"`

    // KeyId 密钥ID。与password必须指定其中的一种。
    // 可调用接口DescribeKeyPairs来获得最新的密钥对信息。
    //                 
    // 关联密钥后，就可以通过对应的私钥来访问实例；密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。
    // 示例值：key-YWD2QFOl
    KeyId *string `json:"keyId,omitempty"`

    // InternetChargeType 网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // InternetMaxBandwidthOut 公网出带宽上限。
    // 单位：Mbps。默认值：1Mbps。不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

    // TrafficPackageSize 流量包订购大小。
    // 单位为TB。该值仅限当 internetChargeType = ByTrafficPackage 生效。
    // 如果没有传则会默认以赠送的流量包大小
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // SubnetId 虚拟子网ID 。
    // 您可以调用DescribeVpcSubnets查询已创建的交换机的相关信息。
    SubnetId *string `json:"subnetId,omitempty"`

    // RaidConfig 磁盘阵列配置。
    RaidConfig *RaidConfig `json:"raidConfig,omitempty"`

    // Partitions 分区配置。
    // 如果未安装操作系统，将不能设置分区
    Partitions []*Partition `json:"partitions,omitempty"`

    // Nic 网卡配置。
    Nic *Nic `json:"nic,omitempty"`

    // IpxeUrl iPXE URL 地址。
    // 传入参数后，将根据指定URL进行iPXE安装， 如果指定为netboot，将使用netboot iPXE方式进行安装。相关帮助文档：Deploy
    //                     a Custom Image Using iPXE
    IpxeUrl *string `json:"ipxeUrl,omitempty"`

    // UserData 用户数据。
    // 在安装实例时可以通过指定用户数据进行配置实例。当实例首次启动时，用户数据将以文本的方式传递到云服务器中，并执行该文本。支持的最大数据大小为 32KB。
    UserData *string `json:"userData,omitempty"`

    // EnablePrimaryIPv6 是否启用实例主IPv6。
    // false为不启用主IPv6，此时将不能为实例配置弹性IPv6。默认为true。
    EnablePrimaryIPv6 *bool `json:"enablePrimaryIPv6,omitempty"`

    // ClusterId 带宽组ID。当 internetChargeType = ByBandwidthCluster 时必传。
    ClusterId *string `json:"clusterId,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

    // Tags 创建实例时关联的标签。
    // 注意: 关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

    // EnableGatewayMode 是否启用网关模式。
    // 启用后，该机器上绑定的IP将作为网关模式使用，用于转发下游网络流量，且仅支持绑定EIP掩码段类IP。开启后需要进入机器完成进一步配置。若未开启则是普通路由模式。
    // 请注意：网关模式默认为不支持，如需使用请联系Console Support。
    EnableGatewayMode *bool `json:"enableGatewayMode,omitempty"`

}

// RaidConfig 实例磁盘阵列配置， 包括自定义raid的配置。
type RaidConfig struct {

    // RaidType Raid类型。
    // 该配置进行快捷raid配置，支持0, 1, 5, 10。
    // raidType和customRaids只能指定其中一个参数。
    RaidType *int `json:"raidType,omitempty"`

    // CustomRaids 自定义Raid配置。
    // 自定义磁盘进行raid的配置。
    // raidType和customRaids只能指定其中一个参数。
    CustomRaids []*CustomRaid `json:"customRaids,omitempty"`

}

// CustomRaid 进行自定义Raid配置时需要的raid级别和指定的磁盘序号。
type CustomRaid struct {

    // RaidType Raid类型。
    // 支持0, 1, 5, 10。
    RaidType *int `json:"raidType,omitempty"`

    // DiskSequence 磁盘序号。
    // 根据机型里的磁盘从1开始顺序编号。如果是多个磁盘序号，则必须连续。
    DiskSequence []int `json:"diskSequence,omitempty"`

}

// Partition 分区配置信息。包括文件类型, 分区大小等。
type Partition struct {

    // FsPath 分区盘符。
    // linux系统：必须为/开头，且第一个为系统分区必须为/。
    // windows系统：支持C~H，第一个系统分区必须指定为C。
    FsPath *string `json:"fsPath,omitempty"`

    // FsType 分区的文件类型。
    // linux系统：支持的值ext2,ext3, ext4, ext类型必须要有。
    // windows系统: 只能为NTFS。
    FsType *string `json:"fsType,omitempty"`

    // Size 分区大小。
    // 单位为GB。
    Size *int `json:"size,omitempty"`

}

// Nic 网卡的相关配置，目前包括公网和内网的网卡名称。
type Nic struct {

    // WanName 公网网卡名称。
    // 只能是数字和大小写字母，且必须以字母开头，长度限制为4-10。
    // 非高可用机型，默认的公网网卡名称为wan0。且不能为lan开头。
    // 高可用机型，默认的公网网卡名称为bond0。
    // 公网名称和内网名称不能相同。
    WanName *string `json:"wanName,omitempty"`

    // LanName 内网网卡名称。
    // 只能是数字和大小写字母，且必须以字母开头，长度限制为4-10。
    // 非高可用机型，默认的内网网卡名称为lan0。且不能为wan开头。
    // 高可用机型，默认的内网网卡名称为bond1。
    // 公网名称和内网名称不能相同。
    LanName *string `json:"lanName,omitempty"`

}

// MarketingInfo 市场营销活动相关信息
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
    // 长度限制：1～128字符。
    Key *string `json:"key,omitempty"`

    // Value 标签值。
    // 长度限制：1～128字符。
    Value *string `json:"value,omitempty"`

}

type CreateInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateInstancesResponseParams `json:"response,omitempty"`

}

type CreateInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceIdSet 实例ID列表。
    // 当通过本接口来创建实例时会返回该参数，表示一个或多个实例ID。返回实例ID列表并不代表实例创建成功，可根据 DescribeInstances 接口查询返回的dataSet中对应实例的状态来判断创建是否完成：如果实例状态由CREATING(创建中)或PENDING变为RUNNING(运行中)，则为创建成功；如果实例找不到或状态变为CREATE_FAILED，表示创建失败。
    InstanceIdSet []string `json:"instanceIdSet,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type DescribeInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID。
    // 最多支持100个ID查询。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // ZoneId 实例所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ResourceGroupId 资源组的ID。
    // 如果不传，则返回该用户可见的所有资源组内的实例。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // InstanceTypeId 实例机型ID。
    // 具体取值可通过调用接口DescribeInstanceTypes来获得最新的规格表。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

    // InternetChargeType 网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // SubnetId 虚拟子网ID 。
    // 可以调用DescribeVpcSubnets查询已创建交换机的相关信息。
    SubnetId *string `json:"subnetId,omitempty"`

    // InstanceStatus 实例状态。
    // 状态类型详见实例状态。
    InstanceStatus *string `json:"instanceStatus,omitempty"`

    // InstanceName 实例显示名称。
    // 如果该值以*结尾，则对instanceName进行模糊匹配，否则将进行精确匹配。
    InstanceName *string `json:"instanceName,omitempty"`

    // Hostname 实例的主机名。
    // 如果该值以*结尾，则对hostname进行模糊匹配，否则将进行精确匹配。
    Hostname *string `json:"hostname,omitempty"`

    // PublicIpAddresses 公网ipv4地址。
    PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`

    // PrivateIpAddresses 内网子网的ipv4地址。
    PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

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

type DescribeInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 结果集。
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

    // Hostname 实例的主机名。
    Hostname *string `json:"hostname,omitempty"`

    // InstanceTypeId 实例机型ID。
    InstanceTypeId *string `json:"instanceTypeId,omitempty"`

    // InstanceType 实例机型配置信息。
    InstanceType *InstanceType `json:"instanceType,omitempty"`

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // ImageName 镜像名称。
    ImageName *string `json:"imageName,omitempty"`

    // IpxeUrl IpxeUrl。
    IpxeUrl *string `json:"ipxeUrl,omitempty"`

    // InstanceChargeType 实例计费类型。
    // PREPAID：预付费，即包年包月。 POSTPAID：后付费。
    InstanceChargeType *string `json:"instanceChargeType,omitempty"`

    // BandwidthOutMbps 公网出口带宽。
    // 单位：Mbps。
    // 0 代表无限制，但是不会超过机型的最大上限。
    BandwidthOutMbps *int `json:"bandwidthOutMbps,omitempty"`

    // InternetChargeType 网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // Period 购买实例的时长。
    // 单位：月。
    // 后付费实例该字段为null。
    Period *int `json:"period,omitempty"`

    // PrimaryPublicIpAddress 实例的母IP。
    PrimaryPublicIpAddress *string `json:"primaryPublicIpAddress,omitempty"`

    // PrimaryPublicIPv6Address 实例的主IPv6地址。
    PrimaryPublicIPv6Address *string `json:"primaryPublicIPv6Address,omitempty"`

    // PublicIpAddresses 实例公网IPv4列表。
    // 如果机器的主IP未加入到公网组网接口，那么主IP将无法使用，且该字段也不会返回该IP。
    PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`

    // PrivateIpAddresses 实例内网IP列表。
    PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`

    // Ipv6Addresses 实例的IPv6地址。
    // 注意：此字段可能返回null，表示取不到有效值。
    Ipv6Addresses []string `json:"ipv6Addresses,omitempty"`

    // SubnetIds 实例所属的内网组网ID列表。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // CreateTime 创建时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 到期时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // ResourceGroupId 实例所属资源组的ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 实例所属资源组的名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // InstanceStatus 实例状态。
    // 状态类型详见实例状态。
    InstanceStatus *string `json:"instanceStatus,omitempty"`

    // TrafficPackageSize 流量包订购大小。
    // 单位为TB。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // RaidConfig 磁盘阵列配置。
    RaidConfig *RaidConfig `json:"raidConfig,omitempty"`

    // Partitions 分区配置。
    Partitions []*Partition `json:"partitions,omitempty"`

    // Nic 网卡配置。
    Nic *Nic `json:"nic,omitempty"`

    // AutoRenew 是否自动续费。
    // 对于预付费实例，取消订阅后，该字段值将返回 false
    AutoRenew *bool `json:"autoRenew,omitempty"`

    // KeyId 安装的SSH密钥ID。
    KeyId *string `json:"keyId,omitempty"`

    // Tags 资源关联的标签信息。
    Tags *Tags `json:"tags,omitempty"`

    // EnableGatewayMode 网关模式是否开启。
    EnableGatewayMode *bool `json:"enableGatewayMode,omitempty"`

}

// Tags 描述资源关联的标签信息。
type Tags struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

type ModifyInstancesAttributeRequest struct {
    *common.BaseRequest

    // InstanceName 实例名称。
    // 不得超过64个字符。仅支持输入字母、数字、-和英文句点(.)。
    InstanceName *string `json:"instanceName,omitempty"`

    // InstanceIds 一个或多个待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。每次请求批量实例的上限为100。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type ModifyInstancesAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type StartInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。每次请求批量实例的上限为100。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type StartInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type StopInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。每次请求批量实例的上限为100。
    InstanceIds []string `json:"instanceIds,omitempty"`

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
    // 可通过DescribeInstances接口返回值中的instanceId获取。每次请求批量实例的上限为100。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type RebootInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ReinstallInstanceRequest struct {
    *common.BaseRequest

    // ImageId 指定有效的镜像ID。
    // 可通过以下方式获取可用的镜像ID：通过调用接口 DescribeImages ，传入instanceTypeId获取当前机型支持的镜像列表，取返回信息中的imageId字段；也可以不指定镜像，如果不指定镜像，后续可以通过IPMI进行安装。使用iPXE安装镜像，请指定ipxeUrl字段，且该字段不必传。
    ImageId *string `json:"imageId,omitempty"`

    // Hostname 实例的主机名。
    // 不得超过64个字符。仅支持输入字母、数字、-和英文句点(.)。
    // 默认值为hostname。
    Hostname *string `json:"hostname,omitempty"`

    // Password 实例的密码。
    // 必须是 8-16 个字符，包含大写字母、小写字母、数字和特殊字符。特殊符号可以是：1~!@$^*-_=+。该密码也是作为IPMI登录的密码，请妥善保管。
    // 密钥与密码必须并且只能指定其中一个。
    Password *string `json:"password,omitempty"`

    // KeyId 密钥ID。与password必须指定其中的一种。
    // 可调用接口DescribeKeyPairs来获得最新的密钥对信息。
    // 关联密钥后，就可以通过对应的私钥来访问实例；密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。
    // 示例值：key-YWD2QFOl
    KeyId *string `json:"keyId,omitempty"`

    // RaidConfig 磁盘阵列配置。
    RaidConfig *RaidConfig `json:"raidConfig,omitempty"`

    // Partitions 分区配置。
    // 如果未安装操作系统，将不能设置分区。
    Partitions []*Partition `json:"partitions,omitempty"`

    // Nic 网卡的配置。
    Nic *Nic `json:"nic,omitempty"`

    // IpxeUrl iPXE URL 地址。
    // 传入参数后，将根据指定URL进行iPXE安装， 如果指定为netboot，将使用netboot iPXE方式进行安装。相关帮助文档：Deploy a Custom Image Using iPXE
    IpxeUrl *string `json:"ipxeUrl,omitempty"`

    // UserData 用户数据。
    // 在安装实例时可以通过指定用户数据进行配置实例。当实例首次启动时，用户数据将以文本的方式传递到云服务器中，并执行该文本。支持的最大数据大小为 32KB。
    UserData *string `json:"userData,omitempty"`

    // InstanceId 待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

}

type ReinstallInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyInstancesResourceGroupRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID列表。
    // 每次请求允许操作的实例数量上限是100。
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

type TerminateInstanceRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

}

type TerminateInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RenewInstanceRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

}

type RenewInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ReleaseInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。每次请求批量实例的上限为100。
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
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

    // BandwidthOutMbps 带宽大小。
    BandwidthOutMbps *int `json:"bandwidthOutMbps,omitempty"`

}

type InquiryPriceInstanceBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceInstanceBandwidthResponseParams `json:"response,omitempty"`

}

type InquiryPriceInstanceBandwidthResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // BandwidthPrice 公网带宽价格。
    // 可能有多个价格，比如流量包计费，包含包的价格和用量超出包后的价格。
    BandwidthPrice []*PriceItem `json:"bandwidthPrice,omitempty"`

}

type ModifyInstanceBandwidthRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

    // BandwidthOutMbps 带宽。
    // 范围：1~机型最大。
    BandwidthOutMbps *int `json:"bandwidthOutMbps,omitempty"`

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
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

}

type CancelInstanceBandwidthDowngradeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type InquiryPriceInstanceTrafficPackageRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

    // TrafficPackageSize 流量包大小。
    // 必须是0.05的倍数。
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
    // 可能有多个价格，比如流量包计费，包含包的价格和用量超出包后的价格。
    TrafficPackagePrice []*PriceItem `json:"trafficPackagePrice,omitempty"`

}

type ModifyInstanceTrafficPackageRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

    // TrafficPackageSize 流量包大小。
    // 必须是0.05的倍数。
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
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

}

type CancelInstanceTrafficPackageDowngradeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeInstanceInternetStatusRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

}

type DescribeInstanceInternetStatusResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceInternetStatusResponseParams `json:"response,omitempty"`

}

type DescribeInstanceInternetStatusResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceId 实例ID
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 实例名称
    InstanceName *string `json:"instanceName,omitempty"`

    // InternetMaxBandwidthOut 当前实例带宽
    InternetMaxBandwidthOut *int `json:"internetMaxBandwidthOut,omitempty"`

    // ModifiedInternetMaxBandwidthOut 实例修改带宽
    ModifiedInternetMaxBandwidthOut *int `json:"modifiedInternetMaxBandwidthOut,omitempty"`

    // ModifiedBandwidthStatus 实例带宽状态
    // Processing: 变更中
    // Enable: 可用
    // WaitToEnable: 下周期变更
    ModifiedBandwidthStatus *string `json:"modifiedBandwidthStatus,omitempty"`

    // TrafficPackageSize 当前实例流量包
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // ModifiedTrafficPackageSize 实例修改流量包
    ModifiedTrafficPackageSize *float64 `json:"modifiedTrafficPackageSize,omitempty"`

    // ModifiedTrafficPackageStatus 实例流量包状态
    // Processing: 变更中
    // Enable: 可用
    // WaitToEnable: 下周期变更
    ModifiedTrafficPackageStatus *string `json:"modifiedTrafficPackageStatus,omitempty"`

}

type DescribeInstanceTrafficRequest struct {
    *common.BaseRequest

    // StartTime 查询开始时间。
    // 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    EndTime *string `json:"endTime,omitempty"`

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type DescribeInstanceTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceTrafficResponseParams `json:"response,omitempty"`

}

type DescribeInstanceTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 带宽数据列表。
    DataList []*InstanceTrafficData `json:"dataList,omitempty"`

    // In95 入口带宽95值。
    In95 *int64 `json:"in95,omitempty"`

    // In95Time 入口带宽95值时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
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
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    Out95Time *string `json:"out95Time,omitempty"`

    // OutAvg 出口带宽平均值。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出口带宽最大值。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出口带宽最小值。
    OutMin *int64 `json:"outMin,omitempty"`

    // OutTotal 出口带宽总流量。
    OutTotal *int64 `json:"outTotal,omitempty"`

    // MaxBandwidth95ValueMbps 最大带宽95值。单位：Mbps。
    MaxBandwidth95ValueMbps *float64 `json:"maxBandwidth95ValueMbps,omitempty"`

    // TotalUnit 总流量单位。例如：B。
    TotalUnit *string `json:"totalUnit,omitempty"`

    // Unit 带宽值单位。例如：bps。
    Unit *string `json:"unit,omitempty"`

}

// InstanceTrafficData 实例带宽数据。
type InstanceTrafficData struct {

    // InternetRX 入口带宽。单位：bps。
    InternetRX *int64 `json:"internetRX,omitempty"`

    // InternetTX 出口带宽。单位：bps。
    InternetTX *int64 `json:"internetTX,omitempty"`

    // Time 数据时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    Time *string `json:"time,omitempty"`

}

type DescribeInstancesMonitorHealthRequest struct {
    *common.BaseRequest

    // InstanceIds 一个或多个待操作的实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。每次请求批量实例的上限为10。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type DescribeInstancesMonitorHealthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstancesMonitorHealthResponseParams `json:"response,omitempty"`

}

type DescribeInstancesMonitorHealthResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // MonitorHealthList 硬件状态信息列表。
    MonitorHealthList []*InstanceHealth `json:"monitorHealthList,omitempty"`

}

// InstanceHealth 实例状态信息。
type InstanceHealth struct {

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // CpuStatus CPU状态。
    // <ul><li>OK：硬件状态正常。</li><li>WARNING：硬件原生告警。</li><li>UNKNOWN：数据未采集到。</li></ul>
    CpuStatus *string `json:"cpuStatus,omitempty"`

    // DiskStatus Disk状态。
    // <ul><li>OK：硬件状态正常。</li><li>WARNING：硬件原生告警。</li><li>UNKNOWN：数据未采集到。</li></ul>
    DiskStatus *string `json:"diskStatus,omitempty"`

    // IpmiPing Ipmi IP状态。
    // <ul><li>OK：ICMP探测正常。</li><li>CRITICAL：ICMP探测失败。</li><li>UNKNOWN：数据未采集到。</li></ul>
    IpmiPing *string `json:"ipmiPing,omitempty"`

    // IpmiStatus Ipmi状态。
    // <ul><li>OK：ICMP探测正常。</li><li>WARNING：硬件原生告警。</li><li>UNKNOWN：数据未采集到。</li></ul>
    IpmiStatus *string `json:"ipmiStatus,omitempty"`

    // MemoryStatus Memory状态。
    // <ul><li>OK：硬件状态正常。</li><li>WARNING：硬件原生告警。</li><li>UNKNOWN：数据未采集到。</li></ul>
    MemoryStatus *string `json:"memoryStatus,omitempty"`

    // PsuStatus Power Supply状态。
    // <ul><li>OK：硬件状态正常。</li><li>WARNING：硬件原生告警。</li><li>UNKNOWN：数据未采集到。</li></ul>
    PsuStatus *string `json:"psuStatus,omitempty"`

    // WanPortStatus 服务器公网口连接的交换机端口的状态。
    // <ul><li>OK：硬件状态正常。</li><li>WARNING：硬件原生告警。</li><li>UNKNOWN：数据未采集到。</li></ul>
    WanPortStatus *string `json:"wanPortStatus,omitempty"`

    // FanStatus 风扇状态。
    // <ul><li>OK：硬件状态正常。</li><li>WARNING：硬件原生告警。</li><li>UNKNOWN：数据未采集到。</li></ul>
    FanStatus *string `json:"fanStatus,omitempty"`

    // ServerBrand 服务器供应商品牌。
    ServerBrand *string `json:"serverBrand,omitempty"`

    // ServerModel 服务器供应商型号。
    ServerModel *string `json:"serverModel,omitempty"`

    // CpuTemp 超微 Supermicro 对于刀片机单 CPU 的温度，范围[0-100]，为空代表取不到值。
    CpuTemp *int `json:"cpuTemp,omitempty"`

    // Cpu0Temp CPU0 的温度，范围[0-100]，为空代表取不到值。
    Cpu0Temp *int `json:"cpu0Temp,omitempty"`

    // Cpu1Temp CPU1 的温度，范围[0-100]，为空代表取不到值。
    Cpu1Temp *int `json:"cpu1Temp,omitempty"`

    // Cpu2Temp CPU2 的温度，范围[0-100]，为空代表取不到值。
    Cpu2Temp *int `json:"cpu2Temp,omitempty"`

    // InletTemp 进入服务器的空气温度，可简单认为是服务器所在机房的温度。
    InletTemp *int `json:"inletTemp,omitempty"`

    // TempUnit 温度单位，目前只有Celsius，即摄氏温度。
    TempUnit *string `json:"tempUnit,omitempty"`

}

type DescribeAvailableIpv4ResourcesRequest struct {
    *common.BaseRequest

    // ChargeType 计费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // ZoneId Cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

}

type DescribeAvailableIpv4ResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAvailableIpv4ResourcesResponseParams `json:"response,omitempty"`

}

type DescribeAvailableIpv4ResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // AvailableIpv4Resources 可用资源的集合。
    AvailableIpv4Resources []*AvailableIpv4Resource `json:"availableIpv4Resources,omitempty"`

}

// AvailableIpv4Resource 可用的Ipv4 Cidr Block资源。
type AvailableIpv4Resource struct {

    // ZoneId Cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Netmask 掩码。
    Netmask *int `json:"netmask,omitempty"`

    // SellStatus 售卖的状态。
    // <ul><li>SELL：表示实例可购买，且库存>10。</li><li>SELL_SHORTAGE: 表示可购买，但是库存<10台。</li><li>SOLD_OUT：表示实例已售罄。</li></ul>
    SellStatus *string `json:"sellStatus,omitempty"`

    // CidrType CIDR地址块的类型。
    CidrType *string `json:"cidrType,omitempty"`

}

type InquiryPriceCreateIpv4BlockRequest struct {
    *common.BaseRequest

    // ZoneId cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ChargeType 付费类型。
    // PREPAID：预付费，即包年包月
    // POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // CidrType CIDR地址块的类型。
    CidrType *string `json:"cidrType,omitempty"`

    // ChargePrepaid 预付费模式。
    // 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长等属性。若指定Cidr Block的付费模式为预付费则该参数必传。
    ChargePrepaid *ChargePrepaid `json:"chargePrepaid,omitempty"`

    // Netmask 购买的掩码。
    // 可以从DescribeAvailableIpv4Resource接口中获取可用的掩码列表。
    Netmask *int `json:"netmask,omitempty"`

    // Amount 购买的数量。
    Amount *int `json:"amount,omitempty"`

}

type InquiryPriceCreateIpv4BlockResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceCreateIpv4BlockResponseParams `json:"response,omitempty"`

}

type InquiryPriceCreateIpv4BlockResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Price Cidr Block价格。
    Price *PriceItem `json:"price,omitempty"`

}

type CreateIpv4BlockRequest struct {
    *common.BaseRequest

    // ZoneId Cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Name Cidr Block的名称。
    // 不得超过64个字符。
    Name *string `json:"name,omitempty"`

    // CidrType CIDR地址块的类型
    CidrType *string `json:"cidrType,omitempty"`

    // ChargeType 付费类型。
    // PREPAID：预付费，即包年包月
    // POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // ChargePrepaid 预付费模式。
    // 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长等属性。若指定Cidr Block的付费模式为预付费则该参数必传。
    ChargePrepaid *ChargePrepaid `json:"chargePrepaid,omitempty"`

    // Netmask 购买的掩码。
    // 可以从DescribeAvailableIpv4Resource接口中获取可用的掩码列表。
    Netmask *int `json:"netmask,omitempty"`

    // Amount 购买的数量。
    Amount *int `json:"amount,omitempty"`

    // ResourceGroupId Cidr Block所属的资源组ID。
    // 如果指定的区域内存在可用的VLAN，则会忽略该参数自动使用与VLAN相同的资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

    // Tags 创建Cidr时关联的标签。
    // 注意: 关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

type CreateIpv4BlockResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateIpv4BlockResponseParams `json:"response,omitempty"`

}

type CreateIpv4BlockResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // CidrBlockIds Cidr Block ID列表。
    CidrBlockIds []string `json:"cidrBlockIds,omitempty"`

}

type DescribeAvailableIpv6ResourcesRequest struct {
    *common.BaseRequest

    // ZoneId Cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

}

type DescribeAvailableIpv6ResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAvailableIpv6ResourcesResponseParams `json:"response,omitempty"`

}

type DescribeAvailableIpv6ResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // AvailableIpv6Resources 可用资源的集合。
    AvailableIpv6Resources []*AvailableIpv6Resource `json:"availableIpv6Resources,omitempty"`

}

// AvailableIpv6Resource 可用的Ipv6 Cidr Block资源。
type AvailableIpv6Resource struct {

    // ZoneId Cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // SellStatus 售卖的状态。
    // <ul><li>SELL：表示实例可购买，且库存>10。</li><li>SELL_SHORTAGE: 表示可购买，但是库存<10台。</li><li>SOLD_OUT：表示实例已售罄。</li></ul>
    SellStatus *string `json:"sellStatus,omitempty"`

}

type CreateIpv6BlockRequest struct {
    *common.BaseRequest

    // ZoneId Cidr Block所属的可用区 ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Name Cidr Block的名称。
    // 不得超过64个字符。
    Name *string `json:"name,omitempty"`

    // Amount 购买的数量。
    Amount *int `json:"amount,omitempty"`

    // ResourceGroupId Cidr Block所属的资源组ID。
    // 如果指定的区域内存在可用的VLAN，则会忽略该参数自动使用与VLAN相同的资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Tags 创建Cidr时关联的标签。
    // 注意: 关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

type CreateIpv6BlockResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateIpv6BlockResponseParams `json:"response,omitempty"`

}

type CreateIpv6BlockResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CidrBlockIds Cidr Block ID列表。
    CidrBlockIds []string `json:"cidrBlockIds,omitempty"`

}

type DescribeCidrBlocksRequest struct {
    *common.BaseRequest

    // CidrBlockIds Cidr Block ID。
    // 最多支持100个ID查询。
    CidrBlockIds []string `json:"cidrBlockIds,omitempty"`

    // CidrBlock CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // CidrBlockName Cidr Block名称。
    CidrBlockName *string `json:"cidrBlockName,omitempty"`

    // ZoneId Cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // CidrBlockType CIDR地址块的类型。
    CidrBlockType *string `json:"cidrBlockType,omitempty"`

    // Gateway 网关地址。
    Gateway *string `json:"gateway,omitempty"`

    // ChargeType 计费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // ResourceGroupId 资源组的ID。
    // 如果不传，则返回该用户可见的所有资源组内的Cidr Block。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeCidrBlocksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCidrBlocksResponseParams `json:"response,omitempty"`

}

type DescribeCidrBlocksResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 结果集。
    DataSet []*CidrBlockInfo `json:"dataSet,omitempty"`

}

// CidrBlockInfo CIDR 地址快的信息。
type CidrBlockInfo struct {

    // CidrBlockId Cidr Block唯一ID。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

    // CidrBlockType CIDR的类型。
    CidrBlockType *string `json:"cidrBlockType,omitempty"`

    // CidrType CIDR地址块的类型。
    CidrType *string `json:"cidrType,omitempty"`

    // CidrBlockName Cidr Block名称。
    CidrBlockName *string `json:"cidrBlockName,omitempty"`

    // ZoneId Cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // CidrBlock CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // Gateway 网关地址。
    Gateway *string `json:"gateway,omitempty"`

    // AvailableIpStart 可用IP的开头。
    AvailableIpStart *string `json:"availableIpStart,omitempty"`

    // AvailableIpEnd 可用IP的结尾。
    AvailableIpEnd *string `json:"availableIpEnd,omitempty"`

    // AvailableIpCount 可用IP的数量。
    AvailableIpCount *int `json:"availableIpCount,omitempty"`

    // InstanceIds 已绑定的实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // Status Cidr Block状态。
    Status *string `json:"status,omitempty"`

    // ChargeType 计费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // CreateTime 创建时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpireTime 到期时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    ExpireTime *string `json:"expireTime,omitempty"`

    // ResourceGroupId 所属资源组的ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 所属资源组的名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // Tags 资源关联的标签信息。
    Tags *Tags `json:"tags,omitempty"`

}

type ModifyCidrBlocksAttributeRequest struct {
    *common.BaseRequest

    // Name Cidr Block的名称。
    // 不得超过64个字符。
    Name *string `json:"name,omitempty"`

    // CidrBlockIds 一个或多个待操作的Cidr Block ID。
    // 可通过DescribeCidrBlocks接口返回值中的cidrBlockId获取。
    // 每次请求批量实例的上限为100。
    CidrBlockIds []string `json:"cidrBlockIds,omitempty"`

}

type ModifyCidrBlocksAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RenewCidrBlockRequest struct {
    *common.BaseRequest

    // CidrBlockId 待操作的Cidr Block ID。
    // 可通过DescribeCidrBlocks接口返回值中的cidrBlockId获取。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

}

type RenewCidrBlockResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type TerminateCidrBlockRequest struct {
    *common.BaseRequest

    // CidrBlockId 待操作的Cidr Block ID。
    // 可通过DescribeCidrBlocks接口返回值中的cidrBlockId获取。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

}

type TerminateCidrBlockResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ReleaseCidrBlocksRequest struct {
    *common.BaseRequest

    // CidrBlockIds 一个或多个待操作的Cidr Block ID。
    // 可通过DescribeCidrBlocks接口返回值中的cidrBlockId获取。
    // 每次请求批量实例的上限为100。
    CidrBlockIds []string `json:"cidrBlockIds,omitempty"`

}

type ReleaseCidrBlocksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeInstanceAvailableCidrBlockRequest struct {
    *common.BaseRequest

    // InstanceId 实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

    // CidrBlockType CIDR地址块的类型。
    CidrBlockType *string `json:"cidrBlockType,omitempty"`

}

type DescribeInstanceAvailableCidrBlockResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceAvailableCidrBlockResponseParams `json:"response,omitempty"`

}

type DescribeInstanceAvailableCidrBlockResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceAvailableCidrBlocks 实例可用的Cidr Block列表。
    InstanceAvailableCidrBlocks []*InstanceAvailableCidrBlock `json:"instanceAvailableCidrBlocks,omitempty"`

}

// InstanceAvailableCidrBlock 实例可用的CIDR地址块。
type InstanceAvailableCidrBlock struct {

    // CidrBlockId Cidr Block唯一ID。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

    // ZoneId Cidr Block所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // CidrBlockType CIDR的类型。
    CidrBlockType *string `json:"cidrBlockType,omitempty"`

    // CidrType CIDR地址块的类型。
    CidrType *string `json:"cidrType,omitempty"`

    // CidrBlock CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // AvailableIps 可用的IP列表。
    AvailableIps []string `json:"availableIps,omitempty"`

    // AvailableIpCount 可用的IP数量。
    AvailableIpCount *int `json:"availableIpCount,omitempty"`

}

type DescribeCidrBlockIpsRequest struct {
    *common.BaseRequest

    // InstanceId 实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

    // Ip IP。
    Ip *string `json:"ip,omitempty"`

    // CidrBlockId Cidr Block ID。
    // 可通过DescribeCidrBlocks接口返回值中的cidrBlockId获取。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

}

type DescribeCidrBlockIpsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCidrBlockIpsResponseParams `json:"response,omitempty"`

}

type DescribeCidrBlockIpsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CidrBlockIps 结果集。
    CidrBlockIps []*CidrBlockIp `json:"cidrBlockIps,omitempty"`

}

// CidrBlockIp CIDR 地址块里的IP的信息。
type CidrBlockIp struct {

    // CidrBlockId Cidr Block唯一ID。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

    // CidrBlockType CIDR的类型。
    CidrBlockType *string `json:"cidrBlockType,omitempty"`

    // Ip IP。
    Ip *string `json:"ip,omitempty"`

    // InstanceId 绑定的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // Status IP的状态。
    Status *string `json:"status,omitempty"`

}

type BindCidrBlockIpsRequest struct {
    *common.BaseRequest

    // IpBindList 待绑定的IP参数列表。
    IpBindList []*IpBindParam `json:"ipBindList,omitempty"`

    // CidrBlockId 待操作的Cidr Block ID。
    // 可通过DescribeCidrBlocks接口返回值中的cidrBlockId获取。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

}

// IpBindParam Cidr Block Ip绑定参数。
type IpBindParam struct {

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // Ip IP。
    // 可通过DescribeInstanceAvailableCidrBlock接口返回值中的availableIps获取。
    Ip *string `json:"ip,omitempty"`

}

type BindCidrBlockIpsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type UnbindCidrBlockIpsRequest struct {
    *common.BaseRequest

    // IpList 待解绑的IP列表。
    IpList []string `json:"ipList,omitempty"`

    // CidrBlockId 待操作的Cidr Block ID。
    // 可通过DescribeCidrBlocks接口返回值中的cidrBlockId获取。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

}

type UnbindCidrBlockIpsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type CreateByoipRequest struct {
    *common.BaseRequest

    // IpType IP 类型。
    IpType *string `json:"ipType,omitempty"`

    // Cidr 宣告IPv4或IPv6地址段。
    Cidr *string `json:"cidr,omitempty"`

    // Asn ASN号。
    Asn *int64 `json:"asn,omitempty"`

    // PublicVirtualInterfaceId 公网VLAN 唯一标识。
    PublicVirtualInterfaceId *string `json:"publicVirtualInterfaceId,omitempty"`

}

type CreateByoipResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateByoipResponseParams `json:"response,omitempty"`

}

// CreateByoipResponseParams 
type CreateByoipResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ByoipId 创建成功的 BYOIP ID。
    ByoipId *string `json:"byoipId,omitempty"`

    // CidrBlockId 创建成功的 CIDR ID。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

}

type InquiryPriceCreateEipAddressRequest struct {
    *common.BaseRequest

    // ZoneId EIP所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // EipChargeType 付费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    EipChargeType *string `json:"eipChargeType,omitempty"`

    // EipChargePrepaid 预付费模式。
    // 即包年包月相关参数设置。
    // 通过该参数可以指定包年包月实例的购买时长等属性。
    // 若指定实例的付费模式为预付费则该参数必传。
    EipChargePrepaid *ChargePrepaid `json:"eipChargePrepaid,omitempty"`

    // Amount 指定创建EIP的数量。
    Amount *int `json:"amount,omitempty"`

    // Netmask 购买的掩码。
    // 可以从DescribeEipAvailableResources接口中获取可用的掩码列表。
    Netmask *int `json:"netmask,omitempty"`

}

type InquiryPriceCreateEipAddressResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceCreateEipAddressResponseParams `json:"response,omitempty"`

}

type InquiryPriceCreateEipAddressResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // EipPrice EIP价格。
    EipPrice *PriceItem `json:"eipPrice,omitempty"`

}

type DescribeEipAddressesRequest struct {
    *common.BaseRequest

    // EipIds 取值可以由多个EIP ID共同组成。最多支持100个ID查询。
    EipIds []string `json:"eipIds,omitempty"`

    // EipChargeType 付费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    EipChargeType *string `json:"eipChargeType,omitempty"`

    // IpAddress IP地址。
    IpAddress *string `json:"ipAddress,omitempty"`

    // ZoneId EIP所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ResourceGroupId 资源组的ID。
    // 如果不传，则返回该用户可见的所有资源组内的EIP。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // EipStatus EIP状态。
    EipStatus *string `json:"eipStatus,omitempty"`

    // InstanceId 机器实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 机器实例名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeEipAddressesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipAddressesResponseParams `json:"response,omitempty"`

}

type DescribeEipAddressesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的EIP总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet EIP列表。
    DataSet []*EipAddress `json:"dataSet,omitempty"`

}

// EipAddress Eip信息。
type EipAddress struct {

    // EipId EIP唯一ID。
    EipId *string `json:"eipId,omitempty"`

    // ZoneId EIP所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // IpAddress IP地址。
    IpAddress *string `json:"ipAddress,omitempty"`

    // Netmask IP掩码。
    Netmask *int `json:"netmask,omitempty"`

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 实例名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // EipChargeType 付费类型。
    // PREPAID：预付费，即包年包月。POSTPAID：后付费。
    EipChargeType *string `json:"eipChargeType,omitempty"`

    // Period 购买EIP的时长。
    // 单位：月。
    // 后付费EIP该字段为null。
    Period *int `json:"period,omitempty"`

    // CreateTime 创建时间。
    // 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 到期时间。
    // 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    // 注意：后付费模式本项为null。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // EipStatus EIP状态
    EipStatus *string `json:"eipStatus,omitempty"`

    // Tags 资源关联的标签信息。
    Tags *Tags `json:"tags,omitempty"`

}

type AllocateEipAddressesRequest struct {
    *common.BaseRequest

    // ZoneId EIP所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // EipChargeType 付费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    EipChargeType *string `json:"eipChargeType,omitempty"`

    // EipChargePrepaid 预付费模式。
    // 即包年包月相关参数设置。
    // 通过该参数可以指定包年包月实例的购买时长等属性。
    // 若指定实例的付费模式为预付费则该参数必传。
    EipChargePrepaid *ChargePrepaid `json:"eipChargePrepaid,omitempty"`

    // ResourceGroupId 资源组ID。
    // 如果不指定，则会放入默认资源组。如果用户没有默认资源组权限， 则请求将会失败。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Amount 指定创建EIP的数量。
    Amount *int `json:"amount,omitempty"`

    // Netmask 购买的掩码。
    // 可以从DescribeEipAvailableResources接口中获取可用的掩码列表。
    Netmask *int `json:"netmask,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

    // Tags 创建EIP时关联的标签。
    // 注意: 关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

type AllocateEipAddressesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *AllocateEipAddressesResponseParams `json:"response,omitempty"`

}

type AllocateEipAddressesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // EipIdSet EIP 的ID列表。
    // 当通过本接口来创建EIP时会返回该参数，表示一个或多个EIP ID。返回EIP ID列表并不代表EIP创建成功，可根据 DescribeEipAddresses 接口查询对应EIP ID的状态来判断创建是否完成；如果EIP状态由CREATING(创建中)变为AVAILABLE，则为创建成功。
    EipIdSet []string `json:"eipIdSet,omitempty"`

    // OrderNumber 订单编号。
    // 当eipChargeType为PREPAID时会返回。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type ModifyEipAddressesResourceGroupRequest struct {
    *common.BaseRequest

    // EipIds 弹性IP ID列表。
    // 每次请求允许操作的IP数量上限是100。
    EipIds []string `json:"eipIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type ModifyEipAddressesResourceGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeInstanceAvailableEipResourcesRequest struct {
    *common.BaseRequest

    // InstanceId 机器实例ID 。
    // 可通过DescribeInstances接口返回的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

}

type DescribeInstanceAvailableEipResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceAvailableEipResourcesResponseParams `json:"response,omitempty"`

}

type DescribeInstanceAvailableEipResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceEipResources 实例可用EIP列表。
    InstanceEipResources []*InstanceAvailableEip `json:"instanceEipResources,omitempty"`

}

// InstanceAvailableEip EIP信息。
type InstanceAvailableEip struct {

    // EipId 一个EIP ID。
    // 可通过DescribeEipAddresses接口返回值中的eipId获取。
    EipId *string `json:"eipId,omitempty"`

    // IpAddress IP地址。
    IpAddress *string `json:"ipAddress,omitempty"`

    // Netmask IP掩码。
    Netmask *int `json:"netmask,omitempty"`

}

type AssociateEipAddressRequest struct {
    *common.BaseRequest

    // EipId 一个EIP ID。
    // 可通过DescribeEipAddresses接口返回值中的eipId获取。
    EipId *string `json:"eipId,omitempty"`

    // InstanceId 机器实例ID。
    // 可通过DescribeInstances接口返回值中的instanceId获取。
    InstanceId *string `json:"instanceId,omitempty"`

}

type AssociateEipAddressResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type UnAssociateEipAddressRequest struct {
    *common.BaseRequest

    // EipId 一个EIP ID。
    // 可通过DescribeEipAddresses接口返回值中的eipId获取。
    EipId *string `json:"eipId,omitempty"`

}

type UnAssociateEipAddressResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeEipAvailableResourcesRequest struct {
    *common.BaseRequest

    // EipChargeType 付费类型。
    // PREPAID：预付费，即包年包月。
    // POSTPAID：后付费。
    EipChargeType *string `json:"eipChargeType,omitempty"`

    // ZoneId EIP所属的可用区ID。
    // 不传则查询所有区域可用的EIP。
    ZoneId *string `json:"zoneId,omitempty"`

}

type DescribeEipAvailableResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipAvailableResourcesResponseParams `json:"response,omitempty"`

}

type DescribeEipAvailableResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // EipResources 购买EIP区域列表。
    EipResources []*EipAvailable `json:"eipResources,omitempty"`

}

// EipAvailable 购买EIP资源区域。
type EipAvailable struct {

    // ZoneId EIP所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Status EIP是否售卖。
    // 取值范围：
    // <ul><li>SELL：表示EIP可购买，且库存>10台。</li><li>SELL_SHORTAGE: 表示可购买，但是库存<10台。</li><li>SOLD_OUT：表示EIP已售罄。</li></ul>
    Status *string `json:"status,omitempty"`

    // Netmask IP掩码。
    Netmask *int `json:"netmask,omitempty"`

}

type TerminateEipAddressRequest struct {
    *common.BaseRequest

    // EipId 一个EIP ID。
    // 可通过DescribeEipAddresses接口返回值中的eipId获取。
    EipId *string `json:"eipId,omitempty"`

}

type TerminateEipAddressResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RenewEipAddressRequest struct {
    *common.BaseRequest

    // EipId 一个EIP ID。
    // 可通过DescribeEipAddresses接口返回值中的eipId获取。
    EipId *string `json:"eipId,omitempty"`

}

type RenewEipAddressResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ReleaseEipAddressesRequest struct {
    *common.BaseRequest

    // EipIds 一个或多个待操作的EIP ID。
    // 可通过DescribeEipAddresses接口返回值中的eipId获取。
    EipIds []string `json:"eipIds,omitempty"`

}

type ReleaseEipAddressesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeVpcAvailableRegionsRequest struct {
    *common.BaseRequest

    // ZoneId 所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // VpcRegionId VPC的节点ID。
    VpcRegionId *string `json:"vpcRegionId,omitempty"`

}

type DescribeVpcAvailableRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeVpcAvailableRegionsResponseParams `json:"response,omitempty"`

}

type DescribeVpcAvailableRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // VpcRegionSet VpcRegionInfo结果集。
    VpcRegionSet []*VpcRegionInfo `json:"vpcRegionSet,omitempty"`

}

// VpcRegionInfo VPC 节点的信息。
type VpcRegionInfo struct {

    // VpcRegionId VPC的节点ID。
    VpcRegionId *string `json:"vpcRegionId,omitempty"`

    // VpcRegionName VPC的节点名称。
    VpcRegionName *string `json:"vpcRegionName,omitempty"`

    // ZoneIds Zone ID 列表。
    ZoneIds []string `json:"zoneIds,omitempty"`

}

type CreateVpcRequest struct {
    *common.BaseRequest

    // VpcRegionId VPC的节点ID。
    VpcRegionId *string `json:"vpcRegionId,omitempty"`

    // CidrBlock VPC的CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // VpcName VPC的名称。
    // 范围1到64个字符。仅支持输入字母、数字、-和英文句点(.)。
    VpcName *string `json:"vpcName,omitempty"`

    // ResourceGroupId VPC所在的资源组ID。
    // 如果不指定，则会放入默认资源组， 如果用户没有默认资源组权限， 则请求将会失败。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Tags 创建VPC时关联的标签。
    // 注意: 关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

type CreateVpcResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateVpcResponseParams `json:"response,omitempty"`

}

type CreateVpcResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // VpcId VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

}

type DescribeVpcsRequest struct {
    *common.BaseRequest

    // VpcIds VPC ID。
    // 取值可以由多个VPC ID组成一个。最多支持100个ID查询。
    VpcIds []string `json:"vpcIds,omitempty"`

    // VpcName VPC的名称。
    VpcName *string `json:"vpcName,omitempty"`

    // CidrBlock VPC的CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // VpcStatus VPC的状态。
    VpcStatus *string `json:"vpcStatus,omitempty"`

    // VpcRegionId VPC的节点ID。
    VpcRegionId *string `json:"vpcRegionId,omitempty"`

    // ResourceGroupId 资源组的ID。
    // 如果不传，则返回该用户可见的所有资源组内的VPC。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeVpcsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeVpcsResponseParams `json:"response,omitempty"`

}

type DescribeVpcsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet VpcInfo结果集。
    DataSet []*VpcInfo `json:"dataSet,omitempty"`

}

// VpcInfo Vpc Info的信息。
type VpcInfo struct {

    // VpcId VPC唯一ID。
    VpcId *string `json:"vpcId,omitempty"`

    // VpcRegionId VPC的节点ID。
    VpcRegionId *string `json:"vpcRegionId,omitempty"`

    // VpcRegionName VPC的节点名称。
    VpcRegionName *string `json:"vpcRegionName,omitempty"`

    // VpcName VPC的名称。
    VpcName *string `json:"vpcName,omitempty"`

    // CidrBlock VPC的CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // CreateTime 创建时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // VpcStatus VPC的状态。
    VpcStatus *string `json:"vpcStatus,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // Tags 资源关联的标签信息。
    Tags *Tags `json:"tags,omitempty"`

}

type ModifyVpcsAttributeRequest struct {
    *common.BaseRequest

    // VpcName VPC名称。
    // 范围1到64个字符。仅支持输入字母、数字、-和英文句点(.)。
    VpcName *string `json:"vpcName,omitempty"`

    // VpcIds 一个或多个待操作的VPC ID。
    // 可通过DescribeVpcs接口返回值中的vpcId获取。
    // 每次请求批量VPC的上限为100。
    VpcIds []string `json:"vpcIds,omitempty"`

}

type ModifyVpcsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteVpcRequest struct {
    *common.BaseRequest

    // VpcId VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

}

type DeleteVpcResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeLoadBalancerZonesRequest struct {
    *common.BaseRequest

    // ChargeType 付费类型。
    // PREPAID：预付费，即包年包月 POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

}

type DescribeLoadBalancerZonesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeLoadBalancerZonesResponseParams `json:"response,omitempty"`

}

type DescribeLoadBalancerZonesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ZoneIdSet zoneId结果集。
    ZoneIdSet []string `json:"zoneIdSet,omitempty"`

}

type DescribeLoadBalancerSpecsRequest struct {
    *common.BaseRequest

    // ChargeType 付费类型。
    // PREPAID：预付费，即包年包月 POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // ZoneId 可用区的节点ID。
    ZoneId *string `json:"zoneId,omitempty"`

}

type DescribeLoadBalancerSpecsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeLoadBalancerSpecsResponseParams `json:"response,omitempty"`

}

type DescribeLoadBalancerSpecsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SpecSet 规格集合。
    SpecSet []*LoadBalancerSpec `json:"specSet,omitempty"`

}

// LoadBalancerSpec LoadBalancerSpec 信息
type LoadBalancerSpec struct {

    // SpecName 规格名称。
    SpecName *string `json:"specName,omitempty"`

    // MaxConnection 最大连接数。
    MaxConnection *int `json:"maxConnection,omitempty"`

    // Cps CPS。
    Cps *int `json:"cps,omitempty"`

    // Qps QPS。
    Qps *int `json:"qps,omitempty"`

}

type CreateLoadBalancerRequest struct {
    *common.BaseRequest

    // ClientToken 用于保证请求的幂等性。
    ClientToken *string `json:"clientToken,omitempty"`

    // ZoneId 可用区的节点ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // LoadBalancerName 负载均衡名称。
    LoadBalancerName *string `json:"loadBalancerName,omitempty"`

    // ChargeType 付费类型。
    // PREPAID：预付费，即包年包月 POSTPAID：后付费。
    ChargeType *string `json:"chargeType,omitempty"`

    // InstanceChargePrepaid 预付费模式。
    // 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长等属性。若指定实例的付费模式为预付费则该参数必传。
    InstanceChargePrepaid *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`

    // Bandwidth 带宽值。
    // 取值范围：1-1024。
    // 单位：Mbps
    Bandwidth *int `json:"bandwidth,omitempty"`

    // IpType 指定负载均衡绑定的IP的类型:
    // 可用值：
    // <ul><li>IPv4</li><li>IPv6</li></ul>
    IpType *string `json:"ipType,omitempty"`

    // VipCount 额外购买VIP数量。
    // 负载均衡会默认绑定1个指定类型的IP。
    VipCount *int `json:"vipCount,omitempty"`

    // SubnetId 子网ID。
    // 若是创建VIP4类型时，必须指定子网ID。该实例的主备IP将在该子网中获取。
    SubnetId *string `json:"subnetId,omitempty"`

    // CidrBlockId CIDR ID。
    // 创建VIP6类型时需要指定CIDR ID。如果该可用区中存在唯一一个CIDR则可以不指定，默认使用该CIDR。
    CidrBlockId *string `json:"cidrBlockId,omitempty"`

    // MasterIp 主IP。
    MasterIp *string `json:"masterIp,omitempty"`

    // BackupIp 备IP。
    BackupIp *string `json:"backupIp,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

    // Tags 创建负载均衡时关联的标签。
    // 注意: 关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

type CreateLoadBalancerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateLoadBalancerResponseParams `json:"response,omitempty"`

}

type CreateLoadBalancerResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // LoadBalancerId 负载均衡ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type DescribeLoadBalancersRequest struct {
    *common.BaseRequest

    // LoadBalancerIds 负载均衡实例ID集合。
    LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

    // ZoneId 可用区 ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // LoadBalancerName 负载均衡实例的名称。支持模糊搜索。
    LoadBalancerName *string `json:"loadBalancerName,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。默认为20，最大为1000。
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

type DescribeLoadBalancersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeLoadBalancersResponseParams `json:"response,omitempty"`

}

type DescribeLoadBalancersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的LB总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet LoadBalancer集合。
    DataSet []*LoadBalancerInfo `json:"dataSet,omitempty"`

}

// LoadBalancerInfo LoadBalancerInfo 信息
type LoadBalancerInfo struct {

    // LoadBalancerId LoadBalancerId唯一ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ZoneId LoadBalancer所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // LoadBalancerName LoadBalancer的名称。
    LoadBalancerName *string `json:"loadBalancerName,omitempty"`

    // SpecName 规格名称。
    SpecName *string `json:"specName,omitempty"`

    // VipList LoadBalancer的VIP集合。
    VipList []*LoadBalancerIp `json:"vipList,omitempty"`

    // ChargeType 计费方式。
    ChargeType *string `json:"chargeType,omitempty"`

    // Period 周期。
    Period *int `json:"period,omitempty"`

    // CreateTime 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-ddTHH:mm:ssZ`。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 过期时间。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // Status LoadBalancer的状态。
    Status *string `json:"status,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // MasterIp 主IP。
    MasterIp *string `json:"masterIp,omitempty"`

    // BackupIp 备IP。
    BackupIp *string `json:"backupIp,omitempty"`

    // IpType IPv4。
    IpType *string `json:"ipType,omitempty"`

    // Bandwidth 带宽。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // IsWorking 是否工作。
    IsWorking *bool `json:"isWorking,omitempty"`

    // ListenerList 监听器集合。
    ListenerList []*ListenerInfo `json:"listenerList,omitempty"`

    // BackendList 后端服务器集合。
    BackendList []*BackendInfo `json:"backendList,omitempty"`

    // Tags 资源关联的标签信息。
    Tags *Tags `json:"tags,omitempty"`

}

// LoadBalancerIp LoadBalancerIp 信息
type LoadBalancerIp struct {

    // VipId VIP的ID。
    VipId *string `json:"vipId,omitempty"`

    // IpAddress IP地址。
    IpAddress *string `json:"ipAddress,omitempty"`

    // Type IP类型
    // DEFAULT、EXTRA
    Type *string `json:"type,omitempty"`

    // Status 状态。
    Status *string `json:"status,omitempty"`

}

// ListenerInfo ListenerInfo 信息
type ListenerInfo struct {

    // LoadBalancerId LoadBalancerId唯一ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerId Listener的ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // ListenerName Listener的名称。
    ListenerName *string `json:"listenerName,omitempty"`

    // Status 状态。
    Status *string `json:"status,omitempty"`

    // Port 监听端口。
    Port *string `json:"port,omitempty"`

    // Protocol 监听协议。
    Protocol *string `json:"protocol,omitempty"`

    // BackendProtocol 后段服务器协议。
    BackendProtocol *string `json:"backendProtocol,omitempty"`

    // Scheduler 调度算法。
    Scheduler *string `json:"scheduler,omitempty"`

    // Kind 流量转发模式。
    Kind *string `json:"kind,omitempty"`

    // HealthCheck 健康检查。
    HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

    // Notify 通知地址。
    Notify *Notify `json:"notify,omitempty"`

    // CreateTime 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-ddTHH:mm:ssZ`。
    CreateTime *string `json:"createTime,omitempty"`

}

// HealthCheck HealthCheck 信息
type HealthCheck struct {

    // CheckEnabled 检查开关。
    CheckEnabled *bool `json:"checkEnabled,omitempty"`

    // CheckType 检查类型。
    CheckType *string `json:"checkType,omitempty"`

    // CheckConnectTimeout 连接时间。
    CheckConnectTimeout *int `json:"checkConnectTimeout,omitempty"`

    // CheckRetry 重试次数。
    CheckRetry *int `json:"checkRetry,omitempty"`

    // CheckDelayBeforeRetry 超时时间。
    CheckDelayBeforeRetry *int `json:"checkDelayBeforeRetry,omitempty"`

    // CheckIntervalTime 重试间隔。
    CheckIntervalTime *int `json:"checkIntervalTime,omitempty"`

    // CheckPort 检查端口。
    CheckPort *int `json:"checkPort,omitempty"`

    // HttpVersion Http版本。
    HttpVersion *string `json:"httpVersion,omitempty"`

    // HttpCheckPath Http地址。
    HttpCheckPath *string `json:"httpCheckPath,omitempty"`

    // HttpCheckDigest Http摘要认证。
    HttpCheckDigest *int `json:"httpCheckDigest,omitempty"`

    // HttpCode 状态码。
    HttpCode *int `json:"httpCode,omitempty"`

    // MiscCheckPath Misc地址。
    MiscCheckPath *string `json:"miscCheckPath,omitempty"`

    // MiscTimeout Misc超时时间。
    MiscTimeout *int `json:"miscTimeout,omitempty"`

}

// Notify Notify 信息
type Notify struct {

    // Enable 通知开关。
    Enable *bool `json:"enable,omitempty"`

    // ApiAddress 地址。
    ApiAddress *string `json:"apiAddress,omitempty"`

}

// BackendInfo BackendInfo 信息
type BackendInfo struct {

    // ListenerId Listener的ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // BackendId Backend的ID。
    BackendId *string `json:"backendId,omitempty"`

    // BackendName Backend的名称。
    BackendName *string `json:"backendName,omitempty"`

    // Status 状态。
    Status *string `json:"status,omitempty"`

    // Port 端口。
    Port *string `json:"port,omitempty"`

    // Weight 权重。
    Weight *int `json:"weight,omitempty"`

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 实例名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // CreateTime 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-ddTHH:mm:ssZ`。
    CreateTime *string `json:"createTime,omitempty"`

}

type ModifyLoadBalancersNameRequest struct {
    *common.BaseRequest

    // LoadBalancerName 负载均衡名称。
    // 不得超过64个字符。仅支持输入字母、数字、-和英文句点(.)。
    LoadBalancerName *string `json:"loadBalancerName,omitempty"`

    // LoadBalancerIds 负载均衡ID集合。
    LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

}

type ModifyLoadBalancersNameResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type TerminateLoadBalancerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例的ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type TerminateLoadBalancerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RestoreLoadBalancerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例的ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type RestoreLoadBalancerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *RestoreLoadBalancerResponseParams `json:"response,omitempty"`

}

type RestoreLoadBalancerResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type ReleaseLoadBalancerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例的ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type ReleaseLoadBalancerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type CreateListenerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例 ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerName 监听器名称。
    // 不得超过64个字符。仅支持输入字母、数字、-和英文句点(.)。
    ListenerName *string `json:"listenerName,omitempty"`

    // PortList 监听端口。取值范围：0~65535。
    // 示例值：["6569"]
    PortList []int `json:"portList,omitempty"`

    // ClientToken 用于保证请求的幂等性。
    ClientToken *string `json:"clientToken,omitempty"`

    // Protocol 监听协议。取值：
    // <ul><li>TCP</li><li>UDP</li></ul>
    Protocol *string `json:"protocol,omitempty"`

    // BackendProtocol 后端转发协议。取值：
    // <ul><li>TCP</li><li>UDP</li></ul>
    BackendProtocol *string `json:"backendProtocol,omitempty"`

    // Scheduler 调度算法。取值：
    // <ul><li>wrr（默认值）：加权轮询，权重值越高的后端服务器，被轮询到的概率也越高。</li><li>rr：轮询，按照访问顺序依次将外部请求分发到后端服务器。</li></ul>
    Scheduler *string `json:"scheduler,omitempty"`

    // Kind 工作模式。取值：
    // <ul><li>FNAT(全局网络地址转换)</li><li>DR(直接路由)</li></ul>
    Kind *string `json:"kind,omitempty"`

    // HealthCheck 健康检查的信息。
    HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

    // Notify 警报。
    Notify *Notify `json:"notify,omitempty"`

}

type CreateListenerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateListenerResponseParams `json:"response,omitempty"`

}

type CreateListenerResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ListenerId 监听器的ID。
    ListenerId *string `json:"listenerId,omitempty"`

}

type DescribeListenersRequest struct {
    *common.BaseRequest

    // ListenerIds 监听器ID列表。
    // 最多支持100个。
    ListenerIds []string `json:"listenerIds,omitempty"`

    // LoadBalancerIds 负载均衡实例ID列表。
    // 最多支持100个。
    LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

    // ListenerName 监听器名称。
    // 支持模糊搜索。
    ListenerName *string `json:"listenerName,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

}

type DescribeListenersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeListenersResponseParams `json:"response,omitempty"`

}

type DescribeListenersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的Listener总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 监听器集合。
    DataSet []*ListenerInfo `json:"dataSet,omitempty"`

}

type ModifyListenerAttributeRequest struct {
    *common.BaseRequest

    // ListenerId 监听器的ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // ListenerName 监听器名称。
    // 不得超过64个字符。仅支持输入字母、数字、-和英文句点(.)。
    ListenerName *string `json:"listenerName,omitempty"`

    // Protocol 监听协议。取值：
    // <ul><li>TCP</li><li>UDP</li></ul>
    Protocol *string `json:"protocol,omitempty"`

    // BackendProtocol 后端转发协议。取值：
    // <ul><li>TCP</li><li>UDP</li></ul>
    BackendProtocol *string `json:"backendProtocol,omitempty"`

    // Scheduler 调度算法。取值：
    // <ul><li>wrr（默认值）：加权轮询，权重值越高的后端服务器，被轮询到的概率也越高。</li><li>rr：轮询，按照访问顺序依次将外部请求分发到后端服务器。</li></ul>
    Scheduler *string `json:"scheduler,omitempty"`

    // Kind 工作模式。取值：
    // <ul><li>NAT</li><li>DR</li></ul>
    Kind *string `json:"kind,omitempty"`

    // HealthCheck 健康检查。
    HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

    // Notify 警报。
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

    // ListenerId 监听器的ID。
    ListenerId *string `json:"listenerId,omitempty"`

}

type DeleteListenerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type CreateLoadBalancerVIPsRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡的ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // Count 额外创建IP数量。
    Count *int `json:"count,omitempty"`

}

type CreateLoadBalancerVIPsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateLoadBalancerVIPsResponseParams `json:"response,omitempty"`

}

type CreateLoadBalancerVIPsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单号。若LB是VIP6类型则返回空
    OrderNumber *string `json:"orderNumber,omitempty"`

    // VipIdSet IP ID集合。
    VipIdSet []string `json:"vipIdSet,omitempty"`

}

type DeleteLoadBalancerVIPRequest struct {
    *common.BaseRequest

    // VipId 负责均衡IP的ID。
    VipId *string `json:"vipId,omitempty"`

}

type DeleteLoadBalancerVIPResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeBackendsRequest struct {
    *common.BaseRequest

    // ListenerId 监听器的ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // BackendIds 后端配置服务器的ID集合。
    BackendIds []string `json:"backendIds,omitempty"`

    // BackendName 后端配置服务器的名称。
    BackendName *string `json:"backendName,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

}

type DescribeBackendsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBackendsResponseParams `json:"response,omitempty"`

}

type DescribeBackendsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的后端配置服务器总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 后端配置服务器集合。
    DataSet []*BackendInfo `json:"dataSet,omitempty"`

}

type RegisterBackendRequest struct {
    *common.BaseRequest

    // ListenerId 监听器的ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // BackendName 后端配置服务器名称。
    // 不得超过64个字符。仅支持输入字母、数字、-和英文句点(.)。
    BackendName *string `json:"backendName,omitempty"`

    // InstanceId 裸金属实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // ClientToken 用于保证请求的幂等性。
    ClientToken *string `json:"clientToken,omitempty"`

    // PortList 端口列表。
    PortList []int `json:"portList,omitempty"`

    // Weight 权重。
    Weight *int `json:"weight,omitempty"`

}

type RegisterBackendResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *RegisterBackendResponseParams `json:"response,omitempty"`

}

type RegisterBackendResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // BackendId 后端配置服务器的ID。
    BackendId *string `json:"backendId,omitempty"`

}

type DeregisterBackendRequest struct {
    *common.BaseRequest

    // BackendId 后端配置服务器的ID。
    BackendId *string `json:"backendId,omitempty"`

}

type DeregisterBackendResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeManagedInstancesRequest 
type DescribeManagedInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID。
    // 取值可以由多个实例ID组成一个。
    // 最多支持100个ID查询。
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
    // 默认值：20。
    PageSize *int `json:"pageSize,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type DescribeManagedInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeManagedInstancesResponseParams `json:"response,omitempty"`

}

// DescribeManagedInstancesResponseParams 
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

// DescribeManagedInstanceTrafficRequest 
type DescribeManagedInstanceTrafficRequest struct {
    *common.BaseRequest

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // StartTime 查询开始时间。
    // 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    EndTime *string `json:"endTime,omitempty"`

}

type DescribeManagedInstanceTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeManagedInstanceTrafficResponseParams `json:"response,omitempty"`

}

type DescribeManagedInstanceTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 带宽数据列表。
    DataList []*InstanceTrafficData `json:"dataList,omitempty"`

    // In95 入口带宽95值。
    In95 *int64 `json:"in95,omitempty"`

    // In95Time 入口带宽95值时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
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
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    Out95Time *string `json:"out95Time,omitempty"`

    // OutAvg 出口带宽平均值。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出口带宽最大值。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出口带宽最小值。
    OutMin *int64 `json:"outMin,omitempty"`

    // OutTotal 出口带宽总流量。
    OutTotal *int64 `json:"outTotal,omitempty"`

    // MaxBandwidth95ValueMbps 最大带宽95值。单位：Mbps。
    MaxBandwidth95ValueMbps *float64 `json:"maxBandwidth95ValueMbps,omitempty"`

    // TotalUnit 总流量单位。例如：B。
    TotalUnit *string `json:"totalUnit,omitempty"`

    // Unit 带宽值单位。例如：bps。
    Unit *string `json:"unit,omitempty"`

}

