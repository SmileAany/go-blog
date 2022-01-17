package api

import (
	"crm/app/http/request"
	category2 "crm/app/http/request/category"
	"crm/app/response"
	"crm/model"
	"crm/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryController struct {

}

// GetCategory 获取到所有分类
func (category *CategoryController) GetCategory(c *gin.Context)  {
	var categoryModel model.Category

	categories := categoryModel.GetAllByUserIdCategory()

	var result = make([]map[string]string,len(categories))

	for i,v := range categories {
		result[i] = map[string]string{
			"name" : v.Name,
		}
	}

	utils.GetSuccessResponse(c,&utils.Response{
		Message: "success",
		Data:    result,
		Status:  "success",
		Code:    200,
		Errors:  []string{},
	})
}

// CreateCategory 新增分类
func (category *CategoryController) CreateCategory(c *gin.Context) {
	//验证表单请求
	var form category2.CreateCategoryRequest

	var response response.Response

	if err := c.ShouldBindJSON(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form,err)).SetCode(422).ResponseError(c)

		return
	}

	var categoryModel model.Category

	//验证分类名称是否已存在
	if categoryModel.GetCategoryByCategoryName(form.Name) > 0 {
		response.SetMessage("分类名称已存在").ResponseError(c)

		return
	}

	data := map[string]string{
		"name" : form.Name,
		"describe" : form.Describe,
	}

	if !categoryModel.CreateCategory(data) {
		response.SetMessage("创建失败").ResponseError(c)

		return
	}

	response.SetMessage("创建成功").ResponseSuccess(c)
}

// UpdateCategory 编辑分类
func (category *CategoryController) UpdateCategory(c *gin.Context) {
	//验证表单请求
	var form category2.UpdateCategoryRequest

	var response response.Response

	if err := c.ShouldBindJSON(&form); err != nil {
		response.SetMessage(request.GetErrorMessage(form, err)).SetCode(422).ResponseError(c)

		return
	}

	var categoryModel model.Category

	//验证分类id是否存在
	if !categoryModel.GetCategoryStatusByCategoryId(form.Id) {
		response.SetMessage("分类Id 不存在").ResponseError(c)

		return
	}

	//验证分类名称是否已存在
	if categoryModel.GetCategoryByCategoryName(form.Name) > 1 {
		response.SetMessage("分类名称已存在").ResponseError(c)

		return
	}

	data := map[string]string{
		"id" : strconv.Itoa(int(form.Id)),
		"name" : form.Name,
		"describe" : form.Describe,
	}

	if !categoryModel.UpdateCategory(data) {
		response.SetMessage("编辑失败").ResponseError(c)

		return
	}

	response.SetMessage("编辑成功").ResponseSuccess(c)
}




