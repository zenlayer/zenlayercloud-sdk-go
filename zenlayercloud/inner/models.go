package inner

import "gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/common"

// CreateIpBlockEventRequest BMC / VOB IP 封堵事件请求
type CreateIpBlockEventRequest struct {
	*common.BaseRequest
	CID        string                 `json:"cid,omitempty"`         // 客户ID
	IP         string                 `json:"ip,omitempty"`          // 牵引IP
	Time       string                 `json:"time,omitempty"`        // 事件时间
	Operation  int 					  `json:"operation,omitempty"`   // 事件类型，1解封2封堵
	BPS        string                 `json:"bps,omitempty"`         // 牵引峰值
	ZoneID     int                    `json:"zoneId,omitempty"`      // 地区字段
	DivertMode string                 `json:"divertMode,omitempty"`  // 自动封堵还是手动封堵：automatically/manually
	GroupName  string                 `json:"groupName,omitempty"`   // 所属业务域分组名称
}

// CreateIpBlockEventResponse BMC / VOB IP 封堵事件响应
type CreateIpBlockEventResponse struct {
	*common.BaseResponse

	RequestId string `json:"requestId,omitempty"`

	Response struct {
		RequestId string `json:"requestId,omitempty"`
	} `json:"response"`
}