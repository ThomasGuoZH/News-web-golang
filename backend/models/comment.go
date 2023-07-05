package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model        //包含评论id、评论时间
	Title      string `db:"title"` // 属于哪条新闻
	Author     string `db:"author"` //评论者名字
	Content    string `db:"content"` // 评论内容
	Type       int    `db:"type"` //评论类型如果为0表示是一级评论，1为二级评论
	ParentId   uint   `db:"parent_id"` //若type为0，则ParentId为自己的id，若type=1，则为父评论的id
	Likes      int    `db:"likes" gorm:"default:0"`
	Channel    string `db:"channel"`
}

// 数据表名字
func (Comment) TableName() string {
	return "comment"
}

// 发布父评论
type ParentComment struct {
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Content string `json:"content" binding:"required"`
	Channel string `json:"channel" binding:"required"`
}

// 发布子评论
type ChildComment struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ParentId string `json:"parentId" binding:"required"`
	Channel  string `json:"channel" binding:"required"`
}
