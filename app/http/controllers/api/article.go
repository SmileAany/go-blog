package api

import (
	"crm/app/http/request"
	"crm/app/http/request/article"
	"crm/app/response"
	"crm/app/services"
	"crm/model"
	"crm/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleController struct {

}

// CreateArticle 创建文章
func (a *ArticleController) CreateArticle(c *gin.Context)  {
	//验证表单请求
	var form article.CreateArticleRequest

	var response response.Response

	if err := c.ShouldBindJSON(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	service := services.Jwt{}

	claims, _ := service.ParseToken(c.GetHeader("Authorization"))

	timeDate := time.Now()

	//默认截取200字符，
	content := []rune(utils.TrimHtml(form.Content))

	var abstract []rune

	for i := 0; i < len(content) / 5; i++ {
		abstract = append(abstract,content[i])
	}

	var articleModel model.Article = model.Article{
		Title:   form.Title,
		Content: form.Content,
		Abstract: string(abstract),
		Status: byte(form.Status),
		CategoryId : uint64(form.CategoryId),
		UserId: claims.UserId,
		CreatedAt: &timeDate,
		UpdatedAt: &timeDate,
		DeletedAt: nil,
	}

	article,err := articleModel.CreateArticle(articleModel)

	if err != nil || article.Id == 0 {
		response.SetMessage("文章创建失败，请联系管理员").ResponseError(c)

		return
	}

	response.SetMessage("创建成功").ResponseSuccess(c)
}

// DeleteArticle 删除文章
func (a *ArticleController) DeleteArticle(c *gin.Context) {
	//验证表单请求
	var form article.DeleteArticleRequest

	var response response.Response

	if err := c.ShouldBindQuery(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	//先判断传递的id是否是有效
	var articleModel model.Article

	article,_ := articleModel.GetArticleDetails(form.Id)

	if article.Id == 0 {
		response.SetMessage("文章id 异常").SetCode(400).ResponseError(c)

		return
	}

	service := services.Jwt{}

	claims, _ := service.ParseToken(c.GetHeader("Authorization"))

	if article.UserId != claims.UserId {
		response.SetMessage("操作越权").SetCode(400).ResponseError(c)

		return
	}

	timeDate := time.Now()

	article.DeletedAt = &timeDate

	if article,_ := articleModel.UpdateArticleDetails(article); article.DeletedAt == nil {
		response.SetMessage("删除失败，请联系管理员").SetCode(400).ResponseError(c)

		return
	}

	response.SetMessage("删除成功").ResponseSuccess(c)
}

// UpdateArticle 文章编辑
func (a *ArticleController) UpdateArticle(c *gin.Context) {
	//验证表单请求
	var form article.UpdateArticleRequest

	var response response.Response

	if err := c.ShouldBindQuery(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	//先判断传递的id是否是有效
	var articleModel model.Article

	article,_ := articleModel.GetArticleDetails(form.Id)

	if article.Id == 0 {
		response.SetMessage("文章id 异常").SetCode(400).ResponseError(c)

		return
	}

	service := services.Jwt{}

	claims, _ := service.ParseToken(c.GetHeader("Authorization"))

	if article.UserId != claims.UserId {
		response.SetMessage("操作越权").SetCode(400).ResponseError(c)

		return
	}

	//默认截取200字符，
	content := []rune(utils.TrimHtml(form.Content))

	var abstract []rune

	for i := 0; i < len(content) / 5; i++ {
		abstract = append(abstract,content[i])
	}

	timeDate := time.Now()

	article = model.Article{
		Title:   form.Title,
		Content: form.Content,
		Abstract: string(abstract),
		Status: byte(form.Status),
		CategoryId : uint64(form.CategoryId),
		UserId: claims.UserId,
		UpdatedAt: &timeDate,
	}

	if article,_ := articleModel.UpdateArticleDetails(article); *(article.UpdatedAt) != timeDate {
		response.SetMessage("编辑失败，请联系管理员").SetCode(400).ResponseError(c)

		return
	}

	response.SetMessage("编辑成功").ResponseSuccess(c)
}

// GetList 获取到文章列表
func (a *ArticleController) GetList(c *gin.Context) {
	//验证表单请求
	var form article.GetListRequest

	var response response.Response

	if err := c.ShouldBindQuery(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	var articleModel model.Article

	data := map[string]int{
		"page": form.Page,
	}

	articles,total,_ := articleModel.GetArticles(data)

	//转换数据格式
	var result = make([]map[string]interface{},len(articles))

	for index,article := range articles {
		result[index] = map[string]interface{}{
			"id" : article.Id,
			"title" : article.Title,
			"abstract" : article.Abstract,
			"content" : article.Content,
			"auth" :  article.User.Username,
			"create_date" : article.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	response.SetData(map[string]interface{}{
		"total": total,
		"page":  form.Page,
		"list":  result,
	}).ResponseSuccess(c)
}

// RecordVisit 统计访问
func (a *ArticleController) RecordVisit(c *gin.Context,articleId int) {
	token := c.GetHeader("Authorization")
	service := &services.Jwt{}
	ip := c.ClientIP()

	var articleVisitModel model.ArticleVisit

	if claim,err := service.ParseToken(token); err != nil{
		status := articleVisitModel.GetStatusByIdAndIP(map[string]interface{}{
			"id" : articleId,
			"ip" : ip,
		})

		if status == 0 {
			articleVisitModel.CreateArticleRecord(map[string]interface{}{
				"article_id" : articleId,
				"ip"      : ip,
			})
		}
	} else {
		userId := claim.UserId

		status := articleVisitModel.GetStatusByIdAndUserId(map[string]int{
			"id" : articleId,
			"userId" : int(userId),
		})

		if status == 0 {
			articleVisitModel.CreateArticleRecord(map[string]interface{}{
				"article_id" : articleId,
				"user_id" : userId,
				"ip"      : ip,
			})
		}
	}
}

// GetDetails 查看文章详情
func (a *ArticleController) GetDetails(c *gin.Context) {
	//验证表单请求
	var form article.GetDetailsRequest

	var response response.Response

	if err := c.ShouldBindQuery(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	var articleModel model.Article

	article,_ := articleModel.GetArticleDetails(form.Id)

	if article.Id == 0 {
		response.SetMessage("文章id 异常").SetCode(400).ResponseError(c)

		return
	}

	response.SetData(map[string]interface{}{
		"content"     : article.Content,
		"create_date" : article.CreatedAt.Format("2006-01-02 15:04:05"),
		"auth"        : article.User.Username,
		"visit_count" : len(article.ArticleVisit),
		"category"    : article.Category.Name,
	}).ResponseSuccess(c)
}