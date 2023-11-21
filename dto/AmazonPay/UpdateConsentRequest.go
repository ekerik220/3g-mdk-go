package AmazonPay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type UpdateConsentRequest struct {
	TxnVersion     string          `json:"txnVersion,omitempty"`
	DummyRequest   string          `json:"dummyRequest,omitempty"`
	MerchantCcid   string          `json:"merchantCcid,omitempty"`
	PayNowIdParam  *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId        string          `json:"orderId,omitempty"`
	Amount         string          `json:"amount,omitempty"`
	FrequencyUnit  string          `json:"frequencyUnit,omitempty"`
	FrequencyValue string          `json:"frequencyValue,omitempty"`
	NoteToBuyer    string          `json:"noteToBuyer,omitempty"`
}
