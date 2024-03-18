package sdn

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeDatacentersRequest struct {
	*common.BaseRequest

	// 数据中心四字码。
	DcIds []string `json:"dcIds,omitempty"`

	// 筛选是否支持新建数据中心端口的dc
	IsPortAvailable bool `json:"isPortAvailable,omitempty"`
}

type DescribeDatacentersResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeDatacentersResponseParams `json:"response"`
}

type DescribeDatacentersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// 端口id
	DcSet []*DatacenterInfo `json:"dcSet,omitempty"`
}

type CreatePortRequest struct {
	*common.BaseRequest

	// 数据中心四字码。
	DcId string `json:"dcId,omitempty"`

	// 端口名称。
	PortName string `json:"portName,omitempty"`

	// 端口备注信息。
	PortRemarks string `json:"portRemarks,omitempty"`

	// 端口规格。
	PortType string `json:"portType,omitempty"`

	// 贵司商业实体名称。
	BusinessEntityName string `json:"businessEntityName,omitempty"`
}

type CreatePortResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreatePortResponseParams `json:"response"`
}

type CreatePortResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// 端口id
	PortId string `json:"portId,omitempty"`
}

type DescribeDataCenterPortPriceRequest struct {
	*common.BaseRequest

	// 数据中心四字码。
	DcId string `json:"dcId,omitempty"`
}

type DescribeDataCenterPortPriceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeDataCenterPortPriceResponseParams `json:"response"`
}

type DescribeDataCenterPortPriceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// 结果集。
	PortPriceSet []*PortPrice `json:"portPriceSet,omitempty"`
}

type PortPrice struct {
	// 数据中心四字码。
	DcId string `json:"dcId,omitempty"`

	// 端口 的规格。
	PortType string `json:"portType,omitempty"`

	// 端口 天价格。
	PortDailyPrice float64 `json:"portDailyPrice,omitempty"`

	// 端口 月价格。
	PortMonthlyPrice float64 `json:"portMonthlyPrice,omitempty"`
}

type DescribePortsRequest struct {
	*common.BaseRequest

	// 数据中心四字码。
	PortIds []string `json:"portIds,omitempty"`

	// 数据中心四字码。
	DcId string `json:"dcId,omitempty"`

	// 端口名称。
	PortName string `json:"portName,omitempty"`

	// 端口备注信息。
	PortRemarks string `json:"portRemarks,omitempty"`

	// 城市名称。
	CityName string `json:"cityName,omitempty"`

	// 返回的分页数。
	// 默认为1。
	PageNum int `json:"pageNum,omitempty"`

	// 返回的分页大小。
	// 默认为20，最大为1000。
	PageSize int `json:"pageSize,omitempty"`
}

type DescribePortsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribePortsResponseParams `json:"response,omitempty"`
}

type DescribePortsResponseParams struct {
	RequestId  string      `json:"requestId,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	DataSet    []*PortInfo `json:"dataSet,omitempty"`
}

type PortInfo struct {
	// 端口 的Id。
	PortId string `json:"portId,omitempty"`

	// 端口 的名称。
	PortName string `json:"portName,omitempty"`

	// 端口 的备注。
	PortRemarks string `json:"portRemarks,omitempty"`

	// 端口 规格。
	PortType string `json:"portType,omitempty"`

	// 端口 连通性的状态。
	ConnectionStatus string `json:"connectionStatus,omitempty"`

	// 端口 业务状态。
	PortStatus string `json:"portStatus,omitempty"`

	// 数据中心 的Id。
	DcId string `json:"dcId,omitempty"`

	// 数据中心 的名称。
	DcName string `json:"dcName,omitempty"`

	// 城市 的名称。
	CityName string `json:"cityName,omitempty"`

	// 区域 的名称。
	AreaName string `json:"areaName,omitempty"`

	// LOA 的状态。
	LoaStatus string `json:"loaStatus,omitempty"`

	// LOA 的下载地址。
	LoaDownloadUrl string `json:"loaDownloadUrl,omitempty"`

	// 商业实体的名称。用于LOA。
	BusinessEntityName string `json:"businessEntityName,omitempty"`

	// 创建时间。
	// 格式为：YYYY-MM-DDThh:mm:ssZ。
	CreatedTime string `json:"createdTime,omitempty"`

	// 到期时间。
	// 格式为：YYYY-MM-DDThh:mm:ssZ。
	ExpiredTime string `json:"expiredTime,omitempty"`

	// 端口计费类型。
	// PREPAID：预付费，即包年包月。 POSTPAID：后付费。
	PortChargeType string `json:"portChargeType,omitempty"`

	// 购买端口的时长。
	// 单位：月。
	// 后付费实例该字段为null。
	Period int `json:"period,omitempty"`
}

type DescribePortTrafficRequest struct {
	*common.BaseRequest

	// 端口id。
	PortId string `json:"portId,omitempty"`

	// 查询开始时间。
	// 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	StartTime string `json:"startTime,omitempty"`

	// 查询结束时间。
	// 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime string `json:"endTime,omitempty"`
}

type DescribePortTrafficResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribePortTrafficResponseParams `json:"response,omitempty"`
}

type DescribePortTrafficResponseParams struct {
	RequestId string `json:"requestId,omitempty"`
	// 带宽数据列表。
	DataList []*TrafficData `json:"dataList,omitempty"`

	// 入口带宽95值。
	In95 float64 `json:"in95,omitempty"`

	// 入口带宽平均值。
	InAvg int `json:"inAvg,omitempty"`

	// 入口带宽最大值。
	InMax int `json:"inMax,omitempty"`

	// 入口带宽最小值。
	InMin int `json:"inMin,omitempty"`

	// 入口带宽总流量。
	InTotal int `json:"inTotal,omitempty"`

	// 出口带宽95值。
	Out95 int `json:"out95,omitempty"`

	// 出口带宽平均值。
	OutAvg int `json:"outAvg,omitempty"`

	// 出口带宽最大值。
	OutMax int `json:"outMax,omitempty"`

	// 出口带宽最小值。
	OutMin int `json:"outMin,omitempty"`

	// 带宽值单位。例如：bps。
	Unit string `json:"unit,omitempty"`
}

type TrafficData struct {
	// 入口带宽。单位：bps。
	InternetRX int `json:"internetRX,omitempty"`

	// 出口带宽。单位：bps。
	InternetTX int `json:"internetTX,omitempty"`

	// 数据时间。
	// 格式为：YYYY-MM-DDThh:mm:ssZ。
	Time string `json:"time,omitempty"`
}

type DescribePortUsableVlanRequest struct {
	*common.BaseRequest

	// 数据中心四字码。
	DcId string `json:"dcId,omitempty"`

	// 端口id。
	PortId string `json:"portId,omitempty"`
}

type DescribePortUsableVlanResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribePortUsableVlanResponseParams `json:"response,omitempty"`
}

type DescribePortUsableVlanResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// VLAN 范围起始值。
	Start int `json:"start,omitempty"`

	// VLAN 范围结束值。
	End int `json:"end,omitempty"`

	// 占用中的vlan列表。
	InuseVlanList []int `json:"inuseVlanList,omitempty"`
}

type DestroyPortRequest struct {
	*common.BaseRequest

	// 端口id。
	PortId string `json:"portId,omitempty"`
}

type DestroyPortResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type ModifyPortAttributeRequest struct {
	*common.BaseRequest

	// 端口id。
	PortId string `json:"portId,omitempty"`

	// 端口名称。
	PortName string `json:"portName,omitempty"`

	// 端口备注信息。
	PortRemarks string `json:"portRemarks,omitempty"`

	// 贵司商业实体名称。
	BusinessEntityName string `json:"businessEntityName,omitempty"`
}

type ModifyPortAttributeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type TerminatePortRequest struct {
	*common.BaseRequest

	// 端口id。
	PortId string `json:"portId,omitempty"`
}

type TerminatePortResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type RenewPortRequest struct {
	*common.BaseRequest

	// 端口id。
	PortId string `json:"portId,omitempty"`
}

type RenewPortResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DescribePrivateConnectsRequest struct {
	*common.BaseRequest

	// 二层网络专线ID列表。
	PrivateConnectIds []string `json:"privateConnectIds,omitempty"`

	// 二层网络专线的名称。
	PrivateConnectName string `json:"privateConnectName,omitempty"`

	// 连通性状态。
	ConnectivityStatus string `json:"connectivityStatus,omitempty"`

	// 二层网络专线的状态。
	PrivateConnectStatus string `json:"privateConnectStatus,omitempty"`

	// 二层网络专线连接点的类型
	EndpointTypes []string `json:"endpointTypes,omitempty"`

	// 资源组的ID。
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// 返回的分页数。
	// 默认为1。
	PageNum int `json:"pageNum,omitempty"`

	// 返回的分页大小。
	// 默认为20，最大为1000。
	PageSize int `json:"pageSize,omitempty"`
}

type DescribePrivateConnectsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribePrivateConnectsResponseParams `json:"response,omitempty"`
}

type DescribePrivateConnectsResponseParams struct {
	RequestId  string            `json:"requestId,omitempty"`
	TotalCount int               `json:"totalCount,omitempty"`
	DataSet    []*PrivateConnect `json:"dataSet,omitempty"`
}

type PrivateConnect struct {
	// 二层专线网络ID。
	PrivateConnectId string `json:"privateConnectId,omitempty"`

	// 二层专线网络名称
	PrivateConnectName string `json:"privateConnectName,omitempty"`

	// 连通性状态。
	ConnectivityStatus string `json:"connectivityStatus,omitempty"`

	// 连接点A的基本信息。
	EndpointA PrivateConnectEndpoint `json:"endpointA,omitempty"`

	// 连接点Z的基本信息。
	EndpointZ PrivateConnectEndpoint `json:"endpointZ,omitempty"`

	// 专线的业务状态。
	PrivateConnectStatus string `json:"privateConnectStatus,omitempty"`

	// 二层专线网络的最大带宽限速。
	BandwidthMbps int64 `json:"bandwidthMbps,omitempty"`

	// 所属的资源组ID。
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// 所属的资源组名称。
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// 创建时间。
	CreateTime string `json:"createTime,omitempty"`

	// 过期时间。
	// 预付费资源可用。
	ExpiredTime string `json:"expiredTime,omitempty"`

	// 回收时间。
	RecycledTime string `json:"recycledTime,omitempty"`
}

type PrivateConnectEndpoint struct {
	// 连接点的ID。
	EndpointId string `json:"endpointId,omitempty"`

	// 连接点的名称。
	EndpointName string `json:"endpointName,omitempty"`

	//公有云区域ID。
	//有且为Google, Tencent, AWS 类型的接入点可取到值。
	CloudRegionId string `json:"cloudRegionId,omitempty"`

	//云接入的账号。
	//有且为Google, Tencent, AWS 类型的接入点可取到值。
	CloudAccountId string `json:"cloudAccountId,omitempty"`

	// 连接点的类型。
	EndpointType string `json:"endpointType,omitempty"`

	// 连接点所属的数据中心信息。
	DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

	// 连接点的VLAN ID。
	// 范围 1～4096。
	VlanId int `json:"vlanId,omitempty"`

	// 连接点的连通性状态。
	ConnectivityStatus string `json:"connectivityStatus,omitempty"`
}

type DatacenterInfo struct {
	// 数据中心四字码
	DcId string `json:"dcId,omitempty"`

	// 数据中心名称。
	DcName string `json:"dcName,omitempty"`

	// 数据中心地址。
	DcAddress string `json:"dcAddress,omitempty"`

	// 城市名称。
	CityName string `json:"cityName,omitempty"`

	// 国家名称。
	CountryName string `json:"countryName,omitempty"`

	// 地区名称。
	AreaName string `json:"areaName,omitempty"`

	// 数据中心所在地理位置的经度。
	Longitude float64 `json:"longitude,omitempty"`

	//数据中心所在地理位置的纬度
	Latitude float64 `json:"latitude,omitempty"`
}

type DescribePrivateConnectAvailablePortsRequest struct {
	*common.BaseRequest

	// 端口ID列表
	PortIds []string `json:"portIds,omitempty"`
	// 数据中心ID。
	DcId string `json:"dcId,omitempty"`

	// 返回的分页数。
	// 默认为1。
	PageNum int `json:"pageNum,omitempty"`

	// 返回的分页大小。
	// 默认为20，最大为100。
	PageSize int `json:"pageSize,omitempty"`
}

type DescribePrivateConnectAvailablePortsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribePrivateConnectAvailablePortsResponseParams `json:"response,omitempty"`
}

type DescribePrivateConnectAvailablePortsResponseParams struct {
	RequestId  string      `json:"requestId,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	DataSet    []*PortInfo `json:"dataSet,omitempty"`
}

type CreatePrivateConnectRequest struct {
	*common.BaseRequest

	// 二层网络专线的名称。
	PrivateConnectName string `json:"privateConnectName,omitempty"`

	// 创建二层网络其中一端的连接点（A）
	EndpointA CreateEndpointParam `json:"endpointA,omitempty"`

	// 创建二层网络另一端的连接点（Z）
	EndpointZ CreateEndpointParam `json:"endpointZ,omitempty"`

	// 二层网络专线最大带宽。
	// 范围是1～500。
	// 默认值是1 ，单位：Mbps。
	BandwidthMbps int `json:"bandwidthMbps,omitempty"`

	// 资源组的ID。
	// 如果不传，则会放到默认资源组。
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
}

type CreateEndpointParam struct {
	// 二层专线网络ID。
	SubnetId string `json:"subnetId,omitempty"`

	// 数据中心端口的ID。
	PortId string `json:"portId,omitempty"`

	// VLAN ID。
	VlanId int `json:"vlanId,omitempty"`

	//云平台账号。
	//Google 云此处为Pairing Key
	CloudAccountId string `json:"cloudAccountId,omitempty"`

	//公有云区域ID。
	CloudRegionId string `json:"cloudRegionId,omitempty"`

	//连接云接入点的数据中心ID。
	DcId string `json:"dcId,omitempty"`

	//云连接类型可选值：
	//AWS / TENCENT / GOOGLE
	CloudType string `json:"cloudType,omitempty"`
}

type CreatePrivateConnectResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreatePrivateConnectResponseParams `json:"response,omitempty"`
}

type CreatePrivateConnectResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// 二层网络专线ID。
	PrivateConnectId string `json:"privateConnectId,omitempty"`
}

type ModifyPrivateConnectsAttributeRequest struct {
	*common.BaseRequest

	// 二层网络专线的ID列表。
	PrivateConnectIds []string `json:"privateConnectIds,omitempty"`

	// 二层网络专线的名称。
	PrivateConnectName string `json:"privateConnectName,omitempty"`
}

type ModifyPrivateConnectsAttributeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`
}

type InquiryCreatePrivateConnectPriceRequest struct {
	*common.BaseRequest

	// 创建二层网络其中一端的连接点（A）
	EndpointA CreateEndpointParam `json:"endpointA,omitempty"`

	// 创建二层网络另一端的连接点（Z）
	EndpointZ CreateEndpointParam `json:"endpointZ,omitempty"`

	// 二层网络专线的最大带宽限制。
	BandwidthMbps int `json:"bandwidthMbps,omitempty"`
}

type InquiryCreatePrivateConnectPriceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *InquiryCreatePrivateConnectPriceResponseParams `json:"response,omitempty"`
}

type InquiryCreatePrivateConnectPriceResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// 二层网络专线带宽价格。
	PrivateConnectPrice Price `json:"privateConnectPrice,omitempty"`
}

type Price struct {
	// 预付费的原价。
	// 预付费模式使用，后付费该值为 null。
	OriginalPrice *float64 `json:"originalPrice,omitempty"`

	// 预付费的折扣价。
	// 预付费模式使用，后付费该值为 null。
	DiscountPrice *float64 `json:"discountPrice,omitempty"`

	// 折扣大小。
	Discount *float64 `json:"discount,omitempty"`

	// 后付费的单元原始价格。
	// 后付费模式使用，如果价格为阶梯价格，该项为null。
	UnitPrice *float64 `json:"unitPrice,omitempty"`

	// 后付费的单元折后价格。
	// 后付费模式使用，如果价格为阶梯价格，该项为null。
	DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`

	// 后付费计价单元。
	// 后付费模式使用，可取值范围：
	// HOUR: 表示计价单元是按每小时来计算。
	// DAY: 表示计价单元是按天来计算。
	// MONTH: 表示计价单元是按月来计算，95计费则是这种。
	ChargeUnit string `json:"chargeUnit,omitempty"`

	// 后付费阶梯价格。
	// 后付费模式使用，如果非阶梯价格，该项为null。
	StepPrices []*StepPrice `json:"stepPrices,omitempty"`
}

type StepPrice struct {
	// 阶梯用量的开始。
	StepStart *float64 `json:"stepStart,omitempty"`

	// 阶梯用量的结束。
	StepEnd *float64 `json:"stepEnd,omitempty"`

	// 当前阶梯的单元原始价格。
	// 后付费模式使用。
	UnitPrice *float64 `json:"unitPrice,omitempty"`

	// 当前阶梯的单元折后价格。
	// 后付费模式使用。
	DiscountUnitPrice *float64 `json:"discountUnitPrice,omitempty"`
}

type DeletePrivateConnectRequest struct {
	*common.BaseRequest

	// 二层网络专线ID。
	PrivateConnectId string `json:"privateConnectId,omitempty"`
}

type DeletePrivateConnectResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DestroyPrivateConnectRequest struct {
	*common.BaseRequest

	// 二层网络专线ID。
	PrivateConnectId string `json:"privateConnectId,omitempty"`
}

type DestroyPrivateConnectResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type RenewPrivateConnectRequest struct {
	*common.BaseRequest

	// 二层网络专线ID。
	PrivateConnectId string `json:"privateConnectId,omitempty"`
}

type RenewPrivateConnectResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DescribePrivateConnectTrafficRequest struct {
	*common.BaseRequest

	// 二层网络专线ID。
	PrivateConnectId string `json:"privateConnectId,omitempty"`

	// 查询开始时间。
	// 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	StartTime string `json:"startTime,omitempty"`

	// 查询结束时间。
	// 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 默认为当前时间。
	EndTime string `json:"endTime,omitempty"`
}

type DescribePrivateConnectTrafficResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribePrivateConnectTrafficResponseParams `json:"response,omitempty"`
}

type DescribePrivateConnectTrafficResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// 带宽数据列表。
	DataList []*TrafficData `json:"dataList,omitempty"`

	// 入口带宽95值。
	In95 *float64 `json:"in95,omitempty"`

	// 入口带宽平均值。
	InAvg int `json:"inAvg,omitempty"`

	// 入口带宽最大值。
	InMax int `json:"inMax,omitempty"`

	// 入口带宽最小值。
	InMin int `json:"inMin,omitempty"`

	// 入口带宽总流量。
	InTotal int `json:"inTotal,omitempty"`

	// 出口带宽95值。
	Out95 int `json:"out95,omitempty"`

	// 出口带宽平均值。
	OutAvg int `json:"outAvg,omitempty"`

	// 出口带宽最大值。
	OutMax int `json:"outMax,omitempty"`

	// 出口带宽最小值。
	OutMin int `json:"outMin,omitempty"`

	// 带宽值单位。例如：bps。
	Unit string `json:"unit,omitempty"`
}

type ModifyPrivateConnectBandwidthRequest struct {
	*common.BaseRequest

	// 二层网络专线ID。
	PrivateConnectId string `json:"privateConnectId,omitempty"`

	// 需要修改的带宽限速。
	BandwidthMbps int `json:"bandwidthMbps,omitempty"`
}

type ModifyPrivateConnectBandwidthResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DescribeCloudRoutersRequest struct {
	*common.BaseRequest

	// 三层网络的ID。
	CloudRouterIds []string `json:"cloudRouterIds,omitempty"`

	// 三层网络的名称。
	CloudRouterName string `json:"cloudRouterName,omitempty"`

	// 三层网络的状态。
	CloudRouterStatus string `json:"cloudRouterStatus,omitempty"`

	// 边缘连接点的ID。
	EgdePointId string `json:"egdePointId,omitempty"`

	// 资源组的ID。
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// 返回的分页数。
	PageNum int `json:"pageNum,omitempty"`

	// 返回的分页大小。
	PageSize int `json:"pageSize,omitempty"`
}

type DescribeCloudRoutersResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeCloudRoutersResponseParams `json:"response,omitempty"`
}

type DescribeCloudRoutersResponseParams struct {
	RequestId  string         `json:"requestId,omitempty"`
	TotalCount int            `json:"totalCount,omitempty"`
	DataSet    []*CloudRouter `json:"dataSet,omitempty"`
}

type CloudRouter struct {
	// 三层网络的ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`

	// 三层网络的名称。
	CloudRouterName string `json:"cloudRouterName,omitempty"`

	// 三层网络的描述信息。
	CloudRouterDescription string `json:"cloudRouterDescription,omitempty"`

	// 三层网络中边缘连接点的信息。
	EdgePoints []*CloudRouterEdgePoint `json:"edgePoints,omitempty"`

	// 三层网络的计费类型。
	CloudRouterChargeType string `json:"cloudRouterChargeType,omitempty"`

	// 创建时间。
	CreateTime string `json:"createTime,omitempty"`

	// 到期时间。
	ExpiredTime string `json:"expiredTime,omitempty"`

	// 销毁的时间。
	RecycledTime string `json:"recycledTime,omitempty"`

	// 购买三层网络的时长。
	Period int `json:"period,omitempty"`

	// 三层网络的状态。
	CloudRouterStatus string `json:"cloudRouterStatus,omitempty"`

	// 三层网络的连通性状态。
	ConnectivityStatus string `json:"connectivityStatus,omitempty"`

	// 三层网络所在资源组ID。
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	// 三层网络所在资源组名称
	ResourceGroupName string `json:"resourceGroupName,omitempty"`
}

type CloudRouterEdgePoint struct {
	// 边缘连接点的ID。
	EdgePointId string `json:"edgePointId,omitempty"`

	// 边缘连接点的名称。
	EdgePointName string `json:"edgePointName,omitempty"`

	// 边缘连接点的状态。
	ConnectivityStatus string `json:"connectivityStatus,omitempty"`

	// 边缘连接点所在的数据中心信息。
	DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

	// IP地址。
	IpAddress string `json:"ipAddress,omitempty"`

	// 边缘连接点的类型。
	EdgePointType string `json:"edgePointType,omitempty"`

	// VPC边缘连接点的ID.
	VpcId string `json:"vpcId,omitempty"`

	// 物理端口的ID.
	PortId string `json:"portId,omitempty"`

	// 云连接的区域。
	CloudRegionId string `json:"cloudRegionId,omitempty"`

	// 云连接的账号。
	CloudAccountId string `json:"CloudAccountId,omitempty"`

	// 边缘连接点配置的VLAN ID。
	VlanId int `json:"vlanId,omitempty"`

	// 带宽大小。
	BandwidthMbps int `json:"bandwidthMbps,omitempty"`

	// BGP连接配置信息。
	BgpConnection *BGPConnection `json:"bgpConnection,omitempty"`

	// 静态路由配置信息
	StaticRoutes []*IPRoute `json:"staticRoutes,omitempty"`

	// 创建时间。
	CreateTime string `json:"createTime,omitempty"`
}

type BGPConnection struct {
	// BGP 对等体的 IP 地址。
	PeerIpAddress string `json:"peerIpAddress,omitempty"`

	// 远程 BGP 对等体的 ASN。
	PeerAsn int64 `json:"peerAsn,omitempty"`

	// 本地 BGP 的 ASN。
	LocalAsn int64 `json:"localAsn,omitempty"`

	// 用于验证 BGP MD5 认证的对等体的共享密钥。
	Password string `json:"password,omitempty"`
}

type IPRoute struct {
	// 用于路由到下一跳的 IPv4 前缀。
	Prefix string `json:"prefix,omitempty"`

	// 下一跳IPv4地址。
	NextHop string `json:"nextHop,omitempty"`
}

type DescribeCloudRouterAvailableVpcsRequest struct {
	*common.BaseRequest

	// VPC ID。
	VpcId string `json:"vpcId,omitempty"`

	// 返回的分页数。
	PageNum int `json:"pageNum,omitempty"`

	// 返回的分页大小。
	PageSize int `json:"pageSize,omitempty"`
}

type DescribeCloudRouterAvailableVpcsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeCloudRouterAvailableVpcsResponseParams `json:"response,omitempty"`
}

type DescribeCloudRouterAvailableVpcsResponseParams struct {
	RequestId  string                     `json:"requestId,omitempty"`
	TotalCount int                        `json:"totalCount,omitempty"`
	DataSet    []*CloudRouterAvailableVpc `json:"dataSet,omitempty"`
}

type CloudRouterAvailableVpc struct {
	// VPC ID。
	VpcId string `json:"vpcId,omitempty"`

	// VPC 名称
	VpcName string `json:"vpcName,omitempty"`

	// vpc所在的数据中心。
	DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

	// VPC 的 CIDR。
	CidrBlock string `json:"cidrBlock,omitempty"`
}

type DescribeCloudRouterAvailablePortsRequest struct {
	*common.BaseRequest

	// 物理端口ID。
	PortIds []string `json:"portIds,omitempty"`

	// 数据中心ID。
	DcId string `json:"dcId,omitempty"`

	// 返回的分页数。
	PageNum int `json:"pageNum,omitempty"`

	// 返回的分页大小。
	PageSize int `json:"pageSize,omitempty"`
}

type DescribeCloudRouterAvailablePortsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeCloudRouterAvailablePortsResponseParams `json:"response,omitempty"`
}

type DescribeCloudRouterAvailablePortsResponseParams struct {
	RequestId  string      `json:"requestId,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	DataSet    []*PortInfo `json:"dataSet,omitempty"`
}

type CreateCloudRouterRequest struct {
	*common.BaseRequest

	// 三层网络的名称。
	CloudRouterName string `json:"cloudRouterName,omitempty"`

	// 三层网络的描述信息。
	CloudRouterDescription string `json:"cloudRouterDescription,omitempty"`

	// 创建三层网络中的边缘连接点信息。
	EdgePoints []*CreateCloudRouterEdgePoint `json:"edgePoints,omitempty"`

	// 资源组的ID。
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
}

type CreateCloudRouterEdgePoint struct {
	// 裸金属产品内VPC ID。
	VpcId string `json:"vpcId,omitempty"`

	// 接入的带宽大小。
	BandwidthMbps int `json:"bandwidthMbps,omitempty"`

	// 物理端口ID。
	PortId string `json:"portId,omitempty"`

	// VLAN ID。
	VlanId int `json:"vlanId,omitempty"`

	// IP地址信息。
	IpAddress string `json:"ipAddress,omitempty"`

	// BGP连接配置信息。
	BgpConnection *BGPConnection `json:"bgpConnection,omitempty"`

	// 静态路由配置信息。
	StaticRoutes []*IPRoute `json:"staticRoutes,omitempty"`
}

type CreateCloudRouterResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateCloudRouterResponseParams `json:"response,omitempty"`
}

type CreateCloudRouterResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// 三层网络ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`
}

type ModifyCloudRoutersAttributeRequest struct {
	*common.BaseRequest

	// 三层网络的ID列表。
	CloudRouterIds []string `json:"cloudRouterIds,omitempty"`

	// 三层网络的名称。
	CloudRouterName string `json:"cloudRouterName,omitempty"`

	// 三层网络的描述信息。
	CloudRouterDescription string `json:"cloudRouterDescription,omitempty"`
}

type ModifyCloudRoutersAttributeResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type AddCloudRouterEdgePointsRequest struct {
	*common.BaseRequest

	// 三层网络的ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`

	// 新的连接点信息。。
	EdgePoints []*CreateCloudRouterEdgePoint `json:"edgePoints,omitempty"`
}

type AddCloudRouterEdgePointsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *AddCloudRouterEdgePointsResponseParams `json:"response,omitempty"`
}

type AddCloudRouterEdgePointsResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// 新增的连接点ID。
	EdgePointIds []string `json:"edgePointIds,omitempty"`
}

type DeleteCloudRouterEdgePointRequest struct {
	*common.BaseRequest

	// 要移除的连接点ID。
	EdgePointId string `json:"edgePointId,omitempty"`

	// 连接点所在的三层网络ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`
}

type DeleteCloudRouterEdgePointResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DeleteCloudRouterRequest struct {
	*common.BaseRequest

	// 连接点所在的三层网络ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`
}

type DeleteCloudRouterResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DestroyCloudRouterRequest struct {
	*common.BaseRequest

	// 连接点所在的三层网络ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`
}

type DestroyCloudRouterResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type RenewCloudRouterRequest struct {
	*common.BaseRequest

	// 连接点所在的三层网络ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`
}

type RenewCloudRouterResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type ModifyCloudRouterEdgePointBandwidthRequest struct {
	*common.BaseRequest

	// 三层网络连接点的ID。
	EdgePointId string `json:"edgePointId,omitempty"`

	// 连接点所在的三层网络ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`

	// 需要修改的带宽限速。
	BandwidthMbps int `json:"bandwidthMbps,omitempty"`
}

type ModifyCloudRouterEdgePointBandwidthResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DescribeCloudRouterEdgePointTrafficRequest struct {
	*common.BaseRequest

	// 三层网络连接点的ID。
	EdgePointId string `json:"edgePointId,omitempty"`

	// 查询开始时间。
	StartTime string `json:"startTime,omitempty"`

	// 查询结束时间。
	EndTime string `json:"endTime,omitempty"`
}

type DescribeCloudRouterEdgePointTrafficResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeCloudRouterEdgePointTrafficResponseParams `json:"response,omitempty"`
}

type DescribeCloudRouterEdgePointTrafficResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// 带宽数据列表。
	DataList []*TrafficData `json:"dataList,omitempty"`

	// 入口带宽95值。
	In95 *float64 `json:"in95,omitempty"`

	// 入口带宽平均值。
	InAvg int `json:"inAvg,omitempty"`

	// 入口带宽最大值。
	InMax int `json:"inMax,omitempty"`

	// 入口带宽最小值。
	InMin int `json:"inMin,omitempty"`

	// 入口带宽总流量。
	InTotal int `json:"inTotal,omitempty"`

	// 出口带宽95值。
	Out95 int `json:"out95,omitempty"`

	// 出口带宽平均值。
	OutAvg int `json:"outAvg,omitempty"`

	// 出口带宽最大值。
	OutMax int `json:"outMax,omitempty"`

	// 出口带宽最小值。
	OutMin int `json:"outMin,omitempty"`

	// 带宽值单位。例如：bps。
	Unit string `json:"unit,omitempty"`
}

type ModifyCloudRouterEdgePointRequest struct {
	*common.BaseRequest

	// 三层网络连接点的ID。
	EdgePointId string `json:"edgePointId,omitempty"`

	// 连接点关联的三层网络ID。
	CloudRouterId string `json:"cloudRouterId,omitempty"`

	// BGP连接配置信息。
	BgpConnection *BGPConnection `json:"bgpConnection,omitempty"`

	// 静态路由配置信息。
	StaticRoutes []*IPRoute `json:"staticRoutes,omitempty"`

	// 需要修改的带宽限速。
	BandwidthMbps int `json:"bandwidthMbps,omitempty"`

	// IP地址信息。
	IpAddress string `json:"ipAddress,omitempty"`
}

type ModifyCloudRouterEdgePointResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`
}

type DescribeAWSRegionsRequest struct {
	*common.BaseRequest

	//当前AWS云接入点支持的产品。可用值：
	//PrivateConnect: 二层网络
	//CloudRouter: 三层网络
	Product string `json:"product,omitempty"`
}

type DescribeAWSRegionsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAWSRegionsResponseParams `json:"response,omitempty"`
}

type DescribeAWSRegionsResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// AWS的接入点相关的区域信息。
	CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`
}

type CloudRegion struct {

	// 公有云区域ID。
	CloudRegionId string `json:"cloudRegionId,omitempty"`

	//云接入点所在的数据中心信息。
	DataCenter *DatacenterInfo `json:"dataCenter,omitempty"`

	//当前云节点点支持的产品。
	Products []string `json:"products,omitempty"`
}

type DescribeAWSVlanUsageRequest struct {
	*common.BaseRequest

	DcId string `json:"dcId,omitempty"`
}

type DescribeAWSVlanUsageResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAWSVlanUsageResponseParams `json:"response,omitempty"`
}

type DescribeAWSVlanUsageResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	Start int `json:"start,omitempty"`

	End int `json:"end,omitempty"`

	UsedVlans []int `json:"usedVlans,omitempty"`
}

type DescribeTencentRegionsRequest struct {
	*common.BaseRequest

	//当前Tencent云接入点支持的产品。可用值：
	//PrivateConnect: 二层网络
	//CloudRouter: 三层网络
	Product string `json:"product,omitempty"`
}

type DescribeTencentRegionsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAWSRegionsResponseParams `json:"response,omitempty"`
}

type DescribeTencentRegionsResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// AWS的接入点相关的区域信息。
	CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`
}

type DescribeTencentVlanUsageRequest struct {
	*common.BaseRequest

	DcId string `json:"dcId,omitempty"`
}

type DescribeTencentVlanUsageResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeTencentVlanUsageResponseParams `json:"response,omitempty"`
}

type DescribeTencentVlanUsageResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	Start int `json:"start,omitempty"`

	End int `json:"end,omitempty"`

	UsedVlans []int `json:"usedVlans,omitempty"`
}

type DescribeGoogleRegionsRequest struct {
	*common.BaseRequest

	//Google 配对密钥。
	PairingKey string `json:"pairingKey,omitempty"`

	//当前Google云接入点支持的产品。可用值：
	//PrivateConnect: 二层网络
	//CloudRouter: 三层网络
	Product string `json:"product,omitempty"`
}

type DescribeGoogleRegionsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAWSRegionsResponseParams `json:"response,omitempty"`
}

type DescribeGoogleRegionsResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	// AWS的接入点相关的区域信息。
	CloudRegions []*CloudRegion `json:"cloudRegions,omitempty"`
}

type DescribeGoogleVlanUsageRequest struct {
	*common.BaseRequest

	DcId string `json:"dcId,omitempty"`
}

type DescribeGoogleVlanUsageResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeGoogleVlanUsageResponseParams `json:"response,omitempty"`
}

type DescribeGoogleVlanUsageResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	Start int `json:"start,omitempty"`

	End int `json:"end,omitempty"`

	UsedVlans []int `json:"usedVlans,omitempty"`
}
