package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/heriant0/financial-api/internal/app/schemas"
)

type UserServices interface {
	Register(req schemas.RegisterRequest) error
	Login(req schemas.LoginRequest) (*schemas.LoginResponse, error)
}

type UserController struct {
	service UserServices
}

func NewUserController(service UserServices) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Register(ctx *gin.Context) {
	var req schemas.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err = c.service.Register(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed register user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create user"})
}

func (c *UserController) Login(ctx *gin.Context) {
	var req schemas.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	data, err := c.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (c *UserController) Auth(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	accessToken := splitToken[1]

	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte("apa@jaB0l3h"), nil
	})

	if err != nil {
		fmt.Println("error", err)
		ctx.Abort()
		return
	}

	if !token.Valid {
		fmt.Println("invalid token")
		ctx.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expiredAt := int64(claims["exp_at"].(float64))
		fmt.Println(expiredAt, "expirate at")
		fmt.Println(time.Now().Unix(), "now")
		if time.Now().Unix() > expiredAt {
			fmt.Println("token expired")
			ctx.Abort()
			return
		}
	} else {
		fmt.Println("failed extract token")
		ctx.Abort()
		return
	}

}
