package ScoreAtPay

type CaptureRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	OrderId       string `json:"orderId,omitempty"`
	PdCompanyCode string `json:"pdCompanyCode,omitempty"`
	SlipNo        string `json:"slipNo,omitempty"`
	DeliveryId    string `json:"deliveryId,omitempty"`
}
