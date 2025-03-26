package zbc

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2024-08-08"
	SERVICE    = "zbc"
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

func NewDescribeMonthlyBillSummaryRequest() (request *DescribeMonthlyBillSummaryRequest) {
	request = &DescribeMonthlyBillSummaryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeMonthlyBillSummary")

	return
}

// DescribeMonthlyBillSummary
// 用于获取产品或资源组维度月账单汇总金额信息以及账单总帐信息
//
// Possible error codes to return:
func (c *Client) DescribeMonthlyBillSummary(request *DescribeMonthlyBillSummaryRequest) (response *DescribeMonthlyBillSummaryResponse, err error) {
	response = NewDescribeMonthlyBillSummaryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeMonthlyBillSummaryResponse() (response *DescribeMonthlyBillSummaryResponse) {
	response = &DescribeMonthlyBillSummaryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeBillDetailRequest() (request *DescribeBillDetailRequest) {
	request = &DescribeBillDetailRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeBillDetail")

	return
}

// DescribeBillDetail
// 用于获取月账单明细相关数据信息
//
// Possible error codes to return:
func (c *Client) DescribeBillDetail(request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
	response = NewDescribeBillDetailResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeBillDetailResponse() (response *DescribeBillDetailResponse) {
	response = &DescribeBillDetailResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeCustomerBalanceRequest() (request *DescribeCustomerBalanceRequest) {
	request = &DescribeCustomerBalanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeCustomerBalance")

	return
}

// DescribeCustomerBalance
// 用于获取用户余额，包括账户余额、信用额度、可用金、冻结金等数据信息
//
// Possible error codes to return:
func (c *Client) DescribeCustomerBalance(request *DescribeCustomerBalanceRequest) (response *DescribeCustomerBalanceResponse, err error) {
	response = NewDescribeCustomerBalanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeCustomerBalanceResponse() (response *DescribeCustomerBalanceResponse) {
	response = &DescribeCustomerBalanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeRechargeRefundHistoryRequest() (request *DescribeRechargeRefundHistoryRequest) {
	request = &DescribeRechargeRefundHistoryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeRechargeRefundHistory")

	return
}

// DescribeRechargeRefundHistory
// 用于获取用户的充值记录以及订单退款记录等
//
// Possible error codes to return:
func (c *Client) DescribeRechargeRefundHistory(request *DescribeRechargeRefundHistoryRequest) (response *DescribeRechargeRefundHistoryResponse, err error) {
	response = NewDescribeRechargeRefundHistoryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeRechargeRefundHistoryResponse() (response *DescribeRechargeRefundHistoryResponse) {
	response = &DescribeRechargeRefundHistoryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func NewDescribeTransactionHistoryRequest() (request *DescribeTransactionHistoryRequest) {
	request = &DescribeTransactionHistoryRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeTransactionHistory")

	return
}

// DescribeTransactionHistory
// 用于获取账户的交易记录
//
// Possible error codes to return:
func (c *Client) DescribeTransactionHistory(request *DescribeTransactionHistoryRequest) (response *DescribeTransactionHistoryResponse, err error) {
	response = NewDescribeTransactionHistoryResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeTransactionHistoryResponse() (response *DescribeTransactionHistoryResponse) {
	response = &DescribeTransactionHistoryResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}
