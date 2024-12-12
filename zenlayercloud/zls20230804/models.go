package zls

import "gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeLogsRequest struct {
	*common.BaseRequest

	StartTime string `json:"startTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`

	ResUid string `json:"resUid,omitempty"`

	ResEvent string `json:"resEvent,omitempty"`

	ClientIP string `json:"clientIP,omitempty"`

	Size int `json:"size,omitempty"`

	Cursor []*interface{} `json:"cursor,omitempty"`
}

type DescribeLogsResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeLogsResponseParam `json:"response"`
}

type DescribeLogsResponseParam struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Cursor []*interface{} `json:"cursor,omitempty"`

	DataSet []*LogInfo `json:"dataSet,omitempty"`
}

type LogInfo struct {
	ResUid string `json:"resUid,omitempty"`

	ResType string `json:"resType,omitempty"`

	ResEvent string `json:"resEvent,omitempty"`

	OpsTime string `json:"opsTime,omitempty"`

	EventSource string `json:"eventSource,omitempty"`

	ApiVersion string `json:"apiVersion,omitempty"`

	OpsUser string `json:"opsUser,omitempty"`

	ClientIP string `json:"clientIP,omitempty"`

	Request string `json:"request,omitempty"`

	Response string `json:"response,omitempty"`
}
