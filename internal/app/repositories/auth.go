package repositories

import (
	"fmt"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type AuthRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(DB *sqlx.DB) *AuthRepository {
	return &AuthRepository{DB: DB}
}

func (r *AuthRepository) Find(userId int, refreshToken string) (models.Auth, error) {
	var (
		auth         models.Auth
		sqlStatement = `
			SELECT id, token, auth_type, user_id, expired_at
			FROM auths
			WHERE user_id = $1 AND token =$2
		`
	)

	err := r.DB.QueryRowx(sqlStatement, userId, refreshToken).StructScan(&auth)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Find : %w", err))
		return auth, err
	}

	return auth, nil
}

func (r *AuthRepository) Create(auth models.Auth) error {
	var (
		sqlStatement = `
			INSERT INTO auths (token, auth_type, expired_at, user_id)
			VALUES ($1, $2, $3, $4)
		`
	)

	_, err := r.DB.Exec(sqlStatement, auth.Token, auth.AuthType, auth.ExpirateAt, auth.UserID)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Create : %w", err))
		return err
	}

	return nil
}

func (r *AuthRepository) DeleteAllByUserID(userId int) error {
	var (
		sqlStatement = `
			DELETE FROM auths
			WHERE user_id = $1
		`
	)

	_, err := r.DB.Exec(sqlStatement, userId)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - DeleteAllByUserID : %w", err))
		return err
	}
	return nil
}
