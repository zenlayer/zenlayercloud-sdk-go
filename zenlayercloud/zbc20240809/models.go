package zbc

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type DescribeMonthlyBillSummaryRequest struct {
	*common.BaseRequest

	SummaryType string `json:"summaryType,omitempty"`

	BillMonthly int `json:"billMonthly,omitempty"`
}

type MonthlyBillSummary struct {
	Product string `json:"product,omitempty"`

	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	Spend float64 `json:"spend,omitempty"`

	Voucher float64 `json:"voucher,omitempty"`

	Cash float64 `json:"cash,omitempty"`

	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	Refund float64 `json:"refund,omitempty"`
}

type DescribeMonthlyBillSummaryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	BillMonthly int `json:"billMonthly,omitempty"`

	TotalSpend float64 `json:"totalSpend,omitempty"`

	TotalCash float64 `json:"totalCash,omitempty"`

	TotalRefund float64 `json:"totalRefund,omitempty"`

	TotalVoucher float64 `json:"totalVoucher,omitempty"`

	SummaryInfoList []*MonthlyBillSummary `json:"summaryInfoList,omitempty"`
}

type DescribeMonthlyBillSummaryResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeMonthlyBillSummaryResponseParams `json:"response"`
}

type DescribeBillDetailRequest struct {
	*common.BaseRequest

	PageNum int `json:"pageNum,omitempty"`

	PageSize int `json:"pageSize,omitempty"`

	BillMonthly int `json:"billMonthly,omitempty"`

	OrderSn string `json:"orderSn,omitempty"`

	ResourceId string `json:"resourceId,omitempty"`
}

type BillingDetailInfo struct {
	ResourceId string `json:"resourceId,omitempty"`

	ProductSubitem string `json:"productSubitem,omitempty"`

	Product string `json:"product,omitempty"`

	Amount float64 `json:"amount,omitempty"`

	Voucher float64 `json:"voucher,omitempty"`

	Cash float64 `json:"cash,omitempty"`

	BillingMode string `json:"billingMode,omitempty"`

	OrderSn string `json:"orderSn,omitempty"`

	DeductionTime string `json:"deductionTime,omitempty"`

	Label string `json:"label,omitempty"`

	Location string `json:"location,omitempty"`

	BillMonthly int `json:"billMonthly,omitempty"`

	ResourceGroupId string `json:"resourceGroupId,omitempty"`

	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`
}

type DescribeBillDetailResponseParams struct {

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	TotalSpend float64 `json:"totalSpend,omitempty"`

	TotalVoucher float64 `json:"totalVoucher,omitempty"`

	TotalCash float64 `json:"totalCash,omitempty"`

	TotalRefund float64 `json:"totalRefund,omitempty"`

	TotalCount int `json:"totalCount,omitempty"`

	DataSet []*BillingDetailInfo `json:"dataSet,omitempty"`
}

type DescribeBillDetailResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeBillDetailResponseParams `json:"response"`
}

type DescribeCustomerBalanceRequest struct {
	*common.BaseRequest
}

type DescribeCustomerBalanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	AccountUid string `json:"accountUid,omitempty"`

	CashBalance float64 `json:"cashBalance,omitempty"`

	CreditValue float64 `json:"creditValue,omitempty"`

	PurchasingPower float64 `json:"purchasingPower,omitempty"`

	OrderFrozenAmount float64 `json:"orderFrozenAmount,omitempty"`
}

type DescribeCustomerBalanceResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeCustomerBalanceResponseParams `json:"response"`
}

type DescribeRechargeRefundHistoryRequest struct {
	*common.BaseRequest

	PageNum int `json:"pageNum,omitempty"`

	PageSize int `json:"pageSize,omitempty"`

	State []string `json:"state,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`
}

type RechargeRefundHistory struct {
	TransactionType string `json:"transactionType,omitempty"`

	Amount float64 `json:"amount,omitempty"`

	PaymentMethod string `json:"paymentMethod,omitempty"`

	TransactionId string `json:"transactionId,omitempty"`

	State string `json:"state,omitempty"`

	TransactionTime string `json:"transactionTime,omitempty"`

	AccountUid string `json:"accountUid,omitempty"`
}

type DescribeRechargeRefundHistoryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	TotalCount int `json:"totalCount,omitempty"`

	DataSet []*RechargeRefundHistory `json:"dataSet,omitempty"`
}

type DescribeRechargeRefundHistoryResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeRechargeRefundHistoryResponseParams `json:"response"`
}

type DescribeTransactionHistoryRequest struct {
	*common.BaseRequest

	PageNum int `json:"pageNum,omitempty"`

	PageSize int `json:"pageSize,omitempty"`

	TransactionType []string `json:"transactionType,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`
}

type TransactionHistory struct {
	AccountUid string `json:"accountUid,omitempty"`

	Uid string `json:"uid,omitempty"`

	TransactionTime string `json:"transactionTime,omitempty"`

	OrderSn string `json:"orderSn,omitempty"`

	Description string `json:"description,omitempty"`

	TransactionType string `json:"transactionType,omitempty"`

	PaymentMethod string `json:"paymentMethod,omitempty"`

	Amount float64 `json:"amount,omitempty"`

	CashBalance float64 `json:"cashBalance,omitempty"`

	TradeType string `json:"tradeType,omitempty"`

	TradeNo string `json:"tradeNo,omitempty"`
}

type DescribeTransactionHistoryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	IncomeSum float64 `json:"incomeSum,omitempty"`

	ExpenditureSum float64 `json:"expenditureSum,omitempty"`

	TotalCount int `json:"totalCount,omitempty"`

	DataSet []*TransactionHistory `json:"dataSet,omitempty"`
}

type DescribeTransactionHistoryResponse struct {
	*common.BaseResponse

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId string `json:"requestId,omitempty"`

	Response *DescribeTransactionHistoryResponseParams `json:"response"`
}
