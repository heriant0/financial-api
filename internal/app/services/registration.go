package services

import (
	"errors"
	"fmt"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/heriant0/financial-api/internal/app/schemas"
	"github.com/heriant0/financial-api/internal/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type RegistrationService struct {
	userRepository UserRepository
}

func NewRegistrationService(repository UserRepository) *RegistrationService {
	return &RegistrationService{
		userRepository: repository,
	}
}

func (s *RegistrationService) Register(req *schemas.RegisterRequest) error {
	// check email and username
	isExists, _ := s.userRepository.GetByEmailAndUsername(req.Email, req.Username)

	if isExists.ID > 0 {
		return errors.New("username and email" + req.Email + "already exists")
	}

	hashedPassword := utils.HashPassword(req.Password)
	data := models.User{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	err := s.userRepository.Create(data)
	if err != nil {
		errMsg := fmt.Errorf("error registration - RegistrationService : %w", err)
		log.Error(errMsg)
		return err
	}

	return nil
}

func (s *RegistrationService) GetProfileUser(req schemas.UserProfileRequest) (schemas.UserProfielResponse, error) {
	var response schemas.UserProfielResponse

	user, err := s.userRepository.GetDataUser(req.ID)
	if err != nil {
		errMsg := fmt.Errorf("error service - GetProfileUser : %w", err)
		log.Error(errMsg)
		return response, err
	}

	response.ID = user.ID
	response.Email = user.Email
	response.Username = user.Username
	response.Usersince = user.CreatedAt.Format("02-01-2006")

	return response, nil
}
