package FraudDetectionV2

type Corporate struct {
	Address string `json:"address,omitempty"`

	City string `json:"city,omitempty"`

	Description string `json:"description,omitempty"`

	Fax string `json:"fax,omitempty"`

	Name string `json:"name,omitempty"`

	Phone string `json:"phone,omitempty"`

	PostboxNumber string `json:"postboxNumber,omitempty"`

	PostCode string `json:"postCode,omitempty"`

	Purchase string `json:"purchase,omitempty"`

	State string `json:"state,omitempty"`

	Street string `json:"street,omitempty"`

	Country string `json:"country,omitempty"`

	Suite string `json:"suite,omitempty"`
}
