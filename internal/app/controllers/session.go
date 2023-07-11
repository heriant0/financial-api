package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/schemas"
	"github.com/heriant0/financial-api/internal/pkg/handler"
)

type SessionService interface {
	Login(req *schemas.LoginRequest) (schemas.LoginResponse, error)
	Logout(UserId int) error
	Refresh(req *schemas.RefreshTokenRequest) (schemas.RefreshTokenResponse, error)
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	sessionService SessionService
	tokenMaker     RefreshTokenVerifier
}

func NewSessionController(service SessionService, tokenMaker RefreshTokenVerifier) *SessionController {
	return &SessionController{
		sessionService: service,
		tokenMaker:     tokenMaker,
	}
}

func (c *SessionController) Login(ctx *gin.Context) {
	req := &schemas.LoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response, err := c.sessionService.Login(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "login success", response)

}

func (c *SessionController) Refresh(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("refresh_token")

	if refreshToken == "" {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed get data refresh token")
		return
	}

	sub, err := c.tokenMaker.VerifyRefreshToken(refreshToken)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed verify refresh token")

		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := &schemas.RefreshTokenRequest{
		RefreshToken: refreshToken,
		UserId:       intSub,
	}

	response, err := c.sessionService.Refresh(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed refresh token")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", response)

}

func (c *SessionController) Logout(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.GetString("user_id"))
	err := c.sessionService.Logout(userId)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success logout"})
	handler.ResponseSuccess(ctx, http.StatusOK, "success logout", nil)

}
