package Card

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type CancelRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`

	OrderId             string `json:"orderId,omitempty"`
	OriginalOrderId     string `json:"originalOrderId,omitempty"`
	Amount              string `json:"amount,omitempty"`
	CardNumber          string `json:"cardNumber,omitempty"`
	CardExpire          string `json:"cardExpire,omitempty"`
	CardOptionType      string `json:"cardOptionType,omitempty"`
	CardCenter          string `json:"cardCenter,omitempty"`
	AcquirerCode        string `json:"acquirerCode,omitempty"`
	WithNew             string `json:"withNew,omitempty"`
	WithDirect          string `json:"withDirect,omitempty"`
	CancelDay           string `json:"cancelDay,omitempty"`
	ItemCode            string `json:"itemCode,omitempty"`
	CurrencyUnit        string `json:"currencyUnit,omitempty"`
	SalesDay            string `json:"salesDay,omitempty"`
	Pin                 string `json:"pin,omitempty"`
	PaymentType         string `json:"paymentType,omitempty"`
	Jis1SecondTrack     string `json:"jis1SecondTrack,omitempty"`
	Jis2Track           string `json:"jis2Track,omitempty"`
	UseOriginalOrder    string `json:"useOriginalOrder,omitempty"`
	Jpo                 string `json:"jpo,omitempty"`
	FirstPayment        string `json:"firstPayment,omitempty"`
	BonusFirstPayment   string `json:"bonusFirstPayment,omitempty"`
	WithDirectOnFailure string `json:"withDirectOnFailure,omitempty"`
	McAmount            string `json:"mcAmount,omitempty"`
	PosDataCode         string `json:"posDataCode,omitempty"`
	TerminalId          string `json:"terminalId,omitempty"`
	ChipConditionCode   string `json:"chipConditionCode,omitempty"`
	ExSlipInfo          string `json:"exSlipInfo,omitempty"`
}
