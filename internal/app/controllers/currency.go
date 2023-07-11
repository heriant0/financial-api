package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/schemas"
	"github.com/heriant0/financial-api/internal/pkg/handler"
)

type CurrencyService interface {
	GetList() ([]schemas.CurrencyListReponse, error)
	GetByID(req schemas.CurrencyDetailRequest) (schemas.CurrencyListReponse, error)
	Create(req schemas.CurrencyCreateRequest) error
	Update(req schemas.CurrencyUpdateRequest) error
	DeleteByID(req schemas.CurrencyDeleteRequest) error
}

type CurrencyController struct {
	currencyService CurrencyService
}

func NewCurrencyController(service CurrencyService) *CurrencyController {
	return &CurrencyController{currencyService: service}
}

func (c *CurrencyController) GetList(ctx *gin.Context) {
	response, err := c.currencyService.GetList()
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", response)
}

func (c *CurrencyController) GetByID(ctx *gin.Context) {
	currencyIdStr := ctx.Param("id")
	currencyId, err := strconv.Atoi(currencyIdStr)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed conver currency id")
		return
	}

	req := schemas.CurrencyDetailRequest{ID: currencyId}
	response, err := c.currencyService.GetByID(req)

	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed get data detail")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", response)

}

func (c *CurrencyController) Create(ctx *gin.Context) {
	var req schemas.CurrencyCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = c.currencyService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed create currency"})
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success create currency", nil)

}

func (c *CurrencyController) Update(ctx *gin.Context) {
	currencyIdStr := ctx.Param("id")
	currencyId, err := strconv.Atoi(currencyIdStr)

	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed convert id currency")
		return
	}

	req := schemas.CurrencyUpdateRequest{}
	err = ctx.ShouldBindJSON(&req)

	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed update data currency")
		return
	}

	req.ID = currencyId

	err = c.currencyService.Update(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed update data currency")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "data has been updated", nil)
}

func (c *CurrencyController) Delete(ctx *gin.Context) {
	currencyIdStr := ctx.Param("id")
	currencyId, err := strconv.Atoi(currencyIdStr)

	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "ailed convert id currency")
		return
	}

	req := schemas.CurrencyDeleteRequest{ID: currencyId}
	err = c.currencyService.DeleteByID(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "ailed convert id currency")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "ata has been deleted", nil)
}
