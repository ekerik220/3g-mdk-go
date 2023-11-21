package Carrier

import (
	"3g-mdk-go/dto/PayNowId"
)

type CancelRequest struct {
	TxnVersion        string          `json:"txnVersion,omitempty"`
	DummyRequest      string          `json:"dummyRequest,omitempty"`
	MerchantCcid      string          `json:"merchantCcid,omitempty"`
	PayNowIdParam     *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId           string          `json:"orderId,omitempty"`
	ServiceOptionType string          `json:"serviceOptionType,omitempty"`
	CancelMonth       string          `json:"cancelMonth,omitempty"`
	Amount            string          `json:"amount,omitempty"`
	StoreId           string          `json:"storeId,omitempty"`
	TermId            string          `json:"termId,omitempty"`
	ReceiptNo         string          `json:"receiptNo,omitempty"`
}
