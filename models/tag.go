package models

type Tag struct {
	Id        int    `json:"id" `
	TagName string `json:"tag_name"`
}

func (Tag) TableName() string {
	return "tag"
}