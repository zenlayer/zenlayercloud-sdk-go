package zrm

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeTagsRequest 
type DescribeTagsRequest struct {
    *common.BaseRequest

    // PageNum 页码，默认值1。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 每页面展示数量，默认值20。
    PageSize *int `json:"pageSize,omitempty"`

    // KeySort 标签键排序方式：ascend（正序），descend（倒序）。
    KeySort *string `json:"keySort,omitempty"`

    // CreatedDateSort 创建时间排序方式：ascend（正序），descend（倒序）。
    // 默认倒序。
    CreatedDateSort *string `json:"createdDateSort,omitempty"`

    // TagKeys 筛选的标签键集合。
    TagKeys []string `json:"tagKeys,omitempty"`

    // Tags 筛选的标签集合。
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

type DescribeTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeTagsResponseParams `json:"response,omitempty"`

}

type DescribeTagsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的标签总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 标签列表。
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

    // CreatedDate 标签创建时间。
    CreatedDate *string `json:"createdDate,omitempty"`

}

// CreateTagsRequest 
type CreateTagsRequest struct {
    *common.BaseRequest

    // Tags 创建的标签。
    // 一次性最多创建20个标签。
    Tags []*Tag `json:"tags,omitempty"`

}

type CreateTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteTagsRequest 
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

// DescribeResourceTagsRequest 
type DescribeResourceTagsRequest struct {
    *common.BaseRequest

    // ResourceUuid 资源的唯一标识。
    ResourceUuid *string `json:"resourceUuid,omitempty"`

}

type DescribeResourceTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeResourceTagsResponseParams `json:"response,omitempty"`

}

type DescribeResourceTagsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 该资源绑定的标签总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 该资源绑定的标签列表。
    DataSet []*ResourceTag `json:"dataSet,omitempty"`

}

// ResourceTag 描述资源关联标签的基本信息
type ResourceTag struct {

    // Key 标签键。
    Key *string `json:"key,omitempty"`

    // Value 标签值。
    Value *string `json:"value,omitempty"`

    // CreatedDate 标签创建时间。
    CreatedDate *string `json:"createdDate,omitempty"`

}

// ModifyResourceTagsRequest 
type ModifyResourceTagsRequest struct {
    *common.BaseRequest

    // ResourceUuid 需要绑定的资源唯一标识。
    ResourceUuid *string `json:"resourceUuid,omitempty"`

    // ReplaceTags 需要更新的标签列表，包含标签键和值。
    ReplaceTags []*Tag `json:"replaceTags,omitempty"`

    // DeleteTagKeys 需要解绑的标签列表，包含标签键。
    DeleteTagKeys []string `json:"deleteTagKeys,omitempty"`

}

type ModifyResourceTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// TagBindResourcesRequest 
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

// TagUnbindResourcesRequest 
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

// DescribeResourceByTagsRequest 
type DescribeResourceByTagsRequest struct {
    *common.BaseRequest

    // Tags 查询的标签列表。
    Tags []*Tag `json:"tags,omitempty"`

    // TagKeys 查询的标签键列表。
    TagKeys []string `json:"tagKeys,omitempty"`

}

type DescribeResourceByTagsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeResourceByTagsResponseParams `json:"response,omitempty"`

}

type DescribeResourceByTagsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 绑定了指定标签的资源总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 绑定了指定标签的资源列表。
    DataSet []*ResourceInfo `json:"dataSet,omitempty"`

}

// ResourceInfo 描述资源的基本信息
type ResourceInfo struct {

    // ResourceType 资源类型。
    ResourceType *string `json:"resourceType,omitempty"`

    // ResourceUuid 资源唯一标识。
    ResourceUuid *string `json:"resourceUuid,omitempty"`

}

// DescribeResourceTypesRequest 
type DescribeResourceTypesRequest struct {
    *common.BaseRequest

    // Services 按所属服务筛选，精确匹配，多个服务之间为或的关系。
    // 不传时返回全部服务的资源类型。
    Services []string `json:"services,omitempty"`

    // ResourceTypes 按资源类型筛选，精确匹配，多个类型之间为或的关系。
    // 不传时返回全部资源类型。
    ResourceTypes []string `json:"resourceTypes,omitempty"`

}

type DescribeResourceTypesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeResourceTypesResponseParams `json:"response,omitempty"`

}

type DescribeResourceTypesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的资源类型总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 资源类型列表。
    DataSet []*ResourceTypeInfo `json:"dataSet,omitempty"`

}

// ResourceTypeInfo 描述资源类型的基本信息
type ResourceTypeInfo struct {

    // ResourceType 资源类型。
    ResourceType *string `json:"resourceType,omitempty"`

    // Name 资源类型名称。
    Name *string `json:"name,omitempty"`

    // Service 资源类型所属的服务。
    Service *string `json:"service,omitempty"`

}

// DescribeResourcesRequest 
type DescribeResourcesRequest struct {
    *common.BaseRequest

    // PageNum 页码，默认值1。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 每页面展示数量，默认值20。
    PageSize *int `json:"pageSize,omitempty"`

    // ResourceIds 资源唯一标识列表，精确匹配，多个标识之间为或的关系。
    ResourceIds []string `json:"resourceIds,omitempty"`

    // ResourceName 资源名称，模糊匹配。
    ResourceName *string `json:"resourceName,omitempty"`

    // ResourceGroupId 资源组唯一标识列表，精确匹配，为空代表查询所有资源组。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Tags 查询的标签列表。
    Tags []*Tag `json:"tags,omitempty"`

    // TagKeys 查询的标签键列表。
    TagKeys []string `json:"tagKeys,omitempty"`

}

type DescribeResourcesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeResourcesResponseParams `json:"response,omitempty"`

}

type DescribeResourcesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的资源总数量。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 资源列表。
    DataSet []*ResourceSearchInfo `json:"dataSet,omitempty"`

}

// ResourceSearchInfo 描述资源的基本信息
type ResourceSearchInfo struct {

    // ResourceId 资源唯一标识。
    ResourceId *string `json:"resourceId,omitempty"`

    // Name 资源名称。
    Name *string `json:"name,omitempty"`

    // ResourceType 资源类型。
    // 取值见 ~~DescribeResourceTypes~~。
    ResourceType *string `json:"resourceType,omitempty"`

    // CreateTime 资源创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ResourceGroupId 资源组唯一标识。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // Location 资源所属的地域信息。
    Location *string `json:"location,omitempty"`

    // Service 资源所属的服务。
    Service *string `json:"service,omitempty"`

}

