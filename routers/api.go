package routers

import (
	"crm/app/http/controllers/api"
	"crm/app/http/middleware/jwt"
	"crm/app/http/middleware/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//全局中间件
	r.Use(logger.Logger())

	//用户
	router := r.Group("api/auth")
	{
		controller := new(api.AuthController)

		//用户注册
		router.POST("/user/register",controller.UserRegister)

		//注册账号/邮箱/手机号验证
		router.GET("/register/account/check",controller.CheckAccount)

		//注册验证
		router.POST("/user/register/check",controller.RegisterCheck)

		//用户账号密码登录
		router.POST("/login/account", controller.AccountPasswordLogin)

		//获取到登录验证码
		router.GET("/phone/code",controller.GetCode)

		//用户手机验证码登录
		router.POST("/login/phone", controller.PhoneLogin)

		//用户退出登录
		router.Use(jwt.Jwt())
		{
			router.GET("/logout",controller.Logout)
		}
	}

	//文章
	router = r.Group("api/article")
	{
		controller := new(api.ArticleController)

		//文章列表
		router.GET("/list", controller.GetList)

		//文章详情
		router.GET("/details", controller.GetDetails)

		router.Use(jwt.Jwt())
		{
			//创建文件
			router.POST("/create",controller.CreateArticle)

			//删除文章
			router.DELETE("/delete",controller.DeleteArticle)

			//编辑文章
			router.PUT("/edit",controller.DeleteArticle)
		}
	}

	router = r.Group("api/category")
	{
		controller := new(api.CategoryController)

		//所有分类
		router.GET("/list", controller.GetCategory)
		//新增分类
		router.POST("/create",controller.CreateCategory)
		//编辑分类
		router.PUT("/update",controller.UpdateCategory)
		//分类详情
		//router.GET("/details",controller.CategoryDetails)
	}

	return r
}
