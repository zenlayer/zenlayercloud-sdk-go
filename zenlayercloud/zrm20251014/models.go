package zrm

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// CreateTagsRequest 批量创建标签的请求参数。
type CreateTagsRequest struct {
    *common.BaseRequest

    // Tags 创建的标签。一次性最多创建20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

// Tag 描述一个标签键值对的信息。
type Tag struct {

    // Key 标签键。长度限制：1～64个字符。
    Key *string `json:"key,omitempty"`

    // Value 标签值。长度限制：1～64个字符。
    Value *string `json:"value,omitempty"`

}

type CreateTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteTagsRequest 批量删除标签的请求参数。
type DeleteTagsRequest struct {
    *common.BaseRequest

    // Tags 需要删除的标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type DeleteTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeTagsRequest 标签分页列表的请求参数。
type DescribeTagsRequest struct {
    *common.BaseRequest

    // PageNum 页码，默认值1。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 每页面展示数量，默认值20。
    PageSize *int `json:"pageSize,omitempty"`

    // KeySort 标签键排序。
    KeySort *string `json:"keySort,omitempty"`

    // CreatedDateSort 创建时间排序，默认倒序。
    CreatedDateSort *string `json:"createdDateSort,omitempty"`

    // TagKeys 筛选的标签键集合。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 筛选的标签集合。
    Tags []*Tag `json:"tags,omitempty"`

}

type DescribeTagsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 列表总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 数据列表。
    DataSet []*TagInfo `json:"dataSet,omitempty"`

}

// TagInfo 描述标签的基本信息
type TagInfo struct {

    // Key 标签键。
    Key *string `json:"key,omitempty"`

    // Value 标签值。
    Value *string `json:"value,omitempty"`

    // BindResourceCount 标签下绑定的资源总数。
    BindResourceCount *int `json:"bindResourceCount,omitempty"`

    // CreatedDate 创建时间。
    CreatedDate *string `json:"createdDate,omitempty"`

}

type DescribeTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeTagsResponseParams `json:"response,omitempty"`

}

// TagBindResourcesRequest 标签批量绑定资源的请求参数。
type TagBindResourcesRequest struct {
    *common.BaseRequest

    // Tag 标签，包含标签键和值。
    Tag *Tag `json:"tag,omitempty"`

    // ResourceUuids 需要绑定的资源唯一标识列表。
    ResourceUuids []string `json:"resourceUuids,omitempty"`

}

type TagBindResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// TagUnbindResourcesRequest 标签批量解绑资源的请求参数。
type TagUnbindResourcesRequest struct {
    *common.BaseRequest

    // Tag 标签，包含标签键和值。
    Tag *Tag `json:"tag,omitempty"`

    // ResourceUuids 需要解绑的资源唯一标识列表。
    ResourceUuids []string `json:"resourceUuids,omitempty"`

}

type TagUnbindResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeResourceTagsRequest 获取资源绑定标签的请求参数。
type DescribeResourceTagsRequest struct {
    *common.BaseRequest

    // ResourceUuid 资源的唯一标识。
    ResourceUuid *string `json:"resourceUuid,omitempty"`

}

type DescribeResourceTagsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 列表总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 数据列表。
    DataSet []*ResourceTag `json:"dataSet,omitempty"`

}

// ResourceTag 描述资源关联标签的基本信息
type ResourceTag struct {

    // Key 标签键。
    Key *string `json:"key,omitempty"`

    // Value 标签值。
    Value *string `json:"value,omitempty"`

    // CreatedDate 创建时间。
    CreatedDate *string `json:"createdDate,omitempty"`

}

type DescribeResourceTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeResourceTagsResponseParams `json:"response,omitempty"`

}

