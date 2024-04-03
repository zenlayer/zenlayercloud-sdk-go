package zec

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeInstanceMonitorDataResponse struct {
	*common.BaseResponse

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 requestId。
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeInstanceMonitorDataResponseParam `json:"response"`
}

type DescribeInstanceMonitorDataResponseParam struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 requestId。
	RequestId string `json:"requestId,omitempty"`

	// 指标数据的最大值，单位取决于指标类型。
	MaxValue float64 `json:"maxValue,omitempty"`

	// 指标数据的平均值，单位取决于指标类型。
	AvgValue float64 `json:"avgValue,omitempty"`

	// 指标数据列表。
	DataList []*MetricValue `json:"dataList,omitempty"`
}

type MetricValue struct {
	// 监控指标值
	Value float64 `json:"value,omitempty"`

	// 数据时间戳，单位秒
	TimeInSecond int `json:"timeInSecond,omitempty"`
}

type DescribeInstanceMonitorDataRequest struct {
	*common.BaseRequest

	//实例ID。
	InstanceId string `json:"instanceId,omitempty"`

	//指标类型。
	//INTERNET_INGRESS_BITS: 公网入向带宽，单位bps
	//INTERNET_EGRESS_BITS: 公网出向带宽，单位bps
	MetricType string `json:"metricType,omitempty"`

	//查询开始时间。
	//按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	StartTime string `json:"startTime,omitempty"`

	//查询结束时间。
	//按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime string `json:"endTime,omitempty"`
}
