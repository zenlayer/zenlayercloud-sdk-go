package maintenance

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeMaintenanceAlertsRequest struct {
	*common.BaseRequest

	Product string `json:"product,omitempty"`
}

type DescribeMaintenanceAlertsResponse struct {
	*common.BaseResponse

	RequestId string `json:"requestId,omitempty"`

	Response *DescribeMaintenanceAlertsResponseParams `json:"response"`
}

type DescribeMaintenanceAlertsResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	DataList []*MaintenanceAlertInfo `json:"dataList,omitempty"`
}

type MaintenanceAlertInfo struct {
	InstanceId string `json:"instanceId,omitempty"`
	Location   string `json:"location,omitempty"`
	Impact     string `json:"impact,omitempty"`
	StartTime  string `json:"startTime,omitempty"`
	EndTime    string `json:"endTime,omitempty"`
	Status     string `json:"status,omitempty"`
}
