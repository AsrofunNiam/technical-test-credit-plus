package domain

import "github.com/AsrofunNiam/technical-test-credit-plus/model/web"

type Currencies []Currency
type Currency struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(15);not null"`
}

func (currency *Currency) ToCurrencyResponse() web.CurrencyResponse {
	return web.CurrencyResponse{
		// Required Fields
		ID:   currency.ID,
		Name: currency.Name,
	}
}

func (users Currencies) ToCurrencyResponses() []web.CurrencyResponse {
	currencyResponses := []web.CurrencyResponse{}
	for _, user := range users {
		currencyResponses = append(currencyResponses, user.ToCurrencyResponse())
	}
	return currencyResponses
}
