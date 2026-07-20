package sdn

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// CreatePortRequest 
type CreatePortRequest struct {
    *common.BaseRequest

    // DcId 数据中心ID。
    // 可通过 ~~DescribeDataCenters~~ 接口获取。
    DcId *string `json:"dcId,omitempty"`

    // PortName 数据中心端口名称。
    // 不超过 255 字符，建议使用数据中心四字码+数据中心端口规格。
    PortName *string `json:"portName,omitempty"`

    // PortRemarks 数据中心端口备注信息。
    // 不超过 255 字符。
    PortRemarks *string `json:"portRemarks,omitempty"`

    // PortType 数据中心端口规格。
    // 可通过 ~~DescribeDataCenterPortPrice~~ 接口获取，取值：1G | 10G | 40G | 100G | 400G。
    PortType *string `json:"portType,omitempty"`

    // BusinessEntityName 商业实体名称。
    // 用于 LOA 抬头。
    BusinessEntityName *string `json:"businessEntityName,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

    // Tags 创建数据中心端口时关联的标签。
    // 注意：关联`标签键`不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

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

type CreatePortResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreatePortResponseParams `json:"response,omitempty"`

}

// CreatePortResponseParams 
type CreatePortResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 创建数据中心端口产生的订单号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // PortId 创建成功的数据中心端口 ID。
    PortId *string `json:"portId,omitempty"`

}

// DestroyPortRequest 
type DestroyPortRequest struct {
    *common.BaseRequest

    // PortId 数据中心端口 ID。
    PortId *string `json:"portId,omitempty"`

}

type DestroyPortResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// TerminatePortRequest 
type TerminatePortRequest struct {
    *common.BaseRequest

    // PortId 数据中心端口 ID。
    PortId *string `json:"portId,omitempty"`

}

type TerminatePortResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// RenewPortRequest 
type RenewPortRequest struct {
    *common.BaseRequest

    // PortId 数据中心端口 ID。
    // 仅数据中心端口状态为 RECYCLED 时可恢复。
    PortId *string `json:"portId,omitempty"`

}

type RenewPortResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyPortAttributeRequest 
type ModifyPortAttributeRequest struct {
    *common.BaseRequest

    // PortId 数据中心端口 ID。
    PortId *string `json:"portId,omitempty"`

    // PortName 数据中心端口名称。
    // 最长 255 字符，建议使用数据中心四字码+数据中心端口规格。
    PortName *string `json:"portName,omitempty"`

    // PortRemarks 数据中心端口备注信息。
    // 最长 255 字符。
    PortRemarks *string `json:"portRemarks,omitempty"`

    // BusinessEntityName 商业实体名称。
    // 用于 LOA 抬头。
    // 仅在数据中心端口 LOA 上传前可修改，已上传则不支持修改。
    BusinessEntityName *string `json:"businessEntityName,omitempty"`

}

type ModifyPortAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribePortsRequest 
type DescribePortsRequest struct {
    *common.BaseRequest

    // PageSize 每页数量。
    // 取值范围 1-1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 页码。
    // 从 1 开始。
    PageNum *int `json:"pageNum,omitempty"`

    // PortIds 数据中心端口 ID 列表。
    // 最大支持长度为 100。
    PortIds []string `json:"portIds,omitempty"`

    // PortName 数据中心端口名称。
    // 最长不超过 255 个字符，支持模糊匹配、忽略大小写。
    PortName *string `json:"portName,omitempty"`

    // DcId 数据中心ID。
    // 可通过 ~~DescribeDataCenters~~ 接口获取。
    DcId *string `json:"dcId,omitempty"`

    // CityName 城市名称。
    // 最长不超过 64 个字符，支持模糊匹配、忽略大小写。
    CityName *string `json:"cityName,omitempty"`

    // ResourceGroupId 资源组 ID。
    // 不传则返回该用户可见的所有资源组内的数据中心端口。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribePortsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePortsResponseParams `json:"response,omitempty"`

}

// DescribePortsResponseParams 
type DescribePortsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 数据中心端口列表结果集。
    DataSet []*PortInfo `json:"dataSet,omitempty"`

}

// PortInfo 描述数据中心端口的详细信息。
type PortInfo struct {

    // PortId 数据中心端口 ID。
    PortId *string `json:"portId,omitempty"`

    // PortName 数据中心端口名称。
    PortName *string `json:"portName,omitempty"`

    // PortRemarks 数据中心端口备注信息。
    PortRemarks *string `json:"portRemarks,omitempty"`

    // PortType 数据中心端口规格。
    PortType *string `json:"portType,omitempty"`

    // DcId 数据中心ID。
    DcId *string `json:"dcId,omitempty"`

    // DcName 数据中心名称。
    DcName *string `json:"dcName,omitempty"`

    // CityName 城市名称。
    CityName *string `json:"cityName,omitempty"`

    // AreaName 所在大区名称。
    AreaName *string `json:"areaName,omitempty"`

    // BusinessEntityName 商业实体名称。
    // 用于 LOA 抬头。
    BusinessEntityName *string `json:"businessEntityName,omitempty"`

    // ConnectionStatus 数据中心端口连接状态。
    ConnectionStatus *string `json:"connectionStatus,omitempty"`

    // PortStatus 数据中心端口业务状态。
    PortStatus *string `json:"portStatus,omitempty"`

    // LoaStatus LOA 状态。
    LoaStatus *string `json:"loaStatus,omitempty"`

    // LoaDownloadUrl LOA 下载地址。
    LoaDownloadUrl *string `json:"loaDownloadUrl,omitempty"`

    // CreatedTime 创建时间。
    CreatedTime *string `json:"createdTime,omitempty"`

    // ExpiredTime 到期时间。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // Period 购买时长。
    // 单位月。
    Period *int `json:"period,omitempty"`

    // IsCreateBusinessAllowed 是否允许在该数据中心端口上开通业务。
    IsCreateBusinessAllowed *bool `json:"isCreateBusinessAllowed,omitempty"`

    // Tags 数据中心端口关联的标签。
    Tags *Tags `json:"tags,omitempty"`

}

// Tags 描述资源关联的标签信息。
type Tags struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// DescribeDataCenterPortPriceRequest 
type DescribeDataCenterPortPriceRequest struct {
    *common.BaseRequest

    // DcId 数据中心ID。
    // 可通过 ~~DescribeDataCenters~~ 接口获取。
    DcId *string `json:"dcId,omitempty"`

}

type DescribeDataCenterPortPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeDataCenterPortPriceResponseParams `json:"response,omitempty"`

}

// DescribeDataCenterPortPriceResponseParams 
type DescribeDataCenterPortPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // PortPriceSet 在售数据中心端口类型及价格结果集。
    PortPriceSet []*PortPrice `json:"portPriceSet,omitempty"`

}

// PortPrice 描述数据中心端口的类型、类别及价格信息。
type PortPrice struct {

    // DcId 数据中心ID。
    DcId *string `json:"dcId,omitempty"`

    // PortType 数据中心端口规格。
    PortType *string `json:"portType,omitempty"`

    // PortCategory 数据中心端口类别。
    PortCategory *string `json:"portCategory,omitempty"`

    // PortPrice 数据中心端口价格信息。
    PortPrice *PriceItem `json:"portPrice,omitempty"`

    // Stock 库存数量。
    Stock *int `json:"stock,omitempty"`

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

// DescribePortTrafficRequest 
type DescribePortTrafficRequest struct {
    *common.BaseRequest

    // PortId 数据中心端口 ID。
    PortId *string `json:"portId,omitempty"`

    // StartTime 查询开始时间。
    // ISO8601 标准 UTC 格式：YYYY-MM-DDThh:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // ISO8601 标准 UTC 格式：YYYY-MM-DDThh:mm:ssZ。
    EndTime *string `json:"endTime,omitempty"`

}

type DescribePortTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePortTrafficResponseParams `json:"response,omitempty"`

}

// DescribePortTrafficResponseParams 
type DescribePortTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 流量明细数据列表。
    DataList []*TrafficData `json:"dataList,omitempty"`

    // In95 入方向 95 峰值流量。
    In95 *int64 `json:"in95,omitempty"`

    // InAvg 入方向平均流量。
    InAvg *int64 `json:"inAvg,omitempty"`

    // InMax 入方向最大流量。
    InMax *int64 `json:"inMax,omitempty"`

    // InMin 入方向最小流量。
    InMin *int64 `json:"inMin,omitempty"`

    // Out95 出方向 95 峰值流量。
    Out95 *int64 `json:"out95,omitempty"`

    // OutAvg 出方向平均流量。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出方向最大流量。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出方向最小流量。
    OutMin *int64 `json:"outMin,omitempty"`

    // Unit 流量单位。
    // 如 bps。
    Unit *string `json:"unit,omitempty"`

}

// TrafficData 描述带宽的数据点信息。
type TrafficData struct {

    // InternetRX 入方向带宽值。
    // 单位：bps。
    InternetRX *int64 `json:"internetRX,omitempty"`

    // InternetTX 出方向带宽值。
    // 单位：bps。
    InternetTX *int64 `json:"internetTX,omitempty"`

    // Time 数据时间。
    // 按照ISO8601标准表示，并且使用UTC时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    Time *string `json:"time,omitempty"`

}

// DescribePortUsableVlanRequest 
type DescribePortUsableVlanRequest struct {
    *common.BaseRequest

    // PortId 数据中心端口 ID。
    // 数据中心端口或数据中心至少传一个。
    PortId *string `json:"portId,omitempty"`

    // DcId 数据中心ID。
    // 可通过 ~~DescribeDataCenters~~ 接口获取。
    // 数据中心端口或数据中心至少传一个。
    DcId *string `json:"dcId,omitempty"`

}

type DescribePortUsableVlanResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePortUsableVlanResponseParams `json:"response,omitempty"`

}

// DescribePortUsableVlanResponseParams 
type DescribePortUsableVlanResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // InuseVlanList 占用中的 VLAN 列表。
    InuseVlanList []int `json:"inuseVlanList,omitempty"`

}

// QueryCloudOnrampPriceRequest 
type QueryCloudOnrampPriceRequest struct {
    *common.BaseRequest

    // CloudType 云连接类型。
    // 可选值：AWS | TENCENT | GOOGLE | ALI_CLOUD | AZURE | HUAWEI_CLOUD。
    CloudType *string `json:"cloudType,omitempty"`

    // DcId 云连接的数据中心ID。
    DcId *string `json:"dcId,omitempty"`

    // BandwidthMbps 云连接的最大带宽限制。
    // 默认值是10，单位：Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // VlanId VLAN ID。
    VlanId *int `json:"vlanId,omitempty"`

    // CloudRegionId 公有云区域ID。
    // Google云无需传参。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

}

type QueryCloudOnrampPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *QueryCloudOnrampPriceResponseParams `json:"response,omitempty"`

}

// QueryCloudOnrampPriceResponseParams 
type QueryCloudOnrampPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Price 价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Stock 可用库存数量。
    Stock *int `json:"stock,omitempty"`

}

// QueryDataCenterPortPriceRequest 
type QueryDataCenterPortPriceRequest struct {
    *common.BaseRequest

    // DcId 数据中心ID。
    // 具体取值可通过调用接口~~DescribeDataCenters~~来获得最新的数据中心列表。
    DcId *string `json:"dcId,omitempty"`

    // PortType 数据中心端口规格。
    // 具体取值可通过调用接口~~DescribeDataCenterPortType~~来获得最新的数据中心端口规格表。
    PortType *string `json:"portType,omitempty"`

    // BuildCrossConnectWithAssisted 是否需要Zenlayer协助建设交叉连接。
    // 如果选择true，则价格将包含交叉连接的跳线费用，以及一次性建设费。
    // 默认值为false，即用户需要自行建设交叉连接。
    BuildCrossConnectWithAssisted *bool `json:"buildCrossConnectWithAssisted,omitempty"`

}

type QueryDataCenterPortPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *QueryDataCenterPortPriceResponseParams `json:"response,omitempty"`

}

// QueryDataCenterPortPriceResponseParams 
type QueryDataCenterPortPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CrossConnectPrice 交叉连接价格。
    // 如果当前数据中心不支持Zenlayer协助建设交叉连接，则该值为null。
    CrossConnectPrice *PriceItem `json:"crossConnectPrice,omitempty"`

    // CrossConnectOneTimeConstructionPrice 交叉连接一次性建设费。
    // 如果当前数据中心不支持Zenlayer协助建设交叉连接，则该值为null。
    CrossConnectOneTimeConstructionPrice *PriceItem `json:"crossConnectOneTimeConstructionPrice,omitempty"`

    // DcId 数据中心ID。
    DcId *string `json:"dcId,omitempty"`

    // Price 价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Stock 可用库存数量。
    Stock *int `json:"stock,omitempty"`

}

// QueryDataCenterPortPricesRequest 
type QueryDataCenterPortPricesRequest struct {
    *common.BaseRequest

    // DcIds 数据中心ID列表。
    // 具体取值可通过调用接口~~DescribeDataCenters~~来获得最新的数据中心列表。
    // 最多支持100个ID查询。
    DcIds []string `json:"dcIds,omitempty"`

    // PortType 数据中心端口规格。
    // 具体取值可通过调用接口~~DescribeDataCenterPortType~~来获得最新的数据中心端口规格表。
    PortType *string `json:"portType,omitempty"`

    // BuildCrossConnectWithAssisted 是否需要Zenlayer协助建设交叉连接。
    // 如果选择true，则价格将包含交叉连接的跳线费用，以及一次性建设费。
    // 默认值为false，即用户需要自行建设交叉连接。
    BuildCrossConnectWithAssisted *bool `json:"buildCrossConnectWithAssisted,omitempty"`

}

type QueryDataCenterPortPricesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *QueryDataCenterPortPricesResponseParams `json:"response,omitempty"`

}

// QueryDataCenterPortPricesResponseParams 
type QueryDataCenterPortPricesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Prices 数据中心端口价格列表。
    Prices []*DatacenterPortPrice `json:"prices,omitempty"`

}

// DatacenterPortPrice 描述数据中心端口的价格信息。
type DatacenterPortPrice struct {

    // CrossConnectPrice 交叉连接价格。
    // 如果当前数据中心不支持Zenlayer协助建设交叉连接，则该值为null。
    CrossConnectPrice *PriceItem `json:"crossConnectPrice,omitempty"`

    // CrossConnectOneTimeConstructionPrice 交叉连接一次性建设费。
    // 如果当前数据中心不支持Zenlayer协助建设交叉连接，则该值为null。
    CrossConnectOneTimeConstructionPrice *PriceItem `json:"crossConnectOneTimeConstructionPrice,omitempty"`

    // DcId 数据中心ID。
    DcId *string `json:"dcId,omitempty"`

    // Price 价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Stock 可用库存数量。
    Stock *int `json:"stock,omitempty"`

}

// QueryPrivateConnectPriceRequest 
type QueryPrivateConnectPriceRequest struct {
    *common.BaseRequest

    // InternetType 网络计费类型。
    // 默认值为ByBandwidth。
    InternetType *string `json:"internetType,omitempty"`

    // BandwidthMbps 二层网络专线的最大带宽限制。
    // 默认值是10，单位：Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // EndpointA 二层网络专线A端连接点信息。
    EndpointA *PrivateConnectEndpointInfo `json:"endpointA,omitempty"`

    // EndpointZ 二层网络专线Z端连接点信息。
    EndpointZ *PrivateConnectEndpointInfo `json:"endpointZ,omitempty"`

}

// PrivateConnectEndpointInfo 二层网络专线的连接点信息。根据连接点的类型不同，连接点的参数也是不同的。
type PrivateConnectEndpointInfo struct {

    // DcId 数据中心ID。
    // 具体取值可通过调用接口~~DescribeDataCenters~~来获得最新的数据中心列表。
    DcId *string `json:"dcId,omitempty"`

    // CloudType 云连接类型。
    CloudType *string `json:"cloudType,omitempty"`

    // BandwidthMbps 云连接的最大带宽限制。
    // 默认值是10，单位：Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // VlanId VLAN ID。
    VlanId *int `json:"vlanId,omitempty"`

    // CloudRegionId 公有云区域ID。
    // Google云无需传参。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // PortType 数据中心端口规格。
    // 具体取值可通过调用接口~~DescribeDataCenterPortType~~来获得最新的数据中心端口规格表。
    PortType *string `json:"portType,omitempty"`

    // BuildCrossConnectWithAssisted 是否需要Zenlayer协助建设交叉连接。
    // 如果选择true，则价格将包含交叉连接的跳线费用，以及一次性建设费。
    // 默认值为false，即用户需要自行建设交叉连接。
    BuildCrossConnectWithAssisted *bool `json:"buildCrossConnectWithAssisted,omitempty"`

}

type QueryPrivateConnectPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *QueryPrivateConnectPriceResponseParams `json:"response,omitempty"`

}

// QueryPrivateConnectPriceResponseParams 
type QueryPrivateConnectPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // EndpointAPrice 二层网络专线A端连接点价格信息。
    EndpointAPrice *PrivateConnectEndpointPrice `json:"endpointAPrice,omitempty"`

    // EndpointZPrice 二层网络专线Z端连接点价格信息。
    EndpointZPrice *PrivateConnectEndpointPrice `json:"endpointZPrice,omitempty"`

    // Price 价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Stock 可用库存数量。
    Stock *int `json:"stock,omitempty"`

}

// PrivateConnectEndpointPrice 描述二层网络专线连接点的价格信息。
type PrivateConnectEndpointPrice struct {

    // CrossConnectPrice 交叉连接价格。
    // 连接点类型为数据中心端口并且该数据中心支持Zenlayer协助建设交叉连接时有值。
    CrossConnectPrice *PriceItem `json:"crossConnectPrice,omitempty"`

    // CrossConnectOneTimeConstructionPrice 交叉连接一次性建设费。
    // 连接点类型为数据中心端口并且该数据中心支持Zenlayer协助建设交叉连接时有值。
    CrossConnectOneTimeConstructionPrice *PriceItem `json:"crossConnectOneTimeConstructionPrice,omitempty"`

    // Price 价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Stock 可用库存数量。
    Stock *int `json:"stock,omitempty"`

}

// QueryPrivateConnectBandwidthPriceRequest 
type QueryPrivateConnectBandwidthPriceRequest struct {
    *common.BaseRequest

    // SourceDcId 二层网络其中一端接入点的数据中心ID。
    SourceDcId *string `json:"sourceDcId,omitempty"`

    // DestinationDcId 二层网络另外一端接入点的数据中心ID。
    DestinationDcId *string `json:"destinationDcId,omitempty"`

    // InternetType 网络计费类型。
    // 默认值为ByBandwidth。
    InternetType *string `json:"internetType,omitempty"`

    // BandwidthMbps 二层网络专线的最大带宽限制。
    // 默认值是10，单位：Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // ServiceLevel 二层网络专线服务等级。
    // 默认值为SINGLE_PROTECTED。
    ServiceLevel *string `json:"serviceLevel,omitempty"`

}

type QueryPrivateConnectBandwidthPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *QueryPrivateConnectBandwidthPriceResponseParams `json:"response,omitempty"`

}

// QueryPrivateConnectBandwidthPriceResponseParams 
type QueryPrivateConnectBandwidthPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Price 价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Stock 可用库存数量。
    Stock *int `json:"stock,omitempty"`

}

// QueryCloudRouterBandwidthPriceRequest 
type QueryCloudRouterBandwidthPriceRequest struct {
    *common.BaseRequest

    // DcId 三层网络骨干网的数据中心ID。
    DcId *string `json:"dcId,omitempty"`

    // BandwidthMbps 三层网络的最大带宽限制。
    // 默认值是10，单位：Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

}

type QueryCloudRouterBandwidthPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *QueryCloudRouterBandwidthPriceResponseParams `json:"response,omitempty"`

}

// QueryCloudRouterBandwidthPriceResponseParams 
type QueryCloudRouterBandwidthPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Price 价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Stock 可用库存数量。
    Stock *int `json:"stock,omitempty"`

}

// DescribeGoogleVlanUsageRequest 
type DescribeGoogleVlanUsageRequest struct {
    *common.BaseRequest

    // DcId 接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // PairingKey 云平台配对密钥。
    // 用于按密钥查询对应接入点的 VLAN 使用情况。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeGoogleVlanUsageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeGoogleVlanUsageResponseParams `json:"response,omitempty"`

}

// DescribeGoogleVlanUsageResponseParams 
type DescribeGoogleVlanUsageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start 可用的 VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End 可用的 VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // UsedVlans 已使用的 VLAN 列表。
    UsedVlans []int `json:"usedVlans,omitempty"`

}

// DescribeTencentVlanUsageRequest 
type DescribeTencentVlanUsageRequest struct {
    *common.BaseRequest

    // DcId 接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // PairingKey 云平台配对密钥。
    // 用于按密钥查询对应接入点的 VLAN 使用情况。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeTencentVlanUsageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeTencentVlanUsageResponseParams `json:"response,omitempty"`

}

// DescribeTencentVlanUsageResponseParams 
type DescribeTencentVlanUsageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start 可用的 VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End 可用的 VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // UsedVlans 已使用的 VLAN 列表。
    UsedVlans []int `json:"usedVlans,omitempty"`

}

// DescribeAliCloudVlanUsageRequest 
type DescribeAliCloudVlanUsageRequest struct {
    *common.BaseRequest

    // DcId 接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // PairingKey 云平台配对密钥。
    // 用于按密钥查询对应接入点的 VLAN 使用情况。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeAliCloudVlanUsageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAliCloudVlanUsageResponseParams `json:"response,omitempty"`

}

// DescribeAliCloudVlanUsageResponseParams 
type DescribeAliCloudVlanUsageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start 可用的 VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End 可用的 VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // UsedVlans 已使用的 VLAN 列表。
    UsedVlans []int `json:"usedVlans,omitempty"`

}

// DescribeHuaweiCloudVlanUsageRequest 
type DescribeHuaweiCloudVlanUsageRequest struct {
    *common.BaseRequest

    // DcId 接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // PairingKey 云平台配对密钥。
    // 用于按密钥查询对应接入点的 VLAN 使用情况。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeHuaweiCloudVlanUsageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeHuaweiCloudVlanUsageResponseParams `json:"response,omitempty"`

}

// DescribeHuaweiCloudVlanUsageResponseParams 
type DescribeHuaweiCloudVlanUsageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start 可用的 VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End 可用的 VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // UsedVlans 已使用的 VLAN 列表。
    UsedVlans []int `json:"usedVlans,omitempty"`

}

// DescribeAzureVlanUsageRequest 
type DescribeAzureVlanUsageRequest struct {
    *common.BaseRequest

    // DcId 接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // PairingKey 云平台配对密钥。
    // 用于按密钥查询对应接入点的 VLAN 使用情况。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeAzureVlanUsageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAzureVlanUsageResponseParams `json:"response,omitempty"`

}

// DescribeAzureVlanUsageResponseParams 
type DescribeAzureVlanUsageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start 可用的 VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End 可用的 VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // UsedVlans 已使用的 VLAN 列表。
    UsedVlans []int `json:"usedVlans,omitempty"`

}

// DescribeOracleVlanUsageRequest 
type DescribeOracleVlanUsageRequest struct {
    *common.BaseRequest

    // DcId 接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // PairingKey 云平台配对密钥。
    // 用于按密钥查询对应接入点的 VLAN 使用情况。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeOracleVlanUsageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeOracleVlanUsageResponseParams `json:"response,omitempty"`

}

// DescribeOracleVlanUsageResponseParams 
type DescribeOracleVlanUsageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start 可用的 VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End 可用的 VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // UsedVlans 已使用的 VLAN 列表。
    UsedVlans []int `json:"usedVlans,omitempty"`

}

// DescribeBytePlusVlanUsageRequest 
type DescribeBytePlusVlanUsageRequest struct {
    *common.BaseRequest

    // DcId 接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // PairingKey 云平台配对密钥。
    // 用于按密钥查询对应接入点的 VLAN 使用情况。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeBytePlusVlanUsageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBytePlusVlanUsageResponseParams `json:"response,omitempty"`

}

// DescribeBytePlusVlanUsageResponseParams 
type DescribeBytePlusVlanUsageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start 可用的 VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End 可用的 VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // UsedVlans 已使用的 VLAN 列表。
    UsedVlans []int `json:"usedVlans,omitempty"`

}

// ModifyCloudBandwidthRequest 
type ModifyCloudBandwidthRequest struct {
    *common.BaseRequest

    // CloudPortId 需要修改的云连接 ID。
    CloudPortId *string `json:"cloudPortId,omitempty"`

    // BandwidthMbps 需要修改的带宽限速。
    // 单位 Mbps，需匹配可用带宽阶梯。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

}

type ModifyCloudBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeCloudAvailableBandwidthTiersRequest 
type DescribeCloudAvailableBandwidthTiersRequest struct {
    *common.BaseRequest

    // CloudType 云连接类型（AWS、TENCENT、GOOGLE、AZURE、ALI_CLOUD、HUAWEI_CLOUD、BYTE_PLUS、ORACLE）。
    CloudType *string `json:"cloudType,omitempty"`

    // CloudRegionId 公有云区域 ID（Google 云无需传参）。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // DcId 连接云接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // VlanId VLAN ID。
    VlanId *int `json:"vlanId,omitempty"`

    // CloudPortId 云连接 ID。
    // 修改云连接带宽限速时使用，传入后其他参数将被忽略。
    CloudPortId *string `json:"cloudPortId,omitempty"`

    // CloudAccountId 云平台账号（Google 为 Pairing Key，Azure 为 Service Key）。
    CloudAccountId *string `json:"cloudAccountId,omitempty"`

    // ZoneColor 可用区标识。
    ZoneColor *string `json:"zoneColor,omitempty"`

}

type DescribeCloudAvailableBandwidthTiersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCloudAvailableBandwidthTiersResponseParams `json:"response,omitempty"`

}

// DescribeCloudAvailableBandwidthTiersResponseParams 
type DescribeCloudAvailableBandwidthTiersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // AvailableBandwidthTiers 云连接可用的带宽阶梯列表。
    AvailableBandwidthTiers []int `json:"availableBandwidthTiers,omitempty"`

}

// DescribeAWSRegionsRequest 
type DescribeAWSRegionsRequest struct {
    *common.BaseRequest

    // Product 筛选云节点支持的产品。
    // 可用值：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Product *string `json:"product,omitempty"`

}

type DescribeAWSRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAWSRegionsResponseParams `json:"response,omitempty"`

}

// DescribeAWSRegionsResponseParams 
type DescribeAWSRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRegions 接入点相关的区域信息列表。
    CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`

}

// CloudRegion 描述公有云接入区域及其支持的产品、高可用类型信息。
type CloudRegion struct {

    // CloudRegionId 公有云区域 ID。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // DataCenter 接入点数据中心信息。
    DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

    // Products 该区域支持的产品列表：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Products []string `json:"products,omitempty"`

    // HaTypes 该区域支持的高可用类型列表。
    HaTypes []*HaTypeInfo `json:"haTypes,omitempty"`

}

// DatacenterInfo 数据中心的基本信息。
type DatacenterInfo struct {

    // DcId 数据中心ID。
    DcId *string `json:"dcId,omitempty"`

    // DcName 数据中心名称。
    DcName *string `json:"dcName,omitempty"`

    // DcAddress 数据中心地址。
    DcAddress *string `json:"dcAddress,omitempty"`

    // CityName 数据中心所在城市名称。
    CityName *string `json:"cityName,omitempty"`

    // CountryName 数据中心所在国家名称。
    CountryName *string `json:"countryName,omitempty"`

    // AreaName 数据中心所在区域名称。
    AreaName *string `json:"areaName,omitempty"`

    // Latitude 数据中心所在地纬度。
    Latitude *float64 `json:"latitude,omitempty"`

    // Longitude 数据中心所在地经度。
    Longitude *float64 `json:"longitude,omitempty"`

}

// HaTypeInfo 描述区域内某高可用类型及其使用情况。
type HaTypeInfo struct {

    // HaType 高可用类型。
    HaType *string `json:"haType,omitempty"`

    // IsUsed 该高可用类型是否已被使用。
    IsUsed *bool `json:"isUsed,omitempty"`

}

// DescribeTencentRegionsRequest 
type DescribeTencentRegionsRequest struct {
    *common.BaseRequest

    // Product 筛选云节点支持的产品。
    // 可用值：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Product *string `json:"product,omitempty"`

}

type DescribeTencentRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeTencentRegionsResponseParams `json:"response,omitempty"`

}

// DescribeTencentRegionsResponseParams 
type DescribeTencentRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRegions 接入点相关的区域信息列表。
    CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`

}

// DescribeGoogleRegionsRequest 
type DescribeGoogleRegionsRequest struct {
    *common.BaseRequest

    // Product 筛选云节点支持的产品。
    // 可用值：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Product *string `json:"product,omitempty"`

    // PairingKey Google 配对密钥。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeGoogleRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeGoogleRegionsResponseParams `json:"response,omitempty"`

}

// DescribeGoogleRegionsResponseParams 
type DescribeGoogleRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRegions 接入点相关的区域信息列表。
    CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`

}

// DescribeAzureRegionsRequest 
type DescribeAzureRegionsRequest struct {
    *common.BaseRequest

    // Product 筛选云节点支持的产品。
    // 可用值：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Product *string `json:"product,omitempty"`

    // PairingKey Azure 配对密钥。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeAzureRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAzureRegionsResponseParams `json:"response,omitempty"`

}

// DescribeAzureRegionsResponseParams 
type DescribeAzureRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRegions 接入点相关的区域信息列表。
    CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`

}

// DescribeOracleRegionsRequest 
type DescribeOracleRegionsRequest struct {
    *common.BaseRequest

    // Product 筛选云节点支持的产品。
    // 可用值：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Product *string `json:"product,omitempty"`

    // OcId Oracle 接入账号 ID。
    OcId *string `json:"ocId,omitempty"`

}

type DescribeOracleRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeOracleRegionsResponseParams `json:"response,omitempty"`

}

// DescribeOracleRegionsResponseParams 
type DescribeOracleRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRegions 接入点相关的区域信息列表。
    CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`

}

// DescribeAliCloudRegionsRequest 
type DescribeAliCloudRegionsRequest struct {
    *common.BaseRequest

    // Product 筛选云节点支持的产品。
    // 可用值：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Product *string `json:"product,omitempty"`

}

type DescribeAliCloudRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAliCloudRegionsResponseParams `json:"response,omitempty"`

}

// DescribeAliCloudRegionsResponseParams 
type DescribeAliCloudRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRegions 接入点相关的区域信息列表。
    CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`

}

// DescribeHuaweiCloudRegionsRequest 
type DescribeHuaweiCloudRegionsRequest struct {
    *common.BaseRequest

    // Product 筛选云节点支持的产品。
    // 可用值：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Product *string `json:"product,omitempty"`

}

type DescribeHuaweiCloudRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeHuaweiCloudRegionsResponseParams `json:"response,omitempty"`

}

// DescribeHuaweiCloudRegionsResponseParams 
type DescribeHuaweiCloudRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRegions 接入点相关的区域信息列表。
    CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`

}

// DescribeBytePlusRegionsRequest 
type DescribeBytePlusRegionsRequest struct {
    *common.BaseRequest

    // Product 筛选云节点支持的产品。
    // 可用值：PrivateConnect(二层网络)、CloudRouter(三层网络)。
    Product *string `json:"product,omitempty"`

}

type DescribeBytePlusRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBytePlusRegionsResponseParams `json:"response,omitempty"`

}

// DescribeBytePlusRegionsResponseParams 
type DescribeBytePlusRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRegions 接入点相关的区域信息列表。
    CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`

}

// DescribeAWSVlanUsageRequest 
type DescribeAWSVlanUsageRequest struct {
    *common.BaseRequest

    // DcId 接入点的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // PairingKey 云平台配对密钥。
    // 用于按密钥查询对应接入点的 VLAN 使用情况。
    PairingKey *string `json:"pairingKey,omitempty"`

}

type DescribeAWSVlanUsageResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAWSVlanUsageResponseParams `json:"response,omitempty"`

}

// DescribeAWSVlanUsageResponseParams 
type DescribeAWSVlanUsageResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Start 可用的 VLAN 范围起始值。
    Start *int `json:"start,omitempty"`

    // End 可用的 VLAN 范围结束值。
    End *int `json:"end,omitempty"`

    // UsedVlans 已使用的 VLAN 列表。
    UsedVlans []int `json:"usedVlans,omitempty"`

}

// CreateCloudRouterRequest 
type CreateCloudRouterRequest struct {
    *common.BaseRequest

    // CloudRouterName 三层网络的名称。
    // 长度不超过 255，不传则默认为 cloud-router-{时间戳}。
    CloudRouterName *string `json:"cloudRouterName,omitempty"`

    // CloudRouterDescription 三层网络的描述信息。
    // 长度不超过 255。
    CloudRouterDescription *string `json:"cloudRouterDescription,omitempty"`

    // EdgePoints 边缘连接点信息。
    // 至少需要 2 个连接点。
    EdgePoints []*CreateCloudRouterEdgePoint `json:"edgePoints,omitempty"`

    // ResourceGroupId 资源组 ID。
    // 不传则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

    // Tags 创建三层网络时关联的标签。
    // 注意：关联`标签键`不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

// CreateCloudRouterEdgePoint 创建三层网络边缘连接点的配置信息。
type CreateCloudRouterEdgePoint struct {

    // EdgePointName 连接点的名称。
    // 仅支持字母、数字和短横线，长度不超过 100 字符。
    EdgePointName *string `json:"edgePointName,omitempty"`

    // VpcId VPC 连接点的 VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

    // ZbgId ZBG 连接点的 ZBG ID。
    ZbgId *string `json:"zbgId,omitempty"`

    // PortId 数据中心端口连接点的端口 ID。
    PortId *string `json:"portId,omitempty"`

    // CloudType 云连接点的云厂商类型。
    CloudType *string `json:"cloudType,omitempty"`

    // CloudAccountId 云连接点的云账号 ID。
    CloudAccountId *string `json:"cloudAccountId,omitempty"`

    // DcId 连接点所在的数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // CloudRegionId 云连接点的云地域 ID。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // HaType 云连接点的高可用类型。
    HaType *string `json:"haType,omitempty"`

    // VlanId 连接点的 VLAN ID。
    // 取值范围 2-4000。
    VlanId *int `json:"vlanId,omitempty"`

    // BandwidthMbps 连接点的骨干带宽限速。
    // 单位 Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // CloudBandwidthMbps 云连接点的云侧带宽限速。
    // 单位 Mbps。
    CloudBandwidthMbps *int `json:"cloudBandwidthMbps,omitempty"`

    // IpAddress 连接点的互联 IP 地址及掩码。
    IpAddress *string `json:"ipAddress,omitempty"`

    // BgpConnection 连接点的 BGP 连接配置。
    // 与静态路由二选一。
    BgpConnection *BGPConnection `json:"bgpConnection,omitempty"`

    // StaticRoutes 连接点的静态路由配置。
    // 与 BGP 连接二选一。
    StaticRoutes []*IPRoute `json:"staticRoutes,omitempty"`

    // IpSecTunnel IPSec 隧道类型。
    // 取值 FQDN 或 RemoteIP。
    IpSecTunnel *string `json:"ipSecTunnel,omitempty"`

    // CustomerPublicIP IPSec 客户侧公网 IP 地址。
    CustomerPublicIP *string `json:"customerPublicIP,omitempty"`

    // CustomerPrivateIP IPSec 客户侧私网 IP 地址。
    CustomerPrivateIP *string `json:"customerPrivateIP,omitempty"`

    // VirtualEdgePrivateIP IPSec Virtual Edge 侧私网 IP 地址。
    VirtualEdgePrivateIP *string `json:"virtualEdgePrivateIP,omitempty"`

    // Psk IPSec 预共享密钥（PSK）。
    Psk *string `json:"psk,omitempty"`

    // EnableHealthCheck 是否开启健康检查。
    // 默认不开启。
    EnableHealthCheck *bool `json:"enableHealthCheck,omitempty"`

    // BackupIpSec IPSec 备份隧道配置。
    BackupIpSec *BackupIPSecConfig `json:"backupIpSec,omitempty"`

    // IpSecBgpConnection IPSec 的 BGP 连接配置。
    // 与 IPSec 静态路由二选一。
    IpSecBgpConnection *IPSecBGPConnection `json:"ipSecBgpConnection,omitempty"`

    // IpSecStaticRoutes IPSec 的静态路由配置。
    // 与 IPSec BGP 连接二选一。
    IpSecStaticRoutes []*IPSecStaticRoute `json:"ipSecStaticRoutes,omitempty"`

}

// BGPConnection BGP 连接的配置信息。
type BGPConnection struct {

    // PeerIpAddress 对端互联 IP 地址及掩码。
    PeerIpAddress *string `json:"peerIpAddress,omitempty"`

    // PeerAsn 对端 BGP ASN。
    PeerAsn *int64 `json:"peerAsn,omitempty"`

    // Password BGP 会话的 MD5 密码。
    // 长度不超过 32 字符。
    Password *string `json:"password,omitempty"`

    // LocalAsn 本端 BGP ASN。
    LocalAsn *int64 `json:"localAsn,omitempty"`

}

// IPRoute 静态路由的配置信息。
type IPRoute struct {

    // Prefix 静态路由的目的网段（CIDR）。
    Prefix *string `json:"prefix,omitempty"`

    // NextHop 静态路由的下一跳 IP 地址。
    NextHop *string `json:"nextHop,omitempty"`

}

// BackupIPSecConfig IPSec 备份隧道的配置信息。
type BackupIPSecConfig struct {

    // DcId 备份 IPSec 隧道所在数据中心的 ID。
    DcId *string `json:"dcId,omitempty"`

    // Password 备份 IPSec 隧道的认证密码。
    // 长度 8-64 个字符。
    Password *string `json:"password,omitempty"`

    // CustomerPrivateIP 备份 IPSec 隧道客户侧私网 IP 地址。
    CustomerPrivateIP *string `json:"customerPrivateIP,omitempty"`

    // VirtualEdgePrivateIP 备份 IPSec 隧道 Virtual Edge 侧私网 IP 地址。
    VirtualEdgePrivateIP *string `json:"virtualEdgePrivateIP,omitempty"`

    // Psk 备份 IPSec 隧道的预共享密钥（PSK）。
    Psk *string `json:"psk,omitempty"`

}

// IPSecBGPConnection IPSec 的 BGP 连接配置信息。
type IPSecBGPConnection struct {

    // CustomerAsn 客户侧 BGP ASN。
    CustomerAsn *int64 `json:"customerAsn,omitempty"`

    // VirtualEdgeAsn Virtual Edge 侧 BGP ASN。
    VirtualEdgeAsn *int64 `json:"virtualEdgeAsn,omitempty"`

    // Password BGP 会话的认证密码。
    // 长度 8-64 个字符。
    Password *string `json:"password,omitempty"`

}

// IPSecStaticRoute IPSec 的静态路由配置信息。
type IPSecStaticRoute struct {

    // Cidr 静态路由的目的网段。
    // CIDR 格式。
    Cidr *string `json:"cidr,omitempty"`

}

type CreateCloudRouterResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateCloudRouterResponseParams `json:"response,omitempty"`

}

// CreateCloudRouterResponseParams 
type CreateCloudRouterResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CloudRouterId 三层网络的 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

}

// DeleteCloudRouterEdgePointRequest 
type DeleteCloudRouterEdgePointRequest struct {
    *common.BaseRequest

    // CloudRouterId 连接点所在的三层网络 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

    // EdgePointId 要移除的连接点 ID。
    EdgePointId *string `json:"edgePointId,omitempty"`

}

type DeleteCloudRouterEdgePointResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// AddCloudRouterEdgePointsRequest 
type AddCloudRouterEdgePointsRequest struct {
    *common.BaseRequest

    // CloudRouterId 三层网络的 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

    // EdgePoints 新增的边缘连接点信息。
    EdgePoints []*CreateCloudRouterEdgePoint `json:"edgePoints,omitempty"`

}

type AddCloudRouterEdgePointsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *AddCloudRouterEdgePointsResponseParams `json:"response,omitempty"`

}

// AddCloudRouterEdgePointsResponseParams 
type AddCloudRouterEdgePointsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // EdgePointIds 新增的连接点 ID 列表。
    // 顺序与请求中的连接点保持一致。
    EdgePointIds []string `json:"edgePointIds,omitempty"`

}

// ModifyCloudRoutersAttributeRequest 
type ModifyCloudRoutersAttributeRequest struct {
    *common.BaseRequest

    // CloudRouterIds 三层网络的 ID 列表。
    // 数量不超过 100。
    CloudRouterIds []string `json:"cloudRouterIds,omitempty"`

    // CloudRouterName 三层网络名称。
    // 长度不超过 255，名称和描述至少指定其一。
    CloudRouterName *string `json:"cloudRouterName,omitempty"`

    // CloudRouterDescription 三层网络描述信息。
    // 长度不超过 255，名称和描述至少指定其一。
    CloudRouterDescription *string `json:"cloudRouterDescription,omitempty"`

}

type ModifyCloudRoutersAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeCloudRouterAvailableVpcsRequest 
type DescribeCloudRouterAvailableVpcsRequest struct {
    *common.BaseRequest

    // VpcId VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为 20，最大为 100。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页页码。
    // 默认为 1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeCloudRouterAvailableVpcsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCloudRouterAvailableVpcsResponseParams `json:"response,omitempty"`

}

// DescribeCloudRouterAvailableVpcsResponseParams 
type DescribeCloudRouterAvailableVpcsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的 VPC 总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 可加入三层网络的 VPC 列表。
    DataSet []*CloudRouterAvailableVpc `json:"dataSet,omitempty"`

}

// CloudRouterAvailableVpc 可加入三层网络的 VPC 信息。
type CloudRouterAvailableVpc struct {

    // VpcId VPC 的 ID。
    VpcId *string `json:"vpcId,omitempty"`

    // VpcName VPC 的名称。
    VpcName *string `json:"vpcName,omitempty"`

    // DataCenter VPC 所在的数据中心信息。
    DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

    // CidrBlock VPC 的 CIDR 网段。
    CidrBlock *string `json:"cidrBlock,omitempty"`

}

// DescribeCloudRouterEdgePointTrafficRequest 
type DescribeCloudRouterEdgePointTrafficRequest struct {
    *common.BaseRequest

    // EdgePointId 三层网络连接点的 ID。
    EdgePointId *string `json:"edgePointId,omitempty"`

    // StartTime 查询开始时间。
    // ISO8601 格式，UTC 时区。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // ISO8601 格式，UTC 时区，默认为当前时间。
    EndTime *string `json:"endTime,omitempty"`

}

type DescribeCloudRouterEdgePointTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCloudRouterEdgePointTrafficResponseParams `json:"response,omitempty"`

}

// DescribeCloudRouterEdgePointTrafficResponseParams 
type DescribeCloudRouterEdgePointTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 流量明细数据列表。
    DataList []*TrafficData `json:"dataList,omitempty"`

    // In95 入方向 95 峰值流量。
    In95 *int64 `json:"in95,omitempty"`

    // InAvg 入方向平均流量。
    InAvg *int64 `json:"inAvg,omitempty"`

    // InMax 入方向最大流量。
    InMax *int64 `json:"inMax,omitempty"`

    // InMin 入方向最小流量。
    InMin *int64 `json:"inMin,omitempty"`

    // Out95 出方向 95 峰值流量。
    Out95 *int64 `json:"out95,omitempty"`

    // OutAvg 出方向平均流量。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出方向最大流量。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出方向最小流量。
    OutMin *int64 `json:"outMin,omitempty"`

    // Unit 流量单位。
    // 如 bps。
    Unit *string `json:"unit,omitempty"`

}

// DescribeCloudRouterDCToDCTrafficRequest 
type DescribeCloudRouterDCToDCTrafficRequest struct {
    *common.BaseRequest

    // CloudRouterId 三层网络的 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

    // SourceDcId 源数据中心 ID。
    SourceDcId *string `json:"sourceDcId,omitempty"`

    // DestinationDcId 目的数据中心 ID。
    DestinationDcId *string `json:"destinationDcId,omitempty"`

    // StartTime 查询开始时间。
    // ISO8601 格式，UTC 时区。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // ISO8601 格式，UTC 时区，默认为当前时间。
    EndTime *string `json:"endTime,omitempty"`

}

type DescribeCloudRouterDCToDCTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCloudRouterDCToDCTrafficResponseParams `json:"response,omitempty"`

}

// DescribeCloudRouterDCToDCTrafficResponseParams 
type DescribeCloudRouterDCToDCTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 流量明细数据列表。
    DataList []*TrafficData `json:"dataList,omitempty"`

    // In95 入方向 95 峰值流量。
    In95 *int64 `json:"in95,omitempty"`

    // InAvg 入方向平均流量。
    InAvg *int64 `json:"inAvg,omitempty"`

    // InMax 入方向最大流量。
    InMax *int64 `json:"inMax,omitempty"`

    // InMin 入方向最小流量。
    InMin *int64 `json:"inMin,omitempty"`

    // Out95 出方向 95 峰值流量。
    Out95 *int64 `json:"out95,omitempty"`

    // OutAvg 出方向平均流量。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出方向最大流量。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出方向最小流量。
    OutMin *int64 `json:"outMin,omitempty"`

    // Unit 流量单位。
    // 如 bps。
    Unit *string `json:"unit,omitempty"`

}

// ModifyCloudRouterEdgePointBandwidthRequest 
type ModifyCloudRouterEdgePointBandwidthRequest struct {
    *common.BaseRequest

    // CloudRouterId 连接点关联的三层网络 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

    // EdgePointId 三层网络连接点的 ID。
    EdgePointId *string `json:"edgePointId,omitempty"`

    // BandwidthMbps 需要修改的带宽限速。
    // 单位 Mbps，取值范围 1-500。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

}

type ModifyCloudRouterEdgePointBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyCloudRouterEdgePointRequest 
type ModifyCloudRouterEdgePointRequest struct {
    *common.BaseRequest

    // CloudRouterId 连接点关联的三层网络 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

    // EdgePointId 三层网络连接点的 ID。
    EdgePointId *string `json:"edgePointId,omitempty"`

    // BandwidthMbps 需要修改的带宽限速。
    // 单位 Mbps，取值范围 1-500。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // IpAddress 连接点的互联 IP 地址。
    IpAddress *string `json:"ipAddress,omitempty"`

    // BgpConnection BGP 连接配置。
    // 与静态路由二选一。
    BgpConnection *BGPConnection `json:"bgpConnection,omitempty"`

    // StaticRoutes 静态路由配置。
    // 与 BGP 连接二选一。
    StaticRoutes []*IPRoute `json:"staticRoutes,omitempty"`

}

type ModifyCloudRouterEdgePointResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteCloudRouterRequest 
type DeleteCloudRouterRequest struct {
    *common.BaseRequest

    // CloudRouterId 三层网络的 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

}

type DeleteCloudRouterResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DestroyCloudRouterRequest 
type DestroyCloudRouterRequest struct {
    *common.BaseRequest

    // CloudRouterId 三层网络的 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

}

type DestroyCloudRouterResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// RenewCloudRouterRequest 
type RenewCloudRouterRequest struct {
    *common.BaseRequest

    // CloudRouterId 三层网络的 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

}

type RenewCloudRouterResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeCloudRouterAvailablePortsRequest 
type DescribeCloudRouterAvailablePortsRequest struct {
    *common.BaseRequest

    // PortIds 数据中心端口 ID 列表。
    // 用于按数据中心端口过滤。
    PortIds []string `json:"portIds,omitempty"`

    // DcId 数据中心 ID。
    // 用于按数据中心过滤。
    DcId *string `json:"dcId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为 20，最大为 100。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页页码。
    // 默认为 1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeCloudRouterAvailablePortsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCloudRouterAvailablePortsResponseParams `json:"response,omitempty"`

}

// DescribeCloudRouterAvailablePortsResponseParams 
type DescribeCloudRouterAvailablePortsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据中心端口总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 可加入网络的数据中心端口列表。
    DataSet []*PortInfo `json:"dataSet,omitempty"`

}

// DescribeCloudRoutersRequest 
type DescribeCloudRoutersRequest struct {
    *common.BaseRequest

    // CloudRouterIds 三层网络的 ID 列表。
    // 长度不超过 100 个。
    CloudRouterIds []string `json:"cloudRouterIds,omitempty"`

    // CloudRouterName 三层网络名称。
    // 支持模糊匹配。
    CloudRouterName *string `json:"cloudRouterName,omitempty"`

    // CloudRouterStatus 三层网络的状态。
    CloudRouterStatus *string `json:"cloudRouterStatus,omitempty"`

    // EdgePointId 边缘连接点的 ID。
    EdgePointId *string `json:"edgePointId,omitempty"`

    // ResourceGroupId 资源组 ID。
    // 不传则返回所有可见资源组内的三层网络。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为 20，最大为 1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页页码。
    // 默认为 1。
    PageNum *int `json:"pageNum,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeCloudRoutersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeCloudRoutersResponseParams `json:"response,omitempty"`

}

// DescribeCloudRoutersResponseParams 
type DescribeCloudRoutersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的三层网络总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 三层网络列表。
    DataSet []*CloudRouter `json:"dataSet,omitempty"`

}

// CloudRouter 三层网络的详细信息。
type CloudRouter struct {

    // CloudRouterId 三层网络的 ID。
    CloudRouterId *string `json:"cloudRouterId,omitempty"`

    // CloudRouterName 三层网络的名称。
    CloudRouterName *string `json:"cloudRouterName,omitempty"`

    // CloudRouterDescription 三层网络的描述信息。
    CloudRouterDescription *string `json:"cloudRouterDescription,omitempty"`

    // CreateTime 三层网络的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 三层网络的过期时间。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // RecycledTime 三层网络的回收时间。
    RecycledTime *string `json:"recycledTime,omitempty"`

    // EdgePoints 三层网络的边缘连接点列表。
    EdgePoints []*CloudRouterEdgePoint `json:"edgePoints,omitempty"`

    // CloudRouterStatus 三层网络的业务状态。
    CloudRouterStatus *string `json:"cloudRouterStatus,omitempty"`

    // ConnectivityStatus 三层网络的连通性状态。
    ConnectivityStatus *string `json:"connectivityStatus,omitempty"`

    // ResourceGroupId 三层网络所属的资源组 ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 三层网络所属的资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // Period 三层网络的购买时长（月）。
    Period *int `json:"period,omitempty"`

    // Source 三层网络的业务来源。
    Source *string `json:"source,omitempty"`

    // Tags 三层网络关联的标签。
    Tags *Tags `json:"tags,omitempty"`

}

// CloudRouterEdgePoint 三层网络边缘连接点的详细信息。
type CloudRouterEdgePoint struct {

    // EdgePointId 连接点的 ID。
    EdgePointId *string `json:"edgePointId,omitempty"`

    // EdgePointName 连接点的名称。
    EdgePointName *string `json:"edgePointName,omitempty"`

    // ConnectivityStatus 连接点的连通性状态。
    ConnectivityStatus *string `json:"connectivityStatus,omitempty"`

    // DataCenter 连接点所在的数据中心信息。
    DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

    // EdgePointType 连接点的类型。
    EdgePointType *string `json:"edgePointType,omitempty"`

    // CloudType 云连接点的云厂商类型。
    CloudType *string `json:"cloudType,omitempty"`

    // CloudAccountId 云连接点的云账号 ID。
    CloudAccountId *string `json:"cloudAccountId,omitempty"`

    // CloudRegionId 云连接点的云地域 ID。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // VpcId VPC 连接点的 VPC ID。
    VpcId *string `json:"vpcId,omitempty"`

    // ZbgId ZBG 连接点的 ZBG ID。
    ZbgId *string `json:"zbgId,omitempty"`

    // ZbgRegionId ZBG 连接点所属的region。
    ZbgRegionId *string `json:"zbgRegionId,omitempty"`

    // PortId 数据中心端口连接点的端口 ID。
    PortId *string `json:"portId,omitempty"`

    // SharedChannelId 共享通道ID。
    // 公有云接入点对应的云侧连接标识，仅公有云类型接入点时有值。
    SharedChannelId *string `json:"sharedChannelId,omitempty"`

    // VlanId 连接点的 VLAN ID。
    // 取值范围 2-4000。
    VlanId *int `json:"vlanId,omitempty"`

    // BandwidthMbps 连接点的骨干带宽限速。
    // 单位 Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // BgpConnection 连接点的 BGP 连接配置。
    BgpConnection *BGPConnection `json:"bgpConnection,omitempty"`

    // StaticRoutes 连接点的静态路由配置。
    StaticRoutes []*IPRoute `json:"staticRoutes,omitempty"`

    // CreateTime 连接点的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // IpAddress 连接点的互联 IP 地址。
    IpAddress *string `json:"ipAddress,omitempty"`

}

// DescribeDatacentersRequest 
type DescribeDatacentersRequest struct {
    *common.BaseRequest

    // DcIds 数据中心ID列表。
    DcIds []string `json:"dcIds,omitempty"`

    // IsPortAvailable 筛选是否支持新建数据中心端口的 DC（true：支持，false：不支持）。
    IsPortAvailable *bool `json:"isPortAvailable,omitempty"`

}

type DescribeDatacentersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeDatacentersResponseParams `json:"response,omitempty"`

}

// DescribeDatacentersResponseParams 
type DescribeDatacentersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DcSet 结果集。
    DcSet []*DataCenterAssemble `json:"dcSet,omitempty"`

}

// DataCenterAssemble 描述数据中心的基本信息，包括位置、坐标及数据中心端口可用带宽等。
type DataCenterAssemble struct {

    // DcId 数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // DcName 数据中心名称。
    DcName *string `json:"dcName,omitempty"`

    // DcAddress 数据中心地址。
    DcAddress *string `json:"dcAddress,omitempty"`

    // CityName 所在城市名称。
    CityName *string `json:"cityName,omitempty"`

    // CountryName 所在国家名称。
    CountryName *string `json:"countryName,omitempty"`

    // AreaName 所在大区名称。
    AreaName *string `json:"areaName,omitempty"`

    // Latitude 纬度。
    Latitude *float64 `json:"latitude,omitempty"`

    // Longitude 经度。
    Longitude *float64 `json:"longitude,omitempty"`

    // IsPortAvailable 是否支持新建数据中心端口。
    IsPortAvailable *bool `json:"isPortAvailable,omitempty"`

    // AvailableBandwidth 可用带宽（Mbps）。
    AvailableBandwidth *int `json:"availableBandwidth,omitempty"`

    // RegionId 节点 ID。
    RegionId *string `json:"regionId,omitempty"`

}

// DescribeVirtualEdgeDatacentersRequest 
type DescribeVirtualEdgeDatacentersRequest struct {
    *common.BaseRequest

    // PrimaryDcId 主数据中心 ID。
    // IPSec 接入类型且需要高可用时，用于过滤可用备机房。
    PrimaryDcId *string `json:"primaryDcId,omitempty"`

}

type DescribeVirtualEdgeDatacentersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeVirtualEdgeDatacentersResponseParams `json:"response,omitempty"`

}

// DescribeVirtualEdgeDatacentersResponseParams 
type DescribeVirtualEdgeDatacentersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DcSet 结果集。
    DcSet []*DataCenterAssemble `json:"dcSet,omitempty"`

}

// DescribeBorderGatewayDatacentersRequest 
type DescribeBorderGatewayDatacentersRequest struct {
    *common.BaseRequest

    // DcId 数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // RegionId 节点 ID。
    RegionId *string `json:"regionId,omitempty"`

}

type DescribeBorderGatewayDatacentersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBorderGatewayDatacentersResponseParams `json:"response,omitempty"`

}

// DescribeBorderGatewayDatacentersResponseParams 
type DescribeBorderGatewayDatacentersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DcSet 结果集。
    DcSet []*DataCenterAssemble `json:"dcSet,omitempty"`

}

// DescribeVPCDatacentersRequest 
type DescribeVPCDatacentersRequest struct {
    *common.BaseRequest

    // DcId 数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // RegionId VPC 的节点 ID。
    RegionId *string `json:"regionId,omitempty"`

}

type DescribeVPCDatacentersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeVPCDatacentersResponseParams `json:"response,omitempty"`

}

// DescribeVPCDatacentersResponseParams 
type DescribeVPCDatacentersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DcSet 结果集。
    DcSet []*DataCenterAssemble `json:"dcSet,omitempty"`

}

// DescribeDatacentersWithServiceRequest 
type DescribeDatacentersWithServiceRequest struct {
    *common.BaseRequest

    // CloudRegionId 公有云区域 ID。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // DcId 数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // ServiceType 服务类型。
    ServiceType *string `json:"serviceType,omitempty"`

}

type DescribeDatacentersWithServiceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeDatacentersWithServiceResponseParams `json:"response,omitempty"`

}

// DescribeDatacentersWithServiceResponseParams 
type DescribeDatacentersWithServiceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DcSet 结果集。
    DcSet []*DataCenterWithServiceResponse `json:"dcSet,omitempty"`

}

// DataCenterWithServiceResponse 带服务类型信息的数据中心详情。
type DataCenterWithServiceResponse struct {

    // DcId 数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // DcName 数据中心名称。
    DcName *string `json:"dcName,omitempty"`

    // DcAddress 数据中心地址。
    DcAddress *string `json:"dcAddress,omitempty"`

    // CityName 所在城市名称。
    CityName *string `json:"cityName,omitempty"`

    // CountryName 所在国家名称。
    CountryName *string `json:"countryName,omitempty"`

    // AreaName 所在大区名称。
    AreaName *string `json:"areaName,omitempty"`

    // Latitude 纬度。
    Latitude *float64 `json:"latitude,omitempty"`

    // Longitude 经度。
    Longitude *float64 `json:"longitude,omitempty"`

    // CloudRegionId 公有云区域 ID。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // ServiceTypes 支持的服务类型集合。
    ServiceTypes []string `json:"serviceTypes,omitempty"`

}

// DescribePrivateConnectsRequest 
type DescribePrivateConnectsRequest struct {
    *common.BaseRequest

    // PrivateConnectIds 二层网络专线 ID 列表。
    // 最多支持 100 个 ID 查询。
    PrivateConnectIds []string `json:"privateConnectIds,omitempty"`

    // PrivateConnectName 二层网络专线名称。
    // 模糊匹配。
    PrivateConnectName *string `json:"privateConnectName,omitempty"`

    // ConnectivityStatus 连通性状态过滤。
    ConnectivityStatus *string `json:"connectivityStatus,omitempty"`

    // PrivateConnectStatus 二层网络专线业务状态过滤。
    PrivateConnectStatus *string `json:"privateConnectStatus,omitempty"`

    // EndpointTypes 连接点类型过滤。
    EndpointTypes []string `json:"endpointTypes,omitempty"`

    // ResourceGroupId 资源组 ID。
    // 不传则返回该用户可见的所有资源组内的二层网络专线。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为 20，最大为 1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    // 默认为 1。
    PageNum *int `json:"pageNum,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribePrivateConnectsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePrivateConnectsResponseParams `json:"response,omitempty"`

}

// DescribePrivateConnectsResponseParams 
type DescribePrivateConnectsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的二层网络专线总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 二层网络专线列表。
    DataSet []*PrivateConnect `json:"dataSet,omitempty"`

}

// PrivateConnect 二层网络专线的详细信息。
type PrivateConnect struct {

    // PrivateConnectId 二层网络专线 ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

    // PrivateConnectName 二层网络专线名称。
    PrivateConnectName *string `json:"privateConnectName,omitempty"`

    // ConnectivityStatus 连通性状态。
    ConnectivityStatus *string `json:"connectivityStatus,omitempty"`

    // PrivateConnectStatus 二层网络专线业务状态。
    PrivateConnectStatus *string `json:"privateConnectStatus,omitempty"`

    // BandwidthMbps 最大带宽限制。
    // 单位：Mbps。
    BandwidthMbps *int64 `json:"bandwidthMbps,omitempty"`

    // ResourceGroupId 资源组 ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ExpiredTime 到期时间。
    ExpiredTime *string `json:"expiredTime,omitempty"`

    // RecycledTime 回收时间。
    RecycledTime *string `json:"recycledTime,omitempty"`

    // EndpointA 二层网络专线一端的连接点（A）。
    EndpointA *PrivateConnectEndpoint `json:"endpointA,omitempty"`

    // EndpointZ 二层网络专线另一端的连接点（Z）。
    EndpointZ *PrivateConnectEndpoint `json:"endpointZ,omitempty"`

    // Source 业务来源。
    Source *string `json:"source,omitempty"`

    // Tags 关联的标签。
    Tags *Tags `json:"tags,omitempty"`

}

// PrivateConnectEndpoint 二层网络专线连接点的详细信息。
type PrivateConnectEndpoint struct {

    // EndpointId 连接点 ID。
    EndpointId *string `json:"endpointId,omitempty"`

    // EndpointName 连接点名称。
    EndpointName *string `json:"endpointName,omitempty"`

    // EndpointType 连接点类型。
    EndpointType *string `json:"endpointType,omitempty"`

    // ConnectivityStatus 连通性状态。
    ConnectivityStatus *string `json:"connectivityStatus,omitempty"`

    // VlanId VLAN ID。
    VlanId *int `json:"vlanId,omitempty"`

    // DataCenter 数据中心信息。
    DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

    // CloudRegionId 云地域 ID。
    // 连接点为云连接时返回。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // CloudAccountId 云账号 ID。
    // 连接点为云连接时返回。
    CloudAccountId *string `json:"cloudAccountId,omitempty"`

    // SharedChannelId 共享通道ID。
    // 公有云接入点对应的云侧连接标识，仅公有云类型接入点时有值。
    SharedChannelId *string `json:"sharedChannelId,omitempty"`

}

// CreatePrivateConnectRequest 
type CreatePrivateConnectRequest struct {
    *common.BaseRequest

    // PrivateConnectName 二层网络专线名称。
    // 长度不超过 255，默认为 private-connect-{当前时间戳}。
    PrivateConnectName *string `json:"privateConnectName,omitempty"`

    // BandwidthMbps 二层网络专线最大带宽限制。
    // 单位：Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // ResourceGroupId 资源组 ID。
    // 不传则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // EndpointA 二层网络专线一端的连接点（A）。
    EndpointA *CreateEndpointParam `json:"endpointA,omitempty"`

    // EndpointZ 二层网络专线另一端的连接点（Z）。
    EndpointZ *CreateEndpointParam `json:"endpointZ,omitempty"`

    // Tags 创建二层网络专线时关联的标签。
    // 注意：关联`标签键`不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

    // MarketingOptions 市场营销活动相关信息。
    MarketingOptions *MarketingInfo `json:"marketingOptions,omitempty"`

}

// CreateEndpointParam 创建二层网络专线时连接点的参数。
type CreateEndpointParam struct {

    // PortId 数据中心端口 ID。
    // 连接点为数据中心端口时使用。
    PortId *string `json:"portId,omitempty"`

    // VlanId VLAN ID。
    VlanId *int `json:"vlanId,omitempty"`

    // CloudRegionId 云地域 ID。
    // 连接点为云连接时使用。
    CloudRegionId *string `json:"cloudRegionId,omitempty"`

    // CloudAccountId 云账号 ID。
    // 连接点为云连接时使用。
    CloudAccountId *string `json:"cloudAccountId,omitempty"`

    // HaType 云连接高可用类型。
    HaType *string `json:"haType,omitempty"`

    // CloudBandwidthMbps 云连接带宽限制。
    // 单位：Mbps。
    CloudBandwidthMbps *int `json:"cloudBandwidthMbps,omitempty"`

    // DcId 数据中心 ID。
    DcId *string `json:"dcId,omitempty"`

    // CloudType 云厂商类型。
    CloudType *string `json:"cloudType,omitempty"`

    // EndpointName 连接点名称。
    // 仅允许字母、数字和连字符。
    EndpointName *string `json:"endpointName,omitempty"`

}

type CreatePrivateConnectResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreatePrivateConnectResponseParams `json:"response,omitempty"`

}

// CreatePrivateConnectResponseParams 
type CreatePrivateConnectResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // PrivateConnectId 二层网络专线 ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

}

// ModifyPrivateConnectBandwidthRequest 
type ModifyPrivateConnectBandwidthRequest struct {
    *common.BaseRequest

    // PrivateConnectId 二层网络专线 ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

    // BandwidthMbps 需要修改的带宽限速。
    // 单位：Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

}

type ModifyPrivateConnectBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribePrivateConnectTrafficRequest 
type DescribePrivateConnectTrafficRequest struct {
    *common.BaseRequest

    // PrivateConnectId 二层网络专线 ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

    // StartTime 查询开始时间。
    // ISO8601 UTC 格式：YYYY-MM-DDThh:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // ISO8601 UTC 格式：YYYY-MM-DDThh:mm:ssZ，默认为当前时间。
    EndTime *string `json:"endTime,omitempty"`

}

type DescribePrivateConnectTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePrivateConnectTrafficResponseParams `json:"response,omitempty"`

}

// DescribePrivateConnectTrafficResponseParams 
type DescribePrivateConnectTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 流量明细数据列表。
    DataList []*TrafficData `json:"dataList,omitempty"`

    // In95 入方向 95 峰值流量。
    In95 *int64 `json:"in95,omitempty"`

    // InAvg 入方向平均流量。
    InAvg *int64 `json:"inAvg,omitempty"`

    // InMax 入方向最大流量。
    InMax *int64 `json:"inMax,omitempty"`

    // InMin 入方向最小流量。
    InMin *int64 `json:"inMin,omitempty"`

    // Out95 出方向 95 峰值流量。
    Out95 *int64 `json:"out95,omitempty"`

    // OutAvg 出方向平均流量。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出方向最大流量。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出方向最小流量。
    OutMin *int64 `json:"outMin,omitempty"`

    // Unit 流量单位。
    // 如 bps。
    Unit *string `json:"unit,omitempty"`

}

// DeletePrivateConnectRequest 
type DeletePrivateConnectRequest struct {
    *common.BaseRequest

    // PrivateConnectId 二层网络专线 ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

}

type DeletePrivateConnectResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DestroyPrivateConnectRequest 
type DestroyPrivateConnectRequest struct {
    *common.BaseRequest

    // PrivateConnectId 二层网络专线 ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

}

type DestroyPrivateConnectResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribePrivateConnectAvailablePortsRequest 
type DescribePrivateConnectAvailablePortsRequest struct {
    *common.BaseRequest

    // PortIds 数据中心端口 ID 列表。
    // 用于按数据中心端口过滤。
    PortIds []string `json:"portIds,omitempty"`

    // DcId 数据中心 ID。
    // 用于按数据中心过滤。
    DcId *string `json:"dcId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为 20，最大为 100。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页页码。
    // 默认为 1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribePrivateConnectAvailablePortsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePrivateConnectAvailablePortsResponseParams `json:"response,omitempty"`

}

// DescribePrivateConnectAvailablePortsResponseParams 
type DescribePrivateConnectAvailablePortsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据中心端口总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 可加入网络的数据中心端口列表。
    DataSet []*PortInfo `json:"dataSet,omitempty"`

}

// ModifyPrivateConnectsAttributeRequest 
type ModifyPrivateConnectsAttributeRequest struct {
    *common.BaseRequest

    // PrivateConnectIds 二层网络专线 ID 列表。
    // 数量不得超过 100。
    PrivateConnectIds []string `json:"privateConnectIds,omitempty"`

    // PrivateConnectName 二层网络专线名称。
    // 长度不得超过 255。
    PrivateConnectName *string `json:"privateConnectName,omitempty"`

}

type ModifyPrivateConnectsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// RenewPrivateConnectRequest 
type RenewPrivateConnectRequest struct {
    *common.BaseRequest

    // PrivateConnectId 二层网络专线 ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

}

type RenewPrivateConnectResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// InquiryCreatePrivateConnectPriceRequest 
type InquiryCreatePrivateConnectPriceRequest struct {
    *common.BaseRequest

    // BandwidthMbps 二层网络专线最大带宽限制。
    // 默认为 1，单位：Mbps。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // InternetType 网络计费方式。
    InternetType *string `json:"internetType,omitempty"`

    // CommitBandwidth 保底带宽。
    // 单位：Mbps。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // EndpointA 二层网络专线一端的连接点（A）。
    EndpointA *CreateEndpointParam `json:"endpointA,omitempty"`

    // EndpointZ 二层网络专线另一端的连接点（Z）。
    EndpointZ *CreateEndpointParam `json:"endpointZ,omitempty"`

}

type InquiryCreatePrivateConnectPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryCreatePrivateConnectPriceResponseParams `json:"response,omitempty"`

}

// InquiryCreatePrivateConnectPriceResponseParams 
type InquiryCreatePrivateConnectPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // PrivateConnectPrice 二层网络专线带宽价格。
    PrivateConnectPrice *PriceItem `json:"privateConnectPrice,omitempty"`

    // PrivateConnectBandwidth 二层网络专线带宽。
    PrivateConnectBandwidth *int `json:"privateConnectBandwidth,omitempty"`

    // EndpointAPrice A 端连接点的建设价格。
    EndpointAPrice *PriceItem `json:"endpointAPrice,omitempty"`

    // EndpointABandwidth A 端连接点带宽。
    EndpointABandwidth *int `json:"endpointABandwidth,omitempty"`

    // EndpointZPrice Z 端连接点的建设价格。
    EndpointZPrice *PriceItem `json:"endpointZPrice,omitempty"`

    // EndpointZBandwidth Z 端连接点带宽。
    EndpointZBandwidth *int `json:"endpointZBandwidth,omitempty"`

}

