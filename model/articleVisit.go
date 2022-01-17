package model

import (
	"crm/pkg/model"
	"time"
)

type ArticleVisit struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"user_id"`
	ArticleId  uint64 `json:"article_id"`
	Ip         string `json:"ip"`
	Model
}

func (a *ArticleVisit) TableName() string {
	return "article_visit"
}

// GetStatusByIdAndIP 判断是否存在记录通过id和ip
func (a *ArticleVisit) GetStatusByIdAndIP(parameters map[string]interface{}) int64 {
	var articleVisit ArticleVisit

	var total int64

	model.Database.Model(&articleVisit).Where("id = ? AND ip = ? AND DATE(created_at) = ?",parameters["id"],parameters["ip"],time.Now().Format("2006-01-02")).Count(&total)

	return total
}

// GetStatusByIdAndUserId 判断是否存在记录通过id和userId
func (a *ArticleVisit) GetStatusByIdAndUserId(parameters map[string]int) int64 {
	var articleVisit ArticleVisit

	var total int64

	model.Database.Model(&articleVisit).Where("id = ? AND user_id = ? AND DATE(created_at) = ?",parameters["id"],parameters["userId"],time.Now().Format("2006-01-02")).Count(&total)

	return total
}

// CreateArticleRecord 创建访问记录
func (a *ArticleVisit) CreateArticleRecord(parameters map[string]interface{}) ArticleVisit  {
	var articleVisit ArticleVisit

	model.Database.Model(&articleVisit).Create(parameters)

	return articleVisit
}
