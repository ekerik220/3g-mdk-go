package AmazonPay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type GetAddressResponse struct {
	PayNowIdResponse       PayNowId.Response `json:"payNowIdResponse"`
	ServiceType            string            `json:"serviceType"`
	Mstatus                string            `json:"mstatus"`
	VResultCode            string            `json:"vResultCode"`
	MerrMsg                string            `json:"merrMsg"`
	OrderId                string            `json:"orderId"`
	TxnVersion             string            `json:"txnVersion"`
	CenterResponseDatetime string            `json:"centerResponseDatetime"`
	ShippingName           string            `json:"shippingName"`
	ShippingPhone          string            `json:"shippingPhone"`
	ShippingAddress1       string            `json:"shippingAddress1"`
	ShippingAddress2       string            `json:"shippingAddress2"`
	ShippingAddress3       string            `json:"shippingAddress3"`
	ShippingPrefecture     string            `json:"shippingPrefecture"`
	ShippingPostalCode     string            `json:"shippingPostalCode"`
	BuyerId                string            `json:"buyerId"`
	BuyerEmail             string            `json:"buyerEmail"`
	BuyerName              string            `json:"buyerName"`
	BuyerPhoneNumber       string            `json:"buyerPhoneNumber"`
	PaymentPreferences     string            `json:"paymentPreferences"`
	BillingAddress1        string            `json:"billingAddress1"`
	BillingAddress2        string            `json:"billingAddress2"`
	BillingAddress3        string            `json:"billingAddress3"`
	BillingPrefecture      string            `json:"billingPrefecture"`
	BillingPostalCode      string            `json:"billingPostalCode"`
	PayRedirectUrl         string            `json:"payRedirectUrl"`
	CheckoutSessionId      string            `json:"checkoutSessionId"`
	ResultJson             string            `json:"resultJson"`
}
