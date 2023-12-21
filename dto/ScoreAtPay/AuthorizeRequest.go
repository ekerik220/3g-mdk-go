package ScoreAtPay

type AuthorizeRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	OrderId       string   `json:"orderId,omitempty"`
	Amount        int64    `json:"amount,omitempty"`
	ShopOrderDate string   `json:"shopOrderDate,omitempty"`
	BuyerContact  Contact  `json:"buyerContact,omitempty"`
	PaymentType   int64    `json:"paymentType,omitempty"`
	Delivery      Delivery `json:"delivery,omitempty"`
}
