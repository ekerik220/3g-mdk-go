package PayPay

import (
	"3g-mdk-go/dto/PayNowId"
)

type AuthorizeRequest struct {
	TxnVersion          string          `json:"txnVersion,omitempty"`
	DummyRequest        string          `json:"dummyRequest,omitempty"`
	MerchantCcid        string          `json:"merchantCcid,omitempty"`
	PayNowIdParam       *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId             string          `json:"orderId,omitempty"`
	ServiceOptionType   string          `json:"serviceOptionType,omitempty"`
	AccountingType      string          `json:"accountingType,omitempty"`
	OneTimeCode         string          `json:"oneTimeCode,omitempty"`
	Amount              string          `json:"amount,omitempty"`
	WithCapture         string          `json:"withCapture,omitempty"`
	StoreId             string          `json:"storeId,omitempty"`
	TerminalId          string          `json:"terminalId,omitempty"`
	ReceiptNumber       string          `json:"receiptNumber,omitempty"`
	SettleBizCode       string          `json:"settleBizCode,omitempty"`
	ItemName            string          `json:"itemName,omitempty"`
	ItemId              string          `json:"itemId,omitempty"`
	SuccessUrl          string          `json:"successUrl,omitempty"`
	CancelUrl           string          `json:"cancelUrl,omitempty"`
	ErrorUrl            string          `json:"errorUrl,omitempty"`
	PushUrl             string          `json:"pushUrl,omitempty"`
	TransitionType      string          `json:"transitionType,omitempty"`
	ExtendParameterType string          `json:"extendParameterType,omitempty"`
}
