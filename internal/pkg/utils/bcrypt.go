package utils

import (
	"github.com/heriant0/financial-api/internal/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var cfg config.Config

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cfg.HashCost)

	return string(bytes)
}

func VerifyPassword(plainPassword string, hanshedPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hanshedPassword),
		[]byte(plainPassword),
	)
	return err == nil
}
