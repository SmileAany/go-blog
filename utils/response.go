package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func GetSuccessResponse(c *gin.Context, response *Response) *gin.Context {
	status := response.Status

	if status == "" {
		status = "success"
	}

	code := response.Code

	if code == 0 {
		code = 200
	}

	c.JSON(code, gin.H{
		"data":    response.Data,
		"status":  status,
		"code":    code,
		"message": response.Message,
		"errors":  response.Errors,
	})

	return c
}

func GetErrorResponse(c *gin.Context, response *Response) *gin.Context {
	status := response.Status

	if status == "" {
		status = "error"
	}

	code := response.Code

	if code == 0 {
		code = 500
	}

	c.JSON(code, gin.H{
		"data":    response.Data,
		"status":  status,
		"code":    code,
		"message": response.Message,
		"errors":  response.Errors,
	})

	return c
}
