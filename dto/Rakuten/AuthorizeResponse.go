package Rakuten

import (
	"3g-mdk-go/dto/PayNowId"
)

type AuthorizeResponse struct {
	PayNowIdResponse PayNowId.Response `json:"payNowIdResponse"`
	ServiceType      string            `json:"serviceType"`
	Mstatus          string            `json:"mstatus"`
	VResultCode      string            `json:"vResultCode"`
	MerrMsg          string            `json:"merrMsg"`
	MarchTxn         string            `json:"marchTxn"`
	OrderId          string            `json:"orderId"`
	CustTxn          string            `json:"custTxn"`
	TxnVersion       string            `json:"txnVersion"`
	ResponseContents string            `json:"responseContents"`
	RedirectUrl      string            `json:"redirectUrl"`
	RakutenOrderId   string            `json:"rakutenOrderId"`
	GatewayOrderId   string            `json:"gatewayOrderId"`
	ResultJson       string            `json:"resultJson"`
}
