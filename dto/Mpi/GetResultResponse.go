package Mpi

import (
	"3g-mdk-go/dto/PayNowId"
)

type GetResultResponse struct {
	PayNowIdResponse PayNowId.Response `json:"payNowIdResponse"`

	ServiceType                string `json:"serviceType"`
	Mstatus                    string `json:"mstatus"`
	VResultCode                string `json:"vResultCode"`
	MerrMsg                    string `json:"merrMsg"`
	OrderId                    string `json:"orderId"`
	TxnVersion                 string `json:"txnVersion"`
	RequestId                  string `json:"requestId"`
	ReqAmount                  string `json:"reqAmount"`
	ReqCardNumber              string `json:"reqCardNumber"`
	ReqCurrencyUnit            string `json:"reqCurrencyUnit"`
	MpiMstatus                 string `json:"mpiMstatus"`
	MpiVresultCode             string `json:"mpiVresultCode"`
	TxnType                    string `json:"txnType"`
	CardMstatus                string `json:"cardMstatus"`
	CardTransactionType        string `json:"cardTransactionType"`
	CenterRequestDate          string `json:"centerRequestDate"`
	CenterResponseDate         string `json:"centerResponseDate"`
	ConnectedCenterId          string `json:"connectedCenterId"`
	AcquirerCode               string `json:"acquirerCode"`
	AuthCode                   string `json:"authCode"`
	FdResult                   string `json:"fdResult"`
	DddMessageVersion          string `json:"dddMessageVersion"`
	DddTransactionId           string `json:"dddTransactionId"`
	DddDsTransactionId         string `json:"dddDsTransactionId"`
	DddServerTransactionId     string `json:"dddServerTransactionId"`
	DddTransactionStatus       string `json:"dddTransactionStatus"`
	DddTransactionStatusReason string `json:"dddTransactionStatusReason"`
	DddCavvAlgorithm           string `json:"dddCavvAlgorithm"`
	DddCavv                    string `json:"dddCavv"`
	DddEci                     string `json:"dddEci"`
	ResultJson                 string `json:"resultJson"`
}
