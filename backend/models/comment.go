package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model        //包含评论id、评论时间
	Title      string `db:"title" json:"title" binding:"required"`     // 属于哪条新闻
	Author     string `db:"author" json:"author" binding:"required"`   //评论者名字
	Content    string `db:"content" json:"content" binding:"required"` // 评论内容
	Type       int    `db:"type"`                                      //评论类型如果为0表示是一级评论，1为二级评论
	ParentId   uint   `db:"parentId"`                                  //若type为0，则ParentId为自己的id，若type=1，则为父评论的id
	Likes      int    `db:"likes" gorm:"default:0"`                    //点赞数，初始为0，点赞后加1
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
}

// 发布子评论
type ChildComment struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ParentId string `json:"parentId" binding:"required"`
}
