package AmazonPay

import (
	"3g-mdk-go/dto/PayNowId"
)

type GetAddressRequest struct {
	TxnVersion      string          `json:"txnVersion,omitempty"`
	DummyRequest    string          `json:"dummyRequest,omitempty"`
	MerchantCcid    string          `json:"merchantCcid,omitempty"`
	PayNowIdParam   *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId         string          `json:"orderId,omitempty"`
	AddInfoRespFlag string          `json:"addInfoRespFlag,omitempty"`
}
