package PayPay

import (
	"3g-mdk-go/dto/PayNowId"
)

type RefundRequest struct {
	TxnVersion        string          `json:"txnVersion,omitempty"`
	DummyRequest      string          `json:"dummyRequest,omitempty"`
	MerchantCcid      string          `json:"merchantCcid,omitempty"`
	PayNowIdParam     *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId           string          `json:"orderId,omitempty"`
	ServiceOptionType string          `json:"serviceOptionType,omitempty"`
	StoreId           string          `json:"storeId,omitempty"`
	TerminalId        string          `json:"terminalId,omitempty"`
	ReceiptNumber     string          `json:"receiptNumber,omitempty"`
	SettleBizCode     string          `json:"settleBizCode,omitempty"`
	Amount            string          `json:"amount,omitempty"`
}
