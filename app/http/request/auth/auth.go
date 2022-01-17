package auth

import "crm/app/http/request"

type AccountPasswordLoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (accountPasswordLoginRequest AccountPasswordLoginRequest) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Username.required": "请输入账号",
		"Password.required": "请输入密码",
	}
}

type CheckAccountRequest struct {
	Column string `form:"column" binding:"required"`
	Value string `form:"value" binding:"required"`
}

func (checkAccountRequest CheckAccountRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Column.required": "请输入字段类型",
		"Value.required": "请输入字段值",
	}
}

type UserRegisterRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email string `form:"email" binding:"required"`
	Name string  `form:"name" binding:"required"`
	Phone string `form:"phone" binding:"required"`
}

func (userRegisterRequest UserRegisterRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Username.required": "请输入账号",
		"Password.required": "请输入密码",
		"Email.required": "请输入邮箱",
		"Name.required": "请输入昵称",
		"Phone.required": "请输入手机号",
	}
}

type RegisterCheckRequest struct {
	Code string `form:"code" binding:"required"`
	Email string `form:"email" binding:"required"`
	UserId int `form:"userId" binding:"required"`
}

func (registerCheckRequest RegisterCheckRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Email.required": "请输入邮箱",
		"Code.required": "请输入验证码",
		"UserId.required" : "请传递userId",
	}
}

type GetCodeRequest struct {
	Phone string `form:"phone" binding:"required"`
}

func (getCodeRequest GetCodeRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Phone.required": "请输入手机号",
	}
}

type PhoneLoginRequest struct {
	Phone string `form:"phone" binding:"required"`
	Code string `form:"code" binding:"required"`
}

func (phoneLoginRequest PhoneLoginRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Phone.required": "请输入手机号",
		"Code.required" : "请输入验证码",
	}
}