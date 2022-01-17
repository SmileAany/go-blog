package request

import (
	"github.com/go-playground/validator/v10"
)

// Validator 验证器接口
type Validator interface {
	// GetMessages GetMessage 获取验证器自定义错误信息
	GetMessages() ValidatorMessages
}

// ValidatorMessages 验证器自定义错误信息字典
type ValidatorMessages map[string]string

// GetErrorMessage 获取自定义错误信息
func GetErrorMessage(request Validator, err error) string {
	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := request.GetMessages()[v.Field() + "." + v.Tag()]; exist {
			return message
		}
		return v.Error()
	}
	return "Parameter error"
}