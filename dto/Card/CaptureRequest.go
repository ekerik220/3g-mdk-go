package Card

import (
	"3g-mdk-go/dto/PayNowId"
)

type CaptureRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`

	OrderId           string `json:"orderId,omitempty"`
	OriginalOrderId   string `json:"originalOrderId,omitempty"`
	Amount            string `json:"amount,omitempty"`
	CardNumber        string `json:"cardNumber,omitempty"`
	CardExpire        string `json:"cardExpire,omitempty"`
	CardOptionType    string `json:"cardOptionType,omitempty"`
	AcquirerCode      string `json:"acquirerCode,omitempty"`
	SalesDay          string `json:"salesDay,omitempty"`
	ItemCode          string `json:"itemCode,omitempty"`
	WithDirect        string `json:"withDirect,omitempty"`
	AuthCode          string `json:"authCode,omitempty"`
	CurrencyUnit      string `json:"currencyUnit,omitempty"`
	Jpo               string `json:"jpo,omitempty"`
	FirstPayment      string `json:"firstPayment,omitempty"`
	BonusFirstPayment string `json:"bonusFirstPayment,omitempty"`
	McAmount          string `json:"mcAmount,omitempty"`
	TerminalId        string `json:"terminalId,omitempty"`
	ExSlipInfo        string `json:"exSlipInfo,omitempty"`
}
