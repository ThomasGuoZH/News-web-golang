package models

type Comment struct {
	NewsTitle string     // 属于哪条新闻
	author    string     //评论者名字
	Id        string     //评论id（随机生成）
	Content   string     // 评论内容
	ParentId  int        // 父评论ID，如果为0表示是一级评论，1为二级评论
	Children  []*Comment // 子评论列表
}

// 数据表名字
func (Comment) TableName() string {
	return "comment"
}
