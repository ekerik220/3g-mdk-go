package Carrier

import (
	"3g-mdk-go/dto/PayNowId"
)

type TerminateResponse struct {
	PayNowIdResponse  PayNowId.Response `json:"payNowIdResponse"`
	ServiceType       string            `json:"serviceType"`
	Mstatus           string            `json:"mstatus"`
	VResultCode       string            `json:"vResultCode"`
	MerrMsg           string            `json:"merrMsg"`
	MarchTxn          string            `json:"marchTxn"`
	OrderId           string            `json:"orderId"`
	CustTxn           string            `json:"custTxn"`
	TxnVersion        string            `json:"txnVersion"`
	RedirectUrl       string            `json:"redirectUrl"`
	ResponseContents  string            `json:"responseContents"`
	TerminateDatetime string            `json:"terminateDatetime"`
	ResultJson        string            `json:"resultJson"`
}
