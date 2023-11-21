package FraudDetectionV2

type Risk struct {
	ServiceId string `json:"serviceId,omitempty"`

	Amount string `json:"amount,omitempty"`

	OrderTimestamp string `json:"orderTimestamp,omitempty"`

	MerchantWebsite string `json:"merchantWebsite,omitempty"`

	AccountToken string `json:"accountToken,omitempty"`

	Parameters *Parameters `json:"parameters,omitempty"`
}
