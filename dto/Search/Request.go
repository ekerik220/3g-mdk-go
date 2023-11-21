package Search

type Request struct {
	TxnVersion   string `json:"txnVersion,omitempty"`
	DummyRequest string `json:"dummyRequest,omitempty"`
	MerchantCcid string `json:"merchantCcid,omitempty"`

	RequestId        string      `json:"requestId,omitempty"`
	ServiceTypeCd    []string    `json:"serviceTypeCd,omitempty"`
	NewerFlag        string      `json:"newerFlag,omitempty"`
	ContainDummyFlag string      `json:"containDummyFlag,omitempty"`
	MaxCount         string      `json:"maxCount,omitempty"`
	MasterNames      []string    `json:"masterNames,omitempty"`
	SearchParameters *Parameters `json:"searchParameters,omitempty"`
}
