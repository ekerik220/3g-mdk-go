package AmazonPay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type UpdateAuthorizeResponse struct {
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
	PayRedirectUrl         string            `json:"payRedirectUrl"`
	ResultJson             string            `json:"resultJson"`
}
