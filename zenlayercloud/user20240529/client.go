/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package user

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2024-05-29"
	SERVICE    = "user"
)

type Client struct {
	common.Client
}

func NewClientWithSecretKey(secretKeyId, secretKeyPassword string) (client *Client, err error) {
	return NewClient(common.NewConfig(), secretKeyId, secretKeyPassword)
}

func NewClient(config *common.Config, secretKeyId, secretKeyPassword string) (client *Client, err error) {
	client = &Client{}

	err = client.InitWithCredential(common.NewCredential(secretKeyId, secretKeyPassword))
	if err != nil {
		return nil, err
	}
	err = client.WithConfig(config)

	if err != nil {
		return nil, err
	}
	return client, nil
}


func NewDescribeMembersRequest() (request *DescribeMembersRequest) {
	request = &DescribeMembersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeMembers")

	return
}

func NewDescribeMembersResponse() (response *DescribeMembersResponse) {
	response = &DescribeMembersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeMembers 查询成员列表。用户可以根据名称、邮箱等信息来搜索成员信息。
func (c *Client) DescribeMembers(request *DescribeMembersRequest) (response *DescribeMembersResponse, err error) {
	response = NewDescribeMembersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeMemberRequest() (request *DescribeMemberRequest) {
	request = &DescribeMemberRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeMember")

	return
}

func NewDescribeMemberResponse() (response *DescribeMemberResponse) {
	response = &DescribeMemberResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeMember 查询单个成员信息。
func (c *Client) DescribeMember(request *DescribeMemberRequest) (response *DescribeMemberResponse, err error) {
	response = NewDescribeMemberResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInviteMemberRequest() (request *InviteMemberRequest) {
	request = &InviteMemberRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InviteMember")

	return
}

func NewInviteMemberResponse() (response *InviteMemberResponse) {
	response = &InviteMemberResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InviteMember 邀请单个成员信息。
func (c *Client) InviteMember(request *InviteMemberRequest) (response *InviteMemberResponse, err error) {
	response = NewInviteMemberResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteMemberRequest() (request *DeleteMemberRequest) {
	request = &DeleteMemberRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteMember")

	return
}

func NewDeleteMemberResponse() (response *DeleteMemberResponse) {
	response = &DeleteMemberResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteMember 删除一个成员信息。
func (c *Client) DeleteMember(request *DeleteMemberRequest) (response *DeleteMemberResponse, err error) {
	response = NewDeleteMemberResponse()
	err = c.ApiCall(request, response)
	return
}

func NewSendInviteMemberEmailRequest() (request *SendInviteMemberEmailRequest) {
	request = &SendInviteMemberEmailRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "SendInviteMemberEmail")

	return
}

func NewSendInviteMemberEmailResponse() (response *SendInviteMemberEmailResponse) {
	response = &SendInviteMemberEmailResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// SendInviteMemberEmail 重新发送一个成员信息的邮件。
func (c *Client) SendInviteMemberEmail(request *SendInviteMemberEmailRequest) (response *SendInviteMemberEmailResponse, err error) {
	response = NewSendInviteMemberEmailResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeMemberGroupsRequest() (request *DescribeMemberGroupsRequest) {
	request = &DescribeMemberGroupsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeMemberGroups")

	return
}

func NewDescribeMemberGroupsResponse() (response *DescribeMemberGroupsResponse) {
	response = &DescribeMemberGroupsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeMemberGroups 查询成员组列表。用户可以根据名称信息来搜索成员组信息。
func (c *Client) DescribeMemberGroups(request *DescribeMemberGroupsRequest) (response *DescribeMemberGroupsResponse, err error) {
	response = NewDescribeMemberGroupsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateMemberGroupRequest() (request *CreateMemberGroupRequest) {
	request = &CreateMemberGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateMemberGroup")

	return
}

func NewCreateMemberGroupResponse() (response *CreateMemberGroupResponse) {
	response = &CreateMemberGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateMemberGroup 创建一个成员组。
func (c *Client) CreateMemberGroup(request *CreateMemberGroupRequest) (response *CreateMemberGroupResponse, err error) {
	response = NewCreateMemberGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyMemberGroupRequest() (request *ModifyMemberGroupRequest) {
	request = &ModifyMemberGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyMemberGroup")

	return
}

func NewModifyMemberGroupResponse() (response *ModifyMemberGroupResponse) {
	response = &ModifyMemberGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyMemberGroup 修改一个成员组信息。
func (c *Client) ModifyMemberGroup(request *ModifyMemberGroupRequest) (response *ModifyMemberGroupResponse, err error) {
	response = NewModifyMemberGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteMemberGroupRequest() (request *DeleteMemberGroupRequest) {
	request = &DeleteMemberGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteMemberGroup")

	return
}

func NewDeleteMemberGroupResponse() (response *DeleteMemberGroupResponse) {
	response = &DeleteMemberGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteMemberGroup 删除一个成员组信息。
func (c *Client) DeleteMemberGroup(request *DeleteMemberGroupRequest) (response *DeleteMemberGroupResponse, err error) {
	response = NewDeleteMemberGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateMemberGroupPermissionRequest() (request *CreateMemberGroupPermissionRequest) {
	request = &CreateMemberGroupPermissionRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateMemberGroupPermission")

	return
}

func NewCreateMemberGroupPermissionResponse() (response *CreateMemberGroupPermissionResponse) {
	response = &CreateMemberGroupPermissionResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateMemberGroupPermission 授予一个成员组相关访问策略信息，从而对进行成员组下所有成员进行权限控制。
func (c *Client) CreateMemberGroupPermission(request *CreateMemberGroupPermissionRequest) (response *CreateMemberGroupPermissionResponse, err error) {
	response = NewCreateMemberGroupPermissionResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePermissionsRequest() (request *DescribePermissionsRequest) {
	request = &DescribePermissionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePermissions")

	return
}

func NewDescribePermissionsResponse() (response *DescribePermissionsResponse) {
	response = &DescribePermissionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePermissions 查询权限列表。用户可以根据成员ID、策略名称、成员组等信息来搜索权限信息。
func (c *Client) DescribePermissions(request *DescribePermissionsRequest) (response *DescribePermissionsResponse, err error) {
	response = NewDescribePermissionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateMemberPermissionRequest() (request *CreateMemberPermissionRequest) {
	request = &CreateMemberPermissionRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateMemberPermission")

	return
}

func NewCreateMemberPermissionResponse() (response *CreateMemberPermissionResponse) {
	response = &CreateMemberPermissionResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateMemberPermission 授予成员对于资源组相关的访问策略，从而对资源组进行权限控制。
func (c *Client) CreateMemberPermission(request *CreateMemberPermissionRequest) (response *CreateMemberPermissionResponse, err error) {
	response = NewCreateMemberPermissionResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeletePermissionRequest() (request *DeletePermissionRequest) {
	request = &DeletePermissionRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeletePermission")

	return
}

func NewDeletePermissionResponse() (response *DeletePermissionResponse) {
	response = &DeletePermissionResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeletePermission 删除一个权限信息。
func (c *Client) DeletePermission(request *DeletePermissionRequest) (response *DeletePermissionResponse, err error) {
	response = NewDeletePermissionResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePoliciesRequest() (request *DescribePoliciesRequest) {
	request = &DescribePoliciesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePolicies")

	return
}

func NewDescribePoliciesResponse() (response *DescribePoliciesResponse) {
	response = &DescribePoliciesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePolicies 查询访问策略。
func (c *Client) DescribePolicies(request *DescribePoliciesRequest) (response *DescribePoliciesResponse, err error) {
	response = NewDescribePoliciesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeResourceGroupsRequest() (request *DescribeResourceGroupsRequest) {
	request = &DescribeResourceGroupsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeResourceGroups")

	return
}

func NewDescribeResourceGroupsResponse() (response *DescribeResourceGroupsResponse) {
	response = &DescribeResourceGroupsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeResourceGroups 查询所有资源组列表信息。
func (c *Client) DescribeResourceGroups(request *DescribeResourceGroupsRequest) (response *DescribeResourceGroupsResponse, err error) {
	response = NewDescribeResourceGroupsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateResourceGroupRequest() (request *CreateResourceGroupRequest) {
	request = &CreateResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateResourceGroup")

	return
}

func NewCreateResourceGroupResponse() (response *CreateResourceGroupResponse) {
	response = &CreateResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateResourceGroup 创建一个资源组信息。
func (c *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (response *CreateResourceGroupResponse, err error) {
	response = NewCreateResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAddResourceResourceGroupRequest() (request *AddResourceResourceGroupRequest) {
	request = &AddResourceResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AddResourceResourceGroup")

	return
}

func NewAddResourceResourceGroupResponse() (response *AddResourceResourceGroupResponse) {
	response = &AddResourceResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AddResourceResourceGroup 将资源添加到一个资源组。
func (c *Client) AddResourceResourceGroup(request *AddResourceResourceGroupRequest) (response *AddResourceResourceGroupResponse, err error) {
	response = NewAddResourceResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyResourceGroupRequest() (request *ModifyResourceGroupRequest) {
	request = &ModifyResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyResourceGroup")

	return
}

func NewModifyResourceGroupResponse() (response *ModifyResourceGroupResponse) {
	response = &ModifyResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyResourceGroup 修改一个资源组信息。
func (c *Client) ModifyResourceGroup(request *ModifyResourceGroupRequest) (response *ModifyResourceGroupResponse, err error) {
	response = NewModifyResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteResourceGroupRequest() (request *DeleteResourceGroupRequest) {
	request = &DeleteResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteResourceGroup")

	return
}

func NewDeleteResourceGroupResponse() (response *DeleteResourceGroupResponse) {
	response = &DeleteResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteResourceGroup 删除一个资源组信息。
func (c *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
	response = NewDeleteResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

