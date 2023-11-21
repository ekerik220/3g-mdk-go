package LinePay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type AuthorizeResponse struct {
	PayNowIdResponse PayNowId.Response `json:"payNowIdResponse"`
	ServiceType      string            `json:"serviceType"`
	Mstatus          string            `json:"mstatus"`
	VResultCode      string            `json:"vResultCode"`
	MerrMsg          string            `json:"merrMsg"`
	MarchTxn         string            `json:"marchTxn"`
	OrderId          string            `json:"orderId"`
	LinepayOrderId   string            `json:"linepayOrderId"`
	CustTxn          string            `json:"custTxn"`
	TxnVersion       string            `json:"txnVersion"`
	RedirectWebUrl   string            `json:"redirectWebUrl"`
	RedirectAppUrl   string            `json:"redirectAppUrl"`
	ResultJson       string            `json:"resultJson"`
}
