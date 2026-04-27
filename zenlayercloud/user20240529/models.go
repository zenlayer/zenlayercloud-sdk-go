package user

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"


// DescribeMembersRequest 
type DescribeMembersRequest struct {
    *common.BaseRequest

    // Name 成员名称。
    Name *string `json:"name,omitempty"`

    // Username 成员邮箱。
    Username *string `json:"username,omitempty"`

    // PageNum 分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 分页大小。
    PageSize *int `json:"pageSize,omitempty"`

}

type DescribeMembersResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeMembersResponseParams `json:"response,omitempty"`

}

// DescribeMembersResponseParams 
type DescribeMembersResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 成员列表信息。
    DataSet []*Member `json:"dataSet,omitempty"`

}

// Member 描述成员的响应信息。包括成员关系Id，成员Id，邮箱等。
type Member struct {

    // MemberId 成员列表唯一ID。
    MemberId *string `json:"memberId,omitempty"`

    // UserUid 用户唯一ID。
    UserUid *string `json:"userUid,omitempty"`

    // FirstName 用户名称。
    FirstName *string `json:"firstName,omitempty"`

    // LastName 用户姓氏。
    LastName *string `json:"lastName,omitempty"`

    // Username 用户邮箱。
    Username *string `json:"username,omitempty"`

    // AssignmentStatus 用户邀请状态。
    // ACTIVE：已加入,PENDING: 待确认
    AssignmentStatus *string `json:"assignmentStatus,omitempty"`

    // UpdateTime 更新时间。
    UpdateTime *string `json:"updateTime,omitempty"`

    // CreateTime 加入时间。
    CreateTime *string `json:"createTime,omitempty"`

}

// DescribeMemberRequest 
type DescribeMemberRequest struct {
    *common.BaseRequest

    // UserUid 用户唯一ID。
    UserUid *string `json:"userUid,omitempty"`

}

type DescribeMemberResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeMemberResponseParams `json:"response,omitempty"`

}

// DescribeMemberResponseParams 
type DescribeMemberResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // MemberId 成员列表唯一ID。
    MemberId *string `json:"memberId,omitempty"`

    // UserUid 用户唯一ID。
    UserUid *string `json:"userUid,omitempty"`

    // FirstName 用户名称。
    FirstName *string `json:"firstName,omitempty"`

    // LastName 用户姓氏。
    LastName *string `json:"lastName,omitempty"`

    // Username 用户邮箱。
    Username *string `json:"username,omitempty"`

    // AssignmentStatus 用户邀请状态。
    // ACTIVE：已加入,PENDING: 待确认
    AssignmentStatus *string `json:"assignmentStatus,omitempty"`

    // UpdateTime 更新时间。
    UpdateTime *string `json:"updateTime,omitempty"`

    // CreateTime 加入时间。
    CreateTime *string `json:"createTime,omitempty"`

}

// InviteMemberRequest 
type InviteMemberRequest struct {
    *common.BaseRequest

    // Email 用户邮箱。
    Email *string `json:"email,omitempty"`

    // MemberGroupId 成员组唯一ID。
    MemberGroupId *string `json:"memberGroupId,omitempty"`

    // OpenApi 是否生成API访问密钥。
    OpenApi *bool `json:"openApi,omitempty"`

}

type InviteMemberResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteMemberRequest 
type DeleteMemberRequest struct {
    *common.BaseRequest

    // UserUid 用户唯一ID。
    UserUid *string `json:"userUid,omitempty"`

}

type DeleteMemberResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// SendInviteMemberEmailRequest 
type SendInviteMemberEmailRequest struct {
    *common.BaseRequest

    // UserUid 用户唯一ID。
    UserUid *string `json:"userUid,omitempty"`

}

type SendInviteMemberEmailResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribeMemberGroupsRequest 
type DescribeMemberGroupsRequest struct {
    *common.BaseRequest

    // Name 成员组名称。
    Name *string `json:"name,omitempty"`

    // PageNum 分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 分页大小。
    PageSize *int `json:"pageSize,omitempty"`

}

type DescribeMemberGroupsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeMemberGroupsResponseParams `json:"response,omitempty"`

}

// DescribeMemberGroupsResponseParams 
type DescribeMemberGroupsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 成员组列表信息。
    DataSet []*MemberGroup `json:"dataSet,omitempty"`

}

// MemberGroup 描述成员组的响应信息。包括成员组ID，名称等。
type MemberGroup struct {

    // MemberGroupId 成员组唯一ID。
    MemberGroupId *string `json:"memberGroupId,omitempty"`

    // Name 成员组名称。
    Name *string `json:"name,omitempty"`

    // CreateTime 成员组创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

// CreateMemberGroupRequest 
type CreateMemberGroupRequest struct {
    *common.BaseRequest

    // Name 成员组名称。
    Name *string `json:"name,omitempty"`

}

type CreateMemberGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateMemberGroupResponseParams `json:"response,omitempty"`

}

// CreateMemberGroupResponseParams 
type CreateMemberGroupResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // MemberGroupId 成员组唯一ID。
    MemberGroupId *string `json:"memberGroupId,omitempty"`

}

// ModifyMemberGroupRequest 
type ModifyMemberGroupRequest struct {
    *common.BaseRequest

    // Name 成员组唯一名称。
    Name *string `json:"name,omitempty"`

    // MemberGroupId 成员组唯一ID。
    MemberGroupId *string `json:"memberGroupId,omitempty"`

}

type ModifyMemberGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeleteMemberGroupRequest 
type DeleteMemberGroupRequest struct {
    *common.BaseRequest

    // MemberGroupId 成员组唯一ID。
    MemberGroupId *string `json:"memberGroupId,omitempty"`

}

type DeleteMemberGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// CreateMemberGroupPermissionRequest 
type CreateMemberGroupPermissionRequest struct {
    *common.BaseRequest

    // AllResource 是否为全部资源。
    // 1：是；0：否 ，当0 是resourceGroupUid为必填。
    AllResource *int `json:"allResource,omitempty"`

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Policies 访问策略唯一ID列表。
    Policies []string `json:"policies,omitempty"`

    // MemberGroupId 成员组唯一ID。
    MemberGroupId *string `json:"memberGroupId,omitempty"`

}

type CreateMemberGroupPermissionResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DescribePermissionsRequest 
type DescribePermissionsRequest struct {
    *common.BaseRequest

    // PolicyName 访问策略名称。
    PolicyName *string `json:"policyName,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // UserUid 用户唯一ID。
    UserUid *string `json:"userUid,omitempty"`

    // AllResource 是否为全部资源。
    // 1：是；0：否。
    AllResource *int `json:"allResource,omitempty"`

    // MemberGroupName 用户组名称。
    MemberGroupName *string `json:"memberGroupName,omitempty"`

    // MemberGroupId 用户组唯一ID。
    MemberGroupId *string `json:"memberGroupId,omitempty"`

    // Associated 权限类别 成员或者成员组 默认查询全部。
    // 成员：BY_SELF，成员组：BY_MEMBER_GROUP
    Associated *string `json:"associated,omitempty"`

    // Order 按时间排序默认1。
    // 1：顺序，0：倒序。
    Order *int `json:"order,omitempty"`

    // PageNum 分页数。
    PageNum *int `json:"pageNum,omitempty"`

    // PageSize 分页大小。
    PageSize *int `json:"pageSize,omitempty"`

}

type DescribePermissionsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePermissionsResponseParams `json:"response,omitempty"`

}

// DescribePermissionsResponseParams 
type DescribePermissionsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // TotalCount 符合条件的数据总数。
    TotalCount *int `json:"totalCount,omitempty"`

    // DataSet 权限列表信息。
    DataSet []*Permission `json:"dataSet,omitempty"`

}

// Permission 描述权限的响应信息。包括权限ID，状态，用户Id等。
type Permission struct {

    // PermissionId 权限唯一ID。
    PermissionId *string `json:"permissionId,omitempty"`

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // AllResource 是否为全部资源。
    // 1：是；0：否。
    AllResource *int `json:"allResource,omitempty"`

    // UserUid 用户唯一ID。
    UserUid *string `json:"userUid,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

    // ResourceGroupName 资源组名称。
    ResourceGroupName *string `json:"resourceGroupName,omitempty"`

    // PolicyId 访问策略唯一ID。
    PolicyId *string `json:"policyId,omitempty"`

    // PolicyName 访问策略名称。
    PolicyName *string `json:"policyName,omitempty"`

    // PolicyDesc 访问策略描述。
    PolicyDesc *string `json:"policyDesc,omitempty"`

    // FirstName 用户名称。
    FirstName *string `json:"firstName,omitempty"`

    // LastName 用户姓氏。
    LastName *string `json:"lastName,omitempty"`

    // Username 用户邮箱。
    Username *string `json:"username,omitempty"`

    // MemberGroupName 用户组名称。
    MemberGroupName *string `json:"memberGroupName,omitempty"`

    // MemberGroupId 用户组唯一ID。
    MemberGroupId *string `json:"memberGroupId,omitempty"`

    // Associated 权限类别。
    // 成员：SELF，成员组：USER_GROUP
    Associated *string `json:"associated,omitempty"`

}

// CreateMemberPermissionRequest 
type CreateMemberPermissionRequest struct {
    *common.BaseRequest

    // AllResource 是否为全部资源。
    // 1：是；0：否 ，当0 是resourceGroupUid为必填。
    AllResource *int `json:"allResource,omitempty"`

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Policies 访问策略唯一ID列表。
    Policies []string `json:"policies,omitempty"`

    // Users 用户唯一ID列表。
    Users []string `json:"users,omitempty"`

}

type CreateMemberPermissionResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// DeletePermissionRequest 
type DeletePermissionRequest struct {
    *common.BaseRequest

    // PermissionId 权限唯一ID。
    PermissionId *string `json:"permissionId,omitempty"`

}

type DeletePermissionResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

type DescribePoliciesRequest struct {
    *common.BaseRequest

}

type DescribePoliciesResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribePoliciesResponseParams `json:"response,omitempty"`

}

// DescribePoliciesResponseParams 
type DescribePoliciesResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // Policies 访问策略列表的数据。
    Policies []*Policy `json:"policies,omitempty"`

}

// Policy 描述访问策略的响应信息。包括ID，名称，描述等。
type Policy struct {

    // PolicyId 访问策略唯一ID。
    PolicyId *string `json:"policyId,omitempty"`

    // AllResource 权访问策略是否为全部资源组(1:是；0否)。
    AllResource *int `json:"allResource,omitempty"`

    // Name 访问策略显示名称。
    Name *string `json:"name,omitempty"`

    // Desc 访问策略描述信息。
    Desc *string `json:"desc,omitempty"`

    // AuthorizationTime 访问策略授权次数。
    AuthorizationTime *int `json:"authorizationTime,omitempty"`

}

type DescribeResourceGroupsRequest struct {
    *common.BaseRequest

}

type DescribeResourceGroupsResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *DescribeResourceGroupsResponseParams `json:"response,omitempty"`

}

// DescribeResourceGroupsResponseParams 
type DescribeResourceGroupsResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ResourceGroups 资源组列表信息。
    ResourceGroups []*ResourceGroup `json:"resourceGroups,omitempty"`

}

// ResourceGroup 描述资源组的信息。
type ResourceGroup struct {

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Name 资源组名称。
    Name *string `json:"name,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

// CreateResourceGroupRequest 
type CreateResourceGroupRequest struct {
    *common.BaseRequest

    // Name 资源组名称。
    Name *string `json:"name,omitempty"`

}

type CreateResourceGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *CreateResourceGroupResponseParams `json:"response,omitempty"`

}

// CreateResourceGroupResponseParams 
type CreateResourceGroupResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Name 资源组名称。
    Name *string `json:"name,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

// AddResourceResourceGroupRequest 
type AddResourceResourceGroupRequest struct {
    *common.BaseRequest

    // Resources 资源ID列表。
    Resources []string `json:"resources,omitempty"`

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type AddResourceResourceGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

// ModifyResourceGroupRequest 
type ModifyResourceGroupRequest struct {
    *common.BaseRequest

    // Name 资源组名称。
    Name *string `json:"name,omitempty"`

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type ModifyResourceGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response *ModifyResourceGroupResponseParams `json:"response,omitempty"`

}

// ModifyResourceGroupResponseParams 
type ModifyResourceGroupResponseParams struct {

    RequestId *string `json:"requestId,omitempty"`

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

    // Name 资源组名称。
    Name *string `json:"name,omitempty"`

    // CreateTime 创建时间。
    CreateTime *string `json:"createTime,omitempty"`

}

// DeleteResourceGroupRequest 
type DeleteResourceGroupRequest struct {
    *common.BaseRequest

    // ResourceGroupId 资源组唯一ID。
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`

}

type DeleteResourceGroupResponse struct {
    *common.BaseResponse

    RequestId *string `json:"requestId,omitempty"`

    Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response,omitempty"`

}

