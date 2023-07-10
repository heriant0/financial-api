package repositories

import (
	"errors"
	"fmt"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user models.User) error {
	sqlStatement := `
		INSERT INTO users (username, email, hashed_password)
		VALUES ($1, $2, $3)
	`

	_, err := r.DB.Exec(sqlStatement, user.Username, user.Email, user.HashedPassword)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Create: %w", err))
		return err
	}

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var (
		data         models.User
		sqlStatement = `
			SELECT id, username, email, hashed_password
			FROM users
			WHERE email = $1
			LIMIT 1 
		`
	)
	err := r.DB.QueryRowx(sqlStatement, email).StructScan(&data)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByEmail: %w", err))
		return &data, err
	}

	if data.ID == 0 {
		return nil, errors.New("data not found")
	}

	return &data, nil
}

func (r *UserRepository) GetByEmailAndUsername(email string, username string) (*models.User, error) {
	var (
		data         models.User
		sqlStatement = `
			SELECT id, username, email
			FROM users
			WHERE email = $1 AND username = $2
			LIMIT 1 
		`
	)

	err := r.DB.QueryRowx(sqlStatement, email, username).StructScan(&data)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByEmailAndUsername: %w", err))
		return &data, err
	}

	if data.ID == 0 {
		return nil, errors.New("data not found")
	}

	return &data, nil
}

func (r *UserRepository) GetByID(userID int) (*models.User, error) {
	var (
		data         models.User
		sqlStatement = `
			SELECT id, username, email
			FROM users
			WHERE id = $1
			LIMIT 1 
		`
	)

	err := r.DB.QueryRowx(sqlStatement, userID).StructScan(&data)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByID: %w", err))
		return &data, err
	}

	if data.ID == 0 {
		return nil, errors.New("data not found")
	}

	return &data, nil
}

func (r *UserRepository) Update(user models.User) error {
	var (
		sqlStatement = `
			UPDATE users 
			SET updated_at = NOW(),
				username = $2,
				email = $3
			WHERE id = $1
		`
	)
	result, err := r.DB.Exec(sqlStatement, user.ID, user.Username, user.Email)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Update: %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r *UserRepository) Delete(userId int) error {
	var (
		sqlStatement = `
			DELETE FROM users 
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, userId)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Delete: %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r *UserRepository) GetDataUser(id int) (models.UserProfile, error) {
	var (
		data         models.UserProfile
		sqlStatement = `
			SELECT id, username, email, created_at
			FROM users
			WHERE id = $1
			LIMIT 1 
		`
	)

	err := r.DB.QueryRowx(sqlStatement, id).StructScan(&data)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetDataUser: %w", err))
		return data, err
	}

	return data, nil
}
