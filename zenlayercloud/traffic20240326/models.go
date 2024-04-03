package traffic

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeBandwidthClustersRequest struct {
	*common.BaseRequest
	// BandwidthCluster IDs.
	// You can query up to 100 bandwidthClusters in each request.
	BandwidthClusterIds []string `json:"bandwidthClusterIds,omitempty"`

	// BandwidthCluster name.
	BandwidthClusterName string `json:"bandwidthClusterName,omitempty"`

	// City name to which the bandwidthClusters belong.
	CityName string `json:"cityName,omitempty"`

	// Number of pages returned.
	// Default value: 1
	PageNum int `json:"pageNum,omitempty"`

	// Number of items in the current page result.
	// Default value: 20
	// Maximum value: 1000
	PageSize int `json:"pageSize,omitempty"`
}

type BandwidthClusterInfo struct {
	// BandwidthCluster ID.
	BandwidthClusterId string `json:"bandwidthClusterId,omitempty"`

	// BandwidthCluster name to be displayed.
	BandwidthClusterName string `json:"bandwidthClusterName,omitempty"`

	// Location to which the bandwidthCluster belong.
	Location string `json:"location,omitempty"`

	// Creation time.
	// Use UTC time according to the ISO8601 standard. Format: YYYY-MM-DDThh:mm:ssZ.
	CreateTime string `json:"createTime,omitempty"`
}

type DescribeBandwidthClustersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	// Number of bandwidth clusters meeting the filtering conditions.
	TotalCount int `json:"totalCount,omitempty"`

	// Information on a bandwidth cluster.
	DataSet []*BandwidthClusterInfo `json:"dataSet,omitempty"`
}

type DescribeBandwidthClustersResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeBandwidthClustersResponseParams `json:"response"`
}

type DescribeBandwidthClusterTrafficRequest struct {
	*common.BaseRequest

	BandwidthClusterId string `json:"bandwidthClusterId,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`
}

type DescribeBandwidthClusterTrafficResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *BandwidthClusterTrafficDataResponse `json:"response"`
}

type BandwidthClusterTrafficDataResponse struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	DataList []*BandwidthClusterTrafficData `json:"dataList,omitempty"`

	In95 int64 `json:"in95,omitempty"`

	In95Time string `json:"in95Time,omitempty"`

	InAvg int64 `json:"inAvg,omitempty"`

	InMax int64 `json:"inMax,omitempty"`

	InMin int64 `json:"inMin,omitempty"`

	InTotal int64 `json:"inTotal,omitempty"`

	MaxBandwidth95ValueMbps float64 `json:"maxBandwidth95ValueMbps,omitempty"`

	Out95 int64 `json:"out95,omitempty"`

	Out95Time string `json:"out95Time,omitempty"`

	OutAvg int64 `json:"outAvg,omitempty"`

	OutMax int64 `json:"outMax,omitempty"`

	OutMin int64 `json:"outMin,omitempty"`

	OutTotal int64 `json:"outTotal,omitempty"`

	TotalUnit string `json:"totalUnit,omitempty"`

	Unit string `json:"unit,omitempty"`
}

type BandwidthClusterTrafficData struct {
	InternetRX int64 `json:"internetRX,omitempty"`

	InternetTX int64 `json:"internetTX,omitempty"`

	Time string `json:"time,omitempty"`
}
