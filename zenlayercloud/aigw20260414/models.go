package aigw

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeAiGatewaysRequest 
type DescribeAiGatewaysRequest struct {
    *common.BaseRequest

    // GatewayUuids AI网关UUID列表。
    GatewayUuids []string `json:"gatewayUuids,omitempty"`

    // GatewayName AI网关名称。
    GatewayName *string `json:"gatewayName,omitempty"`

    // Status AI网关状态。
    Status *string `json:"status,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageSize 分页每页数量。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 分页页码。
    PageNum *int `json:"pageNum,omitempty"`

    // TagKeys 根据标签键进行搜索。
    // 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。
    // 最长不得超过20个标签。
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

type DescribeAiGatewaysResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiGatewaysResponseParams `json:"response,omitempty"`

}

// DescribeAiGatewaysResponseParams 
type DescribeAiGatewaysResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet AI网关列表的数据。
    DataSet []*AiGatewayInfo `json:"dataSet,omitempty"`

}

// AiGatewayInfo 描述AI网关实例的信息。包括过期策略，状态，模型限制等。
type AiGatewayInfo struct {

    // GatewayUuid AI网关唯一ID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

    // GatewayName AI网关显示名称。
    GatewayName *string `json:"gatewayName,omitempty"`

    // Enabled 是否启用。
    Enabled *bool `json:"enabled,omitempty"`

    // CreatedAt 创建时间。
    CreatedAt *string `json:"createdAt,omitempty"`

    // UpdatedAt 更新时间。
    UpdatedAt *string `json:"updatedAt,omitempty"`

    // ResourceGroupId AI网关所属的资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // AccessLimit 配额。
    AccessLimit *int `json:"accessLimit,omitempty"`

    // ExpireType 过期时间类型。
    ExpireType *string `json:"expireType,omitempty"`

    // ExpireTime 过期时间。
    ExpireTime *string `json:"expireTime,omitempty"`

    // ModelAccess 模型访问限制。
    ModelAccess *string `json:"modelAccess,omitempty"`

    // Tags AI网关关联的标签。
    Tags *Tags `json:"tags,omitempty"`

}

// Tags 描述资源关联的标签信息。
type Tags struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// CreateAiGatewayRequest 
type CreateAiGatewayRequest struct {
    *common.BaseRequest

    // ModelUuids 需要关联的模型uuid列表。
    // 当modelAccess为`LIMITED`时必填。
    ModelUuids []string `json:"modelUuids,omitempty"`

    // GatewayName AI网关名称。
    // 范围2到63个字符。
    // 仅支持输入字母、数字、-和英文句点(.)。
    // 且必须以数字或字母开头和结尾。
    GatewayName *string `json:"gatewayName,omitempty"`

    // TokenLimit token配额。
    TokenLimit *int `json:"tokenLimit,omitempty"`

    // ExpireType token过期策略类型。
    // 可选范围：LONG_LIVED、CUSTOM。
    ExpireType *string `json:"expireType,omitempty"`

    // ExpireTime token过期时间。
    // 当expireType为`CUSTOM`时必填。
    ExpireTime *string `json:"expireTime,omitempty"`

    // ResourceGroupId 创建后AI网关所在的资源组ID，如不指定则放入默认资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ModelAccess 模型访问权限。
    // 可选范围：UNLIMITED、LIMITED。
    ModelAccess *string `json:"modelAccess,omitempty"`

    // Tags 创建AI网关时关联的标签。
    // 注意：·关联`标签键`不能重复。
    Tags *TagAssociation `json:"tags,omitempty"`

}

// TagAssociation 描述创建资源时同时绑定的标签对的信息。
type TagAssociation struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

type CreateAiGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateAiGatewayResponseParams `json:"response,omitempty"`

}

// CreateAiGatewayResponseParams 
type CreateAiGatewayResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // GatewayUuid AI网关uuid。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

    // Token AI网关token。
    Token *string `json:"token,omitempty"`

}

// StartAiGatewayRequest 
type StartAiGatewayRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

}

type StartAiGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// StopAiGatewayRequest 
type StopAiGatewayRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

}

type StopAiGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteAiGatewayRequest 
type DeleteAiGatewayRequest struct {
    *common.BaseRequest

    // GatewayUuid AI网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

}

type DeleteAiGatewayResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeAiGatewayModelsRequest 
type DescribeAiGatewayModelsRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

}

type DescribeAiGatewayModelsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiGatewayModelsResponseParams `json:"response,omitempty"`

}

// DescribeAiGatewayModelsResponseParams 
type DescribeAiGatewayModelsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Models 关联模型列表。
    Models []*AiGatewayModel `json:"models,omitempty"`

}

// AiGatewayModel 描述AI网关关联模型的信息。包模型uuid和模型名称。
type AiGatewayModel struct {

    // ModelUuid 模型uuid。
    ModelUuid *string `json:"modelUuid,omitempty"`

    // ModelName 模型名称。
    ModelName *string `json:"modelName,omitempty"`

}

// ModifyAiGatewayModelsRequest 
type ModifyAiGatewayModelsRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

    // ModelAccess 模型访问权限。
    ModelAccess *string `json:"modelAccess,omitempty"`

    // ModelUuids 模型UUID列表。
    ModelUuids []string `json:"modelUuids,omitempty"`

}

type ModifyAiGatewayModelsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeAiGatewayIpAclRequest 
type DescribeAiGatewayIpAclRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

}

type DescribeAiGatewayIpAclResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiGatewayIpAclResponseParams `json:"response,omitempty"`

}

// DescribeAiGatewayIpAclResponseParams 
type DescribeAiGatewayIpAclResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // IpAclType IP访问控制类型（白名单/黑名单）。
    IpAclType *string `json:"ipAclType,omitempty"`

    // IpAddresses IP地址列表。
    IpAddresses []string `json:"ipAddresses,omitempty"`

    // Enabled 是否启用。
    Enabled *bool `json:"enabled,omitempty"`

}

// ModifyAiGatewayIpAclRequest 
type ModifyAiGatewayIpAclRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

    // IpAclType IP访问控制类型（白名单/黑名单）。
    IpAclType *string `json:"ipAclType,omitempty"`

    // IpAddresses IP地址列表。
    IpAddresses []string `json:"ipAddresses,omitempty"`

    // Enabled 是否启用。
    Enabled *bool `json:"enabled,omitempty"`

}

type ModifyAiGatewayIpAclResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeAiGatewayExpireTimeRequest 
type DescribeAiGatewayExpireTimeRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

}

type DescribeAiGatewayExpireTimeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiGatewayExpireTimeResponseParams `json:"response,omitempty"`

}

// DescribeAiGatewayExpireTimeResponseParams 
type DescribeAiGatewayExpireTimeResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ExpireType 过期类型。
    ExpireType *string `json:"expireType,omitempty"`

    // ExpireTime 过期时间。
    ExpireTime *string `json:"expireTime,omitempty"`

}

// ModifyAiGatewayExpireTimeRequest 
type ModifyAiGatewayExpireTimeRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

    // ExpireType 过期类型。
    ExpireType *string `json:"expireType,omitempty"`

    // ExpireTime 过期时间。
    ExpireTime *string `json:"expireTime,omitempty"`

}

type ModifyAiGatewayExpireTimeResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeAiGatewayTokenLimitRequest 
type DescribeAiGatewayTokenLimitRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

}

type DescribeAiGatewayTokenLimitResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiGatewayTokenLimitResponseParams `json:"response,omitempty"`

}

// DescribeAiGatewayTokenLimitResponseParams 
type DescribeAiGatewayTokenLimitResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TokenLimit Token配额限制。
    TokenLimit *int `json:"tokenLimit,omitempty"`

}

// ModifyAiGatewayTokenLimitRequest 
type ModifyAiGatewayTokenLimitRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

    // TokenLimit Token限制（访问限制）。
    TokenLimit *int `json:"tokenLimit,omitempty"`

}

type ModifyAiGatewayTokenLimitResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyAiGatewayNameRequest 
type ModifyAiGatewayNameRequest struct {
    *common.BaseRequest

    // GatewayUuid 网关UUID。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

    // GatewayName 网关名称。
    GatewayName *string `json:"gatewayName,omitempty"`

}

type ModifyAiGatewayNameResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribeAiModelsRequest struct {
    *common.BaseRequest

}

type DescribeAiModelsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiModelsResponseParams `json:"response,omitempty"`

}

// DescribeAiModelsResponseParams 
type DescribeAiModelsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ModelSet AI网关模型。
    ModelSet []*AiModel `json:"modelSet,omitempty"`

}

// AiModel 描述AI网关模型的信息。包模型uuid、模型名称、模型类型等。
type AiModel struct {

    // ModelUuid 模型uuid。
    ModelUuid *string `json:"modelUuid,omitempty"`

    // ModelName 模型名称。
    ModelName *string `json:"modelName,omitempty"`

    // ModelType 模型类型。
    ModelType *string `json:"modelType,omitempty"`

    // ProviderName 厂商名称。
    ProviderName *string `json:"providerName,omitempty"`

}

// DescribeAiUsageDetailDataRequest 
type DescribeAiUsageDetailDataRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID。
    InstanceIds *string `json:"instanceIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Start 开始时间戳10位，示例：1659427200。
    Start *int64 `json:"start,omitempty"`

    // End 结束时间戳10位，示例：1659434400。
    End *int64 `json:"end,omitempty"`

}

type DescribeAiUsageDetailDataResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiUsageDetailDataResponseParams `json:"response,omitempty"`

}

// DescribeAiUsageDetailDataResponseParams 
type DescribeAiUsageDetailDataResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Request 请求数明细数据。
    Request []*UsageDataPoint `json:"request,omitempty"`

    // Token 总Token明细数据。
    Token []*UsageDataPoint `json:"token,omitempty"`

    // InputToken 输入Token明细数据。
    InputToken []*UsageDataPoint `json:"inputToken,omitempty"`

    // OutToken 输出Token明细数据。
    OutToken []*UsageDataPoint `json:"outToken,omitempty"`

}

// UsageDataPoint 用量数据点，包含时间戳和各模型的用量值
type UsageDataPoint struct {

    // Time 时间戳（毫秒）。
    Time *int64 `json:"time,omitempty"`

    // Data 各模型的用量值列表。
    Data []*ModelValueItem `json:"data,omitempty"`

}

// ModelValueItem 模型用量值项
type ModelValueItem struct {

    // Model 模型名称。
    Model *string `json:"model,omitempty"`

    // Value 用量值。
    Value *int64 `json:"value,omitempty"`

}

// DescribeAiMonthlyTotalCostRequest 
type DescribeAiMonthlyTotalCostRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID。
    InstanceIds *string `json:"instanceIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Month 月份，格式："2026-2" 或 "2026-02"。
    Month *string `json:"month,omitempty"`

}

type DescribeAiMonthlyTotalCostResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiMonthlyTotalCostResponseParams `json:"response,omitempty"`

}

// DescribeAiMonthlyTotalCostResponseParams 
type DescribeAiMonthlyTotalCostResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCost AI网关月总费用。
    TotalCost *float64 `json:"totalCost,omitempty"`

}

// DescribeAiModelDailyCostRequest 
type DescribeAiModelDailyCostRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID。
    InstanceIds *string `json:"instanceIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Month 月份，格式："2026-2" 或 "2026-02"。
    Month *string `json:"month,omitempty"`

}

type DescribeAiModelDailyCostResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiModelDailyCostResponseParams `json:"response,omitempty"`

}

// DescribeAiModelDailyCostResponseParams 
type DescribeAiModelDailyCostResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Dates 日期列表。
    Dates []string `json:"dates,omitempty"`

    // Series 日模型费用。
    Series []*ModelCostSeries `json:"series,omitempty"`

}

// ModelCostSeries 描述模型日费用信息。
type ModelCostSeries struct {

    // ModelName 模型名称。
    ModelName *string `json:"modelName,omitempty"`

    // Data 模型日费用列表。
    Data []float64 `json:"data,omitempty"`

}

// DescribeAiUsageDataRequest 
type DescribeAiUsageDataRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID。
    InstanceIds *string `json:"instanceIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Start 开始时间戳10位，示例：1659427200。
    Start *int64 `json:"start,omitempty"`

    // End 结束时间戳10位，示例：1659434400。
    End *int64 `json:"end,omitempty"`

}

type DescribeAiUsageDataResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiUsageDataResponseParams `json:"response,omitempty"`

}

// DescribeAiUsageDataResponseParams 
type DescribeAiUsageDataResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalInputTokens 总输入token数。
    TotalInputTokens *int64 `json:"totalInputTokens,omitempty"`

    // TotalOutputTokens 总输出token数。
    TotalOutputTokens *int64 `json:"totalOutputTokens,omitempty"`

    // InstanceModelToken 实例模型用量数据。
    InstanceModelToken []*UsageData `json:"instanceModelToken,omitempty"`

}

// UsageData 描述AI网关模型用量信息。
type UsageData struct {

    // ModelUuid 模型uuid。
    ModelUuid *string `json:"modelUuid,omitempty"`

    // ModelName 模型名称。
    ModelName *string `json:"modelName,omitempty"`

    // TotalTokens 模型总token数。
    TotalTokens *int64 `json:"totalTokens,omitempty"`

    // TotalRequests 模型总request数。
    TotalRequests *int64 `json:"totalRequests,omitempty"`

}

// DescribeAiModelDailyCacheHitRateRequest 
type DescribeAiModelDailyCacheHitRateRequest struct {
    *common.BaseRequest

    // InstanceIds 实例ID。
    InstanceIds *string `json:"instanceIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Month 月份，格式："2026-2" 或 "2026-02"。
    Month *string `json:"month,omitempty"`

}

type DescribeAiModelDailyCacheHitRateResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeAiModelDailyCacheHitRateResponseParams `json:"response,omitempty"`

}

// DescribeAiModelDailyCacheHitRateResponseParams 
type DescribeAiModelDailyCacheHitRateResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Dates 日期列表。
    Dates []string `json:"dates,omitempty"`

    // CacheHit 日模型缓存命中率。
    CacheHit []*ModelCacheHitSeries `json:"cacheHit,omitempty"`

}

// ModelCacheHitSeries 描述模型日缓存命中率信息。
type ModelCacheHitSeries struct {

    // ModelName 模型名称。
    ModelName *string `json:"modelName,omitempty"`

    // Data 模型日缓存命中率列表，格式如 "90%"。
    Data []string `json:"data,omitempty"`

}

