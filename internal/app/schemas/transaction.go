package schemas

type TransactionCreateRequest struct {
	UserId     int     `json:"user_id"`
	CategoryId int     `json:"category_id" binding:"required"`
	CurrencyId int     `json:"currency_id" binding:"required"`
	Type       string  `json:"type" binding:"required"`
	Note       string  `json:"note" binding:"alphanum"`
	Amount     float64 `json:"amount" binding:"number,required"`
	Date       string  `json:"date" binding:"required,datetime=2006-01-02"`
}

type TransactionResponse struct {
	Types    string `json:"type"`
	Amount   string `json:"amount"`
	Category string `json:"category"`
	Note     string `json:"note"`
	Date     string `json:"date"`
}

type TransactionByTypeRequest struct {
	Types string `json:"type"`
}
