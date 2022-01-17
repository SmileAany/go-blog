package model

import (
	"crm/pkg/model"
	"fmt"
	"strconv"
)

type Category struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Describe string `json:"describe"`
	Status uint `json:"status"`
}

// TableName 声明表名
func (c *Category) TableName() string {
	return "category"
}

// GetAllByUserIdCategory 获取到所有分类
func (c *Category) GetAllByUserIdCategory() []Category {
	var category []Category

	model.Database.Select("name").Where("status = ?",1).Find(&category)

	return category
}

// GetCategoryByCategoryName 根据分类名称判断是否存在相同的分类名称
func (c *Category) GetCategoryByCategoryName(categoryName string) int {
	var category []Category

	model.Database.Where("name = ?",categoryName).Select("id").Find(&category)

	return len(category)
}

// CreateCategory 创建分类
func (c *Category) CreateCategory(parameters map[string]string) bool{
	category := Category{
		Name: parameters["name"],
		Describe: parameters["describe"],
		Status: uint(1),
	}

	model.Database.Create(&category)

	if category.Id > 0 {
		return true
	}

	return false
}

// UpdateCategory 编辑分类
func (c *Category) UpdateCategory(parameters map[string]string) bool  {
	id,_ := strconv.Atoi(parameters["id"])

	category := Category{
		Name: parameters["name"],
		Describe: parameters["describe"],
	}

	result := model.Database.Where("id = ?",id).Updates(category)

	if result.Error != nil {
		return false
	}

	return true
}

// GetCategoryStatusByCategoryId 根据分类id 获取到分类详情
func (c *Category) GetCategoryStatusByCategoryId (categoryId uint64) bool {
	var category Category

	model.Database.Select("id").Where("id = ?",categoryId).First(&category)

	fmt.Println(category)

	if category.Id > 0 {
		return true
	}

	return false
}


