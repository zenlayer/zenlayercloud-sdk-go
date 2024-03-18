package zga

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2023-07-06"
	SERVICE    = "zga"
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

func NewDescribeOriginRegionsRequest() (request *DescribeOriginRegionsRequest) {
	request = &DescribeOriginRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeOriginRegions")

	return
}

// DescribeOriginRegions
// 调用本接口用于查询可用源站区域。
//
// Possible error codes to return:
func (c *Client) DescribeOriginRegions(request *DescribeOriginRegionsRequest) (response *DescribeOriginRegionsResponse, err error) {
	response = NewDescribeOriginRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeOriginRegionsResponse() (response *DescribeOriginRegionsResponse) {
	response = &DescribeOriginRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeAccelerateRegionsRequest() (request *DescribeAccelerateRegionsRequest) {
	request = &DescribeAccelerateRegionsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAccelerateRegions")

	return
}

// DescribeAccelerateRegions
// 调用本接口用于查询可用加速区域。
//
// Possible error codes to return:
func (c *Client) DescribeAccelerateRegions(request *DescribeAccelerateRegionsRequest) (response *DescribeAccelerateRegionsResponse, err error) {
	response = NewDescribeAccelerateRegionsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAccelerateRegionsResponse() (response *DescribeAccelerateRegionsResponse) {
	response = &DescribeAccelerateRegionsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
	request = &DescribeCertificatesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCertificates")

	return
}

// DescribeCertificates
// 调用本接口用于查询证书列表。
//
// Possible error codes to return:
func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
	response = NewDescribeCertificatesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
	response = &DescribeCertificatesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewCreateCertificateRequest() (request *CreateCertificateRequest) {
	request = &CreateCertificateRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateCertificate")

	return
}

// CreateCertificate
// 调用本接口用于创建证书。
//
// Possible error codes to return:
func (c *Client) CreateCertificate(request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
	response = NewCreateCertificateResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateCertificateResponse() (response *CreateCertificateResponse) {
	response = &CreateCertificateResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDeleteCertificateRequest() (request *DeleteCertificateRequest) {
	request = &DeleteCertificateRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteCertificate")

	return
}

// DeleteCertificate
// 调用本接口用于删除证书。
//
// Possible error codes to return:
func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
	response = NewDeleteCertificateResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteCertificateResponse() (response *DeleteCertificateResponse) {
	response = &DeleteCertificateResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeAcceleratorsRequest() (request *DescribeAcceleratorsRequest) {
	request = &DescribeAcceleratorsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAccelerators")

	return
}

// DescribeAccelerators
// 调用本接口用于查询多个加速器的信息。用户可以根据ID、名称或者CNAME等条件来查询。
//
// Possible error codes to return:
func (c *Client) DescribeAccelerators(request *DescribeAcceleratorsRequest) (response *DescribeAcceleratorsResponse, err error) {
	response = NewDescribeAcceleratorsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAcceleratorsResponse() (response *DescribeAcceleratorsResponse) {
	response = &DescribeAcceleratorsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewCreateAcceleratorRequest() (request *CreateAcceleratorRequest) {
	request = &CreateAcceleratorRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateAccelerator")

	return
}

// CreateAccelerator
// 调用本接口用于创建加速器。
//
// Possible error codes to return:
func (c *Client) CreateAccelerator(request *CreateAcceleratorRequest) (response *CreateAcceleratorResponse, err error) {
	response = NewCreateAcceleratorResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCreateAcceleratorResponse() (response *CreateAcceleratorResponse) {
	response = &CreateAcceleratorResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDeleteAcceleratorRequest() (request *DeleteAcceleratorRequest) {
	request = &DeleteAcceleratorRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DeleteAccelerator")

	return
}

// DeleteAccelerator
// 调用本接口用于删除加速器。
//
// Possible error codes to return:
func (c *Client) DeleteAccelerator(request *DeleteAcceleratorRequest) (response *DeleteAcceleratorResponse, err error) {
	response = NewDeleteAcceleratorResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDeleteAcceleratorResponse() (response *DeleteAcceleratorResponse) {
	response = &DeleteAcceleratorResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewStartAcceleratorRequest() (request *StartAcceleratorRequest) {
	request = &StartAcceleratorRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StartAccelerator")

	return
}

// StartAccelerator
// 调用本接口用于启动加速器。
//
// Possible error codes to return:
func (c *Client) StartAccelerator(request *StartAcceleratorRequest) (response *StartAcceleratorResponse, err error) {
	response = NewStartAcceleratorResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStartAcceleratorResponse() (response *StartAcceleratorResponse) {
	response = &StartAcceleratorResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewRedeployAcceleratorRequest() (request *RedeployAcceleratorRequest) {
	request = &RedeployAcceleratorRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RedeployAccelerator")

	return
}

// RedeployAccelerator
// 调用本接口用于重新部署加速器。
//
// Possible error codes to return:
func (c *Client) RedeployAccelerator(request *RedeployAcceleratorRequest) (response *RedeployAcceleratorResponse, err error) {
	response = NewRedeployAcceleratorResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRedeployAcceleratorResponse() (response *RedeployAcceleratorResponse) {
	response = &RedeployAcceleratorResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewPauseAcceleratorRequest() (request *PauseAcceleratorRequest) {
	request = &PauseAcceleratorRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "PauseAccelerator")

	return
}

// PauseAccelerator
// 调用本接口用于重新部署加速器。
//
// Possible error codes to return:
func (c *Client) PauseAccelerator(request *PauseAcceleratorRequest) (response *PauseAcceleratorResponse, err error) {
	response = NewPauseAcceleratorResponse()
	err = c.ApiCall(request, response)
	return
}

func NewPauseAcceleratorResponse() (response *PauseAcceleratorResponse) {
	response = &PauseAcceleratorResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorDomainRequest() (request *ModifyAcceleratorDomainRequest) {
	request = &ModifyAcceleratorDomainRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorDomain")

	return
}

// ModifyAcceleratorDomain
// 调用本接口用于修改加速器加速域名。
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorDomain(request *ModifyAcceleratorDomainRequest) (response *ModifyAcceleratorDomainResponse, err error) {
	response = NewModifyAcceleratorDomainResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorDomainResponse() (response *ModifyAcceleratorDomainResponse) {
	response = &ModifyAcceleratorDomainResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorNameRequest() (request *ModifyAcceleratorNameRequest) {
	request = &ModifyAcceleratorNameRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorName")

	return
}

// ModifyAcceleratorName
// 调用本接口用于修改加速器名字。
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorName(request *ModifyAcceleratorNameRequest) (response *ModifyAcceleratorNameResponse, err error) {
	response = NewModifyAcceleratorNameResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorNameResponse() (response *ModifyAcceleratorNameResponse) {
	response = &ModifyAcceleratorNameResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorOriginRequest() (request *ModifyAcceleratorOriginRequest) {
	request = &ModifyAcceleratorOriginRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorOrigin")

	return
}

// ModifyAcceleratorOrigin
// 调用本接口用于修改加速器源站信息。
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorOrigin(request *ModifyAcceleratorOriginRequest) (response *ModifyAcceleratorOriginResponse, err error) {
	response = NewModifyAcceleratorOriginResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorOriginResponse() (response *ModifyAcceleratorOriginResponse) {
	response = &ModifyAcceleratorOriginResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorAccRegionRequest() (request *ModifyAcceleratorAccRegionRequest) {
	request = &ModifyAcceleratorAccRegionRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorAccRegion")

	return
}

// ModifyAcceleratorAccRegion
// 调用本接口用于修改加速器加速区域。
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorAccRegion(request *ModifyAcceleratorAccRegionRequest) (response *ModifyAcceleratorAccRegionResponse, err error) {
	response = NewModifyAcceleratorAccRegionResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorAccRegionResponse() (response *ModifyAcceleratorAccRegionResponse) {
	response = &ModifyAcceleratorAccRegionResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorRuleRequest() (request *ModifyAcceleratorRuleRequest) {
	request = &ModifyAcceleratorRuleRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorRule")

	return
}

// ModifyAcceleratorRule
// 调用本接口用于修改加速器加速规则。
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorRule(request *ModifyAcceleratorRuleRequest) (response *ModifyAcceleratorRuleResponse, err error) {
	response = NewModifyAcceleratorRuleResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorRuleResponse() (response *ModifyAcceleratorRuleResponse) {
	response = &ModifyAcceleratorRuleResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorProtocolOptsRequest() (request *ModifyAcceleratorProtocolOptsRequest) {
	request = &ModifyAcceleratorProtocolOptsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorProtocolOpts")

	return
}

// ModifyAcceleratorProtocolOpts
// 调用本接口用于修改加速器加速规则协议可选配置。
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorProtocolOpts(request *ModifyAcceleratorProtocolOptsRequest) (response *ModifyAcceleratorProtocolOptsResponse, err error) {
	response = NewModifyAcceleratorProtocolOptsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorProtocolOptsResponse() (response *ModifyAcceleratorProtocolOptsResponse) {
	response = &ModifyAcceleratorProtocolOptsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorCertificateRequest() (request *ModifyAcceleratorCertificateRequest) {
	request = &ModifyAcceleratorCertificateRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorCertificate")

	return
}

// ModifyAcceleratorCertificate
// 调用本接口用于修改加速器证书。
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorCertificate(request *ModifyAcceleratorCertificateRequest) (response *ModifyAcceleratorCertificateResponse, err error) {
	response = NewModifyAcceleratorCertificateResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorCertificateResponse() (response *ModifyAcceleratorCertificateResponse) {
	response = &ModifyAcceleratorCertificateResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorHealthCheckRequest() (request *ModifyAcceleratorHealthCheckRequest) {
	request = &ModifyAcceleratorHealthCheckRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorHealthCheck")

	return
}

// ModifyAcceleratorHealthCheck
// 调用本接口用于修改加速器健康检查。
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorHealthCheck(request *ModifyAcceleratorHealthCheckRequest) (response *ModifyAcceleratorHealthCheckResponse, err error) {
	response = NewModifyAcceleratorHealthCheckResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorHealthCheckResponse() (response *ModifyAcceleratorHealthCheckResponse) {
	response = &ModifyAcceleratorHealthCheckResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewModifyAcceleratorAccessControlRequest() (request *ModifyAcceleratorAccessControlRequest) {
	request = &ModifyAcceleratorAccessControlRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyAcceleratorAccessControl")

	return
}

// ModifyAcceleratorAccessControl
// 调用本接口用于配置加速器访问控制
//
// Possible error codes to return:
func (c *Client) ModifyAcceleratorAccessControl(request *ModifyAcceleratorAccessControlRequest) (response *ModifyAcceleratorAccessControlResponse, err error) {
	response = NewModifyAcceleratorAccessControlResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyAcceleratorAccessControlResponse() (response *ModifyAcceleratorAccessControlResponse) {
	response = &ModifyAcceleratorAccessControlResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewOpenAcceleratorAccessControlRequest() (request *OpenAcceleratorAccessControlRequest) {
	request = &OpenAcceleratorAccessControlRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "OpenAcceleratorAccessControl")

	return
}

// OpenAcceleratorAccessControl
// 调用本接口用于开启加速器访问控制。
//
// Possible error codes to return:
func (c *Client) OpenAcceleratorAccessControl(request *OpenAcceleratorAccessControlRequest) (response *OpenAcceleratorAccessControlResponse, err error) {
	response = NewOpenAcceleratorAccessControlResponse()
	err = c.ApiCall(request, response)
	return
}

func NewOpenAcceleratorAccessControlResponse() (response *OpenAcceleratorAccessControlResponse) {
	response = &OpenAcceleratorAccessControlResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewCloseAcceleratorAccessControlRequest() (request *CloseAcceleratorAccessControlRequest) {
	request = &CloseAcceleratorAccessControlRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CloseAcceleratorAccessControl")

	return
}

// CloseAcceleratorAccessControl
// 调用本接口用于关闭加速器访问控制。
//
// Possible error codes to return:
func (c *Client) CloseAcceleratorAccessControl(request *CloseAcceleratorAccessControlRequest) (response *CloseAcceleratorAccessControlResponse, err error) {
	response = NewCloseAcceleratorAccessControlResponse()
	err = c.ApiCall(request, response)
	return
}

func NewCloseAcceleratorAccessControlResponse() (response *CloseAcceleratorAccessControlResponse) {
	response = &CloseAcceleratorAccessControlResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeAcceleratorsAlertsRequest() (request *DescribeAcceleratorsAlertsRequest) {
	request = &DescribeAcceleratorsAlertsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAcceleratorsAlerts")
	return
}

// DescribeAcceleratorsAlerts
// 调用本接口用于查询加速器警报。
//
//	 Possible error codes to return:
//		INVALID_FROM_TIME_MUST_BEFORE_TO_TIME = "Invalid.From.Time.Must.Before.To.Time"
func (c *Client) DescribeAcceleratorsAlerts(request *DescribeAcceleratorsAlertsRequest) (response *DescribeAcceleratorsAlertsResponse, err error) {
	response = NewDescribeAcceleratorsAlertsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAcceleratorsAlertsResponse() (response *DescribeAcceleratorsAlertsResponse) {
	response = &DescribeAcceleratorsAlertsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeAcceleratorLogsRequest() (request *DescribeAcceleratorLogsRequest) {
	request = &DescribeAcceleratorLogsRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAcceleratorLogs")
	return
}

// DescribeAcceleratorLogs
// 调用本接口用于查询加速器代理日志。
//
//	 Possible error codes to return:
//		NOT_EXIST_VIRTUAL_IP = "Not.Exist.Virtual.Ip"
//		INVALID_START_TIME_MUST_BEFORE_END_TIME = "Invalid.Start.Time.Must.Before.End.Time"
func (c *Client) DescribeAcceleratorLogs(request *DescribeAcceleratorLogsRequest) (response *DescribeAcceleratorLogsResponse, err error) {
	response = NewDescribeAcceleratorLogsResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAcceleratorLogsResponse() (response *DescribeAcceleratorLogsResponse) {
	response = &DescribeAcceleratorLogsResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeAcceleratorTrafficRequest() (request *DescribeAcceleratorTrafficRequest) {
	request = &DescribeAcceleratorTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAcceleratorTraffic")
	return
}

// DescribeAcceleratorTraffic
// 调用本接口用于查询加速器带宽流量。
//
//	Possible error codes to return:
//		INVALID_START_TIME_MUST_BEFORE_END_TIME = "Invalid.Start.Time.Must.Before.End.Time"
//		INVALID_REGION_NOT_FOUND = "Invalid.Region.Not.Found"
func (c *Client) DescribeAcceleratorTraffic(request *DescribeAcceleratorTrafficRequest) (response *DescribeAcceleratorTrafficResponse, err error) {
	response = NewDescribeAcceleratorTrafficResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAcceleratorTrafficResponse() (response *DescribeAcceleratorTrafficResponse) {
	response = &DescribeAcceleratorTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}
