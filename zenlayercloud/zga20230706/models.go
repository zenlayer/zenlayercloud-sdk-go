package zga

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeOriginRegionsRequest struct {
	*common.BaseRequest
}

type DescribeOriginRegionsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeOriginRegionsResponseParams `json:"response"`
}

type DescribeOriginRegionsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// 区域信息集合。
	RegionSet []Region `json:"regionSet,omitempty"`
}

type Region struct {
	// 可用区ID 例如：SEL
	RegionId string `json:"regionId,omitempty"`

	// 可用区的名称。
	RegionName string `json:"regionName,omitempty"`
}

type DescribeAccelerateRegionsRequest struct {
	*common.BaseRequest

	// 服务器所在区域ID
	OriginRegionId string `json:"originRegionId,omitempty"`
}

type DescribeAccelerateRegionsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAccelerateRegionsResponseParams `json:"response"`
}

type DescribeAccelerateRegionsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// 区域信息集合。
	RegionSet []*Region `json:"regionSet,omitempty"`
}

type DescribeCertificatesRequest struct {
	*common.BaseRequest

	CertificateIds   []string `json:"certificateIds,omitempty"`
	CertificateLabel string   `json:"certificateLabel,omitempty"`
	San              string   `json:"san,omitempty"`
	ResourceGroupId  string   `json:"resourceGroupId,omitempty"`
	Expired          *bool    `json:"expired,omitempty"`
	PageSize         int      `json:"pageSize,omitempty"`
	PageNum          int      `json:"pageNum,omitempty"`
}

type DescribeCertificatesResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeCertificatesResponseParams `json:"response"`
}

type DescribeCertificatesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	TotalCount int `json:"totalCount,omitempty"`

	DataSet []*CertificateInfo `json:"dataSet,omitempty"`
}

type CertificateInfo struct {
	CertificateId    string   `json:"certificateId,omitempty"`
	CertificateLabel string   `json:"certificateLabel,omitempty"`
	Common           string   `json:"common,omitempty"`
	Fingerprint      string   `json:"fingerprint,omitempty"`
	Issuer           string   `json:"issuer,omitempty"`
	Sans             []string `json:"sans,omitempty"`
	Algorithm        string   `json:"algorithm,omitempty"`
	CreateTime       string   `json:"createTime,omitempty"`
	StartTime        string   `json:"startTime,omitempty"`
	EndTime          string   `json:"endTime,omitempty"`
	Expired          bool     `json:"expired,omitempty"`
	ResourceGroupId  string   `json:"resourceGroupId,omitempty"`
}

type CreateCertificateRequest struct {
	*common.BaseRequest

	CertificateContent string `json:"certificateContent,omitempty"`
	CertificateKey     string `json:"certificateKey,omitempty"`
	CertificateLabel   string `json:"certificateLabel,omitempty"`
	ResourceGroupId    string `json:"resourceGroupId,omitempty"`
}

type CreateCertificateResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateCertificateResponseParams `json:"response"`
}

type CreateCertificateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	CertificateId string `json:"certificateId,omitempty"`
}

type DeleteCertificateRequest struct {
	*common.BaseRequest

	CertificateId string `json:"certificateId,omitempty"`
}

type DeleteCertificateResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type DescribeAcceleratorsRequest struct {
	*common.BaseRequest

	AcceleratorIds     []string `json:"acceleratorIds,omitempty"`
	AcceleratorName    string   `json:"acceleratorName,omitempty"`
	AcceleratorStatus  string   `json:"acceleratorStatus,omitempty"`
	AccelerateRegionId string   `json:"accelerateRegionId,omitempty"`
	Vip                string   `json:"vip,omitempty"`
	Domain             string   `json:"domain,omitempty"`
	Origin             string   `json:"origin,omitempty"`
	OriginRegionId     string   `json:"originRegionId,omitempty"`
	Cname              string   `json:"cname,omitempty"`
	ResourceGroupId    string   `json:"resourceGroupId,omitempty"`
	PageSize           int      `json:"pageSize,omitempty"`
	PageNum            int      `json:"pageNum,omitempty"`
}

type DescribeAcceleratorsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeAcceleratorsResponseParams `json:"response"`
}

type DescribeAcceleratorsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	TotalCount int `json:"totalCount,omitempty"`

	DataSet []*AcceleratorInfo `json:"dataSet,omitempty"`
}

type AcceleratorInfo struct {
	AcceleratorId     string                        `json:"acceleratorId,omitempty"`
	AcceleratorType   string                        `json:"acceleratorType,omitempty"`
	AcceleratorName   string                        `json:"acceleratorName,omitempty"`
	ChargeType        string                        `json:"chargeType,omitempty"`
	Domain            *Domain                       `json:"domain,omitempty"`
	AcceleratorStatus string                        `json:"acceleratorStatus,omitempty"`
	Cname             string                        `json:"cname,omitempty"`
	Origin            *OriginInfo                   `json:"origin,omitempty"`
	AccelerateRegions []*AccelerateRegionInfo       `json:"accelerateRegions,omitempty"`
	L4Listeners       []*AccelerationRuleL4Listener `json:"l4Listeners,omitempty"`
	L7Listeners       []*AccelerationRuleL7Listener `json:"l7Listeners,omitempty"`
	ProtocolOpts      *AccelerationRuleProtocolOpts `json:"protocolOpts,omitempty"`
	Certificate       *CertificateInfo              `json:"certificate,omitempty"`
	AccessControl     *AccessControl                `json:"accessControl,omitempty"`
	HealthCheck       *HealthCheck                  `json:"healthCheck,omitempty"`
	CreateTime        string                        `json:"createTime,omitempty"`
	ResourceGroupId   string                        `json:"resourceGroupId,omitempty"`
}

type Domain struct {
	Domain        string `json:"domain,omitempty"`
	RelateDomains string `json:"relateDomains,omitempty"`
}

type Origin struct {
	OriginRegionId string `json:"originRegionId,omitempty"`
	Origin         string `json:"origin,omitempty"`
	BackupOrigin   string `json:"backupOrigin,omitempty"`
}

type OriginInfo struct {
	OriginRegionId   string `json:"originRegionId,omitempty"`
	OriginRegionName string `json:"originRegionName,omitempty"`
	Origin           string `json:"origin,omitempty"`
	BackupOrigin     string `json:"backupOrigin,omitempty"`
}

type AccelerateRegionInfo struct {
	AccelerateRegionId     string `json:"accelerateRegionId,omitempty"`
	AccelerateRegionName   string `json:"accelerateRegionName,omitempty"`
	AccelerateRegionStatus string `json:"accelerateRegionStatus,omitempty"`
	Vip                    string `json:"vip,omitempty"`
	Bandwidth              int    `json:"bandwidth,omitempty"`
}

type AccelerationRuleL4Listener struct {
	Port          int    `json:"port,omitempty"`
	BackPort      int    `json:"backPort,omitempty"`
	PortRange     string `json:"portRange,omitempty"`
	BackPortRange string `json:"backPortRange,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

type AccelerationRuleL7Listener struct {
	Port          int    `json:"port,omitempty"`
	BackPort      int    `json:"backPort,omitempty"`
	PortRange     string `json:"portRange,omitempty"`
	BackPortRange string `json:"backPortRange,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	BackProtocol  string `json:"backProtocol,omitempty"`
	Host          string `json:"host,omitempty"`
}

type AccelerationRuleProtocolOpts struct {
	Toa           *bool `json:"toa,omitempty"`
	ToaValue      int   `json:"toaValue,omitempty"`
	Websocket     *bool `json:"websocket,omitempty"`
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
	Gzip          *bool `json:"gzip,omitempty"`
}

type AccessControl struct {
	Enable bool                 `json:"enable,omitempty"`
	Rules  []*AccessControlRule `json:"rules,omitempty"`
}

type AccessControlRule struct {
	Listener  string `json:"listener,omitempty"`
	Directory string `json:"directory,omitempty"`
	CidrIp    string `json:"cidrIp,omitempty"`
	Policy    string `json:"policy,omitempty"`
	Note      string `json:"note,omitempty"`
}

type HealthCheck struct {
	Enable bool `json:"enable,omitempty"`
	Alarm  bool `json:"alarm,omitempty"`
	Port   int  `json:"port,omitempty"`
}

type AccelerateRegion struct {
	AccelerateRegionId string `json:"accelerateRegionId,omitempty"`
	Bandwidth          int    `json:"bandwidth,omitempty"`
	Vip                string `json:"vip,omitempty"`
}

type CreateAcceleratorRequest struct {
	*common.BaseRequest

	AcceleratorName   string                        `json:"acceleratorName,omitempty"`
	ChargeType        string                        `json:"chargeType,omitempty"`
	ResourceGroupId   string                        `json:"resourceGroupId,omitempty"`
	CertificateId     string                        `json:"certificateId,omitempty"`
	Domain            *Domain                       `json:"domain,omitempty"`
	Origin            Origin                        `json:"origin,omitempty"`
	AccelerateRegions []AccelerateRegion            `json:"accelerateRegions,omitempty"`
	L4Listeners       []*AccelerationRuleL4Listener `json:"l4Listeners,omitempty"`
	L7Listeners       []*AccelerationRuleL7Listener `json:"l7Listeners,omitempty"`
	ProtocolOpts      *AccelerationRuleProtocolOpts `json:"protocolOpts,omitempty"`
	HealthCheck       *HealthCheck                  `json:"healthCheck,omitempty"`
}

type CreateAcceleratorResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *CreateAcceleratorResponseParams `json:"response"`
}

type CreateAcceleratorResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	AcceleratorId string `json:"acceleratorId,omitempty"`
}

type DeleteAcceleratorRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`
}

type DeleteAcceleratorResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type StartAcceleratorRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`
}

type StartAcceleratorResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type RedeployAcceleratorRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`
}

type RedeployAcceleratorResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type PauseAcceleratorRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`
}

type PauseAcceleratorResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorDomainRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	Domain Domain `json:"domain,omitempty"`
}

type ModifyAcceleratorDomainResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorNameRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	AcceleratorName string `json:"acceleratorName,omitempty"`
}

type ModifyAcceleratorNameResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorOriginRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	Origin Origin `json:"origin,omitempty"`
}

type ModifyAcceleratorOriginResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorAccRegionRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	AccelerateRegions []AccelerateRegion `json:"accelerateRegions,omitempty"`
}

type ModifyAcceleratorAccRegionResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorRuleRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	L4Listeners []*AccelerationRuleL4Listener `json:"l4Listeners,omitempty"`

	L7Listeners []*AccelerationRuleL7Listener `json:"l7Listeners,omitempty"`
}

type ModifyAcceleratorRuleResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorProtocolOptsRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	ProtocolOpts AccelerationRuleProtocolOpts `json:"protocolOpts,omitempty"`
}

type ModifyAcceleratorProtocolOptsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorCertificateRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	CertificateId string `json:"certificateId,omitempty"`
}

type ModifyAcceleratorCertificateResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorHealthCheckRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	HealthCheck HealthCheck `json:"healthCheck,omitempty"`
}

type ModifyAcceleratorHealthCheckResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type ModifyAcceleratorAccessControlRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`

	AccessControlRules []AccessControlRule `json:"accessControlRules,omitempty"`
}

type ModifyAcceleratorAccessControlResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type OpenAcceleratorAccessControlRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`
}

type OpenAcceleratorAccessControlResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}

type CloseAcceleratorAccessControlRequest struct {
	*common.BaseRequest

	AcceleratorId string `json:"acceleratorId,omitempty"`
}

type CloseAcceleratorAccessControlResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}
