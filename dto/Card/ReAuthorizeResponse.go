package Card

import (
	"github.com/ekerik220/3g-mdk-go/dto/FraudDetectionV2"
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type ReAuthorizeResponse struct {
	PayNowIdResponse         PayNowId.Response         `json:"payNowIdResponse"`
	ServiceType              string                    `json:"serviceType"`
	Mstatus                  string                    `json:"mstatus"`
	VResultCode              string                    `json:"vResultCode"`
	MerrMsg                  string                    `json:"merrMsg"`
	MarchTxn                 string                    `json:"marchTxn"`
	OrderId                  string                    `json:"orderId"`
	CustTxn                  string                    `json:"custTxn"`
	TxnVersion               string                    `json:"txnVersion"`
	CardTransactionType      string                    `json:"cardTransactiontype"`
	GatewayRequestDate       string                    `json:"gatewayRequestDate"`
	GatewayResponseDate      string                    `json:"gatewayResponseDate"`
	CenterRequestDate        string                    `json:"centerRequestDate"`
	CenterResponseDate       string                    `json:"centerResponseDate"`
	Pending                  string                    `json:"pending"`
	Loopback                 string                    `json:"loopback"`
	ConnectedCenterId        string                    `json:"connectedCenterId"`
	CenterRequestNumber      string                    `json:"centerRequestNumber"`
	CenterReferenceNumber    string                    `json:"centerReferenceNumber"`
	ReqCardNumber            string                    `json:"reqCardNumber"`
	ReqCardExpire            string                    `json:"reqCardExpire"`
	ReqCardOptionType        string                    `json:"reqCardOptionType"`
	ReqAmount                string                    `json:"reqAmount"`
	ReqMerchantTransaction   string                    `json:"reqMerchantTransaction"`
	ReqReturnReferenceNumber string                    `json:"reqReturnReferenceNumber"`
	ReqAuthCode              string                    `json:"reqAuthCode"`
	ReqAcquirerCode          string                    `json:"reqAcquirerCode"`
	ReqItemCode              string                    `json:"reqItemCode"`
	ReqCardCenter            string                    `json:"reqCardCenter"`
	ReqJpoInformation        string                    `json:"reqJpoInformation"`
	ReqSalesDay              string                    `json:"reqSalesDay"`
	ReqCancelDay             string                    `json:"reqCancelDay"`
	ReqWithCapture           string                    `json:"reqWithCapture"`
	ReqWithDirect            string                    `json:"reqWithDirect"`
	Req3dMessageVersion      string                    `json:"req3dMessageVersion"`
	Req3dTransactionId       string                    `json:"req3dTransactionId"`
	Req3dTransactionStatus   string                    `json:"req3dTransactionStatus"`
	Req3dCavvAlgorithm       string                    `json:"req3dCavvAlgorithm"`
	Req3dCavv                string                    `json:"req3dCavv"`
	Req3dEci                 string                    `json:"req3dEci"`
	ReqSecurityCode          string                    `json:"reqSecurityCode"`
	ReqAuthFlag              string                    `json:"reqAuthFlag"`
	ReqBirthday              string                    `json:"reqBirthday"`
	ReqTel                   string                    `json:"reqTel"`
	ReqFirstKanaName         string                    `json:"reqFirstKanaName"`
	ReqLastKanaName          string                    `json:"reqLastKanaName"`
	ResMerchantTransaction   string                    `json:"resMerchantTransaction"`
	ResReturnReferenceNumber string                    `json:"resReturnReferenceNumber"`
	ResAuthCode              string                    `json:"resAuthCode"`
	ResActionCode            string                    `json:"resActionCode"`
	ResCenterErrorCode       string                    `json:"resCenterErrorCode"`
	ResAuthTerm              string                    `json:"resAuthTerm"`
	ResItemCode              string                    `json:"resItemCode"`
	ResResponseData          string                    `json:"resResponseData"`
	ReqCurrencyUnit          string                    `json:"reqCurrencyUnit"`
	ReqWithNew               string                    `json:"reqWithNew"`
	AcquirerCode             string                    `json:"acquirerCode"`
	TerminalId               string                    `json:"terminalId"`
	FraudDetectionResponse   FraudDetectionV2.Response `json:"fraudDetectionResponse"`
	KindCode                 string                    `json:"kindCode"`
	ResultJson               string                    `json:"resultJson"`
}
