package FraudDetectionV2

type Customer struct {
	MerchantCustomerId string `json:"merchantCustomerId,omitempty"`

	GivenName string `json:"givenName,omitempty"`

	MiddleName string `json:"middleName,omitempty"`

	Surname string `json:"surname,omitempty"`

	Sex string `json:"sex,omitempty"`

	WorkPhone string `json:"workPhone,omitempty"`

	BirthDate string `json:"birthDate,omitempty"`

	Phone string `json:"phone,omitempty"`

	Mobile string `json:"mobile,omitempty"`

	Email string `json:"email,omitempty"`

	CompanyName string `json:"companyName,omitempty"`

	Ip string `json:"ip,omitempty"`

	BrowserFingerprint *BrowserFingerprint `json:"browserFingerprint,omitempty"`

	IdentificationDocId string `json:"identificationDocId,omitempty"`

	Status string `json:"status,omitempty"`
}
