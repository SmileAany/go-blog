package article

import "crm/app/http/request"

type CreateArticleRequest struct {
	Title string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
	CategoryId int `form:"categoryId"`
	Status int `form:"status" binding:"required"`
}

func (createArticleRequest CreateArticleRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Title.required": "请输入标题",
		"Content.required": "请编辑内容",
		"CategoryId.required": "请传递分类ID",
		"Status.required": "请传递状态值",
	}
}

type DeleteArticleRequest struct {
	Id int `form:"id" binding:"required"`
}

func (deleteArticleRequest DeleteArticleRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Id.required": "请传递文章id",
	}
}

type UpdateArticleRequest struct {
	Id int `form:"id" binding:"required"`
	Title string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
	CategoryId int `form:"categoryId"`
	Status int `form:"status" binding:"required"`
}

func (updateArticleRequest UpdateArticleRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Id.required": "请传递文章id",
		"Title.required": "请输入标题",
		"Content.required": "请编辑内容",
		"CategoryId.required": "请传递分类ID",
		"Status.required": "请传递状态值",
	}
}

type GetListRequest struct {
	Page int `form:"page" binding:"required"`
}

func (getListRequest GetListRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Page.required": "请传入分页page",
	}
}

type GetDetailsRequest struct {
	Id int `form:"id" binding:"required"`
}

func (getDetailsRequest GetDetailsRequest) GetMessages()  request.ValidatorMessages {
	return request.ValidatorMessages{
		"Page.required": "请传入文章id",
	}
}