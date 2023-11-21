package AmazonPay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type AuthorizeRequest struct {
	TxnVersion                  string          `json:"txnVersion,omitempty"`
	DummyRequest                string          `json:"dummyRequest,omitempty"`
	MerchantCcid                string          `json:"merchantCcid,omitempty"`
	PayNowIdParam               *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId                     string          `json:"orderId,omitempty"`
	AccountingType              string          `json:"accountingType,omitempty"`
	ConsentAuthType             string          `json:"consentAuthType,omitempty"`
	Amount                      string          `json:"amount,omitempty"`
	WithCapture                 string          `json:"withCapture,omitempty"`
	SuppressShippingAddressView string          `json:"suppressShippingAddressView,omitempty"`
	NoteToBuyer                 string          `json:"noteToBuyer,omitempty"`
	SuccessUrl                  string          `json:"successUrl,omitempty"`
	CancelUrl                   string          `json:"cancelUrl,omitempty"`
	ErrorUrl                    string          `json:"errorUrl,omitempty"`
	AuthorizePushUrl            string          `json:"authorizePushUrl,omitempty"`
	CapturePushUrl              string          `json:"capturePushUrl,omitempty"`
	CancelPushUrl               string          `json:"cancelPushUrl,omitempty"`
	FrequencyUnit               string          `json:"frequencyUnit,omitempty"`
	FrequencyValue              string          `json:"frequencyValue,omitempty"`
	AddressRestrictions         string          `json:"addressRestrictions,omitempty"`
	PayConfirmScreenType        string          `json:"payConfirmScreenType,omitempty"`
	PayConfirmScreenUrl         string          `json:"payConfirmScreenUrl,omitempty"`
	ExtendExpiration            string          `json:"extendExpiration,omitempty"`
}
