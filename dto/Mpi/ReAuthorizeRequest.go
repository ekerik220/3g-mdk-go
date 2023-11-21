package Mpi

import (
	"github.com/ekerik220/3g-mdk-go/dto/FraudDetectionV2"
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type ReAuthorizeRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`

	ServiceOptionType string `json:"serviceOptionType,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	OriginalOrderId string `json:"originalOrderId,omitempty"`

	Amount string `json:"amount,omitempty"`

	CardNumber string `json:"cardNumber,omitempty"`

	CardExpire string `json:"cardExpire,omitempty"`

	CardCenter string `json:"cardCenter,omitempty"`

	AcquirerCode string `json:"acquirerCode,omitempty"`

	Jpo string `json:"jpo,omitempty"`

	WithCapture string `json:"withCapture,omitempty"`

	SalesDay string `json:"salesDay,omitempty"`

	ItemCode string `json:"itemCode,omitempty"`

	SecurityCode string `json:"securityCode,omitempty"`

	Birthday string `json:"birthday,omitempty"`

	Tel string `json:"tel,omitempty"`

	FirstKanaName string `json:"firstKanaName,omitempty"`

	LastKanaName string `json:"lastKanaName,omitempty"`

	CurrencyUnit string `json:"currencyUnit,omitempty"`

	RedirectionUri string `json:"redirectionUri,omitempty"`

	HttpUserAgent string `json:"httpUserAgent,omitempty"`

	HttpAccept string `json:"httpAccept,omitempty"`

	FirstPayment string `json:"firstPayment,omitempty"`

	BonusFirstPayment string `json:"bonusFirstPayment,omitempty"`

	McAmount string `json:"mcAmount,omitempty"`

	PushUrl string `json:"pushUrl,omitempty"`

	BrowserDeviceCategory string `json:"browserDeviceCategory,omitempty"`

	VerifyResultLink string `json:"verifyResultLink,omitempty"`

	VerifyTimeout string `json:"verifyTimeout,omitempty"`

	FraudDetectionV2Request *FraudDetectionV2.Request `json:"fraudDetectionV2Request,omitempty"`

	WithFraudDetection string `json:"withFraudDetection,omitempty"`

	DeviceChannel string `json:"deviceChannel,omitempty"`

	AccountType string `json:"accountType,omitempty"`

	AuthenticationIndicator string `json:"authenticationIndicator,omitempty"`

	MessageCategory string `json:"messageCategory,omitempty"`

	CardholderName string `json:"cardholderName,omitempty"`

	CardholderNameOmitFlag string `json:"cardholderNameOmitFlag,omitempty"`

	CardholderEmail string `json:"cardholderEmail,omitempty"`

	CardholderHomePhoneCountry string `json:"cardholderHomePhoneCountry,omitempty"`

	CardholderHomePhoneNumber string `json:"cardholderHomePhoneNumber,omitempty"`

	CardholderMobilePhoneCountry string `json:"cardholderMobilePhoneCountry,omitempty"`

	CardholderMobilePhoneNumber string `json:"cardholderMobilePhoneNumber,omitempty"`

	CardholderWorkPhoneCountry string `json:"cardholderWorkPhoneCountry,omitempty"`

	CardholderWorkPhoneNumber string `json:"cardholderWorkPhoneNumber,omitempty"`

	BillingAddressCity string `json:"billingAddressCity,omitempty"`

	BillingAddressCountry string `json:"billingAddressCountry,omitempty"`

	BillingAddressLine1 string `json:"billingAddressLine1,omitempty"`

	BillingAddressLine2 string `json:"billingAddressLine2,omitempty"`

	BillingAddressLine3 string `json:"billingAddressLine3,omitempty"`

	BillingPostalCode string `json:"billingPostalCode,omitempty"`

	BillingAddressState string `json:"billingAddressState,omitempty"`

	ShippingAddressCity string `json:"shippingAddressCity,omitempty"`

	ShippingAddressCountry string `json:"shippingAddressCountry,omitempty"`

	ShippingAddressLine1 string `json:"shippingAddressLine1,omitempty"`

	ShippingAddressLine2 string `json:"shippingAddressLine2,omitempty"`

	ShippingAddressLine3 string `json:"shippingAddressLine3,omitempty"`

	ShippingPostalCode string `json:"shippingPostalCode,omitempty"`

	ShippingAddressState string `json:"shippingAddressState,omitempty"`

	CustomerIp string `json:"customerIp,omitempty"`

	WithChallenge string `json:"withChallenge,omitempty"`

	AuthData string `json:"authData,omitempty"`

	AuthMethod string `json:"authMethod,omitempty"`

	AuthTimestamp string `json:"authTimestamp,omitempty"`

	AccountAgeIndicator string `json:"accountAgeIndicator,omitempty"`

	AccountChangeDate string `json:"accountChangeDate,omitempty"`

	AccountChangeIndicator string `json:"accountChangeIndicator,omitempty"`

	AccountDate string `json:"accountDate,omitempty"`

	AccountPasswordChangeDate string `json:"accountPasswordChangeDate,omitempty"`

	AccountPasswordChangeIndicator string `json:"accountPasswordChangeIndicator,omitempty"`

	AccountPurchaseCount string `json:"accountPurchaseCount,omitempty"`

	AccountProvisioningAttempts string `json:"accountProvisioningAttempts,omitempty"`

	AccountDayTransactions string `json:"accountDayTransactions,omitempty"`

	AccountYearTransactions string `json:"accountYearTransactions,omitempty"`

	PaymentAccountAge string `json:"paymentAccountAge,omitempty"`

	PaymentAccountAgeIndicator string `json:"paymentAccountAgeIndicator,omitempty"`

	ShipAddressUsageDate string `json:"shipAddressUsageDate,omitempty"`

	ShipAddressUsageIndicator string `json:"shipAddressUsageIndicator,omitempty"`

	ShipNameIndicator string `json:"shipNameIndicator,omitempty"`

	SuspiciousAccountActivity string `json:"suspiciousAccountActivity,omitempty"`

	RequestorChallengeIndicator string `json:"requestorChallengeIndicator,omitempty"`

	DeliveryEmailAddress string `json:"deliveryEmailAddress,omitempty"`

	PreOrderPurchaseIndicator string `json:"preOrderPurchaseIndicator,omitempty"`

	ReorderItemsIndicator string `json:"reorderItemsIndicator,omitempty"`

	ShipIndicator string `json:"shipIndicator,omitempty"`

	MpiAcquirerCode string `json:"mpiAcquirerCode,omitempty"`

	ExSlipInfo string `json:"exSlipInfo,omitempty"`

	AddressMatchIndicator string `json:"addressMatchIndicator,omitempty"`
}
