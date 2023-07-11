package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/schemas"
	"github.com/heriant0/financial-api/internal/pkg/handler"
)

type RegisterService interface {
	Register(req *schemas.RegisterRequest) error
	GetProfileUser(req schemas.UserProfileRequest) (schemas.UserProfielResponse, error)
}

type RegistrationController struct {
	registerService RegisterService
}

func NewRegistrationConroller(service RegisterService) *RegistrationController {
	return &RegistrationController{registerService: service}
}

func (c *RegistrationController) Register(ctx *gin.Context) {
	req := &schemas.RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = c.registerService.Register(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "register successfully", nil)
}

func (c *RegistrationController) UserProfile(ctx *gin.Context) {
	// userIdStr := ctx.Param("id")
	userIdStr := ctx.GetString("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed get profile user")
		return
	}

	req := schemas.UserProfileRequest{ID: userId}
	response, err := c.registerService.GetProfileUser(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed get profile user")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", response)
}
