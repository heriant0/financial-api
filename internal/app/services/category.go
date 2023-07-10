package services

import (
	"fmt"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/heriant0/financial-api/internal/app/schemas"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository interface {
	GetList() ([]models.Category, error)
	GetById(id int) (models.Category, error)
	Create(data models.Category) error
	Update(data models.Category) error
	Delete(id int) error
}

type CategoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(repository CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: repository}
}

func (s *CategoryService) GetList() ([]schemas.CategoryListResponse, error) {
	var response []schemas.CategoryListResponse

	data, err := s.categoryRepository.GetList()
	if err != nil {
		return response, err
	}

	for _, value := range data {
		var resp schemas.CategoryListResponse
		resp.ID = value.ID
		resp.Name = value.Name
		resp.Description = value.Description
		response = append(response, resp)
	}
	return response, nil
}

func (s *CategoryService) Detail(req schemas.CategoryDetailRequest) (schemas.CategoryDetailResponse, error) {
	var response schemas.CategoryDetailResponse

	data, err := s.categoryRepository.GetById(req.ID)
	if err != nil {
		errMsg := fmt.Errorf("category service - err detail : %w", err)
		log.Error(errMsg)
		return response, err
	}

	response.ID = data.ID
	response.Name = data.Name
	response.Description = data.Description

	return response, nil
}

func (s *CategoryService) Create(req schemas.CategoryCreateRequest) error {
	data := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.categoryRepository.Create(data)
	if err != nil {
		errMsg := fmt.Errorf("category service - err create : %w", err)
		log.Error(errMsg)
		return err
	}

	return nil
}

func (s *CategoryService) Update(req schemas.CategoryUpdateRequest) error {
	var updateData models.Category

	data, err := s.categoryRepository.GetById(req.ID)
	if err != nil {
		return err
	}

	updateData.ID = data.ID
	updateData.Name = req.Name
	updateData.Description = req.Description
	err = s.categoryRepository.Update(updateData)
	if err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) Delete(req schemas.CategoryDeleteRequest) error {
	_, err := s.categoryRepository.GetById(req.ID)
	if err != nil {
		return err
	}

	err = s.categoryRepository.Delete(req.ID)
	if err != nil {
		return err
	}

	return nil
}
