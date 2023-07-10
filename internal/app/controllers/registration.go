package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/schemas"
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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = c.registerService.Register(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "register successfully"})

}

func (c *RegistrationController) UserProfile(ctx *gin.Context) {
	userIdStr := ctx.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed get profile user"})
	}

	req := schemas.UserProfileRequest{ID: userId}
	res, err := c.registerService.GetProfileUser(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed get profile user"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
