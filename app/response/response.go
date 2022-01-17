package response

import (
	"github.com/gin-gonic/gin"
)

// Response 统计返回的数据格式
type Response struct {
	Code int
	Status string
	Message string
	Data interface{}
	Errors interface{}
}

// SetCode 设置code
func (r *Response) SetCode(code int) *Response {
	r.Code = code

	return r
}

// SetMessage 设置message
func (r *Response) SetMessage(message string) *Response {
	r.Message = message

	return r
}

// SetErrors 设置Errors
func (r *Response) SetErrors(errors map[string]string) *Response {
	r.Errors = errors

	return r
}

// SetData 设置data
func (r *Response) SetData(data interface{}) *Response {
	r.Data = data

	return r
}

// ResponseSuccess 成功返回
func (r *Response) ResponseSuccess(c *gin.Context) {
	if r.Code == 0 {
		r.Code = 200
	}

	r.Status = "success"

	if r.Message == "" {
		r.Message = "success"
	}

	if r.Errors == nil {
		r.Errors = []string{}
	}

	if r.Data == nil {
		r.Data = []string{}
	}

	c.JSON(r.Code,gin.H{
		"message" : r.Message,
		"code"    : r.Code,
		"status"  : r.Status,
		"data"    : r.Data,
		"errors"  : r.Errors,
	})
}

// ResponseError 失败返回
func (r *Response) ResponseError (c *gin.Context) {
	if r.Code == 0 {
		r.Code = 400
	}

	r.Status = "error"

	if r.Message == "" {
		r.Message = "error"
	}

	if r.Data == nil {
		r.Data = []string{}
	}


	if r.Errors == nil {
		r.Errors = []string{}
	}

	c.JSON(r.Code,gin.H{
		"message" : r.Message,
		"code"    : r.Code,
		"status"  : r.Status,
		"data"    : r.Data,
		"errors"  : r.Errors,
	})
}