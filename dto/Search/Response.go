package Search

type Response struct {
	ServiceType      string      `json:"serviceType"`
	Mstatus          string      `json:"mstatus"`
	VResultCode      string      `json:"vResultCode"`
	MerrMsg          string      `json:"merrMsg"`
	OrderId          string      `json:"orderId"`
	MarchTxn         string      `json:"marchTxn"`
	CustTxn          string      `json:"custTxn"`
	TxnVersion       string      `json:"txnVersion"`
	OverMaxCountFlag bool        `json:"overMaxCountFlag"`
	MasterInfos      MasterInfos `json:"masterInfos"`
	SearchCount      int         `json:"searchCount"`
	OrderInfos       OrderInfos  `json:"orderInfos"`
	ResultJson       string      `json:"resultJson"`
}

type OrderInfos struct {
	OrderInfo []OrderInfo `json:"orderInfo"`
}

type OrderInfo struct {
	Memo1                string           `json:"memo1"`
	Memo2                string           `json:"memo2"`
	Memo3                string           `json:"memo3"`
	FreeKey              string           `json:"freeKey"`
	AccountId            string           `json:"accountId"`
	Index                int              `json:"index"`
	ServiceTypeCd        string           `json:"serviceTypeCd"`
	OrderId              string           `json:"orderId"`
	OrderStatus          string           `json:"orderStatus"`
	LastSuccessTxnType   string           `json:"lastSuccessTxnType"`
	SuccessDetailTxnType string           `json:"successDetailTxnType"`
	ProperOrderInfo      ProperOrderInfo  `json:"properOrderInfo"`
	TransactionInfos     TransactionInfos `json:"transactionInfos"`
}

type TransactionInfos struct {
	TransactionInfo []TransactionInfo `json:"transactionInfo"`
}

type ProperOrderInfo struct {
	SettlementMethod          string `json:"settlementMethod"`
	SettlementType            string `json:"settlementType"`
	Amount                    string `json:"amount"`
	AuthorizeAmount           string `json:"authorizeAmount"`
	Balance                   string `json:"balance"`
	UsedPoint                 string `json:"usedPoint"`
	UsePoint                  string `json:"usePoint"`
	GivePoint                 string `json:"givePoint"`
	RecruitCoupon             string `json:"recruitCoupon"`
	MerchantCoupon            string `json:"merchantCoupon"`
	SettlementLimit           string `json:"settlementLimit"`
	ForwardMailFlag           string `json:"forwardMailFlag"`
	MerchantMailAddr          string `json:"merchantMailAddr"`
	CancelMailAddr            string `json:"cancelMailAddr"`
	RequestMailAddInfo        string `json:"requestMailAddInfo"`
	CompleteMailAddInfo       string `json:"completeMailAddInfo"`
	ShopName                  string `json:"shopName"`
	CompleteMailFlag          string `json:"completeMailFlag"`
	ConfirmScreenAddInfo      string `json:"confirmScreenAddInfo"`
	CompleteScreenAddInfo     string `json:"completeScreenAddInfo"`
	ScreenTitle               string `json:"screenTitle"`
	CompleteReturnKind        string `json:"completeReturnKind"`
	CompleteReturnUrl         string `json:"completeReturnUrl"`
	CompleteNoticeUrl         string `json:"completeNoticeUrl"`
	SalesType                 string `json:"salesType"`
	Free                      string `json:"free"`
	RefundOrderCtlId          string `json:"refundOrderCtlId"`
	AppUrl                    string `json:"appUrl"`
	OrderKind                 string `json:"orderKind"`
	CompleteDatetime          string `json:"completeDatetime"`
	ReAuthorizeRedirectionUrl string `json:"reAuthorizeRedirectionUrl"`
	TransactionKind           string `json:"transactionKind"`
	UserId                    string `json:"userId"`
	SettlementId              string `json:"settlementId"`
	ReAuthAppUrl              string `json:"reAuthAppUrl"`
	CvsType                   string `json:"cvsType"`
	Name1                     string `json:"name1"`
	Name2                     string `json:"name2"`
	Kana                      string `json:"kana"`
	TelNo                     string `json:"telNo"`
	MailAddr                  string `json:"mailAddr"`
	Free1                     string `json:"free1"`
	Free2                     string `json:"free2"`
	PayLimit                  string `json:"payLimit"`
	PayLimitDatetime          string `json:"payLimitDatetime"`
	ReceiptNo                 string `json:"receiptNo"`
	PaidDatetime              string `json:"paidDatetime"`
	ReceivedDatetime          string `json:"receivedDatetime"`
	StartTxn                  string `json:"startTxn"`
	DddMessageVersion         string `json:"dddMessageVersion"`
	DeviceChannel             string `json:"deviceChannel"`
	AccountType               string `json:"accountType"`
	AuthenticationIndicator   string `json:"authenticationIndicator"`
	MessageCategory           string `json:"messageCategory"`
	RequestCurrencyUnit       string `json:"requestCurrencyUnit"`
	CardExpire                string `json:"cardExpire"`
	TradUrl                   string `json:"tradUrl"`
	InvoiceId                 string `json:"invoiceId"`
	PayerId                   string `json:"payerId"`
	PaymentDatetime           string `json:"paymentDatetime"`
	MerchantRedirectUri       string `json:"merchantRedirectUri"`
	TotalAmount               string `json:"totalAmount"`
	WalletAmount              string `json:"walletAmount"`
	CardAmount                string `json:"cardAmount"`
	CardOrderId               string `json:"cardOrderId"`
	CrServiceType             string `json:"crServiceType"`
	WithCapture               string `json:"withCapture"`
	AccountingType            string `json:"accountingType"`
	ItemInfo                  string `json:"itemInfo"`
	ItemId                    string `json:"itemId"`
	ItemType                  string `json:"itemType"`
	TerminalKind              string `json:"terminalKind"`
	AuthorizeDatetime         string `json:"authorizeDatetime"`
	CaptureDatetime           string `json:"captureDatetime"`
	CancelDatetime            string `json:"cancelDatetime"`
	MpFirstDate               string `json:"mpFirstDate"`
	MpDay                     string `json:"mpDay"`
	MpStatus                  string `json:"mpStatus"`
	MpOrderId                 string `json:"mpOrderId"`
	MpTxnStatusType           string `json:"mpTxnStatusType"`
	MpCaptureDatetime         string `json:"mpCaptureDatetime"`
	MpCancelDatetime          string `json:"mpCancelDatetime"`
	MpTerminateDatetime       string `json:"mpTerminateDatetime"`
	CrOrderId                 string `json:"crOrderId"`
	D3Flag                    string `json:"d3Flag"`
	FletsArea                 string `json:"fletsArea"`
	MerchantRedirectionUrl    string `json:"merchantRedirectionUrl"`
	OricoOrderNo              string `json:"oricoOrderNo"`
	UserNo                    string `json:"userNo"`
	ItemName                  string `json:"itemName"`
	ItemName1                 string `json:"itemName1"`
	ItemCount1                string `json:"itemCount1"`
	ItemAmount1               string `json:"itemAmount1"`
	ItemName2                 string `json:"itemName2"`
	ItemCount2                string `json:"itemCount2"`
	ItemAmount2               string `json:"itemAmount2"`
	ItemName3                 string `json:"itemName3"`
	ItemCount3                string `json:"itemCount3"`
	ItemAmount3               string `json:"itemAmount3"`
	ItemName4                 string `json:"itemName4"`
	ItemCount4                string `json:"itemCount4"`
	ItemAmount4               string `json:"itemAmount4"`
	ItemName5                 string `json:"itemName5"`
	ItemCount5                string `json:"itemCount5"`
	ItemAmount5               string `json:"itemAmount5"`
	TotalItemAmount           string `json:"totalItemAmount"`
	TotalCarriage             string `json:"totalCarriage"`
	Deposit                   string `json:"deposit"`
	ShippingZipCode           string `json:"shippingZipCode"`
	HandlingContractNo        string `json:"handlingContractNo"`
	MemberStoreNo             string `json:"memberStoreNo"`
	ContractDocumentKbn       string `json:"contractDocumentKbn"`
	WebDescriptionId          string `json:"webDescriptionId"`
	RakutenOrderId            string `json:"rakutenOrderId"`
	RecruitOrderId            string `json:"recruitOrderId"`
	LinepayOrderId            string `json:"linepayOrderId"`
	ItemAmount                string `json:"itemAmount"`
	AcquirerCode              string `json:"acquirerCode"`
	CardNumber                string `json:"cardNumber"`
	JpoInformation            string `json:"jpoInformation"`
	VaccDepositStatusType     string `json:"vaccDepositStatusType"`
	TransferExpiredDate       string `json:"transferExpiredDate"`
	ReconcileDate             string `json:"reconcileDate"`
	TotalDepositAmount        string `json:"totalDepositAmount"`
	EntryTransferName         string `json:"entryTransferName"`
	EntryTransferNumber       string `json:"entryTransferNumber"`
	AccountNumber             string `json:"accountNumber"`
	AccountManageType         string `json:"accountManageType"`
	TenpayServiceType         string `json:"tenpayServiceType"`
	ItemDetail                string `json:"itemDetail"`
	ItemLabel                 string `json:"itemLabel"`
	TenpayOrderId             string `json:"tenpayOrderId"`
	PaypayOrderId             string `json:"paypayOrderId"`
	PaypayServiceType         string `json:"paypayServiceType"`
	RefundDatetime            string `json:"refundDatetime"`
	TerminateDatetime         string `json:"terminateDatetime"`
	UserKey                   string `json:"userKey"`
	AlipayxServiceType        string `json:"alipayxServiceType"`
	CenterId                  string `json:"centerId"`
	PaymentSource             string `json:"paymentSource"`
	CvspayType                string `json:"cvspayType"`
	PayType                   string `json:"payType"`
	CvspayOrderId             string `json:"cvspayOrderId"`
	NissenTransactionId       string `json:"nissenTransactionId"`
	LastAuthorResult          string `json:"lastAuthorResult"`
	CaptureAmount             string `json:"captureAmount"`
	RefundableAmount          string `json:"refundableAmount"`
	CenterOrderId             string `json:"centerOrderId"`
	ConsentStatus             string `json:"consentStatus"`
	FrequencyUnit             string `json:"frequencyUnit"`
	FrequencyValue            string `json:"frequencyValue"`
	MerpayServiceType         string `json:"merpayServiceType"`
	ConsentAuthType           string `json:"consentAuthType"`
	PreapprovalId             string `json:"preapprovalId"`
	EposOrderId               string `json:"eposOrderId"`
	UseCredit                 string `json:"useCredit"`
	UseCoupon                 string `json:"useCoupon"`
	OdAmount                  string `json:"odAmount"`
	GatewayOrderId            string `json:"gatewayOrderId"`
	CardBrand                 string `json:"cardBrand"`
	CardLast4                 string `json:"cardLast4"`
	CardInstallments          string `json:"cardInstallments"`
	CardCvc                   string `json:"cardCvc"`
	Card3ds                   string `json:"card3ds"`
	CancelExpirationTime      string `json:"cancelExpirationTime"`
	CaptureExpirationTime     string `json:"captureExpirationTime"`
	UpdateExpirationTime      string `json:"updateExpirationTime"`
	ExtendAuthExpirationTime  string `json:"extendAuthExpirationTime"`
	PayConfirmScreenType      string `json:"payConfirmScreenType"`
	MemberId                  string `json:"memberId"`
	AccountId                 string `json:"accountId"`
}

type TransactionInfo struct {
	TxnId                 string                `json:"txnId"`
	Command               string                `json:"command"`
	Mstatus               string                `json:"mstatus"`
	VResultCode           string                `json:"vResultCode"`
	TxnDatetime           string                `json:"txnDatetime"`
	Amount                string                `json:"amount"`
	ProperTransactionInfo ProperTransactionInfo `json:"properTransactionInfo"`
}

type ProperTransactionInfo struct {
	TxnKind                  string `json:"txnKind"`
	EmTxnType                string `json:"emTxnType"`
	CenterProcDatetime       string `json:"centerProcDatetime"`
	CardType                 string `json:"cardType"`
	CardNo                   string `json:"cardNo"`
	CardBrandCode            string `json:"cardBrandCode"`
	SettlementStatus         string `json:"settlementStatus"`
	CvsTxnType               string `json:"cvsTxnType"`
	PeTxnType                string `json:"peTxnType"`
	ReceiptNo                string `json:"receiptNo"`
	StartDatetime            string `json:"startDatetime"`
	CardTransactionType      string `json:"cardTransactionType"`
	GatewayRequestDate       string `json:"gatewayRequestDate"`
	GatewayResponseDate      string `json:"gatewayResponseDate"`
	CenterRequestDate        string `json:"centerRequestDate"`
	CenterResponseDate       string `json:"centerResponseDate"`
	CenterRequestNumber      string `json:"centerRequestNumber"`
	CenterReferenceNumber    string `json:"centerReferenceNumber"`
	ReqItemCode              string `json:"reqItemCode"`
	ResItemCode              string `json:"resItemCode"`
	ReqReturnReferenceNumber string `json:"reqReturnReferenceNumber"`
	Responsedata             string `json:"responsedata"`
	Pending                  string `json:"pending"`
	Loopback                 string `json:"loopback"`
	ConnectedCenterId        string `json:"connectedCenterId"`
	ReqCardNumber            string `json:"reqCardNumber"`
	ReqCardExpire            string `json:"reqCardExpire"`
	ReqAmount                string `json:"reqAmount"`
	ReqCardOptionType        string `json:"reqCardOptionType"`
	ReqMerchantTransaction   string `json:"reqMerchantTransaction"`
	ReqAuthCode              string `json:"reqAuthCode"`
	ReqAcquirerCode          string `json:"reqAcquirerCode"`
	ReqCardCenter            string `json:"reqCardCenter"`
	ReqJpoInformation        string `json:"reqJpoInformation"`
	ReqSalesDay              string `json:"reqSalesDay"`
	ReqCancelDay             string `json:"reqCancelDay"`
	ReqWithCapture           string `json:"reqWithCapture"`
	ReqWithDirect            string `json:"reqWithDirect"`
	Req3dMessageVersion      string `json:"req3dMessageVersion"`
	Req3dTransactionId       string `json:"req3dTransactionId"`
	Req3dTransactionStatus   string `json:"req3dTransactionStatus"`
	Req3dCavvAlgorithm       string `json:"req3dCavvAlgorithm"`
	Req3dCavv                string `json:"req3dCavv"`
	Req3dEci                 string `json:"req3dEci"`
	Req3dDsTransactionId     string `json:"req3dDsTransactionId"`
	Req3dServerTransactionId string `json:"req3dServerTransactionId"`
	ReqSecurityCode          string `json:"reqSecurityCode"`
	ReqAuthFlag              string `json:"reqAuthFlag"`
	ReqBirthday              string `json:"reqBirthday"`
	ReqTel                   string `json:"reqTel"`
	ReqFirstKanaName         string `json:"reqFirstKanaName"`
	ReqLastKanaName          string `json:"reqLastKanaName"`
	ResMerchantTransaction   string `json:"resMerchantTransaction"`
	ResReturnReferenceNumber string `json:"resReturnReferenceNumber"`
	ResAuthCode              string `json:"resAuthCode"`
	ResActionCode            string `json:"resActionCode"`
	ResCenterErrorCode       string `json:"resCenterErrorCode"`
	ResAuthTerm              string `json:"resAuthTerm"`
	ReqWithNew               string `json:"reqWithNew"`
	Amount                   string `json:"amount"`
	TxnFixed                 string `json:"txnFixed"`
	PpTxnType                string `json:"ppTxnType"`
	CenterTxnId              string `json:"centerTxnId"`
	FeeAmount                string `json:"feeAmount"`
	ExchangeRate             string `json:"exchangeRate"`
	NetRefundAmount          string `json:"netRefundAmount"`
	MpiTransactionType       string `json:"mpiTransactionType"`
	ReqRedirectionUri        string `json:"reqRedirectionUri"`
	CorporationId            string `json:"corporationId"`
	BrandId                  string `json:"brandId"`
	AcquirerBinary           string `json:"acquirerBinary"`
	DsLoginId                string `json:"dsLoginId"`
	CrresStatus              string `json:"crresStatus"`
	VeresStatus              string `json:"veresStatus"`
	ParesStatus              string `json:"paresStatus"`
	ParesSign                string `json:"paresSign"`
	ParesEci                 string `json:"paresEci"`
	AuthResponseCode         string `json:"authResponseCode"`
	VerifyResponseCode       string `json:"verifyResponseCode"`
	Res3dMessageVersion      string `json:"res3dMessageVersion"`
	Res3dTransactionId       string `json:"res3dTransactionId"`
	Res3dDsTransactionId     string `json:"res3dDsTransactionId"`
	Res3dServerTransactionId string `json:"res3dServerTransactionId"`
	Res3dTransactionStatus   string `json:"res3dTransactionStatus"`
	Res3dCavvAlgorithm       string `json:"res3dCavvAlgorithm"`
	Res3dCavv                string `json:"res3dCavv"`
	Res3dEci                 string `json:"res3dEci"`
	AuthRequestDatetime      string `json:"authRequestDatetime"`
	AuthResponseDatetime     string `json:"authResponseDatetime"`
	VerifyRequestDatetime    string `json:"verifyRequestDatetime"`
	VerifyResponseDatetime   string `json:"verifyResponseDatetime"`
	ReqCurrencyUnit          string `json:"reqCurrencyUnit"`
	UpopTxnType              string `json:"upopTxnType"`
	ResUpopSettleAmount      string `json:"resUpopSettleAmount"`
	ResUpopSettleDate        string `json:"resUpopSettleDate"`
	ResUpopSettleCurrency    string `json:"resUpopSettleCurrency"`
	ResUpopExchangeDate      string `json:"resUpopExchangeDate"`
	ResUpopExchangeRate      string `json:"resUpopExchangeRate"`
	ResUpopOrderId           string `json:"resUpopOrderId"`
	CenterTradeId            string `json:"centerTradeId"`
	AlipayTxnType            string `json:"alipayTxnType"`
	SettleAmount             string `json:"settleAmount"`
	SettleCurrency           string `json:"settleCurrency"`
	PaymentTime              string `json:"paymentTime"`
	SettlementTime           string `json:"settlementTime"`
	PayType                  string `json:"payType"`
	CrResultCode             string `json:"crResultCode"`
	DetailCommandType        string `json:"detailCommandType"`
	CrRequestDatetime        string `json:"crRequestDatetime"`
	CrResponseDatetime       string `json:"crResponseDatetime"`
	OricoTxnType             string `json:"oricoTxnType"`
	OrderStateCode           string `json:"orderStateCode"`
	ApprovalNo               string `json:"approvalNo"`
	RequestDate              string `json:"requestDate"`
	LoanPrincipal            string `json:"loanPrincipal"`
	PaymentCount             string `json:"paymentCount"`
	JpyAmount                string `json:"jpyAmount"`
	ResMcpResponseCode       string `json:"resMcpResponseCode"`
	RakutenApiErrorCode      string `json:"rakutenApiErrorCode"`
	RakutenOrderErrorCode    string `json:"rakutenOrderErrorCode"`
	RakutenRequestDatetime   string `json:"rakutenRequestDatetime"`
	RakutenResponseDatetime  string `json:"rakutenResponseDatetime"`
	RecruitErrorCode         string `json:"recruitErrorCode"`
	RecruitRequestDatetime   string `json:"recruitRequestDatetime"`
	RecruitResponseDatetime  string `json:"recruitResponseDatetime"`
	LinepayErrorCode         string `json:"linepayErrorCode"`
	LinepayRequestDatetime   string `json:"linepayRequestDatetime"`
	LinepayResponseDatetime  string `json:"linepayResponseDatetime"`
	AuthCode                 string `json:"authCode"`
	ReferenceNumber          string `json:"referenceNumber"`
	CardVResultCode          string `json:"cardVResultCode"`
	WithReconcile            string `json:"withReconcile"`
	DepositId                string `json:"depositId"`
	RegistrationMethod       string `json:"registrationMethod"`
	DepositDate              string `json:"depositDate"`
	TransferName             string `json:"transferName"`
	TenpayErrorCode          string `json:"tenpayErrorCode"`
	TenpayRequestDatetime    string `json:"tenpayRequestDatetime"`
	TenpayResponseDatetime   string `json:"tenpayResponseDatetime"`
	ResCenterProcessNumber   string `json:"resCenterProcessNumber"`
	ResCenterSendDateTime    string `json:"resCenterSendDateTime"`
	StoreId                  string `json:"storeId"`
	TerminalId               string `json:"terminalId"`
	PaypayResultCode         string `json:"paypayResultCode"`
	SettleBizCode            string `json:"settleBizCode"`
	ReceiptNumber            string `json:"receiptNumber"`
	ItemName                 string `json:"itemName"`
	PaypayRequestDatetime    string `json:"paypayRequestDatetime"`
	PaypayResponseDatetime   string `json:"paypayResponseDatetime"`
	PaypayOnlineTxnId        string `json:"paypayOnlineTxnId"`
	PaypayErrorCode          string `json:"paypayErrorCode"`
	GatewayTransId           string `json:"gatewayTransId"`
	AlipayxResultCode        string `json:"alipayxResultCode"`
	AlipayxRequestDatetime   string `json:"alipayxRequestDatetime"`
	AlipayxResponseDatetime  string `json:"alipayxResponseDatetime"`
	CvspayResultCd           string `json:"cvspayResultCd"`
	CvspayRequestDatetime    string `json:"cvspayRequestDatetime"`
	CvspayResponseDatetime   string `json:"cvspayResponseDatetime"`
	CvspayCancelOrderId      string `json:"cvspayCancelOrderId"`
	AuthorResult             string `json:"authorResult"`
	CenterTransactionId      string `json:"centerTransactionId"`
	CenterResultCode         string `json:"centerResultCode"`
	CenterStateCode          string `json:"centerStateCode"`
	CenterReasonCode         string `json:"centerReasonCode"`
	MerpayErrorCode          string `json:"merpayErrorCode"`
	OperatorId               string `json:"operatorId"`
	SlipCode                 string `json:"slipCode"`
	MerpayResourceId         string `json:"merpayResourceId"`
	MerpayProcessingId       string `json:"merpayProcessingId"`
	MerpayRequestDatetime    string `json:"merpayRequestDatetime"`
	MerpayResponseDatetime   string `json:"merpayResponseDatetime"`
	InquiryCode              string `json:"inquiryCode"`
	UseCredit                string `json:"useCredit"`
	UseCoupon                string `json:"useCoupon"`
	UsePoint                 string `json:"usePoint"`
	EposRequestDatetime      string `json:"eposRequestDatetime"`
	EposResponseDatetime     string `json:"eposResponseDatetime"`
	ReqPaymentType           string `json:"reqPaymentType"`
	ReqResponseCode          string `json:"reqResponseCode"`
	FdResult                 string `json:"fdResult"`
	FrequencyUnit            string `json:"frequencyUnit"`
	FrequencyValue           string `json:"frequencyValue"`
	RakutenApiErrorType      string `json:"rakutenApiErrorType"`
	BankpayErrorCode         string `json:"bankpayErrorCode"`
	BankpayRequestDatetime   string `json:"bankpayRequestDatetime"`
	BankpayResponseDatetime  string `json:"bankpayResponseDatetime"`
}

type MasterInfos struct {
	MasterInfo []MasterInfo `json:"masterInfo"`
}

type MasterInfo struct {
	Name    string  `json:"name"`
	Masters Masters `json:"masters"`
}

type Masters struct {
	BankFinancialInstInfo []BankFinancialInstInfo `json:"bankFinancialInstInfo"`
}

type BankFinancialInstInfo struct {
	BankCode       string `json:"bankCode"`
	DeviceCode     string `json:"deviceCode"`
	BankName       string `json:"bankName"`
	BankKana       string `json:"bankKana"`
	BankIndexChar1 string `json:"bankIndexChar1"`
	BankIndexChar2 string `json:"bankIndexChar2"`
	StartDatetime  string `json:"startDatetime"`
}