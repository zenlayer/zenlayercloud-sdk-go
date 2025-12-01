package pvtdns

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


type AddPrivateZoneRequest struct {
    *common.BaseRequest

    // ZoneName 域名。格式必须域名，最长不超过200个字符，也可以是单独的不带点的域名后缀（长度2-63）。不区分大小写。
    ZoneName *string `json:"zoneName,omitempty"`

    // ProxyPattern 是否开启子域名递归代理。
    ProxyPattern *string `json:"proxyPattern,omitempty"`

    // VpcIds 要绑定的VPC ID列表。
    VpcIds []string `json:"vpcIds,omitempty"`

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Remark 备注信息。最长不超过255个字符。
    Remark *string `json:"remark,omitempty"`

    // Tags 关联的标签。
    Tags *TagAssociation `json:"tags,omitempty"`

}

// TagAssociation 描述创建资源时同时绑定的标签对的信息。
type TagAssociation struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// Tag 描述一个标签键值对的信息。
type Tag struct {

    // Key 标签键。长度限制：1～64个字符。
    Key *string `json:"key,omitempty"`

    // Value 标签值。长度限制：1～64个字符。
    Value *string `json:"value,omitempty"`

}

type AddPrivateZoneResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ZoneId 新增的内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

}

type AddPrivateZoneResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *AddPrivateZoneResponseParams `json:"response,omitempty"`

}

type DescribePrivateZonesRequest struct {
    *common.BaseRequest

    // ZoneIds 根据内网权威域名ID列表过滤。最长不超过100个。
    ZoneIds []string `json:"zoneIds,omitempty"`

    // VpcIds 根据Zone绑定的VPC ID筛选过滤。最长不超过100个。
    VpcIds []string `json:"vpcIds,omitempty"`

    // ZoneName 根据内网权威域名过滤。
    ZoneName *string `json:"zoneName,omitempty"`

    // ResourceGroupId 根据资源组ID过滤。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // PageSize 返回的分页大小，默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数，默认为1。
    PageNum *int `json:"pageNum,omitempty"`

    // TagKeys 根据标签键进行搜索。 最长不得超过20个标签键。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 根据标签进行搜索。 最长不得超过20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribePrivateZonesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 满足过滤条件的内网权威域名总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 返回的内网权威域名列表数据。
    DataSet []*PrivateZone `json:"dataSet,omitempty"`

}

// PrivateZone 描述内网权威域名的信息。
type PrivateZone struct {

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // ZoneName 内网权威域名。
    ZoneName *string `json:"zoneName,omitempty"`

    // ProxyPattern 是否开启子域名递归代理。
    ProxyPattern *string `json:"proxyPattern,omitempty"`

    // Remark 备注信息。最长不超过255个字符。
    Remark *string `json:"remark,omitempty"`

    // RecordCount 解析记录数。
    RecordCount *int `json:"recordCount,omitempty"`

    // VpcIds 绑定的VPC ID列表。
    VpcIds []string `json:"vpcIds,omitempty"`

    // Tags 关联的标签。
    Tags *Tags `json:"tags,omitempty"`

    // ResourceGroup 内网权威域名所在的资源组信息。
    ResourceGroup *ResourceGroupInfo `json:"resourceGroup,omitempty"`

    // CreateTime 内网权威域名的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

// Tags 描述资源关联的标签信息。
type Tags struct {

    // Tags 标签对列表。
    Tags []*Tag `json:"tags,omitempty"`

}

// ResourceGroupInfo 描述资源所在资源组的相关信息，包括资源组名称和ID。
type ResourceGroupInfo struct {

    // ResourceGroupId 资源组ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

}

type DescribePrivateZonesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePrivateZonesResponseParams `json:"response,omitempty"`

}

type ModifyPrivateZoneRequest struct {
    *common.BaseRequest

    // ZoneId 要修改的内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Remark 备注信息。最长不超过255个字符。
    Remark *string `json:"remark,omitempty"`

    // ProxyPattern 是否开启子域名递归代理。
    ProxyPattern *string `json:"proxyPattern,omitempty"`

}

type ModifyPrivateZoneResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeletePrivateZoneRequest struct {
    *common.BaseRequest

    // ZoneId 要删除的内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

}

type DeletePrivateZoneResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type BindPrivateZoneVpcRequest struct {
    *common.BaseRequest

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // VpcIds 要绑定的VPC列表。
    VpcIds []string `json:"vpcIds,omitempty"`

}

type BindPrivateZoneVpcResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type UnbindPrivateZoneVpcRequest struct {
    *common.BaseRequest

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // VpcIds 要解除绑定的的VPC ID列表。
    VpcIds []string `json:"vpcIds,omitempty"`

}

type UnbindPrivateZoneVpcResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// AddPrivateZoneRecordRequest 
type AddPrivateZoneRecordRequest struct {
    *common.BaseRequest

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Type 解析记录类型。
    Type *string `json:"type,omitempty"`

    // RecordName 主机记录。主机记录就是域名前缀，常见用法有www、@、*（泛解析）、mail（常用于邮箱）。例如要解析@.example.com，主机记录要填写"@”，而不是空。
    RecordName *string `json:"recordName,omitempty"`

    // Value 解析记录值。
    Value *string `json:"value,omitempty"`

    // Weight 解析记录权重。
    Weight *int `json:"weight,omitempty"`

    // Ttl 域名本地缓存时间。单位秒，默认值60，可选值5-86400(一天)。
    Ttl *int `json:"ttl,omitempty"`

    // Priority MX记录的优先级。取值越小，优先级越高, 默认为1。
    Priority *int `json:"priority,omitempty"`

    // Remark 备注信息。最长不超过255个字符。
    Remark *string `json:"remark,omitempty"`

    // Line 解析线路。`default` 为默认线路，即全局线路。也可以指定某个节点ID，例如asia-east-1。
    Line *string `json:"line,omitempty"`

    // Status 解析记录的状态。
    Status *string `json:"status,omitempty"`

}

type AddPrivateZoneRecordResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // RecordId 内网权威域名解析记录ID。
    RecordId *string `json:"recordId,omitempty"`

}

type AddPrivateZoneRecordResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *AddPrivateZoneRecordResponseParams `json:"response,omitempty"`

}

// DescribePrivateZoneRecordsRequest 
type DescribePrivateZoneRecordsRequest struct {
    *common.BaseRequest

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // RecordIds 根据解析记录ID进行筛选。
    RecordIds []string `json:"recordIds,omitempty"`

    // Type 根据解析记录类型进行筛选。
    Type *string `json:"type,omitempty"`

    // Value 根据记录值进行筛选。
    Value *string `json:"value,omitempty"`

    // PageSize 返回的分页大小，默认为20，最大为1000。
    PageSize *int `json:"pageSize,omitempty"`

    // PageNum 返回的分页数，默认为1。
    PageNum *int `json:"pageNum,omitempty"`

}

type DescribePrivateZoneRecordsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 满足过滤条件的内网权威域名解析记录总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 返回的内网权威域名解析记录列表数据。
    DataSet []*PrivateZoneRecord `json:"dataSet,omitempty"`

}

// PrivateZoneRecord 描述内网权威域名解析记录的信息。
type PrivateZoneRecord struct {

    // RecordId 解析记录ID。
    RecordId *string `json:"recordId,omitempty"`

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // Type 解析记录类型。
    Type *string `json:"type,omitempty"`

    // RecordName 主机记录。
    RecordName *string `json:"recordName,omitempty"`

    // Value 解析记录值。
    Value *string `json:"value,omitempty"`

    // Weight 解析记录权重。
    Weight *int `json:"weight,omitempty"`

    // Ttl 域名本地缓存时间。单位秒。
    Ttl *int `json:"ttl,omitempty"`

    // Line 解析线路。`default` 为默认线路，即全局线路。也可以指定某个节点ID，例如asia-east-1。
    Line *string `json:"line,omitempty"`

    // Priority `MX`记录的优先级。取值越小，优先级越高。
    Priority *int `json:"priority,omitempty"`

    // Remark 备注信息。
    Remark *string `json:"remark,omitempty"`

    // Status 解析记录状态。
    Status *string `json:"status,omitempty"`

    // CreateTime 解析记录的创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

type DescribePrivateZoneRecordsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePrivateZoneRecordsResponseParams `json:"response,omitempty"`

}

type ModifyPrivateZoneRecordRequest struct {
    *common.BaseRequest

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // RecordId 要修改的内网权威域名解析记录ID。
    RecordId *string `json:"recordId,omitempty"`

    // Value 解析记录值。
    Value *string `json:"value,omitempty"`

    // Weight 解析记录权重。仅A或者AAAA支持。
    Weight *int `json:"weight,omitempty"`

    // Ttl 域名本地缓存时间。单位秒，默认值60，可选值5-86400(一天)。
    Ttl *int `json:"ttl,omitempty"`

    // Priority MX记录的优先级。取值越小，优先级越高, 默认为1。
    Priority *int `json:"priority,omitempty"`

    // Remark 备注信息。最长不超过255个字符。
    Remark *string `json:"remark,omitempty"`

}

type ModifyPrivateZoneRecordResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyPrivateZoneRecordsStatusRequest 
type ModifyPrivateZoneRecordsStatusRequest struct {
    *common.BaseRequest

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // RecordIds 要修改生效状态的内网权威域名解析记录ID。
    RecordIds []string `json:"recordIds,omitempty"`

    // Status 生效状态。
    Status *string `json:"status,omitempty"`

}

type ModifyPrivateZoneRecordsStatusResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DeletePrivateZoneRecordRequest struct {
    *common.BaseRequest

    // ZoneId 内网权威域名ID。
    ZoneId *string `json:"zoneId,omitempty"`

    // RecordIds 要删除的内网权威域名解析记录ID列表。
    RecordIds []string `json:"recordIds,omitempty"`

}

type DeletePrivateZoneRecordResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

