package FraudDetectionV2

type Passenger struct {
	CheckDigit string `json:"checkDigit,omitempty"`

	Dob string `json:"dob,omitempty"`

	Email string `json:"email,omitempty"`

	FrequentFlyerNumber string `json:"frequentFlyerNumber,omitempty"`

	Name string `json:"name,omitempty"`

	Phone string `json:"phone,omitempty"`

	Status string `json:"status,omitempty"`

	TicketNumber string `json:"ticketNumber,omitempty"`

	TicketRestricted string `json:"ticketRestricted,omitempty"`

	Type string `json:"type,omitempty"`

	Legs1AirlineName string `json:"legs1AirlineName,omitempty"`

	Legs1AirlineCode string `json:"legs1AirlineCode,omitempty"`

	Legs1ArrivalAirport string `json:"legs1ArrivalAirport,omitempty"`

	Legs1ArrivalCountry string `json:"legs1ArrivalCountry,omitempty"`

	Legs1ArrivalDate string `json:"legs1ArrivalDate,omitempty"`

	Legs1ArrivalTime string `json:"legs1ArrivalTime,omitempty"`

	Legs1CarrierCode string `json:"legs1CarrierCode,omitempty"`

	Legs1ClassOfService string `json:"legs1ClassOfService,omitempty"`

	Legs1CouponNumber string `json:"legs1CouponNumber,omitempty"`

	Legs1DepartureAirport string `json:"legs1DepartureAirport,omitempty"`

	Legs1DepartureCountry string `json:"legs1DepartureCountry,omitempty"`

	Legs1DepartureDate string `json:"legs1DepartureDate,omitempty"`

	Legs1DepartureTime string `json:"legs1DepartureTime,omitempty"`

	Legs1DepartureTaxAmount string `json:"legs1DepartureTaxAmount,omitempty"`

	Legs1ExchangeTicketNum string `json:"legs1ExchangeTicketNum,omitempty"`

	Legs1FareAmount string `json:"legs1FareAmount,omitempty"`

	Legs1FareBasis string `json:"legs1FareBasis,omitempty"`

	Legs1FeesAmount string `json:"legs1FeesAmount,omitempty"`

	Legs1FlightNumber string `json:"legs1FlightNumber,omitempty"`

	Legs1Restrictions string `json:"legs1Restrictions,omitempty"`

	Legs1StopOverAllowed string `json:"legs1StopOverAllowed,omitempty"`

	Legs1TaxAmount string `json:"legs1TaxAmount,omitempty"`

	Legs1TicketNumber string `json:"legs1TicketNumber,omitempty"`

	Legs2AirlineName string `json:"legs2AirlineName,omitempty"`

	Legs2AirlineCode string `json:"legs2AirlineCode,omitempty"`

	Legs2ArrivalAirport string `json:"legs2ArrivalAirport,omitempty"`

	Legs2ArrivalCountry string `json:"legs2ArrivalCountry,omitempty"`

	Legs2ArrivalDate string `json:"legs2ArrivalDate,omitempty"`

	Legs2ArrivalTime string `json:"legs2ArrivalTime,omitempty"`

	Legs2CarrierCode string `json:"legs2CarrierCode,omitempty"`

	Legs2ClassOfService string `json:"legs2ClassOfService,omitempty"`

	Legs2CouponNumber string `json:"legs2CouponNumber,omitempty"`

	Legs2DepartureAirport string `json:"legs2DepartureAirport,omitempty"`

	Legs2DepartureCountry string `json:"legs2DepartureCountry,omitempty"`

	Legs2DepartureDate string `json:"legs2DepartureDate,omitempty"`

	Legs2DepartureTime string `json:"legs2DepartureTime,omitempty"`

	Legs2DepartureTaxAmount string `json:"legs2DepartureTaxAmount,omitempty"`

	Legs2ExchangeTicketNum string `json:"legs2ExchangeTicketNum,omitempty"`

	Legs2FareAmount string `json:"legs2FareAmount,omitempty"`

	Legs2FareBasis string `json:"legs2FareBasis,omitempty"`

	Legs2FeesAmount string `json:"legs2FeesAmount,omitempty"`

	Legs2FlightNumber string `json:"legs2FlightNumber,omitempty"`

	Legs2Restrictions string `json:"legs2Restrictions,omitempty"`

	Legs2StopOverAllowed string `json:"legs2StopOverAllowed,omitempty"`

	Legs2TaxAmount string `json:"legs2TaxAmount,omitempty"`

	Legs2TicketNumber string `json:"legs2TicketNumber,omitempty"`

	Legs3AirlineName string `json:"legs3AirlineName,omitempty"`

	Legs3AirlineCode string `json:"legs3AirlineCode,omitempty"`

	Legs3ArrivalAirport string `json:"legs3ArrivalAirport,omitempty"`

	Legs3ArrivalCountry string `json:"legs3ArrivalCountry,omitempty"`

	Legs3ArrivalDate string `json:"legs3ArrivalDate,omitempty"`

	Legs3ArrivalTime string `json:"legs3ArrivalTime,omitempty"`

	Legs3CarrierCode string `json:"legs3CarrierCode,omitempty"`

	Legs3ClassOfService string `json:"legs3ClassOfService,omitempty"`

	Legs3CouponNumber string `json:"legs3CouponNumber,omitempty"`

	Legs3DepartureAirport string `json:"legs3DepartureAirport,omitempty"`

	Legs3DepartureCountry string `json:"legs3DepartureCountry,omitempty"`

	Legs3DepartureDate string `json:"legs3DepartureDate,omitempty"`

	Legs3DepartureTime string `json:"legs3DepartureTime,omitempty"`

	Legs3DepartureTaxAmount string `json:"legs3DepartureTaxAmount,omitempty"`

	Legs3ExchangeTicketNum string `json:"legs3ExchangeTicketNum,omitempty"`

	Legs3FareAmount string `json:"legs3FareAmount,omitempty"`

	Legs3FareBasis string `json:"legs3FareBasis,omitempty"`

	Legs3FeesAmount string `json:"legs3FeesAmount,omitempty"`

	Legs3FlightNumber string `json:"legs3FlightNumber,omitempty"`

	Legs3Restrictions string `json:"legs3Restrictions,omitempty"`

	Legs3StopOverAllowed string `json:"legs3StopOverAllowed,omitempty"`

	Legs3TaxAmount string `json:"legs3TaxAmount,omitempty"`

	Legs3TicketNumber string `json:"legs3TicketNumber,omitempty"`

	Legs4AirlineName string `json:"legs4AirlineName,omitempty"`

	Legs4AirlineCode string `json:"legs4AirlineCode,omitempty"`

	Legs4ArrivalAirport string `json:"legs4ArrivalAirport,omitempty"`

	Legs4ArrivalCountry string `json:"legs4ArrivalCountry,omitempty"`

	Legs4ArrivalDate string `json:"legs4ArrivalDate,omitempty"`

	Legs4ArrivalTime string `json:"legs4ArrivalTime,omitempty"`

	Legs4CarrierCode string `json:"legs4CarrierCode,omitempty"`

	Legs4ClassOfService string `json:"legs4ClassOfService,omitempty"`

	Legs4CouponNumber string `json:"legs4CouponNumber,omitempty"`

	Legs4DepartureAirport string `json:"legs4DepartureAirport,omitempty"`

	Legs4DepartureCountry string `json:"legs4DepartureCountry,omitempty"`

	Legs4DepartureDate string `json:"legs4DepartureDate,omitempty"`

	Legs4DepartureTime string `json:"legs4DepartureTime,omitempty"`

	Legs4DepartureTaxAmount string `json:"legs4DepartureTaxAmount,omitempty"`

	Legs4ExchangeTicketNum string `json:"legs4ExchangeTicketNum,omitempty"`

	Legs4FareAmount string `json:"legs4FareAmount,omitempty"`

	Legs4FareBasis string `json:"legs4FareBasis,omitempty"`

	Legs4FeesAmount string `json:"legs4FeesAmount,omitempty"`

	Legs4FlightNumber string `json:"legs4FlightNumber,omitempty"`

	Legs4Restrictions string `json:"legs4Restrictions,omitempty"`

	Legs4StopOverAllowed string `json:"legs4StopOverAllowed,omitempty"`

	Legs4TaxAmount string `json:"legs4TaxAmount,omitempty"`

	Legs4TicketNumber string `json:"legs4TicketNumber,omitempty"`

	Legs5AirlineName string `json:"legs5AirlineName,omitempty"`

	Legs5AirlineCode string `json:"legs5AirlineCode,omitempty"`

	Legs5ArrivalAirport string `json:"legs5ArrivalAirport,omitempty"`

	Legs5ArrivalCountry string `json:"legs5ArrivalCountry,omitempty"`

	Legs5ArrivalDate string `json:"legs5ArrivalDate,omitempty"`

	Legs5ArrivalTime string `json:"legs5ArrivalTime,omitempty"`

	Legs5CarrierCode string `json:"legs5CarrierCode,omitempty"`

	Legs5ClassOfService string `json:"legs5ClassOfService,omitempty"`

	Legs5CouponNumber string `json:"legs5CouponNumber,omitempty"`

	Legs5DepartureAirport string `json:"legs5DepartureAirport,omitempty"`

	Legs5DepartureCountry string `json:"legs5DepartureCountry,omitempty"`

	Legs5DepartureDate string `json:"legs5DepartureDate,omitempty"`

	Legs5DepartureTime string `json:"legs5DepartureTime,omitempty"`

	Legs5DepartureTaxAmount string `json:"legs5DepartureTaxAmount,omitempty"`

	Legs5ExchangeTicketNum string `json:"legs5ExchangeTicketNum,omitempty"`

	Legs5FareAmount string `json:"legs5FareAmount,omitempty"`

	Legs5FareBasis string `json:"legs5FareBasis,omitempty"`

	Legs5FeesAmount string `json:"legs5FeesAmount,omitempty"`

	Legs5FlightNumber string `json:"legs5FlightNumber,omitempty"`

	Legs5Restrictions string `json:"legs5Restrictions,omitempty"`

	Legs5StopOverAllowed string `json:"legs5StopOverAllowed,omitempty"`

	Legs5TaxAmount string `json:"legs5TaxAmount,omitempty"`

	Legs5TicketNumber string `json:"legs5TicketNumber,omitempty"`

	Legs6AirlineName string `json:"legs6AirlineName,omitempty"`

	Legs6AirlineCode string `json:"legs6AirlineCode,omitempty"`

	Legs6ArrivalAirport string `json:"legs6ArrivalAirport,omitempty"`

	Legs6ArrivalCountry string `json:"legs6ArrivalCountry,omitempty"`

	Legs6ArrivalDate string `json:"legs6ArrivalDate,omitempty"`

	Legs6ArrivalTime string `json:"legs6ArrivalTime,omitempty"`

	Legs6CarrierCode string `json:"legs6CarrierCode,omitempty"`

	Legs6ClassOfService string `json:"legs6ClassOfService,omitempty"`

	Legs6CouponNumber string `json:"legs6CouponNumber,omitempty"`

	Legs6DepartureAirport string `json:"legs6DepartureAirport,omitempty"`

	Legs6DepartureCountry string `json:"legs6DepartureCountry,omitempty"`

	Legs6DepartureDate string `json:"legs6DepartureDate,omitempty"`

	Legs6DepartureTime string `json:"legs6DepartureTime,omitempty"`

	Legs6DepartureTaxAmount string `json:"legs6DepartureTaxAmount,omitempty"`

	Legs6ExchangeTicketNum string `json:"legs6ExchangeTicketNum,omitempty"`

	Legs6FareAmount string `json:"legs6FareAmount,omitempty"`

	Legs6FareBasis string `json:"legs6FareBasis,omitempty"`

	Legs6FeesAmount string `json:"legs6FeesAmount,omitempty"`

	Legs6FlightNumber string `json:"legs6FlightNumber,omitempty"`

	Legs6Restrictions string `json:"legs6Restrictions,omitempty"`

	Legs6StopOverAllowed string `json:"legs6StopOverAllowed,omitempty"`

	Legs6TaxAmount string `json:"legs6TaxAmount,omitempty"`

	Legs6TicketNumber string `json:"legs6TicketNumber,omitempty"`

	Legs7AirlineName string `json:"legs7AirlineName,omitempty"`

	Legs7AirlineCode string `json:"legs7AirlineCode,omitempty"`

	Legs7ArrivalAirport string `json:"legs7ArrivalAirport,omitempty"`

	Legs7ArrivalCountry string `json:"legs7ArrivalCountry,omitempty"`

	Legs7ArrivalDate string `json:"legs7ArrivalDate,omitempty"`

	Legs7ArrivalTime string `json:"legs7ArrivalTime,omitempty"`

	Legs7CarrierCode string `json:"legs7CarrierCode,omitempty"`

	Legs7ClassOfService string `json:"legs7ClassOfService,omitempty"`

	Legs7CouponNumber string `json:"legs7CouponNumber,omitempty"`

	Legs7DepartureAirport string `json:"legs7DepartureAirport,omitempty"`

	Legs7DepartureCountry string `json:"legs7DepartureCountry,omitempty"`

	Legs7DepartureDate string `json:"legs7DepartureDate,omitempty"`

	Legs7DepartureTime string `json:"legs7DepartureTime,omitempty"`

	Legs7DepartureTaxAmount string `json:"legs7DepartureTaxAmount,omitempty"`

	Legs7ExchangeTicketNum string `json:"legs7ExchangeTicketNum,omitempty"`

	Legs7FareAmount string `json:"legs7FareAmount,omitempty"`

	Legs7FareBasis string `json:"legs7FareBasis,omitempty"`

	Legs7FeesAmount string `json:"legs7FeesAmount,omitempty"`

	Legs7FlightNumber string `json:"legs7FlightNumber,omitempty"`

	Legs7Restrictions string `json:"legs7Restrictions,omitempty"`

	Legs7StopOverAllowed string `json:"legs7StopOverAllowed,omitempty"`

	Legs7TaxAmount string `json:"legs7TaxAmount,omitempty"`

	Legs7TicketNumber string `json:"legs7TicketNumber,omitempty"`

	Legs8AirlineName string `json:"legs8AirlineName,omitempty"`

	Legs8AirlineCode string `json:"legs8AirlineCode,omitempty"`

	Legs8ArrivalAirport string `json:"legs8ArrivalAirport,omitempty"`

	Legs8ArrivalCountry string `json:"legs8ArrivalCountry,omitempty"`

	Legs8ArrivalDate string `json:"legs8ArrivalDate,omitempty"`

	Legs8ArrivalTime string `json:"legs8ArrivalTime,omitempty"`

	Legs8CarrierCode string `json:"legs8CarrierCode,omitempty"`

	Legs8ClassOfService string `json:"legs8ClassOfService,omitempty"`

	Legs8CouponNumber string `json:"legs8CouponNumber,omitempty"`

	Legs8DepartureAirport string `json:"legs8DepartureAirport,omitempty"`

	Legs8DepartureCountry string `json:"legs8DepartureCountry,omitempty"`

	Legs8DepartureDate string `json:"legs8DepartureDate,omitempty"`

	Legs8DepartureTime string `json:"legs8DepartureTime,omitempty"`

	Legs8DepartureTaxAmount string `json:"legs8DepartureTaxAmount,omitempty"`

	Legs8ExchangeTicketNum string `json:"legs8ExchangeTicketNum,omitempty"`

	Legs8FareAmount string `json:"legs8FareAmount,omitempty"`

	Legs8FareBasis string `json:"legs8FareBasis,omitempty"`

	Legs8FeesAmount string `json:"legs8FeesAmount,omitempty"`

	Legs8FlightNumber string `json:"legs8FlightNumber,omitempty"`

	Legs8Restrictions string `json:"legs8Restrictions,omitempty"`

	Legs8StopOverAllowed string `json:"legs8StopOverAllowed,omitempty"`

	Legs8TaxAmount string `json:"legs8TaxAmount,omitempty"`

	Legs8TicketNumber string `json:"legs8TicketNumber,omitempty"`

	Legs9AirlineName string `json:"legs9AirlineName,omitempty"`

	Legs9AirlineCode string `json:"legs9AirlineCode,omitempty"`

	Legs9ArrivalAirport string `json:"legs9ArrivalAirport,omitempty"`

	Legs9ArrivalCountry string `json:"legs9ArrivalCountry,omitempty"`

	Legs9ArrivalDate string `json:"legs9ArrivalDate,omitempty"`

	Legs9ArrivalTime string `json:"legs9ArrivalTime,omitempty"`

	Legs9CarrierCode string `json:"legs9CarrierCode,omitempty"`

	Legs9ClassOfService string `json:"legs9ClassOfService,omitempty"`

	Legs9CouponNumber string `json:"legs9CouponNumber,omitempty"`

	Legs9DepartureAirport string `json:"legs9DepartureAirport,omitempty"`

	Legs9DepartureCountry string `json:"legs9DepartureCountry,omitempty"`

	Legs9DepartureDate string `json:"legs9DepartureDate,omitempty"`

	Legs9DepartureTime string `json:"legs9DepartureTime,omitempty"`

	Legs9DepartureTaxAmount string `json:"legs9DepartureTaxAmount,omitempty"`

	Legs9ExchangeTicketNum string `json:"legs9ExchangeTicketNum,omitempty"`

	Legs9FareAmount string `json:"legs9FareAmount,omitempty"`

	Legs9FareBasis string `json:"legs9FareBasis,omitempty"`

	Legs9FeesAmount string `json:"legs9FeesAmount,omitempty"`

	Legs9FlightNumber string `json:"legs9FlightNumber,omitempty"`

	Legs9Restrictions string `json:"legs9Restrictions,omitempty"`

	Legs9StopOverAllowed string `json:"legs9StopOverAllowed,omitempty"`

	Legs9TaxAmount string `json:"legs9TaxAmount,omitempty"`

	Legs9TicketNumber string `json:"legs9TicketNumber,omitempty"`

	Legs10AirlineName string `json:"legs10AirlineName,omitempty"`

	Legs10AirlineCode string `json:"legs10AirlineCode,omitempty"`

	Legs10ArrivalAirport string `json:"legs10ArrivalAirport,omitempty"`

	Legs10ArrivalCountry string `json:"legs10ArrivalCountry,omitempty"`

	Legs10ArrivalDate string `json:"legs10ArrivalDate,omitempty"`

	Legs10ArrivalTime string `json:"legs10ArrivalTime,omitempty"`

	Legs10CarrierCode string `json:"legs10CarrierCode,omitempty"`

	Legs10ClassOfService string `json:"legs10ClassOfService,omitempty"`

	Legs10CouponNumber string `json:"legs10CouponNumber,omitempty"`

	Legs10DepartureAirport string `json:"legs10DepartureAirport,omitempty"`

	Legs10DepartureCountry string `json:"legs10DepartureCountry,omitempty"`

	Legs10DepartureDate string `json:"legs10DepartureDate,omitempty"`

	Legs10DepartureTime string `json:"legs10DepartureTime,omitempty"`

	Legs10DepartureTaxAmount string `json:"legs10DepartureTaxAmount,omitempty"`

	Legs10ExchangeTicketNum string `json:"legs10ExchangeTicketNum,omitempty"`

	Legs10FareAmount string `json:"legs10FareAmount,omitempty"`

	Legs10FareBasis string `json:"legs10FareBasis,omitempty"`

	Legs10FeesAmount string `json:"legs10FeesAmount,omitempty"`

	Legs10FlightNumber string `json:"legs10FlightNumber,omitempty"`

	Legs10Restrictions string `json:"legs10Restrictions,omitempty"`

	Legs10StopOverAllowed string `json:"legs10StopOverAllowed,omitempty"`

	Legs10TaxAmount string `json:"legs10TaxAmount,omitempty"`

	Legs10TicketNumber string `json:"legs10TicketNumber,omitempty"`
}
