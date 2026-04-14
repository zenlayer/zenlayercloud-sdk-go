package ipt

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


type DescribeIPTransitDatacentersRequest struct {
    *common.BaseRequest

    // PeerPortId 指定端口ID。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // PeerDcId 指定数据中心ID。
    PeerDcId *string `json:"peerDcId,omitempty"`

}

type DescribeIPTransitDatacentersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeIPTransitDatacentersResponseParams `json:"response,omitempty"`

}

type DescribeIPTransitDatacentersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // SupportSet 结果集。
    SupportSet []*IPTransitDatacenter `json:"supportSet,omitempty"`

}

// IPTransitDatacenter 描述支持IP Transit的数据中心信息，包括数据中心的基本信息以及支持的路由类型等。
type IPTransitDatacenter struct {

    // DataCenter 数据中心信息。
    DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

    // AvailableRoutingTypes 支持的路由类型配置。
    AvailableRoutingTypes []*RemoteIptAvailableRoutingType `json:"availableRoutingTypes,omitempty"`

}

// DatacenterInfo 描述数据中心的信息。
type DatacenterInfo struct {

    // DcId 数据中心ID。
    DcId *string `json:"dcId,omitempty"`

    // DcName 数据中心名称。
    DcName *string `json:"dcName,omitempty"`

    // DcAddress 数据中心地址。
    DcAddress *string `json:"dcAddress,omitempty"`

    // CityName 城市名称。
    CityName *string `json:"cityName,omitempty"`

    // CountryName 国家名称。
    CountryName *string `json:"countryName,omitempty"`

    // AreaName 地区名称。
    AreaName *string `json:"areaName,omitempty"`

    // Latitude 数据中心所在地理位置的维度。
    Latitude *float64 `json:"latitude,omitempty"`

    // Longitude 数据中心所在地理位置的经度。
    Longitude *float64 `json:"longitude,omitempty"`

}

// RemoteIptAvailableRoutingType 可用的路由类型信息。
type RemoteIptAvailableRoutingType struct {

    // RoutingType 路由类型。
    RoutingType *string `json:"routingType,omitempty"`

    // DeliveryType 交付模式。
    DeliveryType *string `json:"deliveryType,omitempty"`

}

type InquiryCreateIPTransitPriceRequest struct {
    *common.BaseRequest

    // PeerPortId 端口的ID。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // IptDcId IP Transit 目的地数据中心ID。
    // 如果不指定，则代表和端口位于同一个数据中心。
    IptDcId *string `json:"iptDcId,omitempty"`

    // InternetType IP Transit的带宽计费方式。
    InternetType *string `json:"internetType,omitempty"`

    // CommitBandwidth 保底带宽。
    // 单位Mbps。
    // 有且仅当internetType=ByInstanceBandwidth95时该字段必传。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // Bandwidth 带宽限速。
    // 单位Mbps。
    // 最小值不能低于10Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // RoutingType 路由类型。
    RoutingType *string `json:"routingType,omitempty"`

    // PublicIPv4BlockSize 公网IPv4地址。
    // 网段范围：24～30
    // 有且仅当路由类型是Static 或 Gateway时必须指定。
    // 目前只允许指定一个公网CIDR。
    PublicIPv4BlockSize []int `json:"publicIPv4BlockSize,omitempty"`

    // BgpRouteType BGP入站路由类型。
    BgpRouteType *string `json:"bgpRouteType,omitempty"`

}

type InquiryCreateIPTransitPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryCreateIPTransitPriceResponseParams `json:"response,omitempty"`

}

type InquiryCreateIPTransitPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // PrivateConnectPrice IP Transit 的 专线价格。
    // 如果IP Transit 和端口位于同一个数据中心，则取值为null。
    PrivateConnectPrice *PriceItem `json:"privateConnectPrice,omitempty"`

    // IptBandwidthPrice IP Transit 的公网带宽价格。
    IptBandwidthPrice *PriceItem `json:"iptBandwidthPrice,omitempty"`

    // PublicIpPrices IP Transit 的公网IP价格。
    PublicIpPrices []*IPPrice `json:"publicIpPrices,omitempty"`

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

// IPPrice IP价格信息。
type IPPrice struct {

    // Price IP CIDR 的价格信息。
    Price *PriceItem `json:"price,omitempty"`

    // Netmask IP CIDR 的网段。
    Netmask *int `json:"netmask,omitempty"`

    // Qty IP CIDR 的数量。
    Qty *int `json:"qty,omitempty"`

}

type CreateIPTransitRequest struct {
    *common.BaseRequest

    // IptName IP Transit的名称。
    // 长度不能超过255。
    IptName *string `json:"iptName,omitempty"`

    // IptDescription IP Transit的描述信息。
    // 长度不能超过255。
    IptDescription *string `json:"iptDescription,omitempty"`

    // PeerPortId 端口的ID。
    // 端口的连通性状态必须是ACTIVE。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // PeerPortVlan VLAN ID。
    // 范围为1000～4000。
    // 必须为未分配的vlan, 可以通过DescribePortUsableVlan来查询一个可以使用的vlan。
    PeerPortVlan *int `json:"peerPortVlan,omitempty"`

    // IptDcId IP Transit 目的地数据中心ID。
    // 如果不指定，则代表和端口位于同一个数据中心。
    IptDcId *string `json:"iptDcId,omitempty"`

    // InternetType IP Transit的带宽计费方式。
    InternetType *string `json:"internetType,omitempty"`

    // CommitBandwidth 保底带宽。
    // 单位Mbps。
    // 有且仅当internetType=ByInstanceBandwidth95时该字段必传。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // Bandwidth 带宽限速。
    // 单位Mbps。
    // 最小值不能低于10Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // RoutingType 路由类型。
    RoutingType *string `json:"routingType,omitempty"`

    // PublicIPv4BlockSize 公网IPv4地址。
    // 网段范围：24～30。
    // 有且仅当路由类型是Static 或 Gateway时必须指定。
    // 目前只允许指定一个公网CIDR。
    PublicIPv4BlockSize []int `json:"publicIPv4BlockSize,omitempty"`

    // Bfd 启用 BFD配置。
    // 如果不传该字段，则默认不启用BFD。
    // 网关模式不支持配置BFD。
    Bfd *BFDConfig `json:"bfd,omitempty"`

    // ResourceGroupId 资源组的ID。
    // 如果不传，则会放到默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Bgp BGP相关配置。
    Bgp *RiptBgpConfig `json:"bgp,omitempty"`

    // Tags 创建IP Transit时关联的标签。
    // 注意：关联标签键不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

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

// RiptBgpConfig BGP相关配置。
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
    // 长度限制：1～128个字符。
    Key *string `json:"key,omitempty"`

    // Value 标签值。
    // 长度限制：1～128个字符。
    Value *string `json:"value,omitempty"`

}

type CreateIPTransitResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateIPTransitResponseParams `json:"response,omitempty"`

}

type CreateIPTransitResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 创建时产生的订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // IptId IP Transit 的 ID。
    IptId *string `json:"iptId,omitempty"`

}

type DescribeIPTransitsRequest struct {
    *common.BaseRequest

    // IptIds IP Transit ID列表。
    // 最大支持长度为100。
    IptIds []string `json:"iptIds,omitempty"`

    // IptName IPT 名称。
    // 支持模糊搜索。
    IptName *string `json:"iptName,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PeerPortId 端口ID。
    // 通过该字段可以筛选与指定端口有关的IP Transit。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // IptDcId 数据中心ID。
    // 具体取值可通过调用接口DescribeDataCenters来获得最新的数据中心列表。
    IptDcId *string `json:"iptDcId,omitempty"`

    // PageSize 返回的分页大小。
    // 默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数。
    // 默认为1。
    PageNum *int `json:"pageNum,omitempty"`

    // TagKeys 根据标签进行搜索。
    // 最长不得超过20个标签。
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

type DescribeIPTransitsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet IP Transit列表结果集。
    DataSet []*IPTransit `json:"dataSet,omitempty"`

}

// IPTransit 描述一个IP Transit资源相关的信息。包括关联的端口，所属资源组等信息。
type IPTransit struct {

    // IptId IP Transit ID。
    IptId *string `json:"iptId,omitempty"`

    // IptName IP Transit 名称。
    IptName *string `json:"iptName,omitempty"`

    // IptDescription IP Transit 描述信息。
    IptDescription *string `json:"iptDescription,omitempty"`

    // DataCenter IP Transit 对应的数据中心。
    DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

    // PeerPortId 连接的端口ID。
    PeerPortId *string `json:"peerPortId,omitempty"`

    // PeerPortName 连接的端口名称。
    PeerPortName *string `json:"peerPortName,omitempty"`

    // PeerDataCenter 端口侧所属数据中心。
    PeerDataCenter *DatacenterInfo `json:"peerDataCenter,omitempty"`

    // DeliveryType 交付类型。
    DeliveryType *string `json:"deliveryType,omitempty"`

    // ResourceGroupId IP Transit 所属资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName IP Transit 所属资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // CreateTime 创建时间。
    // 格式为：YYYY-MM-ddTHH:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // RoutingType 路由类型。
    RoutingType *string `json:"routingType,omitempty"`

    // InternetType 带宽计费方式。
    InternetType *string `json:"internetType,omitempty"`

    // Bandwidth 带宽限速。
    // 单位Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // CommitBandwidth 承诺保底带宽。
    // 仅当带宽计费方式为ByInstanceBandwidth95 可取到值。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

    // Bfd BFD 配置。
    Bfd *BFDConfig `json:"bfd,omitempty"`

    // Bgp BGP相关配置。
    Bgp *RiptBgpConfig `json:"bgp,omitempty"`

    // Interconnect 互联IP地址信息。
    Interconnect *Interconnect `json:"interconnect,omitempty"`

    // PrivateConnectId IP Transit 互联的专线ID。
    PrivateConnectId *string `json:"privateConnectId,omitempty"`

    // PrivateConnectName IP Transit 互联的专线名称。
    PrivateConnectName *string `json:"privateConnectName,omitempty"`

    // PublicIpv4Addresses 公网IP信息。
    PublicIpv4Addresses []*IPAddress `json:"publicIpv4Addresses,omitempty"`

    // IptStatus IP Transit 的业务状态。
    IptStatus *string `json:"iptStatus,omitempty"`

    // Tags 该IP Transit关联的标签。
    Tags *Tags `json:"tags,omitempty"`

}

// Interconnect 描述IP Transit 互联IP地址的信息。
type Interconnect struct {

    // VendorIpv4Address 运营商侧的IPv4地址。
    VendorIpv4Address *string `json:"vendorIpv4Address,omitempty"`

    // CustomerIpv4Address 用户侧的IPv4地址。
    CustomerIpv4Address *string `json:"customerIpv4Address,omitempty"`

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

type ModifyIPTransitBandwidthRequest struct {
    *common.BaseRequest

    // IptId IP Transit ID。
    IptId *string `json:"iptId,omitempty"`

    // Bandwidth 需要修改的带宽限速。
    // 范围：1~1000。
    // 单位：Mbps。
    Bandwidth *int `json:"bandwidth,omitempty"`

    // CommitBandwidth 保底带宽。
    // 当IP Transit的带宽计费方式为95计费时该参数有效，如果不设置，则不会修改保底。
    CommitBandwidth *int `json:"commitBandwidth,omitempty"`

}

type ModifyIPTransitBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyIPTransitsAttributeRequest struct {
    *common.BaseRequest

    // IptIds IP Transit ID 列表。
    // 数量不得超过100。
    IptIds []string `json:"iptIds,omitempty"`

    // IptName IP Transit名称。
    // 不得超过255个字符。
    // 名称和描述信息至少需要有一项指定。
    IptName *string `json:"iptName,omitempty"`

    // IptDescription IP Transit 描述信息。
    // 名称和描述信息至少需要有一项指定。
    IptDescription *string `json:"iptDescription,omitempty"`

}

type ModifyIPTransitsAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

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

type DescribeIPTransitTrafficRequest struct {
    *common.BaseRequest

    // IptId IP Transit ID。
    IptId *string `json:"iptId,omitempty"`

    // StartTime 查询开始时间。
    // 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。
    // 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-ddTHH:mm:ssZ。
    // 默认为当前时间。
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

// TrafficData 带宽数据。
type TrafficData struct {

    // InternetRX 入口带宽。单位：bps。
    InternetRX *int64 `json:"internetRX,omitempty"`

    // InternetTX 出口带宽。单位：bps。
    InternetTX *int64 `json:"internetTX,omitempty"`

    // Time 数据时间。
    // 格式为：YYYY-MM-DDThh:mm:ssZ。
    Time *string `json:"time,omitempty"`

}

