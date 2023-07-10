package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/schemas"
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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})

}

func (c *CategoryController) Create(ctx *gin.Context) {
	var req schemas.CategoryCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err = c.categoryService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed create category"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create category"})
}

func (c *CategoryController) Detail(ctx *gin.Context) {
	categoryIDstr := ctx.Param("id")
	categoryId, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed get data detail"})
	}

	req := schemas.CategoryDetailRequest{ID: categoryId}
	response, err := c.categoryService.Detail(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed get data detail"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})

}

func (c *CategoryController) Update(ctx *gin.Context) {
	categoryIdStr := ctx.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed update data category"})
		return
	}

	req := schemas.CategoryUpdateRequest{}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed update data category"})
		return
	}

	req.ID = categoryId
	err = c.categoryService.Update(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed update data category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "data has been updated"})

}

func (c *CategoryController) Delete(ctx *gin.Context) {
	categoryIdStr := ctx.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed delete data category"})
		return
	}

	req := schemas.CategoryDeleteRequest{ID: categoryId}
	err = c.categoryService.Delete(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "failed delete data category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "data has been deleted"})

}
