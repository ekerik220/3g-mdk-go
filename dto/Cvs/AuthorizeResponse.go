package Cvs

type AuthorizeResponse struct {
	ServiceType  string `json:"serviceType"`
	Mstatus      string `json:"mstatus"`
	VResultCode  string `json:"vResultCode"`
	MerrMsg      string `json:"merrMsg"`
	MarchTxn     string `json:"marchTxn"`
	OrderId      string `json:"orderId"`
	CustTxn      string `json:"custTxn"`
	ReceiptNo    string `json:"receiptNo"`
	HaraikomiUrl string `json:"haraikomiUrl"`
	TxnVersion   string `json:"txnVersion"`
	ResultJson   string `json:"resultJson"`
}
