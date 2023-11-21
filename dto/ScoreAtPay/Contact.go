package ScoreAtPay

type Contact struct {
	FullName     string `json:"fullName,omitempty"`
	FullKanaName string `json:"fullKanaName,omitempty"`

	// ハイフンは省略可能です。利用する場合は必ず2つ指定してください。 例) 03-9999-9999
	//  先頭は必ず 0 とします。
	Tel    string `json:"tel,omitempty"`
	Mobile string `json:"mobile,omitempty"`

	Email       string `json:"email,omitempty"`
	MobileEmail string `json:"mobileEmail,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`

	// 都道府県
	Address1 string `json:"address1,omitempty"`

	// 市区町村
	Address2 string `json:"address2,omitempty"`

	// 市区町村以降
	Address3       string `json:"address3,omitempty"`
	CompanyName    string `json:"companyName,omitempty"`
	DepartmentName string `json:"departmentName,omitempty"`
}
