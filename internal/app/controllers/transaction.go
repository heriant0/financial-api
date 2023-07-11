package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/schemas"
	"github.com/heriant0/financial-api/internal/pkg/handler"
)

type TransactionService interface {
	Create(req schemas.TransactionCreateRequest) error
	GetList(filter string) ([]schemas.TransactionResponse, error)
	GetListByType(types string) ([]schemas.TransactionResponse, error)
}

type TransactionController struct {
	service TransactionService
}

func NewTransactionController(transactionService TransactionService) *TransactionController {
	return &TransactionController{service: transactionService}
}

func (c *TransactionController) Create(ctx *gin.Context) {
	// Get user id
	userIdstr := ctx.GetString("user_id")
	userId, err := strconv.Atoi(userIdstr)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	var request schemas.TransactionCreateRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed create transaction")
		return
	}

	request.UserId = userId
	err = c.service.Create(request)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed create transaction")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "data has been saved", nil)
}

func (c *TransactionController) GetList(ctx *gin.Context) {
	filter := ctx.Query("type")
	response, err := c.service.GetList(filter)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "Successfully get data  transaction", response)

}

func (c *TransactionController) GetDataByType(ctx *gin.Context) {
	types := ctx.Param("types")
	response, err := c.service.GetListByType(types)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
