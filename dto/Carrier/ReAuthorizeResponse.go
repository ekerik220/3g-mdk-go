package Carrier

import (
	"3g-mdk-go/dto/PayNowId"
)

type ReAuthorizeResponse struct {
	PayNowIdResponse PayNowId.Response `json:"payNowIdResponse"`
	ServiceType      string            `json:"serviceType"`
	Mstatus          string            `json:"mstatus"`
	VResultCode      string            `json:"vResultCode"`
	MerrMsg          string            `json:"merrMsg"`
	MarchTxn         string            `json:"marchTxn"`
	OrderId          string            `json:"orderId"`
	OriginalOrderId  string            `json:"originalOrderId"`
	CrOrderId        string            `json:"crOrderId"`
	TxnVersion       string            `json:"txnVersion"`
	TxnTime          string            `json:"txnTime"`
	ResultJson       string            `json:"resultJson"`
}
