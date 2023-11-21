package LinePay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type AuthorizeRequest struct {
	TxnVersion         string          `json:"txnVersion,omitempty"`
	DummyRequest       string          `json:"dummyRequest,omitempty"`
	MerchantCcid       string          `json:"merchantCcid,omitempty"`
	PayNowIdParam      *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId            string          `json:"orderId,omitempty"`
	Amount             string          `json:"amount,omitempty"`
	WithCapture        string          `json:"withCapture,omitempty"`
	ItemId             string          `json:"itemId,omitempty"`
	ItemName           string          `json:"itemName,omitempty"`
	ItemImageUrl       string          `json:"itemImageUrl,omitempty"`
	CheckUseBrowser    string          `json:"checkUseBrowser,omitempty"`
	AppUrlScheme       string          `json:"appUrlScheme,omitempty"`
	UseOriginalApp     string          `json:"useOriginalApp,omitempty"`
	Mid                string          `json:"mid,omitempty"`
	PackageName        string          `json:"packageName,omitempty"`
	LangCd             string          `json:"langCd,omitempty"`
	SuccessUrl         string          `json:"successUrl,omitempty"`
	CancelUrl          string          `json:"cancelUrl,omitempty"`
	ErrorUrl           string          `json:"errorUrl,omitempty"`
	PushUrl            string          `json:"pushUrl,omitempty"`
	OneTimeKey         string          `json:"oneTimeKey,omitempty"`
	PaymentConfirmType string          `json:"paymentConfirmType,omitempty"`
}
