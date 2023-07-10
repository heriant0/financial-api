package repositories

import (
	"errors"
	"fmt"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetList() ([]models.Category, error) {
	var (
		categories   []models.Category
		sqlStatement = "SELECT id, name, description FROM categories"
	)

	// DB Execution
	rows, err := r.DB.Queryx(sqlStatement)
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var category models.Category
		rows.StructScan(&category) // nolint:errcheck
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) Create(data models.Category) error {
	sqlStatement := `
		INSERT INTO	categories (name, description)
		VALUES ($1, $2)
	`
	_, err := r.DB.Exec(sqlStatement, data.Name, data.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) GetById(id int) (models.Category, error) {
	var (
		data         models.Category
		sqlStatement = `
			SELECT id, name, description
			FROM  categories
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

func (r *CategoryRepository) Update(category models.Category) error {
	sqlStatement := `
		UPDATE categories
		SET name = $1,
			description = $2,
			updated_at = NOW()
		WHERE id = $3
	`
	result, err := r.DB.Exec(sqlStatement, category.Name, category.Description, category.ID)
	if err != nil {
		log.Error(fmt.Errorf("error category respository : %w", err))
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r *CategoryRepository) Delete(id int) error {
	sqlStatemen := `
		Delete FROM categories
		WHERE id =$1
	`

	result, err := r.DB.Exec(sqlStatemen, id)

	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no recored affected")
	}
	return nil
}
