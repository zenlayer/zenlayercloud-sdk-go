/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package zec

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2025-09-01"
	SERVICE    = "zec"
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


func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDisks")

	return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
	response = &DescribeDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeDisks 查询云盘的列表信息。
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	response = NewDescribeDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDiskRegionsRequest() (request *DescribeDiskRegionsRequest) {
	request = &DescribeDiskRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDiskRegions")

	return
}

func NewDescribeDiskRegionsResponse() (response *DescribeDiskRegionsResponse) {
	response = &DescribeDiskRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeDiskRegions 支持售卖云硬盘的节点。
func (c *Client) DescribeDiskRegions(request *DescribeDiskRegionsRequest) (response *DescribeDiskRegionsResponse, err error) {
	response = NewDescribeDiskRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateDisksRequest() (request *InquiryPriceCreateDisksRequest) {
	request = &InquiryPriceCreateDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateDisks")

	return
}

func NewInquiryPriceCreateDisksResponse() (response *InquiryPriceCreateDisksResponse) {
	response = &InquiryPriceCreateDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateDisks 创建一块或多块云硬盘的询价。
func (c *Client) InquiryPriceCreateDisks(request *InquiryPriceCreateDisksRequest) (response *InquiryPriceCreateDisksResponse, err error) {
	response = NewInquiryPriceCreateDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateDisksRequest() (request *CreateDisksRequest) {
	request = &CreateDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateDisks")

	return
}

func NewCreateDisksResponse() (response *CreateDisksResponse) {
	response = &CreateDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateDisks 创建一个或多个云硬盘。
func (c *Client) CreateDisks(request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
	response = NewCreateDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
	request = &AttachDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AttachDisks")

	return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
	response = &AttachDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AttachDisks 云硬盘挂在到实例上。
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
	response = NewAttachDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
	request = &DetachDisksRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DetachDisks")

	return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
	response = &DetachDisksResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DetachDisks 将云硬盘从主机实例上卸载。
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
	response = NewDetachDisksResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyDisksAttributesRequest() (request *ModifyDisksAttributesRequest) {
	request = &ModifyDisksAttributesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyDisksAttributes")

	return
}

func NewModifyDisksAttributesResponse() (response *ModifyDisksAttributesResponse) {
	response = &ModifyDisksAttributesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyDisksAttributes 修改一块或多块云硬盘属性。当前接口只支持修改名称。
func (c *Client) ModifyDisksAttributes(request *ModifyDisksAttributesRequest) (response *ModifyDisksAttributesResponse, err error) {
	response = NewModifyDisksAttributesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
	request = &ResizeDiskRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ResizeDisk")

	return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
	response = &ResizeDiskResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ResizeDisk 将一个云硬盘扩容到指定大小。
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
	response = NewResizeDiskResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseDiskRequest() (request *ReleaseDiskRequest) {
	request = &ReleaseDiskRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseDisk")

	return
}

func NewReleaseDiskResponse() (response *ReleaseDiskResponse) {
	response = &ReleaseDiskResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReleaseDisk 删除释放一块云硬盘。
func (c *Client) ReleaseDisk(request *ReleaseDiskRequest) (response *ReleaseDiskResponse, err error) {
	response = NewReleaseDiskResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewDiskRequest() (request *RenewDiskRequest) {
	request = &RenewDiskRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewDisk")

	return
}

func NewRenewDiskResponse() (response *RenewDiskResponse) {
	response = &RenewDiskResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewDisk 恢复云硬盘
func (c *Client) RenewDisk(request *RenewDiskRequest) (response *RenewDiskResponse, err error) {
	response = NewRenewDiskResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyDisksResourceGroupRequest() (request *ModifyDisksResourceGroupRequest) {
	request = &ModifyDisksResourceGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyDisksResourceGroup")

	return
}

func NewModifyDisksResourceGroupResponse() (response *ModifyDisksResourceGroupResponse) {
	response = &ModifyDisksResourceGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyDisksResourceGroup 修改云硬盘所属的资源组。
func (c *Client) ModifyDisksResourceGroup(request *ModifyDisksResourceGroupRequest) (response *ModifyDisksResourceGroupResponse, err error) {
	response = NewModifyDisksResourceGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeDiskCategoryRequest() (request *DescribeDiskCategoryRequest) {
	request = &DescribeDiskCategoryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeDiskCategory")

	return
}

func NewDescribeDiskCategoryResponse() (response *DescribeDiskCategoryResponse) {
	response = &DescribeDiskCategoryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeDiskCategory 获取某个区域支持的云盘类型。
func (c *Client) DescribeDiskCategory(request *DescribeDiskCategoryRequest) (response *DescribeDiskCategoryResponse, err error) {
	response = NewDescribeDiskCategoryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateVpcRequest() (request *CreateVpcRequest) {
	request = &CreateVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateVpc")

	return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
	response = &CreateVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateVpc 创建全球VPC。
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
	response = NewCreateVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteVpcRequest() (request *DeleteVpcRequest) {
	request = &DeleteVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteVpc")

	return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
	response = &DeleteVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteVpc 删除VPC
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
	response = NewDeleteVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyVpcsAttributeRequest() (request *ModifyVpcsAttributeRequest) {
	request = &ModifyVpcsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyVpcsAttribute")

	return
}

func NewModifyVpcsAttributeResponse() (response *ModifyVpcsAttributeResponse) {
	response = &ModifyVpcsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyVpcsAttribute 修改一个或多个VPC的属性。该接口只支持修改VPC的名称。
func (c *Client) ModifyVpcsAttribute(request *ModifyVpcsAttributeRequest) (response *ModifyVpcsAttributeResponse, err error) {
	response = NewModifyVpcsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSubnetRegionsRequest() (request *DescribeSubnetRegionsRequest) {
	request = &DescribeSubnetRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSubnetRegions")

	return
}

func NewDescribeSubnetRegionsResponse() (response *DescribeSubnetRegionsResponse) {
	response = &DescribeSubnetRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeSubnetRegions 查询支持创建子网区域以及是否IPv6。
func (c *Client) DescribeSubnetRegions(request *DescribeSubnetRegionsRequest) (response *DescribeSubnetRegionsResponse, err error) {
	response = NewDescribeSubnetRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
	request = &CreateSubnetRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateSubnet")

	return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
	response = &CreateSubnetResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateSubnet 创建子网。
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
	response = NewCreateSubnetResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
	request = &DeleteSubnetRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteSubnet")

	return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
	response = &DeleteSubnetResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteSubnet 删除一个子网。
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
	response = NewDeleteSubnetResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySubnetAttributeRequest() (request *ModifySubnetAttributeRequest) {
	request = &ModifySubnetAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySubnetAttribute")

	return
}

func NewModifySubnetAttributeResponse() (response *ModifySubnetAttributeResponse) {
	response = &ModifySubnetAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifySubnetAttribute 修改子网属性。包括名称，CIDR等。
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
	response = NewModifySubnetAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySubnetStackTypeRequest() (request *ModifySubnetStackTypeRequest) {
	request = &ModifySubnetStackTypeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySubnetStackType")

	return
}

func NewModifySubnetStackTypeResponse() (response *ModifySubnetStackTypeResponse) {
	response = &ModifySubnetStackTypeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifySubnetStackType 修改子网堆栈类型
func (c *Client) ModifySubnetStackType(request *ModifySubnetStackTypeRequest) (response *ModifySubnetStackTypeResponse, err error) {
	response = NewModifySubnetStackTypeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
	request = &DescribeVpcsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeVpcs")

	return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
	response = &DescribeVpcsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeVpcs 查询私有网络（VPC）列表，用户可以根据 VPC ID、VPC 名称等信息来筛选过滤VPC信息。
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
	response = NewDescribeVpcsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyVpcAttributeRequest() (request *ModifyVpcAttributeRequest) {
	request = &ModifyVpcAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyVpcAttribute")

	return
}

func NewModifyVpcAttributeResponse() (response *ModifyVpcAttributeResponse) {
	response = &ModifyVpcAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyVpcAttribute 修改私有网络（VPC）的相关属性。
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
	response = NewModifyVpcAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
	request = &DescribeSubnetsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSubnets")

	return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
	response = &DescribeSubnetsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeSubnets 查询子网列表信息。可以根据子网ID, 名称等信息筛选查询子网。
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
	response = NewDescribeSubnetsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySubnetsAttributeRequest() (request *ModifySubnetsAttributeRequest) {
	request = &ModifySubnetsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySubnetsAttribute")

	return
}

func NewModifySubnetsAttributeResponse() (response *ModifySubnetsAttributeResponse) {
	response = &ModifySubnetsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifySubnetsAttribute 批量修改子网的属性。
func (c *Client) ModifySubnetsAttribute(request *ModifySubnetsAttributeRequest) (response *ModifySubnetsAttributeResponse, err error) {
	response = NewModifySubnetsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
	request = &CreateSnapshotRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateSnapshot")

	return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
	response = &CreateSnapshotResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateSnapshot 对指定云盘创建快照。
func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
	response = NewCreateSnapshotResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySnapshotsAttributeRequest() (request *ModifySnapshotsAttributeRequest) {
	request = &ModifySnapshotsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySnapshotsAttribute")

	return
}

func NewModifySnapshotsAttributeResponse() (response *ModifySnapshotsAttributeResponse) {
	response = &ModifySnapshotsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifySnapshotsAttribute 修改指定快照的属性。
func (c *Client) ModifySnapshotsAttribute(request *ModifySnapshotsAttributeRequest) (response *ModifySnapshotsAttributeResponse, err error) {
	response = NewModifySnapshotsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
	request = &DeleteSnapshotsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteSnapshots")

	return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
	response = &DeleteSnapshotsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteSnapshots 删除指定快照集合。
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
	response = NewDeleteSnapshotsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
	request = &DescribeSnapshotsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSnapshots")

	return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
	response = &DescribeSnapshotsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeSnapshots 查询快照的详细信息。
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
	response = NewDescribeSnapshotsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewApplySnapshotRequest() (request *ApplySnapshotRequest) {
	request = &ApplySnapshotRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ApplySnapshot")

	return
}

func NewApplySnapshotResponse() (response *ApplySnapshotResponse) {
	response = &ApplySnapshotResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ApplySnapshot 回滚快照到原云盘。
func (c *Client) ApplySnapshot(request *ApplySnapshotRequest) (response *ApplySnapshotResponse, err error) {
	response = NewApplySnapshotResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
	request = &CreateAutoSnapshotPolicyRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateAutoSnapshotPolicy")

	return
}

func NewCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
	response = &CreateAutoSnapshotPolicyResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateAutoSnapshotPolicy 创建一个自动快照策略。
func (c *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
	response = NewCreateAutoSnapshotPolicyResponse()
	err = c.ApiCall(request, response)
	return
}

func NewApplyAutoSnapshotPolicyRequest() (request *ApplyAutoSnapshotPolicyRequest) {
	request = &ApplyAutoSnapshotPolicyRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ApplyAutoSnapshotPolicy")

	return
}

func NewApplyAutoSnapshotPolicyResponse() (response *ApplyAutoSnapshotPolicyResponse) {
	response = &ApplyAutoSnapshotPolicyResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ApplyAutoSnapshotPolicy 云硬盘添加到指定的自动快照策略。
func (c *Client) ApplyAutoSnapshotPolicy(request *ApplyAutoSnapshotPolicyRequest) (response *ApplyAutoSnapshotPolicyResponse, err error) {
	response = NewApplyAutoSnapshotPolicyResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCancelAutoSnapshotPolicyRequest() (request *CancelAutoSnapshotPolicyRequest) {
	request = &CancelAutoSnapshotPolicyRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CancelAutoSnapshotPolicy")

	return
}

func NewCancelAutoSnapshotPolicyResponse() (response *CancelAutoSnapshotPolicyResponse) {
	response = &CancelAutoSnapshotPolicyResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CancelAutoSnapshotPolicy 云硬盘从指定的自动快照策略中移除。
func (c *Client) CancelAutoSnapshotPolicy(request *CancelAutoSnapshotPolicyRequest) (response *CancelAutoSnapshotPolicyResponse, err error) {
	response = NewCancelAutoSnapshotPolicyResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
	request = &DescribeAutoSnapshotPoliciesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAutoSnapshotPolicies")

	return
}

func NewDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
	response = &DescribeAutoSnapshotPoliciesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAutoSnapshotPolicies 查询自动快照策略的列表数据。
func (c *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
	response = NewDescribeAutoSnapshotPoliciesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAutoSnapshotPolicyRequest() (request *ModifyAutoSnapshotPolicyRequest) {
	request = &ModifyAutoSnapshotPolicyRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAutoSnapshotPolicy")

	return
}

func NewModifyAutoSnapshotPolicyResponse() (response *ModifyAutoSnapshotPolicyResponse) {
	response = &ModifyAutoSnapshotPolicyResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyAutoSnapshotPolicy 修改自动快照策略的基本信息。
func (c *Client) ModifyAutoSnapshotPolicy(request *ModifyAutoSnapshotPolicyRequest) (response *ModifyAutoSnapshotPolicyResponse, err error) {
	response = NewModifyAutoSnapshotPolicyResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteAutoSnapshotPoliciesRequest() (request *DeleteAutoSnapshotPoliciesRequest) {
	request = &DeleteAutoSnapshotPoliciesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteAutoSnapshotPolicies")

	return
}

func NewDeleteAutoSnapshotPoliciesResponse() (response *DeleteAutoSnapshotPoliciesResponse) {
	response = &DeleteAutoSnapshotPoliciesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteAutoSnapshotPolicies 删除指定的自动快照策略。
func (c *Client) DeleteAutoSnapshotPolicies(request *DeleteAutoSnapshotPoliciesRequest) (response *DeleteAutoSnapshotPoliciesResponse, err error) {
	response = NewDeleteAutoSnapshotPoliciesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeNetworkInterfaceRegionsRequest() (request *DescribeNetworkInterfaceRegionsRequest) {
	request = &DescribeNetworkInterfaceRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeNetworkInterfaceRegions")

	return
}

func NewDescribeNetworkInterfaceRegionsResponse() (response *DescribeNetworkInterfaceRegionsResponse) {
	response = &DescribeNetworkInterfaceRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeNetworkInterfaceRegions 支持售卖网卡的区域信息
func (c *Client) DescribeNetworkInterfaceRegions(request *DescribeNetworkInterfaceRegionsRequest) (response *DescribeNetworkInterfaceRegionsResponse, err error) {
	response = NewDescribeNetworkInterfaceRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
	request = &DescribeNetworkInterfacesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeNetworkInterfaces")

	return
}

func NewDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
	response = &DescribeNetworkInterfacesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeNetworkInterfaces 查询一个或多个指定网卡的信息。用户可以根据 网卡ID、网卡名称等信息来搜索网卡信息。
func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
	response = NewDescribeNetworkInterfacesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyNetworkInterfaceAttributeRequest() (request *ModifyNetworkInterfaceAttributeRequest) {
	request = &ModifyNetworkInterfaceAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyNetworkInterfaceAttribute")

	return
}

func NewModifyNetworkInterfaceAttributeResponse() (response *ModifyNetworkInterfaceAttributeResponse) {
	response = &ModifyNetworkInterfaceAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyNetworkInterfaceAttribute 修改一张网卡的属性。
func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) (response *ModifyNetworkInterfaceAttributeResponse, err error) {
	response = NewModifyNetworkInterfaceAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyNetworkInterfacesAttributeRequest() (request *ModifyNetworkInterfacesAttributeRequest) {
	request = &ModifyNetworkInterfacesAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyNetworkInterfacesAttribute")

	return
}

func NewModifyNetworkInterfacesAttributeResponse() (response *ModifyNetworkInterfacesAttributeResponse) {
	response = &ModifyNetworkInterfacesAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyNetworkInterfacesAttribute 修改网卡的属性，目前只支持修改网卡的名称。
func (c *Client) ModifyNetworkInterfacesAttribute(request *ModifyNetworkInterfacesAttributeRequest) (response *ModifyNetworkInterfacesAttributeResponse, err error) {
	response = NewModifyNetworkInterfacesAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
	request = &CreateNetworkInterfaceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateNetworkInterface")

	return
}

func NewCreateNetworkInterfaceResponse() (response *CreateNetworkInterfaceResponse) {
	response = &CreateNetworkInterfaceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateNetworkInterface 创建一张辅助网卡。
func (c *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
	response = NewCreateNetworkInterfaceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteNetworkInterfaceRequest() (request *DeleteNetworkInterfaceRequest) {
	request = &DeleteNetworkInterfaceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteNetworkInterface")

	return
}

func NewDeleteNetworkInterfaceResponse() (response *DeleteNetworkInterfaceResponse) {
	response = &DeleteNetworkInterfaceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteNetworkInterface 删除一张辅助网卡。
func (c *Client) DeleteNetworkInterface(request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
	response = NewDeleteNetworkInterfaceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
	request = &AttachNetworkInterfaceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AttachNetworkInterface")

	return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
	response = &AttachNetworkInterfaceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AttachNetworkInterface 网卡绑定实例。
func (c *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
	response = NewAttachNetworkInterfaceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
	request = &DetachNetworkInterfaceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DetachNetworkInterface")

	return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
	response = &DetachNetworkInterfaceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DetachNetworkInterface 将一张网卡从实例上解绑。
func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
	response = NewDetachNetworkInterfaceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssignNetworkInterfaceIpv4Request() (request *AssignNetworkInterfaceIpv4Request) {
	request = &AssignNetworkInterfaceIpv4Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignNetworkInterfaceIpv4")

	return
}

func NewAssignNetworkInterfaceIpv4Response() (response *AssignNetworkInterfaceIpv4Response) {
	response = &AssignNetworkInterfaceIpv4Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssignNetworkInterfaceIpv4 网卡绑定内网IPv4
func (c *Client) AssignNetworkInterfaceIpv4(request *AssignNetworkInterfaceIpv4Request) (response *AssignNetworkInterfaceIpv4Response, err error) {
	response = NewAssignNetworkInterfaceIpv4Response()
	err = c.ApiCall(request, response)
	return
}

func NewBatchAssignNetworkInterfaceIpv4Request() (request *BatchAssignNetworkInterfaceIpv4Request) {
	request = &BatchAssignNetworkInterfaceIpv4Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "BatchAssignNetworkInterfaceIpv4")

	return
}

func NewBatchAssignNetworkInterfaceIpv4Response() (response *BatchAssignNetworkInterfaceIpv4Response) {
	response = &BatchAssignNetworkInterfaceIpv4Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// BatchAssignNetworkInterfaceIpv4 将一张网卡额外添加内网IPv4地址。
func (c *Client) BatchAssignNetworkInterfaceIpv4(request *BatchAssignNetworkInterfaceIpv4Request) (response *BatchAssignNetworkInterfaceIpv4Response, err error) {
	response = NewBatchAssignNetworkInterfaceIpv4Response()
	err = c.ApiCall(request, response)
	return
}

func NewUnassignNetworkInterfaceIpv4Request() (request *UnassignNetworkInterfaceIpv4Request) {
	request = &UnassignNetworkInterfaceIpv4Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnassignNetworkInterfaceIpv4")

	return
}

func NewUnassignNetworkInterfaceIpv4Response() (response *UnassignNetworkInterfaceIpv4Response) {
	response = &UnassignNetworkInterfaceIpv4Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnassignNetworkInterfaceIpv4 将网卡上的内网IPv4地址解绑。
func (c *Client) UnassignNetworkInterfaceIpv4(request *UnassignNetworkInterfaceIpv4Request) (response *UnassignNetworkInterfaceIpv4Response, err error) {
	response = NewUnassignNetworkInterfaceIpv4Response()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeNetworkInterfacePublicIPv6Request() (request *DescribeNetworkInterfacePublicIPv6Request) {
	request = &DescribeNetworkInterfacePublicIPv6Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeNetworkInterfacePublicIPv6")

	return
}

func NewDescribeNetworkInterfacePublicIPv6Response() (response *DescribeNetworkInterfacePublicIPv6Response) {
	response = &DescribeNetworkInterfacePublicIPv6Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeNetworkInterfacePublicIPv6 网卡的公网IPv6信息。
func (c *Client) DescribeNetworkInterfacePublicIPv6(request *DescribeNetworkInterfacePublicIPv6Request) (response *DescribeNetworkInterfacePublicIPv6Response, err error) {
	response = NewDescribeNetworkInterfacePublicIPv6Response()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPricePublicIpv6Request() (request *InquiryPricePublicIpv6Request) {
	request = &InquiryPricePublicIpv6Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPricePublicIpv6")

	return
}

func NewInquiryPricePublicIpv6Response() (response *InquiryPricePublicIpv6Response) {
	response = &InquiryPricePublicIpv6Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPricePublicIpv6 公网Ipv6流量包或固定带宽询价。
func (c *Client) InquiryPricePublicIpv6(request *InquiryPricePublicIpv6Request) (response *InquiryPricePublicIpv6Response, err error) {
	response = NewInquiryPricePublicIpv6Response()
	err = c.ApiCall(request, response)
	return
}

func NewAssignNetworkInterfaceIpv6Request() (request *AssignNetworkInterfaceIpv6Request) {
	request = &AssignNetworkInterfaceIpv6Request{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignNetworkInterfaceIpv6")

	return
}

func NewAssignNetworkInterfaceIpv6Response() (response *AssignNetworkInterfaceIpv6Response) {
	response = &AssignNetworkInterfaceIpv6Response{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssignNetworkInterfaceIpv6 给网卡添加IPv6。
func (c *Client) AssignNetworkInterfaceIpv6(request *AssignNetworkInterfaceIpv6Request) (response *AssignNetworkInterfaceIpv6Response, err error) {
	response = NewAssignNetworkInterfaceIpv6Response()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeImages")

	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeImages 查询某节点支持的镜像列表。
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	response = NewDescribeImagesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateImage")

	return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateImage 用实例创建自定义镜像。
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
	response = NewCreateImageResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyImagesAttributesRequest() (request *ModifyImagesAttributesRequest) {
	request = &ModifyImagesAttributesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyImagesAttributes")

	return
}

func NewModifyImagesAttributesResponse() (response *ModifyImagesAttributesResponse) {
	response = &ModifyImagesAttributesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyImagesAttributes 修改自定义镜像属性。
func (c *Client) ModifyImagesAttributes(request *ModifyImagesAttributesRequest) (response *ModifyImagesAttributesResponse, err error) {
	response = NewModifyImagesAttributesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
	request = &DeleteImagesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteImages")

	return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
	response = &DeleteImagesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteImages 删除一个或多个镜像。
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
	response = NewDeleteImagesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
	request = &DescribeSecurityGroupsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSecurityGroups")

	return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
	response = &DescribeSecurityGroupsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeSecurityGroups 查询安全组列表。
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
	response = NewDescribeSecurityGroupsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySecurityGroupsAttributeRequest() (request *ModifySecurityGroupsAttributeRequest) {
	request = &ModifySecurityGroupsAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySecurityGroupsAttribute")

	return
}

func NewModifySecurityGroupsAttributeResponse() (response *ModifySecurityGroupsAttributeResponse) {
	response = &ModifySecurityGroupsAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifySecurityGroupsAttribute 修改安全组的属性（目前只支持修改安全组的名称）。
func (c *Client) ModifySecurityGroupsAttribute(request *ModifySecurityGroupsAttributeRequest) (response *ModifySecurityGroupsAttributeResponse, err error) {
	response = NewModifySecurityGroupsAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeSecurityGroupRuleRequest() (request *DescribeSecurityGroupRuleRequest) {
	request = &DescribeSecurityGroupRuleRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeSecurityGroupRule")

	return
}

func NewDescribeSecurityGroupRuleResponse() (response *DescribeSecurityGroupRuleResponse) {
	response = &DescribeSecurityGroupRuleResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeSecurityGroupRule 查询指定安全组内的规则。
func (c *Client) DescribeSecurityGroupRule(request *DescribeSecurityGroupRuleRequest) (response *DescribeSecurityGroupRuleResponse, err error) {
	response = NewDescribeSecurityGroupRuleResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
	request = &CreateSecurityGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateSecurityGroup")

	return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
	response = &CreateSecurityGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateSecurityGroup 创建一个安全组。
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
	response = NewCreateSecurityGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
	request = &DeleteSecurityGroupRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteSecurityGroup")

	return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
	response = &DeleteSecurityGroupResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteSecurityGroup 删除一个安全组。
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
	response = NewDeleteSecurityGroupResponse()
	err = c.ApiCall(request, response)
	return
}

func NewConfigureSecurityGroupRulesRequest() (request *ConfigureSecurityGroupRulesRequest) {
	request = &ConfigureSecurityGroupRulesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ConfigureSecurityGroupRules")

	return
}

func NewConfigureSecurityGroupRulesResponse() (response *ConfigureSecurityGroupRulesResponse) {
	response = &ConfigureSecurityGroupRulesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ConfigureSecurityGroupRules 配置安全组规则。
func (c *Client) ConfigureSecurityGroupRules(request *ConfigureSecurityGroupRulesRequest) (response *ConfigureSecurityGroupRulesResponse, err error) {
	response = NewConfigureSecurityGroupRulesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssignSecurityGroupVpcRequest() (request *AssignSecurityGroupVpcRequest) {
	request = &AssignSecurityGroupVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignSecurityGroupVpc")

	return
}

func NewAssignSecurityGroupVpcResponse() (response *AssignSecurityGroupVpcResponse) {
	response = &AssignSecurityGroupVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssignSecurityGroupVpc VPC更换绑定安全组。
func (c *Client) AssignSecurityGroupVpc(request *AssignSecurityGroupVpcRequest) (response *AssignSecurityGroupVpcResponse, err error) {
	response = NewAssignSecurityGroupVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnAssignSecurityGroupVpcRequest() (request *UnAssignSecurityGroupVpcRequest) {
	request = &UnAssignSecurityGroupVpcRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnAssignSecurityGroupVpc")

	return
}

func NewUnAssignSecurityGroupVpcResponse() (response *UnAssignSecurityGroupVpcResponse) {
	response = &UnAssignSecurityGroupVpcResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnAssignSecurityGroupVpc VPC解绑安全组。
func (c *Client) UnAssignSecurityGroupVpc(request *UnAssignSecurityGroupVpcRequest) (response *UnAssignSecurityGroupVpcResponse, err error) {
	response = NewUnAssignSecurityGroupVpcResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipRegionsRequest() (request *DescribeEipRegionsRequest) {
	request = &DescribeEipRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEipRegions")

	return
}

func NewDescribeEipRegionsResponse() (response *DescribeEipRegionsResponse) {
	response = &DescribeEipRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEipRegions 查询支持售卖EIP的区域信息。
func (c *Client) DescribeEipRegions(request *DescribeEipRegionsRequest) (response *DescribeEipRegionsResponse, err error) {
	response = NewDescribeEipRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipInternetChargeTypesRequest() (request *DescribeEipInternetChargeTypesRequest) {
	request = &DescribeEipInternetChargeTypesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEipInternetChargeTypes")

	return
}

func NewDescribeEipInternetChargeTypesResponse() (response *DescribeEipInternetChargeTypesResponse) {
	response = &DescribeEipInternetChargeTypesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEipInternetChargeTypes 查询EIP支持的网络计费模式。
func (c *Client) DescribeEipInternetChargeTypes(request *DescribeEipInternetChargeTypesRequest) (response *DescribeEipInternetChargeTypesResponse, err error) {
	response = NewDescribeEipInternetChargeTypesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipRemoteRegionsRequest() (request *DescribeEipRemoteRegionsRequest) {
	request = &DescribeEipRemoteRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEipRemoteRegions")

	return
}

func NewDescribeEipRemoteRegionsResponse() (response *DescribeEipRemoteRegionsResponse) {
	response = &DescribeEipRemoteRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEipRemoteRegions 查询EIP支持的远程指向的节点信息。
func (c *Client) DescribeEipRemoteRegions(request *DescribeEipRemoteRegionsRequest) (response *DescribeEipRemoteRegionsResponse, err error) {
	response = NewDescribeEipRemoteRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipsRequest() (request *DescribeEipsRequest) {
	request = &DescribeEipsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEips")

	return
}

func NewDescribeEipsResponse() (response *DescribeEipsResponse) {
	response = &DescribeEipsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEips 指定条件查询已创建的弹性IPv4的信息。用户可以根据ID、名称等信息来搜索。
func (c *Client) DescribeEips(request *DescribeEipsRequest) (response *DescribeEipsResponse, err error) {
	response = NewDescribeEipsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateEipsRequest() (request *CreateEipsRequest) {
	request = &CreateEipsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateEips")

	return
}

func NewCreateEipsResponse() (response *CreateEipsResponse) {
	response = &CreateEipsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateEips 创建弹性公网IP。
func (c *Client) CreateEips(request *CreateEipsRequest) (response *CreateEipsResponse, err error) {
	response = NewCreateEipsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteEipRequest() (request *DeleteEipRequest) {
	request = &DeleteEipRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteEip")

	return
}

func NewDeleteEipResponse() (response *DeleteEipResponse) {
	response = &DeleteEipResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteEip 删除指定的弹性公网IP。
func (c *Client) DeleteEip(request *DeleteEipRequest) (response *DeleteEipResponse, err error) {
	response = NewDeleteEipResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewEipRequest() (request *RenewEipRequest) {
	request = &RenewEipRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewEip")

	return
}

func NewRenewEipResponse() (response *RenewEipResponse) {
	response = &RenewEipResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewEip 恢复弹性公网IP
func (c *Client) RenewEip(request *RenewEipRequest) (response *RenewEipResponse, err error) {
	response = NewRenewEipResponse()
	err = c.ApiCall(request, response)
	return
}

func NewConfigEipEgressIpRequest() (request *ConfigEipEgressIpRequest) {
	request = &ConfigEipEgressIpRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ConfigEipEgressIp")

	return
}

func NewConfigEipEgressIpResponse() (response *ConfigEipEgressIpResponse) {
	response = &ConfigEipEgressIpResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ConfigEipEgressIp 指定IP作为出口IP。
func (c *Client) ConfigEipEgressIp(request *ConfigEipEgressIpRequest) (response *ConfigEipEgressIpResponse, err error) {
	response = NewConfigEipEgressIpResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipPriceRequest() (request *DescribeEipPriceRequest) {
	request = &DescribeEipPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEipPrice")

	return
}

func NewDescribeEipPriceResponse() (response *DescribeEipPriceResponse) {
	response = &DescribeEipPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEipPrice 创建公网弹性IP询价。
func (c *Client) DescribeEipPrice(request *DescribeEipPriceRequest) (response *DescribeEipPriceResponse, err error) {
	response = NewDescribeEipPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewChangeEipInternetChargeTypeRequest() (request *ChangeEipInternetChargeTypeRequest) {
	request = &ChangeEipInternetChargeTypeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ChangeEipInternetChargeType")

	return
}

func NewChangeEipInternetChargeTypeResponse() (response *ChangeEipInternetChargeTypeResponse) {
	response = &ChangeEipInternetChargeTypeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ChangeEipInternetChargeType 变更弹性公网IP更网络计费模式。
func (c *Client) ChangeEipInternetChargeType(request *ChangeEipInternetChargeTypeRequest) (response *ChangeEipInternetChargeTypeResponse, err error) {
	response = NewChangeEipInternetChargeTypeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAvailableLanIpRequest() (request *AvailableLanIpRequest) {
	request = &AvailableLanIpRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AvailableLanIp")

	return
}

func NewAvailableLanIpResponse() (response *AvailableLanIpResponse) {
	response = &AvailableLanIpResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AvailableLanIp 查询可供弹性公网IP绑定的网卡及内网IP信息。
func (c *Client) AvailableLanIp(request *AvailableLanIpRequest) (response *AvailableLanIpResponse, err error) {
	response = NewAvailableLanIpResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipTrafficRequest() (request *DescribeEipTrafficRequest) {
	request = &DescribeEipTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEipTraffic")

	return
}

func NewDescribeEipTrafficResponse() (response *DescribeEipTrafficResponse) {
	response = &DescribeEipTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEipTraffic 查询弹性公网IP指定时间段内的流量信息。
func (c *Client) DescribeEipTraffic(request *DescribeEipTrafficRequest) (response *DescribeEipTrafficResponse, err error) {
	response = NewDescribeEipTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssociateEipAddressRequest() (request *AssociateEipAddressRequest) {
	request = &AssociateEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssociateEipAddress")

	return
}

func NewAssociateEipAddressResponse() (response *AssociateEipAddressResponse) {
	response = &AssociateEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssociateEipAddress 批量将弹性公网IP（EIP）绑定到同地域的云产品实例上。
func (c *Client) AssociateEipAddress(request *AssociateEipAddressRequest) (response *AssociateEipAddressResponse, err error) {
	response = NewAssociateEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnassociateEipAddressRequest() (request *UnassociateEipAddressRequest) {
	request = &UnassociateEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnassociateEipAddress")

	return
}

func NewUnassociateEipAddressResponse() (response *UnassociateEipAddressResponse) {
	response = &UnassociateEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnassociateEipAddress 将弹性公网IP（EIP）从绑定的云产品上解绑。
func (c *Client) UnassociateEipAddress(request *UnassociateEipAddressRequest) (response *UnassociateEipAddressResponse, err error) {
	response = NewUnassociateEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReplaceEipAddressRequest() (request *ReplaceEipAddressRequest) {
	request = &ReplaceEipAddressRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReplaceEipAddress")

	return
}

func NewReplaceEipAddressResponse() (response *ReplaceEipAddressResponse) {
	response = &ReplaceEipAddressResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReplaceEipAddress 替换一个或者多个弹性公网IP地址。
func (c *Client) ReplaceEipAddress(request *ReplaceEipAddressRequest) (response *ReplaceEipAddressResponse, err error) {
	response = NewReplaceEipAddressResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyEipAttributeRequest() (request *ModifyEipAttributeRequest) {
	request = &ModifyEipAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyEipAttribute")

	return
}

func NewModifyEipAttributeResponse() (response *ModifyEipAttributeResponse) {
	response = &ModifyEipAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyEipAttribute 修改弹性公网IP属性。
func (c *Client) ModifyEipAttribute(request *ModifyEipAttributeRequest) (response *ModifyEipAttributeResponse, err error) {
	response = NewModifyEipAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyEipBandwidthRequest() (request *ModifyEipBandwidthRequest) {
	request = &ModifyEipBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyEipBandwidth")

	return
}

func NewModifyEipBandwidthResponse() (response *ModifyEipBandwidthResponse) {
	response = &ModifyEipBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyEipBandwidth 调整弹性公网IP的带宽限速。
func (c *Client) ModifyEipBandwidth(request *ModifyEipBandwidthRequest) (response *ModifyEipBandwidthResponse, err error) {
	response = NewModifyEipBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}

func NewChangeEipBindTypeRequest() (request *ChangeEipBindTypeRequest) {
	request = &ChangeEipBindTypeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ChangeEipBindType")

	return
}

func NewChangeEipBindTypeResponse() (response *ChangeEipBindTypeResponse) {
	response = &ChangeEipBindTypeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ChangeEipBindType 弹性公网IP（EIP）更换绑定模式。
func (c *Client) ChangeEipBindType(request *ChangeEipBindTypeRequest) (response *ChangeEipBindTypeResponse, err error) {
	response = NewChangeEipBindTypeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeEipMonitorDataRequest() (request *DescribeEipMonitorDataRequest) {
	request = &DescribeEipMonitorDataRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeEipMonitorData")

	return
}

func NewDescribeEipMonitorDataResponse() (response *DescribeEipMonitorDataResponse) {
	response = &DescribeEipMonitorDataResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeEipMonitorData 查询一段时间的弹性公网IP监控指标数据。
func (c *Client) DescribeEipMonitorData(request *DescribeEipMonitorDataRequest) (response *DescribeEipMonitorDataResponse, err error) {
	response = NewDescribeEipMonitorDataResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
	request = &DescribeZonesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeZones")

	return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
	response = &DescribeZonesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeZones 查询可用区信息。包括名称，所属的节点等。
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	response = NewDescribeZonesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribePoolsRequest() (request *DescribePoolsRequest) {
	request = &DescribePoolsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribePools")

	return
}

func NewDescribePoolsResponse() (response *DescribePoolsResponse) {
	response = &DescribePoolsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribePools 查询公网IP池列表。
func (c *Client) DescribePools(request *DescribePoolsRequest) (response *DescribePoolsResponse, err error) {
	response = NewDescribePoolsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrRegionsRequest() (request *DescribeCidrRegionsRequest) {
	request = &DescribeCidrRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrRegions")

	return
}

func NewDescribeCidrRegionsResponse() (response *DescribeCidrRegionsResponse) {
	response = &DescribeCidrRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCidrRegions 支持售卖CIDR的节点信息。
func (c *Client) DescribeCidrRegions(request *DescribeCidrRegionsRequest) (response *DescribeCidrRegionsResponse, err error) {
	response = NewDescribeCidrRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrPriceRequest() (request *DescribeCidrPriceRequest) {
	request = &DescribeCidrPriceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrPrice")

	return
}

func NewDescribeCidrPriceResponse() (response *DescribeCidrPriceResponse) {
	response = &DescribeCidrPriceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCidrPrice 查询CIDR地址块售卖价格。
func (c *Client) DescribeCidrPrice(request *DescribeCidrPriceRequest) (response *DescribeCidrPriceResponse, err error) {
	response = NewDescribeCidrPriceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCidrsRequest() (request *DescribeCidrsRequest) {
	request = &DescribeCidrsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCidrs")

	return
}

func NewDescribeCidrsResponse() (response *DescribeCidrsResponse) {
	response = &DescribeCidrsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeCidrs 查询CIDR地址块列表
func (c *Client) DescribeCidrs(request *DescribeCidrsRequest) (response *DescribeCidrsResponse, err error) {
	response = NewDescribeCidrsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateCidrRequest() (request *CreateCidrRequest) {
	request = &CreateCidrRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateCidr")

	return
}

func NewCreateCidrResponse() (response *CreateCidrResponse) {
	response = &CreateCidrResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateCidr 创建CIDR地址段。
func (c *Client) CreateCidr(request *CreateCidrRequest) (response *CreateCidrResponse, err error) {
	response = NewCreateCidrResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyCidrAttributeRequest() (request *ModifyCidrAttributeRequest) {
	request = &ModifyCidrAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyCidrAttribute")

	return
}

func NewModifyCidrAttributeResponse() (response *ModifyCidrAttributeResponse) {
	response = &ModifyCidrAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyCidrAttribute 修改CIDR地址段的属性。 目前只能修改名称。
func (c *Client) ModifyCidrAttribute(request *ModifyCidrAttributeRequest) (response *ModifyCidrAttributeResponse, err error) {
	response = NewModifyCidrAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteCidrRequest() (request *DeleteCidrRequest) {
	request = &DeleteCidrRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteCidr")

	return
}

func NewDeleteCidrResponse() (response *DeleteCidrResponse) {
	response = &DeleteCidrResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteCidr 删除CIDR地址段。
func (c *Client) DeleteCidr(request *DeleteCidrRequest) (response *DeleteCidrResponse, err error) {
	response = NewDeleteCidrResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewCidrRequest() (request *RenewCidrRequest) {
	request = &RenewCidrRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewCidr")

	return
}

func NewRenewCidrResponse() (response *RenewCidrResponse) {
	response = &RenewCidrResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewCidr 将一个处于回收站的CIDR地址段恢复回正常。
func (c *Client) RenewCidr(request *RenewCidrRequest) (response *RenewCidrResponse, err error) {
	response = NewRenewCidrResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateBorderGatewayRequest() (request *CreateBorderGatewayRequest) {
	request = &CreateBorderGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateBorderGateway")

	return
}

func NewCreateBorderGatewayResponse() (response *CreateBorderGatewayResponse) {
	response = &CreateBorderGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateBorderGateway 在某节点为VPC创建一个边界网关。
func (c *Client) CreateBorderGateway(request *CreateBorderGatewayRequest) (response *CreateBorderGatewayResponse, err error) {
	response = NewCreateBorderGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBorderGatewaysRequest() (request *DescribeBorderGatewaysRequest) {
	request = &DescribeBorderGatewaysRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBorderGateways")

	return
}

func NewDescribeBorderGatewaysResponse() (response *DescribeBorderGatewaysResponse) {
	response = &DescribeBorderGatewaysResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBorderGateways 查询边界网关列表。
func (c *Client) DescribeBorderGateways(request *DescribeBorderGatewaysRequest) (response *DescribeBorderGatewaysResponse, err error) {
	response = NewDescribeBorderGatewaysResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteBorderGatewayRequest() (request *DeleteBorderGatewayRequest) {
	request = &DeleteBorderGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteBorderGateway")

	return
}

func NewDeleteBorderGatewayResponse() (response *DeleteBorderGatewayResponse) {
	response = &DeleteBorderGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteBorderGateway 删除一个指定的边界网关。
func (c *Client) DeleteBorderGateway(request *DeleteBorderGatewayRequest) (response *DeleteBorderGatewayResponse, err error) {
	response = NewDeleteBorderGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyBorderGatewayAsnRequest() (request *ModifyBorderGatewayAsnRequest) {
	request = &ModifyBorderGatewayAsnRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyBorderGatewayAsn")

	return
}

func NewModifyBorderGatewayAsnResponse() (response *ModifyBorderGatewayAsnResponse) {
	response = &ModifyBorderGatewayAsnResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyBorderGatewayAsn 修改边界网关的ASN。
func (c *Client) ModifyBorderGatewayAsn(request *ModifyBorderGatewayAsnRequest) (response *ModifyBorderGatewayAsnResponse, err error) {
	response = NewModifyBorderGatewayAsnResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyBorderGatewaysAttributeRequest() (request *ModifyBorderGatewaysAttributeRequest) {
	request = &ModifyBorderGatewaysAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyBorderGatewaysAttribute")

	return
}

func NewModifyBorderGatewaysAttributeResponse() (response *ModifyBorderGatewaysAttributeResponse) {
	response = &ModifyBorderGatewaysAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyBorderGatewaysAttribute 修改边界网关的属性，包括名称，路由级别，子网宣告控制等。
func (c *Client) ModifyBorderGatewaysAttribute(request *ModifyBorderGatewaysAttributeRequest) (response *ModifyBorderGatewaysAttributeResponse, err error) {
	response = NewModifyBorderGatewaysAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAvailableNatsRequest() (request *DescribeAvailableNatsRequest) {
	request = &DescribeAvailableNatsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAvailableNats")

	return
}

func NewDescribeAvailableNatsResponse() (response *DescribeAvailableNatsResponse) {
	response = &DescribeAvailableNatsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAvailableNats 获取可绑定边界网关的NAT网关列表。
func (c *Client) DescribeAvailableNats(request *DescribeAvailableNatsRequest) (response *DescribeAvailableNatsResponse, err error) {
	response = NewDescribeAvailableNatsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssignBorderGatewayRequest() (request *AssignBorderGatewayRequest) {
	request = &AssignBorderGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignBorderGateway")

	return
}

func NewAssignBorderGatewayResponse() (response *AssignBorderGatewayResponse) {
	response = &AssignBorderGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssignBorderGateway 边界网关绑定NAT网关。
func (c *Client) AssignBorderGateway(request *AssignBorderGatewayRequest) (response *AssignBorderGatewayResponse, err error) {
	response = NewAssignBorderGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnassignBorderGatewayRequest() (request *UnassignBorderGatewayRequest) {
	request = &UnassignBorderGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnassignBorderGateway")

	return
}

func NewUnassignBorderGatewayResponse() (response *UnassignBorderGatewayResponse) {
	response = &UnassignBorderGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnassignBorderGateway 解绑边界网关
func (c *Client) UnassignBorderGateway(request *UnassignBorderGatewayRequest) (response *UnassignBorderGatewayResponse, err error) {
	response = NewUnassignBorderGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewAssignBorderGatewayRouteRequest() (request *AssignBorderGatewayRouteRequest) {
	request = &AssignBorderGatewayRouteRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "AssignBorderGatewayRoute")

	return
}

func NewAssignBorderGatewayRouteResponse() (response *AssignBorderGatewayRouteResponse) {
	response = &AssignBorderGatewayRouteResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// AssignBorderGatewayRoute 边界网关绑定自定义路由。即自定义路由发布到zbg网关。
func (c *Client) AssignBorderGatewayRoute(request *AssignBorderGatewayRouteRequest) (response *AssignBorderGatewayRouteResponse, err error) {
	response = NewAssignBorderGatewayRouteResponse()
	err = c.ApiCall(request, response)
	return
}

func NewUnassignBorderGatewayRouteRequest() (request *UnassignBorderGatewayRouteRequest) {
	request = &UnassignBorderGatewayRouteRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "UnassignBorderGatewayRoute")

	return
}

func NewUnassignBorderGatewayRouteResponse() (response *UnassignBorderGatewayRouteResponse) {
	response = &UnassignBorderGatewayRouteResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// UnassignBorderGatewayRoute 边界网关路由宣告中移除指定的自定义路由。
func (c *Client) UnassignBorderGatewayRoute(request *UnassignBorderGatewayRouteRequest) (response *UnassignBorderGatewayRouteResponse, err error) {
	response = NewUnassignBorderGatewayRouteResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeZoneInstanceConfigInfosRequest() (request *DescribeZoneInstanceConfigInfosRequest) {
	request = &DescribeZoneInstanceConfigInfosRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeZoneInstanceConfigInfos")

	return
}

func NewDescribeZoneInstanceConfigInfosResponse() (response *DescribeZoneInstanceConfigInfosResponse) {
	response = &DescribeZoneInstanceConfigInfosResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeZoneInstanceConfigInfos 查询可用区售卖的机型信息
func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
	response = NewDescribeZoneInstanceConfigInfosResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeTimeZonesRequest() (request *DescribeTimeZonesRequest) {
	request = &DescribeTimeZonesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeTimeZones")

	return
}

func NewDescribeTimeZonesResponse() (response *DescribeTimeZonesResponse) {
	response = &DescribeTimeZonesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeTimeZones 查询时区信息
func (c *Client) DescribeTimeZones(request *DescribeTimeZonesRequest) (response *DescribeTimeZonesResponse, err error) {
	response = NewDescribeTimeZonesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
	request = &InquiryPriceCreateInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateInstance")

	return
}

func NewInquiryPriceCreateInstanceResponse() (response *InquiryPriceCreateInstanceResponse) {
	response = &InquiryPriceCreateInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateInstance 创建虚拟机实例询价。
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
	response = NewInquiryPriceCreateInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateZecInstancesRequest() (request *CreateZecInstancesRequest) {
	request = &CreateZecInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateZecInstances")

	return
}

func NewCreateZecInstancesResponse() (response *CreateZecInstancesResponse) {
	response = &CreateZecInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateZecInstances 创建一台或多台虚拟机实例。
func (c *Client) CreateZecInstances(request *CreateZecInstancesRequest) (response *CreateZecInstancesResponse, err error) {
	response = NewCreateZecInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstances")

	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstances 查询一台或多台虚拟机实例的信息。用户可以根据实例ID、实例名称等条件来查询实例的详细信息。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	response = NewDescribeInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstancesStatusRequest() (request *DescribeInstancesStatusRequest) {
	request = &DescribeInstancesStatusRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstancesStatus")

	return
}

func NewDescribeInstancesStatusResponse() (response *DescribeInstancesStatusResponse) {
	response = &DescribeInstancesStatusResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstancesStatus 查询实例的状态。
func (c *Client) DescribeInstancesStatus(request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
	response = NewDescribeInstancesStatusResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
	request = &ModifyInstancesAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstancesAttribute")

	return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
	response = &ModifyInstancesAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyInstancesAttribute 修改实例属性（名称）。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
	response = NewModifyInstancesAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StartInstances")

	return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
	response = &StartInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// StartInstances 启动一台或多台虚拟机实例。
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	response = NewStartInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StopInstances")

	return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
	response = &StopInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// StopInstances 关闭一台或多台虚拟机实例。
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	response = NewStopInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
	request = &RebootInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RebootInstances")

	return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
	response = &RebootInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RebootInstances 重启虚拟机实例。
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	response = NewRebootInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewResetInstancePasswordRequest() (request *ResetInstancePasswordRequest) {
	request = &ResetInstancePasswordRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ResetInstancePassword")

	return
}

func NewResetInstancePasswordResponse() (response *ResetInstancePasswordResponse) {
	response = &ResetInstancePasswordResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ResetInstancePassword 重置一台虚拟机实例密码。
func (c *Client) ResetInstancePassword(request *ResetInstancePasswordRequest) (response *ResetInstancePasswordResponse, err error) {
	response = NewResetInstancePasswordResponse()
	err = c.ApiCall(request, response)
	return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
	request = &ResetInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ResetInstance")

	return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
	response = &ResetInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ResetInstance 重装一台虚拟机实例操作系统。
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
	response = NewResetInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStartIpForwardRequest() (request *StartIpForwardRequest) {
	request = &StartIpForwardRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StartIpForward")

	return
}

func NewStartIpForwardResponse() (response *StartIpForwardResponse) {
	response = &StartIpForwardResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// StartIpForward 开启IP转发
func (c *Client) StartIpForward(request *StartIpForwardRequest) (response *StartIpForwardResponse, err error) {
	response = NewStartIpForwardResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStopIpForwardRequest() (request *StopIpForwardRequest) {
	request = &StopIpForwardRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StopIpForward")

	return
}

func NewStopIpForwardResponse() (response *StopIpForwardResponse) {
	response = &StopIpForwardResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// StopIpForward 关闭IP转发
func (c *Client) StopIpForward(request *StopIpForwardRequest) (response *StopIpForwardResponse, err error) {
	response = NewStopIpForwardResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStartAgentMonitorRequest() (request *StartAgentMonitorRequest) {
	request = &StartAgentMonitorRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StartAgentMonitor")

	return
}

func NewStartAgentMonitorResponse() (response *StartAgentMonitorResponse) {
	response = &StartAgentMonitorResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// StartAgentMonitor 开启Agent监控采集。
func (c *Client) StartAgentMonitor(request *StartAgentMonitorRequest) (response *StartAgentMonitorResponse, err error) {
	response = NewStartAgentMonitorResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStopAgentMonitorRequest() (request *StopAgentMonitorRequest) {
	request = &StopAgentMonitorRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StopAgentMonitor")

	return
}

func NewStopAgentMonitorResponse() (response *StopAgentMonitorResponse) {
	response = &StopAgentMonitorResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// StopAgentMonitor 关闭Agent监控采集。
func (c *Client) StopAgentMonitor(request *StopAgentMonitorRequest) (response *StopAgentMonitorResponse, err error) {
	response = NewStopAgentMonitorResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstanceTypeRequest() (request *ModifyInstanceTypeRequest) {
	request = &ModifyInstanceTypeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstanceType")

	return
}

func NewModifyInstanceTypeResponse() (response *ModifyInstanceTypeResponse) {
	response = &ModifyInstanceTypeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyInstanceType 变更实例的规格
func (c *Client) ModifyInstanceType(request *ModifyInstanceTypeRequest) (response *ModifyInstanceTypeResponse, err error) {
	response = NewModifyInstanceTypeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewChangeNicNetworkTypeRequest() (request *ChangeNicNetworkTypeRequest) {
	request = &ChangeNicNetworkTypeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ChangeNicNetworkType")

	return
}

func NewChangeNicNetworkTypeResponse() (response *ChangeNicNetworkTypeResponse) {
	response = &ChangeNicNetworkTypeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ChangeNicNetworkType 更改实例的网卡模式。
func (c *Client) ChangeNicNetworkType(request *ChangeNicNetworkTypeRequest) (response *ChangeNicNetworkTypeResponse, err error) {
	response = NewChangeNicNetworkTypeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseInstancesRequest() (request *ReleaseInstancesRequest) {
	request = &ReleaseInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseInstances")

	return
}

func NewReleaseInstancesResponse() (response *ReleaseInstancesResponse) {
	response = &ReleaseInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ReleaseInstances 销毁一台或多台虚拟机实例。
func (c *Client) ReleaseInstances(request *ReleaseInstancesRequest) (response *ReleaseInstancesResponse, err error) {
	response = NewReleaseInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeVncUrlRequest() (request *DescribeVncUrlRequest) {
	request = &DescribeVncUrlRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeVncUrl")

	return
}

func NewDescribeVncUrlResponse() (response *DescribeVncUrlResponse) {
	response = &DescribeVncUrlResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeVncUrl 获取实例VNC地址。
func (c *Client) DescribeVncUrl(request *DescribeVncUrlRequest) (response *DescribeVncUrlResponse, err error) {
	response = NewDescribeVncUrlResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceMonitorDataRequest() (request *DescribeInstanceMonitorDataRequest) {
	request = &DescribeInstanceMonitorDataRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceMonitorData")

	return
}

func NewDescribeInstanceMonitorDataResponse() (response *DescribeInstanceMonitorDataResponse) {
	response = &DescribeInstanceMonitorDataResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeInstanceMonitorData 查询一段时间的实例的监控指标数据。包括CPU,内存等相关指标数据。
func (c *Client) DescribeInstanceMonitorData(request *DescribeInstanceMonitorDataRequest) (response *DescribeInstanceMonitorDataResponse, err error) {
	response = NewDescribeInstanceMonitorDataResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateNatGatewayRequest() (request *CreateNatGatewayRequest) {
	request = &CreateNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateNatGateway")

	return
}

func NewCreateNatGatewayResponse() (response *CreateNatGatewayResponse) {
	response = &CreateNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateNatGateway 创建NAT网关。
func (c *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
	response = NewCreateNatGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyNatGatewayAttributeRequest() (request *ModifyNatGatewayAttributeRequest) {
	request = &ModifyNatGatewayAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyNatGatewayAttribute")

	return
}

func NewModifyNatGatewayAttributeResponse() (response *ModifyNatGatewayAttributeResponse) {
	response = &ModifyNatGatewayAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyNatGatewayAttribute 修改NAT网关的属性。
func (c *Client) ModifyNatGatewayAttribute(request *ModifyNatGatewayAttributeRequest) (response *ModifyNatGatewayAttributeResponse, err error) {
	response = NewModifyNatGatewayAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeNatGatewayRegionsRequest() (request *DescribeNatGatewayRegionsRequest) {
	request = &DescribeNatGatewayRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeNatGatewayRegions")

	return
}

func NewDescribeNatGatewayRegionsResponse() (response *DescribeNatGatewayRegionsResponse) {
	response = &DescribeNatGatewayRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeNatGatewayRegions 支持售卖NAT网关的区域信息。
func (c *Client) DescribeNatGatewayRegions(request *DescribeNatGatewayRegionsRequest) (response *DescribeNatGatewayRegionsResponse, err error) {
	response = NewDescribeNatGatewayRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
	request = &DescribeNatGatewaysRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeNatGateways")

	return
}

func NewDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
	response = &DescribeNatGatewaysResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeNatGateways 查询NAT网关列表。
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
	response = NewDescribeNatGatewaysResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeNatGatewayDetailRequest() (request *DescribeNatGatewayDetailRequest) {
	request = &DescribeNatGatewayDetailRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeNatGatewayDetail")

	return
}

func NewDescribeNatGatewayDetailResponse() (response *DescribeNatGatewayDetailResponse) {
	response = &DescribeNatGatewayDetailResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeNatGatewayDetail 查询NAT网关详情规则表。
func (c *Client) DescribeNatGatewayDetail(request *DescribeNatGatewayDetailRequest) (response *DescribeNatGatewayDetailResponse, err error) {
	response = NewDescribeNatGatewayDetailResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteNatGatewayRequest() (request *DeleteNatGatewayRequest) {
	request = &DeleteNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteNatGateway")

	return
}

func NewDeleteNatGatewayResponse() (response *DeleteNatGatewayResponse) {
	response = &DeleteNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteNatGateway 删除一个指定的NAT网关。
func (c *Client) DeleteNatGateway(request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
	response = NewDeleteNatGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewNatGatewayRequest() (request *RenewNatGatewayRequest) {
	request = &RenewNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewNatGateway")

	return
}

func NewRenewNatGatewayResponse() (response *RenewNatGatewayResponse) {
	response = &RenewNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RenewNatGateway 续费一个指定的NAT网关。
func (c *Client) RenewNatGateway(request *RenewNatGatewayRequest) (response *RenewNatGatewayResponse, err error) {
	response = NewRenewNatGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateNatGatewayRequest() (request *InquiryPriceCreateNatGatewayRequest) {
	request = &InquiryPriceCreateNatGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateNatGateway")

	return
}

func NewInquiryPriceCreateNatGatewayResponse() (response *InquiryPriceCreateNatGatewayResponse) {
	response = &InquiryPriceCreateNatGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateNatGateway 查询创建NAT网关的价格。
func (c *Client) InquiryPriceCreateNatGateway(request *InquiryPriceCreateNatGatewayRequest) (response *InquiryPriceCreateNatGatewayResponse, err error) {
	response = NewInquiryPriceCreateNatGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateSnatEntryRequest() (request *CreateSnatEntryRequest) {
	request = &CreateSnatEntryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateSnatEntry")

	return
}

func NewCreateSnatEntryResponse() (response *CreateSnatEntryResponse) {
	response = &CreateSnatEntryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateSnatEntry 创建SNAT规则。
func (c *Client) CreateSnatEntry(request *CreateSnatEntryRequest) (response *CreateSnatEntryResponse, err error) {
	response = NewCreateSnatEntryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifySnatEntryRequest() (request *ModifySnatEntryRequest) {
	request = &ModifySnatEntryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifySnatEntry")

	return
}

func NewModifySnatEntryResponse() (response *ModifySnatEntryResponse) {
	response = &ModifySnatEntryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifySnatEntry 修改SNAT规则。
func (c *Client) ModifySnatEntry(request *ModifySnatEntryRequest) (response *ModifySnatEntryResponse, err error) {
	response = NewModifySnatEntryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteSnatEntryRequest() (request *DeleteSnatEntryRequest) {
	request = &DeleteSnatEntryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteSnatEntry")

	return
}

func NewDeleteSnatEntryResponse() (response *DeleteSnatEntryResponse) {
	response = &DeleteSnatEntryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteSnatEntry 删除SNAT规则。
func (c *Client) DeleteSnatEntry(request *DeleteSnatEntryRequest) (response *DeleteSnatEntryResponse, err error) {
	response = NewDeleteSnatEntryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateDnatEntryRequest() (request *CreateDnatEntryRequest) {
	request = &CreateDnatEntryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateDnatEntry")

	return
}

func NewCreateDnatEntryResponse() (response *CreateDnatEntryResponse) {
	response = &CreateDnatEntryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateDnatEntry 创建DNAT规则。
func (c *Client) CreateDnatEntry(request *CreateDnatEntryRequest) (response *CreateDnatEntryResponse, err error) {
	response = NewCreateDnatEntryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyDnatEntryRequest() (request *ModifyDnatEntryRequest) {
	request = &ModifyDnatEntryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyDnatEntry")

	return
}

func NewModifyDnatEntryResponse() (response *ModifyDnatEntryResponse) {
	response = &ModifyDnatEntryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyDnatEntry 修改DNAT规则。
func (c *Client) ModifyDnatEntry(request *ModifyDnatEntryRequest) (response *ModifyDnatEntryResponse, err error) {
	response = NewModifyDnatEntryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteDnatEntryRequest() (request *DeleteDnatEntryRequest) {
	request = &DeleteDnatEntryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteDnatEntry")

	return
}

func NewDeleteDnatEntryResponse() (response *DeleteDnatEntryResponse) {
	response = &DeleteDnatEntryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteDnatEntry 删除DNAT规则。
func (c *Client) DeleteDnatEntry(request *DeleteDnatEntryRequest) (response *DeleteDnatEntryResponse, err error) {
	response = NewDeleteDnatEntryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAvailableBorderGatewayRequest() (request *DescribeAvailableBorderGatewayRequest) {
	request = &DescribeAvailableBorderGatewayRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAvailableBorderGateway")

	return
}

func NewDescribeAvailableBorderGatewayResponse() (response *DescribeAvailableBorderGatewayResponse) {
	response = &DescribeAvailableBorderGatewayResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeAvailableBorderGateway 获取可绑定NAT的边界网关。
func (c *Client) DescribeAvailableBorderGateway(request *DescribeAvailableBorderGatewayRequest) (response *DescribeAvailableBorderGatewayResponse, err error) {
	response = NewDescribeAvailableBorderGatewayResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
	request = &CreateRouteRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateRoute")

	return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
	response = &CreateRouteResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateRoute 创建一个自定义路由。
func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
	response = NewCreateRouteResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyRouteAttributeRequest() (request *ModifyRouteAttributeRequest) {
	request = &ModifyRouteAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyRouteAttribute")

	return
}

func NewModifyRouteAttributeResponse() (response *ModifyRouteAttributeResponse) {
	response = &ModifyRouteAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyRouteAttribute 修改路由的基本信息，目前只允许修改路由的名称。
func (c *Client) ModifyRouteAttribute(request *ModifyRouteAttributeRequest) (response *ModifyRouteAttributeResponse, err error) {
	response = NewModifyRouteAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
	request = &DeleteRouteRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteRoute")

	return
}

func NewDeleteRouteResponse() (response *DeleteRouteResponse) {
	response = &DeleteRouteResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteRoute 删除一条自定义路由。
func (c *Client) DeleteRoute(request *DeleteRouteRequest) (response *DeleteRouteResponse, err error) {
	response = NewDeleteRouteResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeRoutesRequest() (request *DescribeRoutesRequest) {
	request = &DescribeRoutesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeRoutes")

	return
}

func NewDescribeRoutesResponse() (response *DescribeRoutesResponse) {
	response = &DescribeRoutesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeRoutes 查询路由列表。用户可以根据ID、名称等信息来搜索Route信息。路由列表包括系统生成的以及用户创建的路由。
func (c *Client) DescribeRoutes(request *DescribeRoutesRequest) (response *DescribeRoutesResponse, err error) {
	response = NewDescribeRoutesResponse()
	err = c.ApiCall(request, response)
	return
}

