package AmazonPay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type TerminateResponse struct {
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
	ResultJson             string            `json:"resultJson"`
}
