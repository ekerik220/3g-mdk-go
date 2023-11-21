package ScoreAtPay

type Delivery struct {
	DeliveryId string   `json:"deliveryId,omitempty"`
	Contact    Contact  `json:"contact,omitempty"`
	Details    []Detail `json:"details,omitempty"`
}
