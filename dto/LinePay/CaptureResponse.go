package LinePay

import (
	"3g-mdk-go/dto/PayNowId"
)

type CaptureResponse struct {
	PayNowIdResponse PayNowId.Response `json:"payNowIdResponse"`
	ServiceType      string            `json:"serviceType"`
	Mstatus          string            `json:"mstatus"`
	VResultCode      string            `json:"vResultCode"`
	MerrMsg          string            `json:"merrMsg"`
	MarchTxn         string            `json:"marchTxn"`
	OrderId          string            `json:"orderId"`
	CustTxn          string            `json:"custTxn"`
	TxnVersion       string            `json:"txnVersion"`
	CaptureDatetime  string            `json:"captureDatetime"`
	Balance          int               `json:"balance"`
	ResultJson       string            `json:"resultJson"`
}