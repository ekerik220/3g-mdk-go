package Rakuten

import (
	"3g-mdk-go/dto/PayNowId"
)

type AuthorizeRequest struct {
	TxnVersion    string          `json:"txnVersion,omitempty"`
	DummyRequest  string          `json:"dummyRequest,omitempty"`
	MerchantCcid  string          `json:"merchantCcid,omitempty"`
	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId       string          `json:"orderId,omitempty"`
	PayType       string          `json:"payType,omitempty"`
	Amount        string          `json:"amount,omitempty"`
	WithCapture   string          `json:"withCapture,omitempty"`
	ItemId        string          `json:"itemId,omitempty"`
	ItemName      string          `json:"itemName,omitempty"`
	SuccessUrl    string          `json:"successUrl,omitempty"`
	ErrorUrl      string          `json:"errorUrl,omitempty"`
	PushUrl       string          `json:"pushUrl,omitempty"`
}
