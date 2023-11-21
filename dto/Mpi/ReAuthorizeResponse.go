package Mpi

import (
	"3g-mdk-go/dto/PayNowId"
)

type ReAuthorizeResponse struct {
	PayNowIdResponse PayNowId.Response `json:"payNowIdResponse"`

	ServiceType          string `json:"serviceType"`
	Mstatus              string `json:"mstatus"`
	VResultCode          string `json:"vResultCode"`
	MerrMsg              string `json:"merrMsg"`
	MarchTxn             string `json:"marchTxn"`
	OrderId              string `json:"orderId"`
	CustTxn              string `json:"custTxn"`
	TxnVersion           string `json:"txnVersion"`
	MpiTransactiontype   string `json:"mpiTransactiontype"`
	ReqCardNumber        string `json:"reqCardNumber"`
	ReqCardExpire        string `json:"reqCardExpire"`
	ReqAmount            string `json:"reqAmount"`
	ReqAcquirerCode      string `json:"reqAcquirerCode"`
	ReqItemCode          string `json:"reqItemCode"`
	ReqCardCenter        string `json:"reqCardCenter"`
	ReqJpoInformation    string `json:"reqJpoInformation"`
	ReqSalesDay          string `json:"reqSalesDay"`
	ReqWithCapture       string `json:"reqWithCapture"`
	ReqSecurityCode      string `json:"reqSecurityCode"`
	ReqBirthday          string `json:"reqBirthday"`
	ReqTel               string `json:"reqTel"`
	ReqFirstKanaName     string `json:"reqFirstKanaName"`
	ReqLastKanaName      string `json:"reqLastKanaName"`
	ReqCurrencyUnit      string `json:"reqCurrencyUnit"`
	ReqRedirectionUri    string `json:"reqRedirectionUri"`
	ReqHttpUserAgent     string `json:"reqHttpUserAgent"`
	ReqHttpAccept        string `json:"reqHttpAccept"`
	ResResponseContents  string `json:"resResponseContents"`
	ResCorporationId     string `json:"resCorporationId"`
	ResBrandId           string `json:"resBrandId"`
	Res3dMessageVersion  string `json:"res3dMessageVersion"`
	AuthRequestDatetime  string `json:"authRequestDatetime"`
	AuthResponseDatetime string `json:"authResponseDatetime"`
	KindCode             string `json:"kindCode"`
	AuthStartUrl         string `json:"authStartUrl"`
	ResultJson           string `json:"resultJson"`
}
