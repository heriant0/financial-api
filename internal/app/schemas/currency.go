package schemas

type CurrencyListReponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code "`
}

type CurrencyCreateRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CurrencyUpdateRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CurrencyDeleteRequest struct {
	ID int `json:"id"`
}

type CurrencyDetailRequest struct {
	ID int `json:"id"`
}
