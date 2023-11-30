package ScoreAtPay

type CaptureRequest struct {
	OrderId       string `json:"orderId,omitempty"`
	PdCompanyCode string `json:"pdCompanyCode,omitempty"`
	SlipNo        string `json:"slipNo,omitempty"`
	DeliveryId    string `json:"deliveryId,omitempty"`
}
