package PayNowId

type Param struct {
	AccountParam *AccountParam `json:"accountParam,omitempty"`
	ChargeParam  *ChargeParam  `json:"chargeParam,omitempty"`
	OrderParam   *OrderParam   `json:"orderParam,omitempty"`
	Tanking      string        `json:"tanking,omitempty"`
	FreeKey      string        `json:"freeKey,omitempty"`
	Memo1        string        `json:"memo1,omitempty"`
	ReceiptData  string        `json:"receiptData,omitempty"`
	Token        string        `json:"token,omitempty"`
}

type AccountParam struct {
	AccountId            string                `json:"accountId,omitempty"`
	OnetimeToken         string                `json:"onetimeToken,omitempty"`
	OnetimeTokenType     string                `json:"onetimeTokenType,omitempty"`
	OriginalAccountId    string                `json:"originalAccountId,omitempty"`
	PayNowId             string                `json:"payNowId,omitempty"`
	TransData            string                `json:"transData,omitempty"`
	AccountBasicParam    *AccountBasicParam    `json:"accountBasicParam,omitempty"`
	CardParam            *CardParam            `json:"cardParam,omitempty"`
	RecurringChargeParam *RecurringChargeParam `json:"recurringChargeParam,omitempty"`
}

type AccountBasicParam struct {
	CreateDate              string `json:"createDate,omitempty"`
	DeleteDate              string `json:"deleteDate,omitempty"`
	ForceDeleteDate         string `json:"forceDeleteDate,omitempty"`
	DeleteCardInfo          string `json:"deleteCardInfo,omitempty"`
	AddCardInfo             string `json:"addCardInfo,omitempty"`
	DeleteOriginalAccountId string `json:"deleteOriginalAccountId,omitempty"`
	DefaultCardOnly         string `json:"defaultCardOnly,omitempty"`
	CleaningConfig          string `json:"cleaningConfig,omitempty"`
}

type CardParam struct {
	CardId             string `json:"cardId,omitempty"`
	DefaultCard        string `json:"defaultCard,omitempty"`
	DefaultCardId      string `json:"defaultCardId,omitempty"`
	Updater            string `json:"updater,omitempty"`
	Token              string `json:"token,omitempty"`
	CardNumberMaskType string `json:"cardNumberMaskType,omitempty"`
	WithAuthorize      string `json:"withAuthorize,omitempty"`
}

type RecurringChargeParam struct {
	GroupId         string `json:"groupId,omitempty"`
	StartDate       string `json:"startDate,omitempty"`
	EndDate         string `json:"endDate,omitempty"`
	FinalCharge     string `json:"finalCharge,omitempty"`
	OneTimeAmount   string `json:"oneTimeAmount,omitempty"`
	Amount          string `json:"amount,omitempty"`
	RecurringMemo1  string `json:"recurringMemo1,omitempty"`
	RecurringMemo2  string `json:"recurringMemo2,omitempty"`
	RecurringMemo3  string `json:"recurringMemo3,omitempty"`
	UseChargeOption string `json:"useChargeOption,omitempty"`
	SalesDay        string `json:"salesDay,omitempty"`
	AcquireCode     string `json:"acquireCode,omitempty"`
}

type ChargeParam struct {
	GroupId       string `json:"groupId,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
	Type          string `json:"type,omitempty"`
	OneTimeAmount string `json:"oneTimeAmount,omitempty"`
	Amount        string `json:"amount,omitempty"`
	ChargeType    string `json:"chargeType,omitempty"`
	Schedule      string `json:"schedule,omitempty"`
	SalesDay      string `json:"salesDay,omitempty"`
	AcquireCode   string `json:"acquireCode,omitempty"`
}

type OrderParam struct {
	OriginalOrderId    string `json:"originalOrderId,omitempty"`
	CleaningMerchantId string `json:"cleaningMerchantId,omitempty"`
}
