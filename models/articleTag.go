package models

type ArticleTag struct {
	Id        int    `json:"id" `
	ArticleId string `json:"article_id"`
	TagId     string `json:"tag_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (ArticleTag) TableName() string {
	return "article_tag"
}