package CardInfo

import (
	"3g-mdk-go/dto/PayNowId"
)

type Response struct {
	ServiceType      string            `json:"serviceType"`
	Mstatus          string            `json:"mstatus"`
	VResultCode      string            `json:"vResultCode"`
	MerrMsg          string            `json:"merrMsg"`
	MarchTxn         string            `json:"marchTxn"`
	CustTxn          string            `json:"custTxn"`
	TxnVersion       string            `json:"txnVersion"`
	PayNowIdResponse PayNowId.Response `json:"payNowIdResponse"`
	ResultJson       string            `json:"resultJson"`
}
