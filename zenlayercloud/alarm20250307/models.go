package alarm

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeIpBlockEventsRequest struct {
	*common.BaseRequest

	StartTime string `json:"startTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`

	Ip string `json:"ip,omitempty"`
}

type DescribeIpBlockEventsResponse struct {
	*common.BaseResponse

	RequestId string `json:"requestId,omitempty"`

	Response *DescribeIpBlockEventsResponseParams `json:"response"`
}

type DescribeIpBlockEventsResponseParams struct {
	RequestId string `json:"requestId,omitempty"`

	DataList []*IpBlockEventInfo `json:"dataList,omitempty"`
}

type IpBlockEventInfo struct {
	Ip            string `json:"ip,omitempty"`
	InternalIps   string `json:"internalIps,omitempty"`
	Region        string `json:"region,omitempty"`
	ResourceGroup string `json:"resourceGroup,omitempty"`
	BlockTime     string `json:"blockTime,omitempty"`
	UnblockTime   string `json:"unblockTime,omitempty"`
	DivertMode    string `json:"divertMode,omitempty"`
}
