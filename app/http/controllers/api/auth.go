package api

import (
	"crm/app/http/request"
	"crm/app/http/request/auth"
	"crm/app/jobs"
	"crm/app/response"
	"crm/app/services"
	"crm/model"
	"crm/pkg/redis"
	"crm/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthController struct {
}

// UserRegister 用户注册
func (a *AuthController) UserRegister(c *gin.Context)  {
	//验证表单请求
	var form auth.UserRegisterRequest

	var response response.Response

	if err := c.ShouldBindJSON(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	var userModel model.User

	parameters := map[string]string{
		"username" : form.Username,
		"password" : form.Password,
		"email"    : form.Email,
		"name"     : form.Name,
		"phone"    : form.Phone,
	}

	//创建数据
	user,err :=userModel.CreateUser(parameters)

	if err != nil || user.Id == 0 {
		response.SetMessage("创建失败，请联系管理员").ResponseError(c)

		return
	}

	//发送邮件信息
	a.sendRegisterEmail(parameters["email"])

	response.SetMessage("success").SetData(map[string]interface{}{
		"userId" : int(user.Id),
		"email"  : user.Email,
	}).ResponseSuccess(c)
}

// 发送注册验证邮箱
func (a *AuthController) sendRegisterEmail(email string) {
	code := utils.GenerateCode()

	key := "register_code_"+email

	redis.Client.Set(key,code,time.Minute*30)

	htmlData := struct {
		Code string
		Type string
	}{
		Code: code,
		Type: "注册验证",
	}

	//解析html数据
	body := utils.AnalysisHtml("resources/emails/register_code.html", htmlData)

	//采用异步非阻塞进行发送
	go func() {
		jobs.Production(map[string]string{
			"email"   : email,
			"body"    : body,
			"subject" : "注册验证码",
		},"email")
	}()
}

// CheckAccount 验证账号信息
func (a *AuthController) CheckAccount(c *gin.Context) {
	//验证表单请求
	var form auth.CheckAccountRequest

	var response response.Response

	if err := c.ShouldBindQuery(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	//验证email/phone/username
	var userModel model.User

	status := true

	switch form.Column {
	case "email":
		status = userModel.GetUserStatusByEmail(form.Value)
		break
	case "username":
		status = userModel.GetUserStatusByUsername(form.Value)
		break
	default:
		status = userModel.GetUserStatusByPhone(form.Value)
	}

	if !status {
		response.SetMessage(form.Column + " 已存在").ResponseError(c)

		return
	}

	response.ResponseSuccess(c)
}

// RegisterCheck 注册状态码验证
func (a *AuthController) RegisterCheck(c *gin.Context)  {
	//验证表单请求
	var form auth.RegisterCheckRequest

	var response response.Response

	if err := c.ShouldBindJSON(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	//code验证
	key := "register_code_"+form.Email

	value := redis.Client.Get(key)

	if value.Val() != form.Code {
		response.SetMessage("验证码异常").ResponseSuccess(c)

		return
	}

	var userModel model.User

	if _,err := userModel.UpdateUserStatus(form.UserId); err != nil {
		response.SetMessage("验证失败，请联系管理员").ResponseSuccess(c)

		return
	}

	//删除key
	redis.Client.Del(key)

	response.SetMessage("注册成功").ResponseSuccess(c)
}

// AccountPasswordLogin 采用账号密码登录
func (a *AuthController) AccountPasswordLogin(c *gin.Context) {
	//验证表单请求
	var form auth.AccountPasswordLoginRequest

	var response response.Response

	if err := c.ShouldBindJSON(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	var userModel model.User

	parameters := map[string]string{
		"username" : form.Username,
		"password" : form.Password,
	}

	user,_ := userModel.GetUserByUsernameAndPassword(parameters)

	if user.Id == 0 {
		response.SetMessage("用户名或密码错误").ResponseError(c)

		return
	}

	if user.Status == 0 {
		response.SetMessage("账号未进行安全认证").ResponseError(c)

		return
	}

	//加载jwt
	var service services.Jwt

	token,err := service.GenerateToken(user.Id,"platform")

	if err != nil {
		response.SetMessage("登录异常，请联系管理员").ResponseError(c)

		return
	}

	//登录成功
	result := map[string]string{
		"token":token,
	}

	response.SetMessage("登录成功").SetData(result).ResponseSuccess(c)
}

// GetCode 获取到登录验证码
func (a *AuthController) GetCode(c *gin.Context) {
	//验证表单请求
	var form auth.GetCodeRequest

	var response response.Response

	if err := c.ShouldBindQuery(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	var userModel model.User

	user,_ := userModel.GetUserByPhone(form.Phone)

	if user.Id == 0 {
		response.SetMessage("手机号未注册").SetCode(422).ResponseError(c)

		return
	}

	code := utils.GenerateCode()
	key := "code_"+form.Phone

	redis.Client.Set(key,code,time.Minute*5)

	//发送邮件验证码
	a.sendPhoneEmail(user.Phone,user.Email)

	response.SetMessage("验证码已发送").ResponseSuccess(c)
}

//发送登录验证码
func (a *AuthController) sendPhoneEmail(phone,email string)  {
	code := utils.GenerateCode()

	key := "login_code_"+phone

	redis.Client.Set(key,code,time.Minute*30)

	htmlData := struct {
		Code string
		Type string
	}{
		Code: code,
		Type: "登录验证",
	}

	//解析html数据
	body := utils.AnalysisHtml("resources/emails/login_phone_code.html", htmlData)

	//采用异步发送邮件
	go func() {
		jobs.Production(map[string]string{
			"email"   : email,
			"body"    : body,
			"subject" : "登录验证码",
		},"email")
	}()
}

// PhoneLogin 采用手机号和验证码登录
func (a *AuthController) PhoneLogin(c *gin.Context) {
	//验证表单请求
	var form auth.PhoneLoginRequest

	var response response.Response

	if err := c.ShouldBindJSON(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	//验证短信验证码
	key := "code_"+form.Phone
	value := redis.Client.Get(key)

	if value.Val() != form.Code {
		response.SetMessage("验证码异常").ResponseError(c)

		return
	}

	var userModel model.User

	user,_ := userModel.GetUserByPhone(form.Phone)

	if user.Id == 0 {
		response.SetMessage("验证异常，请联系管理员").ResponseError(c)

		return
	}

	if user.Status == 0 {
		response.SetMessage("账号未进行安全认证").ResponseError(c)

		return
	}

	var service services.Jwt

	token,err := service.GenerateToken(user.Id,"platform")

	if err != nil {
		response.SetMessage("验证异常，请联系管理员").ResponseError(c)

		return
	}

	//删除key
	redis.Client.Del(key)

	response.SetData(map[string]interface{}{
		"token":token,
	}).ResponseSuccess(c)
}

// Logout 安全退出
func (a *AuthController) Logout(c *gin.Context) {
	var response response.Response

	response.SetMessage("安全退出").ResponseSuccess(c)
}
