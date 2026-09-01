package zsp

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeTicketsRequest 
type DescribeTicketsRequest struct {
    *common.BaseRequest

    // TicketNumber 工单编号，支持模糊查询。
    TicketNumber *string `json:"ticketNumber,omitempty"`

    // Type 工单类型。
    Type *string `json:"type,omitempty"`

    // Description 问题描述，支持模糊查询。
    Description *string `json:"description,omitempty"`

    // Statuses 工单状态列表。
    // 不填时返回所有状态的工单。
    Statuses []string `json:"statuses,omitempty"`

    // StartTime 创建时间的查询起点。
    // 时间格式：yyyy-MM-ddTHH:mm:ssZ。
    StartTime *string `json:"startTime,omitempty"`

    // EndTime 创建时间的查询终点。
    // 时间格式：yyyy-MM-ddTHH:mm:ssZ。
    EndTime *string `json:"endTime,omitempty"`

    // CreateTimeDesc 返回结果是否按创建时间倒序排列。
    CreateTimeDesc *bool `json:"createTimeDesc,omitempty"`

    // PageSize 分页大小。
    // 当未传递时，默认值为10。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 分页页码。
    // 当未传递时，默认值为1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribeTicketsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeTicketsResponseParams `json:"response,omitempty"`

}

// DescribeTicketsResponseParams 
type DescribeTicketsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 工单结果集。
    DataSet []*TicketItem `json:"dataSet,omitempty"`

}

// TicketItem 工单列表项。
type TicketItem struct {

    // TicketId 工单ID。
    TicketId *string `json:"ticketId,omitempty"`

    // TicketNumber 工单编号。
    // 工单尚在创建中（状态为CREATING）时为null。
    TicketNumber *string `json:"ticketNumber,omitempty"`

    // Type 工单类型。
    Type *string `json:"type,omitempty"`

    // Status 工单状态。
    Status *string `json:"status,omitempty"`

    // Description 问题描述。
    Description *string `json:"description,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

// CreateTicketRequest 
type CreateTicketRequest struct {
    *common.BaseRequest

    // Type 工单类型。
    Type *string `json:"type,omitempty"`

    // Description 问题描述。
    // 长度1到4000个字符。
    Description *string `json:"description,omitempty"`

    // ContactName 联系人姓名。
    // 不填时取当前登录用户的姓名。
    ContactName *string `json:"contactName,omitempty"`

    // ContactEmail 联系邮箱。
    // 工单的后续沟通邮件发送到该邮箱，工单详情也按该邮箱过滤邮件往来记录。
    // 不填时取当前登录用户的邮箱。
    ContactEmail *string `json:"contactEmail,omitempty"`

    // ContactPhone 联系电话。
    ContactPhone *string `json:"contactPhone,omitempty"`

}

type CreateTicketResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateTicketResponseParams `json:"response,omitempty"`

}

// CreateTicketResponseParams 
type CreateTicketResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TicketId 工单ID。
    // 后续查询工单详情使用该ID。
    TicketId *string `json:"ticketId,omitempty"`

}

// DescribeTicketDetailRequest 
type DescribeTicketDetailRequest struct {
    *common.BaseRequest

    // TicketId 工单ID。
    TicketId *string `json:"ticketId,omitempty"`

}

type DescribeTicketDetailResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeTicketDetailResponseParams `json:"response,omitempty"`

}

// DescribeTicketDetailResponseParams 
type DescribeTicketDetailResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TicketId 工单ID。
    TicketId *string `json:"ticketId,omitempty"`

    // TicketNumber 工单编号。
    // 工单尚在创建中（状态为CREATING）时为null。
    TicketNumber *string `json:"ticketNumber,omitempty"`

    // Type 工单类型。
    Type *string `json:"type,omitempty"`

    // Status 工单状态。
    Status *string `json:"status,omitempty"`

    // Description 问题描述。
    Description *string `json:"description,omitempty"`

    // ContactName 联系人姓名。
    ContactName *string `json:"contactName,omitempty"`

    // ContactEmail 联系邮箱。
    ContactEmail *string `json:"contactEmail,omitempty"`

    // ContactPhone 联系电话。
    ContactPhone *string `json:"contactPhone,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // Messages 工单沟通记录，按时间倒序排列，末条为提单时的初始描述。
    // 仅包含与联系邮箱相关的往来邮件。
    Messages []*TicketMessage `json:"messages,omitempty"`

}

// TicketMessage 工单沟通记录。
type TicketMessage struct {

    // Subject 邮件主题。
    Subject *string `json:"subject,omitempty"`

    // Content 邮件正文。
    // 仅保留本次回复内容，不含被引用的历史邮件。
    Content *string `json:"content,omitempty"`

    // FromName 发件人姓名。
    FromName *string `json:"fromName,omitempty"`

    // FromAddress 发件人邮箱。
    FromAddress *string `json:"fromAddress,omitempty"`

    // ToAddress 收件人邮箱。
    ToAddress *string `json:"toAddress,omitempty"`

    // CcAddress 抄送邮箱。
    CcAddress *string `json:"ccAddress,omitempty"`

    // MessageDate 邮件时间。
    MessageDate *string `json:"messageDate,omitempty"`

}

// DescribeTicketCreationQuotaRequest 
type DescribeTicketCreationQuotaRequest struct {
    *common.BaseRequest

}

type DescribeTicketCreationQuotaResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeTicketCreationQuotaResponseParams `json:"response,omitempty"`

}

// DescribeTicketCreationQuotaResponseParams 
type DescribeTicketCreationQuotaResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // CanCreate 当前是否还可以创建工单。
    CanCreate *bool `json:"canCreate,omitempty"`

    // CreatedCountLastHour 最近一小时内已创建的工单数。
    CreatedCountLastHour *int `json:"createdCountLastHour,omitempty"`

    // LimitPerHour 每小时可创建的工单数上限。
    LimitPerHour *int `json:"limitPerHour,omitempty"`

}

