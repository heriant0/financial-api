package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/schemas"
)

type TransactionService interface {
	Create(req schemas.TransactionCreateRequest) error
	GetList() ([]schemas.TransactionResponse, error)
	GetListByType(types string) ([]schemas.TransactionResponse, error)
}

type TransactionController struct {
	service TransactionService
}

func NewTransactionController(transactionService TransactionService) *TransactionController {
	return &TransactionController{service: transactionService}
}

func (c *TransactionController) Create(ctx *gin.Context) {
	var req schemas.TransactionCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed create transaction"})
		return
	}

	err = c.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed create transaction"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "data has been saved"})
}

func (c *TransactionController) GetList(ctx *gin.Context) {
	response, err := c.service.GetList()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})

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
