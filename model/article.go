package model

import (
	"crm/pkg/model"
	"time"
)

type Article struct {
	Id       uint64 `json:"id"`
	UserId   uint64 `json:"userId"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Content  string `json:"content"`
	CategoryId uint64 `json:"categoryId"`
	Status   byte `json:"status"`
	User      User `json:"user" gorm:"foreignKey:UserId"`
	Category  Category `json:"user" gorm:"foreignKey:CategoryId"`
	ArticleVisit []ArticleVisit `json:"article_visit" gorm:"foreignKey:ArticleId"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (a *Article) TableName() string {
	return "articles"
}

// CreateArticle 创建文章
func (a *Article) CreateArticle(article Article) (Article,error) {
	if err := model.Database.Create(&article).Error; err != nil {
		return article,err
	}

	return article, nil
}

// GetArticles 获取到所有文章数据
func (a *Article) GetArticles(data map[string]int) ([]Article, int64,error) {
	var articles []Article

	offset := (data["page"] -1) * 10

	var total int64

	model.Database.Model(&articles).Where("status = ? AND deleted_at IS NULL ",1).Count(&total)

	if err := model.Database.Where("status = ? AND deleted_at IS NULL ",1).Preload("User").Offset(offset).Limit(10).Find(&articles).Error; err != nil {
		return articles, 0,err
	}

	return articles, total,nil
}

// GetArticleDetails 获取到文章详情
func (a *Article) GetArticleDetails(articleId int) (Article,error) {
	var article Article

	if err := model.Database.Where("id = ? ",articleId).Preload("User").Preload("Category").Preload("ArticleVisit").First(&article).Error; err != nil {
		return article,err
	}

	return article,nil
}

// UpdateArticleDetails 编辑文章
func (a *Article) UpdateArticleDetails(article Article) (Article,error)  {
	if err := model.Database.Model(&article).Updates(article).Error; err != nil{
		return article,err
	}

	return article,nil
}