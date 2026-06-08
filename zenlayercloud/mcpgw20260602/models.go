package mcpgw

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeMcpMonthlyBillingRequest 
type DescribeMcpMonthlyBillingRequest struct {
    *common.BaseRequest

    // Month 月份，格式："2026-05"。
    Month *string `json:"month,omitempty"`

}

type DescribeMcpMonthlyBillingResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeMcpMonthlyBillingResponseParams `json:"response,omitempty"`

}

// DescribeMcpMonthlyBillingResponseParams 
type DescribeMcpMonthlyBillingResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Items 按Token+日期维度的计费明细。
    Items []*BillingItem `json:"items,omitempty"`

    // TotalUsage 当月用量合计（字符串，保留精度）。
    TotalUsage *string `json:"totalUsage,omitempty"`

}

// BillingItem 描述MCP网关按月计费明细项。
type BillingItem struct {

    // GatewayUuid MCP网关uuid。
    GatewayUuid *string `json:"gatewayUuid,omitempty"`

    // GatewayName MCP网关名称。
    GatewayName *string `json:"gatewayName,omitempty"`

    // Date 用量日期，格式 yyyy-MM-dd。
    Date *string `json:"date,omitempty"`

    // Unit 计量单位。
    Unit *string `json:"unit,omitempty"`

    // Usage 用量（字符串，保留精度）。
    Usage *string `json:"usage,omitempty"`

    // ModelName 模型名称。
    ModelName *string `json:"modelName,omitempty"`

    // UnitPrice 单价（字符串，保留精度）。
    UnitPrice *string `json:"unitPrice,omitempty"`

    // OriginalPrice 原价（字符串，保留精度）。
    OriginalPrice *string `json:"originalPrice,omitempty"`

}

