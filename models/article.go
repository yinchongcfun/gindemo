package models

import (
	"gindemo/database"
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category";gorm:"foreignkey:CategoryID"`
	Tag        []Tag    `gorm:"many2many:article_tag" json:"tag"`
}

func (Article) TableName() string {
	return "article"
}

func (a *Article) ListArticle1(title string) (articles []Article, err error) {
	query := database.GormPool
	err = query.Model(articles).
		Where("title like ?", "%"+title+"%").
		Preload("Category").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

//func ListArticle2(title string) (articles []Article, err error) {
//	query := database.GormPool
//	err = query.Model(articles).
//		Where("title like ?", "%"+title+"%").
//		Preload("Category").Find(&articles).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return
//	}
//	return
//}

func (a *Article) ListArticle(title string) (Article, error) {
	query := database.GormPool
	var article Article
	query.Where("title like ?", "%"+title+"%").First(&article)
	err := query.Model(&article).Related(&article.Tag, "Tag").Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, nil
	}
	return article, err
}

//func (a *Article) ListArticle3(title string) (Article, error) {
//	query := database.GormPool
//	var article Article
//	query.Where("title like ?", "%"+title+"%").First(&article)
//	err := query.Model(&article).Association("Category").Find(&article.Category).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return article, nil
//	}
//	return article, err
//}

////远程一对多
//func (a *Article) ListArticle4(title string) (Article, error) {
//	query := database.GormPool
//	var article Article
//	query.Where("title like ?", "%"+title+"%").First(&article)
//	err := query.Model(&article).
//		Related(&article.Category).
//		Related(&article.Tag, "tag").
//		Find(&article).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return article, nil
//	}
//	return article, err
//}

//多对多
//func (a *Article) ListArticle(title string) (articles []Article, err error) {
//	query := database.GormPool
//	err = query.Model(articles).
//		Where("title like ?", "%"+title+"%").
//		Preload("Category").
//		Preload("Tag").Find(&articles).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return
//	}
//	return
//}
