package models

import "time"

type Auth struct {
	ID         int       `db:"id"`
	Token      string    `db:"token"`
	AuthType   string    `db:"auth_type"`
	ExpirateAt time.Time `db:"expired_at"`
	UserID     int       `db:"user_id"`
}
