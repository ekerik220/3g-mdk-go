package ScoreAtPay

type ConfirmRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	OrderId string `json:"orderId,omitempty"`
}
