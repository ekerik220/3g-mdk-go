package ScoreAtPay

type Detail struct {
	// 明細名 例) コート
	DetailName     string `json:"detailName,omitempty"`
	DetailPrice    int64  `json:"detailPrice,omitempty"`
	DetailQuantity int64  `json:"detailQuantity,omitempty"`
}
