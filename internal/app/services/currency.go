package services

import (
	"fmt"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/heriant0/financial-api/internal/app/schemas"
	log "github.com/sirupsen/logrus"
)

type CurrencyRepository interface {
	GetList() ([]models.Currency, error)
	GetByID(id int) (models.Currency, error)
	Create(currency models.Currency) error
	Update(currency models.Currency) error
	DeleteByID(id int) error
}

type CurrencyService struct {
	currencyRepository CurrencyRepository
}

func NewCurrencyService(repository CurrencyRepository) *CurrencyService {
	return &CurrencyService{currencyRepository: repository}
}

func (s *CurrencyService) GetList() ([]schemas.CurrencyListReponse, error) {
	var response []schemas.CurrencyListReponse

	data, err := s.currencyRepository.GetList()
	if err != nil {
		return nil, err
	}

	for _, value := range data {
		var resp schemas.CurrencyListReponse

		resp.ID = value.ID
		resp.Name = value.Name
		resp.Code = value.Code
		response = append(response, resp)
	}

	return response, nil
}

func (s *CurrencyService) GetByID(req schemas.CurrencyDetailRequest) (schemas.CurrencyListReponse, error) {
	var response schemas.CurrencyListReponse

	data, err := s.currencyRepository.GetByID(req.ID)
	if err != nil {
		return response, err
	}

	response.ID = data.ID
	response.Name = data.Name
	response.Code = data.Code

	return response, nil
}

func (s *CurrencyService) Create(req schemas.CurrencyCreateRequest) error {
	data := models.Currency{
		Name: req.Name,
		Code: req.Code,
	}

	err := s.currencyRepository.Create(data)
	if err != nil {
		errMessage := fmt.Errorf("currency service - err create: %w", err)
		log.Error(errMessage)
		return err
	}
	return nil
}

func (s *CurrencyService) Update(req schemas.CurrencyUpdateRequest) error {
	var updateData models.Currency

	data, err := s.currencyRepository.GetByID(req.ID)
	if err != nil {
		return err
	}

	updateData.ID = data.ID
	updateData.Name = req.Name
	updateData.Code = req.Code

	err = s.currencyRepository.Update(updateData)
	if err != nil {
		return err
	}

	return nil
}

func (s *CurrencyService) DeleteByID(req schemas.CurrencyDeleteRequest) error {
	_, err := s.currencyRepository.GetByID(req.ID)
	if err != nil {
		return err
	}

	err = s.currencyRepository.DeleteByID(req.ID)
	if err != nil {
		return err
	}
	return nil
}
