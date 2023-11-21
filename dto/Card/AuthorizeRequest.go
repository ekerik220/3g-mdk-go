package Card

import (
	"3g-mdk-go/dto/FraudDetectionV2"
	"3g-mdk-go/dto/PayNowId"
)

type AuthorizeRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`

	OrderId                 string                    `json:"orderId,omitempty"`
	Amount                  string                    `json:"amount,omitempty"`
	CardNumber              string                    `json:"cardNumber,omitempty"`
	CardExpire              string                    `json:"cardExpire,omitempty"`
	CardOptionType          string                    `json:"cardOptionType,omitempty"`
	CardCenter              string                    `json:"cardCenter,omitempty"`
	AcquirerCode            string                    `json:"acquirerCode,omitempty"`
	Jpo                     string                    `json:"jpo,omitempty"`
	WithCapture             string                    `json:"withCapture,omitempty"`
	SalesDay                string                    `json:"salesDay,omitempty"`
	ItemCode                string                    `json:"itemCode,omitempty"`
	DddMessageVersion       string                    `json:"dddMessageVersion,omitempty"`
	DddTransactionId        string                    `json:"dddTransactionId,omitempty"`
	DddTransactionStatus    string                    `json:"dddTransactionStatus,omitempty"`
	DddCavvAlgorithm        string                    `json:"dddCavvAlgorithm,omitempty"`
	DddCavv                 string                    `json:"dddCavv,omitempty"`
	DddEci                  string                    `json:"dddEci,omitempty"`
	DddDsTransactionId      string                    `json:"dddDsTransactionId,omitempty"`
	DddServerTransactionId  string                    `json:"dddServerTransactionId,omitempty"`
	SecurityCode            string                    `json:"securityCode,omitempty"`
	AuthFlag                string                    `json:"authFlag,omitempty"`
	Birthday                string                    `json:"birthday,omitempty"`
	Tel                     string                    `json:"tel,omitempty"`
	FirstKanaName           string                    `json:"firstKanaName,omitempty"`
	LastKanaName            string                    `json:"lastKanaName,omitempty"`
	CurrencyUnit            string                    `json:"currencyUnit,omitempty"`
	Pin                     string                    `json:"pin,omitempty"`
	PaymentType             string                    `json:"paymentType,omitempty"`
	Jis1SecondTrack         string                    `json:"jis1SecondTrack,omitempty"`
	Jis2Track               string                    `json:"jis2Track,omitempty"`
	FirstPayment            string                    `json:"firstPayment,omitempty"`
	BonusFirstPayment       string                    `json:"bonusFirstPayment,omitempty"`
	McAmount                string                    `json:"mcAmount,omitempty"`
	PosDataCode             string                    `json:"posDataCode,omitempty"`
	TerminalId              string                    `json:"terminalId,omitempty"`
	FraudDetectionV2Request *FraudDetectionV2.Request `json:"fraudDetectionV2Request,omitempty"`
	WithFraudDetection      string                    `json:"withFraudDetection,omitempty"`
	ChipConditionCode       string                    `json:"chipConditionCode,omitempty"`
	WithJcnTokenization     string                    `json:"withJcnTokenization,omitempty"`
	JcnToken                string                    `json:"jcnToken,omitempty"`
	JcnTokenVersion         string                    `json:"jcnTokenVersion,omitempty"`
	ExSlipInfo              string                    `json:"exSlipInfo,omitempty"`
}
