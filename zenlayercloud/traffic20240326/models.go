package traffic

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeBandwidthClustersRequest 
type DescribeBandwidthClustersRequest struct {
    *common.BaseRequest

    // BandwidthClusterIds 共享带宽包ID。最多支持100个ID查询。
    BandwidthClusterIds []string `json:"bandwidthClusterIds,omitempty"`

    // BandwidthClusterName 共享带宽包显示名称。支持模糊匹配。
    BandwidthClusterName *string `json:"bandwidthClusterName,omitempty"`

    // CityName 共享带宽包所包含的城市名称。
    CityName *string `json:"cityName,omitempty"`

    // PageNum 返回的分页大小。默认为20，最大为1000。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 返回的分页数。 默认为1。
    PageSize *int `json:"pageSize,omitempty"`

}

// DescribeBandwidthClustersResponseParams 
type DescribeBandwidthClustersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 共享带宽包结果集。
    DataSet []*BandwidthClusterInfo `json:"dataSet,omitempty"`

}

// BandwidthClusterInfo 共享带宽包的基本信息。
type BandwidthClusterInfo struct {

    // BandwidthClusterId 共享带宽包唯一ID。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

    // BandwidthClusterName 共享带宽包显示名称。
    BandwidthClusterName *string `json:"bandwidthClusterName,omitempty"`

    // NetworkType IP 网络类型。
    NetworkType *string `json:"networkType,omitempty"`

    // AreaCode 区域代号。
    AreaCode *string `json:"areaCode,omitempty"`

    // InternetChargeType 带宽计费方式。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // CommitBandwidthMbps 保底带宽。单位：Mbps。
    CommitBandwidthMbps *int `json:"commitBandwidthMbps,omitempty"`

    // Location 所属区域。
    Location *string `json:"location,omitempty"`

    // CreateTime 创建时间。格式为：YYYY-MM-DDThh:mm:ssZ。
    CreateTime *string `json:"createTime,omitempty"`

    // ResourceNumber 带宽包内的资源数量。
    ResourceNumber *int `json:"resourceNumber,omitempty"`

    // Cities 城市信息。
    Cities []*CityInfo `json:"cities,omitempty"`

}

// CityInfo 描述城市的基本信息。
type CityInfo struct {

    // CityName 城市名称。
    CityName *string `json:"cityName,omitempty"`

    // CityCode 城市三字码代号。
    CityCode *string `json:"cityCode,omitempty"`

}

type DescribeBandwidthClustersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBandwidthClustersResponseParams `json:"response,omitempty"`

}

type DescribeBandwidthClusterResourcesRequest struct {
    *common.BaseRequest

    // BandwidthClusterId 共享带宽包ID。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

}

type DescribeBandwidthClusterResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Resources 共享带宽包里的资源信息。
    Resources []*BandwidthClusterResource `json:"resources,omitempty"`

    // TotalCount 共享带宽包里的资源数量。
    TotalCount *int `json:"totalCount,omitempty"`

}

// BandwidthClusterResource 描述带宽包里的资源信息。
type BandwidthClusterResource struct {

    // ResourceId 资源ID。
    ResourceId *string `json:"resourceId,omitempty"`

    // ResourceType 资源类型。
    ResourceType *string `json:"resourceType,omitempty"`

}

type DescribeBandwidthClusterResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBandwidthClusterResourcesResponseParams `json:"response,omitempty"`

}

type DescribeBandwidthClusterAreasRequest struct {
    *common.BaseRequest

}

// DescribeBandwidthClusterAreasResponseParams 
type DescribeBandwidthClusterAreasResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Areas 可开通共享流量包区域列表。
    Areas []*BandwidthClusterAreaInfo `json:"areas,omitempty"`

}

// BandwidthClusterAreaInfo 描述带宽包的区域信息。
type BandwidthClusterAreaInfo struct {

    // AreaCode 带宽包区域代码。
    AreaCode *string `json:"areaCode,omitempty"`

    // AreaName 带宽包区域名称。
    AreaName *string `json:"areaName,omitempty"`

    // NetworkTypes 该地区支持的IP网络类型。
    NetworkTypes []string `json:"networkTypes,omitempty"`

}

type DescribeBandwidthClusterAreasResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBandwidthClusterAreasResponseParams `json:"response,omitempty"`

}

type DescribeBandwidthClusterTrafficRequest struct {
    *common.BaseRequest

    // BandwidthClusterId 共享带宽包ID。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

    // StartTime 查询开始时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 查询结束时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
    EndTime *string `json:"endTime,omitempty"`

    // City 按城市进行过滤，如LAX。该值可以是城市英文名也可以是城市码。指定该值，将会只返回该城市的带宽流量数据。当前带宽组的城市码和城市名称可以在[DescribeBandwidthClusters](../bandwidth-cluster/describebandwidthclusters.md)中的cities字段中查看。
    City *string `json:"city,omitempty"`

}

type DescribeBandwidthClusterTrafficResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // DataList 带宽数据列表。
    DataList []*BandwidthClusterTrafficData `json:"dataList,omitempty"`

    // In95 入口带宽95值。
    In95 *int64 `json:"in95,omitempty"`

    // In95Time 入口带宽95值时间。格式为：YYYY-MM-DDThh:mm:ssZ。
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

    // Out95Time 出口带宽95值时间。格式为：YYYY-MM-DDThh:mm:ssZ。
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

// BandwidthClusterTrafficData 共享带宽包带宽数据点。
type BandwidthClusterTrafficData struct {

    // InternetRX 入口带宽。单位：bps。
    InternetRX *int64 `json:"internetRX,omitempty"`

    // InternetTX 出口带宽。单位：bps。
    InternetTX *int64 `json:"internetTX,omitempty"`

    // Time 数据时间。格式为：YYYY-MM-DDThh:mm:ssZ。
    Time *string `json:"time,omitempty"`

}

type DescribeBandwidthClusterTrafficResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeBandwidthClusterTrafficResponseParams `json:"response,omitempty"`

}

// InquiryBandwidthClusterPriceRequest 
type InquiryBandwidthClusterPriceRequest struct {
    *common.BaseRequest

    // AreaCode 共享带宽包所属地域。 具体地域信息可以查询[DescribeBandwidthClusterAreas](../bandwidth-cluster/describebandwidthclusterareas.md)中的`areaCode`获取。
    AreaCode *string `json:"areaCode,omitempty"`

    // CommitBandwidthMbps 保底带宽值。单位: Mbps。
    CommitBandwidthMbps *int `json:"commitBandwidthMbps,omitempty"`

    // NetworkType IP 网络类型。
    NetworkType *string `json:"networkType,omitempty"`

    // InternetChargeType 带宽计费方式。默认为月95计费，如果需要日峰值计费，请联系Support开通相关权限。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

}

type InquiryBandwidthClusterPriceResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Price 描述带宽包的价格信息。
    Price *PriceItem `json:"price,omitempty"`

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

type InquiryBandwidthClusterPriceResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *InquiryBandwidthClusterPriceResponseParams `json:"response,omitempty"`

}

// CreateBandwidthClusterRequest 
type CreateBandwidthClusterRequest struct {
    *common.BaseRequest

    // AreaCode 共享带宽包所属地域。 具体地域信息可以查询[DescribeBandwidthClusterAreas](../bandwidth-cluster/describebandwidthclusterareas.md)中的`areaCode`获取。
    AreaCode *string `json:"areaCode,omitempty"`

    // CommitBandwidthMbps 保底带宽值。单位: Mbps。
    CommitBandwidthMbps *int `json:"commitBandwidthMbps,omitempty"`

    // NetworkType IP网络类型。城市共享带宽包该字段必填。
    NetworkType *string `json:"networkType,omitempty"`

    // InternetChargeType 带宽计费方式。默认为月95计费，如果需要日峰值计费，请联系Support开通相关权限。
    InternetChargeType *string `json:"internetChargeType,omitempty"`

    // Name 带宽包显示名称。如果未指定，默认会使用地域名称命名。
    Name *string `json:"name,omitempty"`

}

// CreateBandwidthClusterResponseParams 
type CreateBandwidthClusterResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // OrderNumber 订单编号。
    OrderNumber *string `json:"orderNumber,omitempty"`

    // BandwidthClusterId 创建的共享带宽包ID。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

}

type CreateBandwidthClusterResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateBandwidthClusterResponseParams `json:"response,omitempty"`

}

// ModifyBandwidthClusterAttributeRequest 修改共享带宽包属性的请求信息。
type ModifyBandwidthClusterAttributeRequest struct {
    *common.BaseRequest

    // BandwidthClusterId 共享带宽包ID。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

    // Name 要许改的共享带宽包显示名称。最长不超过255个字符。
    Name *string `json:"name,omitempty"`

}

type ModifyBandwidthClusterAttributeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeleteBandwidthClustersRequest struct {
    *common.BaseRequest

    // BandwidthClusterIds 共享带宽包ID列表。最多一次可传20个。
    BandwidthClusterIds []string `json:"bandwidthClusterIds,omitempty"`

}

type DeleteBandwidthClustersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type UpdateBandwidthClusterCommitBandwidthRequest struct {
    *common.BaseRequest

    // BandwidthClusterId 共享带宽包ID。
    BandwidthClusterId *string `json:"bandwidthClusterId,omitempty"`

    // CommitBandwidthMbps 保底带宽。
    CommitBandwidthMbps *int `json:"commitBandwidthMbps,omitempty"`

}

type UpdateBandwidthClusterCommitBandwidthResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// MigrateBandwidthClusterResourcesRequest 共享带宽包资源迁移的请求对象
type MigrateBandwidthClusterResourcesRequest struct {
    *common.BaseRequest

    // TargetBandwidthClusterId 要迁移的目标共享带宽包的ID。
    TargetBandwidthClusterId *string `json:"targetBandwidthClusterId,omitempty"`

    // ResourceIdList 要迁移的资源的ID列表。该资源必须属于其他共享带宽包且资源所处的原共享带宽包的`IP网络类型`以及`区域`需要和目标带宽包一致。
    ResourceIdList []string `json:"resourceIdList,omitempty"`

}

type MigrateBandwidthClusterResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

