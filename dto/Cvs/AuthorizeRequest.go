package Cvs

type AuthorizeRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	ServiceOptionType string `json:"serviceOptionType,omitempty"`
	OrderId           string `json:"orderId,omitempty"`
	Amount            string `json:"amount,omitempty"`
	Name1             string `json:"name1,omitempty"`
	Name2             string `json:"name2,omitempty"`
	Kana              string `json:"kana,omitempty"`
	TelNo             string `json:"telNo,omitempty"`
	PayLimit          string `json:"payLimit,omitempty"`
	PaymentType       string `json:"paymentType,omitempty"`
	Free1             string `json:"free1,omitempty"`
	Free2             string `json:"free2,omitempty"`
	PushUrl           string `json:"pushUrl,omitempty"`
	PayLimitHhmm      string `json:"payLimitHhmm,omitempty"`
}
