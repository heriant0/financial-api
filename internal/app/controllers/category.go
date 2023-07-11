package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/schemas"
	"github.com/heriant0/financial-api/internal/pkg/handler"
)

type CategoryServices interface {
	GetList() ([]schemas.CategoryListResponse, error)
	Detail(req schemas.CategoryDetailRequest) (schemas.CategoryDetailResponse, error)
	Create(req schemas.CategoryCreateRequest) error
	Update(req schemas.CategoryUpdateRequest) error
	Delete(req schemas.CategoryDeleteRequest) error
}

type CategoryController struct {
	categoryService CategoryServices
}

func NewCategoryController(service CategoryServices) *CategoryController {
	return &CategoryController{categoryService: service}
}

func (c *CategoryController) GetList(ctx *gin.Context) {
	response, err := c.categoryService.GetList()
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", response)
}

func (c *CategoryController) Create(ctx *gin.Context) {
	var req schemas.CategoryCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = c.categoryService.Create(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed create category")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create category", nil)
}

func (c *CategoryController) Detail(ctx *gin.Context) {
	categoryIDstr := ctx.Param("id")
	categoryId, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed convert category id")
	}

	req := schemas.CategoryDetailRequest{ID: categoryId}
	response, err := c.categoryService.Detail(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed get data detail")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "", response)
}

func (c *CategoryController) Update(ctx *gin.Context) {
	categoryIdStr := ctx.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed convert category id")
		return
	}

	req := schemas.CategoryUpdateRequest{}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "ffailed update data category")
		return
	}

	req.ID = categoryId
	err = c.categoryService.Update(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "ffailed update data category")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "data has been updated", nil)

}

func (c *CategoryController) Delete(ctx *gin.Context) {
	categoryIdStr := ctx.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed convert category id")
		return
	}

	req := schemas.CategoryDeleteRequest{ID: categoryId}
	err = c.categoryService.Delete(req)
	if err != nil {
		handler.ResponError(ctx, http.StatusUnprocessableEntity, "failed delete data category")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "data has been deleted"})
	handler.ResponseSuccess(ctx, http.StatusOK, "data has been deleted", nil)

}
