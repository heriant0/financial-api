package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

type TokenMaker struct {
	AccessTokenKey       string
	RefreshTokenKey      string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

func NewTokenMaker(
	accessTokenKey string,
	refreshTokenKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) *TokenMaker {
	return &TokenMaker{
		AccessTokenKey:       accessTokenKey,
		RefreshTokenKey:      refreshTokenKey,
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
	}
}

func (maker *TokenMaker) GenerateAccessToken(userId int) (string, time.Time, error) {
	exp := time.Now().Add(maker.AccessTokenDuration)
	key := []byte(maker.AccessTokenKey)
	// key := []byte("testaccesesstoken")
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userId),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Error(fmt.Errorf("error signed token %w", err))
		return "", exp, err

	}

	return tokenString, exp, nil
}

func (maker *TokenMaker) GenerateRefreshToken(userId int) (string, time.Time, error) {
	exp := time.Now().Add(maker.RefreshTokenDuration)
	key := []byte(maker.RefreshTokenKey)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userId),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", exp, err
	}

	return tokenString, exp, nil
}

func (maker *TokenMaker) VerifyAccessToken(tokenString string) (string, error) {
	sub, err := maker.verify(tokenString, maker.AccessTokenKey)
	return sub, err
}

func (maker *TokenMaker) VerifyRefreshToken(tokenString string) (string, error) {
	sub, err := maker.verify(tokenString, maker.RefreshTokenKey)
	return sub, err
}

func (maker *TokenMaker) verify(tokenString string, tokenKey string) (string, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenKey), nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		log.Error(fmt.Errorf("error parsing token %w", err))
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		sub := fmt.Sprint(claims["sub"])
		return sub, nil
	}

	return "", err
}
