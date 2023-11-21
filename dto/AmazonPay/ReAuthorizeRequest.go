package AmazonPay

import (
	"3g-mdk-go/dto/PayNowId"
)

type ReAuthorizeRequest struct {
	TxnVersion       string          `json:"txnVersion,omitempty"`
	DummyRequest     string          `json:"dummyRequest,omitempty"`
	MerchantCcid     string          `json:"merchantCcid,omitempty"`
	PayNowIdParam    *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId          string          `json:"orderId,omitempty"`
	OriginalOrderId  string          `json:"originalOrderId,omitempty"`
	Amount           string          `json:"amount,omitempty"`
	WithCapture      string          `json:"withCapture,omitempty"`
	NoteToBuyer      string          `json:"noteToBuyer,omitempty"`
	AuthorizePushUrl string          `json:"authorizePushUrl,omitempty"`
	CapturePushUrl   string          `json:"capturePushUrl,omitempty"`
	CancelPushUrl    string          `json:"cancelPushUrl,omitempty"`
}
