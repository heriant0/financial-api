package repositories

import (
	"errors"
	"fmt"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CurrencyRepository struct {
	DB *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{DB: db}
}

func (r *CurrencyRepository) GetList() ([]models.Currency, error) {
	var (
		currencies   []models.Currency
		sqlStatement = `SELECT id, name, code FROM currencies`
	)

	rows, err := r.DB.Queryx(sqlStatement)
	if err != nil {
		return currencies, err
	}

	for rows.Next() {
		var currency models.Currency
		rows.StructScan(&currency)
		currencies = append(currencies, currency)
	}
	
	return currencies, nil
}

func (r *CurrencyRepository) GetByID(id int) (models.Currency, error) {
	var (
		data         models.Currency
		sqlStatement = `
			SELECT id, name, code
			FROM currencies
			WHERE id = $1
			LIMIT 1
		`
	)

	err := r.DB.QueryRowx(sqlStatement, id).StructScan(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *CurrencyRepository) Create(currency models.Currency) error {
	sqlStatement := `
		INSERT INTO currencies (name, code)
		VALUES ($1, $2)
	`
	_, err := r.DB.Exec(sqlStatement, currency.Name, currency.Code)

	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - Create : %w", err))
		return err
	}

	return nil
}

func (r *CurrencyRepository) Update(currency models.Currency) error {
	sqlStatement := `
		UPDATE currencies 
		SET name = $1, 
			code = $2,
			updated_at = NOW()
		WHERE id = $3
	`
	result, err := r.DB.Exec(sqlStatement, currency.Name, currency.Code, currency.ID)

	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - UpdateByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no recored affected")
	}

	return nil
}

func (r *CurrencyRepository) DeleteByID(id int) error {
	sqlStatement := `
		DELETE FROM currencies
		WHERE id = $1
	`
	result, err := r.DB.Exec(sqlStatement, id)

	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no recored affected")
	}

	return nil
}
