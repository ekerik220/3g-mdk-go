package PayNowId

type Response struct {
	ProcessId     string          `json:"processId"`
	Status        string          `json:"status"`
	Message       string          `json:"message"`
	Account       Account         `json:"account"`
	OptionResults []OptionResults `json:"optionResults"`
}

type OptionResults struct {
	Options []string `json:"options"`
}

type Account struct {
	AccountId       string            `json:"accountId"`
	TransData       string            `json:"transData"`
	AccountBasic    AccountBasic      `json:"accountBasic"`
	AccountInfo     []AccountInfo     `json:"accountInfo"`
	CardInfo        []CardInfo        `json:"cardInfo"`
	RecurringCharge []RecurringCharge `json:"recurringCharge"`
}

type AccountBasic struct {
	LastName      string `json:"lastName"`
	LastKanaName  string `json:"lastKanaName"`
	FirstName     string `json:"firstName"`
	FirstKanaName string `json:"firstKanaName"`
	MailAddress   string `json:"mailAddress"`
	CreateDate    string `json:"createDate"`
	DeleteDate    string `json:"deleteDate"`
}

type AccountInfo struct {
	AccountInfoId  string `json:"accountInfoId"`
	AccountType    string `json:"accountType"`
	DefaultAccount string `json:"defaultAccount"`
	Recipient      string `json:"recipient"`
	Zip            string `json:"zip"`
	Address        string `json:"address"`
	Tel            string `json:"tel"`
}

type CardInfo struct {
	CardId         string `json:"cardId"`
	CardNumber     string `json:"cardNumber"`
	CardExpire     string `json:"cardExpire"`
	CardholderName string `json:"cardholderName"`
	DefaultCard    string `json:"defaultCard"`
	AcquireCode3   string `json:"acquireCode3"`
	KindCode       string `json:"kindCode"`
	OriginalCardId string `json:"originalCardId"`
}

type RecurringCharge struct {
	GroupId string `json:"groupId"`
	CardId  string `json:"cardId"`
}
