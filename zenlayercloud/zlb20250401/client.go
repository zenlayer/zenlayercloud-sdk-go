/*
 * Zenlayer.com Inc.
 * Copyright (c) 2014-2022 All Rights Reserved.
 */
package zlb

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2025-04-01"
	SERVICE    = "zlb"
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


func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
	request = &DescribeListenersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeListeners")

	return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
	response = &DescribeListenersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeListeners 本接口接口可根据负载均衡器 ID、监听器的协议作为过滤条件获取监听器列表。如果不指定任何过滤条件，则返回指定负载均衡实例下的所有监听器。
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
	response = NewDescribeListenersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
	request = &CreateListenerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateListener")

	return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
	response = &CreateListenerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateListener 查询创建负载均衡监听器。
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
	response = NewCreateListenerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
	request = &DeleteListenerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteListener")

	return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
	response = &DeleteListenerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeleteListener 查询删除一个负载均衡监听器。
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
	response = NewDeleteListenerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyListenerRequest() (request *ModifyListenerRequest) {
	request = &ModifyListenerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyListener")

	return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
	response = &ModifyListenerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyListener 修改负载均衡监听器的属性，包括监听器的名称、健康检查参数、转发方式等。不支持修改监听器的监听协议。
func (c *Client) ModifyListener(request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
	response = NewModifyListenerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeLoadBalancerRegionsRequest() (request *DescribeLoadBalancerRegionsRequest) {
	request = &DescribeLoadBalancerRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeLoadBalancerRegions")

	return
}

func NewDescribeLoadBalancerRegionsResponse() (response *DescribeLoadBalancerRegionsResponse) {
	response = &DescribeLoadBalancerRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeLoadBalancerRegions 查询支持购买负载均衡的区域。
func (c *Client) DescribeLoadBalancerRegions(request *DescribeLoadBalancerRegionsRequest) (response *DescribeLoadBalancerRegionsResponse, err error) {
	response = NewDescribeLoadBalancerRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRegisterBackendRequest() (request *RegisterBackendRequest) {
	request = &RegisterBackendRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RegisterBackend")

	return
}

func NewRegisterBackendResponse() (response *RegisterBackendResponse) {
	response = &RegisterBackendResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RegisterBackend 将一台或多台后端服务绑定到负载均衡的监听器。
func (c *Client) RegisterBackend(request *RegisterBackendRequest) (response *RegisterBackendResponse, err error) {
	response = NewRegisterBackendResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeregisterBackendRequest() (request *DeregisterBackendRequest) {
	request = &DeregisterBackendRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeregisterBackend")

	return
}

func NewDeregisterBackendResponse() (response *DeregisterBackendResponse) {
	response = &DeregisterBackendResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DeregisterBackend 将一台或多台绑定在指定监听器上的后端服务解绑。
func (c *Client) DeregisterBackend(request *DeregisterBackendRequest) (response *DeregisterBackendResponse, err error) {
	response = NewDeregisterBackendResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyBackendRequest() (request *ModifyBackendRequest) {
	request = &ModifyBackendRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyBackend")

	return
}

func NewModifyBackendResponse() (response *ModifyBackendResponse) {
	response = &ModifyBackendResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyBackend 修改一台或多台绑定在指定监听器上的后端服务的配置，包括权重和和后端服务器转发端口。
func (c *Client) ModifyBackend(request *ModifyBackendRequest) (response *ModifyBackendResponse, err error) {
	response = NewModifyBackendResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBackendsRequest() (request *DescribeBackendsRequest) {
	request = &DescribeBackendsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBackends")

	return
}

func NewDescribeBackendsResponse() (response *DescribeBackendsResponse) {
	response = &DescribeBackendsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBackends 查询负载均衡实例的绑定的后端服务列表。
func (c *Client) DescribeBackends(request *DescribeBackendsRequest) (response *DescribeBackendsResponse, err error) {
	response = NewDescribeBackendsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyLoadBalancersAttributeRequest() (request *ModifyLoadBalancersAttributeRequest) {
	request = &ModifyLoadBalancersAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyLoadBalancersAttribute")

	return
}

func NewModifyLoadBalancersAttributeResponse() (response *ModifyLoadBalancersAttributeResponse) {
	response = &ModifyLoadBalancersAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// ModifyLoadBalancersAttribute 修改负载均衡实例的属性。目前仅支持修改负载均衡实例的名称。
func (c *Client) ModifyLoadBalancersAttribute(request *ModifyLoadBalancersAttributeRequest) (response *ModifyLoadBalancersAttributeResponse, err error) {
	response = NewModifyLoadBalancersAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
	request = &DescribeLoadBalancersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeLoadBalancers")

	return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
	response = &DescribeLoadBalancersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeLoadBalancers 查询负载均衡实例列表。
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
	response = NewDescribeLoadBalancersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRestoreLoadBalancerRequest() (request *RestoreLoadBalancerRequest) {
	request = &RestoreLoadBalancerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RestoreLoadBalancer")

	return
}

func NewRestoreLoadBalancerResponse() (response *RestoreLoadBalancerResponse) {
	response = &RestoreLoadBalancerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// RestoreLoadBalancer 将在回收站的负载均衡实例进行恢复。
func (c *Client) RestoreLoadBalancer(request *RestoreLoadBalancerRequest) (response *RestoreLoadBalancerResponse, err error) {
	response = NewRestoreLoadBalancerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminateLoadBalancerRequest() (request *TerminateLoadBalancerRequest) {
	request = &TerminateLoadBalancerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminateLoadBalancer")

	return
}

func NewTerminateLoadBalancerResponse() (response *TerminateLoadBalancerResponse) {
	response = &TerminateLoadBalancerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// TerminateLoadBalancer 本接口用户删除一个指定的负载均衡器实例。
func (c *Client) TerminateLoadBalancer(request *TerminateLoadBalancerRequest) (response *TerminateLoadBalancerResponse, err error) {
	response = NewTerminateLoadBalancerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateLoadBalancer")

	return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// CreateLoadBalancer 创建负载均衡器实例。
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = NewCreateLoadBalancerResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateLoadBalancerRequest() (request *InquiryPriceCreateLoadBalancerRequest) {
	request = &InquiryPriceCreateLoadBalancerRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateLoadBalancer")

	return
}

func NewInquiryPriceCreateLoadBalancerResponse() (response *InquiryPriceCreateLoadBalancerResponse) {
	response = &InquiryPriceCreateLoadBalancerResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// InquiryPriceCreateLoadBalancer 查询创建负载均衡的价格。
func (c *Client) InquiryPriceCreateLoadBalancer(request *InquiryPriceCreateLoadBalancerRequest) (response *InquiryPriceCreateLoadBalancerResponse, err error) {
	response = NewInquiryPriceCreateLoadBalancerResponse()
	err = c.ApiCall(request, response)
	return
}

