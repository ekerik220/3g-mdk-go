package Account

import (
	"github.com/ekerik220/3g-mdk-go/dto/PayNowId"
)

type AddRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`
}

type DeleteRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`
}

type LinkRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`
}

type RestoreRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`
}

type UpdateRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`
}

type CopyRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`
}

type TokenRequest struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	PayNowIdParam *PayNowId.Param `json:"payNowIdParam,omitempty"`
}
