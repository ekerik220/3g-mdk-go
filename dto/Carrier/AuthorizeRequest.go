package Carrier

import (
	"3g-mdk-go/dto/PayNowId"
)

type AuthorizeRequest struct {
	TxnVersion                string          `json:"txnVersion,omitempty"`
	DummyRequest              string          `json:"dummyRequest,omitempty"`
	MerchantCcid              string          `json:"merchantCcid,omitempty"`
	PayNowIdParam             *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId                   string          `json:"orderId,omitempty"`
	ServiceOptionType         string          `json:"serviceOptionType,omitempty"`
	Amount                    string          `json:"amount,omitempty"`
	TerminalKind              string          `json:"terminalKind,omitempty"`
	ItemType                  string          `json:"itemType,omitempty"`
	AccountingType            string          `json:"accountingType,omitempty"`
	WithCapture               string          `json:"withCapture,omitempty"`
	D3Flag                    string          `json:"d3Flag,omitempty"`
	MpFirstDate               string          `json:"mpFirstDate,omitempty"`
	MpDay                     string          `json:"mpDay,omitempty"`
	ItemId                    string          `json:"itemId,omitempty"`
	ItemInfo                  string          `json:"itemInfo,omitempty"`
	SuccessUrl                string          `json:"successUrl,omitempty"`
	CancelUrl                 string          `json:"cancelUrl,omitempty"`
	ErrorUrl                  string          `json:"errorUrl,omitempty"`
	PushUrl                   string          `json:"pushUrl,omitempty"`
	OpenId                    string          `json:"openId,omitempty"`
	SbUid                     string          `json:"sbUid,omitempty"`
	FletsArea                 string          `json:"fletsArea,omitempty"`
	LoginAuId                 string          `json:"loginAuId,omitempty"`
	AuId                      string          `json:"auId,omitempty"`
	PaymentAuthenticationType string          `json:"paymentAuthenticationType,omitempty"`
	BillingToken              string          `json:"billingToken,omitempty"`
	Idauth                    string          `json:"idauth,omitempty"`
	StoreId                   string          `json:"storeId,omitempty"`
	TermId                    string          `json:"termId,omitempty"`
	ReceiptNo                 string          `json:"receiptNo,omitempty"`
}
