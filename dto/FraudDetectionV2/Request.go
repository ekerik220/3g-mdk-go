package FraudDetectionV2

type Request struct {
	FraudDetectionOnly string `json:"fraudDetectionOnly,omitempty"`

	FraudDetectionMode string `json:"fraudDetectionMode,omitempty"`

	TargetServiceTypeCd string `json:"targetServiceTypeCd,omitempty"`

	TargetOrderId string `json:"targetOrderId,omitempty"`

	MerchantTransactionId string `json:"merchantTransactionId,omitempty"`

	TransactionCategory string `json:"transactionCategory,omitempty"`

	Amount string `json:"amount,omitempty"`

	Currency string `json:"currency,omitempty"`

	PaymentBrand string `json:"paymentBrand,omitempty"`

	Card *Card `json:"card,omitempty"`

	BankAccount *BankAccount `json:"bankAccount,omitempty"`

	TaxAmount string `json:"taxAmount,omitempty"`

	Risk *Risk `json:"risk,omitempty"`

	ThreeDSecure *ThreeDSecure `json:"threeDSecure,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	Billing *Billing `json:"billing,omitempty"`

	Shipping *Shipping `json:"shipping,omitempty"`

	Corporate *Corporate `json:"corporate,omitempty"`

	Cart *Cart `json:"cart,omitempty"`

	GiftCard *GiftCard `json:"giftCard,omitempty"`

	Airline *Airline `json:"airline,omitempty"`
}
