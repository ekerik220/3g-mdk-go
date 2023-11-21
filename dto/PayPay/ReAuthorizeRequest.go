package PayPay

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type ReAuthorizeRequest struct {
	TxnVersion                 string          `json:"txnVersion,omitempty"`
	DummyRequest               string          `json:"dummyRequest,omitempty"`
	MerchantCcid               string          `json:"merchantCcid,omitempty"`
	PayNowIdParam              *PayNowId.Param `json:"payNowIdParam,omitempty"`
	OrderId                    string          `json:"orderId,omitempty"`
	ServiceOptionType          string          `json:"serviceOptionType,omitempty"`
	OriginalOrderId            string          `json:"originalOrderId,omitempty"`
	Amount                     string          `json:"amount,omitempty"`
	ItemId                     string          `json:"itemId,omitempty"`
	ItemName                   string          `json:"itemName,omitempty"`
	NsfRecoveryFlag            string          `json:"nsfRecoveryFlag,omitempty"`
	NsfRecoveryExpiredDatetime string          `json:"nsfRecoveryExpiredDatetime,omitempty"`
	PushUrl                    string          `json:"pushUrl,omitempty"`
}
