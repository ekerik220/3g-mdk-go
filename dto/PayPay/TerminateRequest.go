package PayPay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type TerminateRequest struct {
	TxnVersion        string          `json:"txnVersion,omitempty"`
	DummyRequest      string          `json:"dummyRequest,omitempty"`
	MerchantCcid      string          `json:"merchantCcid,omitempty"`
	PayNowIdParam     *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId           string          `json:"orderId,omitempty"`
	ServiceOptionType string          `json:"serviceOptionType,omitempty"`
	Force             string          `json:"force,omitempty"`
	SuccessUrl        string          `json:"successUrl,omitempty"`
	CancelUrl         string          `json:"cancelUrl,omitempty"`
	ErrorUrl          string          `json:"errorUrl,omitempty"`
	PushUrl           string          `json:"pushUrl,omitempty"`
}
