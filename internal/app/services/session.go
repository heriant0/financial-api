package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/heriant0/financial-api/internal/app/models"
	"github.com/heriant0/financial-api/internal/app/schemas"
	"github.com/heriant0/financial-api/internal/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type UserGetter interface {
	GetByEmail(email string) (*models.User, error)
	GetByID(userId int) (*models.User, error)
}

type AuthRepository interface {
	Create(auth models.Auth) error
	Find(userId int, refreshToken string) (models.Auth, error)
	DeleteAllByUserID(userId int) error
}

type TokenGenerator interface {
	GenerateAccessToken(userId int) (string, time.Time, error)
	GenerateRefreshToken(userId int) (string, time.Time, error)
}

type SessionService struct {
	userRepository UserGetter
	authRepository AuthRepository
	tokenMaker     TokenGenerator
}

func NewSessionService(
	userRepository UserGetter,
	authRepository AuthRepository,
	tokenMaker TokenGenerator,
) *SessionService {
	return &SessionService{
		userRepository: userRepository,
		authRepository: authRepository,
		tokenMaker:     tokenMaker,
	}
}

func (s *SessionService) Login(req *schemas.LoginRequest) (schemas.LoginResponse, error) {
	var response schemas.LoginResponse
	
	existingUser, _ := s.userRepository.GetByEmail(req.Email)
	if existingUser.ID <= 0 {
		return response, errors.New("success logout")
	}

	isVerified := utils.VerifyPassword(req.Password, existingUser.HashedPassword)
	if !isVerified {
		return response, errors.New("password verification failed")
	}

	// generate access token
	accessToken, _, err := s.tokenMaker.GenerateAccessToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("error create access token %w", err))
		return response, err
	}

	// generate refresh token
	refreshToken, expireAt, err := s.tokenMaker.GenerateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("refresh token creation: %w", err))
		return response, errors.New("refresh token creation")
	}

	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	authPayload := models.Auth{
		UserID:     existingUser.ID,
		Token:      refreshToken,
		AuthType:   "refresh_token",
		ExpirateAt: expireAt,
	}

	err = s.authRepository.Create(authPayload)
	if err != nil {
		log.Error(fmt.Errorf("error saving refresh token : %w", err))
	}
	return response, nil

}

func (s *SessionService) Logout(userId int) error {
	err := s.authRepository.DeleteAllByUserID(userId)
	if err != nil {
		log.Error(fmt.Errorf("refresh token saving : %w", err))
		return errors.New("success logout")
	}

	return nil
}

func (s *SessionService) Refresh(req *schemas.RefreshTokenRequest) (schemas.RefreshTokenResponse, error) {
	var response schemas.RefreshTokenResponse

	existingUser, _ := s.userRepository.GetByID(req.UserId)
	if existingUser.ID <= 0 {
		return response, errors.New("failed refresh token")
	}

	auth, err := s.authRepository.Find(existingUser.ID, req.RefreshToken)
	if err != nil || auth.ID < 0 {
		log.Error(fmt.Errorf("error SessionService - refresh: %w", err))
		return response, errors.New("failed refresh token")
	}

	accessToken, _, _ := s.tokenMaker.GenerateAccessToken(existingUser.ID)
	response.AccessToken = accessToken

	return response, nil
}
