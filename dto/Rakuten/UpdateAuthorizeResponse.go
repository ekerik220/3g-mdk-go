package Rakuten

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type UpdateAuthorizeResponse struct {
	PayNowIdResponse         PayNowId.Response `json:"payNowIdResponse"`
	ServiceType              string            `json:"serviceType"`
	Mstatus                  string            `json:"mstatus"`
	VResultCode              string            `json:"vResultCode"`
	MerrMsg                  string            `json:"merrMsg"`
	MarchTxn                 string            `json:"marchTxn"`
	OrderId                  string            `json:"orderId"`
	CustTxn                  string            `json:"custTxn"`
	TxnVersion               string            `json:"txnVersion"`
	TransactionDatetime      string            `json:"transactionDatetime"`
	RakutenOrderId           string            `json:"rakutenOrderId"`
	GatewayOrderId           string            `json:"gatewayOrderId"`
	Amount                   string            `json:"amount"`
	UsedPoint                string            `json:"usedPoint"`
	CancelExpirationTime     string            `json:"cancelExpirationTime"`
	CaptureExpirationTime    string            `json:"captureExpirationTime"`
	UpdateExpirationTime     string            `json:"updateExpirationTime"`
	ExtendAuthExpirationTime string            `json:"extendAuthExpirationTime"`
	ResultJson               string            `json:"resultJson"`
}
