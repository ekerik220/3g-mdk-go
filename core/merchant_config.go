package core

import "time"

type MerchantConfig struct {
	Host                   string
	Timeout                time.Duration
	MerchantCcId           string
	MerchantSecretKey      string
	DummyRequest           bool
	ParameterMaskingEnable bool
}

func NewMerchantConfig(merchantCcId string, merchantSecretKey string) *MerchantConfig {
	return &MerchantConfig{
		Host:                   "https://api.veritrans.co.jp",
		Timeout:                60 * time.Second,
		MerchantCcId:           merchantCcId,
		MerchantSecretKey:      merchantSecretKey,
		DummyRequest:           false,
		ParameterMaskingEnable: true,
	}
}
