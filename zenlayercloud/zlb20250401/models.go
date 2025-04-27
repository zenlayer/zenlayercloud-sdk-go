package zlb

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


type DescribeListenersRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例的ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerIds 要查询的负载均衡监听器 ID列表。
    ListenerIds []string `json:"listenerIds,omitempty"`

    // Protocol 要查询的监听器协议类型。
    Protocol *string `json:"protocol,omitempty"`

}

// DescribeListenersResponseParams 
type DescribeListenersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Listeners 监听器列表。
    Listeners []*Listener `json:"listeners,omitempty"`

}

// Listener 描述
type Listener struct {

    // ListenerId 负载均衡监听器 ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // ListenerName 监听器的名称。
    ListenerName *string `json:"listenerName,omitempty"`

    // Protocol 监听器协议。
    Protocol *string `json:"protocol,omitempty"`

    // Port 监听器端口。多个端口使用,分隔。当端口是范围时用`-`连接，例如：10000-10005。如果传多个单独的端口连续，将会被自动聚合为范围端口。
    Port *string `json:"port,omitempty"`

    // HealthCheck 监听器的健康检查信息。
    HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

    // Scheduler 监听器转发的方式。
    Scheduler *string `json:"scheduler,omitempty"`

    // CreateTime 创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

}

// HealthCheck 描述健康检查的相关信息。
type HealthCheck struct {

    // Enabled 是否开启健康检查。
    Enabled *bool `json:"enabled,omitempty"`

    // CheckType 健康检查使用的协议。当开启健康检查时， 该字段必填。
    CheckType *string `json:"checkType,omitempty"`

    // CheckPort 健康检查端口。默认为后端服务的端口，除非您希望指定特定端口，否则建议留空。
    CheckPort *int `json:"checkPort,omitempty"`

    // CheckDelayLoop 健康检查的检查间隔时间。单位:秒。
    CheckDelayLoop *int `json:"checkDelayLoop,omitempty"`

    // CheckConnTimeout 健康检查的连接超时时间。单位:秒。
    CheckConnTimeout *int `json:"checkConnTimeout,omitempty"`

    // CheckRetry 检查重试次数。
    CheckRetry *int `json:"checkRetry,omitempty"`

    // CheckDelayTry 健康检查重试的间隔时间。单位:秒。
    CheckDelayTry *int `json:"checkDelayTry,omitempty"`

    // CheckHttpGetUrl 健康检查路径。仅适用于HTTP_GET的协议。如果指定，必须以/开头。
    CheckHttpGetUrl *string `json:"checkHttpGetUrl,omitempty"`

    // CheckHttpStatusCode 健康检查状态码。仅适用于HTTP_GET的协议。
    CheckHttpStatusCode *int `json:"checkHttpStatusCode,omitempty"`

}

type DescribeListenersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeListenersResponseParams `json:"response,omitempty"`

}

type CreateListenerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例 ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerName 要创建的监听器名称。仅支持字母、数字、连字符 (-) 和句点 (.)，且开头和结尾必须是字母或数字。
    ListenerName *string `json:"listenerName,omitempty"`

    // HealthCheck 健康检查相关参数。如果不传该参数或者healthCheck.enable = false. 将会关闭健康检查。
    HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

    // Protocol 监听器协议。
    Protocol *string `json:"protocol,omitempty"`

    // Scheduler 监听器转发的方式。
    Scheduler *string `json:"scheduler,omitempty"`

    // Port 监听器端口。多个端口使用,分隔。当端口是范围时用`-`连接，例如：10000-10005。端口的取值范围为1～65535。请注意端口不能和该监听器的其他端口有重叠。
    Port *string `json:"port,omitempty"`

}

type CreateListenerResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ListenerId 创建的监听器ID。
    ListenerId *string `json:"listenerId,omitempty"`

}

type CreateListenerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateListenerResponseParams `json:"response,omitempty"`

}

type DeleteListenerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

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

type ModifyListenerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡器ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerId 负载均衡器的监听器ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // ListenerName 负载均衡器的监听器名称。不传则不会进行修改。
    ListenerName *string `json:"listenerName,omitempty"`

    // HealthCheck 负载均衡器的监听器健康检查。 不传则不会进行修改，如果开启或关闭，请设置`HealthCheck.enabled`字段。
    HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

    // Scheduler 负载均衡器的监听器调度方式。不传则不会进行修改。
    Scheduler *string `json:"scheduler,omitempty"`

    // Port 监听器端口。多个端口使用,分隔。当端口是范围时用`-`连接，例如：10000-10005。端口的取值范围为1～65535。不指定将不会进行修改。
    Port *string `json:"port,omitempty"`

}

type ModifyListenerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeLoadBalancerRegionsRequest struct {
    *common.BaseRequest

}

type DescribeLoadBalancerRegionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Regions 节点列表。
    Regions []*Region `json:"regions,omitempty"`

}

// Region 描述节点的基本信息。包括节点ID、名称等。
type Region struct {

    // RegionId 节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // CityName 节点所属的城市名称。
    CityName *string `json:"cityName,omitempty"`

    // CityCode 城市对应的三字码。
    CityCode *string `json:"cityCode,omitempty"`

}

type DescribeLoadBalancerRegionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeLoadBalancerRegionsResponseParams `json:"response,omitempty"`

}

type RegisterBackendRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例的 ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerId 负载均衡监听器的 ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // BackendServers 要绑定的后端服务器列表。注意：同一个实例在同一个监听器里只能存在一个。
    BackendServers []*BackendServer `json:"backendServers,omitempty"`

}

// BackendServer 后端服务器信息
type BackendServer struct {

    // InstanceId 实例ID。可通过 DescribeInstances 接口返回字段中的 InstanceId 字段获取。
    InstanceId *string `json:"instanceId,omitempty"`

    // PrivateIpAddress 实例上绑定网卡的私网IP地址。
    PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

    // Weight 后端服务修改后的转发权重。创建绑定服务器时默认值为100, 修改时如果不指定，则不会进行修改。删除时不需要指定该参数。
    Weight *int `json:"weight,omitempty"`

    // Port 请求转发和健康检查的目标端口。如果为空，将跟随监听器端口配置。删除时不需要指定该参数。
    Port *int `json:"port,omitempty"`

}

type RegisterBackendResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeregisterBackendRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例 ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerId 负载均衡监听器 ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // BackendServers 要删除的后端服务器列表。
    BackendServers []*BackendServer `json:"backendServers,omitempty"`

}

type DeregisterBackendResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type ModifyBackendRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例 ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerId 负载均衡监听器的 ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // BackendServers 要修改的后端服务器集合。
    BackendServers []*BackendServer `json:"backendServers,omitempty"`

}

type ModifyBackendResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeBackendsRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例 ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // ListenerId 负载均衡监听器的 ID。如果未指定，则查询所有后端。
    ListenerId *string `json:"listenerId,omitempty"`

}

type DescribeBackendsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Backends 负载均衡的后端服务器列表。
    Backends []*ListenerBackend `json:"backends,omitempty"`

}

// ListenerBackend 描述监听器后端服务器信息。
type ListenerBackend struct {

    // InstanceId 实例ID。可通过 DescribeInstances 接口返回字段中的 InstanceId 字段获取。
    InstanceId *string `json:"instanceId,omitempty"`

    // PrivateIpAddress 实例上绑定网卡的私网IP地址。
    PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

    // Weight 后端服务修改后的转发权重。创建绑定服务器时默认值为100, 修改时如果不指定，则不会进行修改。删除时不需要指定该参数。
    Weight *int `json:"weight,omitempty"`

    // BackendPort 请求转发和健康检查的目标端口。如果为空，将跟随监听器端口配置。删除时不需要指定该参数。
    BackendPort *int `json:"backendPort,omitempty"`

    // ListenerId 负载均衡监听器 ID。
    ListenerId *string `json:"listenerId,omitempty"`

    // ListenerName 监听器的名称。
    ListenerName *string `json:"listenerName,omitempty"`

    // Protocol 监听器协议。
    Protocol *string `json:"protocol,omitempty"`

    // ListenerPort 监听器端口。多个端口使用,分隔。当端口是范围时用`-`连接，例如：10000-10005。
    ListenerPort *string `json:"listenerPort,omitempty"`

}

type DescribeBackendsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBackendsResponseParams `json:"response,omitempty"`

}

type ModifyLoadBalancersAttributeRequest struct {
    *common.BaseRequest

    // LoadBalancerIds 负载均衡实例ID列表。
    LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

    // LoadBalancerName 负载均衡实例的名称。
    LoadBalancerName *string `json:"loadBalancerName,omitempty"`

}

type ModifyLoadBalancersAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeLoadBalancersRequest struct {
    *common.BaseRequest

    // RegionId 节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // VpcId 根据负载均衡后端服务器所属的VPC网络 ID过滤。
    VpcId *string `json:"vpcId,omitempty"`

    // LoadBalancerIds 负载均衡实例ID集合。实例ID数量上限为100个。
    LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

    // LoadBalancerName 负载均衡实例名称。可以在末尾加入*以支持模糊匹配。
    LoadBalancerName *string `json:"loadBalancerName,omitempty"`

    // PageSize 返回的分页大小。默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页页码。默认为1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeLoadBalancersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 满足过滤条件的负载均衡实例总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 返回的负载均衡实例列表。
    DataSet []*LoadBalancer `json:"dataSet,omitempty"`

}

// LoadBalancer 描述负载均衡实例的信息。
type LoadBalancer struct {

    // RegionId 区域节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // LoadBalancerId 负载均衡器ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

    // LoadBalancerName 负载均衡期名称。
    LoadBalancerName *string `json:"loadBalancerName,omitempty"`

    // VpcId 负载均衡后端服务器所属的VPC网络 ID。
    VpcId *string `json:"vpcId,omitempty"`

    // Status 负载均衡实例的状态。
    Status *string `json:"status,omitempty"`

    // PublicIpAddress 负载均衡实例的 VIP 列表。
    PublicIpAddress []string `json:"publicIpAddress,omitempty"`

    // ListenerCount 负载均衡器下监听器的数量。
    ListenerCount *int64 `json:"listenerCount,omitempty"`

    // CreateTime 创建时间。按照ISO8601标准表示，并且使用UTC时间, 格式为：YYYY-MM-DDThh:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

}

type DescribeLoadBalancersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeLoadBalancersResponseParams `json:"response,omitempty"`

}

type RestoreLoadBalancerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例 ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type RestoreLoadBalancerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type TerminateLoadBalancerRequest struct {
    *common.BaseRequest

    // LoadBalancerId 负载均衡实例 ID。
    LoadBalancerId *string `json:"loadBalancerId,omitempty"`

}

type TerminateLoadBalancerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type CreateLoadBalancerRequest struct {
    *common.BaseRequest

    // RegionId 区域节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // VpcId 负载均衡后端服务器所属的VPC网络 ID。
    VpcId *string `json:"vpcId,omitempty"`

    // LoadBalancerName 负载均衡实例名称。长度为1～64个字符。
    LoadBalancerName *string `json:"loadBalancerName,omitempty"`

    // InternetChargeType IP网络计费模式。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // IpNetworkType IP 的网络类型。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

    // BandwidthMbps EIP的最大出带宽。单位为Mbps。最大限制通常为10000，如果有额外要求， 请联系Support。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // TrafficPackageSize 流量包大小。指定此参数时，IP网络计费模式(`internetChargeType`) 需为`ByTrafficPackage`, 否则该参数不生效。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // BandwidthClusterId 共享带宽包ID。指定此参数时，IP网络计费模式(`internetChargeType`) 需为`BandwidthCluster`, 否则该参数不生效。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

    // ResourceGroupId 资源组ID。如果不指定，则会创建在默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Number 创建负载均衡的个数。
    Number *int `json:"number,omitempty"`

}

type CreateLoadBalancerResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 下单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // LoadBalancerIds 由负载均衡实例唯一 ID 组成的集合。
    LoadBalancerIds []string `json:"loadBalancerIds,omitempty"`

}

type CreateLoadBalancerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateLoadBalancerResponseParams `json:"response,omitempty"`

}

type InquiryPriceCreateLoadBalancerRequest struct {
    *common.BaseRequest

    // RegionId 区域节点ID。
    RegionId *string `json:"regionId,omitempty"`

    // InternetChargeType IP网络计费模式。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // IpNetworkType IP 的网络类型。
    IpNetworkType *string `json:"ipNetworkType,omitempty"`

    // BandwidthMbps EIP的最大出带宽。单位为Mbps。最大限制通常为10000，如果有额外要求， 请联系Support。
    BandwidthMbps *int `json:"bandwidthMbps,omitempty"`

    // TrafficPackageSize 流量包大小。指定此参数时，IP网络计费模式(`internetChargeType`) 需为`ByTrafficPackage`, 否则该参数不生效。
    TrafficPackageSize *float64 `json:"trafficPackageSize,omitempty"`

    // BandwidthClusterId 共享带宽包ID。指定此参数时，IP网络计费模式(`internetChargeType`) 需为`BandwidthCluster`, 否则该参数不生效。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

}

// InquiryPriceCreateLoadBalancerResponseParams 描述创建负载均衡实例的响应。
type InquiryPriceCreateLoadBalancerResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // LoadBalancerInstancePrice 负载均衡实例的价格。
    LoadBalancerInstancePrice *PriceItem `json:"loadBalancerInstancePrice,omitempty"`

    // EipPrice 弹性IP的价格。
    EipPrice *PriceItem `json:"eipPrice,omitempty"`

    // EipNetworkPrice 弹性IP的网络计费的价格。
    EipNetworkPrice *PriceItem `json:"eipNetworkPrice,omitempty"`

    // LcuPrice 负载均衡实例LCU的价格。
    LcuPrice *PriceItem `json:"lcuPrice,omitempty"`

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

type InquiryPriceCreateLoadBalancerResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryPriceCreateLoadBalancerResponseParams `json:"response,omitempty"`

}

