package handler

import "github.com/gin-gonic/gin"

type ResponBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponError(ctx *gin.Context, statusCode int, message string) {
	response := ResponBody{
		Status:  "error",
		Message: message,
	}

	ctx.JSON(statusCode, response)
}

func ResponseSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := ResponBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	ctx.JSON(statusCode, response)
}
