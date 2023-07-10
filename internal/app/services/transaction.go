package services

import (
	"fmt"
	"time"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/heriant0/financial-api/internal/app/schemas"
	"github.com/leekchan/accounting"
	log "github.com/sirupsen/logrus"
)

type TransactionRepository interface {
	Save(data models.Transaction) error
	GetList() ([]models.Transaction, error)
	GetListByType(tsType string) ([]models.Transaction, error)
}

type TransactionService struct {
	repository TransactionRepository
}

func NewTransactironService(transactionRepository TransactionRepository) *TransactionService {
	return &TransactionService{repository: transactionRepository}
}

func (s *TransactionService) Create(req schemas.TransactionCreateRequest) error {
	trxDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		log.Error(fmt.Errorf("invalid date - Create :%w", err))
		return err
	}

	data := models.Transaction{
		UserId:          req.UserId,
		CategoryId:      req.CategoryId,
		CurrencyId:      req.CurrencyId,
		TransactionType: req.Type,
		Note:            req.Note,
		Amount:          req.Amount,
		CreatedAt:       trxDate,
	}

	err = s.repository.Save(data)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionService - Create :%w", err))
		return err
	}

	return nil
}

func (s *TransactionService) GetList() ([]schemas.TransactionResponse, error) {
	var transactions []schemas.TransactionResponse

	data, err := s.repository.GetList()
	if err != nil {
		log.Error(fmt.Errorf("error TransactionService - GetList :%w", err))
		return nil, err
	}

	transactions = mapToTransactionResponse(data)
	return transactions, nil
}

func (s *TransactionService) GetListByType(types string) ([]schemas.TransactionResponse, error) {
	var transactions []schemas.TransactionResponse

	data, err := s.repository.GetListByType(types)

	if err != nil {
		log.Error(fmt.Errorf("error TransactionService - GetListByType :%w", err))
		return transactions, err
	}

	transactions = mapToTransactionResponse(data)

	return transactions, nil
}

func mapToTransactionResponse(transactions []models.Transaction) []schemas.TransactionResponse {
	var response []schemas.TransactionResponse

	for _, t := range transactions {
		transactionResponse := schemas.TransactionResponse{
			Types:    t.TransactionType,
			Category: t.CategoryName,
			Amount:   formatFinancial(t.Amount, t.CurrencyCode),
			Note:     t.Note,
			Date:     t.CreatedAt.Format("02-01-2006"),
		}

		response = append(response, transactionResponse)
	}

	return response
}

func formatFinancial(number float64, symbol string) string {
	ac := accounting.Accounting{Symbol: symbol, Precision: 2, Thousand: ".", Decimal: ","}
	formatted := ac.FormatMoney(number)
	return formatted
}
