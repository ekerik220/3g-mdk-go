package FraudDetectionV2

type RdResponse struct {
	Code                  string `json:"code"`
	RequestId             string `json:"requestId"`
	OrderId               string `json:"orderId"`
	RiskStatusCode        string `json:"riskStatusCode"`
	RiskFraudStatusCode   string `json:"riskFraudStatusCode"`
	RiskResponseCode      string `json:"riskResponseCode"`
	RiskOrderId           string `json:"riskOrderId"`
	RiskNeuralScore       string `json:"riskNeuralScore"`
	RiskRuleCategory      string `json:"riskRuleCategory"`
	Id                    string `json:"id"`
	Description           string `json:"description"`
	MerchantTransactionId string `json:"merchantTransactionId"`
	BuildNumber           string `json:"buildNumber"`
	Timestamp             string `json:"timestamp"`
	Ndc                   string `json:"ndc"`
	ResultDescription     string `json:"resultDescription"`
	ExternalSystemLink    string `json:"externalSystemLink"`
	Action                string `json:"action"`
	RiskReport            string `json:"riskReport"`
	RiskFraudDescription  string `json:"riskFraudDescription"`
	TransactionId         string `json:"transactionId"`
	ParameterErrors       string `json:"parameterErrors"`
	TargetServiceTypeCd   string `json:"targetServiceTypeCd"`
	TargetOrderId         string `json:"targetOrderId"`
}
