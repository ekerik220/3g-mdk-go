package Carrier

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
	RedirectUrl      string            `json:"redirectUrl"`
	ResponseContents string            `json:"responseContents"`
	TxnTime          string            `json:"txnTime"`
	CrOrderId        string            `json:"crOrderId"`
	DiscountAmount   string            `json:"discountAmount"`
	SeqNo            string            `json:"seqNo"`
	ResultJson       string            `json:"resultJson"`
}
