package jwt

type Type string

const(
	SUCCESS Type = "success"

	InvalidParams = "token 缺失"

	TokenFail = "token 异常"

	TokenTimeout = "token 超时"
)