package FraudDetectionV2

type Item struct {
	Quantity string `json:"quantity,omitempty"`

	PartNumber string `json:"partNumber,omitempty"`

	ProductCode string `json:"productCode,omitempty"`

	Sku string `json:"sku,omitempty"`

	MerchantItemId string `json:"merchantItemId,omitempty"`

	Description string `json:"description,omitempty"`

	OriginalPrice string `json:"originalPrice,omitempty"`

	ShippingTrackingNumber string `json:"shippingTrackingNumber,omitempty"`

	Name string `json:"name,omitempty"`

	Discount string `json:"discount,omitempty"`

	GiftMessage string `json:"giftMessage,omitempty"`

	ShippingInstructions string `json:"shippingInstructions,omitempty"`

	ShippingMethod string `json:"shippingMethod,omitempty"`

	Type string `json:"type,omitempty"`

	RecipientAddress string `json:"recipientAddress,omitempty"`

	RecipientApartment string `json:"recipientApartment,omitempty"`

	RecipientCity string `json:"recipientCity,omitempty"`

	RecipientCountry string `json:"recipientCountry,omitempty"`

	RecipientEmail string `json:"recipientEmail,omitempty"`

	RecipientFirstName string `json:"recipientFirstName,omitempty"`

	RecipientLastName string `json:"recipientLastName,omitempty"`

	RecipientMiddleName string `json:"recipientMiddleName,omitempty"`

	RecipientPhone string `json:"recipientPhone,omitempty"`

	RecipientPostcode string `json:"recipientPostcode,omitempty"`

	RecipientSalutation string `json:"recipientSalutation,omitempty"`

	RecipientState string `json:"recipientState,omitempty"`

	RecipientStreet string `json:"recipientStreet,omitempty"`
}
