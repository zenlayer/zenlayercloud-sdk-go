package user

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeResourceGroupsRequest struct {
	*common.BaseRequest
}

type DescribeResourceGroupsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeResourceGroupsResponseParam `json:"response"`
}

type DescribeResourceGroupsResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	ResourceGroups []*ResourceGroupInfo `json:"resourceGroups,omitempty"`
}

type ResourceGroupInfo struct {
	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	Name string `json:"name,omitempty"`

	CreateTime string `json:"createTime,omitempty"`
}

type DescribeResourceTypesRequest struct {
	*common.BaseRequest
}

type DescribeResourceTypesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeResourceTypesResponseParam `json:"response"`
}

type DescribeResourceTypesResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	ResourceTypeList []*string `json:"resourceTypeList,omitempty"`
}

type DescribeResourcesByGroupRequest struct {
	*common.BaseRequest

	PageNum         int    `json:"pageNum,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
	ResourceType    string `json:"resourceType,omitempty"`
	UpdateTimeSort  string `json:"updateTimeSort,omitempty"`
}

type DescribeResourcesByGroupResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeResourcesByGroupResponseParam `json:"response"`
}

type DescribeResourcesByGroupResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string          `json:"requestId,omitempty"`
	PageNum   int             `json:"pageNum,omitempty"`
	PageSize  int             `json:"pageSize,omitempty"`
	Total     int             `json:"total,omitempty"`
	List      []*ResourceInfo `json:"list,omitempty"`
}

type ResourceInfo struct {
	ResourceId   string `json:"resourceId,omitempty"`
	Name         string `json:"name,omitempty"`
	AliasName    string `json:"aliasName,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	UpdateTime   string `json:"updateTime,omitempty"`
}
