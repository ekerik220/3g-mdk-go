package Cvs

type CancelRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	ServiceOptionType string `json:"serviceOptionType,omitempty"`
	OrderId           string `json:"orderId,omitempty"`
}
