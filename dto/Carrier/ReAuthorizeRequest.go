package Carrier

import (
	"3g-mdk-go/dto/PayNowId"
)

type ReAuthorizeRequest struct {
	TxnVersion        string          `json:"txnVersion,omitempty"`
	DummyRequest      string          `json:"dummyRequest,omitempty"`
	MerchantCcid      string          `json:"merchantCcid,omitempty"`
	PayNowIdParam     *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId           string          `json:"orderId,omitempty"`
	OriginalOrderId   string          `json:"originalOrderId,omitempty"`
	ServiceOptionType string          `json:"serviceOptionType,omitempty"`
	WithCapture       string          `json:"withCapture,omitempty"`
	Amount            string          `json:"amount,omitempty"`
}
