package models

import "github.com/jinzhu/gorm"

type Likes struct {
	gorm.Model
	Title     string `db:"title" json:"title"`
	Liker     string `db:"username" json:"liker"`
	Author    string `db:"author" json:"author"`
	Content   string `db:"content" json:"content"`
	CommentId string `db:"commentId" json:"commentId"`
}

// 数据表名字
func (Likes) TableName() string {
	return "likes"
}
