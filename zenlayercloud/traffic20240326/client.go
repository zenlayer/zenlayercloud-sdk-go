package traffic

import (
	"gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2024-03-26"
	SERVICE    = "traffic"
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

func NewDescribeBandwidthClustersRequest() (request *DescribeBandwidthClustersRequest) {
	request = &DescribeBandwidthClustersRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBandwidthClusters")

	return
}

func NewDescribeBandwidthClustersResponse() (response *DescribeBandwidthClustersResponse) {
	response = &DescribeBandwidthClustersResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBandwidthClusters
// This API is used to query the details of bandwidth clusters. You can filter the query results with the BandwidthCluster ID, name, or city name.
func (c *Client) DescribeBandwidthClusters(request *DescribeBandwidthClustersRequest) (response *DescribeBandwidthClustersResponse, err error) {
	response = NewDescribeBandwidthClustersResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBandwidthClusterTrafficRequest() (request *DescribeBandwidthClusterTrafficRequest) {
	request = &DescribeBandwidthClusterTrafficRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBandwidthClusterTraffic")

	return
}

func NewDescribeBandwidthClusterTrafficResponse() (response *DescribeBandwidthClusterTrafficResponse) {
	response = &DescribeBandwidthClusterTrafficResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

// DescribeBandwidthClusterTraffic
// This API is used to query the data transfer in the specified date range.
//
// Possible error codes to return:
// INVALID_CLUSTER_NOT_FOUND = "Invalid.Cluster.Not.Found"
// INVALID_START_TIME_MUST_BEFORE_END_TIME = "Invalid.Start.Time.Must.Before.End.Time"
// INVALID_TIME_RANGE = "Invalid.Time.Range"
// TIME_OUT_OF_RANGE = "Time.Out.Of.Range"
// INVALID_TIME_FORMAT = "Invalid.Time.Format"
func (c *Client) DescribeBandwidthClusterTraffic(request *DescribeBandwidthClusterTrafficRequest) (response *DescribeBandwidthClusterTrafficResponse, err error) {
	response = NewDescribeBandwidthClusterTrafficResponse()
	err = c.ApiCall(request, response)
	return
}
