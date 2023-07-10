package models

type Currency struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Code         string `db:"code"`
}
