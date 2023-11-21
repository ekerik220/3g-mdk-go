package FraudDetectionV2

type Airline struct {
	ThirdPartyBooking string `json:"thirdPartyBooking,omitempty"`

	AgentCode string `json:"agentCode,omitempty"`

	AgentName string `json:"agentName,omitempty"`

	BookingType string `json:"bookingType,omitempty"`

	BookingRefNum string `json:"bookingRefNum,omitempty"`

	TicketDeliveryMethod string `json:"ticketDeliveryMethod,omitempty"`

	TicketIssueAddress string `json:"ticketIssueAddress,omitempty"`

	TicketIssueDate string `json:"ticketIssueDate,omitempty"`

	TotalTaxAmount string `json:"totalTaxAmount,omitempty"`

	TotalFareAmount string `json:"totalFareAmount,omitempty"`

	TotalFeesAmount string `json:"totalFeesAmount,omitempty"`

	Passengers *[]Passenger `json:"passengers,omitempty"`
}
