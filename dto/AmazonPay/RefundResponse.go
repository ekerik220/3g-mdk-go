package AmazonPay

import (
	"3g-mdk-go/dto/PayNowId"
)

type RefundResponse struct {
	PayNowIdResponse       PayNowId.Response `json:"payNowIdResponse"`
	ServiceType            string            `json:"serviceType"`
	Mstatus                string            `json:"mstatus"`
	VResultCode            string            `json:"vResultCode"`
	MerrMsg                string            `json:"merrMsg"`
	MarchTxn               string            `json:"marchTxn"`
	OrderId                string            `json:"orderId"`
	CustTxn                string            `json:"custTxn"`
	TxnVersion             string            `json:"txnVersion"`
	CenterResponseDatetime string            `json:"centerResponseDatetime"`
	CenterOrderId          string            `json:"centerOrderId"`
	CenterTransactionId    string            `json:"centerTransactionId"`
	RefundableAmount       int               `json:"refundableAmount"`
	ResultJson             string            `json:"resultJson"`
}
