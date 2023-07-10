package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

// var secretKey = []byte(cfg.SecretKey)

func BuildJWT(userId int, secretKey string, duration time.Duration) (string, error) {
	exp := time.Now().Add(duration)
	// key := []byte(secretKey)
	// claims := jwt.MapClaims{
	// 	"email":     email,
	// 	"exp_at":    time.Now().Add(24 * time.Hour).Unix(),
	// 	"issued_at": time.Now().Unix(),
	// }

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userId),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, err := token.SignedString(secretKey)
	if err != nil {
		errMsg := fmt.Errorf("user not found : %w", err)
		log.Error(errMsg)
		return "", err
	}
	return stringToken, nil
}

func VerifyJWT(accessToken string, secretKey []byte) error {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expiredAt := int64(claims["exp_at"].(float64))
		if time.Now().Unix() > expiredAt {
			return errors.New("token alrady expired")
		}
	} else {
		return errors.New("failed to extract token")
	}

	return nil
}
