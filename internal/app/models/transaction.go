package models

import "time"

type Transaction struct {
	ID              int       `db:"id"`
	UserId          int       `db:"user_id"`
	CategoryId      int       `db:"category_id"`
	CategoryName    string    `db:"category_name"`
	CurrencyId      int       `db:"currency_id"`
	CurrencyCode    string    `db:"currency_code"`
	TransactionType string    `db:"type"`
	Note            string    `db:"note"`
	Amount          float64   `db:"amount"`
	CreatedAt       time.Time `db:"created_at"`
}
