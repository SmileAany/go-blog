package category

import "crm/app/http/request"

type CreateCategoryRequest struct {
	Name string `form:"name" binding:"required"`
	Describe string `form:"describe" binding:"required"`
}

func (createCategoryRequest CreateCategoryRequest) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Name.required": "请输入分类名称",
		"Describe.required": "请输入分类描述",
	}
}

type UpdateCategoryRequest struct {
	Id uint64 `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
	Describe string `form:"describe" binding:"required"`
}

func (updateCategoryRequest UpdateCategoryRequest) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Id.required": "请输入分类ID",
		"Name.required": "请输入分类名称",
		"Describe.required": "请输入分类描述",
	}
}