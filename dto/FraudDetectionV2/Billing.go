package FraudDetectionV2

type Billing struct {
	HouseNumber1 string `json:"houseNumber1,omitempty"`

	Street1 string `json:"street1,omitempty"`

	Street2 string `json:"street2,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Country string `json:"country,omitempty"`

	Postcode string `json:"postcode,omitempty"`
}
