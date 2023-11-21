package Rakuten

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type UpdateAuthorizeRequest struct {
	TxnVersion        string          `json:"txnVersion,omitempty"`
	DummyRequest      string          `json:"dummyRequest,omitempty"`
	MerchantCcid      string          `json:"merchantCcid,omitempty"`
	PayNowIdParam     *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId           string          `json:"orderId,omitempty"`
	Amount            string          `json:"amount,omitempty"`
	AuthExtensionFlag string          `json:"authExtensionFlag,omitempty"`
}
