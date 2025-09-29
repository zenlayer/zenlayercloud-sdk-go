package zec

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeDisksRequest 查询云盘列表的请求参数。
type DescribeDisksRequest struct {
    *common.BaseRequest

    // DiskIds 根据云盘ID列表筛选。
    DiskIds []string `json:"diskIds,omitempty"`

    // DiskName 根据云盘名称筛选，该字段支持模糊搜索。
    DiskName *string `json:"diskName,omitempty"`

    // DiskStatus 根据云盘的状态进行筛选。
    DiskStatus *string `json:"diskStatus,omitempty"`

    // DiskType 根据云盘的类型进行筛选。
    DiskType *string `json:"diskType,omitempty"`

    // DiskCategory 根据云盘的分类进行筛选。
    DiskCategory *string `json:"diskCategory,omitempty"`

    // InstanceId 根据云盘挂载的实例ID进行筛选。
    InstanceId *string `json:"instanceId,omitempty"`

    // ZoneId 根据云盘所在的可用区进行筛选。
    ZoneId *string `json:"zoneId,omitempty"`

    // PageNum 返回的分页大小。默认为20，最大为1000。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页数。默认为1。
    PageSize *int `json:"pageSize,omitempty"`

    // RegionId 根据云盘所在的节点ID进行筛选。
    RegionId *string `json:"regionId,omitempty"`

    // SnapshotAbility 根据云盘是否有快照能力进行筛选。
    SnapshotAbility *bool `json:"snapshotAbility,omitempty"`

    // ResourceGroupId 根据快照所属的资源组进行筛选。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

// DescribeDisksResponseParams 描述快照列表的响应信息。
type DescribeDisksResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int64 `json:"totalCount,omitempty"`

    // DataSet 云盘的结果集。
    DataSet []*DiskInfo `json:"dataSet,omitempty"`

}

// DiskInfo 描述云盘的基本信息。
type DiskInfo struct {

    // DiskId 云盘的 ID。
    DiskId *string `json:"diskId,omitempty"`

    // DiskName 云盘的名称。
    DiskName *string `json:"diskName,omitempty"`

    // RegionId 云盘所在的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // ZoneId 云盘所在节点的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // DiskType 云盘的类型。
    DiskType *string `json:"diskType,omitempty"`

    // Portable 是否可卸载。
    Portable *bool `json:"portable,omitempty"`

    // DiskCategory 云盘的类别。
    DiskCategory *string `json:"diskCategory,omitempty"`

    // DiskSize 云盘的大小。单位：GiB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskStatus 云盘的状态。
    DiskStatus *string `json:"diskStatus,omitempty"`

    // InstanceId 云盘绑定实例的ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 云盘绑定实例的名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 到期时间。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // Period 周期。
    Period *int `json:"period,omitempty"`

    // ResourceGroupId 云盘所属的资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 云盘所属的资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // Serial 云盘序号。可能为null，表示取不到值。
    Serial *string `json:"serial,omitempty"`

    // SnapshotAbility 是否具体快照能力。
    SnapshotAbility *bool `json:"snapshotAbility,omitempty"`

    // AutoSnapshotPolicyId 云盘关联的自动快照策略ID。
    AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty"`

}

type DescribeDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeDisksResponseParams `json:"response,omitempty"`

}

type DescribeDiskRegionsRequest struct {
    *common.BaseRequest

}

// DescribeDiskRegionsResponseParams 支持售卖云硬盘的节点响应结果。
type DescribeDiskRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // RegionIds 支持售卖云硬盘的节点ID列表。
    RegionIds []string `json:"regionIds,omitempty"`

}

type DescribeDiskRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeDiskRegionsResponseParams `json:"response,omitempty"`

}

// InquiryPriceCreateDisksRequest 创建云硬盘询价的请求参数。
type InquiryPriceCreateDisksRequest struct {
    *common.BaseRequest

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // DiskSize 云硬盘大小。单位：GiB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskAmount 云硬盘数量。
    DiskAmount *int `json:"diskAmount,omitempty"`

    // DiskCategory 云硬盘种类。Basic NVMe SSD: 经济型 NVMe SSD。Standard NVMe SSD: 标准型 NVMe SSD。默认为Standard NVMe SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

}

// InquiryPriceCreateDisksResponseParams 云硬盘的询价响应值。
type InquiryPriceCreateDisksResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataDiskPrice 云硬盘的价格。
    DataDiskPrice *PriceItem `json:"dataDiskPrice,omitempty"`

}

// PriceItem 描述价格的信息。
type PriceItem struct {

    // Discount 折扣大小。如80.0代表8折。
    Discount *float64 `json:"discount,omitempty"`

    // DiscountPrice 后付费的单元折后价格。后付费模式使用，如果价格为阶梯价格，该项为null。
    DiscountPrice *float64 `json:"discountPrice,omitempty"`

    // OriginalPrice 预付费的原价。预付费模式使用，后付费该值为 null。
    OriginalPrice *float64 `json:"originalPrice,omitempty"`

    // UnitPrice 后付费的单元原始价格。后付费模式使用，如果价格为阶梯价格，该项为null。
    UnitPrice *float64 `json:"unitPrice,omitempty"`

    // DiscountUnitPrice 后付费的单元折后价格。后付费模式使用，如果价格为阶梯价格，该项为null。
    DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`

    // ChargeUnit 后付费计价单元。后付费模式使用，可取值范围：<br/>HOUR: 表示计价单元是按每小时来计算。DAY: 表示计价单元是按天来计算。MONTH: 表示计价单元是按月来计算，95计费则是这种。
    ChargeUnit *string `json:"chargeUnit,omitempty"`

    // StepPrices 后付费阶梯价格。后付费模式使用，如果非阶梯价格，该项为null。
    StepPrices []*StepPrice `json:"stepPrices,omitempty"`

    // AmountUnit 用量单位。比如Mbps, LCU等。如果为null, 代表取不到值。
    AmountUnit *string `json:"amountUnit,omitempty"`

    // ExcessUnitPrice 超量原始价格。
    ExcessUnitPrice *float64 `json:"excessUnitPrice,omitempty"`

    // ExcessDiscountUnitPrice 超量折扣后价格。
    ExcessDiscountUnitPrice *float64 `json:"excessDiscountUnitPrice,omitempty"`

    // ExcessAmountUnit 超量用量单位。如果为null, 代表取不到值。
    ExcessAmountUnit *string `json:"excessAmountUnit,omitempty"`

}

// StepPrice 描述阶梯价格的信息。
type StepPrice struct {

    // StepStart 阶梯的起始值。
    StepStart *float64 `json:"stepStart,omitempty"`

    // StepEnd 阶梯的到达值。为null代表最后一级阶梯。
    StepEnd *float64 `json:"stepEnd,omitempty"`

    // UnitPrice 阶梯单价。
    UnitPrice *float64 `json:"unitPrice,omitempty"`

    // DiscountUnitPrice 阶梯折后价。
    DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`

}

type InquiryPriceCreateDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceCreateDisksResponseParams `json:"response,omitempty"`

}

// CreateDisksRequest 创建云硬盘的请求信息。
type CreateDisksRequest struct {
    *common.BaseRequest

    // ZoneId 云硬盘所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // DiskName 云盘名称。范围1到64个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    DiskName *string `json:"diskName,omitempty"`

    // DiskSize 云硬盘大小，单位GiB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskAmount 需要创建的云硬盘的数量。
    DiskAmount *int `json:"diskAmount,omitempty"`

    // InstanceId 云硬盘挂在的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // ResourceGroupId 云硬盘所在的资源组ID。如不指定则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // DiskCategory 云硬盘种类。Basic NVMe SSD: 经济型 NVMe SSD。Standard NVMe SSD: 标准型 NVMe SSD。默认为Standard NVMe SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

    // SnapshotId 使用快照ID进行创建。如果传入则根据此快照创建云硬盘，快照的云盘类型必须为数据盘快照。
    SnapshotId *string `json:"snapshotId,omitempty"`

    // MarketingOptions 市场营销的相关选项。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

}

// MarketingInfo 描述市场活动的相关信息。
type MarketingInfo struct {

    // DiscountCode 使用市场发放的折扣码。如果折扣码不存在，最终折扣将不会生效。
    DiscountCode *string `json:"discountCode,omitempty"`

    // UsePocVoucher 是否使用POC代金券。 如果系统不存在POC代金券，相关创建流程会失败。
    UsePocVoucher *bool `json:"usePocVoucher,omitempty"`

}

// CreateDisksResponseParams 创建云硬盘的响应信息。
type CreateDisksResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DiskIds 创建的云硬盘ID列表。
    DiskIds []string `json:"diskIds,omitempty"`

    // OrderNumber 本次创建对应的订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type CreateDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateDisksResponseParams `json:"response,omitempty"`

}

// AttachDisksRequest 云硬盘挂在到实例的请求信息。
type AttachDisksRequest struct {
    *common.BaseRequest

    // DiskIds 需要挂载的云硬盘ID列表。
    DiskIds []string `json:"diskIds,omitempty"`

    // InstanceId 被挂载的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type AttachDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DetachDisksRequest 云硬盘从主机实例上卸载的请求参数。
type DetachDisksRequest struct {
    *common.BaseRequest

    // DiskIds 要卸载的云盘ID列表。
    DiskIds []string `json:"diskIds,omitempty"`

    // InstanceCheckFlag 是否检测实例的运行状态。 默认为true，即实例关机才允许被卸载。否则必须实例关机才能调用本接口。
    InstanceCheckFlag *bool `json:"instanceCheckFlag,omitempty"`

}

type DetachDisksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyDisksAttributesRequest 修改云盘属性的请求参数。
type ModifyDisksAttributesRequest struct {
    *common.BaseRequest

    // DiskIds 需要修改的云盘ID列表。
    DiskIds []string `json:"diskIds,omitempty"`

    // DiskName 云盘名称。范围1到64个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    DiskName *string `json:"diskName,omitempty"`

}

type ModifyDisksAttributesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ResizeDiskRequest struct {
    *common.BaseRequest

    // DiskId 云硬盘ID。通过DescribeDisks接口查询。
    DiskId *string `json:"diskId,omitempty"`

    // DiskSize 云硬盘扩容后的大小。单位GiB。必须大于当前云硬盘大小。云盘最大限制为32768GB(32TB)。
    DiskSize *int `json:"diskSize,omitempty"`

}

type ResizeDiskResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ReleaseDiskRequest 删除云硬盘的请求参数。
type ReleaseDiskRequest struct {
    *common.BaseRequest

    // DiskId 要删除的云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

}

type ReleaseDiskResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// RenewDiskRequest 恢复云硬盘的请求参数。
type RenewDiskRequest struct {
    *common.BaseRequest

    // DiskId 要恢复的云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

}

type RenewDiskResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyDisksResourceGroupRequest 修改云盘资源组的请求参数。
type ModifyDisksResourceGroupRequest struct {
    *common.BaseRequest

    // DiskIds 要迁移资源组的云盘ID列表。
    DiskIds []string `json:"diskIds,omitempty"`

    // ResourceGroupId 目标资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type ModifyDisksResourceGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeDiskCategoryRequest 查询云硬盘支持的类型的请求信息。
type DescribeDiskCategoryRequest struct {
    *common.BaseRequest

    // ZoneId 根据可用区ID筛选。
    ZoneId *string `json:"zoneId,omitempty"`

    // DiskCategory 根据云硬盘种类筛选。Basic NVMe SSD: 经济型 NVMe SSD。Standard NVMe SSD: 标准型 NVMe SSD。默认为Standard NVMe SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

}

// DescribeDiskCategoryResponseParams 获取某个区域支持的云盘类型响应结果。
type DescribeDiskCategoryResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CategoryZoneSet 可用区支持的云盘类型。
    CategoryZoneSet []*DescribeDiskCategoryItem `json:"categoryZoneSet,omitempty"`

}

// DescribeDiskCategoryItem 描述可用区支持的云盘类型的信息。
type DescribeDiskCategoryItem struct {

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // CategorySet 支持的云硬盘类型。
    CategorySet []string `json:"categorySet,omitempty"`

}

type DescribeDiskCategoryResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeDiskCategoryResponseParams `json:"response,omitempty"`

}

// CreateVpcRequest 创建全球VPC的请求参数。
type CreateVpcRequest struct {
    *common.BaseRequest

    // Name VPC的名称。范围2到63个字符。仅支持输入字母、数字、-和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

    // CidrBlock VPC的CIDR地址段。需要满足以下3种内网段内(10.0.0.0/8, 172.16.0.0/12 and 192.168.0.0/16)。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // Mtu VPC的传输单元 MTU。支持：1300、1500、9000。
    Mtu *int `json:"mtu,omitempty"`

    // EnablePriIpv6 是否开启内网IPv6。一旦开启，后续无法关闭。默认为关闭。
    EnablePriIpv6 *bool `json:"enablePriIpv6,omitempty"`

    // ResourceGroupId VPC所在的资源组ID。如果不指定资源组，则会放到默认的资源组中。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

// CreateVpcResponseParams 创建全球VPC的响应结果。
type CreateVpcResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // VpcId 创建的VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

}

type CreateVpcResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateVpcResponseParams `json:"response,omitempty"`

}

// DeleteVpcRequest 删除VPC的响应信息。
type DeleteVpcRequest struct {
    *common.BaseRequest

    // VpcId 要删除的VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

}

type DeleteVpcResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyVpcsAttributeRequest 批量修改VPC属性的请求参数
type ModifyVpcsAttributeRequest struct {
    *common.BaseRequest

    // VpcIds 需要修改的VPC ID列表。
    VpcIds []string `json:"vpcIds,omitempty"`

    // Name VPC的名称。范围2到63个字符。仅支持输入字母、数字、-和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

}

type ModifyVpcsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeSubnetRegionsRequest 查询支持创建子网区域的请求信息。
type DescribeSubnetRegionsRequest struct {
    *common.BaseRequest

    // RegionIds 根据节点ID过滤。
    RegionIds []string `json:"regionIds,omitempty"`

}

// DescribeSubnetRegionsResponseParams 查询支持创建子网区域的响应信息
type DescribeSubnetRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // RegionSet 支持子网的节点信息。
    RegionSet []*RegionInfo `json:"regionSet,omitempty"`

}

// RegionInfo 节点信息。
type RegionInfo struct {

    // RegionId 节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // RegionName 节点名称。
    RegionName *string `json:"regionName,omitempty"`

    // RegionTitle 节点标题。
    RegionTitle *string `json:"regionTitle,omitempty"`

    // EnablePubIpv6 是否支持IPv6。
    EnablePubIpv6 *bool `json:"enablePubIpv6,omitempty"`

}

type DescribeSubnetRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSubnetRegionsResponseParams `json:"response,omitempty"`

}

// CreateSubnetRequest 创建子网请求参数。
type CreateSubnetRequest struct {
    *common.BaseRequest

    // VpcId 需要添加子网的VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

    // Name 子网名称。范围2到63个字符。仅支持输入字母、数字、-和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

    // RegionId 子网所在的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // StackType 子网的IP堆栈类型。
    StackType *string `json:"stackType,omitempty"`

    // CidrBlock 子网的IPv4 CIDR地址段。如果指定堆栈类型`stackType` 包含 `IPv4`, 则该字段必填。指定的CIDR地址段必须属于VPC的CIDR范围内。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // Ipv6Type IPv6的类型。如果指定堆栈类型`stackType` 包含 `IPv6`, 则该字段必填。
    Ipv6Type *string `json:"ipv6Type,omitempty"`

}

// CreateSubnetResponseParams 创建子网的响应结果。
type CreateSubnetResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SubnetId 创建的子网ID。
    SubnetId *string `json:"subnetId,omitempty"`

}

type CreateSubnetResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateSubnetResponseParams `json:"response,omitempty"`

}

// DeleteSubnetRequest 删除子网的请求参数。
type DeleteSubnetRequest struct {
    *common.BaseRequest

    // SubnetId 要删除的子网ID。
    SubnetId *string `json:"subnetId,omitempty"`

}

type DeleteSubnetResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifySubnetAttributeRequest 描述修改子网属性的请求参数。
type ModifySubnetAttributeRequest struct {
    *common.BaseRequest

    // SubnetId 子网的ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // SubnetName 子网的名称。仅支持输入字母、数字、-和英文句点(.)。且必须以数字或字母开头和结尾。
    SubnetName *string `json:"subnetName,omitempty"`

    // CidrBlock 需要修改的IPv4 CIDR Block。仅支持有IPv4堆栈类型的子网。
    CidrBlock *string `json:"cidrBlock,omitempty"`

}

type ModifySubnetAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifySubnetStackTypeRequest 切换子网堆栈类型的请求参数。
type ModifySubnetStackTypeRequest struct {
    *common.BaseRequest

    // SubnetId 要操作的子网的ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // StackType 子网堆栈类型。目前只支持`IPv4_IPv6`。
    StackType *string `json:"stackType,omitempty"`

    // Ipv6Type IPv6的类型。
    Ipv6Type *string `json:"ipv6Type,omitempty"`

}

// ModifySubnetStackTypeResponseParams 修改子网堆栈类型的响应结果。
type ModifySubnetStackTypeResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Ipv6CidrBlock 分配的IPv6地址段。
    Ipv6CidrBlock *string `json:"ipv6CidrBlock,omitempty"`

}

type ModifySubnetStackTypeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ModifySubnetStackTypeResponseParams `json:"response,omitempty"`

}

type DescribeVpcsRequest struct {
    *common.BaseRequest

    // VpcIds VPC的ID列表。最多可传100个ID。
    VpcIds []string `json:"vpcIds,omitempty"`

    // Name 根据VPC名称过滤。支持模糊查询。
    Name *string `json:"name,omitempty"`

    // CidrBlock 根据VPC的IPv4 CIDR过滤。支持模糊查询。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // ResourceGroupId 根据资源组ID过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

// DescribeVpcsResponseParams 描述Vpc列表的响应结果。
type DescribeVpcsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 匹配筛选条件的vpc总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 返回分页的vpc集合数据。
    DataSet []*VpcInfo `json:"dataSet,omitempty"`

}

// VpcInfo 描述VPC的基本信息。
type VpcInfo struct {

    // VpcId VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

    // Name VPC的名称。
    Name *string `json:"name,omitempty"`

    // CidrBlock VPC的IPv4 CIDR。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // Ipv6CidrBlock VPC的内网IPv6 CIDR。如果为null,说明未开启IPv6。
    Ipv6CidrBlock *string `json:"ipv6CidrBlock,omitempty"`

    // Mtu mtu值。
    Mtu *int `json:"mtu,omitempty"`

    // IsDefault 是否为默认VPC。
    IsDefault *bool `json:"isDefault,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // UsageIpv4Count VPC里已使用IPv4数量。
    UsageIpv4Count *int `json:"usageIpv4Count,omitempty"`

    // UsageIpv6Count VPC里已使用IPv6数量。
    UsageIpv6Count *int `json:"usageIpv6Count,omitempty"`

    // SecurityGroupId 关联的安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // ResourceGroup VPC关联的资源组信息。
    ResourceGroup *ResourceGroupInfo `json:"resourceGroup,omitempty"`

}

// ResourceGroupInfo 描述资源所在资源组的相关信息，包括资源组名称和ID。
type ResourceGroupInfo struct {

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

}

type DescribeVpcsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeVpcsResponseParams `json:"response,omitempty"`

}

// ModifyVpcAttributeRequest 修改VPC的属性请求参数。
type ModifyVpcAttributeRequest struct {
    *common.BaseRequest

    // VpcId VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

    // VpcName VPC的名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    VpcName *string `json:"vpcName,omitempty"`

    // CidrBlock 需要修改的IPv4 CIDR。需要满足属于(10.0.0.0/8, 172.16.0.0/12 and 192.168.0.0/16)范围内。如果VPC存在子网，则修改的CIDR范围必须包含原VPC CIDR。默认VPC不支持修改。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // EnableIPv6 是否开启IPv6内网CIDR。当前仅允许打开(`true`)，一旦设置IPv6, 将无法关闭。
    EnableIPv6 *bool `json:"enableIPv6,omitempty"`

    // SecurityGroupId 修改VPC绑定的安全组ID。如果不指定，则不会修改。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type ModifyVpcAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeSubnetsRequest 查询子网列表的请求信息。
type DescribeSubnetsRequest struct {
    *common.BaseRequest

    // SubnetIds 根据子网的ID进行筛选。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // Name 根据子网的名称进行筛选。该字段支持模糊搜索。
    Name *string `json:"name,omitempty"`

    // CidrBlock 根据子网的CIDR进行筛选。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // RegionId 根据子网所在的节点进行筛选。
    RegionId *string `json:"regionId,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

}

// DescribeSubnetsResponseParams 子网列表的响应信息。
type DescribeSubnetsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 子网的结果集。
    DataSet []*SubnetInfo `json:"dataSet,omitempty"`

}

// SubnetInfo 描述子网的基本信息。
type SubnetInfo struct {

    // SubnetId 子网的ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // RegionId 子网所在节点的ID。
    RegionId *string `json:"regionId,omitempty"`

    // Name 子网的名称。
    Name *string `json:"name,omitempty"`

    // CidrBlock 子网的CIDR地址。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // Ipv6CidrBlock 子网的IPv6 CIDR地址段。如果子网的IP堆栈类型不包括V6,该字段取不到值。
    Ipv6CidrBlock *string `json:"ipv6CidrBlock,omitempty"`

    // StackType 子网的IP堆栈类型。
    StackType *string `json:"stackType,omitempty"`

    // Ipv6Type 子网上IPv6类型。如果子网的IP堆栈类型不包括V6,该字段取不到值。
    Ipv6Type *string `json:"ipv6Type,omitempty"`

    // VpcId 子网所属VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

    // VpcName 子网所属VPC的名称。
    VpcName *string `json:"vpcName,omitempty"`

    // UsageIpv4Count 子网已使用IPv4数量。
    UsageIpv4Count *int `json:"usageIpv4Count,omitempty"`

    // UsageIpv6Count 子网已使用IPv6数量。
    UsageIpv6Count *int `json:"usageIpv6Count,omitempty"`

    // CreateTime 子网的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // IsDefault 子网是否为默认。
    IsDefault *bool `json:"isDefault,omitempty"`

}

type DescribeSubnetsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSubnetsResponseParams `json:"response,omitempty"`

}

// ModifySubnetsAttributeRequest 批量修改子网的请求参数。
type ModifySubnetsAttributeRequest struct {
    *common.BaseRequest

    // SubnetIds 需要修改的子网ID列表。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // Name 修改的子网名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

}

type ModifySubnetsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// CreateSnapshotRequest 描述创建快照的请求信息。
type CreateSnapshotRequest struct {
    *common.BaseRequest

    // DiskId 云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

    // SnapshotName 快照名称。
    SnapshotName *string `json:"snapshotName,omitempty"`

    // RetentionTime 保留的到期时间。格式为：yyyy-MM-ddTHH:mm:ssZ。如果不传，则代表永久保留。指定时间必须在当前时间24小时后。
    RetentionTime *string `json:"retentionTime,omitempty"`

}

// CreateSnapshotResponseParams 创建快照的响应信息。
type CreateSnapshotResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SnapshotId 创建的快照ID。
    SnapshotId *string `json:"snapshotId,omitempty"`

}

type CreateSnapshotResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateSnapshotResponseParams `json:"response,omitempty"`

}

type ModifySnapshotsAttributeRequest struct {
    *common.BaseRequest

    // SnapshotIds 快照ID列表。
    SnapshotIds []string `json:"snapshotIds,omitempty"`

    // SnapshotName 快照名称。
    SnapshotName *string `json:"snapshotName,omitempty"`

    // RetentionTime 快照过期时间。格式为：yyyy-MM-ddTHH:mm:ssZ。如果改成永久保留，请设置`isPermanent`=`true`，如果设置该时间必须设置为当前时间后24小时。
    RetentionTime *string `json:"retentionTime,omitempty"`

    // IsPermanent 该定期快照策略创建的快照是否永久保留。
    IsPermanent *bool `json:"isPermanent,omitempty"`

}

type ModifySnapshotsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteSnapshotsRequest struct {
    *common.BaseRequest

    // SnapshotIds 快照ID列表。
    SnapshotIds []string `json:"snapshotIds,omitempty"`

}

// DeleteSnapshotsResponseParams 删除快照的响应信息。
type DeleteSnapshotsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SnapshotIds 操作失败的快照ID。
    SnapshotIds []string `json:"snapshotIds,omitempty"`

}

type DeleteSnapshotsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DeleteSnapshotsResponseParams `json:"response,omitempty"`

}

// DescribeSnapshotsRequest 描述查询快照列表的请求信息。
type DescribeSnapshotsRequest struct {
    *common.BaseRequest

    // SnapshotIds 根据快照ID列表进行过滤。
    SnapshotIds []string `json:"snapshotIds,omitempty"`

    // ZoneId 快照所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Status 根据快照的状态过滤。
    Status *string `json:"status,omitempty"`

    // DiskIds 按照快照所属的Disk ID列表 过滤。
    DiskIds []string `json:"diskIds,omitempty"`

    // DiskType 根据快照的云盘类型过滤。
    DiskType *string `json:"diskType,omitempty"`

    // SnapshotType 根据快照类型过滤。
    SnapshotType *string `json:"snapshotType,omitempty"`

    // SnapshotName 根据快照显示名称过滤。该字段支持模糊搜索。
    SnapshotName *string `json:"snapshotName,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // ResourceGroupId 根据资源组ID过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

// DescribeSnapshotsResponseParams 描述快照列表的响应信息。
type DescribeSnapshotsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 满足过滤条件的快照总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 返回的快照列表数据。
    DataSet []*SnapshotInfo `json:"dataSet,omitempty"`

}

// SnapshotInfo 描述快照的信息。
type SnapshotInfo struct {

    // SnapshotId 快照唯一ID。
    SnapshotId *string `json:"snapshotId,omitempty"`

    // SnapshotName 快照显示名称。
    SnapshotName *string `json:"snapshotName,omitempty"`

    // ZoneId 快照所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Status 快照的状态。
    Status *string `json:"status,omitempty"`

    // SnapshotType 快照的类型。
    SnapshotType *string `json:"snapshotType,omitempty"`

    // RetentionTime 快照的保留到期时间。如果取不到值，说明快照为永久保留。
    RetentionTime *string `json:"retentionTime,omitempty"`

    // DiskId 云盘ID。
    DiskId *string `json:"diskId,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // DiskAbility 是否具备创建disk的能力。
    DiskAbility *bool `json:"diskAbility,omitempty"`

    // ResourceGroup 所属的资源组信息。
    ResourceGroup *ResourceGroupInfo `json:"resourceGroup,omitempty"`

}

type DescribeSnapshotsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSnapshotsResponseParams `json:"response,omitempty"`

}

type ApplySnapshotRequest struct {
    *common.BaseRequest

    // SnapshotId 快照唯一ID。
    SnapshotId *string `json:"snapshotId,omitempty"`

    // DiskId 云硬盘ID。
    DiskId *string `json:"diskId,omitempty"`

}

type ApplySnapshotResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// CreateAutoSnapshotPolicyRequest 描述创建自动快照策略的请求信息。
type CreateAutoSnapshotPolicyRequest struct {
    *common.BaseRequest

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // AutoSnapshotPolicyName 自动快照策略的名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    AutoSnapshotPolicyName *string `json:"autoSnapshotPolicyName,omitempty"`

    // RepeatWeekDays 自动快照的重复日期。单位为天，周期为星期，例如 1 表示周一。
    RepeatWeekDays []int `json:"repeatWeekDays,omitempty"`

    // Hours 指定定期快照策略的触发时间。使用 UTC 时间，单位为小时。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。
    Hours []int `json:"hours,omitempty"`

    // RetentionDays 自动快照的保留时间，单位为天。如果该值设置-1，则代表永久保留。默认为永久保存。取值范围：-1或[1,65535]。
    RetentionDays *int `json:"retentionDays,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

// CreateAutoSnapshotPolicyResponseParams 创建自动快照策略的响应信息。
type CreateAutoSnapshotPolicyResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // AutoSnapshotPolicyId 自动快照策略的ID。
    AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty"`

}

type CreateAutoSnapshotPolicyResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateAutoSnapshotPolicyResponseParams `json:"response,omitempty"`

}

type ApplyAutoSnapshotPolicyRequest struct {
    *common.BaseRequest

    // AutoSnapshotPolicyId 自动快照策略ID。
    AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty"`

    // DiskIds 要添加的磁盘ID列表。
    DiskIds []string `json:"diskIds,omitempty"`

}

type ApplyAutoSnapshotPolicyResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type CancelAutoSnapshotPolicyRequest struct {
    *common.BaseRequest

    // AutoSnapshotPolicyId 自动快照策略ID。
    AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty"`

    // DiskIds 要移除的磁盘ID列表。
    DiskIds []string `json:"diskIds,omitempty"`

}

type CancelAutoSnapshotPolicyResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeAutoSnapshotPoliciesRequest 自动快照策略列表的请求参数。
type DescribeAutoSnapshotPoliciesRequest struct {
    *common.BaseRequest

    // AutoSnapshotPolicyIds 根据自动快照策略的ID进行过滤。
    AutoSnapshotPolicyIds []string `json:"autoSnapshotPolicyIds,omitempty"`

    // ZoneIds 根据自动快照策略的可用区ID进行过滤。
    ZoneIds []string `json:"zoneIds,omitempty"`

    // AutoSnapshotPolicyName 根据自动快照策略的名称进行过滤。该字段支持模糊搜索。
    AutoSnapshotPolicyName *string `json:"autoSnapshotPolicyName,omitempty"`

    // ResourceGroupId 根据资源组ID进行过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

}

// DescribeAutoSnapshotPoliciesResponseParams 查询自动快照策略的响应值。
type DescribeAutoSnapshotPoliciesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataSet 查询的自动快照策略数据结果。
    DataSet []*AutoSnapshotPolicy `json:"dataSet,omitempty"`

    // TotalCount 符合匹配的总数量。
    TotalCount *int `json:"totalCount,omitempty"`

}

// AutoSnapshotPolicy 描述自动快照策略的相关信息。
type AutoSnapshotPolicy struct {

    // AutoSnapshotPolicyId 自动快照策略ID。
    AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty"`

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // AutoSnapshotPolicyName 自动快照策略的名称。
    AutoSnapshotPolicyName *string `json:"autoSnapshotPolicyName,omitempty"`

    // RepeatWeekDays 自动快照的重复日期。单位为天，周期为星期，例如 1 表示周一。
    RepeatWeekDays []int `json:"repeatWeekDays,omitempty"`

    // Hours 指定定期快照策略的触发时间。使用 UTC 时间，单位为小时。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。
    Hours []int `json:"hours,omitempty"`

    // RetentionDays 自动快照的保留时间，单位为天。如果该值设置-1，则代表永久保留。取值范围：-1或[1,65535]。
    RetentionDays *int `json:"retentionDays,omitempty"`

    // DiskNum 关联的云盘数量。
    DiskNum *int `json:"diskNum,omitempty"`

    // CreateTime 策略创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ResourceGroup 资源组信息。
    ResourceGroup *ResourceGroupInfo `json:"resourceGroup,omitempty"`

    // DiskIdSet 关联的云盘ID。
    DiskIdSet []string `json:"diskIdSet,omitempty"`

}

type DescribeAutoSnapshotPoliciesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAutoSnapshotPoliciesResponseParams `json:"response,omitempty"`

}

type ModifyAutoSnapshotPolicyRequest struct {
    *common.BaseRequest

    // AutoSnapshotPolicyId 自动快照策略的ID。
    AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty"`

    // AutoSnapshotPolicyName 自动快照策略的名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    AutoSnapshotPolicyName *string `json:"autoSnapshotPolicyName,omitempty"`

    // RepeatWeekDays 自动快照的重复日期。单位为天，周期为星期，例如 1 表示周一。
    RepeatWeekDays []int `json:"repeatWeekDays,omitempty"`

    // Hours 指定定期快照策略的触发时间。使用 UTC 时间，单位为小时。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。
    Hours []int `json:"hours,omitempty"`

    // RetentionDays 自动快照的保留时间，单位为天。如果该值设置-1，则代表永久保留。取值范围：-1或[1,65535]。
    RetentionDays *int `json:"retentionDays,omitempty"`

}

type ModifyAutoSnapshotPolicyResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteAutoSnapshotPoliciesRequest struct {
    *common.BaseRequest

    // AutoSnapshotPolicyIds 要删除的自动快照策略ID列表。
    AutoSnapshotPolicyIds []string `json:"autoSnapshotPolicyIds,omitempty"`

}

type DeleteAutoSnapshotPoliciesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeNetworkInterfaceRegionsRequest struct {
    *common.BaseRequest

}

// DescribeNetworkInterfaceRegionsResponseParams 支持网卡的区域信息。
type DescribeNetworkInterfaceRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // RegionIds 支持网卡的节点ID。
    RegionIds []string `json:"regionIds,omitempty"`

}

type DescribeNetworkInterfaceRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeNetworkInterfaceRegionsResponseParams `json:"response,omitempty"`

}

type DescribeNetworkInterfacesRequest struct {
    *common.BaseRequest

    // NicIds 根据网卡ID过滤。最多传入100个ID。
    NicIds []string `json:"nicIds,omitempty"`

    // Name 根据网卡所属的名称过滤。该字段支持模糊搜索。
    Name *string `json:"name,omitempty"`

    // RegionId 根据网卡所属的节点 ID过滤。
    RegionId *string `json:"regionId,omitempty"`

    // VpcId 根据网卡所属的VPC ID过滤。
    VpcId *string `json:"vpcId,omitempty"`

    // SubnetId 根据网卡所属的子网ID过滤。
    SubnetId *string `json:"subnetId,omitempty"`

    // InstanceId 网卡关联的实例ID过滤。
    InstanceId *string `json:"instanceId,omitempty"`

    // Status 根据网卡的状态过滤。
    Status *string `json:"status,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // NicType 根据网卡的类型筛选过滤。
    NicType *string `json:"nicType,omitempty"`

    // ResourceGroupId 根据网卡所属的资源组ID过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // SecurityGroupId 根据网卡所属的安全组ID过滤。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

// DescribeNetworkInterfacesResponseParams 网卡列表的响应结果。
type DescribeNetworkInterfacesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 网卡的结果集。
    DataSet []*NicInfo `json:"dataSet,omitempty"`

}

// NicInfo 描述网卡的相关信息。
type NicInfo struct {

    // NicId 网卡的ID。
    NicId *string `json:"nicId,omitempty"`

    // Name 网卡的名称。
    Name *string `json:"name,omitempty"`

    // Status 网卡状态。
    Status *string `json:"status,omitempty"`

    // NicType 网卡类型。
    NicType *string `json:"nicType,omitempty"`

    // RegionId 所属节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // NicSubnetType 网卡的IP堆栈类型。
    NicSubnetType *string `json:"nicSubnetType,omitempty"`

    // PublicIpList 关联的公网IP。
    PublicIpList []string `json:"publicIpList,omitempty"`

    // PrivateIpList 关联的内网IP。
    PrivateIpList []string `json:"privateIpList,omitempty"`

    // PrimaryIpv4 主的内网IPv4地址。
    PrimaryIpv4 *string `json:"primaryIpv4,omitempty"`

    // PrimaryIpv6 网卡上的主IPv6地址。如果堆栈类型是V4,该值取值为空。
    PrimaryIpv6 *string `json:"primaryIpv6,omitempty"`

    // Ipv6Cidr 网卡上的IPv6地址。如果堆栈类型是V4,该值取值为空。
    Ipv6Cidr *string `json:"ipv6Cidr,omitempty"`

    // SecondaryIpv4s 网卡上的辅助 IPv4 地址。
    SecondaryIpv4s []string `json:"secondaryIpv4s,omitempty"`

    // MacAddress 网卡的MAC地址。
    MacAddress *string `json:"macAddress,omitempty"`

    // InstanceId 所绑定的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // VpcId 所关联VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

    // SubnetId 所关联的子网ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // CreateTime 网卡的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // UpdateTime 网卡的更新时间。
    UpdateTime *string `json:"updateTime,omitempty"`

    // ResourceGroup 网卡所属的资源组信息。
    ResourceGroup *ResourceGroupInfo `json:"resourceGroup,omitempty"`

    // SecurityGroupId 网卡关联的安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type DescribeNetworkInterfacesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeNetworkInterfacesResponseParams `json:"response,omitempty"`

}

// ModifyNetworkInterfaceAttributeRequest 修改网卡属性的请求参数。
type ModifyNetworkInterfaceAttributeRequest struct {
    *common.BaseRequest

    // NicId 需要修改的网卡ID。
    NicId *string `json:"nicId,omitempty"`

    // Name 名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

    // SecurityGroupId 修改网卡绑定的目标安全组ID。目前一张网卡只能关联一个安全组。指定该字段会解绑网卡原来的安全组。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type ModifyNetworkInterfaceAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyNetworkInterfacesAttributeRequest 修改网卡的属性的请求参数。
type ModifyNetworkInterfacesAttributeRequest struct {
    *common.BaseRequest

    // NicIds 需要修改的网卡ID列表。
    NicIds []string `json:"nicIds,omitempty"`

    // Name 名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

}

type ModifyNetworkInterfacesAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// CreateNetworkInterfaceRequest 创建辅助网卡的请求参数。
type CreateNetworkInterfaceRequest struct {
    *common.BaseRequest

    // Name 网卡名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

    // SubnetId 子网的ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // NicStackType 网卡的IP堆栈类型。如果不指定，会根据子网堆栈类型决定：如果子网是V4,则网卡为V4，如果子网是V6,则网卡为V6。如果网卡要V4&V6，请指定该参数。
    NicStackType *string `json:"nicStackType,omitempty"`

    // ResourceGroupId 网卡创建所在的资源组ID，如不指定则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // SecurityGroupId 指定安全组ID。目前一个网卡只能关联1个安全组。如果未指定，会默认用VPC关联下的安全组。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // InternetChargeType 公网IPv6的网络计费方式。当子网的堆栈类型包括V6且为公网时，需要指定。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // Bandwidth 公网IPv6的带宽限速。单位Mbps。当子网的堆栈类型包括V6且为公网时，需要指定。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // PackageSize 公网IPv6的流量包大小。单位为TB。值要求为0或0.1的倍数。当子网的堆栈类型包括V6且为公网时，且网络计费方式是流量计费(`ByTrafficPackage`)需要指定。
    PackageSize *float64 `json:"packageSize,omitempty"`

    // ClusterId 公网IPv6所指定的共享带宽包ID。当子网的堆栈类型包括V6且为公网时，且网络计费方式是共享带宽包计费(`BandwidthCluster`)需要指定。
    ClusterId *string `json:"clusterId,omitempty"`

    // MarketingOptions 市场营销相关的选项。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

}

// CreateNetworkInterfaceResponseParams 创建辅助网卡的响应结果。
type CreateNetworkInterfaceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // NicId 创建的辅助网卡ID。
    NicId *string `json:"nicId,omitempty"`

}

type CreateNetworkInterfaceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateNetworkInterfaceResponseParams `json:"response,omitempty"`

}

// DeleteNetworkInterfaceRequest 删除网卡的请求参数。
type DeleteNetworkInterfaceRequest struct {
    *common.BaseRequest

    // NicId 要删除的网卡ID。
    NicId *string `json:"nicId,omitempty"`

}

type DeleteNetworkInterfaceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// AttachNetworkInterfaceRequest 网卡绑定实例的请求参数。
type AttachNetworkInterfaceRequest struct {
    *common.BaseRequest

    // NicId 需要操作的网卡ID。
    NicId *string `json:"nicId,omitempty"`

    // InstanceId 需要绑定的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type AttachNetworkInterfaceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DetachNetworkInterfaceRequest 解绑网卡的请求参数。
type DetachNetworkInterfaceRequest struct {
    *common.BaseRequest

    // NicId 需要操作的网卡ID。
    NicId *string `json:"nicId,omitempty"`

}

type DetachNetworkInterfaceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// AssignNetworkInterfaceIpv4Request 网卡绑定内网IPv4的请求参数。
type AssignNetworkInterfaceIpv4Request struct {
    *common.BaseRequest

    // NicId 需要绑定的网卡ID。
    NicId *string `json:"nicId,omitempty"`

    // IpAddress 绑定的内网IP地址。该地址必须属于子网的CIDR内，且未被使用。
    IpAddress *string `json:"ipAddress,omitempty"`

}

type AssignNetworkInterfaceIpv4Response struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type BatchAssignNetworkInterfaceIpv4Request struct {
    *common.BaseRequest

    // NicId 要操作的网卡ID。
    NicId *string `json:"nicId,omitempty"`

    // IpAddresses 内网IP地址列表。IP地址必须在网卡所属子网CIDR范围内，且不能是网关/广播/网络地址。
    IpAddresses []string `json:"ipAddresses,omitempty"`

    // IpAddressCount 指定额外绑定的内网IP数量。 该字段和`ipAddresses`必须指定其一，如果都指定，则会以 `ipAddresses` 有效。
    IpAddressCount *int `json:"ipAddressCount,omitempty"`

}

type BatchAssignNetworkInterfaceIpv4ResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // IpAddresses 绑定的内网IP地址。
    IpAddresses []string `json:"ipAddresses,omitempty"`

}

type BatchAssignNetworkInterfaceIpv4Response struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *BatchAssignNetworkInterfaceIpv4ResponseParams `json:"response,omitempty"`

}

type UnassignNetworkInterfaceIpv4Request struct {
    *common.BaseRequest

    // NicId 网卡ID。
    NicId *string `json:"nicId,omitempty"`

    // IpAddress 需要解绑的IPv4地址。该字段已过时，请使用`ipAddresses`，如果两者均指定， 则以`ipAddresses`为准。
    IpAddress *string `json:"ipAddress,omitempty"`

    // IpAddresses 需要解绑的内网IPv4地址。
    IpAddresses []string `json:"ipAddresses,omitempty"`

}

type UnassignNetworkInterfaceIpv4Response struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeNetworkInterfacePublicIPv6Request struct {
    *common.BaseRequest

    // NicId 网卡ID。
    NicId *string `json:"nicId,omitempty"`

}

type DescribeNetworkInterfacePublicIPv6ResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Address 公网IPv6信息。网卡如果没有公网IPv6,则取值为空。
    Address *PublicIpv6CidrAddress `json:"address,omitempty"`

}

// PublicIpv6CidrAddress 公网IPv6的基本信息。
type PublicIpv6CidrAddress struct {

    // Ipv6CidrId IPv6 CIDR的ID。
    Ipv6CidrId *string `json:"ipv6CidrId,omitempty"`

    // Ipv6Cidr IPv6 CIDR的地址。
    Ipv6Cidr *string `json:"ipv6Cidr,omitempty"`

    // PrimaryIpv6Address 网卡的主IPv6地址。
    PrimaryIpv6Address *string `json:"primaryIpv6Address,omitempty"`

    // InternetChargeType IPv6的网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // Bandwidth IPv6的公网带宽限速。单位：Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // TrafficPackageSize IPv6的流量包大小。单位：TB。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // BandwidthCluster 关联的带宽组信息。
    BandwidthCluster *BandwidthClusterInfo `json:"bandwidthCluster,omitempty"`

}

// BandwidthClusterInfo 描述带宽组的基本信息。
type BandwidthClusterInfo struct {

    // BandwidthClusterId 共享带宽包ID。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

    // BandwidthClusterName 共享带宽包名称。
    BandwidthClusterName *string `json:"bandwidthClusterName,omitempty"`

}

type DescribeNetworkInterfacePublicIPv6Response struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeNetworkInterfacePublicIPv6ResponseParams `json:"response,omitempty"`

}

// InquiryPricePublicIpv6Request 公网IPv6流量包或固定带宽询价的请求参数。
type InquiryPricePublicIpv6Request struct {
    *common.BaseRequest

    // RegionId 节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // PackageSize 流量包订购大小。单位为TB。
    PackageSize *float64 `json:"packageSize,omitempty"`

    // Bandwidth 公网出带宽上限。单位：Mbps。不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
    Bandwidth *int `json:"bandwidth,omitempty"`

}

// InquiryPricePublicIpv6ResponseParams 公网Ipv6流量包或固定带宽询价的响应信息。
type InquiryPricePublicIpv6ResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // BandwidthPrice 带宽的价格。
    BandwidthPrice *PriceItem `json:"bandwidthPrice,omitempty"`

}

type InquiryPricePublicIpv6Response struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPricePublicIpv6ResponseParams `json:"response,omitempty"`

}

// AssignNetworkInterfaceIpv6Request 网卡添加IPv6的请求信息。
type AssignNetworkInterfaceIpv6Request struct {
    *common.BaseRequest

    // NicId 要添加IPv6的网卡ID。
    NicId *string `json:"nicId,omitempty"`

    // InternetChargeType 公网IPv6的网络计费方式。当子网的堆栈类型包括V6且为公网时，需要指定。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // Bandwidth 公网IPv6的带宽限速。单位Mbps。当子网的堆栈类型包括V6且为公网时，需要指定。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // PackageSize 公网IPv6的流量包大小。单位为TB。值要求为0或0.1的倍数。当子网的堆栈类型包括V6且为公网时，且网络计费方式是流量计费(`ByTrafficPackage`)需要指定。
    PackageSize *float64 `json:"packageSize,omitempty"`

    // ClusterId 公网IPv6所指定的共享带宽包ID。当子网的堆栈类型包括V6且为公网时，且网络计费方式是共享带宽包计费(`BandwidthCluster`)需要指定。
    ClusterId *string `json:"clusterId,omitempty"`

}

type AssignNetworkInterfaceIpv6Response struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeImagesRequest 描述镜像的请求信息
type DescribeImagesRequest struct {
    *common.BaseRequest

    // ZoneId 查询镜像所在的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ImageIds 镜像ID列表。
    ImageIds []string `json:"imageIds,omitempty"`

    // ImageName 根据镜像名称过滤。该字段支持模糊搜索。
    ImageName *string `json:"imageName,omitempty"`

    // Category 镜像所属分类。
    Category *string `json:"category,omitempty"`

    // ImageType 镜像类型。
    ImageType *string `json:"imageType,omitempty"`

    // OsType 操作系统类型。
    OsType *string `json:"osType,omitempty"`

    // ImageStatus 镜像的状态。
    ImageStatus *string `json:"imageStatus,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

}

// DescribeImagesResponseParams 描述镜像列表的响应信息。
type DescribeImagesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 镜像列表数据。
    DataSet []*Image `json:"dataSet,omitempty"`

}

// Image 描述镜像的基本信息。
type Image struct {

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // ImageName 镜像的名称。
    ImageName *string `json:"imageName,omitempty"`

    // ImageType 镜像的类型。
    ImageType *string `json:"imageType,omitempty"`

    // ImageSize 镜像的大小。
    ImageSize *string `json:"imageSize,omitempty"`

    // ImageDescription 镜像描述信息。
    ImageDescription *string `json:"imageDescription,omitempty"`

    // ImageVersion 镜像的版本。
    ImageVersion *string `json:"imageVersion,omitempty"`

    // ImageStatus 镜像的状态。
    ImageStatus *string `json:"imageStatus,omitempty"`

    // NicNetworkType 镜像支持的网卡类型。
    NicNetworkType []string `json:"nicNetworkType,omitempty"`

    // Category 镜像的分类。
    Category *string `json:"category,omitempty"`

    // OsType 操作系统类型。
    OsType *string `json:"osType,omitempty"`

}

type DescribeImagesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeImagesResponseParams `json:"response,omitempty"`

}

// DescribeSecurityGroupsRequest 查询安全组列表的请求参数。
type DescribeSecurityGroupsRequest struct {
    *common.BaseRequest

    // SecurityGroupIds 根据安全组ID列表筛选。
    SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

    // SecurityGroupName 根据安全组名称筛选。支持模糊搜索。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

}

// DescribeSecurityGroupsResponseParams 查询安全组列表的响应结果。
type DescribeSecurityGroupsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 安全组的结果数据。
    DataSet []*SecurityGroupInfo `json:"dataSet,omitempty"`

}

// SecurityGroupInfo 描述安全组的基本信息。
type SecurityGroupInfo struct {

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // SecurityGroupName 安全组名称。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // Scope 范围。目前只有全球范围(`Global`)。
    Scope *string `json:"scope,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // VpcIds 关联的VPC ID列表。
    VpcIds []string `json:"vpcIds,omitempty"`

    // IsDefault 是否是默认安全组。
    IsDefault *bool `json:"isDefault,omitempty"`

    // NicIdList 关联安全组的网卡ID列表。
    NicIdList []string `json:"nicIdList,omitempty"`

    // NatIdList 关联安全组的NAT网关ID列表。
    NatIdList []string `json:"natIdList,omitempty"`

}

type DescribeSecurityGroupsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSecurityGroupsResponseParams `json:"response,omitempty"`

}

// ModifySecurityGroupsAttributeRequest 修改安全组的属性
type ModifySecurityGroupsAttributeRequest struct {
    *common.BaseRequest

    // SecurityGroupName 安全组名称。范围2到63个字符。仅支持输入字母、数字、-和英文句点(.)。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // SecurityGroupIds 要操作的安全组ID列表。
    SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

}

type ModifySecurityGroupsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeSecurityGroupRuleRequest struct {
    *common.BaseRequest

    // SecurityGroupId 安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

// DescribeSecurityGroupRuleResponseParams 查询安全组规则的响应结果
type DescribeSecurityGroupRuleResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // IngressRuleList 入方向规则列表。
    IngressRuleList []*SecurityGroupRuleInfo `json:"ingressRuleList,omitempty"`

    // EgressRuleList 出方向规则列表。
    EgressRuleList []*SecurityGroupRuleInfo `json:"egressRuleList,omitempty"`

}

// SecurityGroupRuleInfo 安全组规则信息。包括出入方向、端口范围、IP协议等信息。
type SecurityGroupRuleInfo struct {

    // Direction 规则方向。ingress: 入方向。egress：出方向。
    Direction *string `json:"direction,omitempty"`

    // Policy 设置访问权限。accept：接受访问。deny: 拒绝访问。
    Policy *string `json:"policy,omitempty"`

    // Priority 规则优先级。
    Priority *int `json:"priority,omitempty"`

    // IpProtocol 传输层协议。取值大小写敏感, 取值范围：<br/>tcp：TCP协议。udp：UDP协议。icmp：ICMP协议。all：支持所有协议。
    IpProtocol *string `json:"ipProtocol,omitempty"`

    // PortRange 目的端安全组开放的传输层协议相关的端口范围。取值范围：<br/> TCP/UDP协议：取值范围为1~65535。ICMP协议：-1。all：-1。
    PortRange *string `json:"portRange,omitempty"`

    // CidrIp 源端IP地址范围。支持CIDR格式和IPv4格式的IP地址范围。默认值：0.0.XX.XX/0。
    CidrIp *string `json:"cidrIp,omitempty"`

    // Desc 备注,长度在255个以内。
    Desc *string `json:"desc,omitempty"`

}

type DescribeSecurityGroupRuleResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeSecurityGroupRuleResponseParams `json:"response,omitempty"`

}

// CreateSecurityGroupRequest 创建安全组的请求信息。
type CreateSecurityGroupRequest struct {
    *common.BaseRequest

    // Scope 范围。目前只支持`Global`。
    Scope *string `json:"scope,omitempty"`

    // SecurityGroupName 安全组名称。范围1到64个字符。仅支持输入字母、数字、-和英文句点(.)。
    SecurityGroupName *string `json:"securityGroupName,omitempty"`

    // RuleInfos 安全组的规则。
    RuleInfos []*SecurityGroupRuleInfo `json:"ruleInfos,omitempty"`

}

// CreateSecurityGroupResponseParams 创建安全组的响应信息。
type CreateSecurityGroupResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SecurityGroupId 创建的安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type CreateSecurityGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateSecurityGroupResponseParams `json:"response,omitempty"`

}

type DeleteSecurityGroupRequest struct {
    *common.BaseRequest

    // SecurityGroupId 要删除的安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type DeleteSecurityGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ConfigureSecurityGroupRulesRequest 配置安全组规则的请求参数。
type ConfigureSecurityGroupRulesRequest struct {
    *common.BaseRequest

    // SecurityGroupId 要配置的安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // RuleInfos 需要配置的安全组规则列表。注意：配置为全量覆盖。
    RuleInfos []*SecurityGroupRuleInfo `json:"ruleInfos,omitempty"`

}

type ConfigureSecurityGroupRulesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type AssignSecurityGroupVpcRequest struct {
    *common.BaseRequest

    // VpcId 要操作的VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

    // SecurityGroupId 要更换的目标安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type AssignSecurityGroupVpcResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// UnAssignSecurityGroupVpcRequest VPC解绑安全组的请求参数。
type UnAssignSecurityGroupVpcRequest struct {
    *common.BaseRequest

    // VpcId 要解绑的VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

}

type UnAssignSecurityGroupVpcResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeEipRegionsRequest struct {
    *common.BaseRequest

}

// DescribeEipRegionsResponseParams 查询支持售卖EIP的节点响应结果。
type DescribeEipRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // RegionIds 支持售卖EIP的节点ID。
    RegionIds []string `json:"regionIds,omitempty"`

}

type DescribeEipRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipRegionsResponseParams `json:"response,omitempty"`

}

// DescribeEipInternetChargeTypesRequest 查询EIP支持的IP网络类型的请求信息。
type DescribeEipInternetChargeTypesRequest struct {
    *common.BaseRequest

    // RegionId 查询的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // EipV4Type EIP IP线路类型。
    EipV4Type *string `json:"eipV4Type,omitempty"`

}

// DescribeEipInternetChargeTypesResponseParams 查询EIP支持的网络计费模式的响应值。
type DescribeEipInternetChargeTypesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InternetChargeTypes IP支持的网络计费方式。
    InternetChargeTypes []string `json:"internetChargeTypes,omitempty"`

}

type DescribeEipInternetChargeTypesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipInternetChargeTypesResponseParams `json:"response,omitempty"`

}

// DescribeEipRemoteRegionsRequest 
type DescribeEipRemoteRegionsRequest struct {
    *common.BaseRequest

    // RegionId 查询的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // EipV4Type EIP IP线路类型。
    EipV4Type *string `json:"eipV4Type,omitempty"`

}

// DescribeEipRemoteRegionsResponseParams 查询EIP支持的远程指向的节点信息的响应信息。
type DescribeEipRemoteRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // PeerRegionIds 支持的远端的节点ID列表。
    PeerRegionIds []string `json:"peerRegionIds,omitempty"`

}

type DescribeEipRemoteRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipRemoteRegionsResponseParams `json:"response,omitempty"`

}

type DescribeEipsRequest struct {
    *common.BaseRequest

    // EipIds 按照 EIP 的唯一 ID 过滤。
    EipIds []string `json:"eipIds,omitempty"`

    // RegionId 按照 EIP 所属节点ID过滤。
    RegionId *string `json:"regionId,omitempty"`

    // Name 按照 EIP 的显示名称过滤，该字段支持模糊匹配。
    Name *string `json:"name,omitempty"`

    // Status 按照 EIP 的状态过滤。
    Status *string `json:"status,omitempty"`

    // IsDefault 按照 EIP 的默认属性进行过滤。
    IsDefault *bool `json:"isDefault,omitempty"`

    // PageSize 返回的分页大小。默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页页码。默认为1。
    PageNum *int `json:"pageNum,omitempty"`

    // PrivateIpAddress 按照 EIP 绑定的内网 IP 过滤。
    PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

    // IpAddress 按照 EIP 的 IP 过滤。
    IpAddress *string `json:"ipAddress,omitempty"`

    // IpAddresses 按照 EIP 的 IP列表过滤。
    IpAddresses []string `json:"ipAddresses,omitempty"`

    // InstanceId 按照 EIP 绑定的实例 ID 过滤。该字段过滤出该实例所绑定的网卡上的 EIP。
    InstanceId *string `json:"instanceId,omitempty"`

    // AssociatedId 按照 EIP 绑定的资源 ID 过滤。
    AssociatedId *string `json:"associatedId,omitempty"`

    // CidrIds 按照 EIP 所属的CIDR ID列表 过滤。
    CidrIds []string `json:"cidrIds,omitempty"`

    // ResourceGroupId 按照 EIP 所属的资源组ID过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type DescribeEipsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 满足过滤条件的EIP总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 返回的EIP列表。
    DataSet []*EipInfo `json:"dataSet,omitempty"`

}

// EipInfo 描述公网弹性IP的基本信息，包括IPv4地址，IP绑定的关系。
type EipInfo struct {

    // EipId EIP 的唯一 ID。
    EipId *string `json:"eipId,omitempty"`

    // Name EIP 的名称。
    Name *string `json:"name,omitempty"`

    // RegionId 节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // PeerRegionId 对端节点ID。仅当IP为Remote IP时该字段有效。
    PeerRegionId *string `json:"peerRegionId,omitempty"`

    // IsDefault 是否是默认类型。
    IsDefault *bool `json:"isDefault,omitempty"`

    // Status EIP 的状态。
    Status *string `json:"status,omitempty"`

    // PublicIpAddresses 公网IP地址。
    PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`

    // PrivateIpAddress EIP 绑定的内网IP地址。
    PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

    // EipV4Type EIP 的类型。
    EipV4Type *string `json:"eipV4Type,omitempty"`

    // InternetChargeType EIP 的网络计费方式。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // CidrId EIP 来自的CIDR地址段ID。
    CidrId *string `json:"cidrId,omitempty"`

    // NicId EIP 关联的网卡ID。 该字段已废弃，请使用 `associatedId` 字段。
    NicId *string `json:"nicId,omitempty"`

    // AssociatedId EIP 绑定的资源ID。可能为实例ID、网卡ID或者NAT网关ID。
    AssociatedId *string `json:"associatedId,omitempty"`

    // AssociatedType EIP 资源类型。可能为实例ID、网卡ID或者NAT网关ID。
    AssociatedType *string `json:"associatedType,omitempty"`

    // BindType EIP 绑定类型。
    BindType *string `json:"bindType,omitempty"`

    // FlowPackage EIP 流量包大小。 仅当网络计费方式为流量计费时可取到值。该字段可能为null。
    FlowPackage *float64 `json:"flowPackage,omitempty"`

    // Bandwidth EIP 的带宽限速。单位为Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // EipGeoRefs EIP 的地理位置信息。
    EipGeoRefs []*EipGeoInfo `json:"eipGeoRefs,omitempty"`

    // BlockInfoList EIP的封堵阈值。
    BlockInfoList []*BlockInfo `json:"blockInfoList,omitempty"`

    // CreateTime EIP 的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime EIP 的到期时间。该字段可能为null。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // ResourceGroupId EIP 关联的资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName EIP 关联的资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // BandwidthCluster EIP 关联的带宽组ID。
    BandwidthCluster *BandwidthClusterInfo `json:"bandwidthCluster,omitempty"`

}

// EipGeoInfo 基于DB-IP/ipdata/... 第三方IP数据库服务商查询到的 IP 地理信息结果。
type EipGeoInfo struct {

    // DbIp 从DB-IP(db-ip.com)查询的地理信息。
    DbIp *string `json:"dbIp,omitempty"`

    // IpData 从ipData(ipdata.co)查询的地理信息。
    IpData *string `json:"ipData,omitempty"`

    // IpInfo 从IPinfo(ipinfo.io)查询的地理信息。
    IpInfo *string `json:"ipInfo,omitempty"`

    // MaxMind 从MaxMind查询的地理信息。
    MaxMind *string `json:"maxMind,omitempty"`

    // Standard 需要查询EIP的所在的地理信息。
    Standard *string `json:"standard,omitempty"`

    // IsConsistent 查询地理信息是否和所在的地理信息一致。
    IsConsistent *bool `json:"isConsistent,omitempty"`

}

type BlockInfo struct {

    // Ip ip。
    Ip *string `json:"ip,omitempty"`

    // Bps 单位bps。
    Bps *int64 `json:"bps,omitempty"`

    // Pps 单位pps。
    Pps *int64 `json:"pps,omitempty"`

    // InCps 单位个。
    InCps *int64 `json:"inCps,omitempty"`

    // OutCps 单位个。
    OutCps *int64 `json:"outCps,omitempty"`

}

type DescribeEipsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipsResponseParams `json:"response,omitempty"`

}

// CreateEipsRequest 创建EIP的请求参数。
type CreateEipsRequest struct {
    *common.BaseRequest

    // RegionId 创建EIP所在的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // Name EIP的名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

    // InternetChargeType 公网弹性IP的网络计费方式。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // Amount 需要创建EIP的数量。
    Amount *int `json:"amount,omitempty"`

    // EipV4Type 公网弹性IP的线路类型。
    EipV4Type *string `json:"eipV4Type,omitempty"`

    // PrimaryIsp 主公网IP的运营商。该字段仅作用于三线IP(`ThreeLine`)。
    PrimaryIsp *string `json:"primaryIsp,omitempty"`

    // Bandwidth 公网弹性IP的带宽限速。单位：Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // CidrId 指定CIDR ID，使用CIDR内分配弹性IP。该字段和`eipV4Type`不能同时指定。
    CidrId *string `json:"cidrId,omitempty"`

    // PublicIp 从CIDR里指定公网起始IP地址开始创建弹性IP。该字段仅在指定`cidrId`时生效。
    PublicIp *string `json:"publicIp,omitempty"`

    // ResourceGroupId 弹性公网IP所放的资源组ID，如不指定则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // FlowPackage 公网IPv6的流量包大小。单位为TB。值要求为0或0.1的倍数。当子网的堆栈类型包括V6且为公网时，且网络计费方式是流量计费(`ByTrafficPackage`)需要指定。
    FlowPackage *float64 `json:"flowPackage,omitempty"`

    // ClusterId 公网IPv6所指定的共享带宽包ID。当子网的堆栈类型包括V6且为公网时，且网络计费方式是共享带宽包计费(`BandwidthCluster`)需要指定。
    ClusterId *string `json:"clusterId,omitempty"`

    // PeerRegionId 远端的节点ID。
    PeerRegionId *string `json:"peerRegionId,omitempty"`

    // MarketingOptions 市场营销的相关选项。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

}

// CreateEipsResponseParams 创建弹性公网IP的响应结果
type CreateEipsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // EipIds 创建的弹性公网IP ID列表。
    EipIds []string `json:"eipIds,omitempty"`

    // OrderNumber 本次创建的订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type CreateEipsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateEipsResponseParams `json:"response,omitempty"`

}

// DeleteEipRequest 删除弹性公网IP的请求参数。
type DeleteEipRequest struct {
    *common.BaseRequest

    // EipId 要删除弹性公网IP的ID。
    EipId *string `json:"eipId,omitempty"`

}

type DeleteEipResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// RenewEipRequest 恢复弹性公网IP的请求参数。
type RenewEipRequest struct {
    *common.BaseRequest

    // EipId 要恢复弹性公网IP的ID。
    EipId *string `json:"eipId,omitempty"`

}

type RenewEipResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ConfigEipEgressIpRequest 指定IP作为出口IP的请求参数。
type ConfigEipEgressIpRequest struct {
    *common.BaseRequest

    // EipId 要作为出口的公网弹性IP。
    EipId *string `json:"eipId,omitempty"`

}

type ConfigEipEgressIpResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeEipPriceRequest 查询公网弹性IP的价格。
type DescribeEipPriceRequest struct {
    *common.BaseRequest

    // RegionId 创建EIP所在的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // InternetChargeType 公网弹性IP的网络计费方式。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // Amount 需要创建EIP的数量。
    Amount *int `json:"amount,omitempty"`

    // EipV4Type 公网弹性IP的线路类型。
    EipV4Type *string `json:"eipV4Type,omitempty"`

    // Bandwidth 公网弹性IP的带宽限速。单位：Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // FlowPackage 公网IPv6的流量包大小。单位为TB。值要求为0或0.1的倍数。当子网的堆栈类型包括V6且为公网时，且网络计费方式是流量计费(`ByTrafficPackage`)需要指定。
    FlowPackage *float64 `json:"flowPackage,omitempty"`

    // CidrId 指定CIDR ID，使用CIDR内分配弹性IP。该字段和`eipV4Type`不能同时指定。
    CidrId *string `json:"cidrId,omitempty"`

    // ClusterId 公网IPv6所指定的共享带宽包ID。当子网的堆栈类型包括V6且为公网时，且网络计费方式是共享带宽包计费(`BandwidthCluster`)需要指定。
    ClusterId *string `json:"clusterId,omitempty"`

    // PeerRegionId 远端的节点ID。
    PeerRegionId *string `json:"peerRegionId,omitempty"`

}

// DescribeEipPriceResponseParams 查询公网弹性IP的价格的响应结果。
type DescribeEipPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // EipPrice 公网弹性IP的保留价格。如果是通过cidr创建，则保留价格为null。
    EipPrice *PriceItem `json:"eipPrice,omitempty"`

    // BandwidthPrice 公网弹性IP的带宽价格。
    BandwidthPrice *PriceItem `json:"bandwidthPrice,omitempty"`

}

type DescribeEipPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipPriceResponseParams `json:"response,omitempty"`

}

// ChangeEipInternetChargeTypeRequest 变更弹性公网IP更网络计费模式的请求参数。
type ChangeEipInternetChargeTypeRequest struct {
    *common.BaseRequest

    // EipId 要操作的公网弹性IP。
    EipId *string `json:"eipId,omitempty"`

    // InternetChargeType 要变更的目标网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // ClusterId 共享带宽包ID。如果要变更为共享带宽包计费，则需要指定。
    ClusterId *string `json:"clusterId,omitempty"`

}

// ChangeEipInternetChargeTypeResponseParams 变更弹性公网IP更网络计费模式
type ChangeEipInternetChargeTypeResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 变更可能产生的订单号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type ChangeEipInternetChargeTypeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ChangeEipInternetChargeTypeResponseParams `json:"response,omitempty"`

}

// AvailableLanIpRequest 查询可供弹性公网IP绑定的网卡及内网IP信息的请求参数。
type AvailableLanIpRequest struct {
    *common.BaseRequest

    // EipId 要查询的公网弹性IP ID。
    EipId *string `json:"eipId,omitempty"`

}

// AvailableLanIpResponseParams 查询可供弹性公网IP绑定的网卡及内网IP信息。
type AvailableLanIpResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // LanIps 可以绑定的网卡及内网信息。
    LanIps []*PrivateIpInfo `json:"lanIps,omitempty"`

}

// PrivateIpInfo 公网弹性IP可以绑定的网卡及内网信息
type PrivateIpInfo struct {

    // LanIp 网卡上的内网IP地址。
    LanIp *string `json:"lanIp,omitempty"`

    // NicId 网卡ID。
    NicId *string `json:"nicId,omitempty"`

    // NicName 网卡的名称。
    NicName *string `json:"nicName,omitempty"`

    // InstanceId 网卡关联的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 网卡关联的实例名称。
    InstanceName *string `json:"instanceName,omitempty"`

}

type AvailableLanIpResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *AvailableLanIpResponseParams `json:"response,omitempty"`

}

// DescribeEipTrafficRequest 查询Eip指定时间段内的流量信息的请求信息。
type DescribeEipTrafficRequest struct {
    *common.BaseRequest

    // EipId 要查询的弹性IP ID。
    EipId *string `json:"eipId,omitempty"`

    // StartTime 查询开始时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    EndTime *string `json:"endTime,omitempty"`

    // Step 查询数据点间隔。单位为分钟。支持参数：1,5。
    Step *int `json:"step,omitempty"`

    // WanIp 指定IP查询，该字段用于三线IP,可以指定IP地址查询流量。
    WanIp *string `json:"wanIp,omitempty"`

}

// DescribeEipTrafficResponseParams 查询弹性公网IP指定时间段内的流量信息的响应值。
type DescribeEipTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 查询的数据点。
    DataList []*EipTrafficData `json:"dataList,omitempty"`

    // In95 入方向带宽95值。单位：bps。
    In95 *int64 `json:"in95,omitempty"`

    // InAvg 入方向带宽平均值。单位：bps。
    InAvg *int64 `json:"inAvg,omitempty"`

    // InMax 入方向带宽的最大值。单位：bps。
    InMax *int64 `json:"inMax,omitempty"`

    // InMin 入方向带宽的最小值。单位：bps。
    InMin *int64 `json:"inMin,omitempty"`

    // InTotal 入方向的总字节。单位：Byte。
    InTotal *int64 `json:"inTotal,omitempty"`

    // Out95 出方向带宽95值。单位：bps。
    Out95 *int64 `json:"out95,omitempty"`

    // OutAvg 出方向带宽平均值。单位：bps。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出方向带宽最大值。单位：bps。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出方向带宽最小值。单位：bps。
    OutMin *int64 `json:"outMin,omitempty"`

    // OutTotal 入方向的总字节。单位：bps。
    OutTotal *int64 `json:"outTotal,omitempty"`

}

// EipTrafficData 描述IP带宽的数据点信息。包括出/入方向的带宽。
type EipTrafficData struct {

    // InternetRX 入方向带宽值。单位：bps。
    InternetRX *int64 `json:"internetRX,omitempty"`

    // InternetTX 出方向带宽值。单位：bps。
    InternetTX *int64 `json:"internetTX,omitempty"`

    // Time 数据时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    Time *string `json:"time,omitempty"`

}

type DescribeEipTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipTrafficResponseParams `json:"response,omitempty"`

}

// AssociateEipAddressRequest EIP绑定的请求信息。
type AssociateEipAddressRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例的ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // NicId 虚拟网卡的ID。
    NicId *string `json:"nicId,omitempty"`

    // LanIp 要绑定的网卡上的内网IP地址。当IP绑定的是网卡, 则该字段不能为空。
    LanIp *string `json:"lanIp,omitempty"`

    // NatId NAT网关的ID。
    NatId *string `json:"natId,omitempty"`

    // EipIds 要绑定的EIP的ID。EIP 必须是未绑定状态。
    EipIds []string `json:"eipIds,omitempty"`

    // BindType 绑定类型。当绑定的是网卡时生效。默认为普通NAT模式。
    BindType *string `json:"bindType,omitempty"`

}

// AssociateEipAddressResponseParams EIP绑定云产品的响应信息。
type AssociateEipAddressResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // FailedEipIds 绑定失败的IP。
    FailedEipIds []string `json:"failedEipIds,omitempty"`

}

type AssociateEipAddressResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *AssociateEipAddressResponseParams `json:"response,omitempty"`

}

// UnassociateEipAddressRequest 将弹性公网IP（EIP）从绑定的云产品上解绑的请求参数。
type UnassociateEipAddressRequest struct {
    *common.BaseRequest

    // EipIds 要解绑的弹性公网IP的ID列表。
    EipIds []string `json:"eipIds,omitempty"`

}

// UnassociateEipAddressResponseParams 将弹性公网IP（EIP）从绑定的云产品上解绑的响应结果。
type UnassociateEipAddressResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // FailedEipIds 操作失败的弹性IP的ID集合。
    FailedEipIds []string `json:"failedEipIds,omitempty"`

}

type UnassociateEipAddressResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *UnassociateEipAddressResponseParams `json:"response,omitempty"`

}

// ReplaceEipAddressRequest 替换弹性公网IP的请求信息。
type ReplaceEipAddressRequest struct {
    *common.BaseRequest

    // ReplaceIps 需要替换的弹性公网IP信息。
    ReplaceIps []*ReplaceIp `json:"replaceIps,omitempty"`

}

// ReplaceIp 替换的公网IP信息。
type ReplaceIp struct {

    // EipId 需要替换的弹性公网IP ID。
    EipId *string `json:"eipId,omitempty"`

    // OwnIp 原IP。当IP是三线IP(IP线路类型为`ThreeLine`)时需要指定。
    OwnIp *string `json:"ownIp,omitempty"`

    // TargetIp 需要变更的目标IP。如果未指定，将由系统随机分配。
    TargetIp *string `json:"targetIp,omitempty"`

}

// ReplaceEipAddressResponseParams 替换弹性公网IP的响应结果。
type ReplaceEipAddressResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // FailedEipIds 替换失败的IP的ID集合。
    FailedEipIds []string `json:"failedEipIds,omitempty"`

}

type ReplaceEipAddressResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ReplaceEipAddressResponseParams `json:"response,omitempty"`

}

// ModifyEipAttributeRequest 修改EIP属性的请求。
type ModifyEipAttributeRequest struct {
    *common.BaseRequest

    // EipId 公网弹性IP的ID。
    EipId *string `json:"eipId,omitempty"`

    // Name 公网弹性IP的名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

}

type ModifyEipAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyEipBandwidthRequest 修改带宽限速请求。
type ModifyEipBandwidthRequest struct {
    *common.BaseRequest

    // EipId EIP唯一标识ID。
    EipId *string `json:"eipId,omitempty"`

    // Bandwidth 调整带宽限速的目标值。单位Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // CommitBandwidth 保底带宽。单位Mbps。有且仅当为Remote IP，且为选择带宽包计费, 需要指定专线部分的保底带宽。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

}

type ModifyEipBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ChangeEipBindTypeRequest 更换弹性公网IP绑定模式请求信息。
type ChangeEipBindTypeRequest struct {
    *common.BaseRequest

    // EipId 要更换绑定模式的EIP ID。
    EipId *string `json:"eipId,omitempty"`

    // BindType 绑定类型。当绑定的是网卡时生效。默认为普通NAT模式。
    BindType *string `json:"bindType,omitempty"`

}

type ChangeEipBindTypeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeEipMonitorDataRequest 查询弹性公网IP监控指标请求。
type DescribeEipMonitorDataRequest struct {
    *common.BaseRequest

    // EipId EIP唯一标识ID。
    EipId *string `json:"eipId,omitempty"`

    // MetricType EIP监控指标类型。
    MetricType *string `json:"metricType,omitempty"`

    // StartTime 查询开始时间。时间格式：yyyy-MM-ddTHH:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。时间格式：yyyy-MM-ddTHH:mm:ssZ。
    EndTime *string `json:"endTime,omitempty"`

}

// DescribeEipMonitorDataResponseParams 查询弹性IP监控数据的响应信息。
type DescribeEipMonitorDataResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InMaxValue 入方向最大值。
    InMaxValue *float64 `json:"inMaxValue,omitempty"`

    // InAvgValue 入方向平均值。
    InAvgValue *float64 `json:"inAvgValue,omitempty"`

    // InMinValue 入方向最小值。
    InMinValue *float64 `json:"inMinValue,omitempty"`

    // InTotalValue 入方向总和值。
    InTotalValue *float64 `json:"inTotalValue,omitempty"`

    // OutMaxValue 出方向最大值。
    OutMaxValue *float64 `json:"outMaxValue,omitempty"`

    // OutAvgValue 出方向平均值。
    OutAvgValue *float64 `json:"outAvgValue,omitempty"`

    // OutMinValue 出方向最小值。
    OutMinValue *float64 `json:"outMinValue,omitempty"`

    // OutTotalValue 出方向总和值。
    OutTotalValue *float64 `json:"outTotalValue,omitempty"`

    // LoseOutMaxValue 丢失出方向最大值。
    LoseOutMaxValue *float64 `json:"loseOutMaxValue,omitempty"`

    // LoseOutMinValue 丢失出方向最小值。
    LoseOutMinValue *float64 `json:"loseOutMinValue,omitempty"`

    // LoseOutTotalValue 丢失出方向总和值。
    LoseOutTotalValue *float64 `json:"loseOutTotalValue,omitempty"`

    // LoseInMaxValue 丢失入方向最大值。
    LoseInMaxValue *float64 `json:"loseInMaxValue,omitempty"`

    // LoseInMinValue 丢失入方向最小值。
    LoseInMinValue *float64 `json:"loseInMinValue,omitempty"`

    // LoseInTotalValue 丢失入方向总和值。
    LoseInTotalValue *float64 `json:"loseInTotalValue,omitempty"`

    // DataList 监控数据集合。
    DataList []*EipMetricValue `json:"dataList,omitempty"`

}

// EipMetricValue 描述EIP的监控指标数据。
type EipMetricValue struct {

    // Time 数据点时间。
    Time *string `json:"time,omitempty"`

    // InValue 入方向值。
    InValue *float64 `json:"inValue,omitempty"`

    // OutValue 入方向值。
    OutValue *float64 `json:"outValue,omitempty"`

    // LoseIn 丢失入方向值。
    LoseIn *float64 `json:"loseIn,omitempty"`

    // LoseOut 丢失出方向值。
    LoseOut *float64 `json:"loseOut,omitempty"`

}

type DescribeEipMonitorDataResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeEipMonitorDataResponseParams `json:"response,omitempty"`

}

// DescribeZonesRequest 查询可用区的请求信息。
type DescribeZonesRequest struct {
    *common.BaseRequest

    // ZoneIds 根据可用区ID过滤。
    ZoneIds []string `json:"zoneIds,omitempty"`

}

// DescribeZonesResponseParams 查询可用区
type DescribeZonesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ZoneSet 可用区列表。
    ZoneSet []*ZoneInfo `json:"zoneSet,omitempty"`

}

// ZoneInfo 可用区的基本信息。
type ZoneInfo struct {

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // RegionId 可用区所在的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // ZoneName 可用区名称。
    ZoneName *string `json:"zoneName,omitempty"`

    // SupportSecurityGroup 可用区是否支持安全组。该字段已废弃，当前所有节点均支持安全组。
    SupportSecurityGroup *bool `json:"supportSecurityGroup,omitempty"`

}

type DescribeZonesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeZonesResponseParams `json:"response,omitempty"`

}

// DescribePoolsRequest 描述公网IP池的请求参数。
type DescribePoolsRequest struct {
    *common.BaseRequest

    // PoolIds 根据公网IP池的ID。
    PoolIds []string `json:"poolIds,omitempty"`

    // RegionId 根据公网IP池的所在节点ID过滤。
    RegionId *string `json:"regionId,omitempty"`

    // Name 根据公网IP池的名称过滤。支持模糊查询。
    Name *string `json:"name,omitempty"`

    // PageSize 返回的分页大小，默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数，默认为1。
    PageNum *int `json:"pageNum,omitempty"`

}

// DescribePoolsResponseParams 描述公网IP池的请求参数。
type DescribePoolsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 公网IP池列表的结果数据。
    DataSet []*PoolInfo `json:"dataSet,omitempty"`

}

// PoolInfo 描述公网IP池的基本信息。
type PoolInfo struct {

    // PoolId 公网IP池的ID。
    PoolId *string `json:"poolId,omitempty"`

    // RegionId 公网IP池所在的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // Name 公网IP池的名称。
    Name *string `json:"name,omitempty"`

    // CreateTime 公网IP池的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

type DescribePoolsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePoolsResponseParams `json:"response,omitempty"`

}

type DescribeCidrRegionsRequest struct {
    *common.BaseRequest

}

// DescribeCidrRegionsResponseParams 支持售卖CIDR的响应信息。
type DescribeCidrRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // RegionIds 支持售卖CIDR的节点ID。
    RegionIds []string `json:"regionIds,omitempty"`

}

type DescribeCidrRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCidrRegionsResponseParams `json:"response,omitempty"`

}

// DescribeCidrPriceRequest 查询CIDR价格的请求信息。
type DescribeCidrPriceRequest struct {
    *common.BaseRequest

    // RegionId 查询CIDR所在的节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // EipV4Type 查询CIDR的IP线路类型。
    EipV4Type *string `json:"eipV4Type,omitempty"`

    // Netmask 需要查询的CIDR子网数量和掩码信息。
    Netmask *NetmaskInfo `json:"netmask,omitempty"`

}

// NetmaskInfo 描述CIDR掩码的信息。
type NetmaskInfo struct {

    // Netmask 掩码大小。
    Netmask *int `json:"netmask,omitempty"`

    // Amount CIDR的数量。
    Amount *int `json:"amount,omitempty"`

}

// DescribeCidrPriceResponseParams 查询CIDR价格的响应信息。
type DescribeCidrPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CidrPrice CIDR的价格信息。
    CidrPrice *PriceItem `json:"cidrPrice,omitempty"`

}

type DescribeCidrPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCidrPriceResponseParams `json:"response,omitempty"`

}

// DescribeCidrsRequest 查询CIDR列表的请求信息。
type DescribeCidrsRequest struct {
    *common.BaseRequest

    // CidrIds 一个或多个待操作的Cidr ID，根据cidrId进行过滤。
    CidrIds []string `json:"cidrIds,omitempty"`

    // RegionId 根据CIDR所在的节点ID进行过滤。
    RegionId *string `json:"regionId,omitempty"`

    // Name 根据CIDR名称进行过滤，支持模糊查询。
    Name *string `json:"name,omitempty"`

    // CidrBlock 根据CIDR地址进行过滤。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // Source 根据CIDR来源进行过滤。
    Source *string `json:"source,omitempty"`

    // ResourceGroupId 根据资源组ID进行过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageSize 返回的分页大小，默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数，默认为1。
    PageNum *int `json:"pageNum,omitempty"`

}

// DescribeCidrsResponseParams 查询CIDR地址的响应信息。
type DescribeCidrsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 查询CIDR地址的结果数据。
    DataSet []*CidrInfo `json:"dataSet,omitempty"`

}

// CidrInfo CIDR信息详情。
type CidrInfo struct {

    // CidrId CIDR ID。
    CidrId *string `json:"cidrId,omitempty"`

    // RegionId CIDR所属的区域节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // Name CIDR的名称。
    Name *string `json:"name,omitempty"`

    // CidrBlock CIDR地址块，例如：10.0.0.0/16。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // TotalCount CIDR中IP地址的总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // UsedCount CIDR中已被使用的IP地址数量。
    UsedCount *int `json:"usedCount,omitempty"`

    // Source CIDR的来源。如CONSOLE（属于zenlayer）或 BYOIP（客户自带IP）。
    Source *string `json:"source,omitempty"`

    // EipV4Type EIP网络类型。表示该CIDR支持的公网IP线路类型。
    EipV4Type *string `json:"eipV4Type,omitempty"`

    // Netmask 子网掩码。表示CIDR的网络位长度。
    Netmask *int `json:"netmask,omitempty"`

    // PoolId Pool的ID。表示该CIDR所属的公网IP池。
    PoolId *string `json:"poolId,omitempty"`

    // CreateTime CIDR的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime CIDR的到期时间。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // ResourceGroupId 该CIDR所属的资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 该CIDR所属资源组的名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // Status CIDR的状态。
    Status *string `json:"status,omitempty"`

}

type DescribeCidrsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCidrsResponseParams `json:"response,omitempty"`

}

// CreateCidrRequest 创建CIDR请求参数
type CreateCidrRequest struct {
    *common.BaseRequest

    // RegionId 节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // EipV4Type 公网IPv4的网络类型。
    EipV4Type *string `json:"eipV4Type,omitempty"`

    // Netmask CIDR掩码、数量。
    Netmask *NetmaskInfo `json:"netmask,omitempty"`

    // Name CIDR名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。默认会将分配的CIDR地址作为名称。
    Name *string `json:"name,omitempty"`

    // ResourceGroupId 资源组ID。如果不指定，则会加入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // MarketingOptions 市场营销相关的选项。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

}

// CreateCidrResponseParams 描述创建CIDR的响应信息。
type CreateCidrResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CidrIds 创建的CIDR地址段ID列表。
    CidrIds []string `json:"cidrIds,omitempty"`

}

type CreateCidrResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateCidrResponseParams `json:"response,omitempty"`

}

// ModifyCidrAttributeRequest 修改CIDR地址段的属性的请求信息
type ModifyCidrAttributeRequest struct {
    *common.BaseRequest

    // CidrId 要修改CIDR地址段的ID。
    CidrId *string `json:"cidrId,omitempty"`

    // Name 要修改的名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    Name *string `json:"name,omitempty"`

}

type ModifyCidrAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteCidrRequest 描述删除CIDR请求信息。
type DeleteCidrRequest struct {
    *common.BaseRequest

    // CidrId 要删除的CIDR ID。
    CidrId *string `json:"cidrId,omitempty"`

}

type DeleteCidrResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// RenewCidrRequest 描述恢复CIDR请求信息。
type RenewCidrRequest struct {
    *common.BaseRequest

    // CidrId 要恢复的CIDR ID。
    CidrId *string `json:"cidrId,omitempty"`

}

type RenewCidrResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// CreateBorderGatewayRequest 创建边界网关的请求参数。
type CreateBorderGatewayRequest struct {
    *common.BaseRequest

    // RegionId 节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // VpcId VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

    // Label 名称。范围2到63个字符。仅支持输入字母、数字、-/_和英文句点(.)。且必须以数字或字母开头和结尾。
    Label *string `json:"label,omitempty"`

    // Asn ASN号。
    Asn *int `json:"asn,omitempty"`

    // RoutingMode 路由级别。
    RoutingMode *string `json:"routingMode,omitempty"`

    // AdvertisedSubnet 子网宣告控制。
    AdvertisedSubnet *string `json:"advertisedSubnet,omitempty"`

    // AdvertisedCidrs IPv4 Cidr集合。
    AdvertisedCidrs []string `json:"advertisedCidrs,omitempty"`

    // AdvertisedRouteIds 自定义路由ID集合。
    AdvertisedRouteIds []string `json:"advertisedRouteIds,omitempty"`

}

// CreateBorderGatewayResponseParams 创建边界网关的响应值。
type CreateBorderGatewayResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ZbgId 边界网关的ID。
    ZbgId *string `json:"zbgId,omitempty"`

}

type CreateBorderGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateBorderGatewayResponseParams `json:"response,omitempty"`

}

type DescribeBorderGatewaysRequest struct {
    *common.BaseRequest

    // ZbgIds 根据边界网关ID列表过滤。
    ZbgIds []string `json:"zbgIds,omitempty"`

    // Name 根据网关名称过滤。支持模糊搜索。
    Name *string `json:"name,omitempty"`

    // RegionId 根据边界网关所在的节点过滤。
    RegionId *string `json:"regionId,omitempty"`

    // VpcId 根据边界网关所属的VPC ID过滤。
    VpcId *string `json:"vpcId,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

}

// DescribeBorderGatewaysResponseParams 查询边界网关列表的响应结果。
type DescribeBorderGatewaysResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 边界网关的列表数据。
    DataSet []*ZbgInfo `json:"dataSet,omitempty"`

}

// ZbgInfo 描述边界网关的基本信息。
type ZbgInfo struct {

    // ZbgId 边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

    // Name 边界网关名称。
    Name *string `json:"name,omitempty"`

    // VpcId VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

    // RegionId 节点的ID。
    RegionId *string `json:"regionId,omitempty"`

    // Asn ASN号。
    Asn *int `json:"asn,omitempty"`

    // InterConnectCidr 互联IP地址段。
    InterConnectCidr *string `json:"interConnectCidr,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // CloudRouterIds 关联的三层网络ID集合。
    CloudRouterIds []string `json:"cloudRouterIds,omitempty"`

    // RoutingMode 路由模式。
    RoutingMode *string `json:"routingMode,omitempty"`

    // NatId NAT的ID。
    NatId *string `json:"natId,omitempty"`

    // AdvertisedSubnet 子网控制。
    AdvertisedSubnet *string `json:"advertisedSubnet,omitempty"`

    // AdvertisedCidrs IPv4 Cidr集合。
    AdvertisedCidrs []string `json:"advertisedCidrs,omitempty"`

    // AdvertisedRouteIds 自定义路由集合。
    AdvertisedRouteIds []string `json:"advertisedRouteIds,omitempty"`

}

type DescribeBorderGatewaysResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBorderGatewaysResponseParams `json:"response,omitempty"`

}

// DeleteBorderGatewayRequest 删除边界网关的请求信息。
type DeleteBorderGatewayRequest struct {
    *common.BaseRequest

    // ZbgId 要删除的边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

}

type DeleteBorderGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyBorderGatewayAsnRequest 修改边界网关ASN的请求信息。
type ModifyBorderGatewayAsnRequest struct {
    *common.BaseRequest

    // ZbgId 要修改的边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

    // Asn ASN。
    Asn *int `json:"asn,omitempty"`

}

type ModifyBorderGatewayAsnResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyBorderGatewaysAttributeRequest struct {
    *common.BaseRequest

    // ZbgIds 边界网关的ID列表。
    ZbgIds []string `json:"zbgIds,omitempty"`

    // Name 边界网关的名称。
    Name *string `json:"name,omitempty"`

    // RoutingMode 路由级别。
    RoutingMode *string `json:"routingMode,omitempty"`

    // AdvertisedSubnet 子网控制。
    AdvertisedSubnet *string `json:"advertisedSubnet,omitempty"`

    // AdvertisedCidrs IPv4 Cidr集合。
    AdvertisedCidrs []string `json:"advertisedCidrs,omitempty"`

    // Asn 边界网关的ASN。
    Asn *int `json:"asn,omitempty"`

}

type ModifyBorderGatewaysAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeAvailableNatsRequest 请求可用NAT网关的请求信息。
type DescribeAvailableNatsRequest struct {
    *common.BaseRequest

    // ZbgId 要查询需要绑定的边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

}

// DescribeAvailableNatsResponseParams 查询可以绑定边界网关的NAT响应信息。
type DescribeAvailableNatsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // NatIds 可以绑定边界网关的NAT ID集合。
    NatIds []string `json:"natIds,omitempty"`

}

type DescribeAvailableNatsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAvailableNatsResponseParams `json:"response,omitempty"`

}

// AssignBorderGatewayRequest 边界网关绑定NAT网关的请求信息。
type AssignBorderGatewayRequest struct {
    *common.BaseRequest

    // ZbgId 边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

    // NatId NAT网关ID。
    NatId *string `json:"natId,omitempty"`

}

type AssignBorderGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// UnassignBorderGatewayRequest 解绑边界网关的请求信息。
type UnassignBorderGatewayRequest struct {
    *common.BaseRequest

    // ZbgId 要解绑NAT的边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

}

type UnassignBorderGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type AssignBorderGatewayRouteRequest struct {
    *common.BaseRequest

    // ZbgId 边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

    // AdvertisedRouteIds 自定义路由ID集合。
    AdvertisedRouteIds []string `json:"advertisedRouteIds,omitempty"`

}

type AssignBorderGatewayRouteResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// UnassignBorderGatewayRouteRequest 边界网关路由宣告中移除指定的自定义路由的请求信息。
type UnassignBorderGatewayRouteRequest struct {
    *common.BaseRequest

    // ZbgId 边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

    // AdvertisedRouteIds 要移除的自定义路由ID集合。
    AdvertisedRouteIds []string `json:"advertisedRouteIds,omitempty"`

}

type UnassignBorderGatewayRouteResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeZoneInstanceConfigInfosRequest 查询可用区售卖的机型信息请求信息。
type DescribeZoneInstanceConfigInfosRequest struct {
    *common.BaseRequest

    // ZoneId 要查询的可用区ID。例如：asia-east-1a。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceType 要查询的实例规格。例如：z2a.cpu.1。
    InstanceType *string `json:"instanceType,omitempty"`

}

// DescribeZoneInstanceConfigInfosResponseParams 查询可用区售卖的机型信息的响应结果
type DescribeZoneInstanceConfigInfosResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceTypeQuotaSet 实例规格信息。
    InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"instanceTypeQuotaSet,omitempty"`

}

// InstanceTypeQuotaItem 描述实例规格的售卖信息。
type InstanceTypeQuotaItem struct {

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceType 实例规格ID。
    InstanceType *string `json:"instanceType,omitempty"`

    // InstanceTypeName 实例规格的名称。
    InstanceTypeName *string `json:"instanceTypeName,omitempty"`

    // CpuCount CPU核数。单位：核。
    CpuCount *int `json:"cpuCount,omitempty"`

    // Memory 实例内存容量。单位：GiB。
    Memory *int `json:"memory,omitempty"`

    // InternetMaxBandwidthOutLimit 最大出口带宽限制。
    InternetMaxBandwidthOutLimit *int `json:"internetMaxBandwidthOutLimit,omitempty"`

    // WithStock 是否有库存。
    WithStock *bool `json:"withStock,omitempty"`

    // InternetChargeTypes 支持的网络计费类型。
    InternetChargeTypes []string `json:"internetChargeTypes,omitempty"`

}

type DescribeZoneInstanceConfigInfosResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeZoneInstanceConfigInfosResponseParams `json:"response,omitempty"`

}

type DescribeTimeZonesRequest struct {
    *common.BaseRequest

}

// DescribeTimeZonesResponseParams 查询时区的响应结果。
type DescribeTimeZonesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TimeZones 所有的时区。
    TimeZones []string `json:"timeZones,omitempty"`

}

type DescribeTimeZonesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeTimeZonesResponseParams `json:"response,omitempty"`

}

type InquiryPriceCreateInstanceRequest struct {
    *common.BaseRequest

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceType 实例机型。具体取值可通过调用接口[DescribeZoneInstanceConfigInfos](describezoneinstanceconfiginfos.md)来获得最新的规格表。
    InstanceType *string `json:"instanceType,omitempty"`

    // EipV4Type 公网IPv4的线路类型。目前不支持三线IP(`ThreeLine`)。
    EipV4Type *string `json:"eipV4Type,omitempty"`

    // InternetChargeType 公网IP的网络计费类型。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // TrafficPackageSize 流量包订购大小。单位为TB。该值必须在`internetChargeType = ByTrafficPackage`时才会生效。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // Bandwidth 公网出带宽上限。单位：Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // InstanceCount 实例数量。
    InstanceCount *int `json:"instanceCount,omitempty"`

    // SystemDisk 系统盘相关信息。
    SystemDisk *SystemDisk `json:"systemDisk,omitempty"`

    // DataDisk 数据盘相关信息。
    DataDisk *DataDisk `json:"dataDisk,omitempty"`

}

// SystemDisk 描述系统盘的基本信息。
type SystemDisk struct {

    // DiskId 系统盘ID。该参数目前仅用于`DescribeInstances`等查询类接口的返回参数，不可用于`CreateInstances`等写接口的入参。
    DiskId *string `json:"diskId,omitempty"`

    // DiskSize 系统盘大小。单位：GiB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskCategory 云硬盘种类。Basic NVMe SSD: 经济型 NVMe SSD。Standard NVMe SSD: 标准型 NVMe SSD。默认为Standard NVMe SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

}

// DataDisk 描述了数据盘的信息。
type DataDisk struct {

    // DiskId 数据盘ID。该参数目前仅用于`DescribeInstances`等查询类接口的返回参数，不可用于`CreateInstances`等写接口的入参。
    DiskId *string `json:"diskId,omitempty"`

    // DiskName 云盘的名称。该参数目前仅用于DescribeInstances等查询类接口的返回参数。
    DiskName *string `json:"diskName,omitempty"`

    // DiskSize 数据盘大小。单位：GiB。
    DiskSize *int `json:"diskSize,omitempty"`

    // DiskAmount 数据盘数量。该参数仅用于`CreateInstances`,`InquiryPriceCreateInstance`等接口的入参使用。
    DiskAmount *int `json:"diskAmount,omitempty"`

    // Portable 是否可拆卸。该参数仅用于查询的回参，不用于入参。true代表不会随着实例删除而删除。false代表会随着实例删除而删除。
    Portable *bool `json:"portable,omitempty"`

    // DiskCategory 云硬盘种类。Basic NVMe SSD: 经济型 NVMe SSD。Standard NVMe SSD: 标准型 NVMe SSD。默认为Standard NVMe SSD。
    DiskCategory *string `json:"diskCategory,omitempty"`

}

// InquiryPriceCreateInstanceResponseParams 
type InquiryPriceCreateInstanceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SpecPrice 规格的价格。
    SpecPrice *PriceItem `json:"specPrice,omitempty"`

    // GpuPrice GPU规格的价格。
    GpuPrice *PriceItem `json:"gpuPrice,omitempty"`

    // Ipv4Price 公网IPv4的保留价格。
    Ipv4Price *PriceItem `json:"ipv4Price,omitempty"`

    // Ipv4BandwidthPrice 公网IPv4的带宽价格。
    Ipv4BandwidthPrice *PriceItem `json:"ipv4BandwidthPrice,omitempty"`

    // Ipv6Price IPv6的价格。
    Ipv6Price *PriceItem `json:"ipv6Price,omitempty"`

    // Ipv6BandwidthPrice IPv6的带宽价格。
    Ipv6BandwidthPrice *PriceItem `json:"ipv6BandwidthPrice,omitempty"`

    // SystemDiskPrice 系统盘的价格。
    SystemDiskPrice *PriceItem `json:"systemDiskPrice,omitempty"`

    // DataDiskPrice 数据盘的价格。
    DataDiskPrice *PriceItem `json:"dataDiskPrice,omitempty"`

}

type InquiryPriceCreateInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceCreateInstanceResponseParams `json:"response,omitempty"`

}

// CreateZecInstancesRequest 创建虚拟机实例的请求参数。
type CreateZecInstancesRequest struct {
    *common.BaseRequest

    // ZoneId 可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ImageId 指定有效的镜像ID。可以通过[DescribeImages](describeimages.md)取返回信息中的imageId字段。
    ImageId *string `json:"imageId,omitempty"`

    // TimeZone 设置操作系统的时区。
    TimeZone *string `json:"timeZone,omitempty"`

    // InstanceType 实例机型。具体取值可通过调用接口[DescribeZoneInstanceConfigInfos](describezoneinstanceconfiginfos.md)来获得最新的规格表。
    InstanceType *string `json:"instanceType,omitempty"`

    // InstanceName 实例显示名称。范围2到63个字符。仅支持输入字母、数字、-和英文句点(.)。且必须以数字或字母开头和结尾。购买多台实例，可以指定模式串[begin_number,bits]。begin_number：有序数值的起始值，取值支持[0,99999]，默认值为0。bits：有序数值所占的位数，取值支持[1,6]，默认值为6。注意模式串中不得有空格。购买1台时，例如server-[3,3]实例显示为server003；购买2台时，实例显示名分别为server003，server004。支持指定多个模式串，如server-[3,3]-[1,1]。默认值为 instance。
    InstanceName *string `json:"instanceName,omitempty"`

    // Password 实例的密码。与keyId必须指定其中的一种。必须包含以下3种格式的字符：大小写字母: [a-zA-Z]数字: 0-9特殊字符: ~!@$^*-_=+。
    Password *string `json:"password,omitempty"`

    // KeyId 密钥ID。与password必须指定其中的一种。可调用接口DescribeKeyPairs来获得最新的密钥对信息。关联密钥后，就可以通过对应的私钥来访问实例；密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。示例值：key-YWD2QFOl。
    KeyId *string `json:"keyId,omitempty"`

    // InstanceCount 要创建的实例数量。
    InstanceCount *int `json:"instanceCount,omitempty"`

    // SystemDisk 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。即操作系统要求的最小大小。
    SystemDisk *SystemDisk `json:"systemDisk,omitempty"`

    // DataDisks 实例数据盘配置信息。若不指定该参数，则默认不额外购买数据盘。目前只能附带1个数据盘。
    DataDisks []*DataDisk `json:"dataDisks,omitempty"`

    // SecurityGroupId 要配置在实例主网卡的安全组ID。目前只能关联1个安全组。如果未指定，会默认用VPC关联的安全组。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // SubnetId 子网ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // LanIp 分配的内网起始IP。如果内网IP被使用,则会往后分配。
    LanIp *string `json:"lanIp,omitempty"`

    // EnableAgent 是否安装启动Agent。
    EnableAgent *bool `json:"enableAgent,omitempty"`

    // EnableIpForward 是否开启IP转发。
    EnableIpForward *bool `json:"enableIpForward,omitempty"`

    // InternetChargeType 公网IP的网络计费类型。如果不指定，则不会分配公网IP地址。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // TrafficPackageSize 流量包订购大小。单位为TB。该值必须在`internetChargeType = ByTrafficPackage`时才会生效。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // Bandwidth 公网出带宽上限。单位：Mbps。当分配公网IP时需要指定。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // EipBindType 公网IP的绑定模式。当分配公网IP时需要指定。
    EipBindType *string `json:"eipBindType,omitempty"`

    // EipV4Type 公网IPv4的线路类型。当分配公网IP时需要指定。请确保所选子网的堆栈类型支持`IPv4`。目前不支持三线IP随实例一起创建。
    EipV4Type *string `json:"eipV4Type,omitempty"`

    // ClusterId 共享带宽包ID。当网络计费方式是共享带宽包计费(`BandwidthCluster`)时需要指定。
    ClusterId *string `json:"clusterId,omitempty"`

    // ResourceGroupId 创建后实例所在的资源组ID，如不指定则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // MarketingOptions 市场营销的相关选项。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

}

// CreateZecInstancesResponseParams 创建虚拟机实例的响应结果。
type CreateZecInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // InstanceIdSet 虚拟机实例ID列表。
    InstanceIdSet []string `json:"instanceIdSet,omitempty"`

    // Instances 随机器创建的数据盘id集合。如果请求中没有指定数据盘，返回空数组。
    Instances []*DiskWithInstance `json:"instances,omitempty"`

}

// DiskWithInstance 随机器创建的数据盘信息。
type DiskWithInstance struct {

    // InstanceId 实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // DiskIdSet 随机器创建的数据盘ID集合。
    DiskIdSet []string `json:"diskIdSet,omitempty"`

}

type CreateZecInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateZecInstancesResponseParams `json:"response,omitempty"`

}

// DescribeInstancesRequest 描述实例列表的请求参数。
type DescribeInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 根据实例ID列表进行筛选。最大不能超过100个。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // ZoneId 实例所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // Ipv4Address 根据实例关联的IPv4过滤。
    Ipv4Address *string `json:"ipv4Address,omitempty"`

    // Ipv6Address 根据实例关联的IPv6信息过滤。
    Ipv6Address *string `json:"ipv6Address,omitempty"`

    // Status 根据实例的状态过滤。
    Status *string `json:"status,omitempty"`

    // Name 根据实例显示名称过滤。该字段支持模糊搜索。
    Name *string `json:"name,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // ResourceGroupId 根据资源组ID过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

// DescribeInstancesResponseParams 描述实例列表的请求参数。
type DescribeInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 实例列表的数据。
    DataSet []*InstanceInfo `json:"dataSet,omitempty"`

}

// InstanceInfo 描述虚拟机实例的信息。包括规格，状态，网卡等。
type InstanceInfo struct {

    // InstanceId 实例唯一ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceName 实例显示名称。
    InstanceName *string `json:"instanceName,omitempty"`

    // ZoneId 实例所属的可用区ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // InstanceType CPU 规格。如果是GPU实例，该字段取值为null。
    InstanceType *string `json:"instanceType,omitempty"`

    // Cpu CPU 核数。单位：个。
    Cpu *int `json:"cpu,omitempty"`

    // Memory 内存容量。单位：GiB。
    Memory *int `json:"memory,omitempty"`

    // ImageId 镜像ID。
    ImageId *string `json:"imageId,omitempty"`

    // ImageName 镜像名称。
    ImageName *string `json:"imageName,omitempty"`

    // TimeZone 设置的系统时区信息。
    TimeZone *string `json:"timeZone,omitempty"`

    // NicNetworkType 网卡模式。
    NicNetworkType *string `json:"nicNetworkType,omitempty"`

    // Status 实例状态。
    Status *string `json:"status,omitempty"`

    // SystemDisk 系统盘信息。
    SystemDisk *SystemDisk `json:"systemDisk,omitempty"`

    // DataDisks 实例上挂在的数据盘信息。
    DataDisks []*DataDisk `json:"dataDisks,omitempty"`

    // PublicIpAddresses 实例上公网IPv4列表。
    PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`

    // PrivateIpAddresses 实例上内网IP列表。
    PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`

    // KeyId 安装的SSH密钥ID。
    KeyId *string `json:"keyId,omitempty"`

    // SubnetId 实例主网卡关联的子网ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // SecurityGroupId 实例主网卡关联的安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // EnableAgent 是否开启QGA Agent。
    EnableAgent *bool `json:"enableAgent,omitempty"`

    // EnableAgentMonitor 是否开启QGA 监控采集。
    EnableAgentMonitor *bool `json:"enableAgentMonitor,omitempty"`

    // EnableIpForward 是否开启IP转发。
    EnableIpForward *bool `json:"enableIpForward,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 到期时间。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // ResourceGroupId 实例所属的资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 实例所属的资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // Nics 实例上绑定的网卡信息。
    Nics []*NicInfo `json:"nics,omitempty"`

}

type DescribeInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstancesResponseParams `json:"response,omitempty"`

}

// DescribeInstancesStatusRequest 查询实例状态的请求参数。
type DescribeInstancesStatusRequest struct {
    *common.BaseRequest

    // InstanceIds 要查询的实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // PageSize 分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 分页页数。
    PageNum *int `json:"pageNum,omitempty"`

    // ResourceGroupId 根据资源组ID过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

// DescribeInstancesStatusResponseParams 查询实例状态
type DescribeInstancesStatusResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 实例状态数据。
    DataSet []*InstanceStatus `json:"dataSet,omitempty"`

}

// InstanceStatus 描述实例状态的信息。
type InstanceStatus struct {

    // InstanceId 实例的ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceStatus 实例的状态。
    InstanceStatus *string `json:"instanceStatus,omitempty"`

}

type DescribeInstancesStatusResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstancesStatusResponseParams `json:"response,omitempty"`

}

// ModifyInstancesAttributeRequest 修改实例属性（名称）的请求参数。
type ModifyInstancesAttributeRequest struct {
    *common.BaseRequest

    // InstanceIds 待修改属性的实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // InstanceName 实例名称。2～63个字符。仅支持输入字母、数字、-和英文句点(.)。
    InstanceName *string `json:"instanceName,omitempty"`

}

type ModifyInstancesAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// StartInstancesRequest 启动虚拟机实例的参数。
type StartInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 要启动的实例ID。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

// StartInstancesResponseParams 启动虚拟机实例的响应结果。
type StartInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceIds 启动失败的实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type StartInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *StartInstancesResponseParams `json:"response,omitempty"`

}

// StopInstancesRequest 关闭虚拟机实例的请求信息。
type StopInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 待关闭的实例ID。
    InstanceIds []string `json:"instanceIds,omitempty"`

    // ForceShutdown 是否强制关机。
    ForceShutdown *bool `json:"forceShutdown,omitempty"`

}

// StopInstancesResponseParams 关闭虚拟机实例
type StopInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceIds 操作失败的实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type StopInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *StopInstancesResponseParams `json:"response,omitempty"`

}

// RebootInstancesRequest 重启虚拟机实例的请求参数。
type RebootInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 待重启虚拟机实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

// RebootInstancesResponseParams 重启虚拟机实例的响应结果
type RebootInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceIds 重启操作失败的实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type RebootInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *RebootInstancesResponseParams `json:"response,omitempty"`

}

// ResetInstancePasswordRequest 重置虚拟机实例密码的请求参数。
type ResetInstancePasswordRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // Password 密码。必须是8-16位。必须包含以下3种格式的字符：大小写字母: [a-zA-Z]数字: 0-9特殊字符: ~!@$^*-_=+。
    Password *string `json:"password,omitempty"`

}

type ResetInstancePasswordResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ResetInstanceRequest 重装一台虚拟机实例操作系统的请求参数。
type ResetInstanceRequest struct {
    *common.BaseRequest

    // InstanceId 要重装的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // Password 实例的新密码。与keyId不能同时指定。必须包含以下3种格式的字符：大小写字母: [a-zA-Z]数字: 0-9特殊字符: ~!@$^*-_=+。
    Password *string `json:"password,omitempty"`

    // KeyId 密钥ID。与password必须指定其中的一种。可调用接口DescribeKeyPairs来获得最新的密钥对信息。关联密钥后，就可以通过对应的私钥来访问实例；密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。示例值：key-YWD2QFOl。
    KeyId *string `json:"keyId,omitempty"`

    // ImageId 指定重装的的镜像ID。可以通过[DescribeImages](describeimages.md)取返回信息中的`imageId`字段。如果不指定，会根据当前镜像进行重装。
    ImageId *string `json:"imageId,omitempty"`

    // Timezone 操作系统时区设置。
    Timezone *string `json:"timezone,omitempty"`

    // EnableAgent 是否启用 QEMU Guest 代理 (QGA)。
    EnableAgent *bool `json:"enableAgent,omitempty"`

    // InstanceName 修改的实例名称。
    InstanceName *string `json:"instanceName,omitempty"`

}

type ResetInstanceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// StartIpForwardRequest 开启IP转发的请求信息。
type StartIpForwardRequest struct {
    *common.BaseRequest

    // InstanceId 要操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type StartIpForwardResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// StopIpForwardRequest 关闭IP转发的请求信息。
type StopIpForwardRequest struct {
    *common.BaseRequest

    // InstanceId 要操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type StopIpForwardResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// StartAgentMonitorRequest 开启Agent监控采集的请求信息。
type StartAgentMonitorRequest struct {
    *common.BaseRequest

    // InstanceId 要操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type StartAgentMonitorResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// StopAgentMonitorRequest 关闭Agent监控采集的请求信息。
type StopAgentMonitorRequest struct {
    *common.BaseRequest

    // InstanceId 要操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

type StopAgentMonitorResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyInstanceTypeRequest 变更实例规格的请求参数。
type ModifyInstanceTypeRequest struct {
    *common.BaseRequest

    // InstanceId 要变更的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // InstanceType 变更的目标实例规格。
    InstanceType *string `json:"instanceType,omitempty"`

}

// ModifyInstanceTypeResponseParams 变更实例规格的响应结果。
type ModifyInstanceTypeResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 变更产生订单的编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

}

type ModifyInstanceTypeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ModifyInstanceTypeResponseParams `json:"response,omitempty"`

}

// ChangeNicNetworkTypeRequest 更改网卡模式的请求参数。
type ChangeNicNetworkTypeRequest struct {
    *common.BaseRequest

    // InstanceId 待操作的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // NicNetworkType 要更改的网卡模式。
    NicNetworkType *string `json:"nicNetworkType,omitempty"`

}

type ChangeNicNetworkTypeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ReleaseInstancesRequest 释放实例请求参数。
type ReleaseInstancesRequest struct {
    *common.BaseRequest

    // InstanceIds 要释放的实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

// ReleaseInstancesResponseParams 删除实例的响应结果。
type ReleaseInstancesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // InstanceIds 操作失败的实例ID列表。
    InstanceIds []string `json:"instanceIds,omitempty"`

}

type ReleaseInstancesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ReleaseInstancesResponseParams `json:"response,omitempty"`

}

// DescribeVncUrlRequest 获取实例VNC地址的请求信息。
type DescribeVncUrlRequest struct {
    *common.BaseRequest

    // InstanceId 要查询的实例ID。
    InstanceId *string `json:"instanceId,omitempty"`

}

// DescribeVncUrlResponseParams 获取实例VNC地址的响应结果。
type DescribeVncUrlResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Url VNC地址URL。
    Url *string `json:"url,omitempty"`

}

type DescribeVncUrlResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeVncUrlResponseParams `json:"response,omitempty"`

}

// DescribeInstanceMonitorDataRequest 查询实例监控指标请求。
type DescribeInstanceMonitorDataRequest struct {
    *common.BaseRequest

    // InstanceId 实例唯一标识ID。
    InstanceId *string `json:"instanceId,omitempty"`

    // MetricType 实例监控指标类型。
    MetricType *string `json:"metricType,omitempty"`

    // StartTime 查询开始时间。时间格式：yyyy-MM-ddTHH:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。时间格式：yyyy-MM-ddTHH:mm:ssZ。
    EndTime *string `json:"endTime,omitempty"`

}

// DescribeInstanceMonitorDataResponseParams 查询实例监控数据的响应信息。
type DescribeInstanceMonitorDataResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // MaxValue 数据点的最大值。
    MaxValue *float64 `json:"maxValue,omitempty"`

    // MinValue 数据点的最小值。
    MinValue *float64 `json:"minValue,omitempty"`

    // AvgValue 数据点的平均值。
    AvgValue *float64 `json:"avgValue,omitempty"`

    // Metrics 监控数据集合。
    Metrics []*MetricValue `json:"metrics,omitempty"`

}

// MetricValue 描述实例监控指标的数据值。
type MetricValue struct {

    // Time 数据点时间。
    Time *string `json:"time,omitempty"`

    // Value 数据点的值。如果该值为null,表示取不到相应的值。
    Value *float64 `json:"value,omitempty"`

}

type DescribeInstanceMonitorDataResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeInstanceMonitorDataResponseParams `json:"response,omitempty"`

}

type CreateNatGatewayRequest struct {
    *common.BaseRequest

    // RegionId 区域节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // VpcId NAT网关所属的VPC网络ID。
    VpcId *string `json:"vpcId,omitempty"`

    // Name NAT网关的名称。长度为2～63个字符。
    Name *string `json:"name,omitempty"`

    // SubnetIds NAT网关所属的Subnet子网ID集合。如果未指定，则指定区域的所有子网将自动关联NAT网关。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // SecurityGroupId 安全组ID。如果未指定，则指定VPC所属的安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // ResourceGroupId 资源组ID。如果不指定，则会创建在默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type CreateNatGatewayResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 下单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // NatGatewayId NAT网关唯一ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

}

type CreateNatGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateNatGatewayResponseParams `json:"response,omitempty"`

}

// ModifyNatGatewayAttributeRequest 修改Nat网关属性的请求
type ModifyNatGatewayAttributeRequest struct {
    *common.BaseRequest

    // NatGatewayId NAT网关 ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

    // Name NAT网关的名称。长度为2～63个字符。
    Name *string `json:"name,omitempty"`

    // SubnetIds NAT网关的子网ID。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // IsAllSubnet NAT网关对应的子网是否应用所有子网。该字段不能和`subnetIds`同时设置。
    IsAllSubnet *bool `json:"isAllSubnet,omitempty"`

    // IcmpReplyEnabled 是否开启ICMP代回。
    IcmpReplyEnabled *bool `json:"icmpReplyEnabled,omitempty"`

    // SecurityGroupId 修改NAT网关绑定的目标安全组ID。目前一张NAT网关只能关联一个安全组。指定该字段会解绑NAT网关原来的安全组。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

}

type ModifyNatGatewayAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeNatGatewayRegionsRequest struct {
    *common.BaseRequest

}

type DescribeNatGatewayRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // RegionIds 节点ID集合。
    RegionIds []string `json:"regionIds,omitempty"`

}

type DescribeNatGatewayRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeNatGatewayRegionsResponseParams `json:"response,omitempty"`

}

type DescribeNatGatewaysRequest struct {
    *common.BaseRequest

    // RegionId 节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // VpcId 根据NAT网关所属的VPC网络 ID过滤。
    VpcId *string `json:"vpcId,omitempty"`

    // NatGatewayIds NAT网关ID集合。实例ID数量上限为100个。
    NatGatewayIds []string `json:"natGatewayIds,omitempty"`

    // Name NAT网关名称。
    Name *string `json:"name,omitempty"`

    // Status NAT网关状态。
    Status *string `json:"status,omitempty"`

    // SecurityGroupId 根据NAT网关所属的安全组ID过滤。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // PageSize 返回的分页大小。默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页页码。默认为1。
    PageNum *int `json:"pageNum,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type DescribeNatGatewaysResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 满足过滤条件的NAT网关总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 返回的NAT网关列表。
    DataSet []*NatGateway `json:"dataSet,omitempty"`

}

// NatGateway 描述NAT网关的信息。
type NatGateway struct {

    // NatGatewayId NAT网关唯一ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

    // VpcId NAT网关所属的VPC网络ID。
    VpcId *string `json:"vpcId,omitempty"`

    // RegionId 区域节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // Status NAT网关的状态。
    Status *string `json:"status,omitempty"`

    // Name NAT网关的名称。
    Name *string `json:"name,omitempty"`

    // SubnetIds NAT网关所属的Subnet子网ID集合。
    SubnetIds []string `json:"subnetIds,omitempty"`

    // IsAllSubnets 是否节点内所有子网关联了NAT网关。
    IsAllSubnets *bool `json:"isAllSubnets,omitempty"`

    // EipIds NAT网关所关联的EIP ID集合。
    EipIds []string `json:"eipIds,omitempty"`

    // ZbgId 边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

    // IcmpReplyEnabled 是否开启ICMP代回。
    IcmpReplyEnabled *bool `json:"icmpReplyEnabled,omitempty"`

    // SecurityGroupId 边界网关关联的安全组ID。
    SecurityGroupId *string `json:"securityGroupId,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // CreateTime 创建时间。按照ISO8601标准表示，并且使用UTC时间, 格式为：YYYY-MM-DDThh:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 到期时间。按照ISO8601标准表示，并且使用UTC时间, 格式为：YYYY-MM-DDThh:mm:ssZ。
    ExpiredTime *string `json:"expiredTime,omitempty"`

}

type DescribeNatGatewaysResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeNatGatewaysResponseParams `json:"response,omitempty"`

}

type DescribeNatGatewayDetailRequest struct {
    *common.BaseRequest

    // NatGatewayId NAT网关 ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

}

type DescribeNatGatewayDetailResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // NatGatewayId NAT网关唯一ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

    // Name NAT网关名称。
    Name *string `json:"name,omitempty"`

    // Snats SNAT网关规则集合。
    Snats []*SnatEntry `json:"snats,omitempty"`

    // Dnats DNAT网关规则集合。
    Dnats []*DnatEntry `json:"dnats,omitempty"`

}

// SnatEntry 描述SNAT规则的信息。
type SnatEntry struct {

    // SnatEntryId SNAT规则 ID。
    SnatEntryId *string `json:"snatEntryId,omitempty"`

    // Cidrs CIDR网段，不传默认是0.0.0.0/0。`cidrs` 和 `snatSubnets` 不会同时存在。
    Cidrs []string `json:"cidrs,omitempty"`

    // EipIds SNAT规则添加的eip ID集合。
    EipIds []string `json:"eipIds,omitempty"`

    // SnatSubnets SNAT规则添加的subnet ID集合。
    SnatSubnets []*SnatSubnet `json:"snatSubnets,omitempty"`

}

// SnatSubnet Snat规则添加的子网集合。
type SnatSubnet struct {

    // SubnetId 子网的ID。
    SubnetId *string `json:"subnetId,omitempty"`

    // Cidr 子网的CIDR。
    Cidr *string `json:"cidr,omitempty"`

}

// DnatEntry 描述DNAT规则的信息。
type DnatEntry struct {

    // DnatEntryId DNAT规则 ID。
    DnatEntryId *string `json:"dnatEntryId,omitempty"`

    // Status DNAT规则状态。
    Status *string `json:"status,omitempty"`

    // PrivateIp DNAT规则的内网IP地址。
    PrivateIp *string `json:"privateIp,omitempty"`

    // EipId DNAT规则添加的eip ID。
    EipId *string `json:"eipId,omitempty"`

    // Protocol DNAT规则的协议类型。
    Protocol *string `json:"protocol,omitempty"`

    // ListenerPort DNAT规则端口转发的外部端口或端口段，取值范围1-65535。
    ListenerPort *string `json:"listenerPort,omitempty"`

    // InternalPort DNAT规则端口转发的内部端口或端口段，取值范围1-65535。
    InternalPort *string `json:"internalPort,omitempty"`

}

type DescribeNatGatewayDetailResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeNatGatewayDetailResponseParams `json:"response,omitempty"`

}

type DeleteNatGatewayRequest struct {
    *common.BaseRequest

    // NatGatewayId NAT网关 ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

}

type DeleteNatGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type RenewNatGatewayRequest struct {
    *common.BaseRequest

    // NatGatewayId NAT网关 ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

}

type RenewNatGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type InquiryPriceCreateNatGatewayRequest struct {
    *common.BaseRequest

    // RegionId 区域节点ID。
    RegionId *string `json:"regionId,omitempty"`

}

// InquiryPriceCreateNatGatewayResponseParams 描述创建NAT网关的响应。
type InquiryPriceCreateNatGatewayResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // NatGatewayPrice NAT网关的价格。
    NatGatewayPrice *PriceItem `json:"natGatewayPrice,omitempty"`

    // CuPrice NAT网关CU的价格。
    CuPrice *PriceItem `json:"cuPrice,omitempty"`

}

type InquiryPriceCreateNatGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceCreateNatGatewayResponseParams `json:"response,omitempty"`

}

type CreateSnatEntryRequest struct {
    *common.BaseRequest

    // NatGatewayId NAT网关 ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

    // Cidr CIDR网段。与subnetIds必须指定其中的一种。
    Cidr *string `json:"cidr,omitempty"`

    // EipIds SNAT规则添加的eip ID集合。为空则代表与该NAT网关绑定的所有eip。
    EipIds []string `json:"eipIds,omitempty"`

    // SubnetIds Subnet ID集合。与cidr必须指定其中的一种。
    SubnetIds []string `json:"subnetIds,omitempty"`

}

type CreateSnatEntryResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SnatEntryId SNAT规则唯一ID。
    SnatEntryId *string `json:"snatEntryId,omitempty"`

}

type CreateSnatEntryResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateSnatEntryResponseParams `json:"response,omitempty"`

}

type ModifySnatEntryRequest struct {
    *common.BaseRequest

    // SnatEntryId SNAT规则 ID。
    SnatEntryId *string `json:"snatEntryId,omitempty"`

    // Cidr CIDR网段。与subnetIds必须指定其中的一种。
    Cidr *string `json:"cidr,omitempty"`

    // EipIds SNAT规则添加的eip ID集合。为空则代表与该NAT网关绑定的所有eip。
    EipIds []string `json:"eipIds,omitempty"`

    // SubnetIds Subnet ID集合。与cidr必须指定其中的一种。
    SubnetIds []string `json:"subnetIds,omitempty"`

}

type ModifySnatEntryResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteSnatEntryRequest struct {
    *common.BaseRequest

    // SnatEntryId SNAT规则 ID。
    SnatEntryId *string `json:"snatEntryId,omitempty"`

}

type DeleteSnatEntryResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type CreateDnatEntryRequest struct {
    *common.BaseRequest

    // NatGatewayId NAT网关 ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

    // EipId DNAT规则添加的eip ID。
    EipId *string `json:"eipId,omitempty"`

    // PrivateIp DNAT规则的内网IP地址。
    PrivateIp *string `json:"privateIp,omitempty"`

    // Protocol DNAT规则的协议类型。
    Protocol *string `json:"protocol,omitempty"`

    // ListenerPort DNAT规则端口转发的外部端口或端口段，取值范围1-65535。
    ListenerPort *string `json:"listenerPort,omitempty"`

    // InternalPort DNAT规则端口转发的内部端口或端口段，取值范围1-65535。
    InternalPort *string `json:"internalPort,omitempty"`

}

type CreateDnatEntryResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DnatEntryId DNAT规则唯一ID。
    DnatEntryId *string `json:"dnatEntryId,omitempty"`

}

type CreateDnatEntryResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateDnatEntryResponseParams `json:"response,omitempty"`

}

type ModifyDnatEntryRequest struct {
    *common.BaseRequest

    // DnatEntryId DNAT规则 ID。
    DnatEntryId *string `json:"dnatEntryId,omitempty"`

    // EipId 修改DNAT关联的弹性公网IP ID。
    EipId *string `json:"eipId,omitempty"`

    // PrivateIp DNAT规则的内网IP地址。
    PrivateIp *string `json:"privateIp,omitempty"`

    // Protocol DNAT规则的协议类型。
    Protocol *string `json:"protocol,omitempty"`

    // ListenerPort DNAT规则端口转发的外部端口或端口段，取值范围1-65535。
    ListenerPort *string `json:"listenerPort,omitempty"`

    // InternalPort DNAT规则端口转发的内部端口或端口段，取值范围1-65535。
    InternalPort *string `json:"internalPort,omitempty"`

}

type ModifyDnatEntryResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteDnatEntryRequest struct {
    *common.BaseRequest

    // DnatEntryId DNAT规则 ID。
    DnatEntryId *string `json:"dnatEntryId,omitempty"`

}

type DeleteDnatEntryResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeAvailableBorderGatewayRequest 获取可绑定NAT的边界网关的请求信息。
type DescribeAvailableBorderGatewayRequest struct {
    *common.BaseRequest

    // NatGatewayId NAT网关 ID。
    NatGatewayId *string `json:"natGatewayId,omitempty"`

}

// DescribeAvailableBorderGatewayResponseParams 可绑定NAT的边界网关的相应信息。
type DescribeAvailableBorderGatewayResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ZbgId 可绑定NAT的边界网关ID。
    ZbgId *string `json:"zbgId,omitempty"`

}

type DescribeAvailableBorderGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAvailableBorderGatewayResponseParams `json:"response,omitempty"`

}

type CreateRouteRequest struct {
    *common.BaseRequest

    // VpcId VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

    // IpVersion IP类型。支持`IPv4`和`IPv6`两种类型。
    IpVersion *string `json:"ipVersion,omitempty"`

    // RouteType 路由类型。
    RouteType *string `json:"routeType,omitempty"`

    // SourceCidrBlock 源IP地址CIDR。`路由类型`配置`RouteTypePolicy(策略路由)`时需指定。
    SourceCidrBlock *string `json:"sourceCidrBlock,omitempty"`

    // DestinationCidrBlock IPv4或IPv6的目标网段。例如：10.0.1.0/24。该字段必填。
    DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty"`

    // CidrBlock IPv4或IPv6的目标网段。例如：10.0.1.0/24。该字段已废弃，请使用`destinationCidrBlock`。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // Priority 路由优先级。数值越小，优先级越高。
    Priority *int `json:"priority,omitempty"`

    // NextHopId 下一跳资源ID。目前只支持网卡ID。该字段必填。
    NextHopId *string `json:"nextHopId,omitempty"`

    // NextHotId 下一跳资源ID。目前只支持网卡ID。该字段已废弃， 请使用`nextHopId`。
    NextHotId *string `json:"nextHotId,omitempty"`

    // Name 路由名称。名称长度为 2 到 63 个字符，仅支持字母、数字、连字符 (-) 、下划线(_) 、斜杠(/) 、和句点 (.)，且开头和结尾必须是字母或数字。
    Name *string `json:"name,omitempty"`

}

// CreateRouteResponseParams 创建路由的响应信息。
type CreateRouteResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // RouteId 创建的路由ID。
    RouteId *string `json:"routeId,omitempty"`

}

type CreateRouteResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateRouteResponseParams `json:"response,omitempty"`

}

// ModifyRouteAttributeRequest 修改路由信息的请求信息。
type ModifyRouteAttributeRequest struct {
    *common.BaseRequest

    // RouteId 路由ID。
    RouteId *string `json:"routeId,omitempty"`

    // Name 路由名称。名称长度为 2 到 63 个字符，仅支持字母、数字、连字符 (-) 、下划线(_) 、斜杠(/) 、和句点 (.)，且开头和结尾必须是字母或数字。
    Name *string `json:"name,omitempty"`

}

type ModifyRouteAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteRouteRequest 删除路由的请求信息。
type DeleteRouteRequest struct {
    *common.BaseRequest

    // RouteId 路由ID。
    RouteId *string `json:"routeId,omitempty"`

}

type DeleteRouteResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeRoutesRequest struct {
    *common.BaseRequest

    // RouteIds 根据路由ID过滤。最多同时传入100个ID。
    RouteIds []string `json:"routeIds,omitempty"`

    // VpcId 根据路由关联的VPC ID过滤。
    VpcId *string `json:"vpcId,omitempty"`

    // IpVersion 根据IP类型过滤。支持`IPv4`和`IPv6`两种类型。
    IpVersion *string `json:"ipVersion,omitempty"`

    // RouteType 根据路由类型过滤。
    RouteType *string `json:"routeType,omitempty"`

    // Name 根据名称类型过滤。该字段支持模糊搜索。
    Name *string `json:"name,omitempty"`

    // DestinationCidrBlock 根据Ipv4或IPv6的目标网段过滤。例如：10.0.1.0/24。
    DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty"`

    // PageSize 返回的分页大小。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    PageNum *int `json:"pageNum,omitempty"`

}

// DescribeRoutesResponseParams 描述路由列表的响应信息。
type DescribeRoutesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 路由列表数据。
    DataSet []*RouteInfo `json:"dataSet,omitempty"`

}

// RouteInfo 描述路由的相关信息。
type RouteInfo struct {

    // RouteId 路由ID。
    RouteId *string `json:"routeId,omitempty"`

    // Name 路由的名称。
    Name *string `json:"name,omitempty"`

    // VpcId 路由关联的VPC的ID。
    VpcId *string `json:"vpcId,omitempty"`

    // VpcName 路由关联的VPC的名称。
    VpcName *string `json:"vpcName,omitempty"`

    // IpVersion IP类型。支持`IPv4`和`IPv6`两种类型。
    IpVersion *string `json:"ipVersion,omitempty"`

    // Type 路由类型。
    Type *string `json:"type,omitempty"`

    // SourceCidrBlock 源IP地址。当`路由类型`是`RouteTypePolicy(策略路由)`时可取值。
    SourceCidrBlock *string `json:"sourceCidrBlock,omitempty"`

    // DestinationCidrBlock IPv4或IPv6的目标网段。例如：10.0.1.0/24。
    DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty"`

    // CidrBlock IPv4或IPv6的目标网段。例如：10.0.1.0/24。该字段已废弃，请使用`destinationCidrBlock`。
    CidrBlock *string `json:"cidrBlock,omitempty"`

    // Priority 路由优先级。
    Priority *int `json:"priority,omitempty"`

    // NextHopId 下一跳资源ID。
    NextHopId *string `json:"nextHopId,omitempty"`

    // NextHopName 下一跳资源名称。
    NextHopName *string `json:"nextHopName,omitempty"`

    // NextHopType 下一跳的类型。
    NextHopType *string `json:"nextHopType,omitempty"`

    // CreateTime 路由的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

type DescribeRoutesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeRoutesResponseParams `json:"response,omitempty"`

}

