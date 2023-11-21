package Search

type CommonSearchParameter struct {
	OrderId string `json:"orderId,omitempty"`

	OrderStatus []string `json:"orderStatus,omitempty"`

	Command []string `json:"command,omitempty"`

	Mstatus []string `json:"mstatus,omitempty"`

	TxnDatetime *Range `json:"txnDatetime,omitempty"`

	Amount *Range `json:"amount,omitempty"`
}
