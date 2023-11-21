package Carrier

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type CaptureRequest struct {
	TxnVersion        string          `json:"txnVersion,omitempty"`
	DummyRequest      string          `json:"dummyRequest,omitempty"`
	MerchantCcid      string          `json:"merchantCcid,omitempty"`
	PayNowIdParam     *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId           string          `json:"orderId,omitempty"`
	ServiceOptionType string          `json:"serviceOptionType,omitempty"`
	CaptureMonth      string          `json:"captureMonth,omitempty"`
}
