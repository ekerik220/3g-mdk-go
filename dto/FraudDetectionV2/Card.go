package FraudDetectionV2

type Card struct {
	Number string `json:"number,omitempty"`

	ExpiryMonth string `json:"expiryMonth,omitempty"`

	ExpiryYear string `json:"expiryYear,omitempty"`
}
