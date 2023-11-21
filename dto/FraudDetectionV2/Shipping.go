package FraudDetectionV2

type Shipping struct {
	Customer *Customer `json:"customer,omitempty"`

	HouseNumber1 string `json:"houseNumber1,omitempty"`

	Street1 string `json:"street1,omitempty"`

	Street2 string `json:"street2,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Country string `json:"country,omitempty"`

	Postcode string `json:"postcode,omitempty"`

	Cost string `json:"cost,omitempty"`

	TrackingNumber string `json:"trackingNumber,omitempty"`

	Comment string `json:"comment,omitempty"`

	LogisticsProvider string `json:"logisticsProvider,omitempty"`

	Method string `json:"method,omitempty"`
}
