package ScoreAtPay

type AuthorizeResponse struct {
	ServiceType         string  `json:"serviceType,omitempty"`
	Mstatus             string  `json:"mstatus,omitempty"`
	VResultCode         string  `json:"vResultCode,omitempty"`
	MerrMsg             string  `json:"merrMsg,omitempty"`
	MarchTxn            string  `json:"marchTxn,omitempty"`
	OrderId             string  `json:"orderId,omitempty"`
	CustTxn             string  `json:"custTxn,omitempty"`
	TxnVersion          string  `json:"txnVersion,omitempty"`
	ShopTransactionId   string  `json:"shopTransactionId,omitempty"`
	NissenTransactionId string  `json:"nissenTransactionId,omitempty"`
	AuthorResult        string  `json:"authorResult,omitempty"`
	Errors              []Error `json:"errors,omitempty"`
}
