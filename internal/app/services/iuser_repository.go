package services

import "github.com/heriant0/financial-api/internal/app/models"

type UserRepository interface {
	Create(user models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByEmailAndUsername(email string, username string) (*models.User, error)
	GetByID(userId int) (*models.User, error)
	GetDataUser(userId int) (models.UserProfile, error)
	Update(user models.User) error
	Delete(userId int) error
}
