package repositories

import (
	"fmt"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) Save(data models.Transaction) error {
	var (
		sqlStatement = `
			INSERT INTO transactions (user_id, category_id, currency_id, transaction_type, note, amount, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
	)

	_, err := r.DB.Exec(sqlStatement, data.UserId, data.CategoryId, data.CurrencyId, data.TransactionType, data.Note, data.Amount, data.CreatedAt)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionRepository - Create :%w", err))
		return err
	}

	return nil
}

func (r *TransactionRepository) GetList(filter string) ([]models.Transaction, error) {
	var (
		transactions []models.Transaction
		sqlStatement = `
			SELECT
				t.transaction_type as type,
				c."name" as  category_name,
				cr.code as currency_code,
				t.amount as amount,
				t.note as note,
				t.created_at as created_at
			FROM transactions t
			INNER JOIN users u on t.user_id = u.id
			INNER JOIN categories c on t.category_id = c.id
			INNER JOIN currencies cr on t.currency_id = cr.id
		`
	)

	var rows *sqlx.Rows
	var err error
	
	if filter != "" {
		sqlStatement = sqlStatement + "WHERE t.transaction_type = $1"
		rows, err = r.DB.Queryx(sqlStatement, filter)
	} else {
		rows, err = r.DB.Queryx(sqlStatement)
	}

	// rows, err := r.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionRepository - GetList :%w", err))
		return transactions, err
	}

	for rows.Next() {
		var transaction models.Transaction
		err = rows.StructScan(&transaction)
		if err != nil {
			log.Error(fmt.Errorf("error GetList - GetList :%w", err))
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *TransactionRepository) GetListByType(types string) ([]models.Transaction, error) {
	var (
		transactions []models.Transaction
		sqlStatement = `
			SELECT 
				t.transaction_type,
				c."name",
				cr.code,
				t.amount,
				t.note,
				t.created_at 
			FROM transactions t
				INNER JOIN users u on t.user_id = u.id
				INNER JOIN categories c on t.category_id = c.id
				INNER JOIN currencies cr on t.currency_id = cr.id
			WHERE t.transaction_type = $1
		`
	)

	rows, err := r.DB.Queryx(sqlStatement, types)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionRepository - GetListByType :%w", err))
		return transactions, err
	}

	for rows.Next() {
		var transaction models.Transaction
		rows.StructScan(&transaction)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
