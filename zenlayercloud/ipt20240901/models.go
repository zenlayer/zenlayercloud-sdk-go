package ipt

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeIPTransitDatacentersRequest 
type DescribeIPTransitDatacentersRequest struct {
    *common.BaseRequest

    // PeerPortId 对端数据中心端口 ID。
    // 传入时查询以该数据中心端口为接入侧的可连接数据中心列表。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // PeerDcId 对端数据中心 ID。
    PeerDcId *string `json:"peerDcId,omitempty"`

    // ZbgRegionId ZBG 接入节点 ID。
    // 非空时查询以该 ZBG 节点为接入侧的 Router RIPT 可连接 DC 列表。
    ZbgRegionId *string `json:"zbgRegionId,omitempty"`

}

type DescribeIPTransitDatacentersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeIPTransitDatacentersResponseParams `json:"response,omitempty"`

}

// DescribeIPTransitDatacentersResponseParams 
type DescribeIPTransitDatacentersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SupportSet 可连接数据中心结果集。
    SupportSet []*IPTransitDatacenter `json:"supportSet,omitempty"`

}

// IPTransitDatacenter 可连接数据中心信息。
type IPTransitDatacenter struct {

    // DataCenter 数据中心信息。
    DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

    // AvailableRoutingTypes 该数据中心可用的路由模式列表。
    AvailableRoutingTypes []*RemoteIptAvailableRoutingType `json:"availableRoutingTypes,omitempty"`

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

// RemoteIptAvailableRoutingType IP Transit可用路由模式信息。
type RemoteIptAvailableRoutingType struct {

    // RoutingType 路由模式。
    RoutingType *string `json:"routingType,omitempty"`

    // AvailableBgpRouteTypes 可选的 BGP 路由通告类型列表。
    // 仅 `routingType` 为 BGP 时有值。
    AvailableBgpRouteTypes []string `json:"availableBgpRouteTypes,omitempty"`

    // DeliveryType 开通方式。
    DeliveryType *string `json:"deliveryType,omitempty"`

    // PublicInterconnectNetmasks IPv4 公网互联可选掩码列表。
    // 目前仅 30 / 31。
    PublicInterconnectNetmasks []int `json:"publicInterconnectNetmasks,omitempty"`

}

type DescribeIPTransitAvailableAsnsRequest struct {
    *common.BaseRequest

}

type DescribeIPTransitAvailableAsnsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeIPTransitAvailableAsnsResponseParams `json:"response,omitempty"`

}

// DescribeIPTransitAvailableAsnsResponseParams 
type DescribeIPTransitAvailableAsnsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataSet 可用 ASN 列表。
    DataSet []*AsnInfo `json:"dataSet,omitempty"`

}

// AsnInfo 可用 ASN 信息。
type AsnInfo struct {

    // Asn ASN 值。
    Asn *string `json:"asn,omitempty"`

    // AsnType ASN 类型。
    AsnType *string `json:"asnType,omitempty"`

}

// DescribeIPTransitAvailableCidrBlocksRequest 
type DescribeIPTransitAvailableCidrBlocksRequest struct {
    *common.BaseRequest

    // IptDcId 目标数据中心 ID。
    // 传入 `ipUuid` 时可不传，将从该 IP 块所在数据中心自动推导。
    IptDcId *string `json:"iptDcId,omitempty"`

    // RoutingType 路由类型。
    // 不同路由类型下可用掩码范围不同；不传则返回全量掩码。
    RoutingType *string `json:"routingType,omitempty"`

    // ZbgRegionId ZBG 区域 ID。
    // ZBG 场景下必传。
    ZbgRegionId *string `json:"zbgRegionId,omitempty"`

    // IpUuid IP 地址 UUID。
    // 传入后接口会自动推导所属数据中心和路由类型，仅返回掩码长度不小于当前 IP 块的可选项。
    IpUuid *string `json:"ipUuid,omitempty"`

}

type DescribeIPTransitAvailableCidrBlocksResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeIPTransitAvailableCidrBlocksResponseParams `json:"response,omitempty"`

}

// DescribeIPTransitAvailableCidrBlocksResponseParams 
type DescribeIPTransitAvailableCidrBlocksResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Ipv4CidrBlocks 可用 IPv4 CIDR 块列表。
    Ipv4CidrBlocks []*CidrBlock `json:"ipv4CidrBlocks,omitempty"`

    // Ipv6CidrBlocks 可用 IPv6 CIDR 块列表。
    Ipv6CidrBlocks []*CidrBlock `json:"ipv6CidrBlocks,omitempty"`

}

// CidrBlock 可用 CIDR 块信息。
type CidrBlock struct {

    // Netmask 掩码长度。
    // IPv4 范围 24–32，IPv6 范围 48–64。
    Netmask *int `json:"netmask,omitempty"`

    // IpNetworkType IP 网络类型。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

}

// InquiryCreateIPTransitPriceRequest 
type InquiryCreateIPTransitPriceRequest struct {
    *common.BaseRequest

    // PeerPortId 对端数据中心端口 ID。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // IptDcId 本端数据中心 ID。
    // 为空代表本地连接（Local IPT）。
    IptDcId *string `json:"iptDcId,omitempty"`

    // InternetType 网络计费方式。
    InternetType *string `json:"internetType,omitempty"`

    // CommitBandwidth 保底带宽。
    // 单位Mbps。
    // 有且仅当internetType=ByInstanceBandwidth95时该字段必传。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // Bandwidth 带宽限速。
    // 单位Mbps。
    // 最小值不能低于5Mbps。
    // 默认值为5Mbps。
    // 95 计费下必须大于等于 `commitBandwidth`。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // RoutingType 路由模式。
    RoutingType *string `json:"routingType,omitempty"`

    // PublicIPv4BlockSize 公网 IPv4 地址段大小列表。
    // 与 `publicIpList` 互斥，优先级更低。
    PublicIPv4BlockSize []int `json:"publicIPv4BlockSize,omitempty"`

    // BgpRouteType BGP入站路由类型。
    BgpRouteType *string `json:"bgpRouteType,omitempty"`

    // IpType IP 类型（IPV4 / IPV6）。
    // 默认 IPV4。
    IpType *string `json:"ipType,omitempty"`

    // PublicIpList 公网 IP 分配列表。
    // 与 `publicIPv4BlockSize` 互斥，优先级更高。
    // 传此字段时 `publicIPv4BlockSize` 被忽略。
    PublicIpList []*IPTransitIpRequest `json:"publicIpList,omitempty"`

    // ZbgRegionId ZBG 接入节点 ID。
    // 非空时走 Router RIPT 询价流程。
    ZbgRegionId *string `json:"zbgRegionId,omitempty"`

    // HaConfig HA 高可用配置。
    // 非空时询价包含 2 条 VLL 价格。
    HaConfig *IPTransitHaConfig `json:"haConfig,omitempty"`

    // PublicInterconnectNetmask 公网互联块掩码。
    // 非空时响应包含公网互联 IP 块价格。
    PublicInterconnectNetmask *int `json:"publicInterconnectNetmask,omitempty"`

}

// IPTransitIpRequest 公网 IP 分配请求。
type IPTransitIpRequest struct {

    // Netmask CIDR 掩码长度。
    // IPv4 有效范围 24–30，IPv6 有效范围 48–64。
    Netmask *int `json:"netmask,omitempty"`

    // IpType IP 类型（IPV4 / IPV6）。
    IpType *string `json:"ipType,omitempty"`

    // IpNetworkType IP 类型。
    // 默认 BGP_IP（从 IP 池分配）。
    // LOCAL_IP 表示原生 IP。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

    // Amount 购买数量。
    // 指定相同掩码长度的 IP 块数量，默认为 1。
    Amount *int `json:"amount,omitempty"`

}

// IPTransitHaConfig IP Transit HA配置信息。
type IPTransitHaConfig struct {

    // HaMode HA 运行模式。
    HaMode *string `json:"haMode,omitempty"`

    // SecondaryPortId 备链路接入数据中心端口 ID。
    // 与顶层 peerPortId 必须同城不同 DC。
    SecondaryPortId *string `json:"secondaryPortId,omitempty"`

    // SecondaryPortVlanId 备链路数据中心端口 VLAN ID。
    SecondaryPortVlanId *int `json:"secondaryPortVlanId,omitempty"`

}

type InquiryCreateIPTransitPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryCreateIPTransitPriceResponseParams `json:"response,omitempty"`

}

// InquiryCreateIPTransitPriceResponseParams 
type InquiryCreateIPTransitPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // PrivateConnectPrice 二层网络专线价格。
    // 可能为空。
    PrivateConnectPrice *PriceItem `json:"privateConnectPrice,omitempty"`

    // IptBandwidthPrice IP Transit带宽价格。
    IptBandwidthPrice *PriceItem `json:"iptBandwidthPrice,omitempty"`

    // PublicIpPrices 公网 IP 价格列表。
    PublicIpPrices []*IPPrice `json:"publicIpPrices,omitempty"`

    // PublicInterconnectIpPrice 公网互联 IP 价格。
    // 仅 publicInterconnectNetmask 非空时返回。
    PublicInterconnectIpPrice *IPPrice `json:"publicInterconnectIpPrice,omitempty"`

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

// IPPrice IP 价格信息。
type IPPrice struct {

    // Price 价格详情。
    Price *PriceItem `json:"price,omitempty"`

    // Netmask 掩码长度。
    Netmask *int `json:"netmask,omitempty"`

    // Qty 数量。
    Qty *int `json:"qty,omitempty"`

    // IpNetworkType IP 网络类型。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

}

// CreateIPTransitRequest 
type CreateIPTransitRequest struct {
    *common.BaseRequest

    // IptName IP Transit名称。
    IptName *string `json:"iptName,omitempty"`

    // IptDescription IP Transit描述。
    IptDescription *string `json:"iptDescription,omitempty"`

    // PeerPortId 对端数据中心端口 ID。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // PeerPortVlan 对端数据中心端口 VLAN。
    PeerPortVlan *int `json:"peerPortVlan,omitempty"`

    // IptDcId 本端数据中心 ID。
    // 为空代表本地连接（Local IPT）。
    // 传 `haConfig` 创建高可用 IP Transit 时必传。
    IptDcId *string `json:"iptDcId,omitempty"`

    // InternetType 网络计费方式。
    InternetType *string `json:"internetType,omitempty"`

    // CommitBandwidth 保底带宽（Mbps）。
    // 95 计费（internetType=ByInstanceBandwidth95）下必传。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // Bandwidth 带宽（Mbps）。
    // 95 计费（internetType=ByInstanceBandwidth95）下必须大于等于 `commitBandwidth`。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // RoutingType 路由模式。
    RoutingType *string `json:"routingType,omitempty"`

    // PublicIPv4BlockSize 公网 IPv4 地址段大小列表。
    // 与 `publicIpList` 互斥，优先级更低。
    PublicIPv4BlockSize []int `json:"publicIPv4BlockSize,omitempty"`

    // Bfd BFD 配置。
    // 传 `haConfig` 创建高可用 IP Transit 时必传，且后续不允许关闭。
    Bfd *BFDConfig `json:"bfd,omitempty"`

    // ResourceGroupId 资源组 ID。
    // 不传则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Bgp BGP相关配置。
    Bgp *RiptBgpConfig `json:"bgp,omitempty"`

    // Tags 创建CIDR时关联的标签。
    // 注意：关联`标签键`不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

    // PublicIpList 公网 IP 分配列表。
    // 与 `publicIPv4BlockSize` 互斥，优先级更高。
    // 传此字段时 `publicIPv4BlockSize` 被忽略。
    PublicIpList []*IPTransitIpRequest `json:"publicIpList,omitempty"`

    // ZbgRegionId ZBG 接入节点 ID。
    // 非空时走 Router RIPT 流程，与 `haConfig` 互斥。
    // 调用 ~~zec:DescribeInterconnectBorderGatewayRegions~~ 以获取可用的节点信息。
    ZbgRegionId *string `json:"zbgRegionId,omitempty"`

    // HaConfig HA 高可用配置。
    // 非空时走 HA 创建流程，与 `zbgRegionId` 互斥，且此时 `iptDcId` 和 `bfd` 均必传。
    HaConfig *IPTransitHaConfig `json:"haConfig,omitempty"`

    // PublicInterconnectNetmask 公网互联块掩码。
    // 非空启用公网地址互联，仅 BGP / Static 路由支持。
    // 合法值见 ~~DescribeIPTransitDatacenters~~ 响应中 availableRoutingTypes[].publicInterconnectNetmasks。
    PublicInterconnectNetmask *int `json:"publicInterconnectNetmask,omitempty"`

}

// BFDConfig BFD配置。
type BFDConfig struct {

    // BfdTxInterval 发送间隔。
    // 单位：ms。
    // 取值范围：100～60000。
    BfdTxInterval *int `json:"bfdTxInterval,omitempty"`

    // BfdRxInterval 接收间隔。
    // 单位：ms
    // 取值范围：100～60000。
    BfdRxInterval *int `json:"bfdRxInterval,omitempty"`

    // BfdMultiplier 本地检测倍数。
    // 取值范围：3～20。
    BfdMultiplier *int `json:"bfdMultiplier,omitempty"`

}

// RiptBgpConfig BGP相关配置
type RiptBgpConfig struct {

    // RouteType 入站路由类型。
    RouteType *string `json:"routeType,omitempty"`

    // Asn 宣告出站路由的ASN号。
    Asn *int64 `json:"asn,omitempty"`

    // Password 加密认证秘钥。
    Password *string `json:"password,omitempty"`

    // AsnList 宣告出站路由的ASN号列表。
    AsnList []int64 `json:"asnList,omitempty"`

    // AsSetList 宣告出站路由的AS-SET列表。
    AsSetList []string `json:"asSetList,omitempty"`

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

type CreateIPTransitResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateIPTransitResponseParams `json:"response,omitempty"`

}

// CreateIPTransitResponseParams 
type CreateIPTransitResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // IptId IP Transit ID。
    IptId *string `json:"iptId,omitempty"`

}

// DescribeIPTransitsRequest 
type DescribeIPTransitsRequest struct {
    *common.BaseRequest

    // IptIds IP Transit ID 列表。
    // 最多支持 100 个 ID 查询。
    IptIds []string `json:"iptIds,omitempty"`

    // IptName IP Transit名称。
    // 模糊匹配。
    IptName *string `json:"iptName,omitempty"`

    // ResourceGroupId 资源组 ID。
    // 不传则返回该用户可见的所有资源组内的IP Transit。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PeerPortId 对端数据中心端口 ID 过滤。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // IptDcId 本端数据中心 ID 过滤。
    IptDcId *string `json:"iptDcId,omitempty"`

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

type DescribeIPTransitsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeIPTransitsResponseParams `json:"response,omitempty"`

}

// DescribeIPTransitsResponseParams 
type DescribeIPTransitsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet IP Transit结果集。
    DataSet []*IPTransit `json:"dataSet,omitempty"`

}

// IPTransit IP Transit信息。
type IPTransit struct {

    // IptId IP Transit ID。
    IptId *string `json:"iptId,omitempty"`

    // IptName IP Transit名称。
    IptName *string `json:"iptName,omitempty"`

    // IptDescription IP Transit描述。
    IptDescription *string `json:"iptDescription,omitempty"`

    // DataCenter IP Transit所在数据中心。
    DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

    // PeerPortId 对端数据中心端口 ID。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // PeerPortName 对端数据中心端口名称。
    PeerPortName *string `json:"peerPortName,omitempty"`

    // PeerDataCenter 对端数据中心端口所在数据中心。
    PeerDataCenter *DatacenterInfo `json:"peerDataCenter,omitempty"`

    // DeliveryType 开通方式。
    DeliveryType *string `json:"deliveryType,omitempty"`

    // ResourceGroupId 资源组 ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // RoutingType 路由模式。
    RoutingType *string `json:"routingType,omitempty"`

    // InternetType 网络计费方式。
    InternetType *string `json:"internetType,omitempty"`

    // Bandwidth 带宽（Mbps）。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // CommitBandwidth 保底带宽（Mbps）。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // Bfd BFD 配置。
    Bfd *BFDConfig `json:"bfd,omitempty"`

    // Bgp BGP 相关配置。
    Bgp *RiptBgpConfig `json:"bgp,omitempty"`

    // Interconnect 互联地址配置。
    Interconnect *Interconnect `json:"interconnect,omitempty"`

    // PrivateConnectId 关联的 VLL ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

    // PrivateConnectName 关联的 VLL 名称。
    PrivateConnectName *string `json:"privateConnectName,omitempty"`

    // PublicIpv4Addresses 公网 IPv4 地址列表。
    PublicIpv4Addresses []*IPAddress `json:"publicIpv4Addresses,omitempty"`

    // IptStatus 业务状态。
    IptStatus *string `json:"iptStatus,omitempty"`

    // ConnectivityStatus 链路连通性状态。
    ConnectivityStatus *string `json:"connectivityStatus,omitempty"`

    // Tags 该IP Transit关联的标签。
    Tags *Tags `json:"tags,omitempty"`

    // PublicIpAddresses 公网 IP 地址列表。
    PublicIpAddresses []*IPTransitIpAddress `json:"publicIpAddresses,omitempty"`

    // HaMode 高可用模式。
    HaMode *string `json:"haMode,omitempty"`

    // ZbgRegionId ZBG 区域 ID。
    // ZBG 场景下的 IP Transit 将返回此字段。
    ZbgRegionId *string `json:"zbgRegionId,omitempty"`

    // PeerPortType 对端数据中心端口类型。
    PeerPortType *string `json:"peerPortType,omitempty"`

    // HaLinks HA 子链路列表。
    // 非 HA 模式下为 null；HA 模式下含两个子链路对象。
    HaLinks []*HaLink `json:"haLinks,omitempty"`

}

// Interconnect 互联地址配置。
type Interconnect struct {

    // VendorIpv4Address Zenlayer 侧 IPv4 互联地址。
    VendorIpv4Address *string `json:"vendorIpv4Address,omitempty"`

    // CustomerIpv4Address 客户侧 IPv4 互联地址。
    CustomerIpv4Address *string `json:"customerIpv4Address,omitempty"`

    // VendorIpv6Address Zenlayer 侧 IPv6 互联地址。
    VendorIpv6Address *string `json:"vendorIpv6Address,omitempty"`

    // CustomerIpv6Address 客户侧 IPv6 互联地址。
    CustomerIpv6Address *string `json:"customerIpv6Address,omitempty"`

}

// IPAddress 描述IP信息。
type IPAddress struct {

    // IpAddress IP地址。
    IpAddress *string `json:"ipAddress,omitempty"`

    // Netmask 掩码。
    Netmask *int `json:"netmask,omitempty"`

    // GatewayIpAddress 网关IP地址。
    GatewayIpAddress *string `json:"gatewayIpAddress,omitempty"`

}

// Tags 描述资源关联的标签信息。
type Tags struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// IPTransitIpAddress IP Transit公网 IP 地址信息。
type IPTransitIpAddress struct {

    // IpUuid IP 块 UUID。
    // 变更（升降级/删除）时作为 ipUuid 传入。
    IpUuid *string `json:"ipUuid,omitempty"`

    // IpAddress IP 地址（CIDR 表示法，如 192.0.2.0/30）。
    IpAddress *string `json:"ipAddress,omitempty"`

    // Netmask 掩码长度。
    Netmask *int `json:"netmask,omitempty"`

    // GatewayIpAddress 网关 IP。
    GatewayIpAddress *string `json:"gatewayIpAddress,omitempty"`

    // IpType IP 类型（IPV4 / IPV6）。
    IpType *string `json:"ipType,omitempty"`

    // IpNetworkType IP 网络类型（BGP_IP / LOCAL_IP）。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

}

// HaLink HA 子链路信息。
type HaLink struct {

    // IsPrimary 当前是否为主线。
    // ACTIVE_STANDBY 模式下动态反映主备切换状态；ACTIVE_ACTIVE 模式下为 null。
    IsPrimary *bool `json:"isPrimary,omitempty"`

    // IptStatus 子链路业务状态。
    IptStatus *string `json:"iptStatus,omitempty"`

    // ConnectivityStatus 子链路连通性状态。
    ConnectivityStatus *string `json:"connectivityStatus,omitempty"`

    // PrivateConnectId 所属 VLL ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

    // PrivateConnectName 所属 VLL 名称。
    PrivateConnectName *string `json:"privateConnectName,omitempty"`

    // PeerPortId 对端数据中心端口 ID。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // PeerPortName 对端数据中心端口名称。
    PeerPortName *string `json:"peerPortName,omitempty"`

    // PeerDataCenter 数据中心端口所在数据中心。
    PeerDataCenter *DatacenterInfo `json:"peerDataCenter,omitempty"`

    // PeerPortVlan VLAN ID。
    PeerPortVlan *int `json:"peerPortVlan,omitempty"`

    // Interconnect 互联 IP 配置。
    Interconnect *Interconnect `json:"interconnect,omitempty"`

}

// ModifyIPTransitBandwidthRequest 
type ModifyIPTransitBandwidthRequest struct {
    *common.BaseRequest

    // IptId IP Transit ID。
    IptId *string `json:"iptId,omitempty"`

    // Bandwidth 目标带宽（Mbps）。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // CommitBandwidth 保底带宽（Mbps）。
    // 不填则与 bandwidth 相同。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

}

type ModifyIPTransitBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyIPTransitsAttributeRequest 
type ModifyIPTransitsAttributeRequest struct {
    *common.BaseRequest

    // IptIds IP Transit ID 列表。
    // 最多支持 100 个。
    IptIds []string `json:"iptIds,omitempty"`

    // IptName IP Transit名称。
    IptName *string `json:"iptName,omitempty"`

    // IptDescription IP Transit描述。
    IptDescription *string `json:"iptDescription,omitempty"`

}

type ModifyIPTransitsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteIPTransitRequest 
type DeleteIPTransitRequest struct {
    *common.BaseRequest

    // IptId IP Transit ID。
    IptId *string `json:"iptId,omitempty"`

}

type DeleteIPTransitResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeIPTransitTrafficRequest 
type DescribeIPTransitTrafficRequest struct {
    *common.BaseRequest

    // IptId IP Transit ID。
    IptId *string `json:"iptId,omitempty"`

    // StartTime 查询开始时间。
    // ISO8601 UTC 格式：YYYY-MM-DDThh:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // ISO8601 UTC 格式：YYYY-MM-DDThh:mm:ssZ，默认为当前时间。
    EndTime *string `json:"endTime,omitempty"`

}

type DescribeIPTransitTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeIPTransitTrafficResponseParams `json:"response,omitempty"`

}

type DescribeIPTransitTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 带宽数据列表。
    DataList []*TrafficData `json:"dataList,omitempty"`

    // In95 入口带宽95值。
    In95 *int64 `json:"in95,omitempty"`

    // InAvg 入口带宽平均值。
    InAvg *int64 `json:"inAvg,omitempty"`

    // InMax 入口带宽最大值。
    InMax *int64 `json:"inMax,omitempty"`

    // InMin 入口带宽最小值。
    InMin *int64 `json:"inMin,omitempty"`

    // Out95 出口带宽95值。
    Out95 *int64 `json:"out95,omitempty"`

    // OutAvg 出口带宽平均值。
    OutAvg *int64 `json:"outAvg,omitempty"`

    // OutMax 出口带宽最大值。
    OutMax *int64 `json:"outMax,omitempty"`

    // OutMin 出口带宽最小值。
    OutMin *int64 `json:"outMin,omitempty"`

    // Unit 带宽值单位。例如：bps。
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

// InquiryModifyIPTransitPriceRequest 
type InquiryModifyIPTransitPriceRequest struct {
    *common.BaseRequest

    // IptId IP Transit 实例 ID。
    IptId *string `json:"iptId,omitempty"`

    // Type 变配类型。
    // 支持 BANDWIDTH、ADD_CIDR_BLOCK、DEL_CIDR_BLOCK、EXPAND_CIDR_BLOCK、SHRINK_CIDR_BLOCK，BFD/BGP/HA 操作无费用，不允许传入。
    Type *string `json:"type,omitempty"`

    // Bandwidth 目标带宽（Mbps）。
    // type=BANDWIDTH 时必填。
    // 95 计费下必须大于等于 `commitBandwidth`。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // CommitBandwidth 保底带宽（Mbps）。
    // type=BANDWIDTH 时有效，不填则与 `bandwidth` 相同；95 计费（internetType=ByInstanceBandwidth95）下必填，不能用 `bandwidth` 代替。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // PublicIPv4BlockSize IPv4 CIDR 掩码长度（24–32）。
    // type=ADD_CIDR_BLOCK、EXPAND_CIDR_BLOCK、SHRINK_CIDR_BLOCK 时必填。
    PublicIPv4BlockSize *int `json:"publicIPv4BlockSize,omitempty"`

    // IpUuid 目标 IP 块 UUID。
    // type=DEL_CIDR_BLOCK、EXPAND_CIDR_BLOCK、SHRINK_CIDR_BLOCK 时必填。
    IpUuid *string `json:"ipUuid,omitempty"`

    // IpNetworkType IP 网络类型。
    // type=ADD_CIDR_BLOCK 时有效，默认 BGP_IP。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

}

type InquiryModifyIPTransitPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryModifyIPTransitPriceResponseParams `json:"response,omitempty"`

}

// InquiryModifyIPTransitPriceResponseParams 
type InquiryModifyIPTransitPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // PrivateConnectPrice 专线（VLL）价格。
    // BANDWIDTH 类型时可能有值，Router RIPT 为空。
    PrivateConnectPrice *PriceItem `json:"privateConnectPrice,omitempty"`

    // PrivateConnectBandwidth 专线带宽（Mbps）。
    // BANDWIDTH 类型时有值。
    PrivateConnectBandwidth *int `json:"privateConnectBandwidth,omitempty"`

    // IptPrice RIPT 带宽价格。
    IptPrice *PriceItem `json:"iptPrice,omitempty"`

    // IptIpPrices 公网 CIDR 块价格列表。
    // ADD_CIDR_BLOCK、EXPAND_CIDR_BLOCK 类型时有值。
    IptIpPrices []*IPTransitIpPriceItem `json:"iptIpPrices,omitempty"`

    // PublicInterconnectIpPrice 公网互联块价格。
    // 启用公网互联时填充，否则为空。
    PublicInterconnectIpPrice *IPTransitIpPriceItem `json:"publicInterconnectIpPrice,omitempty"`

}

// IPTransitIpPriceItem CIDR 块价格信息。
type IPTransitIpPriceItem struct {

    // Price 价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Netmask CIDR 掩码长度。
    Netmask *int `json:"netmask,omitempty"`

    // Amount 数量。
    Amount *int `json:"amount,omitempty"`

    // IpNetworkType IP 网络类型（BGP_IP / LOCAL_IP）。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

}

// ModifyIPTransitConfigRequest 
type ModifyIPTransitConfigRequest struct {
    *common.BaseRequest

    // IptId IP Transit 实例 ID。
    IptId *string `json:"iptId,omitempty"`

    // Type 变配操作类型。
    Type *string `json:"type,omitempty"`

    // Bandwidth 目标带宽（Mbps）。
    // type=BANDWIDTH 时必填。
    // 95 计费下必须大于等于生效后的 `commitBandwidth`。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // CommitBandwidth 保底带宽（Mbps）。
    // type=BANDWIDTH 时有效，不填则与 `bandwidth` 相同。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // Bfd BFD 配置。
    // type=BFD 时填写；传 null 表示关闭 BFD。
    // 高可用 IP Transit 不允许关闭 BFD。
    Bfd *BFDConfig `json:"bfd,omitempty"`

    // Bgp BGP 配置参数。
    // type=BGP_ROUTE_TYPE、BGP_ASN_AS_SET、BGP_PASSWORD 时必填，并填写对应子字段。
    Bgp *BgpConfigParam `json:"bgp,omitempty"`

    // PublicIPv4BlockSize IPv4 CIDR 掩码长度（24–32）。
    // type=ADD_CIDR_BLOCK、EXPAND_CIDR_BLOCK、SHRINK_CIDR_BLOCK 时必填。
    PublicIPv4BlockSize *int `json:"publicIPv4BlockSize,omitempty"`

    // IpUuid 目标 IP 块 UUID。
    // type=DEL_CIDR_BLOCK、EXPAND_CIDR_BLOCK、SHRINK_CIDR_BLOCK 时必填。
    IpUuid *string `json:"ipUuid,omitempty"`

    // IpNetworkType IP 网络类型。
    // type=ADD_CIDR_BLOCK 时有效，默认 BGP_IP。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

}

// BgpConfigParam BGP 变配参数。
type BgpConfigParam struct {

    // RouteType BGP inbound 路由类型。
    // type=BGP_ROUTE_TYPE 时必填。
    RouteType *string `json:"routeType,omitempty"`

    // AsnList ASN 列表。
    // type=BGP_ASN_AS_SET 时与 `asSetList` 二选一。
    // `asn` 创建后不支持修改。
    AsnList []int64 `json:"asnList,omitempty"`

    // AsSetList AS-SET 列表。
    // type=BGP_ASN_AS_SET 时与 `asnList` 二选一。
    // `asn` 创建后不支持修改。
    AsSetList []string `json:"asSetList,omitempty"`

    // Password BGP MD5 密码（长度 8–64）。
    // type=BGP_PASSWORD 时必填。
    Password *string `json:"password,omitempty"`

}

type ModifyIPTransitConfigResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

